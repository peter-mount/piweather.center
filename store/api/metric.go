package api

import (
	"fmt"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

type Metric struct {
	Metric    string    `json:"metric" xml:"metric,attr,omitempty" yaml:"metric"`
	Time      time.Time `json:"time" xml:"time,attr,omitempty" yaml:"time"`
	Unit      string    `json:"unit" xml:"unit,attr,omitempty" yaml:"unit"`
	Value     float64   `json:"value" xml:",chardata" yaml:"value"`
	Formatted string    `json:"formatted,omitempty" xml:"formatted,attr,omitempty" yaml:"formatted,omitempty"`
	Unix      int64     `json:"unix,omitempty" xml:"unix,attr,omitempty" yaml:"unix,omitempty"`
}

// IsValid returns true if the Metric has a value
func (m Metric) IsValid() bool {
	return !m.Time.IsZero()
}

func (m Metric) String() string {
	return fmt.Sprintf("Metric[%q,%q,%q,%f]", m.Metric, m.Time.Format(time.RFC3339), m.Unit, m.Value)
}

type MetricValue struct {
	Time  time.Time `json:"time" xml:"time,attr"`
	Unit  string    `json:"unit" xml:"unit,attr"`
	Value float64   `json:"value" xml:",chardata"`
}

type MetricGroup struct {
	Metric string        `json:"metric" xml:"metric,attr"`
	Values []MetricValue `json:"values" xml:"values"`
}

func (s MetricGroup) Size() int {
	return len(s.Values)
}

func (s MetricGroup) Get(i int) MetricValue {
	return s.Values[i]
}

func (s MetricGroup) Period() time2.Period {
	return time2.PeriodOf(s.Values[0].Time, s.Values[len(s.Values)-1].Time)
}

func (s MetricGroup) GetYRange() *value.Range {
	r := value.NewRange(s.GetUnit())
	for _, reading := range s.Values {
		u, _ := value.GetUnit(reading.Unit)
		_ = r.Add(u.Value(reading.Value))
	}
	return r
}

func (s MetricGroup) GetUnit() *value.Unit {
	u, _ := value.GetUnit(s.Values[0].Unit)
	return u
}

func (s MetricGroup) ForEach(f func(int, MetricValue)) {
	for i, e := range s.Values {
		f(i, e)
	}
}
