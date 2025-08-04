// Package rgbaGeneric provides access to all std.RGBAGeneric procedures.
package rgbaGeneric

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/normalize"
	"github.com/ignite-laboratories/core/std/num"
)

// From constructs a std.RGBAGeneric[TR, TG, TB, TA] from the individually typed red, green, blue, and alpha values.
//
// NOTE: If you provide a sub-byte size, each channel's value will be modulo-d against 2ⁿ, with 𝑛 being the sub-byte bit-width.
func From[TR num.ExtendedInteger, TG num.ExtendedInteger, TB num.ExtendedInteger, TA num.ExtendedInteger](r TR, g TG, b TB, a TA) std.RGBAGeneric[TR, TG, TB, TA] {
	return std.RGBAGeneric[TR, TG, TB, TA]{}.Set(r, g, b, a)
}

// Normalize returns a std.RGBAGeneric[TOut] ranging from 0.0-1.0.
func Normalize[TR num.ExtendedInteger, TG num.ExtendedInteger, TB num.ExtendedInteger, TA num.ExtendedInteger, TOut num.Float](c std.RGBAGeneric[TR, TG, TB, TA]) std.RGBAGeneric[TOut, TOut, TOut, TOut] {
	return std.RGBAGeneric[TOut, TOut, TOut, TOut]{}.SetRed(normalize.To[TR, TOut](c.Red())).SetGreen(normalize.To[TG, TOut](c.Green())).SetBlue(normalize.To[TB, TOut](c.Blue())).SetAlpha(normalize.To[TA, TOut](c.Alpha()))
}

// ReScale returns a std.RGBAGeneric[TR, TG, TB, TA] scaled up to [0, TChan.Max] from an input bounded in the fully closed interval [0.0, 1.0].
func ReScale[TIn num.Float, TR num.ExtendedInteger, TG num.ExtendedInteger, TB num.ExtendedInteger, TA num.ExtendedInteger](c std.RGBAGeneric[TIn, TIn, TIn, TIn]) std.RGBAGeneric[TR, TG, TB, TA] {
	return std.RGBAGeneric[TR, TG, TB, TA]{}.SetRed(normalize.From[TIn, TR](c.Red())).SetGreen(normalize.From[TIn, TG](c.Green())).SetBlue(normalize.From[TIn, TB](c.Blue())).SetAlpha(normalize.From[TIn, TA](c.Alpha()))
}

// Comparator returns if the two RGBA values are equal in values.
func Comparator[TR num.ExtendedInteger, TG num.ExtendedInteger, TB num.ExtendedInteger, TA num.ExtendedInteger](a std.RGBAGeneric[TR, TG, TB, TA], b std.RGBAGeneric[TR, TG, TB, TA]) bool {
	return a.Red() == b.Red() && a.Green() == b.Green() && a.Blue() == b.Blue() && a.Alpha() == b.Alpha()
}

// Random returns a pseudo-random std.RGBAGeneric[TR, TG, TB, TA] of the provided type using math.Random[TChan].
//
// NOTE: If requesting a floating point type, the resulting number will be bounded in the fully closed interval [0.0, 1.0]
//
// NOTE: If requesting an integer type, the resulting number will be bounded in the fully closed interval [0, TChan.Max]
func Random[TR num.ExtendedInteger, TG num.ExtendedInteger, TB num.ExtendedInteger, TA num.ExtendedInteger]() std.RGBAGeneric[TR, TG, TB, TA] {
	return std.RGBAGeneric[TR, TG, TB, TA]{}.SetRed(std.Random[TR]()).SetGreen(std.Random[TG]()).SetBlue(std.Random[TB]()).SetAlpha(std.Random[TA]())
}
