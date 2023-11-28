package server

import (
	"github.com/peter-mount/go-kernel/v2/rest"
	api2 "github.com/peter-mount/piweather.center/store/api"
	"net/http"
	"time"
)

func (s *Server) latestMetrics(r *rest.Rest) error {
	req := GetRequest(r)

	var metrics []api2.Metric
	var t time.Time

	for _, metric := range s.Latest.Metrics() {
		if req.Match(metric) {
			if rec, exists := s.Latest.Latest(metric); exists {
				v := rec.Value
				metrics = append(metrics, api2.Metric{
					Metric:    metric,
					Time:      rec.Time,
					Unit:      v.Unit().ID(),
					Value:     v.Float(),
					Formatted: v.String(),
					Unix:      rec.Time.Unix(),
				})
				if rec.Time.After(t) {
					t = rec.Time
				}
			}
		}
	}

	response := req.Response()
	if len(metrics) > 0 {
		response.Time = &t
		response.Status = http.StatusOK
		response.Metrics = metrics
	} else {
		response.Status = http.StatusNotFound
	}

	return response.Submit(r)
}

func (s *Server) latestMetric(r *rest.Rest) error {
	req := GetRequest(r)

	response := req.Response()

	rec, exists := s.Latest.Latest(req.Metric)

	if exists {
		v := rec.Value
		response.Status = 200
		response.Result = &api2.MetricValue{
			Time:  rec.Time,
			Unit:  v.Unit().ID(),
			Value: v.Float(),
		}
	} else {
		response.Status = 404
	}

	return response.Submit(r)
}
