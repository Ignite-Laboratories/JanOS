package xyz

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/num"
)

// From creates a new instance of std.XYZ[T] bounded in the fully closed interval [0, max].  If you would like to
// also set the minimum boundary, please use FromFull.
//
// NOTE: If you would like the values to be bound by their type's std.MaxValue[T], do not provide a boundary function.
//
// NOTE: If no boundary function is provided and T is a sub-byte type, std.ImplicitOverflow is automatically chosen.
func From[T num.ExtendedPrimitive](x, y, z T, maxX, maxY, maxZ T) std.XYZ[T] {
	return std.XYZ[T]{}.SetAll(x, y, z, 0, maxX, 0, maxY, 0, maxZ)
}

// FromFull creates a new instance of std.XYZ[T] bounded in the fully closed interval [min, max].
//
// NOTE: If you would like the values to be bound by their type's std.MaxValue[T], do not provide a boundary function.
//
// NOTE: If no boundary function is provided and T is a sub-byte type, std.ImplicitOverflow is automatically chosen.
func FromFull[T num.ExtendedPrimitive](x, y, z T, maxX, maxY, maxZ T) std.XYZ[T] {
	return std.XYZ[T]{}.SetAll(x, y, z, 0, maxX, 0, maxY, 0, maxZ)
}

// Random returns a pseudo-random std.XYZ[T] of the provided type using math.Random[T], with
// each directional component bounded in the fully closed interval [0, min].  If you would like
// the minimum to be above 0, please use RandomFull
func Random[T num.ExtendedPrimitive](maxX, maxY, maxZ T) std.XYZ[T] {
	x := num.RandomBounded[T](0, maxX)
	y := num.RandomBounded[T](0, maxY)
	z := num.RandomBounded[T](0, maxZ)
	return std.XYZ[T]{}.SetAll(x, y, z, 0, maxX, 0, maxY, 0, maxZ)
}

// RandomFull returns a pseudo-random std.XYZ[T] of the provided type using math.Random[T], with
// each directional component bounded in the fully closed interval [max, min].
func RandomFull[T num.ExtendedPrimitive](minX, maxX, minY, maxY, minZ, maxZ T) std.XYZ[T] {
	x := num.RandomBounded[T](minX, maxX)
	y := num.RandomBounded[T](minY, maxY)
	z := num.RandomBounded[T](minZ, maxZ)
	return std.XYZ[T]{}.SetAll(x, y, z, minX, maxX, minY, maxY, minZ, maxZ)
}
