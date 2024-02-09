package weatheregress

import (
	"context"
	"flag"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/util/task"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/tools/weatheregress/lang"
)

type Processor struct {
	Worker    task.Queue `kernel:"worker"`
	script    *lang.Script
	processor lang.Visitor
	action    *action
}

type action struct {
	proc    *Processor
	metric  api.Metric
	metrics []*lang.Metric
}

func (a *action) run(_ context.Context) error {
	// Capture any panics so we don't shut down
	defer func() {
		if err1 := recover(); err1 != nil {
			log.Printf("Panic %q %.3f %v\n",
				a.metric.Metric,
				a.metric.Value,
				err1)
		}
	}()
	a.proc.processAction(a)
	return nil
}

func (s *Processor) Start() error {
	p := lang.NewParser()
	script, err := p.ParseFiles(flag.Args()...)
	if err != nil {
		return err
	}

	s.script = script
	s.processor = lang.NewBuilder().
		Metric(s.metric).
		Publish(s.publish).
		Build()

	return nil
}

// ProcessMetric accepts a metric, checks to see if it's one we are interested in
// and if so places it into the work queue
func (s *Processor) ProcessMetric(metric api.Metric) {
	metrics := s.script.State().GetMetrics(metric.Metric)
	if len(metrics) > 0 {
		act := &action{
			proc:    s,
			metric:  metric,
			metrics: metrics,
		}
		s.Worker.AddTask(act.run)
	}
}

func (s *Processor) processAction(a *action) {
	s.action = a
	for _, m := range a.metrics {
		_ = m.Accept(s.processor)
	}
}

func (s *Processor) metric(v lang.Visitor, m *lang.Metric) error {
	return nil
}

func (s *Processor) publish(_ lang.Visitor, p *lang.Publish) error {
	switch {
	case p.Console:
		log.Printf("%s\n", s.action.metric)
	case p.Amqp != "":
	}
	return nil
}
