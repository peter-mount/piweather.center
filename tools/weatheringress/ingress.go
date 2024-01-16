package weatheringress

import (
	"context"
	"github.com/peter-mount/go-build/version"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/mq/amqp"
	"github.com/peter-mount/piweather.center/mq/mqtt"
	"github.com/peter-mount/piweather.center/store/broker"
	log2 "github.com/peter-mount/piweather.center/store/log"
	"github.com/peter-mount/piweather.center/tools/weatheringress/model"
	"github.com/peter-mount/piweather.center/tools/weatheringress/service"
	"github.com/peter-mount/piweather.center/util/endpoint"
)

// Ingress handles the ability to get data into the system, be it via
// http, amqp, mqtt etc.
type Ingress struct {
	Archiver *log2.Archiver `kernel:"inject"`
	//Latest          memory.Latest             `kernel:"inject"`
	Amqp            amqp.Pool                 `kernel:"inject"`
	Mqtt            mqtt.Pool                 `kernel:"inject"`
	EndpointManager *endpoint.EndpointManager `kernel:"inject"`
	Config          service.Config            `kernel:"inject"`
	DatabaseBroker  broker.DatabaseBroker     `kernel:"inject"`
	//QueueName       *string                   `kernel:"flag,metric-queue,DB queue name,database.ingress"`
	//DBServer        *string                   `kernel:"flag,metric-db,DB url"`
	//mqQueue         *amqp.Queue
	subContext     context.Context      // Common Context
	processVisitor model.VisitorBuilder // Common visitor used by all sources to process data
	updateVisitor  model.Visitor        // Visitor used to process updates sent to the DB to handle calculated updates
}

func (s *Ingress) Start() error {

	// Common context for processing
	s.subContext = s.Archiver.AddContext(context.Background())

	// Visitor that will process an inbound message.
	// This is common to all sources, so we define it here, but they will
	// build it as needed.
	s.processVisitor = model.NewVisitor().
		Sensors(s.Archiver.Archive).
		Reading(s.processReading)

	// Now start the sources
	if err := s.Config.Accept(model.NewVisitor().
		Sensors(s.startAMQP).
		Sensors(s.startHttp).
		WithContext(s.subContext)); err != nil {
		return err
	}

	log.Println(version.Version)
	return nil
}
