package store

import (
	"github.com/peter-mount/piweather.center/graph/chart"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
	"time"
)

type dataSource []*Reading

func (s *Store) GetHistoryBetween(name string, start, end time.Time) chart.DataSource {
	start, end = util.NormalizeTime(start, end)

	var r dataSource
	for _, e := range s.GetHistory(name) {
		if util.TimeBetween(e.Time, start, end) {
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

func (s *dataSource) GetXRange() (time.Time, time.Time) {
	return (*s)[0].Time, (*s)[len(*s)-1].Time
}

func (s *dataSource) GetYRange() (value.Value, value.Value) {
	minVal, maxVal := math.MaxFloat64, -math.MaxFloat64
	for _, reading := range *s {
		v := reading.Value.Float()
		minVal, maxVal = math.Min(minVal, v), math.Max(maxVal, v)
	}

	unit := s.GetUnit()
	return unit.Value(minVal), unit.Value(maxVal)
}

func (s *dataSource) GetUnit() *value.Unit {
	return (*s)[0].Value.Unit()
}

func (s *dataSource) ForEach(f func(int, time.Time, value.Value)) {
	for i, e := range *s {
		f(i, e.Time, e.Value)
	}
}
