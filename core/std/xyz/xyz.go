package xyz

import (
	"github.com/ignite-laboratories/core/math"
	"github.com/ignite-laboratories/core/std"
)

// Random returns a pseudo-random std.XYZ[T] of the provided type using math.RandomNumber[T].
//
// If requesting a floating point type, the resulting number will be bounded
// in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [0, n] - where n is the maximum value of
// the provided type.
func Random[T number.Numeric]() std.XYZ[T] {
	return std.XYZ[T]{
		X: number.RandomNumber[T](),
		Y: number.RandomNumber[T](),
		Z: number.RandomNumber[T](),
	}
}

// RandomUpTo returns a pseudo-random std.XYZ[T] of the provided type bounded in the closed interval [0, max].
func RandomUpTo[T number.Numeric](xUpper T, yUpper T, zUpper T) std.XYZ[T] {
	return std.XYZ[T]{
		X: number.RandomNumberRange[T](number.Tuple[T]{B: xUpper}),
		Y: number.RandomNumberRange[T](number.Tuple[T]{B: yUpper}),
		Z: number.RandomNumberRange[T](number.Tuple[T]{B: zUpper}),
	}
}

// RandomRange returns a pseudo-random std.XYZ[T] of the provided type bounded in the closed interval [min, max].
func RandomRange[T number.Numeric](xRange number.Tuple[T], yRange number.Tuple[T], zRange number.Tuple[T]) std.XYZ[T] {
	return std.XYZ[T]{
		X: number.RandomNumberRange[T](xRange),
		Y: number.RandomNumberRange[T](yRange),
		Z: number.RandomNumberRange[T](zRange),
	}
}

// Normalize32 returns an std.XYZ[float32] ranging from 0.0-1.0.
func Normalize32[T number.Integer](source std.XYZ[T]) std.XYZ[float32] {
	return std.XYZ[float32]{
		X: number.NormalizeToFloat32(source.X),
		Y: number.NormalizeToFloat32(source.Y),
		Z: number.NormalizeToFloat32(source.Z),
	}
}

// Normalize64 returns an std.XYZ[float64] ranging from 0.0-1.0.
func Normalize64[T number.Integer](source std.XYZ[T]) std.XYZ[float64] {
	return std.XYZ[float64]{
		X: number.NormalizeToFloat64(source.X),
		Y: number.NormalizeToFloat64(source.Y),
		Z: number.NormalizeToFloat64(source.Z),
	}
}

// ScaleToType32 returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleToType32[TOut number.Integer](source std.XYZ[float32]) std.XYZ[TOut] {
	return std.XYZ[TOut]{
		X: number.ScaleFloat32ToType[TOut](source.X),
		Y: number.ScaleFloat32ToType[TOut](source.Y),
		Z: number.ScaleFloat32ToType[TOut](source.Z),
	}
}

// ScaleToType64 returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleToType64[TOut number.Integer](source std.XYZ[float64]) std.XYZ[TOut] {
	return std.XYZ[TOut]{
		X: number.ScaleFloat64ToType[TOut](source.X),
		Y: number.ScaleFloat64ToType[TOut](source.Y),
		Z: number.ScaleFloat64ToType[TOut](source.Z),
	}
}

// Comparator returns if the two XYZ values are equal in values.
func Comparator[T number.Numeric](a std.XYZ[T], b std.XYZ[T]) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z
}
