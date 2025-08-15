package shape

import "github.com/ignite-laboratories/core/std/num"

type Cube[T num.Primitive] struct {
	Width  T
	Height T
	Depth  T
}

func NewCube[T num.Primitive](width, height, depth T) Cube[T] {
	return Cube[T]{
		Width:  width,
		Height: height,
		Depth:  depth,
	}
}
