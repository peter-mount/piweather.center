package station

import (
	"context"
	"github.com/peter-mount/piweather.center/weather/value"
)

// CalculatedValue is like a Reading, but it generates its values
// from either other Reading's or from other data like date, time etc
type CalculatedValue struct {
	ID string `json:"-" xml:"-" yaml:"-"`
	// Type of the calculated reading. This will be the name of the registered value.Calculator
	// that will be used for this calculation.
	Type string `json:"type,omitempty" xml:"type,attr,omitempty" yaml:"type,omitempty"`
	// Source lists the ID's of the Readings to pass to the Type.
	// Some types might not need any, others one or more Readings.
	// Note: The order and Value types will be specific to the value.Calculator
	Source []string `json:"source,omitempty" xml:"source,omitempty" yaml:"source,omitempty"`
	// If set, use is the case-insensitive id of the Unit that is required for the result.
	// If not set then the unit of returned valued of the Calculator is used.
	//
	// e.g. the Calculator might return a temperature in Fahrenheit, but we want Celsius.
	// In that instance Use is "Celsius".
	Use string `json:"use,omitempty" xml:"use,attr,omitempty" yaml:"use,omitempty"`
	// Graph is an optional set of graphs to be made available for this calculation.
	// These can only represent this calculation. Composite Graphs are defined elsewhere.
	Graph []*Graph `json:"graph,omitempty" xml:"graph,omitempty" yaml:"graph,omitempty"`
	// calculator to use
	calculator value.Calculator
	sensors    *Sensors
}

func CalculatedValueFromContext(ctx context.Context) *CalculatedValue {
	c := ctx.Value("CalculatedValue")
	if c == nil {
		return nil
	}
	return ctx.Value("CalculatedValue").(*CalculatedValue)
}

func (s *CalculatedValue) WithContext(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, "CalculatedValue", s), nil
}

func (s *CalculatedValue) Accept(v Visitor) error {
	return v.VisitCalculatedValue(s)
}

func (s *CalculatedValue) Calculate(t value.Time, v ...value.Value) (value.Value, error) {
	return s.calculator(t, v...)
}

func (s *CalculatedValue) Calculator() value.Calculator { return s.calculator }

// IsPseudo returns true if this is a Pseudo calculation.
//
// A Pseudo calculation is where the calculation takes no values, just the value.Time.
//
// Examples of this are calculating the altitude of an astronomical object like the sun
// in the sky, whose result is based on time and location but not any reading of any kind.
//
// If a pseudo calculation requires the Use field to be set to the unit it outputs, or the
// unit required in the output.
func (s *CalculatedValue) IsPseudo() bool {
	return s == nil || len(s.Source) == 0
}

func (s *CalculatedValue) Sensors() *Sensors { return s.sensors }