package rgba

import (
	"github.com/ignite-laboratories/core/math"
	"github.com/ignite-laboratories/core/std"
)

// FromHex converts the provided RGBA hex values into a std.std.RGBA[byte].
func FromHex(value uint32) std.RGBA[byte] {
	return std.RGBA[byte]{
		R: byte((value >> 24) & 0xFF),
		G: byte((value >> 16) & 0xFF),
		B: byte((value >> 8) & 0xFF),
		A: byte(value & 0xFF),
	}
}

// Random returns a pseudo-random std.RGBA[T] of the provided type using math.RandomNumber[T].
//
// If requesting a floating point type, the resulting number will be bounded
// in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [0, n] - where n is the maximum value of
// the provided type.
func Random[T math.Numeric]() std.RGBA[T] {
	return std.RGBA[T]{
		R: math.RandomNumber[T](),
		G: math.RandomNumber[T](),
		B: math.RandomNumber[T](),
		A: math.RandomNumber[T](),
	}
}

// RandomUpTo returns a pseudo-random std.RGBA[T] of the provided type bounded in the closed interval [0, max].
func RandomUpTo[T math.Numeric](rUpper T, gUpper T, bUpper T, aUpper T) std.RGBA[T] {
	return std.RGBA[T]{
		R: math.RandomNumberRange[T](math.Tuple[T]{B: rUpper}),
		G: math.RandomNumberRange[T](math.Tuple[T]{B: gUpper}),
		B: math.RandomNumberRange[T](math.Tuple[T]{B: bUpper}),
		A: math.RandomNumberRange[T](math.Tuple[T]{B: aUpper}),
	}
}

// RandomRange returns a pseudo-random std.RGBA[T] of the provided type bounded in the closed interval [min, max].
func RandomRange[T math.Numeric](rRange math.Tuple[T], gRange math.Tuple[T], bRange math.Tuple[T], aRange math.Tuple[T]) std.RGBA[T] {
	return std.RGBA[T]{
		R: math.RandomNumberRange[T](rRange),
		G: math.RandomNumberRange[T](gRange),
		B: math.RandomNumberRange[T](bRange),
		A: math.RandomNumberRange[T](aRange),
	}
}
