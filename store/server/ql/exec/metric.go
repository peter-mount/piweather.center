package exec

import (
	"github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/store/server/ql/lang"
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

type Value struct {
	Time   time.Time
	Value  value.Value
	IsTime bool
	IsNull bool
}

func FromRecord(r record.Record) Value {
	return Value{
		Time:   r.Time,
		Value:  r.Value,
		IsNull: !r.IsValid(),
	}
}

func (ex *executor) metric(v lang.Visitor, s *lang.Metric) error {
	recs := ex.metrics[s.Name]

	// No results then push null
	if len(recs) == 0 {
		ex.push(Value{IsNull: true})
		return nil
	}

	// TODO implement - for now return first metric but should get
	ex.push(recs[0])

	return nil
}
