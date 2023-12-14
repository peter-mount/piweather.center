package exec

import (
	"fmt"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/store/server/ql/lang"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
	"sync"
)

func (ex *executor) function(v lang.Visitor, f *lang.Function) error {
	switch {
	case f.TimeOf:
		return ex.funcTimeOf(v, f)
	default:
		if af, exists := aggregators.GetAggregator(f.Name); exists {
			return ex.runAggregator(v, f, af)
		}
		return fmt.Errorf("unknown function %q", f.Name)
	}
}

type AggregatorFunction struct {
	Initial     func(Value) Value          // Get initial value, nil for first entry
	Reducer     value.Comparator           // reducer(a,b) returns true then take a, false take b
	Calculation value.Calculation          // calculation of a and b
	Aggregator  func(l int, a Value) Value // l=number of entries in set, return Value based on l and result
}

type FunctionMap struct {
	mutex       sync.Mutex
	aggregators map[string]AggregatorFunction
}

func (f *FunctionMap) AddAggregator(n string, ag AggregatorFunction) {
	n = strings.ToLower(n)

	f.mutex.Lock()
	defer f.mutex.Unlock()
	if _, exists := f.aggregators[n]; exists {
		panic(fmt.Errorf("function %q already defined", n))
	}
	f.aggregators[n] = ag
}

func (f *FunctionMap) GetAggregator(n string) (AggregatorFunction, bool) {
	n = strings.ToLower(n)

	f.mutex.Lock()
	defer f.mutex.Unlock()
	ag, exists := f.aggregators[n]
	return ag, exists
}

var aggregators = FunctionMap{
	aggregators: map[string]AggregatorFunction{
		"avg": {
			Calculation: value.Add,
			Aggregator: func(l int, a Value) Value {
				if a.Value.IsValid() {
					a.Value = a.Value.Unit().Value(a.Value.Float() / float64(l))
				}
				return a
			}},
		"count": {
			Aggregator: func(l int, a Value) Value {
				// Return l as a value with the unit of "integer"
				u, _ := value.GetUnit("integer")
				a.Value = u.Value(float64(l))
				return a
			}},
		"first": {Initial: InitialFirst},
		"last":  {Initial: InitialLast},
		"min":   {Reducer: value.LessThan},
		"max":   {Reducer: value.GreaterThan},
		"sum":   {Calculation: value.Add},
	},
}

func (ex *executor) runAggregator(v lang.Visitor, f *lang.Function, agg AggregatorFunction) error {
	return ex.funcEval1(v, f, func(v lang.Visitor, f *lang.Function, val Value) (Value, error) {
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
				af := a.Value.Float()

				// Only check if b is valid
				if b.Value.IsValid() {
					// If a is invalid then take b otherwise pass to the reducer
					if !a.Value.IsValid() {
						a = b
					} else {
						bf, err := b.Value.As(a.Value.Unit())
						if err != nil {
							return Value{}, err
						}

						if !agg.Reducer(af, bf.Float()) {
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
	})
}

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

type funcEvaluator func(v lang.Visitor, f *lang.Function, val Value) (Value, error)

func (ex *executor) funcEval1(v lang.Visitor, f *lang.Function, h funcEvaluator) error {
	err := assertExpressions(f.Pos, f.Expressions, 1, 1)

	if err == nil {
		err = v.Expression(f.Expressions[0])
	}

	if err == nil {
		r, ok := ex.pop()
		// Apply reduction if a valid Value but an invalid singular value
		if ok && !r.Value.IsValid() {
			r, err = h(v, f, r)
		}
		if err == nil {
			ex.push(r)
		}
	}

	if err != nil {
		return err
	}

	return lang.VisitorStop
}

func assertExpressions(p lexer.Position, e []*lang.Expression, min, max int) error {
	if min > max {
		min, max = max, min
	}
	l := len(e)
	if l < min || l > max {
		if min == max {
			return participle.Errorf(p, "require %d expressions", min)
		}
		return participle.Errorf(p, "require %d..%d expressions", min, max)
	}
	return nil
}

// funcTimeOf implements TIMEOF which marks the value as requiring the TIME not the Value of a metric
func (ex *executor) funcTimeOf(v lang.Visitor, f *lang.Function) error {
	if err := assertExpressions(f.Pos, f.Expressions, 0, 1); err != nil {
		return err
	}

	if len(f.Expressions) == 0 {
		ex.push(Value{
			Time:   ex.time,
			IsTime: true,
		})
	} else {
		if err := v.Expression(f.Expressions[0]); err != nil {
			return err
		}

		r, ok := ex.pop()
		if ok {
			r.IsTime = true
			ex.push(r)
		}
	}
	return lang.VisitorStop
}
