package exec

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/piweather.center/store/ql/functions"
	"github.com/peter-mount/piweather.center/store/ql/lang"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
)

func (ex *Executor) windRose(v lang.Visitor, s *lang.WindRose) error {
	wr := ex.result.NewWindRose()

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

	// Tell the visitor to stop processing this Histogram statement
	return lang.VisitorStop
}

func (ex *Executor) windRoseExpression(v lang.Visitor, s *lang.Expression, u *value.Unit) (float64, bool, error) {
	ex.resetStack()
	err := v.Expression(s)

	val, ok := ex.Pop()

	// If invalid but have values attached then get the last value in the set.
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
