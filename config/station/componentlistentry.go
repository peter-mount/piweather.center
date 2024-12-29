package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/store/api"
)

type ComponentListEntry struct {
	Pos        lexer.Position
	Container  *Container  `parser:"( @@"`
	Gauge      *Gauge      `parser:"| @@"`
	MultiValue *MultiValue `parser:"| @@"`
	Text       *Text       `parser:"| @@"`
	Value      *Value      `parser:"| @@ )"`
}

func (c *visitor[T]) ComponentListEntry(d *ComponentListEntry) error {
	var err error
	if d != nil {
		if c.componentListEntry != nil {
			err = c.componentListEntry(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Container(d.Container)
		}

		if err == nil {
			err = c.Gauge(d.Gauge)
		}

		if err == nil {
			err = c.MultiValue(d.MultiValue)
		}

		if err == nil {
			err = c.Text(d.Text)
		}

		if err == nil {
			err = c.Value(d.Value)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (b *builder[T]) ComponentListEntry(f func(Visitor[T], *ComponentListEntry) error) Builder[T] {
	b.componentListEntry = f
	return b
}

func (c *ComponentListEntry) AcceptMetric(v api.Metric) bool {
	switch {
	case c.Gauge != nil:
		return c.Gauge.AcceptMetric(v)
	case c.MultiValue != nil:
		return c.MultiValue.AcceptMetric(v)
	case c.Value != nil:
		return c.Value.AcceptMetric(v)
	default:
		return false
	}
}

func (c *ComponentListEntry) GetID() string {
	switch {
	case c.Gauge != nil:
		return c.Gauge.Component.GetID()
	case c.MultiValue != nil:
		return c.MultiValue.Component.GetID()
	case c.Text != nil:
		return c.Text.GetID()
	case c.Value != nil:
		return c.Value.Component.GetID()
	default:
		return ""
	}
}

func (c *ComponentListEntry) GetType() string {
	switch {
	case c.Container != nil:
		return c.Container.GetType()
	case c.Gauge != nil:
		return c.Gauge.GetType()
	case c.MultiValue != nil:
		return c.MultiValue.GetType()
	case c.Text != nil:
		return c.Text.GetType()
	case c.Value != nil:
		return c.Value.GetType()
	default:
		return ""
	}
}
