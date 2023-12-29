package model

import (
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/weather/value"
)

type Metric struct {
	Metric string      `yaml:"metric"`         // Metric ID
	Unit   string      `yaml:"unit,omitempty"` // Required Unit
	Value  value.Value `yaml:"-"`              // Used by templates
	metric api.Metric
}

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

// String is the same as Value.String() but returns "" if the Value is invalid
// instead of saying "invalid value"
func (m *Metric) String() string {
	if m.Value.IsValid() {
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
