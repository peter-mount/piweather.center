package exec

import (
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/file"
	"github.com/peter-mount/piweather.center/store/server/ql"
	"github.com/peter-mount/piweather.center/store/server/ql/lang"
	"github.com/peter-mount/piweather.center/util"
	"time"
)

type QueryPlan struct {
	QueryRange api.Range      `json:"queryRange"` // Time range for results
	ScanRange  api.Range      `json:"scanRange"`  // Time range to scan for metrics
	Metrics    util.StringSet `json:"metrics"`    // Set of Metrics we require
	query      *lang.Query    // Queries for this plan that share the Range
	store      file.Store     // The actual file Store
	// Used to handle expression offsets, so we can expand the QueryRange to get aggregated Metrics
	offset    time.Duration
	minOffset time.Duration
	maxOffset time.Duration
}

func NewQueryPlan(s file.Store, q *lang.Query) (*QueryPlan, error) {
	qp := &QueryPlan{
		query:   q,
		Metrics: util.NewStringSet(),
		store:   s,
	}

	if err := q.Accept(lang.NewBuilder().
		AliasedExpression(qp.aliasedExpression).
		Metric(qp.addMetric).
		QueryRange(qp.setQueryRange).
		Expression(qp.expression).
		Build()); err != nil {
		return nil, err
	}

	qp.ScanRange = qp.QueryRange.Expand(qp.minOffset, qp.maxOffset)

	return qp, nil
}

func (qp *QueryPlan) aliasedExpression(v lang.Visitor, m *lang.AliasedExpression) error {
	if err := v.Expression(m.Expression); err != nil {
		return err
	}
	return lang.VisitorStop
}

func (qp *QueryPlan) addMetric(_ lang.Visitor, m *lang.Metric) error {
	qp.Metrics.Add(m.Name)
	return nil
}

func (qp *QueryPlan) setQueryRange(_ lang.Visitor, m *lang.QueryRange) error {
	qp.QueryRange = m.Range()
	return nil
}

func (qp *QueryPlan) expression(v lang.Visitor, m *lang.Expression) error {
	if m.Offset != nil {
		old := qp.offset
		qp.offset = qp.offset + m.Offset.Duration

		if qp.offset < qp.minOffset {
			qp.minOffset = qp.offset
		}

		if qp.offset > qp.maxOffset {
			qp.maxOffset = qp.offset
		}

		defer func() {
			qp.offset = old
		}()
	}

	var err error

	switch {
	case m.Metric != nil:
		err = v.Metric(m.Metric)
	case m.Function != nil:
		err = v.Function(m.Function)
	}

	if err != nil {
		return err
	}

	return lang.VisitorStop
}

func (qp *QueryPlan) GetMetric(m string) []ql.Value {
	var e []ql.Value
	q := qp.store.Query(m).
		Between(qp.ScanRange.From, qp.ScanRange.To).
		Build()
	for q.HasNext() {
		e = append(e, ql.FromRecord(q.Next()))
	}
	return e
}
