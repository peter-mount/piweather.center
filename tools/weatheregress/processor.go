package weatheregress

import (
	"encoding/json"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/tools/weatheregress/lang"
	"github.com/rabbitmq/amqp091-go"
)

func (s *Egress) initProcessor() {
	s.processor = lang.NewBuilder().
		Build()
}

// processMetricUpdate accepts a metric from RabbitMQ, updates it in Latest
// then forwards it to any calculations
func (s *Egress) processMetricUpdate(delivery amqp091.Delivery) error {
	var metric api.Metric
	err := json.Unmarshal(delivery.Body, &metric)
	if err == nil {
		err = s.processMetric(metric)
	}
	return err
}

func (s *Egress) processMetric(metric api.Metric) error {
	metrics := s.script.State().GetMetrics(metric.Metric)
	for _, m := range metrics {
		log.Printf("Found %q @ %T\n", metric.Metric, m)
	}
	return nil
}
