package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
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
			if errors.IsVisitorStop(err) {
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

func printText(v Visitor[*printState], d *Text) error {
	return v.Get().Run(d.Pos, func(st *printState) error {
		st.AppendHead("text(").
			AppendComponent(d.Component).
			AppendBody("%q", d.Text).
			AppendFooter(")")
		return nil
	})
}

func (c *Text) GetID() string {
	return c.Component.GetID()
}

func (c *Text) GetType() string {
	return c.Type
}
