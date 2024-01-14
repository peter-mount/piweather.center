package model

import (
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
	"time"
)

type Metric struct {
	Metric string      `yaml:"metric"`         // Metric ID
	Unit   string      `yaml:"unit,omitempty"` // Required Unit
	Time   string      `yaml:"time,omitempty"` // Indicate we want the time
	Value  value.Value `yaml:"-"`              // Used by templates
	metric api.Metric
}

const (
	MetricDateTime  = "datetime"
	MetricDateTimeZ = "datetimez"
	MetricDate      = "date"
	MetricDateZ     = "datez"
	MetricTime      = "time"
	MetricTimeZ     = "timez"
	MetricTimeZone  = "timezone"
)

func (m *Metric) setValue(v api.Metric) {
	if v.IsValid() {
		m.metric = v
		// Get the metric value
		var v value.Value
		if unit, ok := value.GetUnit(m.metric.Unit); ok {
			v = unit.Value(m.metric.Value)
		}

		// Requested an alternate unit then transform
		// but if the transform fails, silently ignore and use the original
		if m.Unit != "" {
			if unit, ok := value.GetUnit(m.Unit); ok {
				v1, err := v.As(unit)
				if err == nil {
					v = v1
				}
			}
		}

		m.Value = v
	}
}

func (m *Metric) MetricTime() time.Time {
	return m.metric.Time
}

func (m *Metric) TimeString() string {
	switch strings.ToLower(m.Time) {
	case MetricDateTime:
		return m.metric.Time.Format("2006-01-02T15:04:05")
	case MetricDateTimeZ:
		return m.metric.Time.Format("2006-01-02T15:04:05Z07:00")
	case MetricDate:
		return m.metric.Time.Format("2006-01-02")
	case MetricDateZ:
		return m.metric.Time.Format("2006-01-02 Z07:00")
	case MetricTime:
		return m.metric.Time.Format("15:04:05")
	case MetricTimeZ:
		return m.metric.Time.Format("15:04:05 Z07:00")
	case MetricTimeZone:
		return m.metric.Time.Format("MST")
	default:
		return ""
	}
}

// String is the same as Value.String() but returns "" if the Value is invalid
// instead of saying "invalid value"
func (m *Metric) String() string {
	if m.Value.IsValid() {
		if m.Time != "" {
			return m.TimeString()
		}
		return m.Value.String()
	}

	return ""
}

// PlainString is the same as Value.PlainString() but returns "" if the Value is invalid
// instead of saying "invalid value"
func (m *Metric) PlainString() string {
	if m.Value.IsValid() {
		return m.Value.PlainString()
	}

	return ""
}

func (m *Metric) Accept(s string) bool {
	return m != nil && m.Metric == s
}
