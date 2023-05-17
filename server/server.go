package server

import (
	"context"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/util/mq"
	"time"
)

// Server represents the primary service running the weather station.
type Server struct {
	Rest   *rest.Server                `kernel:"inject"`
	Amqp   mq.Pool                     `kernel:"inject"`
	Config *map[string]station.Station `kernel:"config,stations"`
}

func (s *Server) Start() error {
	if s.Config == nil || len(*s.Config) == 0 {
		return fmt.Errorf("no configuration provided")
	}

	for name, stationConfig := range *s.Config {
		for sensorName, sensor := range stationConfig.Sensors {
			if amqp := sensor.Source.Amqp; amqp != nil {
				broker := s.Amqp.GetMQ(amqp.Broker)
				if broker == nil {
					return fmt.Errorf("no broker %q defined for %s:%s", amqp.Broker, name, sensorName)
				}
				if err := broker.ConsumeTask(amqp, "tag", func(ctx context.Context) error {
					msg := mq.Delivery(ctx)
					log.Println(string(msg.Body))

					p, err := sensor.FromAMQP(msg)
					if err != nil {
						log.Println(err)
						return err
					}

					log.Println(p.Time().Format(time.RFC3339))
					return nil
				}); err != nil {
					return err
				}
			}
		}
	}
	log.Printf("Config: %v", s.Config)
	return nil
}
