package std

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/num"
)

type Bounded[T num.ExtendedPrimitive] struct {
	value   T
	maximum T
}

func (b Bounded[T]) Value() T {
	return b.value
}

func (b Bounded[T]) Boundary() T {
	return b.maximum
}

func (b Bounded[T]) Set(value T) Bounded[T] {
	if value > b.maximum {
		value = T(uint(value) % uint(b.maximum+1))
	}
	b.value = value
	return b
}

func (b Bounded[T]) SetBoundary(maximum T) Bounded[T] {
	b.maximum = maximum
	return b
}

func (b Bounded[T]) String() string {
	return fmt.Sprintf("%v", b.value)
}
