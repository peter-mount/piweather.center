package memory

import (
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/store/file/record"
	"sync"
	"time"
)

func init() {
	kernel.RegisterAPI((*Latest)(nil), &latest{})
}

// Latest manages an in memory copy of the most recent Record entered into the database
type Latest interface {
	Append(metric string, rec record.Record)
	Latest(metric string) (record.Record, bool)
	Metrics() []string
}

type latest struct {
	mutex   sync.Mutex
	metrics map[string]record.Record
}

func (l *latest) Start() error {
	l.metrics = make(map[string]record.Record)
	return nil
}

func (l *latest) Append(metric string, rec record.Record) {
	if metric == "" || !rec.IsValid() {
		return
	}

	// Truncate time to the second as the DB only has resolution to
	// the second
	rec.Time = rec.Time.Truncate(time.Second)

	l.mutex.Lock()
	defer l.mutex.Unlock()

	// Check that existing entry is not newer than the one being appended
	old, exists := l.metrics[metric]
	if exists && old.IsValid() && old.Time.After(rec.Time) {
		return
	}

	l.metrics[metric] = rec
}

func (l *latest) Latest(metric string) (record.Record, bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	rec, exists := l.metrics[metric]
	return rec, exists
}

func (l *latest) Metrics() []string {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	var metrics []string

	for metric, _ := range l.metrics {
		metrics = append(metrics, metric)
	}

	return metrics
}
