package shape

import "github.com/ignite-laboratories/core/std/num"

type Ellipse[T num.Primitive] struct {
	XAxis T
	YAxis T
}

func NewEllipse[T num.Primitive](xAxis, yAxis T) Ellipse[T] {
	return Ellipse[T]{
		XAxis: xAxis,
		YAxis: yAxis,
	}
}
