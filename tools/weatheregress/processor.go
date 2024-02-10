package weatheregress

import (
	"flag"
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
		Publish(s.publish).
		Build()

	return s.initMq()
}

// ProcessMetric accepts a metric, checks to see if it's one we are interested in
// and if so places it into the work queue
func (s *Processor) ProcessMetric(metric api.Metric) {
	for _, metrics := range s.script.State().GetMetrics(metric.Metric) {
		s.processAction(newAction(metric, metrics))
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
	err := ctx.Metric(a.metrics)
	if err != nil {
		log.Printf("Error %q %.3f %v\n",
			a.metric.Metric,
			a.metric.Value,
			err)
	}
}

func (s *Processor) metric(v lang.Visitor[*action], m *lang.Metric) error {
	var err error
	if m.Statement != nil {
		act := v.GetData()
		exec := act.exec

		scope := exec.GlobalScope()

		scope.Declare("metric")
		scope.Set("metric", act.metric)

		scope.Declare("message")
		scope.Set("message", act.message)

		scope.Declare("routingKey")
		scope.Set("routingKey", act.routingKey)

		exec.Calculator().Reset()
		err = exec.Statements(m.Statement)

		if err == nil {
			act.message, _ = scope.Get("message")

			if k, exists := scope.Get("routingKey"); exists {
				act.routingKey, err = calculator.GetString(k)
			}
		}
	}

	return err
}

func (s *Processor) publish(v lang.Visitor[*action], p *lang.Publish) error {
	switch {
	case p.Console:
		log.Printf("%s\n", v.GetData().message)
	case p.Amqp != "":
		return s.publishAmqp(v, p)
	}
	return nil
}
