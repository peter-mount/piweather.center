package memory

import (
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/weather/value"
	"sync"
	"time"
)

func init() {
	kernel.RegisterAPI((*Latest)(nil), &latest{})
}

// Latest manages an in memory copy of the most recent Record entered into the database
type Latest interface {
	Append(metric string, rec record.Record) bool
	Latest(metric string) (record.Record, bool)
	Metrics() []string
	LatestTime() time.Time
}

type latest struct {
	mutex      sync.Mutex
	metrics    map[string]record.Record
	latestTime time.Time
}

func (l *latest) Start() error {
	l.metrics = make(map[string]record.Record)
	return nil
}

func (l *latest) Append(metric string, rec record.Record) bool {
	if metric == "" || !rec.IsValid() {
		return false
	}

	// Truncate time to the second as the DB only has resolution to
	// the second
	rec.Time = rec.Time.Truncate(time.Second)

	l.mutex.Lock()
	defer l.mutex.Unlock()

	// Check that existing entry is not newer than the one being appended
	old, exists := l.metrics[metric]
	if exists && old.IsValid() && old.Time.After(rec.Time) {
		return false
	}

	l.metrics[metric] = rec

	// Keep latest time value to most recent timestamp
	if rec.Time.After(l.latestTime) {
		l.latestTime = rec.Time
	}

	// If we have an old value then compare the two.
	// If they do not equal each other, or the transform fails due to change of unit
	// then return true to indicate the value has changed.
	if exists {
		notSame, err := old.Value.Compare(rec.Value, value.NotEqual)
		return err != nil || notSame
	}

	return true
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

func (l *latest) LatestTime() time.Time {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	return l.latestTime
}
