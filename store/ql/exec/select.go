package exec

import (
	lang2 "github.com/peter-mount/piweather.center/config/ql"
	"github.com/peter-mount/piweather.center/config/util"
)

func query(v lang2.Visitor[*Executor], s *lang2.Query) error {
	v.Get().setSelectLimit(s.Limit)
	return nil
}

func usingDefinitions(v lang2.Visitor[*Executor], s *lang2.UsingDefinitions) error {
	ex := v.Get()
	for _, e := range s.Defs {
		// Ensure the definition is valid
		if err := v.UsingDefinition(e); err != nil {
			return err
		}
		ex.using[e.Name] = e
	}
	return util.VisitorStop
}

func selectStatement(v lang2.Visitor[*Executor], s *lang2.Select) error {
	ex := v.Get()
	ex.table = ex.result.NewTable()

	ex.save()
	defer ex.restore()
	ex._select = s
	ex.summary = newSummary(s)

	// Select has its own LIMIT defined
	if s.Limit > 0 {
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

	// Summarize if we have data
	if ex.summary.IsValid() {
		ex.summary.summarize(ex.table)
	}

	// Tell the visitor to stop processing this Select statement
	return util.VisitorStop
}

func selectExpression(v lang2.Visitor[*Executor], d *lang2.SelectExpression) error {
	ex := v.Get()
	ex.table.PruneCurrentRow()

	// If we have exceeded the selectLimit then stop here
	if ex.selectLimit > 0 && ex.table.RowCount() >= ex.selectLimit {
		return util.VisitorExit
	}

	ex.row = ex.table.NewRow()

	for i, e := range d.Expressions {
		ex.selectColumn = i
		err := v.AliasedExpression(e)
		if err != nil {
			return err
		}
	}

	return util.VisitorStop
}
