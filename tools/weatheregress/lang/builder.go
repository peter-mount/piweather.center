package lang

type Builder interface {
	Action(func(Visitor, *Action) error) Builder
	Amqp(func(Visitor, *Amqp) error) Builder
	Format(func(Visitor, *Format) error) Builder
	FormatAtom(func(Visitor, *FormatAtom) error) Builder
	FormatExpression(func(Visitor, *FormatExpression) error) Builder
	Metric(f func(Visitor, *Metric) error) Builder
	Publish(f func(Visitor, *Publish) error) Builder
	Script(func(Visitor, *Script) error) Builder
	Build() Visitor
}

type builder struct {
	visitorCommon
}

func NewBuilder() Builder {
	return &builder{}
}

func (b *builder) Build() Visitor {
	return &visitor{visitorCommon: b.visitorCommon}
}

func (b *builder) Action(f func(Visitor, *Action) error) Builder {
	b.action = f
	return b
}

func (b *builder) Amqp(f func(Visitor, *Amqp) error) Builder {
	b.amqp = f
	return b
}

func (b *builder) Publish(f func(Visitor, *Publish) error) Builder {
	b.publish = f
	return b
}

func (b *builder) Format(f func(Visitor, *Format) error) Builder {
	b.format = f
	return b
}

func (b *builder) FormatAtom(f func(Visitor, *FormatAtom) error) Builder {
	b.formatAtom = f
	return b
}

func (b *builder) FormatExpression(f func(Visitor, *FormatExpression) error) Builder {
	b.formatExpression = f
	return b
}

func (b *builder) Metric(f func(Visitor, *Metric) error) Builder {
	b.metric = f
	return b
}

func (b *builder) Script(f func(Visitor, *Script) error) Builder {
	b.script = f
	return b
}
