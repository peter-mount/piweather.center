package model

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

func (c *Component) Accept(v Visitor) error {
	if v == nil {
		return nil
	}
	return v(c)
}
