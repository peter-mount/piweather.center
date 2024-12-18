package exec

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/go-script/errors"
	lang2 "github.com/peter-mount/piweather.center/config/ql"
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

var (
	queryPlanVisitor = lang2.NewBuilder[*QueryPlan]().
		AliasedExpression(qpAliasedExpression).
		Metric(qpMetric).
		QueryRange(qpQueryRange).
		Expression(qpExpression).
		Build()
)

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

	if err := queryPlanVisitor.Clone().
		Set(qp).
		Query(q); err != nil {
		return nil, err
	}

	qp.ScanRange = qp.ScanRange.Add(qp.QueryRange.Expand(qp.minOffset, qp.maxOffset))

	return qp, nil
}

func qpAliasedExpression(v lang2.Visitor[*QueryPlan], m *lang2.AliasedExpression) error {
	var err error
	switch {
	case m.Group != nil:
		err = v.AliasedGroup(m.Group)
	default:
		err = v.Expression(m.Expression)
	}
	if err != nil {
		return err
	}
	return errors.VisitorStop
}

func qpMetric(v lang2.Visitor[*QueryPlan], m *lang2.Metric) error {
	v.Get().Metrics.Add(m.Name)
	return nil
}

func qpQueryRange(v lang2.Visitor[*QueryPlan], m *lang2.QueryRange) error {
	v.Get().QueryRange = m.Range()
	return nil
}

func qpExpression(v lang2.Visitor[*QueryPlan], m *lang2.Expression) error {
	qp := v.Get()

	// Check for modifiers, looking them up if m.Using is Set
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

	return errors.VisitorStop
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
