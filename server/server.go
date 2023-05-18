package server

import (
	"context"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/server/archiver"
	"github.com/peter-mount/piweather.center/server/ecowitt"
	_ "github.com/peter-mount/piweather.center/server/menu"
	_ "github.com/peter-mount/piweather.center/server/view"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/station/payload"
	"github.com/peter-mount/piweather.center/util/mq"
	"github.com/peter-mount/piweather.center/util/template"
	"github.com/peter-mount/piweather.center/weather/store"
	"io"
	"path/filepath"
)

// Server represents the primary service running the weather station.
type Server struct {
	Rest       *rest.Server                `kernel:"inject"`
	Archiver   *archiver.Archiver          `kernel:"inject"`
	Amqp       mq.Pool                     `kernel:"inject"`
	Ecowitt    *ecowitt.Server             `kernel:"inject"`
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
		for _, sensor := range stationConfig.Sensors {
			var err error
			switch {
			case sensor.Source.Amqp != nil:
				err = s.startAMQP(sensor)

			case sensor.Source.EcoWitt != nil:
				err = s.startEcowitt(sensor)
			}
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Server) startAMQP(sensor *station.Sensors) error {
	amqp := sensor.Source.Amqp

	broker := s.Amqp.GetMQ(amqp.Broker)
	if broker == nil {
		return fmt.Errorf("no broker %q defined for %s:%s", amqp.Broker, sensor.ID)
	}

	return broker.ConsumeTask(amqp, "tag", func(ctx context.Context) error {
		msg := mq.Delivery(ctx)

		p, err := payload.FromAMQP(sensor.ID, sensor.Format, sensor.Timestamp, msg)
		if err != nil {
			log.Println(err)
			return err
		}

		return sensor.Process(p.AddContext(s.subContext))
	})
}

func (s *Server) startEcowitt(sensor *station.Sensors) error {
	return s.Ecowitt.RegisterEndpoint(sensor, func(ctx context.Context) error {
		r := rest.GetRest(ctx)
		body, _ := io.ReadAll(r.Request().Body)

		p, err := payload.FromBytes(sensor.ID, sensor.Format, sensor.Timestamp, body)
		if err != nil {
			log.Println(err)
			return err
		}

		return sensor.Process(p.AddContext(s.subContext))
	})
}
