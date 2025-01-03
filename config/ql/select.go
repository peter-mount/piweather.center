package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
)

type Select struct {
	Pos lexer.Position

	Expression *SelectExpression `parser:"'select' @@"`
	Limit      int               `parser:"( 'limit' @Number )?"`
}

func (v *visitor[T]) Select(b *Select) error {
	var err error
	if b != nil {
		if v._select != nil {
			err = v._select(v, b)
			if errors.IsVisitorStop(err) || errors.IsVisitorExit(err) {
				return nil
			}
		}

		if err == nil {
			err = v.SelectExpression(b.Expression)
		}
	}
	return err
}

func initSelect(_ Visitor[*parserState], s *Select) error {
	return assertLimit(s.Pos, s.Limit)
}

func (b *builder[T]) Select(f func(Visitor[T], *Select) error) Builder[T] {
	b.common._select = f
	return b
}
