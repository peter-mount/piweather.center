package store

import (
	"github.com/peter-mount/piweather.center/graph/chart"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

type dataSource []*Reading

func (s *Store) GetHistoryBetween(name string, start, end time.Time) chart.DataSource {
	start, end = time2.NormalizeTime(start, end)

	var r dataSource
	for _, e := range s.GetHistory(name) {
		if time2.Between(e.Time, start, end) {
			r = append(r, e)
		}
	}
	return &r
}

func (s *dataSource) Size() int {
	return len(*s)
}

func (s *dataSource) Get(i int) (time.Time, value.Value) {
	e := (*s)[i]
	return e.Time, e.Value
}

func (s *dataSource) Period() time2.Period {
	return time2.PeriodOf((*s)[0].Time, (*s)[len(*s)-1].Time)
}

func (s *dataSource) GetYRange() *value.Range {
	r := value.NewRange(s.GetUnit())
	for _, reading := range *s {
		_ = r.Add(reading.Value)
	}
	return r
}

func (s *dataSource) GetUnit() *value.Unit {
	return (*s)[0].Value.Unit()
}

func (s *dataSource) ForEach(f func(int, time.Time, value.Value)) {
	for i, e := range *s {
		f(i, e.Time, e.Value)
	}
}
