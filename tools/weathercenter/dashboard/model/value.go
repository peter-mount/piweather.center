package model

import (
	"github.com/peter-mount/piweather.center/store/api"
)

func init() {
	f := func() Instance { return &Value{} }
	Register("barometer", f)
	Register("compass", f)
	Register("gauge", f)
	Register("inclinometer", f)
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
	Metric    []Metric `yaml:"metric,omitempty"` // Multiple Metrics
}

// Process a Metric
func (c *Value) Process(m api.Metric, r *Response) {
	if len(c.Metric) > 0 {
		a := Action{ID: c.ID}

		for i, e := range c.Metric {
			var s string
			var post bool

			switch {
			// Time set and no metric then use the most recent time as the value
			case e.Time != "" && e.Metric == "":
				t := e.MetricTime()
				if t.IsZero() || t.Before(m.Time) {
					e.setValue(m)
					s = e.TimeString()
					post = true
				}

			case e.Metric == "":
				break

			case e.Accept(m.Metric):
				e.setValue(m)
				s = e.Value.String()
				post = true

			default:
				break
			}

			if post {
				if e.Value.IsValid() {
					a = a.Add(i, api.Metric{
						Metric:    m.Metric,
						Time:      m.Time,
						Unit:      e.Value.Unit().ID(),
						Value:     e.Value.Float(),
						Formatted: s,
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
