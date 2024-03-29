package ql

import (
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
)

type QueryVisitor interface {
	time.TimeVisitor
	units.UnitsVisitor
	Query(*Query) error
	Select(*Select) error
	SelectExpression(*SelectExpression) error
	AliasedExpression(*AliasedExpression) error
	Expression(*Expression) error
	ExpressionModifier(*ExpressionModifier) error
	Function(*Function) error
	Metric(*Metric) error
	QueryRange(*QueryRange) error
	UsingDefinitions(*UsingDefinitions) error
	UsingDefinition(*UsingDefinition) error
	Histogram(*Histogram) error
	WindRose(*WindRose) error
}
