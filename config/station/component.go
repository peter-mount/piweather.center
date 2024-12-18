package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/store/api"
)

type ComponentType interface {
	GetType() string
}

type ComponentId interface {
	GetID() string
}

type ComponentDefinition interface {
	Definition() any
}

type ComponentProcessor interface {
	Process(ComponentStore, api.Metric)
}

type ComponentStore interface {
	// Add to a store.
	//
	// t component type
	//
	// id component id
	//
	// i & s metric index & suffix within the Component
	//
	// m api.Metric to add
	Add(t, id string, i int, s string, m api.Metric)
}

type Component struct {
	Pos lexer.Position
	//Type      string     `yaml:"type"`            // type of component - required
	Title     string     `parser:"('title' @String)?"` // title, optional based on component
	Class     string     `parser:"('class' @String)?"` // optional CSS class
	Style     string     `parser:"('style' @String)?"` // optional inline CSS
	ID        string     // Unique ID, generated on load
	dashboard *Dashboard // link to dashboard
}

func (c *visitor[T]) Component(d *Component) error {
	var err error
	if d != nil {
		if c.component != nil {
			err = c.component(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (b *builder[T]) Component(f func(Visitor[T], *Component) error) Builder[T] {
	b.component = f
	return b
}

func (c *Component) GetID() string {
	return c.ID
}

func (c *Component) GetType() string {
	return "component"
}
