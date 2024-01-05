package weatheringress

import (
	"context"
	"encoding/json"
	"github.com/peter-mount/go-build/version"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/mq/amqp"
	"github.com/peter-mount/piweather.center/mq/mqtt"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/station/service"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/broker"
	"github.com/peter-mount/piweather.center/store/client"
	"github.com/peter-mount/piweather.center/store/file/record"
	log2 "github.com/peter-mount/piweather.center/store/log"
	"github.com/peter-mount/piweather.center/store/memory"
	"github.com/peter-mount/piweather.center/util/endpoint"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/rabbitmq/amqp091-go"
)

// Ingress handles the ability to get data into the system, be it via
// http, amqp, mqtt etc.
type Ingress struct {
	Archiver        *log2.Archiver            `kernel:"inject"`
	Latest          memory.Latest             `kernel:"inject"`
	Amqp            amqp.Pool                 `kernel:"inject"`
	Mqtt            mqtt.Pool                 `kernel:"inject"`
	EndpointManager *endpoint.EndpointManager `kernel:"inject"`
	Config          service.Config            `kernel:"inject"`
	DatabaseBroker  broker.DatabaseBroker     `kernel:"inject"`
	QueueName       *string                   `kernel:"flag,metric-queue,DB queue name,database.ingress"`
	DBServer        *string                   `kernel:"flag,metric-db,DB url"`
	mqQueue         *amqp.Queue
	subContext      context.Context        // Common Context
	processVisitor  station.VisitorBuilder // Common visitor used by all sources to process data
	updateVisitor   station.Visitor        // Visitor used to process updates sent to the DB to handle calculated updates
}

func (s *Ingress) Start() error {

	// Get latest metrics from DB
	if err := s.loadLatestMetrics(); err != nil {
		return err
	}

	// Common context for processing
	s.subContext = s.Archiver.AddContext(context.Background())

	// Visitor that will process an inbound message.
	// This is common to all sources, so we define it here, but they will
	// build it as needed.
	s.processVisitor = station.NewVisitor().
		Sensors(value.ResetMap).
		Sensors(s.Archiver.Archive).
		Reading(s.processReading).
		CalculatedValue(s.calculate).
		Output(s.databaseReading)

	// Now start the sources
	if err := s.Config.Accept(station.NewVisitor().
		Sensors(s.startAMQP).
		Sensors(s.startEcowitt).
		WithContext(value.WithMap(s.subContext))); err != nil {
		return err
	}

	s.mqQueue = &amqp.Queue{
		Name:       *s.QueueName,
		Durable:    true,
		AutoDelete: false,
	}

	err := s.DatabaseBroker.ConsumeKeys(s.mqQueue, "ingress", s.processMetricUpdate, "metric.#")

	if err == nil {
		log.Println(version.Version)
	}

	return nil
}

// loadLatestMetrics retrieves the current metrics from the DB server
func (s *Ingress) loadLatestMetrics() error {
	if *s.DBServer != "" {
		c := &client.Client{Url: *s.DBServer}
		r, err := c.LatestMetrics()
		if err != nil {
			return err
		}
		for _, m := range r.Metrics {
			s.storeLatest(m)
		}
	}
	return nil
}

// processMetricUpdate accepts a metric from RabbitMQ, updates it in Latest
// then forwards it to any calculations
func (s *Ingress) processMetricUpdate(delivery amqp091.Delivery) error {
	var metric api.Metric
	err := json.Unmarshal(delivery.Body, &metric)
	if err == nil {
		s.storeLatest(metric)
	}
	return err
}

func (s *Ingress) storeLatest(metric api.Metric) {
	u, ok := value.GetUnit(metric.Unit)
	if ok {
		log.Println("Store", metric.Metric)
		updated := s.Latest.Append(metric.Metric, record.Record{
			Time:  metric.Time,
			Value: u.Value(metric.Value),
		})

		if updated {
			metric.Formatted = u.String(metric.Value)
			metric.Unix = metric.Time.Unix()
			//
			//// Update websocket clients only if we have updated
			//b, err := json.Marshal(&metric)
			//if err == nil {
			//	s.liveServer.Send(b)
			//}
			//
			//// Also notify any listeners of this new metric
			//s.listener.Notify(metric)
		}
	}
}
