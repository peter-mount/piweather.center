package server

import (
	"encoding/json"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/rabbitmq/amqp091-go"
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
		Value(response.Sort())

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

	err := s.Store.Append(metric.Metric, record.Record{
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

// recordMetricAmqp accepts a metric from RabbitMQ and passes it to recordMetric
func (s *Server) recordMetricAmqp(delivery amqp091.Delivery) error {
	var metric api.Metric
	err := json.Unmarshal(delivery.Body, &metric)
	if err == nil {
		resp := s.recordMetric(metric)
		// Allow 2xx for successful status codes
		if (resp.Status / 100) != 2 {
			err = fmt.Errorf("recordMetricAmqp status %d %s %q", resp.Status, resp.Message, resp.Source)
		}
	}
	return err
}
