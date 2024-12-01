package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util/units"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
)

type MetricAccept interface {
	AcceptMetric(api.Metric) bool
}

type MetricList struct {
	Pos     lexer.Position
	Metrics []*Metric `parser:"@@*"`
}

func (m *MetricList) AcceptMetric(v api.Metric) bool {
	if m != nil {
		for _, e := range m.Metrics {
			return e.AcceptMetric(v)
		}
	}
	return false
}

type GetMetric interface {
	GetMetric(string) (api.Metric, bool)
}

func (m *MetricList) GetValues(f GetMetric) []value.Value {
	var r []value.Value
	if m != nil && len(m.Metrics) > 0 {
		for _, m := range m.Metrics {
			if metric, exists := f.GetMetric(m.Name); exists {
				v, _ := metric.ToValue()
				v, _ = m.Convert(v)
				r = append(r, v)
			} else {
				r = append(r, value.Value{})
			}
		}
	}
	return r
}

type Metric struct {
	Pos  lexer.Position
	Name string      `parser:"@String"`
	Unit *units.Unit `parser:"@@?"`
}

func (m *Metric) AcceptMetric(v api.Metric) bool {
	return m != nil && v.Metric == m.Name
}

func (m *Metric) Convert(v value.Value) (value.Value, error) {
	if m.Unit != nil {
		return v.As(m.Unit.Unit())
	}
	return v, nil
}

type MetricPattern struct {
	Pos     lexer.Position
	Pattern string           `parser:"@String"` // the pattern, either "", "*abc", "abd*" or "*abc*"
	Type    util.PatternType // the type of match
	Prefix  string           // prefix for server/sensor
}

func (m *MetricPattern) AcceptMetric(v api.Metric) bool {
	return m != nil && m.Match(v.Metric)
}

func (m *MetricPattern) Match(s string) bool {
	if strings.HasPrefix(s, m.Prefix) {
		return m.Type.Match(strings.TrimPrefix(s, m.Prefix), m.Pattern)
	}
	return false
}
