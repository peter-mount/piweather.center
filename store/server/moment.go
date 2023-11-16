package server

import (
	api2 "github.com/peter-mount/piweather.center/store/api"
	"time"
)

// queryMetricBetween returns the values of a metric between two times.
// This will search up to the previous 6 hours for those metrics which are
// not regularly submitted.
func (s *Server) queryMetricBetween(metric string, from, to time.Time, r *api2.Response) {
	query := s.Store.Query(metric).
		Between(from, to).
		Build()

	for query.HasNext() {
		rec := query.Next()
		if rec.IsValid() {
			val := rec.Value
			r.Metrics = append(r.Metrics, api2.Metric{
				Metric: metric,
				Time:   rec.Time,
				Unit:   val.Unit().ID(),
				Value:  val.Float(),
			})
		}
	}
}
