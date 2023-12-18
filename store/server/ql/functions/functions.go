package functions

import (
	"github.com/peter-mount/piweather.center/store/server/ql"
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
)

var functions = FunctionMap{
	functions: map[string]Function{
		// average - avg(metric) or avg(metric, ...)
		"avg": {
			MinArg:      1,
			Initial:     InitialInvalid,
			Calculation: value.Add,
			Aggregator: func(l int, a ql.Value) ql.Value {
				if a.Value.IsValid() && l > 0 {
					a.Value = a.Value.Value(a.Value.Float() / float64(l))
				}
				return a
			},
		},
		// ceil returns the least integer value greater than or equal to x
		"ceil": {
			// Here reduce to get the current max value, then aggregate to the ceiling
			MinArg:     1,
			Initial:    InitialLast,
			Aggregator: MathAggregator(math.Ceil),
		},
		// ceil returns the least integer value greater than or equal to x
		// Unlike ceil() this will return the ceil of the highest value in the set.
		// e.g. ceilAll(metric) is the same as ceil(max(metric))
		"ceilall": {
			// Here reduce to get the current max value, then aggregate to the ceiling
			MinArg:     1,
			Reducer:    value.GreaterThan,
			Aggregator: MathAggregator(math.Ceil),
		},
		// count(metric) the number of entries in the metrics set
		"count": {
			Args: 1,
			Aggregator: func(l int, a ql.Value) ql.Value {
				// Return l as a value with the unit of "integer"
				return ql.Value{
					Time:  a.Time,
					Value: value.Integer.Value(float64(l)),
				}
			},
		},
		// first(metric) the first valid entry in the metrics set
		"first": {
			Args:    1,
			Initial: InitialFirst,
			Reducer: value.TrueComparator,
		},
		// floor returns the greatest integer value less than or equal to x
		"floor": {
			// Here reduce to get the current min value, then aggregate to the floor
			MinArg:     1,
			Initial:    InitialLast,
			Aggregator: MathAggregator(math.Floor),
		},
		// floorAll returns the greatest integer value less than or equal to x.
		// Unlike floor() this will return the floor of the lowest value in the set.
		// e.g. floorAll(metric) is the same as floor(min(metric))
		"floorall": {
			// Here reduce to get the current min value, then aggregate to the floor
			MinArg:     1,
			Reducer:    value.LessThan,
			Aggregator: MathAggregator(math.Floor),
		},
		// last(metric) the last valid entry in the metrics set
		"last": {
			Args:    1,
			Initial: InitialLast,
			Reducer: value.FalseComparator,
		},
		// minimum - min(metric) or min(metric, ...)
		"min": {
			MinArg:  1,
			Reducer: value.LessThan,
		},
		// maximum - max(metric) or max(metric, ...)
		"max": {
			MinArg:  1,
			Reducer: value.GreaterThan,
		},
		// sum - sum(metric) or sum(metric, ...)
		"sum": {
			Args:        1,
			Initial:     InitialInvalid,
			Calculation: value.Add,
		},
		// timeof metric
		// timeof() sets the value to be the time of the row
		// timeof(metric) sets the value to that of the metric
		"timeof": {
			MinArg:   0,
			MaxArg:   1,
			Function: funcTimeOf,
		},
		// trend takes two values and returns the trend between them:
		// 0 = steady
		// 1 = rising
		// -1 = falling
		// null = invalid or no data
		"trend": {
			Args:     2,
			Function: funcTrend,
		},
	},
}
