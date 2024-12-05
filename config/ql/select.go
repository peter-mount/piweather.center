package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util"
)

type Select struct {
	Pos lexer.Position

	Expression *SelectExpression `parser:"'select' @@"`
	Limit      int               `parser:"( 'limit' @Int )?"`
}

func (v *visitor[T]) Select(b *Select) error {
	var err error
	if b != nil {
		if v._select != nil {
			err = v._select(v, b)
			if util.IsVisitorStop(err) || util.IsVisitorExit(err) {
				return nil
			}
		}

		if err == nil {
			err = v.SelectExpression(b.Expression)
		}
	}
	return err
}

func (b *builder[T]) Select(f func(Visitor[T], *Select) error) Builder[T] {
	b.common._select = f
	return b
}
