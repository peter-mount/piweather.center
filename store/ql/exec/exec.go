package exec

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-script/errors"
	lang2 "github.com/peter-mount/piweather.center/config/ql"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/ql"
	"github.com/peter-mount/piweather.center/store/ql/functions"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"net/http"
	"strings"
	"time"
)

type Executor struct {
	exState
	qp          *QueryPlan                        // QueryPlan to execute
	result      *api.Result                       // Query Results
	table       *api.Table                        // Current table
	row         *api.Row                          // Current row
	metrics     map[string][]ql.Value             // Collected data for each metric
	stack       []ql.Value                        // Stack for expressions
	using       map[string]*lang2.UsingDefinition // Using aliases
	colResolver *colResolver                      // Used when resolving columns
}

type exState struct {
	prevState   *exState  // link to previous station
	time        time.Time // Query time
	timeRange   api.Range // Query range
	selectLimit int       // Max number of rows to return in a query
}

func (ex *Executor) save() {
	old := ex.exState
	ex.exState.prevState = &old
}

func (ex *Executor) restore() {
	if ex.exState.prevState != nil {
		ex.exState = *ex.exState.prevState
	}
}

func (ex *Executor) Time() time.Time {
	return ex.time
}

func (qp *QueryPlan) Execute(result *api.Result) error {
	ex := &Executor{
		qp:          qp,
		result:      result,
		metrics:     make(map[string][]ql.Value),
		using:       make(map[string]*lang2.UsingDefinition),
		colResolver: newColResolver(),
		exState: exState{
			timeRange: qp.QueryRange,
		},
	}

	err := ex.run()

	if err == nil {
		ex.result.Finalise()
	} else {
		ex.result.Status = http.StatusInternalServerError
		ex.result.Message = err.Error()
		// Remove all results as we are failing
		ex.result.Table = nil
	}

	return err
}

func (ex *Executor) run() error {
	qp := ex.qp

	_ = qp.Metrics.ForEach(ex.getMetric)

	return lang2.NewBuilder().
		Query(ex.query).
		UsingDefinitions(ex.usingDefinitions).
		Histogram(ex.histogram).
		Select(ex.selectStatement).
		TableSelect(ex.tableSelect).
		WindRose(ex.windRose).
		SelectExpression(ex.selectExpression).
		AliasedExpression(ex.aliasedExpression).
		Expression(ex.expression).
		ExpressionModifier(ex.expressionModifier).
		Function(ex.function).
		Metric(ex.metric).
		Build().
		Query(qp.query)
}

func (ex *Executor) getMetric(m string) error {
	ex.metrics[m] = ex.qp.GetMetric(m)
	return nil
}

func (ex *Executor) query(_ lang2.Visitor, s *lang2.Query) error {
	ex.setSelectLimit(s.Limit)
	return nil
}

func (ex *Executor) setSelectLimit(l int) {
	ex.selectLimit = l
	if ex.selectLimit < 0 {
		ex.selectLimit = 0
	}
}

func (ex *Executor) usingDefinitions(v lang2.Visitor, s *lang2.UsingDefinitions) error {
	for _, e := range s.Defs {
		// Ensure the definition is valid
		if err := v.UsingDefinition(e); err != nil {
			return err
		}
		ex.using[e.Name] = e
	}
	return util.VisitorStop
}

