package xyzw

import (
	"github.com/ignite-laboratories/core/math"
	"github.com/ignite-laboratories/core/std"
)

// Random returns a pseudo-random std.XYZW[T] of the provided type using math.RandomNumber[T].
//
// If requesting a floating point type, the resulting number will be bounded
// in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [0, n] - where n is the maximum value of
// the provided type.
func Random[T math.Numeric]() std.XYZW[T] {
	return std.XYZW[T]{
		X: math.RandomNumber[T](),
		Y: math.RandomNumber[T](),
		Z: math.RandomNumber[T](),
		W: math.RandomNumber[T](),
	}
}

// RandomUpTo returns a pseudo-random std.XYZW[T] of the provided type bounded in the closed interval [0, max].
func RandomUpTo[T math.Numeric](xUpper T, yUpper T, zUpper T, wUpper T) std.XYZW[T] {
	return std.XYZW[T]{
		X: math.RandomNumberRange[T](math.Tuple[T]{B: xUpper}),
		Y: math.RandomNumberRange[T](math.Tuple[T]{B: yUpper}),
		Z: math.RandomNumberRange[T](math.Tuple[T]{B: zUpper}),
		W: math.RandomNumberRange[T](math.Tuple[T]{B: wUpper}),
	}
}

// RandomRange returns a pseudo-random std.XYZW[T] of the provided type bounded in the closed interval [min, max].
func RandomRange[T math.Numeric](xRange math.Tuple[T], yRange math.Tuple[T], zRange math.Tuple[T], wRange math.Tuple[T]) std.XYZW[T] {
	return std.XYZW[T]{
		X: math.RandomNumberRange[T](xRange),
		Y: math.RandomNumberRange[T](yRange),
		Z: math.RandomNumberRange[T](zRange),
		W: math.RandomNumberRange[T](wRange),
	}
}

// Normalize32 returns an XYZW[float32] ranging from 0.0-1.0.
func Normalize32[T math.Integer](source std.XYZW[T]) std.XYZW[float32] {
	return std.XYZW[float32]{
		X: math.NormalizeToFloat32(source.X),
		Y: math.NormalizeToFloat32(source.Y),
		Z: math.NormalizeToFloat32(source.Z),
		W: math.NormalizeToFloat32(source.W),
	}
}

// Normalize64 returns an XYZW[float64] ranging from 0.0-1.0.
func Normalize64[T math.Integer](source std.XYZW[T]) std.XYZW[float64] {
	return std.XYZW[float64]{
		X: math.NormalizeToFloat64(source.X),
		Y: math.NormalizeToFloat64(source.Y),
		Z: math.NormalizeToFloat64(source.Z),
		W: math.NormalizeToFloat64(source.W),
	}
}

// ScaleToType32 returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleToType32[TOut math.Integer](source std.XYZW[float32]) std.XYZW[TOut] {
	return std.XYZW[TOut]{
		X: math.ScaleFloat32ToType[TOut](source.X),
		Y: math.ScaleFloat32ToType[TOut](source.Y),
		Z: math.ScaleFloat32ToType[TOut](source.Z),
		W: math.ScaleFloat32ToType[TOut](source.W),
	}
}

// ScaleToType64 returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleToType64[TOut math.Integer](source std.XYZW[float64]) std.XYZW[TOut] {
	return std.XYZW[TOut]{
		X: math.ScaleFloat64ToType[TOut](source.X),
		Y: math.ScaleFloat64ToType[TOut](source.Y),
		Z: math.ScaleFloat64ToType[TOut](source.Z),
		W: math.ScaleFloat64ToType[TOut](source.W),
	}
}

// Comparator returns if the two XYZW values are equal in values.
func Comparator[T math.Numeric](a std.XYZW[T], b std.XYZW[T]) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z && a.W == b.W
}
