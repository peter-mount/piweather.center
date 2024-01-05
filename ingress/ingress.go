package ingress

import (
	"context"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	mq "github.com/peter-mount/piweather.center/mq/amqp"
	"github.com/peter-mount/piweather.center/mq/mqtt"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/station/payload"
	"github.com/peter-mount/piweather.center/station/service"
	api2 "github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/broker"
	log2 "github.com/peter-mount/piweather.center/store/log"
	_ "github.com/peter-mount/piweather.center/tools/weathercenter/menu"
	_ "github.com/peter-mount/piweather.center/tools/weathercenter/view"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/util/endpoint"
	"github.com/peter-mount/piweather.center/weather/value"
	"io"
	"net/http"
)

// Ingress handles the ability to get data into the system, be it via
// http, amqp, mqtt etc.
type Ingress struct {
	Archiver        *log2.Archiver            `kernel:"inject"`
	Amqp            mq.Pool                   `kernel:"inject"`
	Mqtt            mqtt.Pool                 `kernel:"inject"`
	EndpointManager *endpoint.EndpointManager `kernel:"inject"`
	Config          service.Config            `kernel:"inject"`
	DatabaseBroker  broker.DatabaseBroker     `kernel:"inject"`
	subContext      context.Context           // Common Context
	processVisitor  station.VisitorBuilder    // Common visitor used by all sources to process data
}

func (s *Ingress) Start() error {

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
					WithContext(p.AddContext(value.WithMap(s.subContext))).
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
					WithContext(p.AddContext(value.WithMap(s.subContext))).
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
				return s.processVisitor.WithContext(p.AddContext(value.WithMap(s.subContext))).
					VisitSensors(sensor)
			}
		})
}

func (s *Ingress) databaseReading(ctx context.Context) error {
	payloadEntry := payload.GetPayload(ctx)

	metric := api2.Metric{Time: payloadEntry.Time().UTC()}

	values := value.MapFromContext(ctx)
	for _, key := range values.GetKeys() {
		val := values.Get(key)

		if val.IsValid() {

			metric.Metric = key
			metric.Value = val.Float()
			metric.Unit = val.Unit().ID()

			err := s.DatabaseBroker.PublishMetric(metric)
			if err != nil {
				return err
			}
		} else {
			// Has happened, not sure if down to altered ID's etc
			log.Printf("Invalid Metric %q", key)
		}
	}

	return nil
}

func (s *Ingress) processReading(ctx context.Context) error {
	r := station.ReadingFromContext(ctx)
	values := value.MapFromContext(ctx)
	if r.Unit() != nil {
		p := payload.GetPayload(ctx)

		str, ok := p.Get(r.Source)
		if !ok {
			// FIXME warn/fail if not found?
			return nil
		}

		if f, ok := util.ToFloat64(str); ok {
			// Convert to Type unit then transform to Use unit
			v, err := r.Value(f)
			if err != nil {
				// Ignore, should only happen if the result is
				// invalid as we checked the transform previously
				return nil
			}

			values.Put(r.ID, v)
		}
	}
	return nil
}

func (s *Ingress) calculate(ctx context.Context) error {
	// Get value.Time from Station and Payload
	sensors := station.SensorsFromContext(ctx)
	p := payload.GetPayload(ctx)
	t := sensors.Station().LatLong().Time(p.Time())

	calc := station.CalculatedValueFromContext(ctx)

	values := value.MapFromContext(ctx)
	args := values.GetAll(calc.Source...)

	result, err := calc.Calculate(t, args...)
	if err != nil {
		return err
	}

	values.Put(calc.ID, result)

	return nil
}
