package file

import (
	"github.com/peter-mount/piweather.center/store/file/record"
	"sync"
)

// Latest manages an in memory copy of the most recent Record entered into the database
type Latest struct {
	mutex   sync.Mutex
	metrics map[string]record.Record
}

func (l *Latest) Start() error {
	l.metrics = make(map[string]record.Record)
	return nil
}

func (l *Latest) Append(metric string, rec record.Record) {
	if metric == "" || !rec.IsValid() {
		return
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()

	// Check that existing entry is not newer than the one being appended
	old, exists := l.metrics[metric]
	if exists && old.IsValid() && old.Time.After(rec.Time) {
		return
	}

	l.metrics[metric] = rec
}

func (l *Latest) Latest(metric string) (record.Record, bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	rec, exists := l.metrics[metric]
	return rec, exists
}

func (l *Latest) Metrics() []string {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	var metrics []string

	for metric, _ := range l.metrics {
		metrics = append(metrics, metric)
	}

	return metrics
}
