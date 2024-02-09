package weatheregress

import (
	"flag"
	"fmt"
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-script/calculator"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/tools/weatheregress/lang"
)

type Processor struct {
	script    *lang.Script
	processor lang.Visitor[*action]
}

func (s *Processor) Start() error {
	p := lang.NewParser()
	script, err := p.ParseFiles(flag.Args()...)
	if err != nil {
		return err
	}

	s.script = script
	s.processor = lang.NewBuilder[*action]().
		Metric(s.metric).
		Format(s.format).
		FormatExpression(s.formatExpression).
		FormatAtom(s.formatAtom).
		Publish(s.publish).
		Build()

	return nil
}

// ProcessMetric accepts a metric, checks to see if it's one we are interested in
// and if so places it into the work queue
func (s *Processor) ProcessMetric(metric api.Metric) {
	metrics := s.script.State().GetMetrics(metric.Metric)
	if len(metrics) > 0 {
		s.processAction(&action{
			metric:  metric,
			metrics: metrics,
			message: metric.String(),
			calc:    calculator.New(),
		})
	}
}

func (s *Processor) processAction(a *action) {
	// Capture any panics so we don't shut down
	defer func() {
		if err1 := recover(); err1 != nil {
			log.Printf("Panic %q %.3f %v\n",
				a.metric.Metric,
				a.metric.Value,
				err1)
		}
	}()
	// Get a new visitor with our data attached
	ctx := s.processor.SetData(a)
	for _, m := range a.metrics {
		_ = ctx.Metric(m)
	}
}

func (s *Processor) metric(v lang.Visitor[*action], m *lang.Metric) error {
	return nil
}

func (s *Processor) format(v lang.Visitor[*action], f *lang.Format) error {
	act := v.GetData()

	var args []any
	for _, e := range f.Expressions {
		act.calc.Reset()
		if err := v.FormatExpression(e); err != nil {
			return err
		}

		val, err := act.calc.Pop()
		if err != nil {
			return err
		}
		args = append(args, val)
	}

	act.message = fmt.Sprintf(f.Format, args...)
	return nil
}

func (s *Processor) formatExpression(v lang.Visitor[*action], f *lang.FormatExpression) error {
	err := v.FormatAtom(f.Left)

	if err == nil && f.Op != "" {
		err = v.FormatExpression(f.Right)

		if err == nil {
			err = v.GetData().calc.Op2(f.Op)
		}
	}

	return err
}

func (s *Processor) formatAtom(v lang.Visitor[*action], f *lang.FormatAtom) error {
	act := v.GetData()
	calc := act.calc
	metric := act.metric

	switch {
	case f.Metric:
		calc.Push(metric.Metric)

	case f.Value:
		calc.Push(metric.Value)

	case f.UnixTime:
		calc.Push(metric.Time.Unix())

	case f.String != nil:
		calc.Push(*f.String)

	default:
		return participle.Errorf(f.Pos, "invalid atom")
	}

	return nil
}

func (s *Processor) publish(v lang.Visitor[*action], p *lang.Publish) error {
	switch {
	case p.Console:
		log.Printf("%s\n", v.GetData().message)
	case p.Amqp != "":
	}
	return nil
}
