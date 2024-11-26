package weatheringress

import (
	"context"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/tools/weatheringress/model"
	"github.com/peter-mount/piweather.center/tools/weatheringress/payload"
	mq "github.com/peter-mount/piweather.center/util/mq/amqp"
	"github.com/peter-mount/piweather.center/util/mq/mqtt"
	"io"
	"net/http"
)

func (s *Ingress) startAMQP(ctx context.Context) error {
	sensor := model.SensorsFromContext(ctx)
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
	sensor := model.SensorsFromContext(ctx)
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

func (s *Ingress) startHttp(ctx context.Context) error {
	sensor := model.SensorsFromContext(ctx)
	if sensor.Source.Http == nil {
		return nil
	}

	var bodyFunc func(*rest.Rest) []byte
	switch sensor.Source.Http.Method {
	case http.MethodGet:
		bodyFunc = getBody
	case http.MethodPost:
		bodyFunc = postBody
	default:
		return fmt.Errorf("method %q is not valid", sensor.Source.Http.Method)
	}

	return s.EndpointManager.RegisterHttpEndpoint(
		"inbound",
		"/api/inbound/"+sensor.Source.Http.Path,
		sensor.ID,
		sensor.Name,
		sensor.Source.Http.Method,
		"http",
		func(ctx context.Context) error {
			r := rest.GetRest(ctx)

			body := bodyFunc(r)

			if sensor.Source.Http.Debug {
				log.Printf("%s:%s:%s %q", sensor.ID, sensor.Name, sensor.Source.Http.Method, string(body))
			}

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

func getBody(r *rest.Rest) []byte {
	return []byte(r.Request().URL.RawQuery)
}

func postBody(r *rest.Rest) []byte {
	body, _ := io.ReadAll(r.Request().Body)
	return body
}
