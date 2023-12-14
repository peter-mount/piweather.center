package exec

import (
	"github.com/peter-mount/piweather.center/store/file"
	"github.com/peter-mount/piweather.center/store/memory"
	"github.com/peter-mount/piweather.center/store/server/ql/lang"
	"strings"
)

type QueryPlan struct {
	metrics map[string]interface{} // Set of metrics we require
	_range  lang.Range             // Time range to scan
	query   *lang.Query            // Queries for this plan that share the Range
	store   file.Store             // The actual file Store
	latest  memory.Latest          // Needed when selecting all (*)
}

func (qp *QueryPlan) String() string {
	var s []string
	s = append(s,
		"QueryPlan:",
		"    Range:"+qp._range.String())

	m := "  Metrics: "
	ml := len(m)
	for k, _ := range qp.metrics {
		m = m + " " + k
		if len(k) > 80 {
			s = append(s, m)
			m = strings.Repeat(" ", ml)
		}
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
		SelectExpression(qp.addSelectExpression).
		Metric(qp.addMetric).
		QueryRange(qp.setQueryRange).
		Build()); err != nil {
		return nil, err
	}

	return qp, nil
}

func (qp *QueryPlan) addSelectExpression(_ lang.Visitor, m *lang.SelectExpression) error {
	// If we have all set then include all metrics
	if m.All {
		for _, n := range qp.latest.Metrics() {
			qp.metrics[n] = nil
		}
	}
	return nil
}

func (qp *QueryPlan) addMetric(_ lang.Visitor, m *lang.Metric) error {
	qp.metrics[m.Name] = nil
	return nil
}

func (qp *QueryPlan) setQueryRange(_ lang.Visitor, m *lang.QueryRange) error {
	qp._range = m.Range()
	return nil
}
