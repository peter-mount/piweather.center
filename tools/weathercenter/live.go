package weathercenter

import (
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/weather/value"
)

// loadLatestMetrics retrieves the current metrics from the DB server
func (s *Server) loadLatestMetrics() error {
	//if *s.DBServer != "" {
	//	c := &client.Client{Url: *s.DBServer}
	//	r, err := c.LatestMetrics()
	//	if err != nil {
	//		return err
	//	}
	//	if r != nil {
	//		for _, m := range r.Metrics {
	//			s.storeLatest(m)
	//		}
	//	}
	//}
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

			// Notify the station for this metric
			responses := s.ViewService.Stations.Notify(metric)

			// Send any responses (one per dashboard the metric was used on) to the appropriate clients
			if len(responses) > 0 {
				for _, response := range responses {
					live := s.ViewService.GetLive(response.Station, response.Dashboard)
					if live != nil {
						live.Notify(response)
					}
				}
			}
		}
	}
}
