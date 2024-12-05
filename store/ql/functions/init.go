package functions

import (
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/ql"
)

func ValidateQuery(d *ql.Query) error {
	return ql.NewBuilder().
		Function(initFunction).
		Build().
		Query(d)
}

func ValidateExpression(d *ql.Expression) error {
	return ql.NewBuilder().
		Function(initFunction).
		Build().
		Expression(d)
}

func initFunction(_ ql.Visitor, d *ql.Function) error {
	if HasFunction(d.Name) {
		return nil
	}
	return errors.Errorf(d.Pos, "unknown function %q", d.Name)
}
