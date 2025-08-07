// Package rgba provides access to all std.RGBA procedures.
package rgba

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/normalize"
	"github.com/ignite-laboratories/core/std/num"
)

// From constructs a std.RGBA[TInt] value from individual red, green, blue, and alpha components of type TInt.
//
// NOTE: If you provide a sub-byte size, each channel's value will be modulo-d against 2ⁿ, with 𝑛 being the sub-byte bit-width.
func From[TInt num.ExtendedInteger](r, g, b, a TInt) std.RGBA[TInt] {
	return FromHex[TInt]((uint32(r) << 24) | (uint32(g) << 16) | (uint32(b) << 8) | uint32(a))
}

// FromHex converts the provided uint32 value into a std.RGBA[TInt].  For example -
//
//	rgba.FromHex(0xaabbccdd):
//	    R: 170 [0xAA]
//	    G: 187 [0xBB]
//	    B: 204 [0xCC]
//	    A: 221 [0xDD]
//
// NOTE: If you provide a sub-byte size, each channel of the 32-bit value will be modulo-d against 2ⁿ, with 𝑛 being the sub-byte bit-width.
func FromHex[TInt num.ExtendedInteger](value uint32) std.RGBA[TInt] {
	r := TInt((value >> 24) & 0xFF)
	g := TInt((value >> 16) & 0xFF)
	b := TInt((value >> 8) & 0xFF)
	a := TInt(value & 0xFF)
	return std.RGBA[TInt]{}.Set(r, g, b, a)
}

// Normalize returns an RGBA[TOut] ranging from 0.0-1.0.
func Normalize[TIn num.ExtendedPrimitive, TOut num.Float](c std.RGBA[TIn]) std.RGBA[TOut] {
	r := normalize.To[TIn, TOut](c.Red())
	g := normalize.To[TIn, TOut](c.Green())
	b := normalize.To[TIn, TOut](c.Blue())
	a := normalize.To[TIn, TOut](c.Alpha())
	return std.RGBA[TOut]{}.Set(r, g, b, a)
}

// ReScale returns an RGBA[TOut] scaled up to [0, TIn.Max] from an input bounded in the fully closed interval [0.0, 1.0].
func ReScale[TIn num.Float, TOut num.ExtendedInteger](c std.RGBA[TIn]) std.RGBA[TOut] {
	r := normalize.From[TIn, TOut](c.Red())
	g := normalize.From[TIn, TOut](c.Green())
	b := normalize.From[TIn, TOut](c.Blue())
	a := normalize.From[TIn, TOut](c.Alpha())
	return std.RGBA[TOut]{}.Set(r, g, b, a)
}

// Comparator returns if the two RGBA values are equal in values.
func Comparator[T num.ExtendedPrimitive](a std.RGBA[T], b std.RGBA[T]) bool {
	return a.Red() == b.Red() && a.Green() == b.Green() && a.Blue() == b.Blue() && a.Alpha() == b.Alpha()
}

// Random returns a pseudo-random std.RGBA[T] of the provided type using math.Random[T].
//
// If requesting a floating point type, the resulting number will be bounded
// in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [0, n] - where n is the maximum value of
// the provided type.
func Random[T num.ExtendedPrimitive]() std.RGBA[T] {
	r := num.Random[T]()
	g := num.Random[T]()
	b := num.Random[T]()
	a := num.Random[T]()
	return std.RGBA[T]{}.Set(r, g, b, a)
}

// RandomUpTo returns a pseudo-random std.RGBA[T] of the provided type with each channel bounded within its provided closed interval of [0, max].
func RandomUpTo[T num.ExtendedPrimitive](rUpper T, gUpper T, bUpper T, aUpper T) std.RGBA[T] {
	r := num.RandomBounded[T](0, rUpper)
	g := num.RandomBounded[T](0, gUpper)
	b := num.RandomBounded[T](0, bUpper)
	a := num.RandomBounded[T](0, aUpper)
	return std.RGBA[T]{}.Set(r, g, b, a)
}

// RandomRange returns a pseudo-random std.RGBA[T] of the provided type bounded in the closed interval [min, max].
func RandomRange[T num.ExtendedPrimitive](minimum, maximum T) std.RGBA[T] {
	r := num.RandomBounded[T](minimum, maximum)
	g := num.RandomBounded[T](minimum, maximum)
	b := num.RandomBounded[T](minimum, maximum)
	a := num.RandomBounded[T](minimum, maximum)
	return std.RGBA[T]{}.Set(r, g, b, a)
}
