package rgba

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/sys/number"
	"github.com/ignite-laboratories/core/sys/number/normalize"
)

// Normalize returns an RGBA[TOut] ranging from 0.0-1.0.
func Normalize[TIn number.Numeric, TOut number.Float](c std.RGBA[TIn]) std.RGBA[TOut] {
	return std.RGBA[TOut]{
		R: normalize.To[TIn, TOut](c.R),
		G: normalize.To[TIn, TOut](c.G),
		B: normalize.To[TIn, TOut](c.B),
		A: normalize.To[TIn, TOut](c.A),
	}
}

// ReScale returns an RGBA[TOut] scaled up to [0, TIn.Max] from an input bounded in the fully closed interval [0.0, 1.0].
func ReScale[TIn number.Float, TOut number.Integer](c std.RGBA[TIn]) std.RGBA[TOut] {
	return std.RGBA[TOut]{
		R: normalize.From[TIn, TOut](c.R),
		G: normalize.From[TIn, TOut](c.G),
		B: normalize.From[TIn, TOut](c.B),
		A: normalize.From[TIn, TOut](c.A),
	}
}

// Comparator returns if the two RGBA values are equal in values.
func Comparator[T number.Numeric](a std.RGBA[T], b std.RGBA[T]) bool {
	return a.R == b.R && a.G == b.G && a.B == b.B && a.A == b.A
}

// From converts the provided value into a std.RGBA[byte].  For example:
//
//	rgba.From(0xaabbccdd):
//	    R: 170 [0xAA]
//	    G: 187 [0xBB]
//	    B: 204 [0xCC]
//	    A: 221 [0xDD]
func From(value uint32) std.RGBA[byte] {
	return std.RGBA[byte]{
		R: byte((value >> 24) & 0xFF),
		G: byte((value >> 16) & 0xFF),
		B: byte((value >> 8) & 0xFF),
		A: byte(value & 0xFF),
	}
}

// Random returns a pseudo-random std.RGBA[T] of the provided type using math.Random[T].
//
// If requesting a floating point type, the resulting number will be bounded
// in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [0, n] - where n is the maximum value of
// the provided type.
func Random[T number.Numeric]() std.RGBA[T] {
	return std.RGBA[T]{
		R: number.Random[T](),
		G: number.Random[T](),
		B: number.Random[T](),
		A: number.Random[T](),
	}
}

// RandomUpTo returns a pseudo-random std.RGBA[T] of the provided type with each channel bounded within its provided closed interval of [0, max].
func RandomUpTo[T number.Numeric](rUpper T, gUpper T, bUpper T, aUpper T) std.RGBA[T] {
	return std.RGBA[T]{
		R: number.RandomBounded[T](0, rUpper),
		G: number.RandomBounded[T](0, gUpper),
		B: number.RandomBounded[T](0, bUpper),
		A: number.RandomBounded[T](0, aUpper),
	}
}

// RandomRange returns a pseudo-random std.RGBA[T] of the provided type bounded in the closed interval [min, max].
func RandomRange[T number.Numeric](minimum, maximum T) std.RGBA[T] {
	return std.RGBA[T]{
		R: number.RandomBounded[T](minimum, maximum),
		G: number.RandomBounded[T](minimum, maximum),
		B: number.RandomBounded[T](minimum, maximum),
		A: number.RandomBounded[T](minimum, maximum),
	}
}
