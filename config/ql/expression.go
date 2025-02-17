package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
)

// Expression handles function calls or direct metric values
type Expression struct {
	Pos      lexer.Position
	Function *Function             `parser:"( @@"`
	Metric   *Metric               `parser:"| @@ )"`
	Using    string                `parser:"( 'using' @String"`
	Modifier []*ExpressionModifier `parser:"| (@@)+ )?"`
}

func (v *visitor[T]) Expression(b *Expression) error {
	var err error
	if b != nil {
		if v.expression != nil {
			err = v.expression(v, b)
		}

		if errors.IsVisitorStop(err) {
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

func initExpression(v Visitor[*parserState], s *Expression) error {
	p := v.Get()

	if s.Using != "" && !p.usingNames.Contains(s.Using) {
		return errors.Errorf(s.Pos, "%q undefined", s.Using)
	}
	return nil
}

func (b *builder[T]) Expression(f func(Visitor[T], *Expression) error) Builder[T] {
	b.common.expression = f
	return b
}
