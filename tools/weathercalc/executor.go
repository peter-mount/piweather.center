package weathercalc

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/store/memory"
	"github.com/peter-mount/piweather.center/tools/weathercalc/lang"
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

type executor struct {
	script *lang.Script
	calc   *Calculation
	latest memory.Latest
	stack  []StackEntry
	time   value.Time
	stacks [][]StackEntry
}

type StackEntry struct {
	Time  time.Time
	Value value.Value
}

func (se StackEntry) IsValid() bool {
	return !se.Time.IsZero() && se.Value.IsValid()
}

func (e *executor) save() {
	e.stacks = append(e.stacks, e.stack)
	e.resetStack()
}

func (e *executor) restore() {
	l := len(e.stacks)
	if l == 0 {
		e.stack = nil
	} else {
		e.stack = e.stacks[l-1]
		e.stacks = e.stacks[:l-1]
	}
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

func (e *executor) peek() (StackEntry, bool) {
	if e.stackEmpty() {
		return StackEntry{}, false
	}
	return e.stack[len(e.stack)-1], true
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
	}
}

func (calc *Calculator) calculateResult(c *Calculation) (value.Value, time.Time, error) {
	e := executor{
		script: calc.Script(),
		calc:   c,
		latest: calc.Latest,
		time:   calc.script.State.GetLocation(c.Src().At).Time(),
	}

	err := c.Src().Accept(lang.NewBuilder().
		Calculation(e.calculation).
		Current(e.current).
		Expression(e.expression).
		Function(e.function).
		Metric(e.metric).
		Unit(e.unit).
		UseFirst(e.useFirst).
		Build())

	r, exists := e.pop()
	if err != nil {
		log.Printf("calc %q %v", c.ID(), err)
	}
	if err != nil || !exists || !r.IsValid() {
		return value.Value{}, time.Time{}, err
	}

	log.Printf("Result %q %s %s", c.ID(), r.Value, r.Time.Format(time.RFC3339))

	return r.Value, r.Time, nil
}

func (e *executor) calculation(v lang.Visitor, b *lang.Calculation) error {
	e.resetStack()

	// If usefirst set check to see we have a latest value, if not then set the default value
	if b.UseFirst != nil {
		err := b.UseFirst.Accept(v)
		if err != nil {
			return err
		}
	}

	err := b.Expression.Accept(v)
	if err != nil {
		return err
	}

	val, _ := e.peek()

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
		if err != nil {
			// Use this so the user is told the file/line of the error
			return participle.Errorf(b.Pos, "%s", err.Error())
		}
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
	if exists && rec.IsValid() {
		e.push(rec.Time, rec.Value)
	} else {
		e.pushNull()
	}
	return nil
}

func (e *executor) useFirst(_ lang.Visitor, b *lang.UseFirst) error {
	rec, exists := e.latest.Latest(e.calc.ID())
	if !exists {
		rec, exists = e.latest.Latest(b.Metric.Name)
		if exists {
			// Set the new value then VisitorStop to tell
			// calculation() to terminate
			e.latest.Append(e.calc.ID(), rec)
			return lang.VisitorStop
		}
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
	e.save()

	calc, err := value.GetCalculator(b.Name)
	if err != nil {
		e.restore()
		return participle.Errorf(b.Pos, "%s", err.Error())
	}

	var t time.Time
	var args []value.Value
	for _, exp := range b.Expressions {
		err = exp.Accept(v)
		if err != nil {
			e.restore()
			return err
		}

		arg, _ := e.pop()
		// Not valid then stop here
		if !arg.IsValid() {
			e.restore()
			e.pushNull()
			return nil
		}

		if t.IsZero() || t.Before(arg.Time) {
			t = arg.Time
		}
		args = append(args, arg.Value)
	}

	if t.IsZero() {
		t = time.Now()
	}
	e.time.SetTime(t)

	val, err := calc(e.time, args...)
	if err == nil {
		e.restore()
		e.push(t, val)
	}

	return err
}
