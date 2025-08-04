// Package rgbGeneric provides access to all std.RGBGeneric procedures.
package rgbGeneric

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/normalize"
	"github.com/ignite-laboratories/core/std/num"
)

// From constructs a std.RGBGeneric[TR, TG, TB] from individual red, green, blue, and alpha components of their own type.
//
// NOTE: If you provide a sub-byte size, each channel's value will be modulo-d against 2ⁿ, with 𝑛 being the sub-byte bit-width.
func From[TR num.ExtendedInteger, TG num.ExtendedInteger, TB num.ExtendedInteger](r TR, g TG, b TB) std.RGBGeneric[TR, TG, TB] {
	return std.RGBGeneric[TR, TG, TB]{}.Set(r, g, b)
}

// Normalize returns a std.RGBGeneric[TOut] ranging from 0.0-1.0.
func Normalize[TR num.ExtendedInteger, TG num.ExtendedInteger, TB num.ExtendedInteger, TOut num.Float](c std.RGBGeneric[TR, TG, TB]) std.RGBGeneric[TOut, TOut, TOut] {
	return std.RGBGeneric[TOut, TOut, TOut]{}.SetRed(normalize.To[TR, TOut](c.Red())).SetGreen(normalize.To[TG, TOut](c.Green())).SetBlue(normalize.To[TB, TOut](c.Blue()))
}

// ReScale returns a std.RGBGeneric[TR, TG, TB] scaled up to [0, TChan.Max] from an input bounded in the fully closed interval [0.0, 1.0].
func ReScale[TIn num.Float, TR num.ExtendedInteger, TG num.ExtendedInteger, TB num.ExtendedInteger](c std.RGBGeneric[TIn, TIn, TIn]) std.RGBGeneric[TR, TG, TB] {
	return std.RGBGeneric[TR, TG, TB]{}.SetRed(normalize.From[TIn, TR](c.Red())).SetGreen(normalize.From[TIn, TG](c.Green())).SetBlue(normalize.From[TIn, TB](c.Blue()))
}

// Comparator returns if the two RGB values are equal in values.
func Comparator[TR num.ExtendedInteger, TG num.ExtendedInteger, TB num.ExtendedInteger](a std.RGBGeneric[TR, TG, TB], b std.RGBGeneric[TR, TG, TB]) bool {
	return a.Red() == b.Red() && a.Green() == b.Green() && a.Blue() == b.Blue()
}

// Random returns a pseudo-random std.RGBGeneric[TR, TG, TB] of the provided type using math.Random[TChan].
//
// NOTE: If requesting a floating point type, the resulting number will be bounded in the fully closed interval [0.0, 1.0]
//
// NOTE: If requesting an integer type, the resulting number will be bounded in the fully closed interval [0, TChan.Max]
func Random[TR num.ExtendedInteger, TG num.ExtendedInteger, TB num.ExtendedInteger]() std.RGBGeneric[TR, TG, TB] {
	return std.RGBGeneric[TR, TG, TB]{}.SetRed(std.Random[TR]()).SetGreen(std.Random[TG]()).SetBlue(std.Random[TB]())
}
