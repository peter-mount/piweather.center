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
	ID      string `json:"-" xml:"-" yaml:"-"`
	Source  string `json:"source" xml:"source,attr" yaml:"source"`
	Type    string `json:"type,omitempty" xml:"type,attr,omitempty" yaml:"type,omitempty"`
	Use     string `json:"use,omitempty" xml:"use,attr,omitempty" yaml:"use,omitempty"`
	unit    *value.Unit
	useUnit *value.Unit
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
