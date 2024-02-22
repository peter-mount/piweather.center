package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type UsingDefinitions struct {
	Pos  lexer.Position
	Defs []*UsingDefinition `parser:" 'declare' @@ (',' @@)* "`
}

type UsingDefinition struct {
	Pos      lexer.Position
	Name     string                `parser:"@String 'as'"`
	Modifier []*ExpressionModifier `parser:"(@@)+"`
}
