package exec

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/server/ql/lang"
	"strings"
	"time"
)

type Executor struct {
	qp          *QueryPlan         // QueryPlan to execute
	result      *api.Result        // Query Results
	table       *api.Table         // Current table
	row         *api.Row           // Current row
	metrics     map[string][]Value // Collected data for each metric
	time        time.Time          // Query time
	timeRange   lang.Range         // Query range
	stack       []Value            // Stack for expressions
	colResolver *colResolver       // Used when resolving columns
}

func (qp *QueryPlan) Execute() (*api.Result, error) {
	ex := &Executor{
		qp:          qp,
		result:      &api.Result{},
		metrics:     make(map[string][]Value),
		colResolver: newColResolver(),
		timeRange:   qp._range,
	}

	if err := ex.run(); err != nil {
		return nil, err
	}

	ex.result.Finalise()

	return ex.result, nil
}

func (ex *Executor) run() error {
	qp := ex.qp

	for m, _ := range qp.metrics {
		var e []Value
		q := qp.store.Query(m).
			Between(qp._range.From, qp._range.To).
			Build()
		for q.HasNext() {
			e = append(e, FromRecord(q.Next()))
		}
		ex.metrics[m] = e
	}

	if err := qp.query.Accept(lang.NewBuilder().
		Select(ex.selectStatement).
		SelectExpression(ex.selectExpression).
		AliasedExpression(ex.aliasedExpression).
		Function(ex.function).
		Metric(ex.metric).
		Build()); err != nil {
		return err
	}

	return nil
}

func (ex *Executor) selectStatement(v lang.Visitor, s *lang.Select) error {
	ex.table = ex.result.NewTable()

	if s.Expression != nil {
		if s.Expression.All {
			// TODO handle this if we keep it
			return participle.Errorf(s.Pos, "Select * unsupported")
		} else {
			// Create the required columns
			for _, ae := range s.Expression.Expressions {
				ex.table.AddColumn(ex.colResolver.resolveColumn(ae))
			}
		}

		// Now the row data
		it := ex.timeRange.Iterator()
		for it.HasNext() {
			ex.time = it.Next()

			if err := v.SelectExpression(s.Expression); err != nil {
				return err
			}
		}
	}

	// Tell the visitor to stop processing this Select statement
	return lang.VisitorStop
}

func (ex *Executor) selectExpression(v lang.Visitor, s *lang.SelectExpression) error {
	ex.row = ex.table.NewRow()
	return nil
}

func (ex *Executor) aliasedExpression(v lang.Visitor, s *lang.AliasedExpression) error {
	ex.resetStack()

	err := v.Expression(s.Expression)

	val, ok := ex.pop()

	// If invalid but have values attached then get the last value in the set.
	// Required with metrics without an aggregation function around them
	if !val.IsTime && !val.Value.IsValid() && len(val.Values) > 0 {
		val = InitialLast(val)
	}

	switch {
	case err != nil:
		log.Println(err)
		ex.row.AddNull()

	case !ok,
		val.IsNull():
		ex.row.AddNull()

	case val.IsTime:
		ex.row.AddDynamic(val.Time, val.Time.Format(time.RFC3339))

	default:
		col := ex.table.Columns[len(ex.row.Columns)]
		val1, err := col.Transform(val.Value)
		if err != nil {
			return err
		}
		ex.row.AddValue(val.Time, val1)
	}

	return lang.VisitorStop
}

type colResolver struct {
	visitor lang.Visitor
	path    []string
}

func newColResolver() *colResolver {
	r := &colResolver{}
	r.visitor = lang.NewBuilder().
		AliasedExpression(r.aliasedExpression).
		Function(r.function).
		Metric(r.metric).
		Build()

	return r
}

func (r *colResolver) append(s ...string) {
	r.path = append(r.path, s...)
}

func (r *colResolver) resolveColumn(v *lang.AliasedExpression) api.Column {
	return api.Column{Name: r.resolveName(v)}
}

func (r *colResolver) resolveName(v *lang.AliasedExpression) string {
	r.path = nil
	_ = r.visitor.AliasedExpression(v)
	return strings.Join(r.path, "")
}

func (r *colResolver) aliasedExpression(v lang.Visitor, f *lang.AliasedExpression) error {
	if f.As != "" {
		r.append(f.As)
		return lang.VisitorStop
	}
	return nil
}

func (r *colResolver) function(v lang.Visitor, f *lang.Function) error {
	r.append(f.Name, "(")
	for i, e := range f.Expressions {
		if i > 0 {
			r.append(",")
		}
		if err := v.Expression(e); err != nil {
			return err
		}
	}
	r.append(")")
	return lang.VisitorStop
}

func (r *colResolver) metric(_ lang.Visitor, f *lang.Metric) error {
	r.append(f.Name)
	return nil
}
