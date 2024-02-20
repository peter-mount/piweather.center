package lang

type Builder[T any] interface {
	At(func(Visitor[T], *At) error) Builder[T]
	Cron(func(Visitor[T], *Cron) error) Builder[T]
	Every(f func(Visitor[T], *Every) error) Builder[T]
	Rule(f func(Visitor[T], *Rule) error) Builder[T]
	Schedule(f func(Visitor[T], *Schedule) error) Builder[T]
	Script(func(Visitor[T], *Script) error) Builder[T]
	TaskRule(f func(Visitor[T], *TaskRule) error) Builder[T]
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

func (b *builder[T]) At(f func(Visitor[T], *At) error) Builder[T] {
	b.at = f
	return b
}

func (b *builder[T]) Cron(f func(Visitor[T], *Cron) error) Builder[T] {
	b.cron = f
	return b
}

func (b *builder[T]) Every(f func(Visitor[T], *Every) error) Builder[T] {
	b.every = f
	return b
}

func (b *builder[T]) Rule(f func(Visitor[T], *Rule) error) Builder[T] {
	b.rule = f
	return b
}

func (b *builder[T]) Schedule(f func(Visitor[T], *Schedule) error) Builder[T] {
	b.schedule = f
	return b
}

func (b *builder[T]) Script(f func(Visitor[T], *Script) error) Builder[T] {
	b.script = f
	return b
}

func (b *builder[T]) TaskRule(f func(Visitor[T], *TaskRule) error) Builder[T] {
	b.taskRule = f
	return b
}
