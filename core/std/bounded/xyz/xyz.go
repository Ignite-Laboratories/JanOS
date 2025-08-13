package xyz

import (
	"github.com/ignite-laboratories/core/std/bounded"
	"github.com/ignite-laboratories/core/std/num"
)

// From creates a new instance of std.XYZ[T] with each direction bounded in the fully closed interval [0, max].  If you
// would like to also set the minimum boundaries, please use FromFull.
func From[T num.Primitive](x, y, z T, maxX, maxY, maxZ T) bounded.XYZ[T] {
	return bounded.XYZ[T]{}.SetAll(x, y, z, 0, maxX, 0, maxY, 0, maxZ)
}

// FromFull creates a new instance of std.XYZ[T] with each direction bounded in the fully closed interval [min, max].
func FromFull[T num.Primitive](x, y, z T, minX, maxX, minY, maxY, minZ, maxZ T) bounded.XYZ[T] {
	return bounded.XYZ[T]{}.SetAll(x, y, z, minX, maxX, minY, maxY, minZ, maxZ)
}

// Random returns a pseudo-random std.XYZ[T] of the provided type using math.Random[T], with
// each directional component bounded in the fully closed interval [0, max].  If you would like
// the minimum to be above 0, please use RandomFull
func Random[T num.Primitive](maxX, maxY, maxZ T) bounded.XYZ[T] {
	x := num.RandomWithinRange[T](0, maxX)
	y := num.RandomWithinRange[T](0, maxY)
	z := num.RandomWithinRange[T](0, maxZ)
	return bounded.XYZ[T]{}.SetAll(x, y, z, 0, maxX, 0, maxY, 0, maxZ)
}

// RandomFull returns a pseudo-random std.XYZ[T] of the provided type using math.Random[T], with
// each directional component bounded in the fully closed interval [min, max].
func RandomFull[T num.Primitive](minX, maxX, minY, maxY, minZ, maxZ T) bounded.XYZ[T] {
	x := num.RandomWithinRange[T](minX, maxX)
	y := num.RandomWithinRange[T](minY, maxY)
	z := num.RandomWithinRange[T](minZ, maxZ)
	return bounded.XYZ[T]{}.SetAll(x, y, z, minX, maxX, minY, maxY, minZ, maxZ)
}

// ScaleToType normalizes the std.XY[T] directional components into unit vectors and then scales them to a new std.XY[TOut].
func ScaleToType[TIn num.Primitive, TOut num.Primitive](value bounded.XYZ[TIn]) bounded.XYZ[TOut] {
	return bounded.XYZ[TOut]{
		X: bounded.ScaleToType[TOut](value.X),
		Y: bounded.ScaleToType[TOut](value.Y),
		Z: bounded.ScaleToType[TOut](value.Z),
	}
}
