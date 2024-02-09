package lang

type Builder[T any] interface {
	Action(func(Visitor[T], *Action) error) Builder[T]
	Amqp(func(Visitor[T], *Amqp) error) Builder[T]
	Format(func(Visitor[T], *Format) error) Builder[T]
	FormatAtom(func(Visitor[T], *FormatAtom) error) Builder[T]
	FormatExpression(func(Visitor[T], *FormatExpression) error) Builder[T]
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

func (b *builder[T]) Amqp(f func(Visitor[T], *Amqp) error) Builder[T] {
	b.amqp = f
	return b
}

func (b *builder[T]) Publish(f func(Visitor[T], *Publish) error) Builder[T] {
	b.publish = f
	return b
}

func (b *builder[T]) Format(f func(Visitor[T], *Format) error) Builder[T] {
	b.format = f
	return b
}

func (b *builder[T]) FormatAtom(f func(Visitor[T], *FormatAtom) error) Builder[T] {
	b.formatAtom = f
	return b
}

func (b *builder[T]) FormatExpression(f func(Visitor[T], *FormatExpression) error) Builder[T] {
	b.formatExpression = f
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
