package server

import (
	"github.com/peter-mount/go-build/version"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/mq/amqp"
	"github.com/peter-mount/piweather.center/store/file"
	"github.com/peter-mount/piweather.center/store/memory"
)

type Server struct {
	Web     *rest.Server  `kernel:"inject"`
	Amqp    amqp.Pool     `kernel:"inject"`
	Store   file.Store    `kernel:"inject"`
	Latest  memory.Latest `kernel:"inject"`
	mq      *amqp.MQ
	mqQueue *amqp.Queue
}

const (
	METRIC        = "metric"
	metricPrefix  = "/metric"
	metricPattern = "/{" + METRIC + ":.{1,}}"
	POST          = "POST"
	GET           = "GET"
	brokerName    = "database"
)

func (s *Server) Init(_ kernel.Kernel) error {
	if s.Web.Port == 8080 {
		s.Web.Port = 9001
	}
	return nil
}

func (s *Server) PostInit() error {

	s.Web.Handle("/latest", s.latestMetrics).Methods(GET)

	s.Web.Handle(metricPrefix+metricPattern, s.latestMetric).Queries("latest", "").Methods(GET)
	s.Web.Handle(metricPrefix+metricPattern, s.queryToday).Queries("today", "").Methods(GET)
	s.Web.Handle(metricPrefix+metricPattern, s.queryTodayUTC).Queries("todayUTC", "").Methods(GET)

	s.Web.Handle("/record", s.record).Methods(POST)
	s.Web.Handle("/recordMultiple", s.recordMultiple).Methods(POST)

	return nil
}

func (s *Server) Start() error {
	s.mq = s.Amqp.GetMQ(brokerName)
	if s.mq == nil {
		return nil
	}

	s.mq.ConnectionName = "PIWeatherCenter Database"
	s.mq.Version = version.Version

	s.mqQueue = &amqp.Queue{
		Name: "database.ingress",
		Binding: []amqp.Binding{
			{
				Topic: s.mq.Exchange,
				Key:   "metric.#",
			},
		},
		Durable:    true,
		AutoDelete: false,
	}

	err := s.mqQueue.Bind(s.mq)
	if err == nil {
		err = s.mqQueue.Start("database ingress", false, s.recordMetricAmqp)
	}

	if err == nil {
		log.Println(version.Version)
	}

	return err
}

func (s *Server) Stop() {
	if s.mqQueue != nil {
		s.mqQueue.Stop()
		s.mqQueue = nil
	}
}
