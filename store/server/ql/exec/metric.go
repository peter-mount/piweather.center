package exec

import (
	"github.com/peter-mount/go-kernel/v2/log"
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
	IsNull bool
}

func FromRecord(r record.Record) Value {
	return Value{
		Time:   r.Time,
		Value:  r.Value,
		IsNull: !r.IsValid(),
	}
}

func (ex *executor) metric(_ lang.Visitor, s *lang.Metric) error {
	r := lang.RangeFrom(ex.time, ex.timeRange.Every)
	vals := ex.findMetrics(s.Name, r)

	log.Printf("metric %q %d", s.Name, len(vals))
	// No results then push null
	if len(vals) == 0 {
		ex.push(Value{IsNull: true})
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
