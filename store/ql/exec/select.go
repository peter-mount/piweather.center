package exec

import (
	lang2 "github.com/peter-mount/piweather.center/config/ql"
	"github.com/peter-mount/piweather.center/config/util"
)

func (ex *Executor) selectStatement(v lang2.Visitor[*Executor], s *lang2.Select) error {
	ex.table = ex.result.NewTable()

	ex.save()
	defer ex.restore()
	ex._select = s

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

	//fmt.Println("Summarize")
	//if err := v.Summarize(s.Summarize); err != nil {
	//	fmt.Println("Summarize", err)
	//	return err
	//}

	// Tell the visitor to stop processing this Select statement
	return util.VisitorStop
}
