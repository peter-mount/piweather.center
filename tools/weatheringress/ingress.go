package weatheringress

import (
	"context"
	"github.com/peter-mount/go-build/version"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/store/broker"
	"github.com/peter-mount/piweather.center/tools/weatherarchive"
	"github.com/peter-mount/piweather.center/tools/weatheringress/model"
	"github.com/peter-mount/piweather.center/util/endpoint"
	"github.com/peter-mount/piweather.center/util/mq/amqp"
	"github.com/peter-mount/piweather.center/util/mq/mqtt"
)

// Ingress handles the ability to get data into the system, be it via
// http, amqp, mqtt etc.
type Ingress struct {
	Archiver        *weatherarchive.Archiver  `kernel:"inject"`
	Amqp            amqp.Pool                 `kernel:"inject"`
	Mqtt            mqtt.Pool                 `kernel:"inject"`
	EndpointManager *endpoint.EndpointManager `kernel:"inject"`
	Loader          model.Loader              `kernel:"inject"`
	DatabaseBroker  broker.DatabaseBroker     `kernel:"inject"`
	subContext      context.Context           // Common Context
	processVisitor  model.VisitorBuilder      // Common visitor used by all sources to process data
}

func (s *Ingress) Start() error {
	// Originally this was set up with specific entries in it which are no longer present.
	// It's kept in case we need this again in the future, so for now it's the default context.
	s.subContext = context.Background()

	// Visitor that will process an inbound message.
	// This is common to all sources, so we define it here, but they will
	// build it as needed.
	s.processVisitor = model.NewVisitor().
		Sensors(s.Archiver.Archive).
		Reading(s.processReading)

	// Now start the sources
	if err := s.Loader.Accept(model.NewVisitor().
		Sensors(s.startAMQP).
		Sensors(s.startHttp).
		WithContext(s.subContext)); err != nil {
		return err
	}

	log.Println(version.Version)
	return nil
}
