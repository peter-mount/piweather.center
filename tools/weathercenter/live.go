package weathercenter

import (
	"encoding/json"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/client"
	"github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/rabbitmq/amqp091-go"
)

// recordMetricAmqp accepts a metric from RabbitMQ, stores it in Latest
// then forwards it to any websocket clients
func (s *Server) recordMetricAmqp(delivery amqp091.Delivery) error {
	var metric api.Metric
	err := json.Unmarshal(delivery.Body, &metric)
	if err == nil {
		s.storeLatest(metric)
	}
	return err
}

// loadLatestMetrics retrieves the current metrics from the DB server
func (s *Server) loadLatestMetrics() error {
	if *s.DBServer != "" {
		c := &client.Client{Url: *s.DBServer}
		r, err := c.LatestMetrics()
		if err != nil {
			return err
		}
		for _, m := range r.Metrics {
			s.storeLatest(m)
		}
	}
	return nil
}

func (s *Server) storeLatest(metric api.Metric) {
	u, ok := value.GetUnit(metric.Unit)
	if ok {
		updated := s.Latest.Append(metric.Metric, record.Record{
			Time:  metric.Time,
			Value: u.Value(metric.Value),
		})

		// Update websocket clients only if we have updated
		if updated {
			metric.Formatted = u.String(metric.Value)
			metric.Unix = metric.Time.Unix()
			b, err := json.Marshal(&metric)
			if err == nil {
				s.liveServer.Send(b)
			}
		}
	}
}
