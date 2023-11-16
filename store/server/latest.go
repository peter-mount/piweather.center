package server

import (
	"github.com/peter-mount/go-kernel/v2/rest"
	api2 "github.com/peter-mount/piweather.center/store/api"
)

func (s *Server) latestMetrics(r *rest.Rest) error {
	req := getRequest(r)

	var metrics []api2.Metric

	for _, metric := range s.Latest.Metrics() {
		if req.Match(metric) {
			if rec, exists := s.Latest.Latest(metric); exists {
				v := rec.Value
				metrics = append(metrics, api2.Metric{
					Metric: metric,
					Time:   rec.Time,
					Unit:   v.Unit().ID(),
					Value:  v.Float(),
				})
			}
		}
	}

	response := api2.Response{
		Status:  200,
		Metrics: metrics,
	}

	r.Status(response.Status).
		ContentType(r.GetHeader("Content-Type")).
		Value(response.Sort())

	return nil
}

func (s *Server) latestMetric(r *rest.Rest) error {
	req := getRequest(r)

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

	r.Status(response.Status).
		ContentType(r.GetHeader("Content-Type")).
		Value(response.Sort())

	return nil
}
