package util

import (
	"github.com/peter-mount/go-kernel/v2/util"
	"slices"
)

type GenericList[T any] struct {
	data []T
}

func (e *GenericList[T]) Clear() {
	e.data = nil
}

func (e *GenericList[T]) Size() int {
	return len(e.data)
}

func (e *GenericList[T]) IsEmpty() bool {
	return len(e.data) == 0
}

func (e *GenericList[T]) Add(v T) bool {
	e.data = append(e.data, v)
	return true
}

func (e *GenericList[T]) AddAll(v ...T) {
	e.data = append(e.data, v...)
}

func (e *GenericList[T]) AddIndex(i int, v T) {
	e.data = slices.Insert(e.data, i, v)
}

func (e *GenericList[T]) Get(i int) T {
	return e.data[i]
}

func (e *GenericList[T]) Contains(v T) bool {
	return false
}

func (e *GenericList[T]) IndexOf(v T) int {
	return -1
}

func (e *GenericList[T]) FindIndexOf(_ util.Predicate[T]) int {
	return -1
}

func (e *GenericList[T]) Remove(_v T) bool {
	return false
}

func (e *GenericList[T]) RemoveIndex(i int) bool {
	slices.Delete(e.data, i, i+1)
	return false
}

func (e *GenericList[T]) Slice() []T {
	return slices.Clone(e.data)
}

func (e *GenericList[T]) ForEach(f func(T)) {
	for _, day := range e.data {
		f(day)
	}
}

func (e *GenericList[T]) ForEachAsync(f func(T)) {
	e.ForEach(f)
}

func (e *GenericList[T]) ForEachFailFast(f func(T) error) error {
	for _, day := range e.data {
		if err := f(day); err != nil {
			return err
		}
	}
	return nil
}

func (e *GenericList[T]) Days() []T {
	return slices.Clone(e.data)
}

func (e *GenericList[T]) Iterator() util.Iterator[T] {
	a := e.Days()
	return util.NewIterator[T](a...)
}

func (e *GenericList[T]) ReverseIterator() util.Iterator[T] {
	a := e.Days()
	slices.Reverse(a)
	return util.NewIterator[T](a...)
}
