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

func (ex *executor) metric(_ lang.Visitor, s *lang.Metric) error {
	r := lang.RangeFrom(ex.time, ex.timeRange.Every)
	vals := ex.findMetrics(s.Name, r)

	// No results then push null
	if len(vals) == 0 {
		ex.push(Value{})
		return nil
	}

	// Take time from the first result
	ex.push(Value{
		Time:   vals[0].Time,
		Values: vals,
	})

	return nil
}

func (ex *executor) findMetrics(n string, times lang.Range) []Value {
	var r []Value

	recs := ex.metrics[n]

	if len(recs) > 0 {
		for _, e := range recs {
			if times.Contains(e.Time) {
				r = append(r, e)
			}
		}
	}

	return r
}
