package server

import (
	"github.com/peter-mount/go-kernel/v2/rest"
	api2 "github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/file"
)

func (s *Server) queryToday(r *rest.Rest) error {
	return s.query(r, r.Var(METRIC), func(b file.QueryBuilder) {
		b.Today()
	})
}

func (s *Server) queryTodayUTC(r *rest.Rest) error {
	return s.query(r, r.Var(METRIC), func(b file.QueryBuilder) {
		b.TodayUTC()
	})
}

func (s *Server) query(r *rest.Rest, metric string, qbf func(file.QueryBuilder)) error {

	qb := s.Store.Query(metric)
	qbf(qb)
	query := qb.Build()

	response := api2.Response{
		Status: 200,
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

	r.Status(response.Status).
		ContentType(r.GetHeader("Content-Type")).
		Value(response.Sort())

	return nil
}
