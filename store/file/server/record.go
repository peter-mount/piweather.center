package server

import (
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/store/file"
	"github.com/peter-mount/piweather.center/store/file/api"
	"github.com/peter-mount/piweather.center/weather/value"
)

// record implements the /record api
func (s *Server) record(r *rest.Rest) error {
	metric := api.Metric{}
	if err := r.Body(&metric); err != nil {
		return err
	}

	response := s.recordMetric(metric)

	r.Status(response.Status).
		ContentType(r.GetHeader("Content-Type")).
		Value(response)

	return nil
}

func (s *Server) recordMultiple(r *rest.Rest) error {
	var metrics []api.Metric
	if err := r.Body(&metrics); err != nil {
		return err
	}

	response := api.Response{
		Status:  404,
		Message: "Nothing imported",
	}

	for _, metric := range metrics {
		response = s.recordMetric(metric)
	}

	r.Status(response.Status).
		ContentType(r.GetHeader("Content-Type")).
		Value(response)

	return nil
}

func (s *Server) recordMetric(metric api.Metric) api.Response {

	if metric.Time.IsZero() {
		return api.Response{
			Status:  500,
			Message: "Invalid Metric",
		}
	}

	unit, exists := value.GetUnit(metric.Unit)
	if !exists {
		return api.Response{
			Status:  500,
			Message: "Unknown unit",
			Source:  metric.Unit,
		}
	}

	val := unit.Value(metric.Value)
	if !val.IsValid() {
		return api.Response{
			Status:  500,
			Message: "Value invalid for unit",
			Source:  metric.Unit,
		}
	}

	err := s.Store.Append(metric.Metric, file.Record{
		Time:  metric.Time,
		Value: val,
	})
	if err != nil {
		return api.Response{
			Status:  500,
			Message: "Failed to store",
			Source:  err.Error(),
		}
	}

	return api.Response{Status: 200}
}
