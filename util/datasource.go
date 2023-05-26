package util

import (
	"fmt"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

// DataSource represents a collection of values to be plotted
type DataSource interface {
	// Size of the DataSource
	Size() int
	// Get a specific entry in the DataSource
	Get(int) (time.Time, value.Value)
	// Period returns the Period of the entries within the DataSource
	Period() time2.Period
	// GetYRange returns the Range of values in the DataSource
	GetYRange() *value.Range
	// GetUnit returns the Unit of the values in the DataSource
	GetUnit() *value.Unit
	// ForEach calls a function for each entry in the DataSource
	ForEach(func(int, time.Time, value.Value))
}

// SliceDataSource returns a DataSource over a slice of Entry's
func SliceDataSource(s []Entry, period time2.Period, unit *value.Unit) (DataSource, error) {
	yRange := value.NewRange(unit)
	for _, e := range s {
		if err := yRange.Add(e.Value); err != nil {
			return nil, err
		}
	}

	return &sliceDataSource{
		s:      s,
		period: period,
		yRange: yRange,
		unit:   unit,
	}, nil
}

type sliceDataSource struct {
	s      []Entry
	period time2.Period
	yRange *value.Range
	unit   *value.Unit
}

type Entry struct {
	Time  time.Time
	Value value.Value
}

func (s *sliceDataSource) Size() int { return len(s.s) }

func (s *sliceDataSource) Get(i int) (time.Time, value.Value) {
	e := s.s[i]
	return e.Time, e.Value
}

func (s *sliceDataSource) Period() time2.Period { return s.period }

func (s *sliceDataSource) GetYRange() *value.Range { return s.yRange }

func (s *sliceDataSource) GetUnit() *value.Unit { return s.unit }

func (s *sliceDataSource) ForEach(f func(int, time.Time, value.Value)) {
	for i, e := range s.s {
		f(i, e.Time, e.Value)
	}
}

// PseudoDataSource returns a DataSource based on the output of a calculator
func PseudoDataSource(c value.Calculator, period time2.Period, unit *value.Unit, stepSize time.Duration, t value.Time) (DataSource, error) {
	if stepSize < time.Minute {
		return nil, fmt.Errorf("invalid stepSize %v", stepSize)
	}

	var s []Entry

	tm := period.Start()
	for tm.Before(period.End()) {
		t.SetTime(tm)
		v, err := c(t)
		if err != nil {
			return nil, err
		}
		s = append(s, Entry{
			Time:  tm,
			Value: v,
		})

		tm = tm.Add(stepSize)
	}

	return SliceDataSource(s, period, unit)
}

// LimitedPseudoDataSource is more optimal than PseudoDataSource as it doesn't have to precalculate the results as the min/max Value's
// are provided, so calculations are made on the fly.
func LimitedPseudoDataSource(c value.Calculator, period time2.Period, unit *value.Unit, stepSize time.Duration, min, max value.Value, t value.Time) (DataSource, error) {
	if stepSize < time.Minute {
		return nil, fmt.Errorf("invalid stepSize %v", stepSize)
	}

	yRange := value.NewRange(unit)
	_ = yRange.Add(min)
	_ = yRange.Add(max)

	return &limitedDataSource{
		c:        c,
		t:        t,
		size:     int(period.DurationMinutes() / stepSize.Minutes()),
		stepSize: stepSize,
		period:   period,
		unit:     unit,
		yRange:   yRange,
	}, nil
}

type limitedDataSource struct {
	c        value.Calculator
	t        value.Time
	size     int
	stepSize time.Duration
	period   time2.Period
	unit     *value.Unit
	yRange   *value.Range
}

func (s *limitedDataSource) Size() int { return s.size }

func (s *limitedDataSource) Get(i int) (time.Time, value.Value) {
	t := s.t.Clone()
	t.SetTime(s.period.Start().Add(s.stepSize * time.Duration(i)))
	v, _ := s.c(t)
	return t.Time(), v
}

func (s *limitedDataSource) Period() time2.Period { return s.period }

func (s *limitedDataSource) GetYRange() *value.Range { return s.yRange }

func (s *limitedDataSource) GetUnit() *value.Unit { return s.unit }

func (s *limitedDataSource) ForEach(f func(int, time.Time, value.Value)) {
	t := s.t.Clone()
	t.SetTime(s.period.Start())

	i := 0
	for s.period.Contains(t.Time()) {
		v, _ := s.c(t)
		f(i, t.Time(), v)
		i++
		t.SetTime(t.Time().Add(s.stepSize))
	}
}
