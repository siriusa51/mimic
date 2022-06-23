package vector

import (
	"fmt"
)

type vector[T any] struct {
	data []T
}

func New[T any](items ...T) *vector[T] {
	v := &vector[T]{}
	for _, item := range items {
		v.PushBack(item)
	}
	return v
}

func (v *vector[T]) String() string {
	return fmt.Sprintf("%v", v.data)
}

func (v *vector[T]) Size() int {
	return len(v.data)
}

func (v *vector[T]) Cap() int {
	return cap(v.data)
}

func (v *vector[T]) Empty() bool {
	return v.Size() == 0
}

func (v *vector[T]) Clear() {
	v.data = v.data[:0]
}

func (v *vector[T]) Resize(size int) {
	if size <= v.Size() {
		return
	}

	v.data = v.data[:size]
}

func (v *vector[T]) Data() []T {
	return v.data
}

func (v *vector[T]) Copy() *vector[T] {
	return New[T](v.data...)
}

func (v *vector[T]) PushBack(value T) {
	v.data = append(v.data, value)
}

func (v *vector[T]) PushFront(value T) {
	nv := make([]T, v.Size()+1)
	copy(nv[1:], v.data)
	nv[0] = value
	v.data = nv
}

func (v *vector[T]) Insert(index int, value T) {
	if index < 0 || index > v.Size() {
		return
	}

	v.data = append(v.data[:index+1], v.data[index:]...)
	v.data[index] = value
}

func (v *vector[T]) RemoveAt(index int) {
	v.RemoveRange(index, index+1)
}

func (v *vector[T]) RemoveRange(start, stop int) {
	if start >= stop || start < 0 || stop > v.Size() {
		return
	}

	size := v.Size() - (stop - start)
	copy(v.data[start:], v.data[stop:])
	v.data = v.data[:size]
}

func (v *vector[T]) PopFront() T {
	val := v.data[0]
	v.data = v.data[1:]
	return val
}

func (v *vector[T]) PopBack() T {
	index := v.Size() - 1
	val := v.data[index]
	v.data = v.data[:index]
	return val
}

func (v *vector[T]) At(index int) T {
	return v.data[index]
}

func (v *vector[T]) Front() T {
	return v.At(0)
}

func (v *vector[T]) Back() T {
	return v.At(v.Size() - 1)
}

func (v *vector[T]) Extend(other *vector[T]) {
	v.data = append(v.data, other.data...)
}

func (v *vector[T]) ForRange() <-chan T {
	c := make(chan T, 0)
	go func() {
		defer close(c)
		for _, val := range v.data {
			c <- val
		}
	}()
	return c
}
