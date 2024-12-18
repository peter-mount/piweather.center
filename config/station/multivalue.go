package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/store/api"
)

type MultiValue struct {
	Pos       lexer.Position
	Component *Component     `parser:"'multivalue' '(' @@"`
	Pattern   *MetricPattern `parser:"@@"`
	Time      bool           `parser:"@'time'? ')'"`
}

func (c *visitor[T]) MultiValue(d *MultiValue) error {
	var err error
	if d != nil {
		if c.multiValue != nil {
			err = c.multiValue(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Component(d.Component)
		}

		if err == nil {
			err = c.MetricPattern(d.Pattern)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (s *initState) multiValue(_ Visitor[*initState], d *MultiValue) error {
	// Ensure Component exists, require by templates
	if d.Component == nil {
		d.Component = &Component{}
	}

	return nil
}

func (b *builder[T]) MultiValue(f func(Visitor[T], *MultiValue) error) Builder[T] {
	b.multiValue = f
	return b
}

func (c *MultiValue) AcceptMetric(v api.Metric) bool {
	return c != nil && c.Pattern.AcceptMetric(v)
}

func (c *MultiValue) GetID() string {
	return c.Component.GetID()
}

func (c *MultiValue) GetType() string {
	return "multivalue"
}
