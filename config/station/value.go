package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/store/api"
)

type Value struct {
	Pos       lexer.Position
	Type      string      `parser:"@('value') '('"`
	Component *Component  `parser:"@@"`
	Label     string      `parser:"@String"`
	Metrics   *MetricList `parser:"@@ ')'"`
}

func (c *visitor[T]) Value(d *Value) error {
	var err error
	if d != nil {
		if c.value != nil {
			err = c.value(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Component(d.Component)
		}

		if err == nil {
			err = c.MetricList(d.Metrics)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func initValue(_ Visitor[*initState], d *Value) error {
	// Ensure Component exists, require by templates
	if d.Component == nil {
		d.Component = &Component{}
	}
	return nil
}

func (b *builder[T]) Value(f func(Visitor[T], *Value) error) Builder[T] {
	b.value = f
	return b
}

func (c *Value) AcceptMetric(v api.Metric) bool {
	return c != nil && c.Metrics.AcceptMetric(v)
}

func (c *Value) GetID() string {
	return c.Component.GetID()
}

func (c *Value) GetType() string {
	return c.Type
}
