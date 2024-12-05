package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/time"
)

type ExpressionModifier struct {
	Pos    lexer.Position
	Range  *QueryRange    `parser:"( @@"`
	Offset *time.Duration `parser:"| 'offset' @@ )"`
}

func (v *visitor) ExpressionModifier(b *ExpressionModifier) error {
	var err error
	if b != nil {
		if v.expressionModifier != nil {
			err = v.expressionModifier(v, b)
		}
		if util.IsVisitorStop(err) {
			return nil
		}
		if err == nil {
			err = v.Duration(b.Offset)
		}
		if err == nil {
			err = v.QueryRange(b.Range)
		}
	}
	return err
}

func (b *builder) ExpressionModifier(f func(Visitor, *ExpressionModifier) error) Builder {
	b.common.expressionModifier = f
	return b
}
