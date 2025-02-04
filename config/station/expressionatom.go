package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util/units"
)

type ExpressionAtom struct {
	Pos      lexer.Position
	Current  *Current            `parser:"( @@"`   // Get the current value of calculation
	Function *Function           `parser:"| @@"`   // Generic Function Call
	Location *LocationExpression `parser:"| @@"`   // Return values from the stations location
	Metric   *MetricExpression   `parser:"| @@ )"` // Metric reference
	Using    *units.Unit         `parser:"(@@)?"`  // Optional target Unit
}

func (c *visitor[T]) ExpressionAtom(b *ExpressionAtom) error {
	var err error
	if b != nil {
		if c.expressionAtom != nil {
			err = c.expressionAtom(c, b)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = visitExpressionAtom[T](c, b)
		}

		err = errors.Error(b.Pos, err)
	}
	return err
}

func visitExpressionAtom[T any](v Visitor[T], d *ExpressionAtom) error {
	var err error
	if d != nil {
		switch {
		case d.Current != nil:
			err = v.Current(d.Current)

		case d.Function != nil:
			err = v.Function(d.Function)

		case d.Location != nil:
			err = v.LocationExpression(d.Location)

		case d.Metric != nil:
			err = v.MetricExpression(d.Metric)
		}

		if err == nil && d.Using != nil {
			err = v.Unit(d.Using)
		}
	}
	return err
}

func (b *builder[T]) ExpressionAtom(f func(Visitor[T], *ExpressionAtom) error) Builder[T] {
	b.expressionAtom = f
	return b
}

func printExpressionAtom(v Visitor[*printState], d *ExpressionAtom) error {
	err := visitExpressionAtom(v, d)
	if err == nil {
		err = errors.VisitorStop
	}
	return err
}
