package station

import (
	"context"
	"github.com/peter-mount/piweather.center/server/store"
	"github.com/peter-mount/piweather.center/station/payload"
	"github.com/peter-mount/piweather.center/util"
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
}

func ReadingFromContext(ctx context.Context) *Reading {
	return ctx.Value("Reading").(*Reading)
}

func (s *Reading) WithContext(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, "Reading", s), nil
}

func (s *Reading) Accept(v Visitor) error {
	return v.VisitReading(s)
}

func InitReading(ctx context.Context) error {
	parent := SensorsFromContext(ctx)
	s := ReadingFromContext(ctx)
	s.ID = parent.ID + "." + ctx.Value("ReadingId").(string)

	// If not ok then we will ignore the reading
	if u, ok := value.GetUnit(s.Type); ok {
		s.unit = u

		// Use the same unit, unless we declare an alternate
		s.useUnit = u
		if s.Use != "" {
			// TODO if the unit is not ok or there's no transform then this will fail over to use the src unit
			if u, ok := value.GetUnit(s.Use); ok && value.TransformAvailable(s.useUnit, u) {
				s.useUnit = u
			}
		}

		// Register the reading with the final unit
		store.FromContext(ctx).DeclareReading(s.ID, s.useUnit)
	}
	return nil
}

func ProcessReading(ctx context.Context) error {
	s := ReadingFromContext(ctx)
	if s.unit != nil {
		p := payload.GetPayload(ctx)

		str, ok := p.Get(s.Source)
		if !ok {
			// FIXME warn/fail if not found?
			return nil
		}

		if f, ok := util.ToFloat64(str); ok {
			// Convert to Type unit then transform to Use unit
			v, err := s.unit.Value(f).As(s.useUnit)
			if err != nil {
				// Ignore, should only happen if the result is
				// invalid as we checked the transform previously
				return nil
			}

			store.FromContext(ctx).Record(s.ID, v, p.Time())
		}
	}
	return nil
}
