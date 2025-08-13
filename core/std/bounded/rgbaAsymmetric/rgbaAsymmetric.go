// Package rgbaAsymmetric provides access to all std.RGBAAsymmetric procedures.
package rgbaAsymmetric

import (
	"github.com/ignite-laboratories/core/std/bounded"
	"github.com/ignite-laboratories/core/std/num"
)

// From creates a new instance of std.RGBAAsymmetric[TR, TG, TB, TA] with each direction bounded in the fully closed interval [0, T.max], where
// T.max is the implied maximum value of that color channel's type.
func From[TR num.Primitive, TG num.Primitive, TB num.Primitive, TA num.Primitive](r TR, g TG, b TB, a TA) bounded.RGBAAsymmetric[TR, TG, TB, TA] {
	return bounded.RGBAAsymmetric[TR, TG, TB, TA]{}.Set(r, g, b, a)
}

// Random returns a pseudo-random std.RGBAAsymmetric[TR, TG, TB, TA] of the provided type using math.Random[TR, TG, TB], with
// each color channel bounded in the fully closed interval [0, T.Max]
func Random[TR num.Primitive, TG num.Primitive, TB num.Primitive, TA num.Primitive]() bounded.RGBAAsymmetric[TR, TG, TB, TA] {
	return bounded.RGBAAsymmetric[TR, TG, TB, TA]{
		R: bounded.Random[TR](),
		G: bounded.Random[TG](),
		B: bounded.Random[TB](),
		A: bounded.Random[TA](),
	}
}

// ScaleToType normalizes the std.RGBAAsymmetric[TInR, TInG, TInB, TInA] directional components into unit vectors and then scales them to a new std.RGBAAsymmetric[TOutR, TOutG, TOutB, TOutA].
func ScaleToType[TInR num.Primitive, TInG num.Primitive, TInB num.Primitive, TInA num.Primitive, TOutR num.Primitive, TOutG num.Primitive, TOutB num.Primitive, TOutA num.Primitive](value bounded.RGBAAsymmetric[TInR, TInG, TInB, TInA]) bounded.RGBAAsymmetric[TOutR, TOutG, TOutB, TOutA] {
	return bounded.RGBAAsymmetric[TOutR, TOutG, TOutB, TOutA]{
		R: bounded.ScaleToType[TInR, TOutR](value.R),
		G: bounded.ScaleToType[TInG, TOutG](value.G),
		B: bounded.ScaleToType[TInB, TOutB](value.B),
		A: bounded.ScaleToType[TInA, TOutA](value.A),
	}
}
