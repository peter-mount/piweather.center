package model

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/store/api"
	"strings"
)

func init() {
	f := func() Instance { return &MultiValue{} }
	Register("multivalue", f)
}

// MultiValue is similar to Value except it works on a wildcard when
// retrieving metric values and uses the metric name and value as the columns
// instead of Label and Metric within Value.
//
// Optionally, you can request the time of the current value as a third column.
//
// Metrics are shown in alphabetical order.
type MultiValue struct {
	Component   `yaml:",inline"`
	Pattern     string   `yaml:"pattern"` // Metric pattern
	Time        bool     `yaml:"time"`    // Time included if true
	Metrics     []Metric // list of metric values
	pattern     string   // Pattern in Metric minus trailing .*
	initialized bool     // set after setup
}

func (c *MultiValue) init(d *Dashboard) {
	c.Component.init(d)

	pattern := strings.TrimSpace(c.Pattern)
	log.Printf("****\n%q pattern %q", c.ID, c.pattern)

	switch {
	case pattern == "*":
		// matches all metrics
		c.pattern = ""

	case strings.HasSuffix(pattern, ".*"):
		// set pattern to be string before *
		c.pattern = strings.TrimSuffix(pattern, "*")

	default:
		// Bail out as not set
		return
	}

	c.initialized = true
	log.Printf("****\n%q pattern %q init %v", c.ID, c.pattern, c.initialized)
}

// Process a Metric
func (c *MultiValue) Process(m api.Metric, r *Response) {
	if !c.initialized {
		return
	}

	if c.pattern == "" || strings.HasPrefix(m.Metric, c.pattern) {
		a := Action{ID: c.ID}

		found := false
		for i, e := range c.Metrics {
			if e.Accept(m.Metric) {
				found = true
				e.setValue(m)
				a = c.add(a, i, e)

				// Must update as slice is not a pointer
				c.Metrics[i] = e
			}
		}

		// not found then add it - will happen on first run
		if !found {
			e := NewMetric(m)
			c.Metrics = append(c.Metrics, e)
			a = c.add(a, len(c.Metrics)-1, e)
		}

		r.Add(c.Type, a)
	}
}

func (c *MultiValue) add(a Action, i int, e Metric) Action {
	if e.Value.IsValid() {
		m := api.Metric{
			Metric:    e.Metric,
			Time:      e.metric.Time,
			Unit:      e.Value.Unit().ID(),
			Value:     e.Value.Float(),
			Formatted: e.Value.String(),
			Unix:      e.metric.Unix,
		}
		a = a.Add(i, m)

		if c.Time {
			m.Formatted = e.TimeString()
			a = a.AddSuffix(i, "T", m)
		}
	}

	return a
}
