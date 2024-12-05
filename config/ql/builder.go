package ql

import (
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
)

type Builder[T any] interface {
	AliasedExpression(func(Visitor[T], *AliasedExpression) error) Builder[T]
	Expression(func(Visitor[T], *Expression) error) Builder[T]
	ExpressionModifier(func(Visitor[T], *ExpressionModifier) error) Builder[T]
	Duration(func(Visitor[T], *time.Duration) error) Builder[T]
	Function(func(Visitor[T], *Function) error) Builder[T]
	Histogram(func(Visitor[T], *Histogram) error) Builder[T]
	Metric(func(Visitor[T], *Metric) error) Builder[T]
	Query(func(Visitor[T], *Query) error) Builder[T]
	QueryRange(func(Visitor[T], *QueryRange) error) Builder[T]
	Select(func(Visitor[T], *Select) error) Builder[T]
	SelectExpression(func(Visitor[T], *SelectExpression) error) Builder[T]
	Summarize(func(Visitor[T], *Summarize) error) Builder[T]
	TableSelect(func(Visitor[T], *TableSelect) error) Builder[T]
	Time(func(Visitor[T], *time.Time) error) Builder[T]
	Unit(func(Visitor[T], *units.Unit) error) Builder[T]
	UsingDefinition(func(Visitor[T], *UsingDefinition) error) Builder[T]
	UsingDefinitions(func(Visitor[T], *UsingDefinitions) error) Builder[T]
	WindRose(f func(Visitor[T], *WindRose) error) Builder[T]

	Build() Visitor[T]
}

func NewBuilder[T any]() Builder[T] {
	return &builder[T]{}
}

type builder[T any] struct {
	common[T]
}

func (b *builder[T]) Build() Visitor[T] {
	return &visitor[T]{common: b.common}
}

func (b *builder[T]) Duration(f func(Visitor[T], *time.Duration) error) Builder[T] {
	b.common.duration = f
	return b
}

func (b *builder[T]) Time(f func(Visitor[T], *time.Time) error) Builder[T] {
	b.common.time = f
	return b
}

func (b *builder[T]) Unit(f func(Visitor[T], *units.Unit) error) Builder[T] {
	b.common.unit = f
	return b
}
