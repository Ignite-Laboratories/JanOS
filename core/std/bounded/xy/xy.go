package xy

import (
	"github.com/ignite-laboratories/core/std/bounded"
	"github.com/ignite-laboratories/core/std/num"
)

// From creates a new instance of std.XY[T] with each direction bounded in the fully closed interval [0, max].  If you
// would like to also set the minimum boundaries, please use FromFull.
func From[T num.Primitive](x, y T, maxX, maxY T) bounded.XY[T] {
	return bounded.XY[T]{}.SetAll(x, y, 0, maxX, 0, maxY)
}

// FromFull creates a new instance of std.XY[T] with each direction bounded in the fully closed interval [min, max].
func FromFull[T num.Primitive](x, y T, minX, maxX, minY, maxY T) bounded.XY[T] {
	return bounded.XY[T]{}.SetAll(x, y, minX, maxX, minY, maxY)
}

// Random returns a pseudo-random std.XY[T] of the provided type using math.Random[T], with
// each directional component bounded in the fully closed interval [0, max].  If you would like
// the minimum to be above 0, please use RandomFull
func Random[T num.Primitive](maxX, maxY T) bounded.XY[T] {
	x := num.RandomWithinRange[T](0, maxX)
	y := num.RandomWithinRange[T](0, maxY)
	return bounded.XY[T]{}.SetAll(x, y, 0, maxX, 0, maxY)
}

// RandomFull returns a pseudo-random std.XY[T] of the provided type using math.Random[T], with
// each directional component bounded in the fully closed interval [min, max].
func RandomFull[T num.Primitive](minX, maxX, minY, maxY T) bounded.XY[T] {
	x := num.RandomWithinRange[T](minX, maxX)
	y := num.RandomWithinRange[T](minY, maxY)
	return bounded.XY[T]{}.SetAll(x, y, minX, maxX, minY, maxY)
}

// ScaleToType normalizes the std.XY[T] directional components into unit vectors and then scales them to a new std.XY[TOut].
func ScaleToType[TIn num.Primitive, TOut num.Primitive](value bounded.XY[TIn]) bounded.XY[TOut] {
	return bounded.XY[TIn]{
		X: bounded.ScaleToType[TOut](value.X),
		Y: bounded.ScaleToType[TOut](value.Y),
	}
}
