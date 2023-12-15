package exec

import (
	"github.com/peter-mount/piweather.center/store/file"
	"github.com/peter-mount/piweather.center/store/memory"
	"github.com/peter-mount/piweather.center/store/server/ql/lang"
	"strings"
	"time"
)

type QueryPlan struct {
	metrics    map[string]interface{} // Set of metrics we require
	queryRange lang.Range             // Time range for results
	scanRange  lang.Range             // Time range to scan for metrics
	query      *lang.Query            // Queries for this plan that share the Range
	store      file.Store             // The actual file Store
	latest     memory.Latest          // Needed when selecting all (*)
	// Used to handle expression offsets, so we can expand the queryRange to get aggregated metrics
	offset    time.Duration
	minOffset time.Duration
	maxOffset time.Duration
}

func (qp *QueryPlan) String() string {
	var s []string
	s = append(s,
		"QueryPlan:",
		"Data Range: "+qp.queryRange.String(),
		"Scan Range: "+qp.scanRange.String())

	m := "   Metrics:"
	ml := len(m)
	for k, _ := range qp.metrics {
		if len(m)+len(k) > 80 {
			s = append(s, m)
			m = strings.Repeat(" ", ml)
		}
		m = m + " " + k
	}
	if len(m) > ml {
		s = append(s, m)
	}

	return strings.Join(s, "\n")
}

func NewQueryPlan(s file.Store, l memory.Latest, q *lang.Query) (*QueryPlan, error) {
	qp := &QueryPlan{
		query:   q,
		metrics: make(map[string]interface{}),
		store:   s,
		latest:  l,
	}

	if err := q.Accept(lang.NewBuilder().
		AliasedExpression(qp.aliasedExpression).
		Metric(qp.addMetric).
		QueryRange(qp.setQueryRange).
		Expression(qp.expression).
		Build()); err != nil {
		return nil, err
	}

	qp.scanRange = qp.queryRange.Expand(qp.minOffset, qp.maxOffset)

	return qp, nil
}

func (qp *QueryPlan) aliasedExpression(v lang.Visitor, m *lang.AliasedExpression) error {
	if err := v.Expression(m.Expression); err != nil {
		return err
	}
	return lang.VisitorStop
}

func (qp *QueryPlan) addMetric(_ lang.Visitor, m *lang.Metric) error {
	qp.metrics[m.Name] = nil
	return nil
}

func (qp *QueryPlan) setQueryRange(_ lang.Visitor, m *lang.QueryRange) error {
	qp.queryRange = m.Range()
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

func (qp *QueryPlan) GetMetric(m string) []Value {
	var e []Value
	q := qp.store.Query(m).
		Between(qp.scanRange.From, qp.scanRange.To).
		Build()
	for q.HasNext() {
		e = append(e, FromRecord(q.Next()))
	}
	return e
}
