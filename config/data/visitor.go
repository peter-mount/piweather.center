package data

type DataVisitor[T any] interface {
	GetData() T
	SetData(T)
}

type DataVisitorCommon[T any] struct {
	data T
}

func (v *DataVisitorCommon[T]) GetData() T {
	return v.data
}

func (v *DataVisitorCommon[T]) SetData(data T) {
	v.data = data
}
