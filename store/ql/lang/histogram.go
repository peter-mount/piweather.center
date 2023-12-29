package lang

import "github.com/alecthomas/participle/v2/lexer"

type Histogram struct {
	Pos lexer.Position

	Expression *AliasedExpression `parser:"'HISTOGRAM' @@"`
}

func (a *Histogram) Accept(v Visitor) error {
	return v.Histogram(a)
}

type WindRose struct {
	Pos lexer.Position

	Degrees *Expression      `parser:"'WINDROSE' @@"`
	Speed   *Expression      `parser:"',' @@"`
	Options []WindRoseOption `parser:"('AS' @@ (',' @@)+ )?"`
}

type WindRoseOption struct {
	Pos   lexer.Position
	Rose  bool `parser:"( @'ROSE'"`
	Table bool `parser:"| @'TABLE')"`
}

func (a *WindRose) Accept(v Visitor) error {
	return v.WindRose(a)
}
