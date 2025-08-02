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
func Random[T math.Numeric]() std.XY[T] {
	return std.XY[T]{
		X: math.RandomNumber[T](),
		Y: math.RandomNumber[T](),
	}
}

// RandomUpTo returns a pseudo-random std.XY[T] of the provided type bounded in the closed interval [0, max].
func RandomUpTo[T math.Numeric](xUpper T, yUpper T) std.XY[T] {
	return std.XY[T]{
		X: math.RandomNumberRange[T](math.Tuple[T]{B: xUpper}),
		Y: math.RandomNumberRange[T](math.Tuple[T]{B: yUpper}),
	}
}

// RandomRange returns a pseudo-random std.XY[T] of the provided type bounded in the closed interval [min, max].
func RandomRange[T math.Numeric](xRange math.Tuple[T], yRange math.Tuple[T]) std.XY[T] {
	return std.XY[T]{
		X: math.RandomNumberRange[T](xRange),
		Y: math.RandomNumberRange[T](yRange),
	}
}

// Normalize32 returns an std.XY[float32] ranging from 0.0-1.0.
func Normalize32[T math.Integer](source std.XY[T]) std.XY[float32] {
	return std.XY[float32]{
		X: math.NormalizeToFloat32(source.X),
		Y: math.NormalizeToFloat32(source.Y),
	}
}

// Normalize64 returns an XYZ[float64] ranging from 0.0-1.0.
func Normalize64[T math.Integer](source std.XY[T]) std.XY[float64] {
	return std.XY[float64]{
		X: math.NormalizeToFloat64(source.X),
		Y: math.NormalizeToFloat64(source.Y),
	}
}

// ScaleToType32 returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleToType32[TOut math.Integer](source std.XY[float32]) std.XY[TOut] {
	return std.XY[TOut]{
		X: math.ScaleFloat32ToType[TOut](source.X),
		Y: math.ScaleFloat32ToType[TOut](source.Y),
	}
}

// ScaleToType64 returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleToType64[TOut math.Integer](source std.XY[float64]) std.XY[TOut] {
	return std.XY[TOut]{
		X: math.ScaleFloat64ToType[TOut](source.X),
		Y: math.ScaleFloat64ToType[TOut](source.Y),
	}
}

// Comparator returns if the two XY values are equal in values.
func Comparator[T math.Numeric](a std.XY[T], b std.XY[T]) bool {
	return a.X == b.X && a.Y == b.Y
}
