package shape

import "github.com/ignite-laboratories/core/std/num"

type Sphere[T num.Primitive] struct {
	Radius T
}

func NewSphere[T num.Primitive](radius T) Sphere[T] {
	return Sphere[T]{
		Radius: radius,
	}
}
