package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
)

type Container struct {
	Pos        lexer.Position
	Type       string         `parser:"@('container' | 'col' | 'row') '('"`
	Component  *Component     `parser:"@@"`
	Components *ComponentList `parser:"@@ ')'"`
}

func (c *visitor[T]) Container(d *Container) error {
	var err error
	if d != nil {
		if c.container != nil {
			err = c.container(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = visitContainer[T](c, d)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}
func visitContainer[T any](v Visitor[T], d *Container) error {
	var err error
	if d != nil {
		err = v.Component(d.Component)

		if err == nil {
			err = v.ComponentList(d.Components)
		}
	}
	return err
}

func initContainer(_ Visitor[*initState], d *Container) error {
	// Ensure Component exists, require by templates
	if d.Component == nil {
		d.Component = &Component{}
	}
	// Ensure we have an entry present so we don't need to check this in templates
	if d.Components == nil {
		d.Components = &ComponentList{}
	}

	return nil
}

func (b *builder[T]) Container(f func(Visitor[T], *Container) error) Builder[T] {
	b.container = f
	return b
}

func printContainer(v Visitor[*printState], d *Container) error {
	return v.Get().
		Start().
		AppendHead("%s(", d.Type).
		AppendFooter(")").
		EndError(d.Pos, visitContainer(v, d))
}

func (c *Container) GetType() string {
	return c.Type
}
