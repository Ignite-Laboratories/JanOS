// Package rgb provides access to all std.RGB procedures.
package rgb

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/normalize"
	"github.com/ignite-laboratories/core/std/num"
)

// From constructs a std.RGB[TInt] value from individual red, green, blue, and alpha components of type TInt.
//
// NOTE: If you provide a sub-byte size, each channel's value will be modulo-d against 2ⁿ, with 𝑛 being the sub-byte bit-width.
func From[TInt num.ExtendedInteger](r, g, b TInt) std.RGB[TInt] {
	return FromHex[TInt]((uint32(r) << 24) | (uint32(g) << 16) | (uint32(b) << 8))
}

// FromHex converts the provided uint32 value into a std.RGB[TInt].  For example -
//
//	rgba.FromHex(0xaabbcc):
//	    R: 170 [0xAA]
//	    G: 187 [0xBB]
//	    B: 204 [0xCC]
//
// NOTE: If you provide a sub-byte size, each channel of the 32-bit value will be modulo-d against 2ⁿ, with 𝑛 being the sub-byte bit-width.
func FromHex[TInt num.ExtendedInteger](value uint32) std.RGB[TInt] {
	r := TInt((value >> 24) & 0xFF)
	g := TInt((value >> 16) & 0xFF)
	b := TInt((value >> 8) & 0xFF)
	return std.RGB[TInt]{}.Set(r, g, b)
}

// Normalize returns an RGBA[TOut] ranging from 0.0-1.0.
func Normalize[TIn num.ExtendedPrimitive, TOut num.Float](c std.RGB[TIn]) std.RGB[TOut] {
	r := normalize.To[TIn, TOut](c.Red())
	g := normalize.To[TIn, TOut](c.Green())
	b := normalize.To[TIn, TOut](c.Blue())
	return std.RGB[TOut]{}.Set(r, g, b)
}

// ReScale returns an RGBA[TOut] scaled up to [0, TIn.Max] from an input bounded in the fully closed interval [0.0, 1.0].
func ReScale[TIn num.Float, TOut num.ExtendedInteger](c std.RGB[TIn]) std.RGB[TOut] {
	r := normalize.From[TIn, TOut](c.Red())
	g := normalize.From[TIn, TOut](c.Green())
	b := normalize.From[TIn, TOut](c.Blue())
	return std.RGB[TOut]{}.Set(r, g, b)
}

// Comparator returns if the two RGBA values are equal in values.
func Comparator[T num.ExtendedPrimitive](a std.RGB[T], b std.RGB[T]) bool {
	return a.R == b.R && a.G == b.G && a.B == b.B
}

// Random returns a pseudo-random std.RGB[T] of the provided type using math.Random[T].
//
// If requesting a floating point type, the resulting number will be bounded
// in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [0, n] - where n is the maximum value of
// the provided type.
func Random[T num.ExtendedPrimitive]() std.RGB[T] {
	r := std.Random[T]()
	g := std.Random[T]()
	b := std.Random[T]()
	return std.RGB[T]{}.Set(r, g, b)
}

// RandomUpTo returns a pseudo-random std.RGB[T] of the provided type with each channel bounded within its provided closed interval of [0, max].
func RandomUpTo[T num.ExtendedPrimitive](rUpper T, gUpper T, bUpper T) std.RGB[T] {
	r := std.RandomBounded[T](0, rUpper)
	g := std.RandomBounded[T](0, gUpper)
	b := std.RandomBounded[T](0, bUpper)
	return std.RGB[T]{}.Set(r, g, b)
}

// RandomRange returns a pseudo-random std.RGB[T] of the provided type bounded in the closed interval [min, max].
func RandomRange[T num.ExtendedPrimitive](minimum, maximum T) std.RGB[T] {
	r := std.RandomBounded[T](minimum, maximum)
	g := std.RandomBounded[T](minimum, maximum)
	b := std.RandomBounded[T](minimum, maximum)
	return std.RGB[T]{}.Set(r, g, b)
}
