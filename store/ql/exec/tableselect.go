package exec

import (
	lang2 "github.com/peter-mount/piweather.center/config/ql"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/store/ql/functions"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"time"
)

func tableSelect(v lang2.Visitor[*Executor], s *lang2.TableSelect) error {
	ex := v.Get()

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
