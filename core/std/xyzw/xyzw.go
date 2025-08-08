package xyzw

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/num"
)

// From creates a new instance of std.XYZW[T] bounded in the fully closed interval [0, max].  If you would like to
// also set the minimum boundary, please use FromFull.
//
// NOTE: If you would like the values to be bound by their type's std.MaxValue[T], do not provide a boundary function.
//
// NOTE: If no boundary function is provided and T is a sub-byte type, std.ImplicitOverflow is automatically chosen.
func From[T num.ExtendedPrimitive](x, y, z T, w float64, maxX, maxY, maxZ T) std.XYZW[T] {
	return std.XYZW[T]{}.SetAll(x, y, z, w, 0, maxX, 0, maxY, 0, maxZ)
}

// FromFull creates a new instance of std.XYZW[T] bounded in the fully closed interval [min, max].
//
// NOTE: If you would like the values to be bound by their type's std.MaxValue[T], do not provide a boundary function.
//
// NOTE: If no boundary function is provided and T is a sub-byte type, std.ImplicitOverflow is automatically chosen.
func FromFull[T num.ExtendedPrimitive](x, y, z T, w float64, minX, maxX, minY, maxY, minZ, maxZ T) std.XYZW[T] {
	return std.XYZW[T]{}.SetAll(x, y, z, w, minX, maxX, minY, maxY, minZ, maxZ)
}

// Random returns a pseudo-random std.XYZW[T] of the provided type using math.Random[T], with
// each directional component bounded in the fully closed interval [0, min].  If you would like
// the minimum to be above 0, please use RandomFull
//
// NOTE: W will always return as 1.0
func Random[T num.ExtendedPrimitive](maxX, maxY, maxZ T) std.XYZW[T] {
	x := num.RandomBounded[T](0, maxX)
	y := num.RandomBounded[T](0, maxY)
	z := num.RandomBounded[T](0, maxZ)
	return std.XYZW[T]{}.SetAll(x, y, z, 1.0, 0, maxX, 0, maxY, 0, maxZ)
}

// RandomFull returns a pseudo-random std.XYZW[T] of the provided type using math.Random[T], with
// each directional component bounded in the fully closed interval [max, min].
//
// NOTE: W will always return as 1.0
func RandomFull[T num.ExtendedPrimitive](minX, maxX, minY, maxY, minZ, maxZ T) std.XYZW[T] {
	x := num.RandomBounded[T](minX, maxX)
	y := num.RandomBounded[T](minY, maxY)
	z := num.RandomBounded[T](minZ, maxZ)
	return std.XYZW[T]{}.SetAll(x, y, z, 1.0, minX, maxX, minY, maxY, minZ, maxZ)
}
