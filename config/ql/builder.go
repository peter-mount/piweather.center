package ql

import (
	"github.com/peter-mount/piweather.center/config/util/ql"
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
)

type Builder interface {
	Query(func(ql.QueryVisitor, *ql.Query) error) Builder

	Select(func(ql.QueryVisitor, *ql.Select) error) Builder
	SelectExpression(func(ql.QueryVisitor, *ql.SelectExpression) error) Builder
	AliasedExpression(func(ql.QueryVisitor, *ql.AliasedExpression) error) Builder
	Expression(func(ql.QueryVisitor, *ql.Expression) error) Builder
	ExpressionModifier(func(ql.QueryVisitor, *ql.ExpressionModifier) error) Builder
	Function(func(ql.QueryVisitor, *ql.Function) error) Builder
	Metric(func(ql.QueryVisitor, *ql.Metric) error) Builder
	QueryRange(func(ql.QueryVisitor, *ql.QueryRange) error) Builder
	Time(func(ql.QueryVisitor, *time.Time) error) Builder
	Duration(func(ql.QueryVisitor, *time.Duration) error) Builder
	Unit(func(ql.QueryVisitor, *units.Unit) error) Builder
	UsingDefinitions(func(ql.QueryVisitor, *ql.UsingDefinitions) error) Builder
	UsingDefinition(func(ql.QueryVisitor, *ql.UsingDefinition) error) Builder

	Histogram(f func(ql.QueryVisitor, *ql.Histogram) error) Builder
	WindRose(f func(ql.QueryVisitor, *ql.WindRose) error) Builder

	Build() ql.QueryVisitor
}

func NewBuilder() Builder {
	return &builder{}
}

type builder struct {
	common
}

func (b *builder) Build() ql.QueryVisitor {
	v := &visitor{}
	v.common = b.common
	return v
}

func (b *builder) Query(f func(ql.QueryVisitor, *ql.Query) error) Builder {
	b.common.query = f
	return b
}

func (b *builder) Select(f func(ql.QueryVisitor, *ql.Select) error) Builder {
	b.common._select = f
	return b
}

func (b *builder) SelectExpression(f func(ql.QueryVisitor, *ql.SelectExpression) error) Builder {
	b.common.selectExpression = f
	return b
}

func (b *builder) AliasedExpression(f func(ql.QueryVisitor, *ql.AliasedExpression) error) Builder {
	b.common.aliasedExpression = f
	return b
}

func (b *builder) Expression(f func(ql.QueryVisitor, *ql.Expression) error) Builder {
	b.common.expression = f
	return b
}

func (b *builder) ExpressionModifier(f func(ql.QueryVisitor, *ql.ExpressionModifier) error) Builder {
	b.common.expressionModifier = f
	return b
}

func (b *builder) Function(f func(ql.QueryVisitor, *ql.Function) error) Builder {
	b.common.function = f
	return b
}

func (b *builder) Metric(f func(ql.QueryVisitor, *ql.Metric) error) Builder {
	b.common.metric = f
	return b
}

func (b *builder) QueryRange(f func(ql.QueryVisitor, *ql.QueryRange) error) Builder {
	b.common.queryRange = f
	return b
}

func (b *builder) Time(f func(ql.QueryVisitor, *time.Time) error) Builder {
	b.common.time = f
	return b
}

func (b *builder) Duration(f func(ql.QueryVisitor, *time.Duration) error) Builder {
	b.common.duration = f
	return b
}

func (b *builder) Unit(f func(ql.QueryVisitor, *units.Unit) error) Builder {
	b.common.unit = f
	return b
}

func (b *builder) UsingDefinitions(f func(ql.QueryVisitor, *ql.UsingDefinitions) error) Builder {
	b.common.usingDefinitions = f
	return b
}

func (b *builder) UsingDefinition(f func(ql.QueryVisitor, *ql.UsingDefinition) error) Builder {
	b.common.usingDefinition = f
	return b
}

func (b *builder) Histogram(f func(ql.QueryVisitor, *ql.Histogram) error) Builder {
	b.common.histogram = f
	return b
}

func (b *builder) WindRose(f func(ql.QueryVisitor, *ql.WindRose) error) Builder {
	b.common.windRose = f
	return b
}
