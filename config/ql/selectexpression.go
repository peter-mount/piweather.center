package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util"
)

type SelectExpression struct {
	Pos lexer.Position

	Expressions []*AliasedExpression `parser:"@@ ( ',' @@ )*"`
}

func (v *visitor[T]) SelectExpression(b *SelectExpression) error {
	var err error
	if b != nil {
		if v.selectExpression != nil {
			err = v.selectExpression(v, b)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			for _, e := range b.Expressions {
				err = v.AliasedExpression(e)
				if err != nil {
					return err
				}
			}
		}
	}
	return err
}

func (b *builder[T]) SelectExpression(f func(Visitor[T], *SelectExpression) error) Builder[T] {
	b.common.selectExpression = f
	return b
}
