package exec

import (
	lang2 "github.com/peter-mount/piweather.center/config/ql"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/ql"
	"net/http"
	"time"
)

type Executor struct {
	exState
	qp           *QueryPlan                        // QueryPlan to execute
	result       *api.Result                       // Query Results
	table        *api.Table                        // Current table
	row          *api.Row                          // Current row
	metrics      map[string][]ql.Value             // Collected data for each metric
	stack        []ql.Value                        // Stack for expressions
	using        map[string]*lang2.UsingDefinition // Using aliases
	selectColumn int                               // The select expression index being evaluated
}

type exState struct {
	prevState   *exState      // link to previous station
	time        time.Time     // Query time
	timeRange   api.Range     // Query range
	_select     *lang2.Select // Select being processed
	selectLimit int           // Max number of rows to return in a query
	inGroup     bool          // True if we are inside a group
	summary     *summary      // Summary
}

var (
	execVisitor = lang2.NewBuilder[*Executor]().
		AliasedExpression(aliasedExpression).
		Expression(expression).
		ExpressionModifier(expressionModifier).
		Function(function).
		Histogram(histogram).
		Metric(metric).
		Query(query).
		Select(selectStatement).
		Summarize(summarize).
		TableSelect(tableSelect).
		UsingDefinitions(usingDefinitions).
		WindRose(windRose).
		Build()
)

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
		qp:      qp,
		result:  result,
		metrics: make(map[string][]ql.Value),
		using:   make(map[string]*lang2.UsingDefinition),
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

	return execVisitor.Clone().Set(ex).Query(qp.query)
}

func (ex *Executor) getMetric(m string) error {
	ex.metrics[m] = ex.qp.GetMetric(m)
	return nil
}

func (ex *Executor) setSelectLimit(l int) {
	ex.selectLimit = l
	if ex.selectLimit < 0 {
		ex.selectLimit = 0
	}
}
