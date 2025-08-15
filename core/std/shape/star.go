package shape

import "github.com/ignite-laboratories/core/std/num"

type Star[T num.Primitive] struct {
	Points      int
	InnerRadius int
	OuterRadius int
}

func NewStar[T num.Primitive](points, innerRadius, outerRadius int) Star[T] {
	return Star[T]{
		Points:      points,
		InnerRadius: innerRadius,
		OuterRadius: outerRadius,
	}
}
