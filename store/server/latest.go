package server

import (
	"github.com/peter-mount/go-kernel/v2/rest"
	api2 "github.com/peter-mount/piweather.center/store/api"
	"sort"
	"strings"
)

func (s *Server) latestMetrics(r *rest.Rest) error {

	var metrics []api2.Metric

	keys := s.Latest.Metrics()
	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, metric := range keys {
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
	metric := strings.ReplaceAll(r.Var(METRIC), "/", ".")

	response := api2.Response{Metric: metric}

	rec, exists := s.Latest.Latest(metric)

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
