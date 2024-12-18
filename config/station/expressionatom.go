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

		if err == nil && b.Current != nil {
			err = c.Current(b.Current)
		}

		if err == nil && b.Function != nil {
			err = c.Function(b.Function)
		}

		if err == nil && b.Metric != nil {
			err = c.MetricExpression(b.Metric)
		}

		if err == nil && b.Location != nil {
			err = c.LocationExpression(b.Location)
		}

		if err == nil && b.Using != nil {
			err = c.Unit(b.Using)
		}

		err = errors.Error(b.Pos, err)
	}
	return err
}

func (b *builder[T]) ExpressionAtom(f func(Visitor[T], *ExpressionAtom) error) Builder[T] {
	b.expressionAtom = f
	return b
}
