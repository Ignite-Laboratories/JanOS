// Package rgba provides access to all std.RGBA procedures.
package rgba

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/bounded"
	"github.com/ignite-laboratories/core/std/num"
)

// From creates a new instance of std.RGBA[T] with each direction bounded in the fully closed interval [0, T.max], where
// T.max is the implied maximum value of that color channel's type.
func From[T num.ExtendedPrimitive](r, g, b, a T) std.RGBA[T] {
	return std.RGBA[T]{}.Set(r, g, b, a)
}

// FromHex treats each hexadecimal component as of the input value as a unit vector representing the typed color space,
// allowing you to create implicitly scaled colors from hexadecimal values.  For example -
//
//	rgba.FromHex[byte](0xAABBCC**):
//	    R: 170 [0xAA]
//	    G: 187 [0xBB]
//	    B: 204 [0xCC]
//
//	rgba.FromHex[num.Nibble](0xAABBCC**):
//	    R: 10 [0xAA]
//	    G: 11 [0xBB]
//	    B: 12 [0xCC]
func FromHex[T num.ExtendedPrimitive](value uint32) std.RGBA[T] {
	r := byte((value >> 24) & 0xFF)
	g := byte((value >> 16) & 0xFF)
	b := byte((value >> 8) & 0xFF)
	a := byte(value & 0xFF)
	return ScaleToType[byte, T](From[byte](r, g, b, a))
}

// Random returns a pseudo-random std.RGBA[T] of the provided type using math.Random[T], with
// each color channel bounded in the fully closed interval [0, T.Max]
func Random[T num.ExtendedPrimitive]() std.RGBA[T] {
	return std.RGBA[T]{
		R: bounded.Random[T](),
		G: bounded.Random[T](),
		B: bounded.Random[T](),
		A: bounded.Random[T](),
	}
}

// RandomUpTo returns a pseudo-random std.RGBA[T] of the provided type using math.Random[T], with
// each color channel bounded in the fully closed interval [0, maximum]
func RandomUpTo[T num.ExtendedPrimitive](maximum T) std.RGBA[T] {
	return std.RGBA[T]{
		R: bounded.RandomSubset[T](0, maximum),
		G: bounded.RandomSubset[T](0, maximum),
		B: bounded.RandomSubset[T](0, maximum),
		A: bounded.RandomSubset[T](0, maximum),
	}
}

// RandomRange returns a pseudo-random std.RGBA[T] of the provided type using math.Random[T], with
// each color channel bounded in the fully closed interval [minimum, maximum]
func RandomRange[T num.ExtendedPrimitive](minimum, maximum T) std.RGBA[T] {
	return std.RGBA[T]{
		R: bounded.RandomSubset[T](minimum, maximum),
		G: bounded.RandomSubset[T](minimum, maximum),
		B: bounded.RandomSubset[T](minimum, maximum),
		A: bounded.RandomSubset[T](minimum, maximum),
	}
}

// ScaleToType normalizes the std.RGBA[T] directional components into unit vectors and then scales them to a new std.RGBA[TOut].
func ScaleToType[TIn num.ExtendedPrimitive, TOut num.ExtendedPrimitive](value std.RGBA[TIn]) std.RGBA[TOut] {
	return std.RGBA[TOut]{
		R: bounded.ScaleToType[TIn, TOut](value.R),
		G: bounded.ScaleToType[TIn, TOut](value.G),
		B: bounded.ScaleToType[TIn, TOut](value.B),
		A: bounded.ScaleToType[TIn, TOut](value.A),
	}
}
