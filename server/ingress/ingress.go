package ingress

import (
	"context"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/homeassistant"
	"github.com/peter-mount/piweather.center/influxdb"
	mq "github.com/peter-mount/piweather.center/mq/amqp"
	"github.com/peter-mount/piweather.center/mq/mqtt"
	"github.com/peter-mount/piweather.center/server/api"
	_ "github.com/peter-mount/piweather.center/server/menu"
	_ "github.com/peter-mount/piweather.center/server/view"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/station/payload"
	"github.com/peter-mount/piweather.center/station/service"
	"github.com/peter-mount/piweather.center/store"
	"github.com/peter-mount/piweather.center/weather/value"
	"io"
	"net/http"
)

// Ingress handles the ability to get data into the system, be it via
// http, amqp, mqtt etc.
type Ingress struct {
	Archiver        *store.Archiver        `kernel:"inject"`
	Amqp            mq.Pool                `kernel:"inject"`
	Mqtt            mqtt.Pool              `kernel:"inject"`
	EndpointManager *api.EndpointManager   `kernel:"inject"`
	Config          service.Config         `kernel:"inject"`
	Store           *store.Store           `kernel:"inject"`
	HomeAssistant   homeassistant.Service  `kernel:"inject"`
	InfluxDB        influxdb.Pool          `kernel:"inject"`
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
		CalculatedValue(s.Store.Calculate).
		Output(s.HomeAssistant.StoreReading).
		Output(s.InfluxDB.StoreReading)

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

	if err := queue.Bind(broker); err != nil {
		return err
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
			switch {
			case err != nil:
				log.Println(err)
				return err

			case p == nil:
				return nil

			default:
				return s.processVisitor.
					WithContext(p.AddContext(s.subContext)).
					VisitSensors(sensor)
			}
		})

	if err == nil {
		err = queue.Start(sensor.ID, false, mq.ContextTask(task, context.Background()))
	}
	return err
}

func (s *Ingress) startMQTT(ctx context.Context) error {
	sensor := station.SensorsFromContext(ctx)
	if sensor.Source.Amqp == nil {
		return nil
	}

	queue := sensor.Source.Mqtt

	broker := s.Mqtt.GetMQ(queue.Broker)
	if broker == nil {
		return fmt.Errorf("no broker %q defined for %s", queue.Broker, sensor.ID)
	}

	task, err := s.EndpointManager.RegisterEndpoint(
		"mqtt",
		queue.Broker+":"+queue.Topic,
		sensor.ID,
		sensor.Name,
		"MQTT",
		sensor.Format,
		func(ctx context.Context) error {
			msg := mqtt.Delivery(ctx)

			p, err := payload.FromMQTT(sensor.ID, sensor.Format, sensor.Timestamp, msg)
			switch {
			case err != nil:
				log.Println(err)
				return err

			case p == nil:
				return nil

			default:
				return s.processVisitor.
					WithContext(p.AddContext(s.subContext)).
					VisitSensors(sensor)
			}
		})

	queue.AddHandler(mqtt.ContextTask(task, context.Background()))
	if err == nil {
		err = queue.Bind(broker)
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
			switch {
			case err != nil:
				log.Println(err)
				return err

			case p == nil:
				return nil

			default:
				return s.processVisitor.WithContext(p.AddContext(s.subContext)).
					VisitSensors(sensor)
			}
		})
}
