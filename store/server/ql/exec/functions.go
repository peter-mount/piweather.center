package exec

import (
	"fmt"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-kernel/v2/log"
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
		af, exists := functions.Get(f.Name)
		if !exists {
			return fmt.Errorf("unknown function %q", f.Name)
		}
		return ex.runAggregator(v, f, af)
	}
}

type AggregatorFunction struct {
	Initial    func(Value) Value          // Get initial value, nil for first entry
	Reducer    value.Comparator           // reducer(a,b) true take a, false take b
	Aggregator func(l int, a Value) Value // l=number of entries in set, return Value based on l and result
}

type FunctionMap struct {
	mutex sync.Mutex
	funcs map[string]AggregatorFunction
}

func (f *FunctionMap) Add(n string, ag AggregatorFunction) {
	n = strings.ToLower(n)

	f.mutex.Lock()
	defer f.mutex.Unlock()
	if _, exists := f.funcs[n]; exists {
		panic(fmt.Errorf("function %q already defined", n))
	}
	f.funcs[n] = ag
}

func (f *FunctionMap) Get(n string) (AggregatorFunction, bool) {
	n = strings.ToLower(n)

	f.mutex.Lock()
	defer f.mutex.Unlock()
	ag, exists := f.funcs[n]
	return ag, exists
}

var functions = FunctionMap{
	funcs: map[string]AggregatorFunction{
		"min":    {Reducer: value.LessThan},
		"max":    {Reducer: value.GreaterThan},
		"latest": {Initial: InitialLast},
		"first":  {Initial: InitialFirst},
	},
}

func (ex *executor) runAggregator(v lang.Visitor, f *lang.Function, agg AggregatorFunction) error {
	return ex.funcEval1(v, f, func(v lang.Visitor, f *lang.Function, val Value) (Value, error) {
		log.Printf("func %q", f.Name)
		var a Value
		if agg.Initial == nil {
			a = InitialFirst(val)
		} else {
			a = agg.Initial(a)
		}

		l := len(val.Values)
		log.Printf("%s(%d) -> %s %v", f.Name, l, a.Value, a.IsNull)
		if agg.Reducer != nil {
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
		}

		// Ensure we have 1 entry - e.g. prevent divide by zero
		if l < 1 {
			l = 1
		}

		if agg.Aggregator != nil {
			a = agg.Aggregator(l, a)
		}

		a.IsNull = !a.Value.IsValid()
		log.Printf("%s(%d) -> %s %v", f.Name, l, a.Value, a.IsNull)
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
