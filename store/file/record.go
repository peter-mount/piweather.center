package file

import (
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

type RecordHandler interface {
	// Size returns the fixed record size
	Size() int
	// Read a record
	Read([]byte) (Record, error)
	// Append a record
	Append([]byte, Record) []byte
}

type Record struct {
	Time  time.Time
	Value value.Value
}

func (r Record) IsValid() bool {
	return !r.Time.IsZero() && r.Value.IsValid()
}

func (r Record) Equals(b Record) bool {
	v, err := r.Value.Equals(b.Value)
	return err == nil && v && r.Time.Equal(b.Time)
}

func AssertRecordLength(r RecordHandler, b []byte) error {
	return AssertLength(r.Size(), b)
}
