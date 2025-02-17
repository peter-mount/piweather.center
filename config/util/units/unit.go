package units

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/weather/value"
)

// Unit allows for Unit selection
type Unit struct {
	Pos   lexer.Position
	Using string `parser:"'unit' @String"`
	unit  *value.Unit
}

func (s *Unit) Unit() *value.Unit {
	if s == nil {
		return nil
	}
	return s.unit
}

func (s *Unit) Convert(v value.Value) (value.Value, error) {
	if s == nil || s.unit == nil || !v.IsValid() {
		return v, nil
	}
	return v.As(s.unit)
}

func (s *Unit) Value(f float64) value.Value {
	if s != nil && s.unit != nil {
		return s.unit.Value(f)
	}
	return value.Value{}
}

func (s *Unit) Init() error {
	if s == nil || s.Using == "" {
		return nil
	}
	u, exists := value.GetUnit(s.Using)
	if exists {
		s.unit = u
		return nil
	}
	return participle.Errorf(s.Pos, "unsupported unit %q", s.Using)
}
