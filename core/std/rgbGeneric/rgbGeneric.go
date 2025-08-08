// Package rgbGeneric provides access to all std.RGBGeneric procedures.
package rgbGeneric

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/bounded"
	"github.com/ignite-laboratories/core/std/num"
)

// From creates a new instance of std.RGBGeneric[TR, TG, TB] with each direction bounded in the fully closed interval [0, T.max], where
// T.max is the implied maximum value of that color channel's type.
func From[TR num.ExtendedPrimitive, TG num.ExtendedPrimitive, TB num.ExtendedPrimitive](r TR, g TG, b TB) std.RGBGeneric[TR, TG, TB] {
	return std.RGBGeneric[TR, TG, TB]{}.Set(r, g, b)
}

// Random returns a pseudo-random std.RGBGeneric[TR, TG, TB] of the provided type using math.Random[TR, TG, TB], with
// each color channel bounded in the fully closed interval [0, T.Max]
func Random[TR num.ExtendedPrimitive, TG num.ExtendedPrimitive, TB num.ExtendedPrimitive]() std.RGBGeneric[TR, TG, TB] {
	return std.RGBGeneric[TR, TG, TB]{
		R: bounded.Random[TR](),
		G: bounded.Random[TG](),
		B: bounded.Random[TB](),
	}
}

// ScaleToType normalizes the std.RGBGeneric[TInR, TInG, TInB] directional components into unit vectors and then scales them to a new std.RGBGeneric[TOutR, TOutG, TOutB].
func ScaleToType[TInR num.ExtendedPrimitive, TInG num.ExtendedPrimitive, TInB num.ExtendedPrimitive, TOutR num.ExtendedPrimitive, TOutG num.ExtendedPrimitive, TOutB num.ExtendedPrimitive](value std.RGBGeneric[TInR, TInG, TInB]) std.RGBGeneric[TOutR, TOutG, TOutB] {
	return std.RGBGeneric[TOutR, TOutG, TOutB]{
		R: bounded.ScaleToType[TInR, TOutR](value.R),
		G: bounded.ScaleToType[TInG, TOutG](value.G),
		B: bounded.ScaleToType[TInB, TOutB](value.B),
	}
}
