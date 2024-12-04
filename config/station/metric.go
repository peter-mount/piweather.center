package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util/units"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
	"time"
)

// MetricAccept is an interface to a type that can test to see if it would accept the supplied metric
type MetricAccept interface {
	// AcceptMetric returns true if this type accepts the named metric
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

// GetMetric is an interface a type implements to retrieve a metric value
type GetMetric interface {
	// GetMetric returns an api.Metric by name
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

// Metric represents a metric the containing type requires.
// This is used either as inbound, e.g. we need to retrieve the named metric,
// or outbound as in the target metric in a calculation
type Metric struct {
	Pos  lexer.Position
	Name string      `parser:"@String"` // metric name
	Unit *units.Unit `parser:"@@?"`     // optional Unit we require the metric to be in
}

func (m *Metric) AcceptMetric(v api.Metric) bool {
	return m != nil && v.Metric == m.Name
}

// Convert converts the passed Value to that of the Metric based on the requested Unit
func (m *Metric) Convert(v value.Value) (value.Value, error) {
	if m.Unit != nil {
		return v.As(m.Unit.Unit())
	}
	return v, nil
}

// MetricExpression represents a Metric within a Calculation
type MetricExpression struct {
	Pos    lexer.Position
	Metric *Metric `parser:"@@"`                    // Metric reference
	Offset string  `parser:"( 'offset' @String )?"` // optional offset in time, usually negative in the past
	// the parsed value of Offset
	offset time.Duration
}

// HasOffset returns true of Offset is defined and the parsed value is not 0
func (m *MetricExpression) HasOffset() bool {
	return m.offset != 0
}

// GetOffset returns the parsed offset, 0 if undefined
func (m *MetricExpression) GetOffset() time.Duration {
	return m.offset
}

// MetricPattern defines we require any metrics that match a simple pattern, either "*", "*abc", "abd*" or "*abc*".
// Note: "" or "**" is equivalent to "*".
type MetricPattern struct {
	Pos     lexer.Position
	Pattern string `parser:"@String"` // the pattern
	// the type of match
	patternType util.PatternType
	// prefix for server/sensor
	prefix string
}

func (m *MetricPattern) AcceptMetric(v api.Metric) bool {
	return m != nil && m.Match(v.Metric)
}

// Match returns true if the metric name passed matches the pattern
func (m *MetricPattern) Match(s string) bool {
	if strings.HasPrefix(s, m.prefix) {
		return m.patternType.Match(strings.TrimPrefix(s, m.prefix), m.Pattern)
	}
	return false
}
