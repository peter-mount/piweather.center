package server

import (
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/store/file"
	"github.com/peter-mount/piweather.center/store/file/api"
)

// record implements the /record api
func (s *Server) queryToday(r *rest.Rest) error {
	return s.query(r, r.Var(METRIC), func(b file.QueryBuilder) {
		b.Today()
	})
}

func (s *Server) query(r *rest.Rest, metric string, qbf func(file.QueryBuilder)) error {

	qb := s.Store.Query(metric)
	qbf(qb)
	query := qb.Build()

	response := api.Response{
		Status: 200,
		Metric: metric,
	}

	for query.HasNext() {
		rec := query.Next()
		val := rec.Value
		response.Results = append(response.Results, api.MetricValue{
			Time:  rec.Time,
			Unit:  val.Unit().ID(),
			Value: val.Float(),
		})
	}

	r.Status(response.Status).
		ContentType(r.GetHeader("Content-Type")).
		Value(response)

	return nil
}
