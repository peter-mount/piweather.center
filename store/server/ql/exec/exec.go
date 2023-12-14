package exec

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/server/ql/lang"
	"strings"
	"time"
)

type executor struct {
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
	ex := &executor{
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

func (ex *executor) run() error {
	qp := ex.qp

	log.Printf("Retrieving data")
	for m, _ := range qp.metrics {
		var e []Value
		q := qp.store.Query(m).
			Between(qp._range.From, qp._range.To).
			Build()
		for q.HasNext() {
			e = append(e, FromRecord(q.Next()))
		}
		ex.metrics[m] = e
		log.Printf("%q got %d", m, len(e))
	}

	log.Printf("Querying data")
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

func (ex *executor) selectStatement(v lang.Visitor, s *lang.Select) error {
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

func (ex *executor) selectExpression(v lang.Visitor, s *lang.SelectExpression) error {
	ex.row = ex.table.NewRow()
	return nil
}

func (ex *executor) aliasedExpression(v lang.Visitor, s *lang.AliasedExpression) error {
	ex.resetStack()

	err := v.Expression(s.Expression)

	val, ok := ex.pop()

	// If invalid but have values attached then get the last value in the set.
	// Required with metrics without an aggregation function around them
	if !val.Value.IsValid() && len(val.Values) > 0 {
		val = InitialLast(val)
	}

	switch {
	case err == nil && !ok,
		val.IsNull:
		ex.row.Add(api.Cell{Type: api.CellNull})

	case err == nil && val.IsTime:
		ex.row.Add(api.Cell{
			Type:   api.CellString,
			Time:   val.Time,
			String: val.Time.Format(time.RFC3339),
		})

	case err == nil:
		col := ex.table.Columns[len(ex.row.Columns)]
		val1, err := col.Transform(val.Value)
		if err != nil {
			return err
		} else {
			ex.row.Add(api.Cell{
				Type:   api.CellString,
				Time:   val.Time,
				String: val1.PlainString(),
			})
		}

	case err != nil:
		log.Println(err)
		ex.row.Add(api.Cell{
			Type:   api.CellNull,
			String: "???",
		})
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
