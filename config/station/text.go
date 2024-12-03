package station

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type Text struct {
	Pos       lexer.Position
	Type      string     `parser:"@('text') '('"`
	Component *Component `parser:"@@"`
	Text      string     `parser:"@String ')'"`
}

func (c *Text) GetID() string {
	return c.Component.GetID()
}

func (c *Text) GetType() string {
	return c.Type
}
