package weatheregress

import (
	"flag"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/tools/weatheregress/lang"
)

type Processor struct {
	script    *lang.Script
	processor lang.Visitor[*action]
}

type action struct {
	metric  api.Metric
	metrics []*lang.Metric
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

func (s *Processor) publish(v lang.Visitor[*action], p *lang.Publish) error {
	switch {
	case p.Console:
		log.Printf("%s\n", v.GetData().metric)
	case p.Amqp != "":
	}
	return nil
}
