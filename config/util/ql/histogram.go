package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type Histogram struct {
	Pos lexer.Position

	Expression *AliasedExpression `parser:"'histogram' @@"`
}

type WindRose struct {
	Pos lexer.Position

	Degrees *Expression      `parser:"'windrose' @@"`
	Speed   *Expression      `parser:"',' @@"`
	Options []WindRoseOption `parser:"('AS' @@ (',' @@)* )?"`
}

type WindRoseOption struct {
	Pos   lexer.Position
	Rose  bool `parser:"( @'rose'"`
	Count bool `parser:"| @'count'"`
	Max   bool `parser:"| @'max')"`
}
