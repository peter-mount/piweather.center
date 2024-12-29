package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/weather/value"
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

func (c *visitor[T]) MetricList(d *MetricList) error {
	var err error
	if d != nil {
		if c.metricList != nil {
			err = c.metricList(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		for _, s := range d.Metrics {
			err = c.Metric(s)
			if err != nil {
				break
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (b *builder[T]) MetricList(f func(Visitor[T], *MetricList) error) Builder[T] {
	b.metricList = f
	return b
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
