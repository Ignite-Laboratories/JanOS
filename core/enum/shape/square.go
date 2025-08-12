package shape

import "github.com/ignite-laboratories/core/std/num"

type Square[T num.Primitive] struct {
	Width  T
	Height T
}

func NewSquare[T num.Primitive](width, height T) Square[T] {
	return Square[T]{
		Width:  width,
		Height: height,
	}
}
