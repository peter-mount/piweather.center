package functions

import (
	"fmt"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	ql3 "github.com/peter-mount/piweather.center/config/ql"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/store/ql"
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
	"strings"
	"sync"
)

// Function definition
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

type Aggregator func(l int, a ql.Value) ql.Value

type FunctionInitialiser func(ql.Value) ql.Value

// IsAggregator returns true if one of Reducer, Calculation or Aggregator is defined.
func (f Function) IsAggregator() bool {
	// Note: Initial is not checked as it's valid to be nil as it will default to InitialFirst
	return f.Reducer != nil || f.Calculation != nil || f.Aggregator != nil
}

type FunctionHandler func(ql.Executor, ql3.QueryVisitor, *ql3.Function, []ql.Value) error

type FunctionMap struct {
	mutex     sync.Mutex
	functions map[string]Function
}

func AddFunction(n string, ag Function) *FunctionMap {
	functions.AddFunction(n, ag)
	return &functions
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

func HasFunction(n string) bool {
	return functions.HasFunction(n)
}

func (f *FunctionMap) HasFunction(n string) bool {
	_, exists := f.GetFunction(n)
	return exists
}

func GetFunction(n string) (Function, bool) {
	return functions.GetFunction(n)
}

func (f *FunctionMap) GetFunction(n string) (Function, bool) {
	n = strings.ToLower(n)

	f.mutex.Lock()
	defer f.mutex.Unlock()
	ag, exists := f.functions[n]
	return ag, exists
}

func (f Function) Run(ex ql.Executor, v ql3.QueryVisitor, fn *ql3.Function) error {

	if err := assertExpressions(fn.Pos, fn.Name, fn.Expressions, f); err != nil {
		return err
	}

	// Process the arguments, applying any aggregator against each one
	var args []ql.Value
	for _, e := range fn.Expressions {
		err := v.Expression(e)
		if err != nil {
			return err
		}

		val, _ := ex.Pop()

		if f.IsAggregator() {
			val, err = f.runAggregator(val)
			if err != nil {
				return err
			}
		}

		args = append(args, val)
	}

	// We have a function so call it with the arguments
	if f.Function != nil {

		// Aggregate the arguments into a single value
		if f.AggregateArguments {
			val, err := f.runAggregator(ql.Value{Time: ex.Time(), Values: args})
			if err != nil {
				return err
			}
			args = []ql.Value{val}
		}

		if err := f.Function(ex, v, fn, args); err != nil {
			return err
		}
	}

	// No function but we are an Aggregator then ensure we push a result
	if f.Function == nil && f.IsAggregator() {
		switch len(args) {
		// No args then push null
		case 0:
			ex.Push(ql.Value{Time: ex.Time()})

		// 1 arg so use it
		case 1:
			ex.Push(args[0])

		// Aggregate the args to get the final result
		default:
			val, err := f.runAggregator(ql.Value{Time: ex.Time(), Values: args})
			if err != nil {
				return err
			}
			ex.Push(val)
		}
	}

	return util.VisitorStop
}

func (f Function) runAggregator(val ql.Value) (ql.Value, error) {
	var a ql.Value

	if f.Initial == nil {
		a = InitialFirst(val)
	} else {
		a = f.Initial(val)
	}

	l := len(val.Values)
	switch {
	case f.Reducer != nil:
		for _, b := range val.Values {
			// Only check if b is valid
			if b.Value.IsValid() {
				// If a is invalid then take b otherwise pass to the reducer
				if !a.Value.IsValid() {
					a = b
				} else {
					c, err := a.Value.Compare(b.Value, f.Reducer)
					switch {
					case err != nil:
						return ql.Value{}, err

					// If false then use b, if true keep a
					case !c:
						a = b
					}
				}
			}
		}

	case f.Calculation != nil:
		for _, b := range val.Values {
			// Only check if b is valid
			if b.Value.IsValid() {
				// If a is invalid then take b otherwise pass to the reducer
				if !a.Value.IsValid() {
					a = b
				} else {
					nv, err := a.Value.Calculate(b.Value, f.Calculation)
					if err != nil {
						return ql.Value{}, err
					}
					a.Value = nv
				}
			}
		}
	}

	if a.Value.IsValid() && f.Aggregator != nil {
		a = f.Aggregator(l, a)
	}

	return a, nil
}

// InitialFirst returns the first valid metric in a Value.
// If the Value has no valid metrics then it returns the original value.
func InitialFirst(v ql.Value) ql.Value {
	if len(v.Values) > 0 {
		for _, e := range v.Values {
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
func InitialLast(v ql.Value) ql.Value {
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
func InitialInvalid(_ ql.Value) ql.Value {
	return ql.Value{}
}

func assertExpressions(p lexer.Position, n string, e []*ql3.Expression, agg Function) error {
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
func funcTimeOf(ex ql.Executor, v ql3.QueryVisitor, f *ql3.Function, args []ql.Value) error {
	switch len(args) {
	case 0:
		ex.Push(ql.Value{
			Time:   ex.Time(),
			IsTime: true,
		})

	case 1:
		if err := v.Expression(f.Expressions[0]); err != nil {
			return err
		}

		r, ok := ex.Pop()
		if ok {
			// if an invalid time then return nul
			if r.Time.IsZero() {
				r.Value = value.Value{}
			} else {
				r.IsTime = true
			}
		}
		ex.Push(r)

	default:
		return participle.Errorf(f.Pos, "Invalid station %d args expected 0..1", len(args))
	}

	return util.VisitorStop
}

func funcTrend(ex ql.Executor, _ ql3.QueryVisitor, _ *ql3.Function, args []ql.Value) error {
	r := ql.Value{Time: ex.Time()}

	if len(args) == 2 {
		r1, r2 := args[0], args[1]

		// Ensure we have a single and not a set Value.
		// If you want to use First instead then declare that in the ql
		r1 = InitialLast(r1)
		r2 = InitialLast(r2)

		// If both are still valid then do r2-r1 and set the Value to:
		// 0 if both are the same - e.g. steady
		// 1 if r2 is greater than r1 - e.g. trending up
		// -1 if r2 is less than r1 - e.g. trending down
		if r1.Value.IsValid() && r2.Value.IsValid() {

			// Ensure r1 is temporally before r2
			if r1.Time.After(r2.Time) {
				r1, r2 = r2, r1
			}

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

	ex.Push(r)

	return util.VisitorStop
}

type SingleMathOperation func(float64) float64

// MathAggregator applies a SingleMathOperation against the result
func MathAggregator(o SingleMathOperation) Aggregator {
	return func(_ int, v ql.Value) ql.Value {
		if v.Value.IsValid() {
			v.Value = v.Value.Value(o(v.Value.Float()))
		}
		return v
	}
}
