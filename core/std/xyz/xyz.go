package xyz

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/num"
)

// From creates a new instance of std.XYZ[T] with each direction bounded in the fully closed interval [0, max].  If you
// would like to also set the minimum boundaries, please use FromFull.
func From[T num.ExtendedPrimitive](x, y, z T, maxX, maxY, maxZ T) std.XYZ[T] {
	return std.XYZ[T]{}.SetAll(x, y, z, 0, maxX, 0, maxY, 0, maxZ)
}

// FromFull creates a new instance of std.XYZ[T] with each direction bounded in the fully closed interval [min, max].
func FromFull[T num.ExtendedPrimitive](x, y, z T, minX, maxX, minY, maxY, minZ, maxZ T) std.XYZ[T] {
	return std.XYZ[T]{}.SetAll(x, y, z, minX, maxX, minY, maxY, minZ, maxZ)
}

// Random returns a pseudo-random std.XYZ[T] of the provided type using math.Random[T], with
// each directional component bounded in the fully closed interval [0, max].  If you would like
// the minimum to be above 0, please use RandomFull
func Random[T num.ExtendedPrimitive](maxX, maxY, maxZ T) std.XYZ[T] {
	x := num.RandomWithinRange[T](0, maxX)
	y := num.RandomWithinRange[T](0, maxY)
	z := num.RandomWithinRange[T](0, maxZ)
	return std.XYZ[T]{}.SetAll(x, y, z, 0, maxX, 0, maxY, 0, maxZ)
}

// RandomFull returns a pseudo-random std.XYZ[T] of the provided type using math.Random[T], with
// each directional component bounded in the fully closed interval [min, max].
func RandomFull[T num.ExtendedPrimitive](minX, maxX, minY, maxY, minZ, maxZ T) std.XYZ[T] {
	x := num.RandomWithinRange[T](minX, maxX)
	y := num.RandomWithinRange[T](minY, maxY)
	z := num.RandomWithinRange[T](minZ, maxZ)
	return std.XYZ[T]{}.SetAll(x, y, z, minX, maxX, minY, maxY, minZ, maxZ)
}

// ScaleToType normalizes the std.XY[T] directional components into unit vectors and then scales them to a new std.XY[TOut].
func ScaleToType[TIn num.ExtendedPrimitive, TOut num.ExtendedPrimitive](value std.XYZ[TIn]) std.XYZ[TOut] {
	return std.XYZ[TOut]{
		X: bounded.ScaleToType[TOut](value.X),
		Y: bounded.ScaleToType[TOut](value.Y),
		Z: bounded.ScaleToType[TOut](value.Z),
	}
}
