package server

import (
	"github.com/peter-mount/go-kernel/v2/rest"
	api2 "github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/file"
	"net/http"
	"time"
)

func (s *Server) queryMetricToday(r *rest.Rest) error {
	return s.queryMetric(r, r.Var(METRIC), func(b file.QueryBuilder) {
		b.Today()
	})
}

func (s *Server) queryMetricTodayUTC(r *rest.Rest) error {
	return s.queryMetric(r, r.Var(METRIC), func(b file.QueryBuilder) {
		b.TodayUTC()
	})
}

// queryMetric returns the values of a single metric based on a query
func (s *Server) queryMetric(r *rest.Rest, metric string, qbf func(file.QueryBuilder)) error {

	qb := s.Store.Query(metric)
	qbf(qb)
	query := qb.Build()

	response := api2.Response{
		Metric: metric,
	}

	for query.HasNext() {
		rec := query.Next()
		val := rec.Value
		response.Results = append(response.Results, api2.MetricValue{
			Time:  rec.Time,
			Unit:  val.Unit().ID(),
			Value: val.Float(),
		})
	}

	if len(response.Results) > 0 {
		response.Status = http.StatusOK
	} else {
		response.Status = http.StatusNotFound
	}

	return response.Submit(r)
}

// queryMetricAt handles a single metric with the at query parameter
func (s *Server) queryMetricAt(r *rest.Rest) error {
	req := GetRequest(r)

	if req.At.IsZero() {
		r.Status(http.StatusBadRequest)
		return nil
	}
	from := req.At.Add(-time.Hour)

	response := req.Response()

	val := s.queryLatestBetween(req.Metric, from, req.At)
	if val.Time.IsZero() {
		response.Status = http.StatusNotFound
	} else {
		response.Status = http.StatusOK
		response.Result = &val
	}

	return response.Submit(r)
}

// queryAllAt returns all metric values at a specific timestamp
func (s *Server) queryAllAt(r *rest.Rest) error {
	req := GetRequest(r)

	if req.At.IsZero() {
		r.Status(http.StatusBadRequest)
		return nil
	}

	from := req.At.Add(-time.Hour)

	response := req.Response()
	response.Status = http.StatusOK

	for _, metric := range s.Latest.Metrics() {
		if req.Match(metric) {
			val := s.queryLatestBetween(metric, from, req.At)
			if !val.Time.IsZero() {
				response.Metrics = append(response.Metrics, api2.Metric{
					Metric: metric,
					Time:   val.Time,
					Unit:   val.Unit,
					Value:  val.Value,
				})
			}
		}
	}

	return response.Submit(r)
}

// queryBetween returns all values for a single metric between two times
func (s *Server) queryBetween(r *rest.Rest) error {
	req := GetRequest(r)

	response := req.Response()

	switch {
	// Ensure from & to are set, from is before to and
	// a max duration of 24 hours
	case req.From.IsZero(),
		req.To.IsZero(),
		req.From.After(req.To),
		req.To.Sub(req.From) > 24*time.Hour:
		response.Status = http.StatusBadRequest

	default:
		q := s.Store.Query(req.Metric).
			Between(req.From, req.To).
			Build()

		for q.HasNext() {
			rec := q.Next()
			if rec.IsValid() {
				val := rec.Value
				response.Results = append(response.Results, api2.MetricValue{
					Time:  rec.Time,
					Unit:  val.Unit().ID(),
					Value: val.Float(),
				})
			}
		}

		if len(response.Results) > 0 {
			response.Status = http.StatusOK
		} else {
			response.Status = http.StatusNotFound
		}
	}

	return response.Submit(r)
}

// queryLatestBetween returns the most recent value of metric up to and including t.
// This will search up to the previous 6 hours for those metrics which are
// not regularly submitted.
func (s *Server) queryLatestBetween(metric string, from, to time.Time) api2.MetricValue {
	query := s.Store.Query(metric).
		Between(from, to).
		Build()

	var ret api2.MetricValue

	for query.HasNext() {
		rec := query.Next()
		if rec.IsValid() && (ret.Time.IsZero() || ret.Time.Before(rec.Time)) {
			val := rec.Value
			ret = api2.MetricValue{
				Time:  rec.Time,
				Unit:  val.Unit().ID(),
				Value: val.Float(),
			}
		}
	}

	return ret
}
