package calc

import (
	"github.com/peter-mount/piweather.center/config/location"
	"github.com/peter-mount/piweather.center/config/misc"
)

type Builder[T any] interface {
	location.LocationBuilder[T]
	Calculation(func(Visitor[T], *Calculation) error) Builder[T]
	CronTab(func(Visitor[T], *misc.CronTab) error) Builder[T]
	Current(func(Visitor[T], *Current) error) Builder[T]
	Expression(func(Visitor[T], *Expression) error) Builder[T]
	Function(func(Visitor[T], *Function) error) Builder[T]
	Load(f func(Visitor[T], *Load) error) Builder[T]
	Metric(func(Visitor[T], *Metric) error) Builder[T]
	Script(func(Visitor[T], *Script) error) Builder[T]
	Unit(f func(Visitor[T], *misc.Unit) error) Builder[T]
	UseFirst(f func(Visitor[T], *UseFirst) error) Builder[T]
	Build() Visitor[T]
}

type builder[T any] struct {
	location.LocationBuilderBase[T]
	visitorCommon[T]
}

func NewBuilder[T any]() Builder[T] {
	return &builder[T]{}
}

func (b *builder[T]) Build() Visitor[T] {
	return &visitor[T]{
		visitorCommon: b.visitorCommon,
		LocationVisitorBase: location.LocationVisitorBase[T]{
			LocationVisitorCommon: b.LocationBuilderBase.LocationVisitorCommon,
		},
	}
}

func (b *builder[T]) Script(f func(Visitor[T], *Script) error) Builder[T] {
	b.script = f
	return b
}

func (b *builder[T]) Calculation(f func(Visitor[T], *Calculation) error) Builder[T] {
	b.calculation = f
	return b
}

func (b *builder[T]) Load(f func(Visitor[T], *Load) error) Builder[T] {
	b.load = f
	return b
}

func (b *builder[T]) CronTab(f func(Visitor[T], *misc.CronTab) error) Builder[T] {
	b.cronTab = f
	return b
}

func (b *builder[T]) Expression(f func(Visitor[T], *Expression) error) Builder[T] {
	b.expression = f
	return b
}

func (b *builder[T]) Unit(f func(Visitor[T], *misc.Unit) error) Builder[T] {
	b.unit = f
	return b
}

func (b *builder[T]) Current(f func(Visitor[T], *Current) error) Builder[T] {
	b.current = f
	return b
}

func (b *builder[T]) Function(f func(Visitor[T], *Function) error) Builder[T] {
	b.function = f
	return b
}

func (b *builder[T]) Metric(f func(Visitor[T], *Metric) error) Builder[T] {
	b.metric = f
	return b
}

func (b *builder[T]) UseFirst(f func(Visitor[T], *UseFirst) error) Builder[T] {
	b.useFirst = f
	return b
}
