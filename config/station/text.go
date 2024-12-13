package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
)

type Text struct {
	Pos       lexer.Position
	Type      string     `parser:"@('text') '('"`
	Component *Component `parser:"@@"`
	Text      string     `parser:"@String ')'"`
}

func (c *visitor[T]) Text(d *Text) error {
	var err error
	if d != nil {
		if c.text != nil {
			err = c.text(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Component(d.Component)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (b *builder[T]) Text(f func(Visitor[T], *Text) error) Builder[T] {
	b.text = f
	return b
}

func (c *Text) GetID() string {
	return c.Component.GetID()
}

func (c *Text) GetType() string {
	return c.Type
}
