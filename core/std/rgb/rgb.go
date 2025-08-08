// Package rgb provides access to all std.RGB procedures.
package rgb

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/bounded"
	"github.com/ignite-laboratories/core/std/num"
)

// From creates a new instance of std.RGB[T] with each direction bounded in the fully closed interval [0, T.max], where
// T.max is the implied maximum value of that color channel's type.
func From[T num.ExtendedPrimitive](r, g, b T) std.RGB[T] {
	return std.RGB[T]{}.Set(r, g, b)
}

// FromHex treats each hexadecimal component as of the input value as a unit vector representing the typed color space,
// allowing you to create implicitly scaled colors from hexadecimal values.  For example -
//
//	rgb.FromHex[byte](0xAABBCC**):
//	    R: 170 [0xAA]
//	    G: 187 [0xBB]
//	    B: 204 [0xCC]
//
//	rgbb.FromHex[num.Nibble](0xAABBCC**):
//	    R: 10 [0xAA]
//	    G: 11 [0xBB]
//	    B: 12 [0xCC]
func FromHex[T num.ExtendedPrimitive](value uint32) std.RGB[T] {
	r := byte((value >> 24) & 0xFF)
	g := byte((value >> 16) & 0xFF)
	b := byte((value >> 8) & 0xFF)
	return ScaleToType[byte, T](From[byte](r, g, b))
}

// Random returns a pseudo-random std.RGB[T] of the provided type using math.Random[T], with
// each color channel bounded in the fully closed interval [0, T.Max]
func Random[T num.ExtendedPrimitive]() std.RGB[T] {
	return std.RGB[T]{
		R: bounded.Random[T](),
		G: bounded.Random[T](),
		B: bounded.Random[T](),
	}
}

// RandomUpTo returns a pseudo-random std.RGB[T] of the provided type using math.Random[T], with
// each color channel bounded in the fully closed interval [0, maximum]
func RandomUpTo[T num.ExtendedPrimitive](maximum T) std.RGB[T] {
	return std.RGB[T]{
		R: bounded.RandomSubset[T](0, maximum),
		G: bounded.RandomSubset[T](0, maximum),
		B: bounded.RandomSubset[T](0, maximum),
	}
}

// RandomRange returns a pseudo-random std.RGB[T] of the provided type using math.Random[T], with
// each color channel bounded in the fully closed interval [minimum, maximum]
func RandomRange[T num.ExtendedPrimitive](minimum, maximum T) std.RGB[T] {
	return std.RGB[T]{
		R: bounded.RandomSubset[T](minimum, maximum),
		G: bounded.RandomSubset[T](minimum, maximum),
		B: bounded.RandomSubset[T](minimum, maximum),
	}
}

// ScaleToType normalizes the std.RGB[T] directional components into unit vectors and then scales them to a new std.RGB[TOut].
func ScaleToType[TIn num.ExtendedPrimitive, TOut num.ExtendedPrimitive](value std.RGB[TIn]) std.RGB[TOut] {
	return std.RGB[TOut]{
		R: bounded.ScaleToType[TIn, TOut](value.R),
		G: bounded.ScaleToType[TIn, TOut](value.G),
		B: bounded.ScaleToType[TIn, TOut](value.B),
	}
}
