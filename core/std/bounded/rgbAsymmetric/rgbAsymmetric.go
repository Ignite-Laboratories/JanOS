// Package rgbAsymmetric provides access to all std.RGBAsymmetric procedures.
package rgbAsymmetric

import (
	"github.com/ignite-laboratories/core/std/bounded"
	"github.com/ignite-laboratories/core/std/num"
)

// From creates a new instance of std.RGBAsymmetric[TR, TG, TB] with each direction bounded in the fully closed interval [0, T.max], where
// T.max is the implied maximum value of that color channel's type.
func From[TR num.Primitive, TG num.Primitive, TB num.Primitive](r TR, g TG, b TB) bounded.RGBAsymmetric[TR, TG, TB] {
	return bounded.RGBAsymmetric[TR, TG, TB]{}.Set(r, g, b)
}

// Random returns a pseudo-random std.RGBAsymmetric[TR, TG, TB] of the provided type using math.Random[TR, TG, TB], with
// each color channel bounded in the fully closed interval [0, T.Max]
func Random[TR num.Primitive, TG num.Primitive, TB num.Primitive]() bounded.RGBAsymmetric[TR, TG, TB] {
	return bounded.RGBAsymmetric[TR, TG, TB]{
		R: bounded.Random[TR](),
		G: bounded.Random[TG](),
		B: bounded.Random[TB](),
	}
}

// ScaleToType normalizes the std.RGBAsymmetric[TInR, TInG, TInB] directional components into unit vectors and then scales them to a new std.RGBAsymmetric[TOutR, TOutG, TOutB].
func ScaleToType[TInR num.Primitive, TInG num.Primitive, TInB num.Primitive, TOutR num.Primitive, TOutG num.Primitive, TOutB num.Primitive](value bounded.RGBAsymmetric[TInR, TInG, TInB]) bounded.RGBAsymmetric[TOutR, TOutG, TOutB] {
	return bounded.RGBAsymmetric[TOutR, TOutG, TOutB]{
		R: bounded.ScaleToType[TInR, TOutR](value.R),
		G: bounded.ScaleToType[TInG, TOutG](value.G),
		B: bounded.ScaleToType[TInB, TOutB](value.B),
	}
}
