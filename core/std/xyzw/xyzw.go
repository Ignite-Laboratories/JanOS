package xyzw

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/bounded"
	"github.com/ignite-laboratories/core/std/num"
)

// From creates a new instance of std.XYZW[T] with each direction bounded in the fully closed interval [0, max].  If you
// would like to also set the minimum boundaries, please use FromFull.
func From[T num.Primitive](x, y, z T, w float64, maxX, maxY, maxZ T) std.XYZW[T] {
	return std.XYZW[T]{}.Set(x, y, z, w, 0, maxX, 0, maxY, 0, maxZ)
}

// FromBounded creates a new instance of std.XYZW[T] with each direction bounded in the fully closed interval [min, max].
func FromBounded[T num.Primitive](x, y, z T, w float64, minX, maxX, minY, maxY, minZ, maxZ T) std.XYZW[T] {
	return std.XYZW[T]{}.Set(x, y, z, w, minX, maxX, minY, maxY, minZ, maxZ)
}

// Random returns a pseudo-random std.XYZW[T] of the provided type using math.Random[T], with
// each directional component bounded in the fully closed interval [0, max].  If you would like
// the minimum to be above 0, please use RandomFull
//
// NOTE: W will always return as 1.0
func Random[T num.Primitive](maxX, maxY, maxZ T) std.XYZW[T] {
	x := num.RandomWithinRange[T](0, maxX)
	y := num.RandomWithinRange[T](0, maxY)
	z := num.RandomWithinRange[T](0, maxZ)
	return std.XYZW[T]{}.Set(x, y, z, 1.0, 0, maxX, 0, maxY, 0, maxZ)
}

// RandomFull returns a pseudo-random std.XYZW[T] of the provided type using math.Random[T], with
// each directional component bounded in the fully closed interval [min, max].
//
// NOTE: W will always return as 1.0
func RandomFull[T num.Primitive](minX, maxX, minY, maxY, minZ, maxZ T) std.XYZW[T] {
	x := num.RandomWithinRange[T](minX, maxX)
	y := num.RandomWithinRange[T](minY, maxY)
	z := num.RandomWithinRange[T](minZ, maxZ)
	return std.XYZW[T]{}.Set(x, y, z, 1.0, minX, maxX, minY, maxY, minZ, maxZ)
}

// ScaleToType normalizes the std.XYZW[T] directional components into unit vectors and then scales them to a new std.XYZW[TOut].
//
// NOTE: W will remain unchanged.
func ScaleToType[TIn num.Primitive, TOut num.Primitive](value std.XYZW[TIn]) std.XYZW[TOut] {
	return std.XYZW[TOut]{
		X: bounded.ScaleToType[TOut](value.X),
		Y: bounded.ScaleToType[TOut](value.Y),
		Z: bounded.ScaleToType[TOut](value.Z),
		W: value.W,
	}
}
