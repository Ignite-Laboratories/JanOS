// Package rgb provides access to all std.RGB procedures.
package rgb

import (
	"github.com/ignite-laboratories/core/std/bounded"
	"github.com/ignite-laboratories/core/std/num"
)

// From creates a new instance of std.RGB[T] with each direction bounded in the fully closed interval [0, T.max], where
// T.max is the implied maximum value of that color channel's type.
func From[T num.Primitive](r, g, b T) bounded.RGB[T] {
	return bounded.RGB[T]{}.Set(r, g, b)
}

// FromHex treats each hexadecimal component of the input value as a unit vector representing the typed color space,
// allowing you to implicitly scale up to a color space from a hexadecimal value.  For example -
//
//	rgba.FromHex[byte](0xAABBCCDD):
//	    R: 170 [0xAA] (255 * 0.66666)
//	    G: 187 [0xBB] (255 * 0.73333)
//	    B: 204 [0xCC] (255 * 0.8)
//	    A: 221 [0xCC] (255 * 0.86666)
//
//	rgba.FromHex[num.Nibble](0xAABBCCDD):
//	    R: 10 [0xAA] (15 * 0.66666)
//	    G: 11 [0xBB] (15 * 0.73333)
//	    B: 12 [0xCC] (15 * 0.8)
//	    A: 13 [0xDD] (15 * 0.86666)
func FromHex[T num.Primitive](value uint32) bounded.RGB[T] {
	r := byte((value >> 24) & 0xFF)
	g := byte((value >> 16) & 0xFF)
	b := byte((value >> 8) & 0xFF)
	return ScaleToType[byte, T](From[byte](r, g, b))
}

// Random returns a pseudo-random std.RGB[T] of the provided type using math.Random[T], with
// each color channel bounded in the fully closed interval [0, T.Max]
func Random[T num.Primitive]() bounded.RGB[T] {
	return bounded.RGB[T]{
		R: bounded.Random[T](),
		G: bounded.Random[T](),
		B: bounded.Random[T](),
	}
}

// RandomUpTo returns a pseudo-random std.RGB[T] of the provided type using math.Random[T], with
// each color channel bounded in the fully closed interval [0, maximum]
func RandomUpTo[T num.Primitive](maximum T) bounded.RGB[T] {
	return bounded.RGB[T]{
		R: bounded.RandomSubset[T](0, maximum),
		G: bounded.RandomSubset[T](0, maximum),
		B: bounded.RandomSubset[T](0, maximum),
	}
}

// RandomRange returns a pseudo-random std.RGB[T] of the provided type using math.Random[T], with
// each color channel bounded in the fully closed interval [minimum, maximum]
func RandomRange[T num.Primitive](minimum, maximum T) bounded.RGB[T] {
	return bounded.RGB[T]{
		R: bounded.RandomSubset[T](minimum, maximum),
		G: bounded.RandomSubset[T](minimum, maximum),
		B: bounded.RandomSubset[T](minimum, maximum),
	}
}

// ScaleToType normalizes the std.RGB[T] directional components into unit vectors and then scales them to a new std.RGB[TOut].
func ScaleToType[TIn num.Primitive, TOut num.Primitive](value bounded.RGB[TIn]) bounded.RGB[TOut] {
	return bounded.RGB[TOut]{
		R: bounded.ScaleToType[TIn, TOut](value.R),
		G: bounded.ScaleToType[TIn, TOut](value.G),
		B: bounded.ScaleToType[TIn, TOut](value.B),
	}
}
