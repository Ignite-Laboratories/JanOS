package xyz

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/sys/number"
	"github.com/ignite-laboratories/core/sys/number/normalize"
)

// From creates a new instance of std.XYZ[T] with the provided values.
func From[T number.Numeric](x, y, z T) std.XYZ[T] {
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
func Random[T number.Numeric]() std.XYZ[T] {
	return std.XYZ[T]{
		X: number.Random[T](),
		Y: number.Random[T](),
		Z: number.Random[T](),
	}
}

// RandomUpTo returns a pseudo-random std.XYZ[T] of the provided type bounded in the closed interval [0, max].
func RandomUpTo[T number.Numeric](xUpper T, yUpper T, zUpper T) std.XYZ[T] {
	return std.XYZ[T]{
		X: number.RandomBounded[T](0, xUpper),
		Y: number.RandomBounded[T](0, yUpper),
		Z: number.RandomBounded[T](0, zUpper),
	}
}

// RandomRange returns a pseudo-random std.XYZ[T] of the provided type bounded in the closed interval [min, max].
func RandomRange[T number.Numeric](minimum, maximum T) std.XYZ[T] {
	return std.XYZ[T]{
		X: number.RandomBounded[T](minimum, maximum),
		Y: number.RandomBounded[T](minimum, maximum),
		Z: number.RandomBounded[T](minimum, maximum),
	}
}

// Normalize returns an std.XYZ[TOut] ranging from 0.0-1.0.
func Normalize[TIn number.Numeric, TOut number.Float](source std.XYZ[TIn]) std.XYZ[TOut] {
	return std.XYZ[TOut]{
		X: normalize.To[TIn, TOut](source.X),
		Y: normalize.To[TIn, TOut](source.Y),
		Z: normalize.To[TIn, TOut](source.Z),
	}
}

// ReScale returns an std.XYZ[TOut] scaled up to [0, TIn.Max] from an input bounded in the fully closed interval [0.0, 1.0].
func ReScale[TIn number.Float, TOut number.Integer](source std.XYZ[TIn]) std.XYZ[TOut] {
	return std.XYZ[TOut]{
		X: normalize.From[TIn, TOut](source.X),
		Y: normalize.From[TIn, TOut](source.Y),
		Z: normalize.From[TIn, TOut](source.Z),
	}
}

// Comparator returns if the two std.XYZ values are equal in values.
func Comparator[T number.Numeric](a std.XYZ[T], b std.XYZ[T]) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z
}
