package lang

type Builder interface {
	Calculation(func(Visitor, *Calculation) error) Builder
	CronTab(func(Visitor, *CronTab) error) Builder
	Current(func(Visitor, *Current) error) Builder
	Expression(func(Visitor, *Expression) error) Builder
	Function(func(Visitor, *Function) error) Builder
	Location(func(Visitor, *Location) error) Builder
	Metric(func(Visitor, *Metric) error) Builder
	Script(func(Visitor, *Script) error) Builder
	Unit(f func(Visitor, *Unit) error) Builder
	UseFirst(f func(Visitor, *UseFirst) error) Builder
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

func (b *builder) Script(f func(Visitor, *Script) error) Builder {
	b.script = f
	return b
}

func (b *builder) Location(f func(Visitor, *Location) error) Builder {
	b.location = f
	return b
}

func (b *builder) Calculation(f func(Visitor, *Calculation) error) Builder {
	b.calculation = f
	return b
}

func (b *builder) CronTab(f func(Visitor, *CronTab) error) Builder {
	b.cronTab = f
	return b
}

func (b *builder) Expression(f func(Visitor, *Expression) error) Builder {
	b.expression = f
	return b
}

func (b *builder) Unit(f func(Visitor, *Unit) error) Builder {
	b.unit = f
	return b
}

func (b *builder) Current(f func(Visitor, *Current) error) Builder {
	b.current = f
	return b
}

func (b *builder) Function(f func(Visitor, *Function) error) Builder {
	b.function = f
	return b
}

func (b *builder) Metric(f func(Visitor, *Metric) error) Builder {
	b.metric = f
	return b
}

func (b *builder) UseFirst(f func(Visitor, *UseFirst) error) Builder {
	b.useFirst = f
	return b
}
