package server

import (
	"github.com/peter-mount/go-build/version"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/mq/amqp"
	"github.com/peter-mount/piweather.center/store/broker"
	"github.com/peter-mount/piweather.center/store/file"
	"github.com/peter-mount/piweather.center/store/memory"
)

type Server struct {
	Web            *rest.Server          `kernel:"inject"`
	Amqp           amqp.Pool             `kernel:"inject"`
	Store          file.Store            `kernel:"inject"`
	Latest         memory.Latest         `kernel:"inject"`
	DatabaseBroker broker.DatabaseBroker `kernel:"inject"`
	QueueName      *string               `kernel:"flag,metric-queue,DB queue name,database.ingress"`
	mqQueue        *amqp.Queue
}

const (
	METRIC        = "metric"
	metricPrefix  = "/metric"
	metricPattern = "/{" + METRIC + ":.{1,}}"
	AT            = "at"
	FROM          = "from"
	TO            = "to"
	LATEST        = "latest"
	TODAY         = "today"
	TODAYUTC      = "todayUTC"
	FILTER        = "filter"
	POST          = "POST"
	GET           = "GET"
)

func (s *Server) Init(_ kernel.Kernel) error {
	if s.Web.Port == 8080 {
		s.Web.Port = 9001
	}
	return nil
}

func (s *Server) PostInit() error {

	if *s.QueueName == "" {
		*s.QueueName = "database.ingress"
	}

	// /latest returns the latest values of all metrics
	s.Web.Handle(metricPrefix, s.latestMetrics).Queries(LATEST, "").Methods(GET)
	s.Web.Handle(metricPrefix, s.queryAllAt).Queries(AT, "").Methods(GET)

	// queries against an individual metric
	s.Web.Handle(metricPrefix+metricPattern, s.latestMetric).Queries(LATEST, "").Methods(GET)
	s.Web.Handle(metricPrefix+metricPattern, s.queryToday).Queries(TODAY, "").Methods(GET)
	s.Web.Handle(metricPrefix+metricPattern, s.queryTodayUTC).Queries(TODAYUTC, "").Methods(GET)
	s.Web.Handle(metricPrefix+metricPattern, s.queryAt).Queries(AT, "").Methods(GET)
	s.Web.Handle(metricPrefix+metricPattern, s.queryBetween).Queries(FROM, "", TO, "").Methods(GET)

	// record a metric over http - not normally used as amqp is normally used
	s.Web.Handle("/record", s.record).Methods(POST)
	s.Web.Handle("/recordMultiple", s.recordMultiple).Methods(POST)

	return nil
}

func (s *Server) Start() error {

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

func (s *Server) Stop() {
	if s.mqQueue != nil {
		s.mqQueue.Stop()
		s.mqQueue = nil
	}
}
