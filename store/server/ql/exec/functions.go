package exec

import (
	"fmt"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/store/server/ql/lang"
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
	"strings"
	"sync"
)

func (ex *Executor) function(v lang.Visitor, f *lang.Function) error {
	if af, exists := functions.GetFunction(f.Name); exists {
		return ex.runFunction(v, f, af)
	}
	return fmt.Errorf("unknown function %q", f.Name)
}

type Function struct {
	// Args is the number of arguments a Function requires.
	// This overrides MinArg and MaxArg if MinArg < Args > MaxArg
	Args int

	// MinArg is the minimum number of arguments a Function accepts.
	// Default 0
	MinArg int

	// MaxArg is the maximum number of arguments a Function accepts.
	// Default 0, if this is less than MinArg then the function supports MaxInt arguments
	MaxArg int

	// Initial is a function to be called to get the initial valur for an aggregator.
	// If not set then InitialFirst is used.
	Initial FunctionInitialiser

	// Reducer is a value.Comparator used to determine which value to keep.
	// If comparator(a,b) returns true then a is used, otherwise b.
	Reducer value.Comparator

	// Calculation is a value.Calculation which performs an operation against two Value's
	// returning a new one.
	Calculation value.Calculation // calculation of a and b

	// Aggregator is a function which will convert a Value based on the number of values
	// passed through the aggregator.
	//
	// l=number of entries in set, return Value based on l and result
	Aggregator Aggregator

	// Function is called after any aggregations have been performed.
	Function FunctionHandler // Handler for specific functions

	// AggregateArguments if set will apply the aggregator against the arguments before passing
	// the single result to the Function.
	// This only applies if Function is set.
	AggregateArguments bool
}

type Aggregator func(l int, a Value) Value

// NullAggregator is an Aggregator which simply returns the value unchanged.
func NullAggregator(_ int, a Value) Value {
	return a
}

type FunctionInitialiser func(Value) Value

// IsAggregator returns true if one of Reducer, Calculation or Aggregator is defined.
func (f Function) IsAggregator() bool {
	// Note: Initial is not checked as it's valid to be nil as it will default to InitialFirst
	return f.Reducer != nil || f.Calculation != nil || f.Aggregator != nil
}

type FunctionHandler func(*Executor, lang.Visitor, *lang.Function, []Value) error

type FunctionMap struct {
	mutex     sync.Mutex
	functions map[string]Function
}

func (f *FunctionMap) AddFunction(n string, ag Function) {
	n = strings.ToLower(n)

	f.mutex.Lock()
	defer f.mutex.Unlock()
	if _, exists := f.functions[n]; exists {
		panic(fmt.Errorf("function %q already defined", n))
	}
	f.functions[n] = ag
}

func (f *FunctionMap) GetFunction(n string) (Function, bool) {
	n = strings.ToLower(n)

	f.mutex.Lock()
	defer f.mutex.Unlock()
	ag, exists := f.functions[n]
	return ag, exists
}

