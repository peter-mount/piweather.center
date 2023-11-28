package weathercenter

import (
	"github.com/peter-mount/go-build/version"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/mq/amqp"
	"github.com/peter-mount/piweather.center/store/broker"
	"github.com/peter-mount/piweather.center/store/memory"
	_ "github.com/peter-mount/piweather.center/tools/weathercenter/menu"
	"github.com/peter-mount/piweather.center/tools/weathercenter/template"
	"github.com/peter-mount/piweather.center/tools/weathercenter/view"
	"github.com/peter-mount/piweather.center/tools/weathercenter/ws"
	"path/filepath"
)

// Server represents the primary service running the fully integrated weather station.
type Server struct {
	Rest *rest.Server `kernel:"inject"`
	//Config    service.Config    `kernel:"inject"`
	_              *view.Home            `kernel:"inject"`
	Templates      *template.Manager     `kernel:"inject"`
	Latest         memory.Latest         `kernel:"inject"`
	Amqp           amqp.Pool             `kernel:"inject"`
	DatabaseBroker broker.DatabaseBroker `kernel:"inject"`
	QueueName      *string               `kernel:"flag,metric-queue,DB queue name,database.web"`
	DBServer       *string               `kernel:"flag,metric-db,DB url"`
	mqQueue        *amqp.Queue
	liveServer     *ws.Server
}

func (s *Server) Start() error {
	// Static content to the webserver
	rootDir := filepath.Dir(s.Templates.GetRootDir())
	staticDir := filepath.Join(rootDir, "static")
	log.Printf("Static content: %s", staticDir)
	s.Rest.Static("/static", staticDir)

	// Get latest metrics from DB
	if err := s.loadLatestMetrics(); err != nil {
		return err
	}

	// Websocket handler for live metrics
	s.liveServer = ws.NewServer()
	s.Rest.HandleFunc("/live", s.liveServer.Handle)
	go s.liveServer.Run()

	s.mqQueue = &amqp.Queue{
		Name:       *s.QueueName,
		Durable:    true,
		AutoDelete: false,
	}

	err := s.DatabaseBroker.ConsumeKeys(s.mqQueue, "ingress", s.recordMetricAmqp, "metric.#")

	if err == nil {
		log.Println(version.Version)
	}

	return err
}
