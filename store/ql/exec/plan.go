package exec

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/go-script/errors"
	lang2 "github.com/peter-mount/piweather.center/config/ql"
	util2 "github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/file"
	"github.com/peter-mount/piweather.center/store/ql"
	"github.com/peter-mount/piweather.center/util"
	"time"
)

type QueryPlan struct {
	qpState
	QueryRange api.Range      `json:"queryRange"` // Time range for results
	ScanRange  api.Range      `json:"scanRange"`  // Time range to scan for metrics
	Metrics    util.StringSet `json:"metrics"`    // Set of Metrics we require
	query      *lang2.Query   // Queries for this plan that share the Range
	store      file.Store     // The actual file Store
	// Used to handle expression offsets, so we can expand the QueryRange to get aggregated Metrics
	minOffset time.Duration
	maxOffset time.Duration
}

type qpState struct {
	prevState *qpState
	offset    time.Duration
}

func (qp *QueryPlan) save() {
	old := qp.qpState
	qp.qpState.prevState = &old
}

func (qp *QueryPlan) restore() {
	if qp.qpState.prevState != nil {
		qp.qpState = *qp.prevState
	}
}

func NewQueryPlan(s file.Store, q *lang2.Query) (*QueryPlan, error) {

	// We must have a QueryRange, and it cannot reference "row"
	if q.QueryRange == nil || q.QueryRange.IsRow() {
		return nil, participle.Errorf(q.Pos, "invalid QueryRange")
	}

	qp := &QueryPlan{
		query:   q,
		Metrics: util.NewStringSet(),
		store:   s,
	}

	if err := lang2.New().
		AliasedExpression(qp.aliasedExpression).
		Metric(qp.addMetric).
		QueryRange(qp.setQueryRange).
		Expression(qp.expression).
		Build().
		Query(q); err != nil {
		return nil, err
	}

	qp.ScanRange = qp.ScanRange.Add(qp.QueryRange.Expand(qp.minOffset, qp.maxOffset))

	return qp, nil
}

func (qp *QueryPlan) aliasedExpression(v lang2.Visitor, m *lang2.AliasedExpression) error {
	if err := v.Expression(m.Expression); err != nil {
		return err
	}
	return util2.VisitorStop
}

func (qp *QueryPlan) addMetric(_ lang2.Visitor, m *lang2.Metric) error {
	qp.Metrics.Add(m.Name)
	return nil
}

func (qp *QueryPlan) setQueryRange(_ lang2.Visitor, m *lang2.QueryRange) error {
	qp.QueryRange = m.Range()
	return nil
}

func (qp *QueryPlan) expression(v lang2.Visitor, m *lang2.Expression) error {

	// Check for modifiers, looking them up if m.Using is set
	mods := m.Modifier
	if m.Using != "" && qp.query.Using != nil {
		for _, def := range qp.query.Using.Defs {
			if def.Name == m.Using {
				mods = def.Modifier
			}
		}

		// Should not happen
		if mods == nil {
			return errors.Errorf(m.Pos, "%q is undefined", m.Using)
		}
	}

	if mods != nil {
		qp.save()
		defer qp.restore()

		for _, e := range mods {
			if e.Offset != nil {
				qp.offset = qp.offset + e.Offset.Duration(0)

				if qp.offset < qp.minOffset {
					qp.minOffset = qp.offset
				}

				if qp.offset > qp.maxOffset {
					qp.maxOffset = qp.offset
				}
			}

			if e.Range != nil {
				qp.ScanRange = qp.ScanRange.Add(e.Range.Range())
			}
		}
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

	return util2.VisitorStop
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
