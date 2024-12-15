package expression

import (
	"fmt"
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-script/errors"
	station2 "github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/units"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/client"
	"github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/store/memory"
	_ "github.com/peter-mount/piweather.center/weather/forecast"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

type Executor interface {
	CalculateResult(*station2.Calculation) (value.Value, time.Time, error)
	Evaluate(*station2.Expression) (value.Value, time.Time, error)
}

type executor struct {
	//service *Calculator
	dbServer        string        // Database server url, "" for none
	latest          memory.Latest // Latest metric service
	visitor         station2.Visitor[*executor]
	currentMetricId string         // metric to return with the current operator
	calc            *Calculation   // Calculation being processed
	stack           []StackEntry   // expression stack
	stacks          [][]StackEntry // saved expression stacks
	time            value.Time     // Time for expression queries
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
	e.stack = nil
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

func (e *executor) CalculateResult(c *station2.Calculation) (value.Value, time.Time, error) {

	err := e.visitor.Calculation(c)

	r, exists := e.pop()
	if err != nil {
		log.Printf("calc %q %v", e.currentMetricId, err)
	}
	if err != nil || !exists || !r.IsValid() {
		return value.Value{}, time.Time{}, err
	}

	if log.IsVerbose() {
		log.Printf("Result %q %s %s", e.currentMetricId, r.Value, r.Time.Format(time.RFC3339))
	}

	return r.Value, r.Time, nil
}

func (e *executor) Evaluate(c *station2.Expression) (value.Value, time.Time, error) {

	err := e.visitor.Expression(c)

	r, exists := e.pop()
	if err != nil {
		log.Printf("eval %q %v", e.currentMetricId, err)
	}
	if err != nil || !exists || !r.IsValid() {
		return value.Value{}, time.Time{}, err
	}

	if log.IsVerbose() {
		log.Printf("eval result %q %s %s", e.currentMetricId, r.Value, r.Time.Format(time.RFC3339))
	}

	return r.Value, r.Time, nil
}

func NewExecutor(currentMetricId string, t value.Time, dbServer string, latest memory.Latest) Executor {
	e := &executor{
		dbServer:        dbServer,
		currentMetricId: currentMetricId,
		latest:          latest,
		// Time at the location
		time: t,
	}

	e.visitor = station2.NewBuilder[*executor]().
		Calculation(e.calculation).
		Current(e.current).
		Expression(e.expression).
		Function(e.function).
		LocationExpression(e.locationExpression).
		Metric(e.metric).
		MetricExpression(e.metricExpression).
		Unit(e.unit).
		UseFirst(e.useFirst).
		Build().
		Set(e)

	return e
}

func (e *executor) calculation(v station2.Visitor[*executor], b *station2.Calculation) error {
	e.resetStack()

	// If usefirst set check to see we have a latest value, if not then set the default value
	if b.UseFirst != nil {
		err := v.UseFirst(b.UseFirst)
		if err != nil {
			return err
		}
	}

	if b.Expression == nil {
		// Handle no AS clause - result is the target metric
		r, exists := e.latest.Latest(b.Target)
		if !exists {
			return util.VisitorStop
		}
		e.push(r.Time, r.Value)
	} else {
		// Evaluate the expression
		err := v.Expression(b.Expression)
		if err != nil {
			return err
		}
	}
	val, _ := e.peek()

	e.setMetric(b.Target, val)

	return util.VisitorStop
}

func (e *executor) expression(v station2.Visitor[*executor], b *station2.Expression) error {
	var err error
	switch {
	case b.Current != nil:
		err = v.Current(b.Current)
	case b.Function != nil:
		err = v.Function(b.Function)
	case b.Metric != nil:
		err = v.MetricExpression(b.Metric)
	case b.Location != nil:
		err = v.LocationExpression(b.Location)
	}

	if err == nil && b.Using != nil {
		err = v.Unit(b.Using)
		if err != nil {
			// Use this so the user is told the file/line of the error
			return participle.Errorf(b.Pos, "%s", err.Error())
		}
	}

	if err != nil {
		return err
	}

	return util.VisitorStop
}

func (e *executor) current(_ station2.Visitor[*executor], _ *station2.Current) error {
	return e.metricImpl(e.calc.ID())
}

func (e *executor) metric(_ station2.Visitor[*executor], b *station2.Metric) error {
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

func (e *executor) metricExpression(v station2.Visitor[*executor], b *station2.MetricExpression) error {
	var err error
	var res *api.Result

	// Offset so call DB to get the metric after the offset
	if b.HasOffset() {
		if b.GetOffset() >= 0 {
			return participle.Errorf(b.Pos, "Expected offset into the past, got %q", b.Offset)
		}

		q := fmt.Sprintf(`between "now" add %q and "now" limit 1 select timeof(),first(%s)`, b.Offset, b.Metric.Name)

		cl := client.Client{Url: e.dbServer}
		res, err = cl.Query(q)
		if err == nil {
			if len(res.Table) > 0 {
				if t := res.Table[0]; !t.IsEmpty() {
					if r := t.Rows[0]; r.Size() > 1 {
						tc := r.Cell(0)
						vc := r.Cell(1)
						if vc.Value.IsValid() {
							e.push(tc.Time, vc.Value)
							return util.VisitorStop
						}
					}
				}
			}

			// No result so set null
			e.pushNull()
			return util.VisitorStop
		}
	}

	return errors.Error(b.Pos, err)
}

func (e *executor) useFirst(_ station2.Visitor[*executor], b *station2.UseFirst) error {
	rec, exists := e.latest.Latest(e.calc.ID())
	if !exists {
		rec, exists = e.latest.Latest(b.Metric.Name)
		if exists {
			// Set the new value then VisitorStop to tell
			// calculation() to terminate
			e.latest.Append(e.calc.ID(), rec)
			return util.VisitorStop
		}
	}
	return nil
}

func (e *executor) unit(_ station2.Visitor[*executor], b *units.Unit) error {
	v, present := e.pop()
	if present {
		nv, err := v.Value.As(b.Unit())
		if err != nil {
			return err
		}
		e.push(v.Time, nv)
	} else {
		e.pushNull()
	}
	return nil
}

func (e *executor) function(v station2.Visitor[*executor], b *station2.Function) error {
	e.save()

	calc, err := value.GetCalculator(b.Name)
	if err != nil {
		e.restore()
		return participle.Errorf(b.Pos, "%s", err.Error())
	}

	var t time.Time
	var args []value.Value
	for _, exp := range b.Expressions {
		err = v.Expression(exp)
		if err != nil {
			e.restore()
			return err
		}

		arg, _ := e.pop()
		// Not valid then stop here
		if !arg.IsValid() {
			e.restore()
			e.pushNull()
			return util.VisitorStop
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
		return util.VisitorStop
	}

	return err
}

func (e *executor) locationExpression(v station2.Visitor[*executor], b *station2.LocationExpression) error {
	var err error
	if b != nil {
		var val value.Value

		loc := e.calc.Station().Location
		ll := loc.LatLong()
		switch {
		case b.Altitude:
			val = measurement.Meters.Value(loc.Altitude)
		case b.Latitude:
			val = measurement.Degree.Value(ll.Latitude.Deg())
		case b.Longitude:
			val = measurement.Degree.Value(ll.Longitude.Deg())
		default:
		}

		if val.IsValid() {
			e.push(time.Now(), val)
		} else {
			e.pushNull()
		}
	}
	return err
}
