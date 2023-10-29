package store

import (
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/cron"
	"github.com/peter-mount/piweather.center/station/service"
	"github.com/peter-mount/piweather.center/store/file"
	"github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/store/memory"
	"github.com/peter-mount/piweather.center/weather/value"
	"golang.org/x/net/context"
	"sort"
	"strings"
	"time"
)

func init() {
	kernel.RegisterAPI((*Store)(nil), &store{})
}

type Store interface {
	AddContext(ctx context.Context) context.Context
	Record(name string, value value.Value, recTime time.Time)
	Latest(name string) (record.Record, bool)
	Metrics() []string
	GetMetricBetween(metric string, from, to time.Time) []record.Record
	GetHistory(metric string) []record.Record
}

type store struct {
	Cron   *cron.CronService `kernel:"inject"`
	Config service.Config    `kernel:"inject"`
	Memory memory.Latest     `kernel:"inject"`
	File   file.Store        `kernel:"inject"`
}

func StoreFromContext(ctx context.Context) Store {
	return ctx.Value("local.store").(Store)
}

func (s *store) AddContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "local.store", s)
}

func (s *store) Record(name string, value value.Value, recTime time.Time) {
	name = strings.ToLower(name)

	rec := record.Record{
		Time:  recTime,
		Value: value,
	}

	_ = s.File.Append(name, rec)
	s.Memory.Append(name, rec)
}

func (s *store) Latest(name string) (record.Record, bool) {
	return s.Memory.Latest(name)
}

func (s *store) Metrics() []string {
	return s.Memory.Metrics()
}

func (s *store) GetMetrics(query file.Query) []record.Record {
	records := file.GetAllRecords(query)
	sort.SliceStable(records, func(i, j int) bool {
		return records[i].Time.Before(records[j].Time)
	})
	return records
}

func (s *store) GetMetricBetween(metric string, from, to time.Time) []record.Record {
	return s.GetMetrics(s.File.Query(metric).
		Between(from, to).
		Build())
}

func (s *store) GetHistory(metric string) []record.Record {
	return s.GetMetrics(s.File.Query(metric).
		Today().
		Build())
}
