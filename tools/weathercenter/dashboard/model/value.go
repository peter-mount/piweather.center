package model

import (
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/weather/value"
)

func init() {
	f := func() Instance { return &Value{} }
	Register("barometer", f)
	Register("compass", f)
	Register("gauge", f)
	Register("rain-gauge", f)
	Register("value", f)
}

// Value represents a distinct component displaying values
type Value struct {
	Component `yaml:",inline"`
	Label     string   `yaml:"label"`            // optional label
	Min       *float64 `yaml:"min,omitempty"`    // Min axis value
	Max       *float64 `yaml:"max,omitempty"`    // Max axis value
	Ticks     *float64 `yaml:"ticks,omitempty"`  // Number of ticks on axis
	Unit      string   `yaml:"unit,omitempty"`   // Unit for display, defaults to first metric
	Metric    []Metric `yaml:"metric,omitempty"` // Multiple metrics
}

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

// Process a Metric
func (c *Value) Process(m api.Metric, r *Response) {
	if len(c.Metric) > 0 {
		a := Action{ID: c.ID}

		for i, e := range c.Metric {
			if e.Accept(m.Metric) {
				e.setValue(m)
				if e.Value.IsValid() {
					a = a.Add(i, api.Metric{
						Metric:    m.Metric,
						Time:      m.Time,
						Unit:      e.Value.Unit().ID(),
						Value:     e.Value.Float(),
						Formatted: e.Value.String(),
						Unix:      m.Unix,
					})
				}

				// required as e is not a pointer
				c.Metric[i] = e
			}
		}
		r.Add(c.Type, a)
	}
}
