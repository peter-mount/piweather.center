package ql

type Builder interface {
	Query(func(Visitor, *Query) error) Builder

	Select(func(Visitor, *Select) error) Builder
	SelectExpression(func(Visitor, *SelectExpression) error) Builder
	AliasedExpression(func(Visitor, *AliasedExpression) error) Builder
	Expression(func(Visitor, *Expression) error) Builder
	ExpressionModifier(func(Visitor, *ExpressionModifier) error) Builder
	Function(func(Visitor, *Function) error) Builder
	Metric(func(Visitor, *Metric) error) Builder
	QueryRange(func(Visitor, *QueryRange) error) Builder
	Time(func(Visitor, *Time) error) Builder
	Duration(func(Visitor, *Duration) error) Builder
	UsingDefinitions(func(Visitor, *UsingDefinitions) error) Builder
	UsingDefinition(func(Visitor, *UsingDefinition) error) Builder

	Histogram(f func(Visitor, *Histogram) error) Builder
	WindRose(f func(Visitor, *WindRose) error) Builder

	Build() Visitor
}

func NewBuilder() Builder {
	return &builder{}
}

type builder struct {
	common
}

func (b *builder) Build() Visitor {
	v := &visitor{}
	v.common = b.common
	return v
}

func (b *builder) Query(f func(Visitor, *Query) error) Builder {
	b.common.query = f
	return b
}

func (b *builder) Select(f func(Visitor, *Select) error) Builder {
	b.common._select = f
	return b
}

func (b *builder) SelectExpression(f func(Visitor, *SelectExpression) error) Builder {
	b.common.selectExpression = f
	return b
}

func (b *builder) AliasedExpression(f func(Visitor, *AliasedExpression) error) Builder {
	b.common.aliasedExpression = f
	return b
}

func (b *builder) Expression(f func(Visitor, *Expression) error) Builder {
	b.common.expression = f
	return b
}

func (b *builder) ExpressionModifier(f func(Visitor, *ExpressionModifier) error) Builder {
	b.common.expressionModifier = f
	return b
}

func (b *builder) Function(f func(Visitor, *Function) error) Builder {
	b.common.function = f
	return b
}

func (b *builder) Metric(f func(Visitor, *Metric) error) Builder {
	b.common.metric = f
	return b
}

func (b *builder) QueryRange(f func(Visitor, *QueryRange) error) Builder {
	b.common.queryRange = f
	return b
}

func (b *builder) Time(f func(Visitor, *Time) error) Builder {
	b.common.time = f
	return b
}

func (b *builder) Duration(f func(Visitor, *Duration) error) Builder {
	b.common.duration = f
	return b
}

func (b *builder) UsingDefinitions(f func(Visitor, *UsingDefinitions) error) Builder {
	b.common.usingDefinitions = f
	return b
}

func (b *builder) UsingDefinition(f func(Visitor, *UsingDefinition) error) Builder {
	b.common.usingDefinition = f
	return b
}

func (b *builder) Histogram(f func(Visitor, *Histogram) error) Builder {
	b.common.histogram = f
	return b
}

func (b *builder) WindRose(f func(Visitor, *WindRose) error) Builder {
	b.common.windRose = f
	return b
}
