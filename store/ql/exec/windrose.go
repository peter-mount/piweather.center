package exec

import (
	"github.com/alecthomas/participle/v2"
	ql2 "github.com/peter-mount/piweather.center/config/ql"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/ql/functions"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
	"strconv"
)

func windRose(v ql2.Visitor[*Executor], s *ql2.WindRose) error {
	ex := v.Get()

	wr := api.NewWindRose()

	it := ex.timeRange.Iterator()
	for it.HasNext() {
		ex.time = it.Next()

		degrees, b0, err := ex.windRoseExpression(v, s.Degrees, measurement.Degree)
		if err != nil {
			return err
		}

		speed, b1, err := ex.windRoseExpression(v, s.Speed, measurement.MetersPerSecond)
		if err != nil {
			return err
		}

		if b0 && b1 {
			wr.Add(degrees, speed)
		}
	}

	// Ensure all statistics are correct
	wr.Finalise()

	for _, opt := range s.Options {
		switch {
		case opt.Rose:
			ex.result.AddWindRose(wr)

		case opt.Count:
			ex.windRoseTable(wr, func(bucket *api.WindRoseBucket) float64 {
				return float64(bucket.Count)
			})

		case opt.Max:
			ex.windRoseTable(wr, func(bucket *api.WindRoseBucket) float64 {
				return bucket.Max
			})

		}

	}

	// Tell the visitor to stop processing this Histogram statement
	return util.VisitorStop
}

func (ex *Executor) windRoseTable(wr *api.WindRose, f func(*api.WindRoseBucket) float64) {
	t := ex.result.NewTable()
	for i := 0; i < len(wr.Buckets); i++ {
		t.AddColumn(&api.Column{
			Index: i,
			Name:  "b" + strconv.Itoa(i),
			Unit:  "Float",
		})
	}

	r := t.NewRow()
	for i := 0; i < len(wr.Buckets); i++ {
		r.AddValue(ex.time, value.Float.Value(f(wr.Buckets[i])))
	}
}

func (ex *Executor) windRoseExpression(v ql2.Visitor[*Executor], s *ql2.Expression, u *value.Unit) (float64, bool, error) {
	ex.resetStack()
	err := v.Expression(s)

	val, ok := ex.Pop()

	// If invalid but have values attached then get the last value in the Set.
	// Required with metrics without an aggregation function around them
	if !val.IsTime && !val.Value.IsValid() && len(val.Values) > 0 {
		val = functions.InitialLast(val)
	}

	var f float64
	var b bool
	if err == nil && ok && !val.IsNull() {
		if val.IsTime {
			err = participle.Errorf(s.Pos, "time not supported here")
		} else {
			val1, err := val.Value.As(u)
			b = err == nil
			f = val1.Float()
		}
	}

	return f, b, err
}
