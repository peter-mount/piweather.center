package station

import (
	"context"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/weather/store"
	"github.com/peter-mount/piweather.center/weather/value"
)

// Reading defines a sensor available within a collection
type Reading struct {
	ID     string `json:"-" xml:"-" yaml:"-"`
	Source string `json:"source" xml:"source,attr" yaml:"source"`
	Type   string `json:"type,omitempty" xml:"type,attr,omitempty" yaml:"type,omitempty"`
	unit   *value.Unit
}

func (s *Reading) init(ctx context.Context) error {
	parent := ctx.Value("Sensors").(*Sensors)
	s.ID = parent.ID + "." + ctx.Value("ReadingId").(string)
	if u, ok := value.GetUnit(s.Type); ok {
		s.unit = u
		ctx.Value("Store").(*store.Store).DeclareReading(s.ID, s.unit)
	}
	return nil
}

func (s *Reading) process(ctx context.Context) error {
	if s.unit != nil {
		p := GetPayload(ctx)
		str, ok := p.Get(s.Source)
		if !ok {
			// FIXME warn/fail if not found?
			return nil
		}

		if f, ok := util.ToFloat64(str); ok {
			store.FromContext(ctx).Record(s.ID, s.unit.Value(f), p.Time())
		}
	}
	return nil
}
