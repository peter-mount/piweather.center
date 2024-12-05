package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/units"
)

// AliasedExpression handles expression AS name to create aliases
type AliasedExpression struct {
	Pos lexer.Position

	Expression *Expression `parser:"@@"`
	Unit       *units.Unit `parser:"( @@ )?"`
	As         string      `parser:"( 'as' @String )?"`
	Summarize  *Summarize  `parser:"( @@ )?"`
}

func (v *visitor[T]) AliasedExpression(b *AliasedExpression) error {
	var err error
	if b != nil {
		if v.aliasedExpression != nil {
			err = v.aliasedExpression(v, b)
		}
		if util.IsVisitorStop(err) {
			return nil
		}
		if err == nil {
			err = v.Unit(b.Unit)
		}
		if err == nil {
			err = v.Expression(b.Expression)
		}
		if err == nil {
			err = v.Summarize(b.Summarize)
		}
	}
	return err
}

func (b *builder[T]) AliasedExpression(f func(Visitor[T], *AliasedExpression) error) Builder[T] {
	b.common.aliasedExpression = f
	return b
}