package ql

import (
	"github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

type Value struct {
	Time   time.Time
	Value  value.Value
	Values []Value
	IsTime bool
}

func (v Value) IsNull() bool {
	return !(v.IsTime || v.Value.IsValid())
}

func FromRecord(r record.Record) Value {
	return Value{
		Time:  r.Time,
		Value: r.Value,
	}
}

type Executor interface {
	Time() time.Time
	Push(v Value)
	Pop() (Value, bool)
}
