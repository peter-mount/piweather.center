package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util"
)

type Histogram struct {
	Pos lexer.Position

	Expression *AliasedExpression `parser:"'histogram' @@"`
}

func (v *visitor) Histogram(b *Histogram) error {
	var err error
	if b != nil {
		if v.histogram != nil {
			err = v.histogram(v, b)
			if util.IsVisitorStop(err) || util.IsVisitorExit(err) {
				return nil
			}
		}
		if err == nil {
			err = v.AliasedExpression(b.Expression)
		}
	}
	return err
}

func (b *builder) Histogram(f func(Visitor, *Histogram) error) Builder {
	b.common.histogram = f
	return b
}
