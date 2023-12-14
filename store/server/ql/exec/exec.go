package exec

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/server/ql/lang"
	"time"
)

type executor struct {
	qp      *QueryPlan         // QueryPlan to execute
	result  *api.Result        // Query Results
	table   *api.Table         // Current table
	row     *api.Row           // Current row
	metrics map[string][]Value // Collected data for each metric
	stack   []Value            // Stack for expressions
}

func (qp *QueryPlan) Execute() (*api.Result, error) {
	ex := &executor{
		qp:      qp,
		result:  &api.Result{},
		metrics: make(map[string][]Value),
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
		Expression(ex.expression).
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
		} else {
			// Create the required columns
			for i, ae := range s.Expression.Expressions {
				c := api.Column{}
				if ae.As != "" {
					c.Name = ae.As
				} else {
					// FIXME resolve the src metric name or based on function like postgresql does
					c.Name = fmt.Sprintf("col%d", i)
				}
				ex.table.AddColumn(c)
			}

			// Now the row data
			// TODO scan the time range here
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
	log.Printf("ae %v", s)

	ex.resetStack()

	err := v.Expression(s.Expression)
	if err != nil {
		log.Printf("ae err %v", err)
	}

	val, ok := ex.pop()
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

func (ex *executor) expression(v lang.Visitor, s *lang.Expression) error {
	log.Printf("ex %v", s)
	return nil
}
