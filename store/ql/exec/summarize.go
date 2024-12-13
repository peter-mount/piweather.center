package exec

import (
	lang2 "github.com/peter-mount/piweather.center/config/ql"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/ql"
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
	return &summary{}
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

func (s *summary) ensureCapacity(i int) {
	for i >= len(s.results) {
		s.results = append(s.results, ql.Value{})
		s.functions = append(s.functions, nil)
	}
}

func (s *summary) Set(i int, v value.Value) {
	if s != nil {
		s.ensureCapacity(i)
		e := s.results[i]
		e.Values = append(e.Values, ql.Value{Value: v})
		s.results[i] = e
	}
}

func (s *summary) With(i int, n string) {
	s.ensureCapacity(i)
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
	r.RowType = api.RowTypeSummary
	for _, v := range s.results {
		switch {
		case v.IsNull():
			r = r.AddNull()
		case v.IsTime:
			r = r.AddDynamic(v.Time, v.Time.Format(time.RFC3339))
		case v.Value.IsValid():
			r = r.AddValue(time.Time{}, v.Value)
		default:
			r = r.AddNull()
		}
	}
}
