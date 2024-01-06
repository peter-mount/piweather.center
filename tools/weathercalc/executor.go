package weathercalc

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/store/memory"
	"github.com/peter-mount/piweather.center/tools/weathercalc/functions"
	"github.com/peter-mount/piweather.center/tools/weathercalc/lang"
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

type executor struct {
	script *lang.Script
	calc   *Calculation
	latest memory.Latest
	stack  []StackEntry
}

type StackEntry struct {
	Time  time.Time
	Value value.Value
}

func (se StackEntry) IsValid() bool {
	return !se.Time.IsZero() && se.Value.IsValid()
}

func (e *executor) resetStack() {
	e.stack = nil
}

func (e *executor) pushNull() {
	e.stack = append(e.stack, StackEntry{})
}

func (e *executor) push(t time.Time, v value.Value) {
	e.stack = append(e.stack, StackEntry{Time: t, Value: v})
}

func (e *executor) pop() (StackEntry, bool) {
	if e.stackEmpty() {
		return StackEntry{}, false
	}
	sl := len(e.stack) - 1
	r := e.stack[sl]
	e.stack = e.stack[:sl]
	return r, true
}

func (e *executor) stackEmpty() bool {
	return len(e.stack) == 0
}

func (e *executor) setMetric(m string, v StackEntry) {
	if v.IsValid() {
		e.latest.Append(m, record.Record{
			Time:  v.Time,
			Value: v.Value,
		})
		log.Printf("set %q %s %s", m, v.Value, v.Time.Format(time.RFC3339))
	}
}

func (calc *Calculator) calculateResult(c *Calculation) (value.Value, time.Time, error) {
	log.Printf("Calculating %q", c.ID())

	e := executor{
		script: calc.Calculations.Script(),
		calc:   c,
		latest: calc.Latest,
	}

	err := e.calc.Src().Accept(lang.NewBuilder().
		Calculation(e.calculation).
		Current(e.current).
		Expression(e.expression).
		Function(e.function).
		Metric(e.metric).
		Unit(e.unit).
		Build())

	r, empty := e.pop()
	if err != nil || !empty {
		return value.Value{}, time.Time{}, err
	}

	log.Printf("result %q = %s %s", c.ID(), r.Value, r.Time.Format(time.RFC3339))

	return r.Value, r.Time, nil
}

func (e *executor) calculation(v lang.Visitor, b *lang.Calculation) error {
	e.resetStack()

	var val StackEntry

	// If usefirst set check to see we have a latest value, if not then set the default value
	if b.UseFirst != nil {
		_, exists := e.latest.Latest(b.Target)
		if !exists {
			err := b.UseFirst.Accept(v)
			if err != nil {
				return err
			}
			val, _ = e.pop()
		}
	}

	if !val.IsValid() {
		err := b.Expression.Accept(v)
		if err != nil {
			return err
		}

		val, _ = e.pop()
	}

	e.setMetric(b.Target, val)

	return lang.VisitorStop
}

func (e *executor) expression(v lang.Visitor, b *lang.Expression) error {
	var err error
	switch {
	case b.Current != nil:
		err = b.Current.Accept(v)
	case b.Function != nil:
		err = b.Function.Accept(v)
	case b.Metric != nil:
		err = b.Metric.Accept(v)
	}

	if err == nil && b.Using != nil {
		err = b.Using.Accept(v)
	}

	if err != nil {
		return err
	}

	return lang.VisitorStop
}

func (e *executor) current(_ lang.Visitor, _ *lang.Current) error {
	return e.metricImpl(e.calc.ID())
}

func (e *executor) metric(_ lang.Visitor, b *lang.Metric) error {
	return e.metricImpl(b.Name)
}

func (e *executor) metricImpl(n string) error {
	rec, exists := e.latest.Latest(n)
	if exists {
		e.push(rec.Time, rec.Value)
	} else {
		e.pushNull()
	}
	return nil
}

func (e *executor) unit(_ lang.Visitor, b *lang.Unit) error {
	v, empty := e.pop()
	if !empty {
		nv, err := v.Value.As(b.Unit())
		if err != nil {
			return err
		}
		e.push(v.Time, nv)
	}
	return nil
}

func (e *executor) function(v lang.Visitor, b *lang.Function) error {
	f, exists := functions.LookupFunction(b.Name)
	if !exists {
		return participle.Errorf(b.Pos, "function %q undefined", b.Name)
	}

	var args []StackEntry
	for _, exp := range b.Expressions {
		err := exp.Accept(v)
		if err != nil {
			return err
		}
		se, _ := e.pop()
		args = append(args, se)
	}

	switch {
	case f.Calculator != nil:
	case f.Op != nil:
		return e.op(b.Pos, f.Op, args)
	case f.MathOp != nil:
		return e.mathOp(b.Pos, f.MathOp, args)
	case f.MathBiOp != nil:
		return e.mathBiOp(b.Pos, f.MathBiOp, args)
	}
	return nil
}

func getTime(args []StackEntry) time.Time {
	var t time.Time
	var v *value.Unit
	for _, arg := range args {
		if t.IsZero() || t.Before(arg.Time) {
			t = arg.Time
		}
		if v == nil && arg.Value.IsValid() {
			v = arg.Value.Unit()
		}
	}
	if t.IsZero() {
		t = time.Now().UTC()
	}
	return t
}

func (e *executor) pushValue(f float64, args []StackEntry) {
	var t time.Time
	var v *value.Unit
	for _, arg := range args {
		if t.IsZero() || t.Before(arg.Time) {
			t = arg.Time
		}
		if v == nil && arg.Value.IsValid() {
			v = arg.Value.Unit()
		}
	}
	if t.IsZero() {
		t = time.Now().UTC()
	}
	if v == nil {
		v = value.Float
	}

	e.push(t, v.Value(f))
}

func (e *executor) op(pos lexer.Position, f func(value.Value) (value.Value, error), args []StackEntry) error {
	if len(args) != 1 {
		return participle.Errorf(pos, "expected 1 arg")
	}
	v, err := f(args[0].Value)
	if err != nil {
		return err
	}
	e.push(getTime(args), v)
	return nil
}

func (e *executor) biOp(pos lexer.Position, f func(value.Value, value.Value) (value.Value, error), args []StackEntry) error {
	if len(args) != 2 {
		return participle.Errorf(pos, "expected 2 args")
	}

	v, err := f(args[0].Value, args[1].Value)
	if err != nil {
		return err
	}
	e.push(getTime(args), v)
	return nil
}

func (e *executor) triOp(pos lexer.Position, f func(value.Value, value.Value, value.Value) (value.Value, error), args []StackEntry) error {
	if len(args) != 3 {
		return participle.Errorf(pos, "expected 3 args")
	}

	v, err := f(args[0].Value, args[1].Value, args[2].Value)
	if err != nil {
		return err
	}
	e.push(getTime(args), v)
	return nil
}

func (e *executor) mathOp(pos lexer.Position, f func(float64) float64, args []StackEntry) error {
	if len(args) != 1 {
		return participle.Errorf(pos, "expected 1 arg")
	}

	e.pushValue(f(args[0].Value.Float()), args)
	return nil
}

func (e *executor) mathBiOp(pos lexer.Position, f func(float64, float64) float64, args []StackEntry) error {
	if len(args) != 2 {
		return participle.Errorf(pos, "expected 2 args")
	}

	e.pushValue(f(args[0].Value.Float(), args[1].Value.Float()), args)
	return nil
}
