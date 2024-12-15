package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
)

type Expression struct {
	Pos  lexer.Position
	Left *ExpressionLevel1 `parser:"@@"`
}

func (c *visitor[T]) Expression(b *Expression) error {
	var err error
	if b != nil {
		if c.expression != nil {
			err = c.expression(c, b)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.ExpressionLevel1(b.Left)
		}

		err = errors.Error(b.Pos, err)
	}
	return err
}

func (b *builder[T]) Expression(f func(Visitor[T], *Expression) error) Builder[T] {
	b.expression = f
	return b
}
