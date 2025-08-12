package shape

import "github.com/ignite-laboratories/core/std/num"

type Ellipsoid[T num.Primitive] struct {
	XAxis T
	YAxis T
	ZAxis T
}

func NewEllipsoid[T num.Primitive](xAxis, yAxis, zAxis T) Ellipsoid[T] {
	return Ellipsoid[T]{
		XAxis: xAxis,
		YAxis: yAxis,
		ZAxis: zAxis,
	}
}
