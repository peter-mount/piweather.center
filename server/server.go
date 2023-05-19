package server

import (
	"context"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/server/archiver"
	"github.com/peter-mount/piweather.center/server/ecowitt"
	_ "github.com/peter-mount/piweather.center/server/menu"
	"github.com/peter-mount/piweather.center/server/store"
	_ "github.com/peter-mount/piweather.center/server/view"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/station/payload"
	"github.com/peter-mount/piweather.center/util/mq"
	"github.com/peter-mount/piweather.center/util/template"
	"io"
	"path/filepath"
)

// Server represents the primary service running the weather station.
type Server struct {
	Rest       *rest.Server       `kernel:"inject"`
	Archiver   *archiver.Archiver `kernel:"inject"`
	Amqp       mq.Pool            `kernel:"inject"`
	Ecowitt    *ecowitt.Server    `kernel:"inject"`
	Config     *station.Stations  `kernel:"config,stations"`
	Templates  *template.Manager  `kernel:"inject"`
	Store      *store.Store       `kernel:"inject"`
	subContext context.Context    // Common Context
}

func (s *Server) Start() error {
	if s.Config == nil || len(*s.Config) == 0 {
		return fmt.Errorf("no configuration provided")
	}

	// Static content to the webserver
	rootDir := filepath.Dir(s.Templates.GetRootDir())
	staticDir := filepath.Join(rootDir, "static")
	log.Printf("Static content: %s", staticDir)
	s.Rest.Static("/static", staticDir)

	// Common context for processing
	s.subContext = s.Archiver.AddContext(s.Store.AddContext(context.Background()))

	// Initialise the station config so ID's etc. are correct
	if err := station.NewVisitor().
		Station(station.InitStation).
		Sensors(station.InitSensors).
		Reading(station.InitReading).
		WithContext(s.subContext).
		VisitStations(s.Config); err != nil {
		return err
	}

	// Visitor that will process an inbound message.
	// This is common to all sources, so we define it here, but they will
	// build it as needed.
	processVisitor := station.NewVisitor().
		Sensors(s.Archiver.Archive).
		Reading(station.ProcessReading)

	// Now we preload data from storage to give us some recent
	// history, then we start each data source, so we can get fresh data sent to us.
	if err := station.NewVisitor().
		Sensors(s.Archiver.Preload).
		Sensors(s.startAMQP).
		Sensors(s.startEcowitt).
		WithContext(context.WithValue(s.subContext, "processVisitor", processVisitor)).
		VisitStations(s.Config); err != nil {
		return err
	}

	return nil
}

func (s *Server) startAMQP(ctx context.Context) error {
	sensor := station.SensorsFromContext(ctx)
	if sensor.Source.Amqp == nil {
		return nil
	}

	amqp := sensor.Source.Amqp

	broker := s.Amqp.GetMQ(amqp.Broker)
	if broker == nil {
		return fmt.Errorf("no broker %q defined for %s", amqp.Broker, sensor.ID)
	}

	processVisitor := ctx.Value("processVisitor").(station.VisitorBuilder)

	return broker.ConsumeTask(amqp, "tag", func(ctx context.Context) error {
		msg := mq.Delivery(ctx)

		p, err := payload.FromAMQP(sensor.ID, sensor.Format, sensor.Timestamp, msg)
		if err != nil {
			log.Println(err)
			return err
		}

		return processVisitor.WithContext(p.AddContext(s.subContext)).
			VisitSensors(sensor)
	})
}

func (s *Server) startEcowitt(ctx context.Context) error {
	sensor := station.SensorsFromContext(ctx)
	if sensor.Source.EcoWitt == nil {
		return nil
	}
	processVisitor := ctx.Value("processVisitor").(station.VisitorBuilder)

	return s.Ecowitt.RegisterEndpoint(sensor, func(ctx context.Context) error {
		r := rest.GetRest(ctx)
		body, _ := io.ReadAll(r.Request().Body)

		p, err := payload.FromBytes(sensor.ID, sensor.Format, sensor.Timestamp, body)
		if err != nil {
			log.Println(err)
			return err
		}

		return processVisitor.WithContext(p.AddContext(s.subContext)).
			VisitSensors(sensor)
	})
}
