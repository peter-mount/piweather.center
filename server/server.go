package server

import (
	"context"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	_ "github.com/peter-mount/piweather.center/server/menu"
	_ "github.com/peter-mount/piweather.center/server/view"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/util/mq"
	"github.com/peter-mount/piweather.center/util/template"
	"github.com/peter-mount/piweather.center/weather/store"
	"path/filepath"
	"time"
)

// Server represents the primary service running the weather station.
type Server struct {
	Rest      *rest.Server                `kernel:"inject"`
	Amqp      mq.Pool                     `kernel:"inject"`
	Config    *map[string]station.Station `kernel:"config,stations"`
	Templates *template.Manager           `kernel:"inject"`
	Store     *store.Store                `kernel:"inject"`
}

func (s *Server) Start() error {
	for id, stationConfig := range *s.Config {
		stationConfig.ID = id
		ctx := context.WithValue(context.Background(), "Station", stationConfig)
		ctx = context.WithValue(ctx, "Store", s.Store)
		if err := stationConfig.Init(ctx); err != nil {
			return err
		}
	}

	rootDir := filepath.Dir(s.Templates.GetRootDir())
	staticDir := filepath.Join(rootDir, "static")
	log.Printf("Static content: %s", staticDir)
	s.Rest.Static("/static", staticDir)

	return s.startBrokers()
}

func (s *Server) startBrokers() error {
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

					return sensor.Process(p.AddContext(s.Store.AddContext(context.Background())))
				}); err != nil {
					return err
				}
			}
		}
	}
	log.Printf("Config: %v", s.Config)
	return nil
}
