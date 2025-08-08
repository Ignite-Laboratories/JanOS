// Package rgbaGeneric provides access to all std.RGBAGeneric procedures.
package rgbaGeneric

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/bounded"
	"github.com/ignite-laboratories/core/std/num"
)

// From creates a new instance of std.RGBAGeneric[TR, TG, TB, TA] with each direction bounded in the fully closed interval [0, T.max], where
// T.max is the implied maximum value of that color channel's type.
func From[TR num.ExtendedPrimitive, TG num.ExtendedPrimitive, TB num.ExtendedPrimitive, TA num.ExtendedPrimitive](r TR, g TG, b TB, a TA) std.RGBAGeneric[TR, TG, TB, TA] {
	return std.RGBAGeneric[TR, TG, TB, TA]{}.Set(r, g, b, a)
}

// Random returns a pseudo-random std.RGBAGeneric[TR, TG, TB, TA] of the provided type using math.Random[TR, TG, TB], with
// each color channel bounded in the fully closed interval [0, T.Max]
func Random[TR num.ExtendedPrimitive, TG num.ExtendedPrimitive, TB num.ExtendedPrimitive, TA num.ExtendedPrimitive]() std.RGBAGeneric[TR, TG, TB, TA] {
	return std.RGBAGeneric[TR, TG, TB, TA]{
		R: bounded.Random[TR](),
		G: bounded.Random[TG](),
		B: bounded.Random[TB](),
		A: bounded.Random[TA](),
	}
}

// ScaleToType normalizes the std.RGBAGeneric[TInR, TInG, TInB, TInA] directional components into unit vectors and then scales them to a new std.RGBAGeneric[TOutR, TOutG, TOutB, TOutA].
func ScaleToType[TInR num.ExtendedPrimitive, TInG num.ExtendedPrimitive, TInB num.ExtendedPrimitive, TInA num.ExtendedPrimitive, TOutR num.ExtendedPrimitive, TOutG num.ExtendedPrimitive, TOutB num.ExtendedPrimitive, TOutA num.ExtendedPrimitive](value std.RGBAGeneric[TInR, TInG, TInB, TInA]) std.RGBAGeneric[TOutR, TOutG, TOutB, TOutA] {
	return std.RGBAGeneric[TOutR, TOutG, TOutB, TOutA]{
		R: bounded.ScaleToType[TInR, TOutR](value.R),
		G: bounded.ScaleToType[TInG, TOutG](value.G),
		B: bounded.ScaleToType[TInB, TOutB](value.B),
		A: bounded.ScaleToType[TInA, TOutA](value.A),
	}
}
