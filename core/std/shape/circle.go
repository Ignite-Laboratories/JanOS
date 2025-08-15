package shape

import "github.com/ignite-laboratories/core/std/num"

// TODO: A circle should take in a coordinate type

type Circle[T num.Primitive] struct {
	Radius T
}

func NewCircle[T num.Primitive](radius T) Circle[T] {
	return Circle[T]{
		Radius: radius,
	}
}
