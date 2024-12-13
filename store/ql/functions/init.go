package functions

import (
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/ql"
)

var (
	validateVisitor = ql.NewBuilder[any]().
		Function(func(_ ql.Visitor[any], d *ql.Function) error {
			if HasFunction(d.Name) {
				return nil
			}
			return errors.Errorf(d.Pos, "unknown function %q", d.Name)
		}).
		Build()
)

func ValidateQuery(d *ql.Query) error {
	return validateVisitor.Clone().Query(d)
}

func ValidateExpression(d *ql.Expression) error {
	return validateVisitor.Clone().Expression(d)
}
