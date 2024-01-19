package server

import (
	"encoding/json"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/rabbitmq/amqp091-go"
	"sort"
)

// record implements the /record api
func (s *Server) record(r *rest.Rest) error {
	metric := api.Metric{}
	if err := r.Body(&metric); err != nil {
		return err
	}

	response := s.recordMetric(metric)

	return response.Submit(r)
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

	// ensure in time order
	sort.SliceStable(metrics, func(i, j int) bool {
		return metrics[i].Time.Before(metrics[j].Time)
	})

	// Collate by metric id
	m := make(map[string][]api.Metric)
	for _, metric := range metrics {
		m[metric.Metric] = append(m[metric.Metric], metric)
	}

	// Bulk import for each metric. This is several orders of magnitude faster than just recording each one in sequence
	for k, v := range m {
		response = s.recordMetricBulk(k, v, response)
	}

	return response.Submit(r)
}

func (s *Server) recordMetricBulk(n string, metrics []api.Metric, resp api.Response) api.Response {
	defer s.Store.Sync(n)
	for _, metric := range metrics {
		return s.recordMetricImpl(metric, true)
	}
	return resp
}

func (s *Server) recordMetric(metric api.Metric) api.Response {
	return s.recordMetricImpl(metric, false)
}

func (s *Server) recordMetricImpl(metric api.Metric, bulk bool) api.Response {

	if metric.Time.IsZero() {
		return api.Response{
			Status:  500,
			Message: "Invalid Metric",
			Metric:  metric.Metric,
		}
	}

	unit, exists := value.GetUnit(metric.Unit)
	if !exists {
		return api.Response{
			Status:  500,
			Message: "Unknown unit",
			Metric:  metric.Metric,
			Source:  metric.Unit,
		}
	}

	val := unit.Value(metric.Value)
	if !val.IsValid() {
		return api.Response{
			Status:  500,
			Message: "Value invalid for unit",
			Metric:  metric.Metric,
			Source:  metric.Unit,
		}
	}

	rec := record.Record{
		Time:  metric.Time,
		Value: val,
	}

	var err error
	if bulk {
		err = s.Store.AppendBulk(metric.Metric, rec)
	} else {
		err = s.Store.Append(metric.Metric, rec)
	}

	if err != nil {
		return api.Response{
			Status:  500,
			Message: "Failed to store",
			Metric:  metric.Metric,
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
