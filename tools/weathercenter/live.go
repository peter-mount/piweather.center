package weathercenter

import (
	"encoding/json"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/client"
	"github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/weather/value"
)

// loadLatestMetrics retrieves the current metrics from the DB server
func (s *Server) loadLatestMetrics() error {
	if *s.DBServer != "" {
		c := &client.Client{Url: *s.DBServer}
		r, err := c.LatestMetrics()
		if err != nil {
			return err
		}
		if r != nil {
			for _, m := range r.Metrics {
				s.storeLatest(m)
			}
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

		if updated {
			metric.Formatted = u.String(metric.Value)
			metric.Unix = metric.Time.Unix()

			// Update websocket clients only if we have updated
			b, err := json.Marshal(&metric)
			if err == nil {
				s.liveServer.Send(b)
			}

			// Also notify any listeners of this new metric
			s.listener.Notify(metric)
		}
	}
}
