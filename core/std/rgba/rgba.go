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
func Random[T number.Numeric]() std.RGBA[T] {
	return std.RGBA[T]{
		R: number.RandomNumber[T](),
		G: number.RandomNumber[T](),
		B: number.RandomNumber[T](),
		A: number.RandomNumber[T](),
	}
}

// RandomUpTo returns a pseudo-random std.RGBA[T] of the provided type bounded in the closed interval [0, max].
func RandomUpTo[T number.Numeric](rUpper T, gUpper T, bUpper T, aUpper T) std.RGBA[T] {
	return std.RGBA[T]{
		R: number.RandomNumberRange[T](number.Tuple[T]{B: rUpper}),
		G: number.RandomNumberRange[T](number.Tuple[T]{B: gUpper}),
		B: number.RandomNumberRange[T](number.Tuple[T]{B: bUpper}),
		A: number.RandomNumberRange[T](number.Tuple[T]{B: aUpper}),
	}
}

// RandomRange returns a pseudo-random std.RGBA[T] of the provided type bounded in the closed interval [min, max].
func RandomRange[T number.Numeric](rRange number.Tuple[T], gRange number.Tuple[T], bRange number.Tuple[T], aRange number.Tuple[T]) std.RGBA[T] {
	return std.RGBA[T]{
		R: number.RandomNumberRange[T](rRange),
		G: number.RandomNumberRange[T](gRange),
		B: number.RandomNumberRange[T](bRange),
		A: number.RandomNumberRange[T](aRange),
	}
}
