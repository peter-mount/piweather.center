package model

import (
	"github.com/peter-mount/piweather.center/store/api"
	"gopkg.in/yaml.v3"
)

func init() {
	f := func() Instance { return &Container{} }
	Register("col", f)
	Register("container", f)
	Register("row", f)
}

// Container represents a collection of Components that will be rendered together
type Container struct {
	Component  `yaml:",inline"`
	Components ComponentList `yaml:"components"`
}

type initializer interface {
	init(*Dashboard)
}

func (c *Container) init(d *Dashboard) {
	c.Component.init(d)
	for _, e := range c.Components {
		if ct, ok := e.(initializer); ok {
			ct.init(d)
		}
	}
}

// Process a Metric by sending it to all Component's within the Container
func (c *Container) Process(m api.Metric, r *Response) {
	for _, e := range c.Components {
		if d, ok := e.(Processor); ok {
			d.Process(m, r)
		}
	}
}

func (c *Container) Init(db string) {
	for _, e := range c.Components {
		if d, ok := e.(Init); ok {
			d.Init(db)
		}
	}
}

// ComponentList holds a list of dynamic Component implementations.
// When a ComponentList is unmarshalled from the yaml, the components are of the correct type
// based on the Component As field.
type ComponentList []Instance

func (c *ComponentList) UnmarshalYAML(types *yaml.Node) error {
	for _, n := range types.Content {
		o, err := Decode(n)
		if err != nil {
			return err
		}

		*c = append(*c, o)
	}

	return nil
}
