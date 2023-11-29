package model

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/dashboard/registry"
	"gopkg.in/yaml.v3"
)

func init() {
	f := func() registry.Component { return &Container{} }
	registry.Register("container", f)
	registry.Register("row", f)
}

// Dashboard is the top level Component representing an entire dashboard.
// Initially it's the same as Container, however it will have additional fields in the future.
type Dashboard struct {
	Container `yaml:",inline"`
}

// Container represents a collection of Components that will be rendered together
type Container struct {
	Component  `yaml:",inline"`
	Components ComponentList `yaml:"components"`
}

// Component is the common fields available to all Component's.
//
// When defining a new component, this must be defined at the start of the struct using:
//
// Component  `yaml:",inline"`
//
// Doing this ensures that the fields are decoded correctly
type Component struct {
	Type  string `yaml:"type"`            // type of component - required
	Title string `yaml:"title,omitempty"` // title, optional based on component
	Style string `yaml:"style,omitempty"` // optional css for this component only
}

// GetType returns the type of component
func (c *Component) GetType() string {
	if c == nil {
		return ""
	}
	return c.Type
}

func (c *Component) Accept(v registry.Visitor) error {
	if v == nil {
		return nil
	}
	return v(c)
}

// ComponentList holds a list of dynamic Component implementations.
// When a ComponentList is unmarshalled from the yaml, the components are of the correct type
// based on the Component Type field.
type ComponentList []registry.Component

func (c *ComponentList) UnmarshalYAML(unmarshal func(interface{}) error) error {
	types := make([]yaml.Node, 0)
	err := unmarshal(&types)
	if err != nil {
		return err
	}

	for _, n := range types {
		o, err := registry.Decode(n)
		if err != nil {
			return err
		}

		*c = append(*c, o)
	}

	return nil
}

func Debug(d *Dashboard) {
	log.Printf("Dashboard %q %q", d.Type, d.Title)
	debugList(0, d.Components)
}
func debugList(level int, l ComponentList) {
	log.Printf("%03d START", level)
	for _, child := range l {
		log.Printf("%03d found %q", level, child.GetType())
		switch child.GetType() {
		case "container":
			debugList(level+1, child.(*Container).Components)
		}
	}
}
