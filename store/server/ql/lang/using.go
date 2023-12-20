package lang

import "github.com/alecthomas/participle/v2/lexer"

type UsingDefinitions struct {
	Pos  lexer.Position
	Defs []*UsingDefinition `parser:" 'DECLARE' @@ (',' @@)* "`
}

func (u *UsingDefinitions) Accept(v Visitor) error {
	return v.UsingDefinitions(u)
}

type UsingDefinition struct {
	Pos      lexer.Position
	Name     string                `parser:"@String 'AS'"`
	Modifier []*ExpressionModifier `parser:"(@@)+"`
}

func (u *UsingDefinition) Accept(v Visitor) error {
	return v.UsingDefinition(u)
}
