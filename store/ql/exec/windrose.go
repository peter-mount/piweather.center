package exec

import (
	"github.com/peter-mount/piweather.center/store/ql/lang"
	"github.com/peter-mount/piweather.center/weather/value"
)

type windrose struct {
	min value.Value
	max value.Value
}

func (ex *Executor) windrose(v lang.Visitor, s *lang.WindRose) error {
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
	return lang.VisitorStop
}
