package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type UsingDefinitions struct {
	Pos  lexer.Position
	Defs []*UsingDefinition `parser:" 'DECLARE' @@ (',' @@)* "`
}

type UsingDefinition struct {
	Pos      lexer.Position
	Name     string                `parser:"@String 'AS'"`
	Modifier []*ExpressionModifier `parser:"(@@)+"`
}
