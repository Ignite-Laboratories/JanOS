package rgb

import (
	"github.com/ignite-laboratories/core/math"
	"github.com/ignite-laboratories/core/std"
)

// FromHex converts the provided RGB hex values into a std.RGBA[byte].
//
// The alpha channel can optionally be provided, otherwise it defaults to 0.
func FromHex(value uint32, alpha ...byte) std.RGBA[byte] {
	a := byte(0)
	if len(alpha) > 0 {
		a = alpha[0]
	}

	return std.RGBA[byte]{
		R: byte((value >> 16) & 0xFF),
		G: byte((value >> 8) & 0xFF),
		B: byte(value & 0xFF),
		A: a,
	}
}

// Random returns a pseudo-random RGB[T] of the provided type using math.RandomNumber[T].
//
// If requesting a floating point type, the resulting number will be bounded
// in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [0, n] - where n is the maximum value of
// the provided type.
//
// The alpha channel can optionally be provided, otherwise it defaults to 0.
func Random[T number.Numeric](alpha ...T) std.RGBA[T] {
	a := T(0)
	if len(alpha) > 0 {
		a = alpha[0]
	}

	return std.RGBA[T]{
		R: number.RandomNumber[T](),
		G: number.RandomNumber[T](),
		B: number.RandomNumber[T](),
		A: a,
	}
}

// RandomUpTo returns a pseudo-random RGB[T] of the provided type bounded in the closed interval [0, max].
//
// The alpha channel can optionally be provided, otherwise it defaults to 0.
func RandomUpTo[T number.Numeric](rUpper T, gUpper T, bUpper T, alpha ...T) std.RGBA[T] {
	a := T(0)
	if len(alpha) > 0 {
		a = alpha[0]
	}

	return std.RGBA[T]{
		R: number.RandomNumberRange[T](number.Tuple[T]{B: rUpper}),
		G: number.RandomNumberRange[T](number.Tuple[T]{B: gUpper}),
		B: number.RandomNumberRange[T](number.Tuple[T]{B: bUpper}),
		A: a,
	}
}

// RandomRange returns a pseudo-random RGB[T] of the provided type bounded in the closed interval [min, max].
func RandomRange[T number.Numeric](rRange number.Tuple[T], gRange number.Tuple[T], bRange number.Tuple[T], alpha ...T) std.RGBA[T] {
	a := T(0)
	if len(alpha) > 0 {
		a = alpha[0]
	}

	return std.RGBA[T]{
		R: number.RandomNumberRange[T](rRange),
		G: number.RandomNumberRange[T](gRange),
		B: number.RandomNumberRange[T](bRange),
		A: a,
	}
}
