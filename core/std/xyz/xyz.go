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
func Random[T math.Numeric]() std.XYZ[T] {
	return std.XYZ[T]{
		X: math.RandomNumber[T](),
		Y: math.RandomNumber[T](),
		Z: math.RandomNumber[T](),
	}
}

// RandomUpTo returns a pseudo-random std.XYZ[T] of the provided type bounded in the closed interval [0, max].
func RandomUpTo[T math.Numeric](xUpper T, yUpper T, zUpper T) std.XYZ[T] {
	return std.XYZ[T]{
		X: math.RandomNumberRange[T](math.Tuple[T]{B: xUpper}),
		Y: math.RandomNumberRange[T](math.Tuple[T]{B: yUpper}),
		Z: math.RandomNumberRange[T](math.Tuple[T]{B: zUpper}),
	}
}

// RandomRange returns a pseudo-random std.XYZ[T] of the provided type bounded in the closed interval [min, max].
func RandomRange[T math.Numeric](xRange math.Tuple[T], yRange math.Tuple[T], zRange math.Tuple[T]) std.XYZ[T] {
	return std.XYZ[T]{
		X: math.RandomNumberRange[T](xRange),
		Y: math.RandomNumberRange[T](yRange),
		Z: math.RandomNumberRange[T](zRange),
	}
}

// Normalize32 returns an std.XYZ[float32] ranging from 0.0-1.0.
func Normalize32[T math.Integer](source std.XYZ[T]) std.XYZ[float32] {
	return std.XYZ[float32]{
		X: math.NormalizeToFloat32(source.X),
		Y: math.NormalizeToFloat32(source.Y),
		Z: math.NormalizeToFloat32(source.Z),
	}
}

// Normalize64 returns an std.XYZ[float64] ranging from 0.0-1.0.
func Normalize64[T math.Integer](source std.XYZ[T]) std.XYZ[float64] {
	return std.XYZ[float64]{
		X: math.NormalizeToFloat64(source.X),
		Y: math.NormalizeToFloat64(source.Y),
		Z: math.NormalizeToFloat64(source.Z),
	}
}

// ScaleToType32 returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleToType32[TOut math.Integer](source std.XYZ[float32]) std.XYZ[TOut] {
	return std.XYZ[TOut]{
		X: math.ScaleFloat32ToType[TOut](source.X),
		Y: math.ScaleFloat32ToType[TOut](source.Y),
		Z: math.ScaleFloat32ToType[TOut](source.Z),
	}
}

// ScaleToType64 returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleToType64[TOut math.Integer](source std.XYZ[float64]) std.XYZ[TOut] {
	return std.XYZ[TOut]{
		X: math.ScaleFloat64ToType[TOut](source.X),
		Y: math.ScaleFloat64ToType[TOut](source.Y),
		Z: math.ScaleFloat64ToType[TOut](source.Z),
	}
}

// Comparator returns if the two XYZ values are equal in values.
func Comparator[T math.Numeric](a std.XYZ[T], b std.XYZ[T]) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z
}
