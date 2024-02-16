package exec

import (
	lang2 "github.com/peter-mount/piweather.center/config/ql"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/ql"
)

func (ex *Executor) metric(_ lang2.Visitor, s *lang2.Metric) error {
	r := api.RangeFrom(ex.time, ex.timeRange.Every)
	vals := ex.findMetrics(s.Name, r)

	// No results then Push null
	if len(vals) == 0 {
		ex.Push(ql.Value{})
		return nil
	}

	// Take time from the first result
	ex.Push(ql.Value{
		Time:   vals[0].Time,
		Values: vals,
	})

	return nil
}

func (ex *Executor) findMetrics(n string, times api.Range) []ql.Value {
	var r []ql.Value

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
