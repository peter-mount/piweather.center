package station

import (
	"context"
	"github.com/peter-mount/piweather.center/weather/value"
)

// Reading defines a sensor available within a collection
type Reading struct {
	ID string `json:"-" xml:"-" yaml:"-"`
	// Source the source within the received data for a reading
	Source string `json:"source" xml:"source,attr" yaml:"source"`
	// Type of the reading. This must match the case-insensitive id of a Unit.
	// If absent or not a valid Unit id, then this is taken as a placeholder and the reading/graphs are ignored.
	Type string `json:"type,omitempty" xml:"type,attr,omitempty" yaml:"type,omitempty"`
	// If set, use is the case-insensitive id of the Unit that is required for the Reading.
	// If not set then Type is the unit used.
	//
	// e.g. the device might provide temperature in Fahrenheit, but we want Celsius.
	// In that instance Type is "Fahrenheit" and Use is "Celsius".
	Use string `json:"use,omitempty" xml:"use,attr,omitempty" yaml:"use,omitempty"`
	// Graph is an optional set of graphs to be made available for this reading.
	// These can only represent this reading. Composite Graphs are defined elsewhere.
	Graph []*Graph `json:"graph,omitempty" xml:"graph,omitempty" yaml:"graph,omitempty"`
	// unit resolved from Type
	unit *value.Unit
	// useUnit either unit from Type or resolved from Use if defined.
	useUnit *value.Unit
	sensors *Sensors
}

func ReadingFromContext(ctx context.Context) *Reading {
	r := ctx.Value("Reading")
	if r == nil {
		return nil
	}
	return r.(*Reading)
}

func (s *Reading) WithContext(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, "Reading", s), nil
}

func (s *Reading) Accept(v Visitor) error {
	return v.VisitReading(s)
}

// Value returns f in the Type unit and returns the Value in the Use unit.
func (s *Reading) Value(f float64) (value.Value, error) {
	return s.unit.Value(f).As(s.useUnit)
}

func (s *Reading) Unit() *value.Unit {
	return s.useUnit
}

func (s *Reading) Sensors() *Sensors { return s.sensors }
