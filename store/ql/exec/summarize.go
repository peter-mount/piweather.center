package exec

import (
	lang2 "github.com/peter-mount/piweather.center/config/ql"
	"github.com/peter-mount/piweather.center/store/ql"
	"github.com/peter-mount/piweather.center/store/ql/exec/utils"
	"time"
)

func summarize(v lang2.Visitor[*Executor], d *lang2.Summarize) error {
	ex := v.Get()

	// Lookup the aggregators based on the original query
	funcs, err := utils.GetAggregators(ex._select)
	if err != nil {
		return err
	}

	// Now aggregate the table's results
	table := ex.table
	rc := table.RowCount()
	cc := table.ColumnCount()
	summary := make([]ql.Value, cc)
	for cn, s := range summary {
		if f := funcs[cn]; f != nil {
			// Add the column's values into the summary column
			for rn := 0; rn < rc; rn++ {
				cell := table.GetRow(rn).Cell(cn)
				if cell.Value.IsValid() {
					s.Values = append(s.Values, ql.Value{Value: cell.Value})
				}
			}

			// Now run the aggregator to reduce the summary column
			if newS, err := f.RunAggregator(s); err == nil {
				summary[cn] = newS
			}
		}
	}

	// Finally add a new row with the summary
	r := table.NewRow()
	for _, s := range summary {
		switch {
		case s.IsNull():
			r = r.AddNull()
		case s.IsTime:
			r = r.AddDynamic(s.Time, s.Time.Format(time.RFC3339))
		case s.Value.IsValid():
			r = r.AddValue(time.Time{}, s.Value)
		default:
			r = r.AddNull()
		}
	}

	return nil
}
