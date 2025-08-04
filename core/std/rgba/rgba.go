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
	return FromHex[TInt](
		(uint32(r) << 24) |
			(uint32(g) << 16) |
			(uint32(b) << 8) |
			uint32(a))
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
	overflow := uint64(0)
	var zero TInt
	switch any(zero).(type) {
	case num.Crumb, num.Note, num.Nibble, num.Flake, num.Morsel, num.Shred, num.Run, num.Scale, num.Riff, num.Hook:
		overflow = std.MaxValue[TInt]() + 1
	case float32, float64:
		panic("floating point types are reserved for normalized values, please create a std.RGBA using an integer")
	}

	if overflow > 0 {
		return std.RGBA[TInt]{
			R: TInt(uint64((value>>24)&0xFF) % overflow),
			G: TInt(uint64((value>>16)&0xFF) % overflow),
			B: TInt(uint64((value>>8)&0xFF) % overflow),
			A: TInt(uint64(value&0xFF) % overflow),
		}
	}
	return std.RGBA[TInt]{
		R: TInt((value >> 24) & 0xFF),
		G: TInt((value >> 16) & 0xFF),
		B: TInt((value >> 8) & 0xFF),
		A: TInt(value & 0xFF),
	}
}

// Normalize returns an RGBA[TOut] ranging from 0.0-1.0.
func Normalize[TIn num.ExtendedPrimitive, TOut num.Float](c std.RGBA[TIn]) std.RGBA[TOut] {
	return std.RGBA[TOut]{
		R: normalize.To[TIn, TOut](c.R),
		G: normalize.To[TIn, TOut](c.G),
		B: normalize.To[TIn, TOut](c.B),
		A: normalize.To[TIn, TOut](c.A),
	}
}

// ReScale returns an RGBA[TOut] scaled up to [0, TIn.Max] from an input bounded in the fully closed interval [0.0, 1.0].
func ReScale[TIn num.Float, TOut num.ExtendedInteger](c std.RGBA[TIn]) std.RGBA[TOut] {
	return std.RGBA[TOut]{
		R: normalize.From[TIn, TOut](c.R),
		G: normalize.From[TIn, TOut](c.G),
		B: normalize.From[TIn, TOut](c.B),
		A: normalize.From[TIn, TOut](c.A),
	}
}

// Comparator returns if the two RGBA values are equal in values.
func Comparator[T num.ExtendedPrimitive](a std.RGBA[T], b std.RGBA[T]) bool {
	return a.R == b.R && a.G == b.G && a.B == b.B && a.A == b.A
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
	return std.RGBA[T]{
		R: std.Random[T](),
		G: std.Random[T](),
		B: std.Random[T](),
		A: std.Random[T](),
	}
}

// RandomUpTo returns a pseudo-random std.RGBA[T] of the provided type with each channel bounded within its provided closed interval of [0, max].
func RandomUpTo[T num.ExtendedPrimitive](rUpper T, gUpper T, bUpper T, aUpper T) std.RGBA[T] {
	return std.RGBA[T]{
		R: std.RandomBounded[T](0, rUpper),
		G: std.RandomBounded[T](0, gUpper),
		B: std.RandomBounded[T](0, bUpper),
		A: std.RandomBounded[T](0, aUpper),
	}
}

// RandomRange returns a pseudo-random std.RGBA[T] of the provided type bounded in the closed interval [min, max].
func RandomRange[T num.ExtendedPrimitive](minimum, maximum T) std.RGBA[T] {
	return std.RGBA[T]{
		R: std.RandomBounded[T](minimum, maximum),
		G: std.RandomBounded[T](minimum, maximum),
		B: std.RandomBounded[T](minimum, maximum),
		A: std.RandomBounded[T](minimum, maximum),
	}
}