var functions = FunctionMap{
	functions: map[string]Function{
		// average - avg(metric) or avg(metric, ...)
		"avg": {
			MinArg:      1,
			Initial:     InitialInvalid,
			Calculation: value.Add,
			Aggregator: func(l int, a Value) Value {
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
			Reducer:    value.GreaterThan,
			Aggregator: MathAggregator(math.Ceil),
		},
		// count(metric) the number of entries in the metrics set
		"count": {
			Args: 1,
			Aggregator: func(l int, a Value) Value {
				// Return l as a value with the unit of "integer"
				a.Value = value.Integer.Value(float64(l))
				return a
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

func (ex *Executor) runFunction(v lang.Visitor, f *lang.Function, agg Function) error {

	if err := assertExpressions(f.Pos, f.Name, f.Expressions, agg); err != nil {
		return err
	}

	// Process the arguments, applying any aggregator against each one
	var args []Value
	for _, e := range f.Expressions {
		err := v.Expression(e)
		if err != nil {
			return err
		}

		val, _ := ex.pop()
		if agg.IsAggregator() {
			val, err = ex.runAggregator(agg, val)
			if err != nil {
				return err
			}
		}

		args = append(args, val)
	}

	// We have a function so call it with the arguments
	if agg.Function != nil {
		if agg.AggregateArguments {
			val, err := ex.runAggregator(agg, Value{Time: ex.time, Values: args})
			if err == nil {
				err = agg.Function(ex, v, f, []Value{val})
			}
			if err != nil {
				return err
			}
		} else {
			if err := agg.Function(ex, v, f, args); err != nil {
				return err
			}
		}
	}

	// No function but we are an Aggregator then ensure we push a result
	if agg.Function == nil && agg.IsAggregator() {
		switch len(args) {
		// No args then push null
		case 0:
			ex.push(Value{Time: ex.time})

		// 1 arg so use it
		case 1:
			ex.push(args[0])

		// Aggregate the args to get the final result
		default:
			val, err := ex.runAggregator(agg, Value{Time: ex.time, Values: args})
			if err != nil {
				return err
			}
			ex.push(val)
		}
	}

	return lang.VisitorStop
}

func (ex *Executor) runAggregator(agg Function, val Value) (Value, error) {
	var a Value

	if agg.Initial == nil {
		a = InitialFirst(val)
	} else {
		a = agg.Initial(val)
	}

	l := len(val.Values)
	switch {
	case agg.Reducer != nil:
		for _, b := range val.Values {
			// Only check if b is valid
			if b.Value.IsValid() {
				// If a is invalid then take b otherwise pass to the reducer
				if !a.Value.IsValid() {
					a = b
				} else {
					c, err := a.Value.Compare(b.Value, agg.Reducer)
					switch {
					case err != nil:
						return Value{}, err

					// If false then use b, if true keep a
					case !c:
						a = b
					}
				}
			}
		}

	case agg.Calculation != nil:
		for _, b := range val.Values {
			// Only check if b is valid
			if b.Value.IsValid() {
				// If a is invalid then take b otherwise pass to the reducer
				if !a.Value.IsValid() {
					a = b
				} else {
					nv, err := a.Value.Calculate(b.Value, agg.Calculation)
					if err != nil {
						return Value{}, err
					}
					a.Value = nv
				}
			}
		}
	}

	if a.Value.IsValid() && agg.Aggregator != nil {
		a = agg.Aggregator(l, a)
	}

	return a, nil
}

// InitialFirst returns the first valid metric in a Value.
// If the Value has no valid metrics then it returns the original value.
func InitialFirst(v Value) Value {
	if len(v.Values) > 0 {
		for i := 0; i < len(v.Values); i++ {
			e := v.Values[i]
			if e.Value.IsValid() {
				return e
			}
		}
		return v.Values[0]
	}
	return v
}

// InitialLast returns the last valid metric in a Value.
// If the Value has no valid metrics then it returns the original value.
func InitialLast(v Value) Value {
	if len(v.Values) > 0 {
		for i := len(v.Values) - 1; i >= 0; i-- {
			e := v.Values[i]
			if e.Value.IsValid() {
				return e
			}
		}
		return v.Values[len(v.Values)-1]
	}
	return v
}

// InitialInvalid returns an invalid Value.
// It's used for functions using Calculation's, where we want to start the aggregation
// without an initial starting value.
func InitialInvalid(_ Value) Value {
	return Value{}
}

func assertExpressions(p lexer.Position, n string, e []*lang.Expression, agg Function) error {
	// Here start with MinArg & MaxArg.
	// Override with Args if it's greater than either of them.
	// Enforce min>=0 but if max<min or negative then set max to MaxInt
	min, max := agg.MinArg, agg.MaxArg
	if agg.Args > min && agg.Args > max {
		min, max = agg.Args, agg.Args
	}
	if min < 0 {
		min = 0
	}
	if max < min || max < 0 {
		max = math.MaxInt
	}

	l := len(e)
	if l < min || l > max {
		if min == max {
			return participle.Errorf(p, "%s require %d expressions", n, min)
		}
		return participle.Errorf(p, "%s require %d..%d expressions", n, min, max)
	}

	return nil
}

// funcTimeOf implements TIMEOF which marks the value as requiring the TIME not the Value of a metric
func funcTimeOf(ex *Executor, v lang.Visitor, f *lang.Function, args []Value) error {
	switch len(args) {
	case 0:
		ex.push(Value{
			Time:   ex.time,
			IsTime: true,
		})

	case 1:
		if err := v.Expression(f.Expressions[0]); err != nil {
			return err
		}

		r, ok := ex.pop()
		if ok {
			r.IsTime = true
			ex.push(r)
		}

	default:
		return participle.Errorf(f.Pos, "Invalid state %d args expected 0..1", len(args))
	}

	return lang.VisitorStop
}

func funcTrend(ex *Executor, _ lang.Visitor, _ *lang.Function, args []Value) error {
	r := Value{Time: ex.time}

	if len(args) == 2 {
		r1, r2 := args[0], args[1]

		// Ensure r1 is temporally before r2
		if r1.Time.After(r2.Time) {
			r1, r2 = r2, r1
		}

		// Ensure we have a single and not a set Value.
		// If you want to use First instead then declare that in the ql
		r1 = InitialLast(r1)
		r2 = InitialLast(r2)

		// If both are still valid then do r2-r1 and set the Value to:
		// 0 if both are the same - e.g. steady
		// 1 if r2 is greater than r1 - e.g. trending up
		// -1 if r2 is less than r1 - e.g. trending down
		if r1.Value.IsValid() && r2.Value.IsValid() {
			d, err := r2.Value.Subtract(r1.Value)
			df := d.Float()
			if err == nil {
				switch {
				case value.IsZero(df):
					r.Value = value.Integer.Value(0)
				case value.IsNegative(df):
					r.Value = value.Integer.Value(-1)
				default:
					r.Value = value.Integer.Value(1)
				}
			}
		}
	}

	ex.push(r)

	return lang.VisitorStop
}

type SingleMathOperation func(float64) float64

// MathAggregator applies a SingleMathOperation against the result
func MathAggregator(o SingleMathOperation) Aggregator {
	return func(_ int, v Value) Value {
		if v.Value.IsValid() {
			v.Value = v.Value.Value(o(v.Value.Float()))
		}
		return v
	}
}
