package xy

import (
	"github.com/ignite-laboratories/core/math"
	"github.com/ignite-laboratories/core/std"
)

// Random returns a pseudo-random std.XY[T] of the provided type using math.RandomNumber[T].
//
// If requesting a floating point type, the resulting number will be bounded
// in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [0, n] - where n is the maximum value of
// the provided type.
func Random[T number.Numeric]() std.XY[T] {
	return std.XY[T]{
		X: number.RandomNumber[T](),
		Y: number.RandomNumber[T](),
	}
}

// RandomUpTo returns a pseudo-random std.XY[T] of the provided type bounded in the closed interval [0, max].
func RandomUpTo[T number.Numeric](xUpper T, yUpper T) std.XY[T] {
	return std.XY[T]{
		X: number.RandomNumberRange[T](number.Tuple[T]{B: xUpper}),
		Y: number.RandomNumberRange[T](number.Tuple[T]{B: yUpper}),
	}
}

// RandomRange returns a pseudo-random std.XY[T] of the provided type bounded in the closed interval [min, max].
func RandomRange[T number.Numeric](xRange number.Tuple[T], yRange number.Tuple[T]) std.XY[T] {
	return std.XY[T]{
		X: number.RandomNumberRange[T](xRange),
		Y: number.RandomNumberRange[T](yRange),
	}
}

// Normalize32 returns an std.XY[float32] ranging from 0.0-1.0.
func Normalize32[T number.Integer](source std.XY[T]) std.XY[float32] {
	return std.XY[float32]{
		X: number.NormalizeToFloat32(source.X),
		Y: number.NormalizeToFloat32(source.Y),
	}
}

// Normalize64 returns an XYZ[float64] ranging from 0.0-1.0.
func Normalize64[T number.Integer](source std.XY[T]) std.XY[float64] {
	return std.XY[float64]{
		X: number.NormalizeToFloat64(source.X),
		Y: number.NormalizeToFloat64(source.Y),
	}
}

// ScaleToType32 returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleToType32[TOut number.Integer](source std.XY[float32]) std.XY[TOut] {
	return std.XY[TOut]{
		X: number.ScaleFloat32ToType[TOut](source.X),
		Y: number.ScaleFloat32ToType[TOut](source.Y),
	}
}

// ScaleToType64 returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleToType64[TOut number.Integer](source std.XY[float64]) std.XY[TOut] {
	return std.XY[TOut]{
		X: number.ScaleFloat64ToType[TOut](source.X),
		Y: number.ScaleFloat64ToType[TOut](source.Y),
	}
}

// Comparator returns if the two XY values are equal in values.
func Comparator[T number.Numeric](a std.XY[T], b std.XY[T]) bool {
	return a.X == b.X && a.Y == b.Y
}
