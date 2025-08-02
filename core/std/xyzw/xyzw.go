package xyzw

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/normalize"
	"github.com/ignite-laboratories/core/std/num"
)

// From creates a new instance of std.XYZW[T] with the provided values.
func From[T num.ExtendedPrimitive](x, y, z, w T) std.XYZW[T] {
	return std.XYZW[T]{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}

// Random returns a pseudo-random std.XYZW[T] of the provided type using math.Random[T].
//
// If requesting a floating point type, the resulting number will be bounded
// in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [0, n] - where n is the maximum value of
// the provided type.
func Random[T num.ExtendedPrimitive]() std.XYZW[T] {
	return std.XYZW[T]{
		X: std.Random[T](),
		Y: std.Random[T](),
		Z: std.Random[T](),
		W: std.Random[T](),
	}
}

// RandomUpTo returns a pseudo-random std.XYZW[T] of the provided type bounded in the closed interval [0, max].
func RandomUpTo[T num.ExtendedPrimitive](xUpper T, yUpper T, zUpper T, wUpper T) std.XYZW[T] {
	return std.XYZW[T]{
		X: std.RandomBounded[T](0, xUpper),
		Y: std.RandomBounded[T](0, yUpper),
		Z: std.RandomBounded[T](0, zUpper),
		W: std.RandomBounded[T](0, wUpper),
	}
}

// RandomRange returns a pseudo-random std.XYZW[T] of the provided type bounded in the closed interval [min, max].
func RandomRange[T num.ExtendedPrimitive](minimum, maximum T) std.XYZW[T] {
	return std.XYZW[T]{
		X: std.RandomBounded[T](minimum, maximum),
		Y: std.RandomBounded[T](minimum, maximum),
		Z: std.RandomBounded[T](minimum, maximum),
		W: std.RandomBounded[T](minimum, maximum),
	}
}

// Normalize returns an std.XYZW[TOut] ranging from 0.0-1.0.
func Normalize[TIn num.ExtendedPrimitive, TOut num.Float](source std.XYZW[TIn]) std.XYZW[TOut] {
	return std.XYZW[TOut]{
		X: normalize.To[TIn, TOut](source.X),
		Y: normalize.To[TIn, TOut](source.Y),
		Z: normalize.To[TIn, TOut](source.Z),
		W: normalize.To[TIn, TOut](source.W),
	}
}

// ReScale returns an std.XYZW[TOut] scaled up to [0, TIn.Max] from an input bounded in the fully closed interval [0.0, 1.0].
func ReScale[TIn num.Float, TOut num.Integer](source std.XYZW[TIn]) std.XYZW[TOut] {
	return std.XYZW[TOut]{
		X: normalize.From[TIn, TOut](source.X),
		Y: normalize.From[TIn, TOut](source.Y),
		Z: normalize.From[TIn, TOut](source.Z),
		W: normalize.From[TIn, TOut](source.W),
	}
}

// Comparator returns if the two std.XYZW values are equal in values.
func Comparator[T num.ExtendedPrimitive](a std.XYZW[T], b std.XYZW[T]) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z && a.W == b.W
}