func (ex *Executor) selectStatement(v lang2.Visitor, s *lang2.Select) error {
	ex.table = ex.result.NewTable()

	// Select has its own LIMIT defined
	if s.Limit > 0 {
		ex.save()
		defer ex.restore()
		ex.setSelectLimit(s.Limit)
	}

	if s.Expression != nil {
		// Create the required columns
		for _, ae := range s.Expression.Expressions {
			col := ex.colResolver.resolveColumn(ae)
			if ae.Unit != nil {
				col.SetUnit(ae.Unit.Unit())
			}
			ex.table.AddColumn(col)
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
	return util.VisitorStop
}

func (ex *Executor) selectExpression(_ lang2.Visitor, _ *lang2.SelectExpression) error {
	ex.table.PruneCurrentRow()

	// If we have exceeded the selectLimit then stop here
	if ex.selectLimit > 0 && ex.table.RowCount() >= ex.selectLimit {
		return util.VisitorExit
	}

	ex.row = ex.table.NewRow()
	return nil
}

func (ex *Executor) expression(v lang2.Visitor, s *lang2.Expression) error {
	var err error

	// If offset defined, temporarily adjust the current time by that offset
	if s.Using != "" || s.Modifier != nil {
		ex.save()
		defer ex.restore()

		// Resolve the modifier if we are declaring using
		mod := s.Modifier
		if s.Using != "" {
			uDef, exists := ex.using[s.Using]
			if !exists {
				// Should not happen as we checked before running
				return errors.Errorf(s.Pos, "panic: %q missing", s.Using)
			}
			mod = uDef.Modifier
		}

		for _, e := range mod {
			if err == nil {
				err = v.ExpressionModifier(e)
			}
		}
	}

	if err == nil {
		switch {
		case s.Metric != nil:
			err = v.Metric(s.Metric)
		case s.Function != nil:
			err = v.Function(s.Function)
		}
	}

	if err != nil {
		return err
	}

	return util.VisitorStop
}

func (ex *Executor) expressionModifier(v lang2.Visitor, s *lang2.ExpressionModifier) error {
	var err error

	if s.Offset != nil {
		ex.time = ex.time.Add(s.Offset.Duration(ex.timeRange.Every))
	}

	if s.Range != nil {
		if s.Range.IsRow() {
			err = s.Range.SetTime(ex.time, ex.timeRange.Every, v)
		}

		r := s.Range.Range()
		ex.time = r.From
		ex.timeRange.Every = r.Duration()
	}

	return err
}

func (ex *Executor) aliasedExpression(v lang2.Visitor, s *lang2.AliasedExpression) error {
	ex.resetStack()

	err := v.Expression(s.Expression)

	val, ok := ex.Pop()

	// If invalid but have values attached then get the last value in the set.
	// Required with metrics without an aggregation function around them
	if !val.IsTime && !val.Value.IsValid() && len(val.Values) > 0 {
		val = functions.InitialLast(val)
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
		col := ex.table.Columns[ex.row.Size()]
		val1, err := col.Transform(val.Value)
		if err != nil {
			return err
		}
		ex.row.AddValue(val.Time, val1)
	}

	return util.VisitorStop
}

type colResolver struct {
	visitor lang2.Visitor
	path    []string
}

func newColResolver() *colResolver {
	r := &colResolver{}
	r.visitor = lang2.NewBuilder().
		AliasedExpression(r.aliasedExpression).
		Function(r.function).
		Metric(r.metric).
		Build()

	return r
}

func (r *colResolver) append(s ...string) {
	r.path = append(r.path, s...)
}

func (r *colResolver) resolveColumn(v *lang2.AliasedExpression) *api.Column {
	return &api.Column{Name: r.resolveName(v)}
}

func (r *colResolver) resolveName(v *lang2.AliasedExpression) string {
	r.path = nil
	_ = r.visitor.AliasedExpression(v)
	return strings.Join(r.path, "")
}

func (r *colResolver) aliasedExpression(_ lang2.Visitor, f *lang2.AliasedExpression) error {
	if f.As != "" {
		r.append(f.As)
		return util.VisitorStop
	}
	return nil
}

func (r *colResolver) function(v lang2.Visitor, f *lang2.Function) error {
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
	return util.VisitorStop
}

func (r *colResolver) metric(_ lang2.Visitor, f *lang2.Metric) error {
	r.append(f.Name)
	return nil
}

func (ex *Executor) tableSelect(v lang2.Visitor, s *lang2.TableSelect) error {
	ex.table = ex.result.NewTable()

	// The unit required
	dataUnit := s.Unit.Unit()

	var t0 time.Time

	// Now the row data
	it := ex.timeRange.Iterator()
	for it.HasNext() {
		ex.time = it.Next()

		// Calculate the time
		t1 := ex.time
		if s.Time != nil {
			ex.resetStack()
			if err := v.Expression(s.Time); err != nil {
				return err
			}

			val, ok := ex.Pop()
			if ok && val.IsTime {
				t1 = ex.time
			}
		}

		// Now the value
		ex.resetStack()
		if err := v.Expression(s.Metric); err != nil {
			return err
		}

		val, ok := ex.Pop()
		if ok && !val.IsNull() && !val.IsTime {
			// If a set then reduce it to the last value
			if !val.Value.IsValid() && len(val.Values) > 0 {
				val = functions.InitialLast(val)
			}

			// Use first values Unit if not specified in the query
			if dataUnit == nil && val.Value.IsValid() {
				dataUnit = val.Value.Unit()
			}

			// Transform to the required unit and add to the current row
			val1, err := val.Value.As(dataUnit)
			if err != nil {
				return err
			}

			// Start a new row if first or starting a new day
			if t0.IsZero() || !time2.SameDay(t0, t1) {
				t0 = t1
				ex.table.PruneCurrentRow()
				ex.row = ex.table.NewRow()
				ex.row.AddDynamic(t0, t0.Format(time.RFC3339))
			}

			ex.row.AddValue(t1, val1)
		}
	}

	// Tell the visitor to stop processing this Select statement
	return util.VisitorStop
}
