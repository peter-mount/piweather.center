package location

type LocationVisitor[T any] interface {
	Location(*Location) error
}

type LocationVisitorCommon[T any] struct {
	location func(LocationVisitor[T], *Location) error
}

type LocationBuilder[T any] interface {
	Location(func(LocationVisitor[T], *Location) error)
}

type LocationBuilderBase[T any] struct {
	LocationVisitorCommon[T]
}

func (b *LocationBuilderBase[T]) Location(f func(LocationVisitor[T], *Location) error) {
	b.location = f
}

type LocationVisitorBase[T any] struct {
	LocationVisitorCommon[T]
}

func (v *LocationVisitorBase[T]) Location(b *Location) error {
	var err error
	if b != nil && v.location != nil {
		err = v.location(v, b)
	}
	return err
}
