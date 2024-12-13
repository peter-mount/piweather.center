package weathercalc

import (
	"encoding/json"
	"github.com/peter-mount/go-build/version"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/broker"
	"github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/store/memory"
	"github.com/peter-mount/piweather.center/util/mq/amqp"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/rabbitmq/amqp091-go"
)

type Service struct {
	Latest         memory.Latest         `kernel:"inject"`
	DatabaseBroker broker.DatabaseBroker `kernel:"inject"`
	Calculator     *Calculator           `kernel:"inject"`
	Daemon         *kernel.Daemon        `kernel:"inject"`
	QueueName      *string               `kernel:"flag,metric-queue,DB queue name,database.calc"`
	mqQueue        *amqp.Queue
}

func (s *Service) Start() error {

	s.mqQueue = &amqp.Queue{
		Name:       *s.QueueName,
		Durable:    true,
		AutoDelete: false,
	}

	err := s.DatabaseBroker.ConsumeKeys(s.mqQueue, "calc", s.processMetricUpdate, "metric.#")

	if err == nil {
		log.Println(version.Version)
	}

	// Mark the application as a daemon
	s.Daemon.SetDaemon()

	return nil
}

// processMetricUpdate accepts a metric from RabbitMQ, updates it in Latest
// then forwards it to any calculations
func (s *Service) processMetricUpdate(delivery amqp091.Delivery) error {
	var metric api.Metric
	err := json.Unmarshal(delivery.Body, &metric)
	if err == nil {
		s.storeLatest(metric)
	}
	return err
}

func (s *Service) storeLatest(metric api.Metric) {
	u, ok := value.GetUnit(metric.Unit)
	if ok {
		updated := s.Latest.Append(metric.Metric, record.Record{
			Time:  metric.Time,
			Value: u.Value(metric.Value),
		})

		if updated {
			s.Calculator.Accept(metric)
		}
	}
}
