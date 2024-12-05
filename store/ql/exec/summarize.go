package exec

import (
	lang2 "github.com/peter-mount/piweather.center/config/ql"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/ql"
	"github.com/peter-mount/piweather.center/store/ql/exec/utils"
	"github.com/peter-mount/piweather.center/store/ql/functions"
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

func summarize(v lang2.Visitor[*Executor], d *lang2.Summarize) error {
	ex := v.Get()

	// We have a summarize clause so mark the summary as valid
	ex.summary.SetValid()

	if d.With != "" {
		ex.summary.With(ex.selectColumn, d.With)
	}

	return nil
}

type summary struct {
	functions []*functions.Function // Aggregators to use
	results   []ql.Value            // Summary results
	valid     bool
}

func newSummary(d *lang2.Select) *summary {
	l := len(d.Expression.Expressions)
	return &summary{
		functions: make([]*functions.Function, l),
		results:   make([]ql.Value, l),
	}
}

func (s *summary) IsValid() bool {
	return s.valid
}

func (s *summary) SetValid() {
	s.valid = true
}

func (s *summary) IsAggregated(i int) bool {
	if i < len(s.functions) {
		return s.functions[i] != nil
	}
	return false
}

func (s *summary) Set(i int, v value.Value) {
	if s != nil {
		if i < len(s.results) {
			e := s.results[i]
			e.Values = append(e.Values, ql.Value{Value: v})
			s.results[i] = e
		}
	}
}

func (s *summary) With(i int, n string) {
	if s.functions[i] == nil {
		if f, exists := functions.GetFunction(n); exists && f.IsAggregator() {
			s.functions[i] = &f
		}
	}
}

func (s *summary) summarize(table *api.Table) {
	// Aggregate the results
	for cn, v := range s.results {
		if f := s.functions[cn]; f != nil {
			if newS, err := f.RunAggregator(v); err == nil {
				s.results[cn] = newS
			}
		}
	}

	// Finally add a new row with the summary
	r := table.NewRow()
	for _, s := range s.results {
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
}

func summarizex(v lang2.Visitor[*Executor], d *lang2.Summarize) error {
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
