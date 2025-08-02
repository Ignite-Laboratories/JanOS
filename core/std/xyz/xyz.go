package xyz

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/normalize"
	"github.com/ignite-laboratories/core/std/num"
)

// From creates a new instance of std.XYZ[T] with the provided values.
func From[T num.ExtendedPrimitive](x, y, z T) std.XYZ[T] {
	return std.XYZ[T]{
		X: x,
		Y: y,
		Z: z,
	}
}

// Random returns a pseudo-random std.XYZ[T] of the provided type using math.Random[T].
//
// If requesting a floating point type, the resulting number will be bounded
// in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [0, n] - where n is the maximum value of
// the provided type.
func Random[T num.ExtendedPrimitive]() std.XYZ[T] {
	return std.XYZ[T]{
		X: std.Random[T](),
		Y: std.Random[T](),
		Z: std.Random[T](),
	}
}

// RandomUpTo returns a pseudo-random std.XYZ[T] of the provided type bounded in the closed interval [0, max].
func RandomUpTo[T num.ExtendedPrimitive](xUpper T, yUpper T, zUpper T) std.XYZ[T] {
	return std.XYZ[T]{
		X: std.RandomBounded[T](0, xUpper),
		Y: std.RandomBounded[T](0, yUpper),
		Z: std.RandomBounded[T](0, zUpper),
	}
}

// RandomRange returns a pseudo-random std.XYZ[T] of the provided type bounded in the closed interval [min, max].
func RandomRange[T num.ExtendedPrimitive](minimum, maximum T) std.XYZ[T] {
	return std.XYZ[T]{
		X: std.RandomBounded[T](minimum, maximum),
		Y: std.RandomBounded[T](minimum, maximum),
		Z: std.RandomBounded[T](minimum, maximum),
	}
}

// Normalize returns an std.XYZ[TOut] ranging from 0.0-1.0.
func Normalize[TIn num.ExtendedPrimitive, TOut num.Float](source std.XYZ[TIn]) std.XYZ[TOut] {
	return std.XYZ[TOut]{
		X: normalize.To[TIn, TOut](source.X),
		Y: normalize.To[TIn, TOut](source.Y),
		Z: normalize.To[TIn, TOut](source.Z),
	}
}

// ReScale returns an std.XYZ[TOut] scaled up to [0, TIn.Max] from an input bounded in the fully closed interval [0.0, 1.0].
func ReScale[TIn num.Float, TOut num.Integer](source std.XYZ[TIn]) std.XYZ[TOut] {
	return std.XYZ[TOut]{
		X: normalize.From[TIn, TOut](source.X),
		Y: normalize.From[TIn, TOut](source.Y),
		Z: normalize.From[TIn, TOut](source.Z),
	}
}

// Comparator returns if the two std.XYZ values are equal in values.
func Comparator[T num.ExtendedPrimitive](a std.XYZ[T], b std.XYZ[T]) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z
}
