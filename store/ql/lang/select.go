package lang

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/weather/value"
)

type Select struct {
	Pos lexer.Position

	Expression *SelectExpression `parser:"'SELECT' @@"`
	Limit      int               `parser:"( 'LIMIT' @Int )?"`
}

func (a *Select) Accept(v Visitor) error {
	return v.Select(a)
}

type SelectExpression struct {
	Pos lexer.Position

	Expressions []*AliasedExpression `parser:"@@ ( ',' @@ )*"`
}

func (a *SelectExpression) Accept(v Visitor) error {
	return v.SelectExpression(a)
}

// AliasedExpression handles expression AS name to create aliases
type AliasedExpression struct {
	Pos lexer.Position

	Expression *Expression `parser:"@@"`
	Unit       string      `parser:"( 'UNIT' @String )?"`
	As         string      `parser:"( 'AS' @String )?"`
	unit       *value.Unit
}

func (a *AliasedExpression) Accept(v Visitor) error {
	return v.AliasedExpression(a)
}

func (a *AliasedExpression) GetUnit() *value.Unit {
	if a == nil {
		return nil
	}
	return a.unit
}

func (a *AliasedExpression) SetUnit(u *value.Unit) {
	a.unit = u
}

// Expression handles function calls or direct metric values
type Expression struct {
	Pos      lexer.Position
	Function *Function             `parser:"( @@"`
	Metric   *Metric               `parser:"| @@ )"`
	Using    string                `parser:"( 'USING' @String"`
	Modifier []*ExpressionModifier `parser:"| (@@)+ )?"`
}

func (a *Expression) Accept(v Visitor) error {
	return v.Expression(a)
}

type ExpressionModifier struct {
	Pos    lexer.Position
	Range  *QueryRange `parser:"( @@"`
	Offset *Duration   `parser:"| 'OFFSET' @@ )"`
}

// Function handles function calls
type Function struct {
	Pos lexer.Position

	Name        string        `parser:"@Ident"`
	Expressions []*Expression `parser:"'(' (@@ (',' @@)*)? ')'"`
}

func (a *Function) Accept(v Visitor) error {
	return v.Function(a)
}

// Metric handles a metric reference
type Metric struct {
	Pos lexer.Position

	Metric []string `parser:"@Ident ( '.' @Ident )*"`
	Name   string
}

func (a *Metric) Accept(v Visitor) error {
	return v.Metric(a)
}
