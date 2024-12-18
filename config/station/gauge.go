package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util/units"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/weather/value"
)

type Gauge struct {
	Pos       lexer.Position
	Type      string      `parser:"@('gauge'|'barometer'|'compass'|'inclinometer'|'raingauge') '('"`
	Component *Component  `parser:"@@"`
	Label     string      `parser:"@String"`
	Unit      *units.Unit `parser:"(@@)?"`
	Axis      *Axis       `parser:"(@@)?"`
	Metrics   *MetricList `parser:"@@ ')'"`
}

func (c *visitor[T]) Gauge(d *Gauge) error {
	var err error
	if d != nil {
		if c.gauge != nil {
			err = c.gauge(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Component(d.Component)
		}

		if err == nil {
			err = c.Unit(d.Unit)
		}

		if err == nil {
			err = c.Axis(d.Axis)
		}

		if err == nil {
			err = c.MetricList(d.Metrics)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func initGauge(_ Visitor[*initState], d *Gauge) error {
	var err error

	// Ensure Component exists, require by templates
	if d.Component == nil {
		d.Component = &Component{}
	}

	if d.Metrics == nil || len(d.Metrics.Metrics) == 0 {
		// We must have at least 1 metric for gauges
		err = errors.Errorf(d.Pos, "No metrics provided for Gauge")
	}

	return errors.Error(d.Pos, err)
}

func (b *builder[T]) Gauge(f func(Visitor[T], *Gauge) error) Builder[T] {
	b.gauge = f
	return b
}

func (c *Gauge) AcceptMetric(v api.Metric) bool {
	return c != nil && c.Metrics.AcceptMetric(v)
}

func (c *Gauge) GetID() string {
	return c.Component.GetID()
}

func (c *Gauge) GetType() string {
	return c.Type
}

func (c *Gauge) Convert(v value.Value) (value.Value, error) {
	var err error

	// Convert v to either the specified unit or that of the first metric
	if c.Unit == nil {
		// This is safe as the parser ensures that metrics contain at least 1 metric for gauges
		v, err = c.Metrics.Metrics[0].Convert(v)
	} else {
		v, err = c.Unit.Convert(v)
	}

	return v, errors.Error(c.Pos, err)
}

func (c *Gauge) ConvertAll(vals []value.Value) ([]value.Value, error) {
	var err error

	if c.Unit != nil {
		for i, v := range vals {
			vals[i], err = c.Unit.Convert(v)
			if err != nil {
				return nil, errors.Error(c.Metrics.Metrics[i].Pos, err)
			}
		}
	}

	return vals, nil
}
