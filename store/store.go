package store

import (
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/cron"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/station/payload"
	"github.com/peter-mount/piweather.center/station/service"
	"github.com/peter-mount/piweather.center/store/file"
	"github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/store/memory"
	"github.com/peter-mount/piweather.center/util"
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
	ProcessReading(ctx context.Context) error
	Calculate(ctx context.Context) error
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

func (s *store) ProcessReading(ctx context.Context) error {
	r := station.ReadingFromContext(ctx)
	values := value.MapFromContext(ctx)
	if r.Unit() != nil {
		p := payload.GetPayload(ctx)

		str, ok := p.Get(r.Source)
		if !ok {
			// FIXME warn/fail if not found?
			return nil
		}

		if f, ok := util.ToFloat64(str); ok {
			// Convert to Type unit then transform to Use unit
			v, err := r.Value(f)
			if err != nil {
				// Ignore, should only happen if the result is
				// invalid as we checked the transform previously
				return nil
			}

			values.Put(r.ID, v)
		}
	}
	return nil
}

func (s *store) Calculate(ctx context.Context) error {
	// Get value.Time from Station and Payload
	sensors := station.SensorsFromContext(ctx)
	p := payload.GetPayload(ctx)
	t := sensors.Station().LatLong().Time(p.Time())

	calc := station.CalculatedValueFromContext(ctx)

	values := value.MapFromContext(ctx)
	args := values.GetAll(calc.Source...)

	result, err := calc.Calculate(t, args...)
	if err != nil {
		return err
	}

	values.Put(calc.ID, result)

	return nil
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
