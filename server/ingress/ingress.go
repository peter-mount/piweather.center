package ingress

import (
	"context"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	mq "github.com/peter-mount/piweather.center/mq/amqp"
	"github.com/peter-mount/piweather.center/server/api"
	"github.com/peter-mount/piweather.center/server/archiver"
	_ "github.com/peter-mount/piweather.center/server/menu"
	"github.com/peter-mount/piweather.center/server/store"
	_ "github.com/peter-mount/piweather.center/server/view"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/station/payload"
	"github.com/peter-mount/piweather.center/weather/value"
	"io"
	"net/http"
)

// Ingress handles the ability to get data into the system, be it via
// http, amqp, mqtt etc.
type Ingress struct {
	Archiver        *archiver.Archiver     `kernel:"inject"`
	Amqp            mq.Pool                `kernel:"inject"`
	EndpointManager *api.EndpointManager   `kernel:"inject"`
	Config          station.Config         `kernel:"inject"`
	Store           *store.Store           `kernel:"inject"`
	subContext      context.Context        // Common Context
	processVisitor  station.VisitorBuilder // Common visitor used by all sources to process data
}

func (s *Ingress) Start() error {
	// Common context for processing
	s.subContext = s.Archiver.AddContext(s.Store.AddContext(value.WithMap(context.Background())))

	// Visitor that will process an inbound message.
	// This is common to all sources, so we define it here, but they will
	// build it as needed.
	s.processVisitor = station.NewVisitor().
		Sensors(value.ResetMap).
		Sensors(s.Archiver.Archive).
		Reading(s.Store.ProcessReading).
		CalculatedValue(s.Store.Calculate)

	// Now we preload data from storage to give us some recent history
	if err := s.Config.Accept(station.NewVisitor().
		Sensors(value.ResetMap).
		Sensors(s.Archiver.Preload).
		WithContext(s.subContext)); err != nil {
		return err
	}

	// Now start the sources
	if err := s.Config.Accept(station.NewVisitor().
		Sensors(s.startAMQP).
		Sensors(s.startEcowitt).
		WithContext(s.subContext)); err != nil {
		return err
	}

	return nil
}

func (s *Ingress) startAMQP(ctx context.Context) error {
	sensor := station.SensorsFromContext(ctx)
	if sensor.Source.Amqp == nil {
		return nil
	}

	queue := sensor.Source.Amqp

	broker := s.Amqp.GetMQ(queue.Broker)
	if broker == nil {
		return fmt.Errorf("no broker %q defined for %s", queue.Broker, sensor.ID)
	}

	task, err := s.EndpointManager.RegisterEndpoint(
		"amqp",
		queue.Broker+":"+queue.Name,
		sensor.ID,
		sensor.Name,
		"AMQP",
		sensor.Format,
		func(ctx context.Context) error {
			msg := mq.Delivery(ctx)

			p, err := payload.FromAMQP(sensor.ID, sensor.Format, sensor.Timestamp, msg)
			if err != nil {
				log.Println(err)
				return err
			}

			return s.processVisitor.
				WithContext(p.AddContext(s.subContext)).
				VisitSensors(sensor)
		})

	if err == nil {
		err = broker.ConsumeTask(queue, sensor.ID, task)
	}
	return err
}

func (s *Ingress) startEcowitt(ctx context.Context) error {
	sensor := station.SensorsFromContext(ctx)
	if sensor.Source.EcoWitt == nil {
		return nil
	}

	return s.EndpointManager.RegisterHttpEndpoint(
		"inbound",
		"/api/inbound/"+sensor.Source.EcoWitt.Path,
		sensor.ID,
		sensor.Name,
		http.MethodPost,
		"ecowitt",
		func(ctx context.Context) error {
			r := rest.GetRest(ctx)
			body, _ := io.ReadAll(r.Request().Body)

			p, err := payload.FromBytes(sensor.ID, sensor.Format, sensor.Timestamp, body)
			if err != nil {
				log.Println(err)
				return err
			}

			return s.processVisitor.WithContext(p.AddContext(s.subContext)).
				VisitSensors(sensor)
		})
}
