package model

import (
	uuid "github.com/peter-mount/go.uuid"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/tools/weathercenter/dashboard/registry"
	"gopkg.in/yaml.v3"
	"strconv"
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
	Live      bool   `yaml:"live,omitempty"` // If true then dashboard can have live updates
	Uuid      string `yaml:"-"`              // Uuid of dashboard - generated
	idSeq     int    // Used in initialising the ID's
	uuid      uuid.UUID
}

func (c *Dashboard) Init(uuid uuid.UUID) {
	c.uuid = uuid
	c.Uuid = uuid.String()
	c.idSeq = 0
	c.Container.init(c)
}

func (c *Dashboard) NextId() string {
	c.idSeq++
	return uuid.NewV3(c.uuid, strconv.Itoa(c.idSeq)).String()
}

func (c *Dashboard) Process(m api.Metric, r *Response) {
	// Set the response Uuid to the Dashboard.
	// This allows the front end to detect a dashboard change.
	r.Uuid = c.Uuid
	c.Container.Process(m, r)
}

// Container represents a collection of Components that will be rendered together
type Container struct {
	Component  `yaml:",inline"`
	Components ComponentList `yaml:"components"`
}

func (c *Container) init(d *Dashboard) {
	c.Component.init(d)
	for _, e := range c.Components {
		if ct, ok := e.(*Container); ok {
			ct.init(d)
		} else if ct, ok := e.(*Value); ok {
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

func (c *Component) init(d *Dashboard) {
	c.ID = d.NextId()
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
