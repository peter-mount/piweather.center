package egress

import "github.com/peter-mount/piweather.center/config/util/amqp"

type Builder[T any] interface {
	Action(func(EgressVisitor[T], *Action) error) Builder[T]
	Amqp(func(EgressVisitor[T], *amqp.Amqp) error) Builder[T]
	Metric(f func(EgressVisitor[T], *Metric) error) Builder[T]
	Publish(f func(EgressVisitor[T], *Publish) error) Builder[T]
	Script(func(EgressVisitor[T], *Script) error) Builder[T]
	Build() EgressVisitor[T]
}

type builder[T any] struct {
	visitorCommon[T]
}

func NewBuilder[T any]() Builder[T] {
	return &builder[T]{}
}

func (b *builder[T]) Build() EgressVisitor[T] {
	return &visitor[T]{visitorCommon: b.visitorCommon}
}

func (b *builder[T]) Action(f func(EgressVisitor[T], *Action) error) Builder[T] {
	b.action = f
	return b
}

func (b *builder[T]) Amqp(f func(EgressVisitor[T], *amqp.Amqp) error) Builder[T] {
	b.amqp = f
	return b
}

func (b *builder[T]) Publish(f func(EgressVisitor[T], *Publish) error) Builder[T] {
	b.publish = f
	return b
}

func (b *builder[T]) Metric(f func(EgressVisitor[T], *Metric) error) Builder[T] {
	b.metric = f
	return b
}

func (b *builder[T]) Script(f func(EgressVisitor[T], *Script) error) Builder[T] {
	b.script = f
	return b
}
