package exec

import (
	ql2 "github.com/peter-mount/piweather.center/config/ql"
	"github.com/peter-mount/piweather.center/config/util"
)

//type histogram struct {
//	min value.Value
//	max value.Value
//}

func histogram(v ql2.Visitor[*Executor], s *ql2.Histogram) error {
	ex := v.Get()

	ex.table = ex.result.NewTable()

	if s.Expression != nil {
		//// Create the required columns
		//for _, ae := range s.Expression.Expressions {
		//	col := ex.colResolver.resolveColumn(ae)
		//	col.SetUnit(ae.GetUnit())
		//	ex.table.AddColumn(col)
		//}
		//
		//// Now the row data
		//it := ex.timeRange.Iterator()
		//for it.HasNext() {
		//	ex.time = it.Next()
		//
		//	if err := v.SelectExpression(s.Expression); err != nil {
		//		return err
		//	}
		//}
	}

	// Tell the visitor to stop processing this Histogram statement
	return util.VisitorStop
}
