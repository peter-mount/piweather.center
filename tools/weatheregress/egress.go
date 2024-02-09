package weatheregress

import (
	"encoding/json"
	"github.com/peter-mount/go-build/version"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/mq/amqp"
	"github.com/peter-mount/piweather.center/mq/mqtt"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/broker"
	"github.com/rabbitmq/amqp091-go"
)

type Egress struct {
	Amqp           amqp.Pool             `kernel:"inject"`
	Mqtt           mqtt.Pool             `kernel:"inject"`
	DatabaseBroker broker.DatabaseBroker `kernel:"inject"`
	Processor      *Processor            `kernel:"inject"`
	Daemon         *kernel.Daemon        `kernel:"inject"`
	QueueName      *string               `kernel:"flag,metric-queue,DB queue name,database.calc"`
	mqQueue        *amqp.Queue
}

func (s *Egress) Start() error {
	s.mqQueue = &amqp.Queue{
		Name:       *s.QueueName,
		Durable:    true,
		AutoDelete: false,
	}

	err := s.DatabaseBroker.ConsumeKeys(s.mqQueue, "egress", s.processMetricUpdate, "metric.#")
	if err != nil {
		return err
	}

	log.Println(version.Version)

	// Mark the application as a daemon
	s.Daemon.SetDaemon()

	return nil
}

// processMetricUpdate accepts a metric from RabbitMQ, updates it in Latest
// then forwards it to any calculations
func (s *Egress) processMetricUpdate(delivery amqp091.Delivery) error {
	var metric api.Metric
	err := json.Unmarshal(delivery.Body, &metric)
	if err == nil {
		s.Processor.ProcessMetric(metric)
	}
	return err
}
