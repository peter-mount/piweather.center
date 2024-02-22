package units

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/weather/value"
)

// Unit allows for Unit selection
type Unit struct {
	Pos   lexer.Position
	Using string `parser:"'using' @String"`
	unit  *value.Unit
}

func (s *Unit) Unit() *value.Unit {
	return s.unit
}

func (s *Unit) Init() error {
	u, exists := value.GetUnit(s.Using)
	if exists {
		s.unit = u
		return nil
	}
	return participle.Errorf(s.Pos, "unsupported unit %q", s.Using)
}
