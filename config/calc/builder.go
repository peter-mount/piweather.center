package calc

import (
	"github.com/peter-mount/piweather.center/config/util/location"
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
)

type Builder[T any] interface {
	location.LocationBuilder[T]
	Calculation(func(CalcVisitor[T], *Calculation) error) Builder[T]
	CronTab(func(CalcVisitor[T], *time.CronTab) error) Builder[T]
	Current(func(CalcVisitor[T], *Current) error) Builder[T]
	Expression(func(CalcVisitor[T], *Expression) error) Builder[T]
	Function(func(CalcVisitor[T], *Function) error) Builder[T]
	Load(f func(CalcVisitor[T], *Load) error) Builder[T]
	Metric(func(CalcVisitor[T], *Metric) error) Builder[T]
	Script(func(CalcVisitor[T], *Script) error) Builder[T]
	Unit(f func(CalcVisitor[T], *units.Unit) error) Builder[T]
	UseFirst(f func(CalcVisitor[T], *UseFirst) error) Builder[T]
	Build() CalcVisitor[T]
}

type builder[T any] struct {
	location.LocationBuilderBase[T]
	visitorCommon[T]
}

func NewBuilder[T any]() Builder[T] {
	return &builder[T]{}
}

func (b *builder[T]) Build() CalcVisitor[T] {
	return &visitor[T]{
		visitorCommon: b.visitorCommon,
		LocationVisitorBase: location.LocationVisitorBase[T]{
			LocationVisitorCommon: b.LocationBuilderBase.LocationVisitorCommon,
		},
	}
}

func (b *builder[T]) Script(f func(CalcVisitor[T], *Script) error) Builder[T] {
	b.script = f
	return b
}

func (b *builder[T]) Calculation(f func(CalcVisitor[T], *Calculation) error) Builder[T] {
	b.calculation = f
	return b
}

func (b *builder[T]) Load(f func(CalcVisitor[T], *Load) error) Builder[T] {
	b.load = f
	return b
}

func (b *builder[T]) CronTab(f func(CalcVisitor[T], *time.CronTab) error) Builder[T] {
	b.cronTab = f
	return b
}

func (b *builder[T]) Expression(f func(CalcVisitor[T], *Expression) error) Builder[T] {
	b.expression = f
	return b
}

func (b *builder[T]) Unit(f func(CalcVisitor[T], *units.Unit) error) Builder[T] {
	b.unit = f
	return b
}

func (b *builder[T]) Current(f func(CalcVisitor[T], *Current) error) Builder[T] {
	b.current = f
	return b
}

func (b *builder[T]) Function(f func(CalcVisitor[T], *Function) error) Builder[T] {
	b.function = f
	return b
}

func (b *builder[T]) Metric(f func(CalcVisitor[T], *Metric) error) Builder[T] {
	b.metric = f
	return b
}

func (b *builder[T]) UseFirst(f func(CalcVisitor[T], *UseFirst) error) Builder[T] {
	b.useFirst = f
	return b
}
