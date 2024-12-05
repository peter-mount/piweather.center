package ql

import (
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
)

type Builder interface {
	Query(func(QueryVisitor, *Query) error) Builder

	Select(func(QueryVisitor, *Select) error) Builder
	SelectExpression(func(QueryVisitor, *SelectExpression) error) Builder
	AliasedExpression(func(QueryVisitor, *AliasedExpression) error) Builder
	Expression(func(QueryVisitor, *Expression) error) Builder
	ExpressionModifier(func(QueryVisitor, *ExpressionModifier) error) Builder
	Function(func(QueryVisitor, *Function) error) Builder
	Metric(func(QueryVisitor, *Metric) error) Builder
	QueryRange(func(QueryVisitor, *QueryRange) error) Builder
	Time(func(QueryVisitor, *time.Time) error) Builder
	Duration(func(QueryVisitor, *time.Duration) error) Builder
	Unit(func(QueryVisitor, *units.Unit) error) Builder
	UsingDefinitions(func(QueryVisitor, *UsingDefinitions) error) Builder
	UsingDefinition(func(QueryVisitor, *UsingDefinition) error) Builder

	Histogram(f func(QueryVisitor, *Histogram) error) Builder
	WindRose(f func(QueryVisitor, *WindRose) error) Builder

	TableSelect(func(QueryVisitor, *TableSelect) error) Builder

	Build() QueryVisitor
}

func NewBuilder() Builder {
	return &builder{}
}

type builder struct {
	common
}

func (b *builder) Build() QueryVisitor {
	v := &visitor{}
	v.common = b.common
	return v
}

func (b *builder) Query(f func(QueryVisitor, *Query) error) Builder {
	b.common.query = f
	return b
}

func (b *builder) Select(f func(QueryVisitor, *Select) error) Builder {
	b.common._select = f
	return b
}

func (b *builder) SelectExpression(f func(QueryVisitor, *SelectExpression) error) Builder {
	b.common.selectExpression = f
	return b
}

func (b *builder) AliasedExpression(f func(QueryVisitor, *AliasedExpression) error) Builder {
	b.common.aliasedExpression = f
	return b
}

func (b *builder) Expression(f func(QueryVisitor, *Expression) error) Builder {
	b.common.expression = f
	return b
}

func (b *builder) ExpressionModifier(f func(QueryVisitor, *ExpressionModifier) error) Builder {
	b.common.expressionModifier = f
	return b
}

func (b *builder) Function(f func(QueryVisitor, *Function) error) Builder {
	b.common.function = f
	return b
}

func (b *builder) Metric(f func(QueryVisitor, *Metric) error) Builder {
	b.common.metric = f
	return b
}

func (b *builder) QueryRange(f func(QueryVisitor, *QueryRange) error) Builder {
	b.common.queryRange = f
	return b
}

func (b *builder) Time(f func(QueryVisitor, *time.Time) error) Builder {
	b.common.time = f
	return b
}

func (b *builder) Duration(f func(QueryVisitor, *time.Duration) error) Builder {
	b.common.duration = f
	return b
}

func (b *builder) Unit(f func(QueryVisitor, *units.Unit) error) Builder {
	b.common.unit = f
	return b
}

func (b *builder) UsingDefinitions(f func(QueryVisitor, *UsingDefinitions) error) Builder {
	b.common.usingDefinitions = f
	return b
}

func (b *builder) UsingDefinition(f func(QueryVisitor, *UsingDefinition) error) Builder {
	b.common.usingDefinition = f
	return b
}

func (b *builder) Histogram(f func(QueryVisitor, *Histogram) error) Builder {
	b.common.histogram = f
	return b
}

func (b *builder) WindRose(f func(QueryVisitor, *WindRose) error) Builder {
	b.common.windRose = f
	return b
}

func (b *builder) TableSelect(f func(QueryVisitor, *TableSelect) error) Builder {
	b.common.tableSelect = f
	return b
}
