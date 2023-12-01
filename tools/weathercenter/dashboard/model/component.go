package model

import (
	"github.com/peter-mount/go-kernel/v2/log"
	uuid "github.com/peter-mount/go.uuid"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/tools/weathercenter/dashboard/registry"
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
	Live      bool `yaml:"live,omitempty"` // If true then dashboard can have live updates
}

// Container represents a collection of Components that will be rendered together
type Container struct {
	Component  `yaml:",inline"`
	Components ComponentList `yaml:"components"`
}

func (c *Container) Init() {
	c.Component.Init()
	for _, e := range c.Components {
		if d, ok := e.(*Container); ok {
			d.Init()
		} else if d, ok := e.(*Value); ok {
			d.Init()
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

// Component is the common fields available to all Component's.
//
// When defining a new component, this must be defined at the start of the struct using:
//
// Component  `yaml:",inline"`
//
// Doing this ensures that the fields are decoded correctly
type Component struct {
	ID    string `yaml:"-"`               // Unique ID, generated on load
	Type  string `yaml:"type"`            // type of component - required
	Title string `yaml:"title,omitempty"` // title, optional based on component
	Class string `yaml:"class,omitempty"` // optional CSS class
	Style string `yaml:"style,omitempty"` // optional inline CSS
}

func (c *Component) Init() {
	if c.ID == "" {
		u, _ := uuid.NewV1()
		c.ID = u.String()
	}
}

// Process a Metric
//func (c *Component) Process(m api.Metric) {
//	// Do nothing, we override this on each type
//}

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
