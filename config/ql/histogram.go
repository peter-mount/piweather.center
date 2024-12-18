package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
)

type Histogram struct {
	Pos lexer.Position

	Expression *AliasedExpression `parser:"'histogram' @@"`
}

func (v *visitor[T]) Histogram(b *Histogram) error {
	var err error
	if b != nil {
		if v.histogram != nil {
			err = v.histogram(v, b)
			if errors.IsVisitorStop(err) || errors.IsVisitorExit(err) {
				return nil
			}
		}
		if err == nil {
			err = v.AliasedExpression(b.Expression)
		}
	}
	return err
}

func (b *builder[T]) Histogram(f func(Visitor[T], *Histogram) error) Builder[T] {
	b.common.histogram = f
	return b
}
