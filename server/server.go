package server

import (
	"context"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/server/archiver"
	_ "github.com/peter-mount/piweather.center/server/menu"
	_ "github.com/peter-mount/piweather.center/server/view"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/station/payload"
	"github.com/peter-mount/piweather.center/util/mq"
	"github.com/peter-mount/piweather.center/util/template"
	"github.com/peter-mount/piweather.center/weather/store"
	"path/filepath"
)

// Server represents the primary service running the weather station.
type Server struct {
	Rest       *rest.Server                `kernel:"inject"`
	Archiver   *archiver.Archiver          `kernel:"inject"`
	Amqp       mq.Pool                     `kernel:"inject"`
	Config     *map[string]station.Station `kernel:"config,stations"`
	Templates  *template.Manager           `kernel:"inject"`
	Store      *store.Store                `kernel:"inject"`
	subContext context.Context             // Common Context
}

func (s *Server) Start() error {
	// Common context for processing
	s.subContext = s.Archiver.AddContext(s.Store.AddContext(context.Background()))

	for id, stationConfig := range *s.Config {
		stationConfig.ID = id
		ctx := context.WithValue(s.subContext, "Station", stationConfig)
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

	for _, stationConfig := range *s.Config {
		for sensorName, sensor := range stationConfig.Sensors {
			if amqp := sensor.Source.Amqp; amqp != nil {
				if err := s.startAMQP(sensorName, sensor, amqp); err != nil {
					return err
				}
			}
		}
	}
	log.Printf("Config: %v", s.Config)
	return nil
}

func (s *Server) startAMQP(sensorName string, sensor *station.Sensors, amqp *mq.Queue) error {
	broker := s.Amqp.GetMQ(amqp.Broker)
	if broker == nil {
		return fmt.Errorf("no broker %q defined for %s:%s", amqp.Broker, sensor.ID, sensorName)
	}
	return broker.ConsumeTask(amqp, "tag", func(ctx context.Context) error {
		msg := mq.Delivery(ctx)
		log.Println(string(msg.Body))

		p, err := payload.FromAMQP(sensor.ID, sensor.Format, sensor.Timestamp, msg)
		if err != nil {
			log.Println(err)
			return err
		}

		return sensor.Process(p.AddContext(s.subContext))
	})
}
