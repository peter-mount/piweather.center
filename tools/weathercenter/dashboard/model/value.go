package model

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/tools/weathercenter/dashboard/registry"
	"github.com/peter-mount/piweather.center/weather/value"
)

func init() {
	f := func() registry.Component { return &Value{} }
	registry.Register("value", f)
	registry.Register("rain-gauge", f)
	registry.Register("compass", f)
}

// Value represents a distinct component displaying values
type Value struct {
	Component `yaml:",inline"`
	Label     string   `yaml:"label"`             // optional label
	Min       *float64 `yaml:"min,omitempty"`     // Min axis value
	Max       *float64 `yaml:"max,omitempty"`     // Max axis value
	Metric    *Metric  `yaml:"metric,omitempty"`  // Single metric for value
	Metrics   []Metric `yaml:"metrics,omitempty"` // Multiple metrics
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
	update := false

	if c.Metric.Accept(m.Metric) {
		log.Printf("%q %v", c.Metric.Metric, m)
		c.Metric.setValue(m)
		update = update || c.Metric.Value.IsValid()
	}

	if len(c.Metrics) > 0 {
		for i, e := range c.Metrics {
			e.setValue(m)
			update = update || e.Value.IsValid()

			// required as e is not a pointer
			c.Metrics[i] = e
		}
	}

	// If notifiable then add it to the response
	if update {
		r.Add(c.Type, c.ID)
	}
}
