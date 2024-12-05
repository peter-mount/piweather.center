package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util"
)

// Expression handles function calls or direct metric values
type Expression struct {
	Pos      lexer.Position
	Function *Function             `parser:"( @@"`
	Metric   *Metric               `parser:"| @@ )"`
	Using    string                `parser:"( 'using' @String"`
	Modifier []*ExpressionModifier `parser:"| (@@)+ )?"`
}

func (v *visitor) Expression(b *Expression) error {
	var err error
	if b != nil {
		if v.expression != nil {
			err = v.expression(v, b)
		}

		if util.IsVisitorStop(err) {
			return nil
		}

		for _, e := range b.Modifier {
			if err == nil {
				err = v.ExpressionModifier(e)
			}
		}

		if err == nil {
			err = v.Function(b.Function)
		}

		if err == nil {
			err = v.Metric(b.Metric)
		}
	}

	return err
}

func (b *builder) Expression(f func(Visitor, *Expression) error) Builder {
	b.common.expression = f
	return b
}
