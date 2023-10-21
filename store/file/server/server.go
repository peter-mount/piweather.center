package server

import (
	"github.com/peter-mount/go-build/version"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/store/file"
)

type Server struct {
	Web    *rest.Server `kernel:"inject"`
	Store  file.Store   `kernel:"inject"`
	Latest *file.Latest `kernel:"inject"`
}

const (
	METRIC        = "metric"
	metricPrefix  = "/metric"
	metricPattern = "/{" + METRIC + ":.{1,}}"
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

	s.Web.Handle("/latest", s.latestMetrics).Methods(GET)

	s.Web.Handle(metricPrefix+metricPattern, s.latestMetric).Queries("latest", "").Methods(GET)
	s.Web.Handle(metricPrefix+metricPattern, s.queryToday).Queries("today", "").Methods(GET)
	s.Web.Handle(metricPrefix+metricPattern, s.queryTodayUTC).Queries("todayUTC", "").Methods(GET)

	s.Web.Handle("/record", s.record).Methods(POST)
	s.Web.Handle("/recordMultiple", s.recordMultiple).Methods(POST)

	return nil
}

func (s *Server) Start() error {
	log.Println(version.Version)
	return nil
}

func (s *Server) Stop() {
}
