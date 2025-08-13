// Package rgba provides access to all std.RGBA procedures.
package rgba

import (
	"github.com/ignite-laboratories/core/std/bounded"
	"github.com/ignite-laboratories/core/std/num"
)

// From creates a new instance of std.RGBA[T] with each direction bounded in the fully closed interval [0, T.max], where
// T.max is the implied maximum value of that color channel's type.
func From[T num.Primitive](r, g, b, a T) bounded.RGBA[T] {
	return bounded.RGBA[T]{}.Set(r, g, b, a)
}

// FromHex treats each hexadecimal component of the input value as a unit vector representing the typed color space,
// allowing you to implicitly scale up to a color space from a hexadecimal value.  For example -
//
//	rgba.FromHex[byte](0xAABBCC**):
//	    R: 170 [0xAA] (255 * 0.66666)
//	    G: 187 [0xBB] (255 * 0.73333)
//	    B: 204 [0xCC] (255 * 0.8)
//
//	rgba.FromHex[num.Nibble](0xAABBCC**):
//	    R: 10 [0xAA] (15 * 0.66666)
//	    G: 11 [0xBB] (15 * 0.73333)
//	    B: 12 [0xCC] (15 * 0.8)
func FromHex[T num.Primitive](value uint32) bounded.RGBA[T] {
	r := byte((value >> 24) & 0xFF)
	g := byte((value >> 16) & 0xFF)
	b := byte((value >> 8) & 0xFF)
	a := byte(value & 0xFF)
	return ScaleToType[byte, T](From[byte](r, g, b, a))
}

// Random returns a pseudo-random std.RGBA[T] of the provided type using math.Random[T], with
// each color channel bounded in the fully closed interval [0, T.Max]
func Random[T num.Primitive]() bounded.RGBA[T] {
	return bounded.RGBA[T]{
		R: bounded.Random[T](),
		G: bounded.Random[T](),
		B: bounded.Random[T](),
		A: bounded.Random[T](),
	}
}

// RandomUpTo returns a pseudo-random std.RGBA[T] of the provided type using math.Random[T], with
// each color channel bounded in the fully closed interval [0, maximum]
func RandomUpTo[T num.Primitive](maximum T) bounded.RGBA[T] {
	return bounded.RGBA[T]{
		R: bounded.RandomSubset[T](0, maximum),
		G: bounded.RandomSubset[T](0, maximum),
		B: bounded.RandomSubset[T](0, maximum),
		A: bounded.RandomSubset[T](0, maximum),
	}
}

// RandomRange returns a pseudo-random std.RGBA[T] of the provided type using math.Random[T], with
// each color channel bounded in the fully closed interval [minimum, maximum]
func RandomRange[T num.Primitive](minimum, maximum T) bounded.RGBA[T] {
	return bounded.RGBA[T]{
		R: bounded.RandomSubset[T](minimum, maximum),
		G: bounded.RandomSubset[T](minimum, maximum),
		B: bounded.RandomSubset[T](minimum, maximum),
		A: bounded.RandomSubset[T](minimum, maximum),
	}
}

// ScaleToType normalizes the std.RGBA[T] directional components into unit vectors and then scales them to a new std.RGBA[TOut].
func ScaleToType[TIn num.Primitive, TOut num.Primitive](value bounded.RGBA[TIn]) bounded.RGBA[TOut] {
	return bounded.RGBA[TOut]{
		R: bounded.ScaleToType[TIn, TOut](value.R),
		G: bounded.ScaleToType[TIn, TOut](value.G),
		B: bounded.ScaleToType[TIn, TOut](value.B),
		A: bounded.ScaleToType[TIn, TOut](value.A),
	}
}
