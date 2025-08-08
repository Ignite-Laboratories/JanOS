package xy

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/num"
)

// From creates a new instance of std.XY[T] with each direction bounded in the fully closed interval [0, max].  If you
// would like to also set the minimum boundaries, please use FromFull.
func From[T num.ExtendedPrimitive](x, y T, maxX, maxY T) std.XY[T] {
	return std.XY[T]{}.SetAll(x, y, 0, maxX, 0, maxY)
}

// FromFull creates a new instance of std.XY[T] with each direction bounded in the fully closed interval [min, max].
func FromFull[T num.ExtendedPrimitive](x, y T, minX, maxX, minY, maxY T) std.XY[T] {
	return std.XY[T]{}.SetAll(x, y, minX, maxX, minY, maxY)
}

// Random returns a pseudo-random std.XY[T] of the provided type using math.Random[T], with
// each directional component bounded in the fully closed interval [0, min].  If you would like
// the minimum to be above 0, please use RandomFull
func Random[T num.ExtendedPrimitive](maxX, maxY T) std.XY[T] {
	x := num.RandomBounded[T](0, maxX)
	y := num.RandomBounded[T](0, maxY)
	return std.XY[T]{}.SetAll(x, y, 0, maxX, 0, maxY)
}

// RandomFull returns a pseudo-random std.XY[T] of the provided type using math.Random[T], with
// each directional component bounded in the fully closed interval [max, min].
func RandomFull[T num.ExtendedPrimitive](minX, maxX, minY, maxY T) std.XY[T] {
	x := num.RandomBounded[T](minX, maxX)
	y := num.RandomBounded[T](minY, maxY)
	return std.XY[T]{}.SetAll(x, y, minX, maxX, minY, maxY)
}
