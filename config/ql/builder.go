package ql

import (
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
)

type Builder interface {
	AliasedExpression(func(Visitor, *AliasedExpression) error) Builder
	Expression(func(Visitor, *Expression) error) Builder
	ExpressionModifier(func(Visitor, *ExpressionModifier) error) Builder
	Duration(func(Visitor, *time.Duration) error) Builder
	Function(func(Visitor, *Function) error) Builder
	Histogram(f func(Visitor, *Histogram) error) Builder
	Metric(func(Visitor, *Metric) error) Builder
	Query(func(Visitor, *Query) error) Builder
	QueryRange(func(Visitor, *QueryRange) error) Builder
	Select(func(Visitor, *Select) error) Builder
	SelectExpression(func(Visitor, *SelectExpression) error) Builder
	TableSelect(func(Visitor, *TableSelect) error) Builder
	Time(func(Visitor, *time.Time) error) Builder
	Unit(func(Visitor, *units.Unit) error) Builder
	UsingDefinition(func(Visitor, *UsingDefinition) error) Builder
	UsingDefinitions(func(Visitor, *UsingDefinitions) error) Builder
	WindRose(f func(Visitor, *WindRose) error) Builder

	Build() Visitor
}

func New() Builder {
	return &builder{}
}

type builder struct {
	common
}

func (b *builder) Build() Visitor {
	return &visitor{common: b.common}
}

func (b *builder) Duration(f func(Visitor, *time.Duration) error) Builder {
	b.common.duration = f
	return b
}

func (b *builder) Time(f func(Visitor, *time.Time) error) Builder {
	b.common.time = f
	return b
}

func (b *builder) Unit(f func(Visitor, *units.Unit) error) Builder {
	b.common.unit = f
	return b
}
