package egress

import "github.com/peter-mount/piweather.center/config/util/amqp"

type Builder[T any] interface {
	Action(func(Visitor[T], *Action) error) Builder[T]
	Amqp(func(Visitor[T], *amqp.Amqp) error) Builder[T]
	Metric(f func(Visitor[T], *Metric) error) Builder[T]
	Publish(f func(Visitor[T], *Publish) error) Builder[T]
	Script(func(Visitor[T], *Script) error) Builder[T]
	Build() Visitor[T]
}

type builder[T any] struct {
	visitorCommon[T]
}

func NewBuilder[T any]() Builder[T] {
	return &builder[T]{}
}

func (b *builder[T]) Build() Visitor[T] {
	return &visitor[T]{visitorCommon: b.visitorCommon}
}

func (b *builder[T]) Action(f func(Visitor[T], *Action) error) Builder[T] {
	b.action = f
	return b
}

func (b *builder[T]) Amqp(f func(Visitor[T], *amqp.Amqp) error) Builder[T] {
	b.amqp = f
	return b
}

func (b *builder[T]) Publish(f func(Visitor[T], *Publish) error) Builder[T] {
	b.publish = f
	return b
}

func (b *builder[T]) Metric(f func(Visitor[T], *Metric) error) Builder[T] {
	b.metric = f
	return b
}

func (b *builder[T]) Script(f func(Visitor[T], *Script) error) Builder[T] {
	b.script = f
	return b
}
