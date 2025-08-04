// Package rgbaGeneric provides access to all std.RGBAGeneric procedures.
package rgbaGeneric

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/normalize"
	"github.com/ignite-laboratories/core/std/num"
)

func overflowValue[T num.ExtendedInteger](value T) T {
	var zero T
	switch any(zero).(type) {
	case num.Crumb, num.Note, num.Nibble, num.Flake, num.Morsel, num.Shred, num.Run, num.Scale, num.Riff, num.Hook:
		overflow := std.MaxValue[T]() + 1
		return T(int(value) % int(overflow))
	case float32, float64:
		panic("floating point types are reserved for normalized values, please create a std.RGBA using an integer")
	}
	return value
}

// From constructs a std.RGBAGeneric[TR, TG, TB, TA] from the individually typed red, green, blue, and alpha values.
//
// NOTE: If you provide a sub-byte size, each channel's value will be modulo-d against 2ⁿ, with 𝑛 being the sub-byte bit-width.
func From[TR num.ExtendedInteger, TG num.ExtendedInteger, TB num.ExtendedInteger, TA num.ExtendedInteger](r TR, g TG, b TB, a TA) std.RGBAGeneric[TR, TG, TB, TA] {
	return std.RGBAGeneric[TR, TG, TB, TA]{
		R: overflowValue(r),
		G: overflowValue(g),
		B: overflowValue(b),
		A: overflowValue(a),
	}
}

// Normalize returns a std.RGBAGeneric[TOut] ranging from 0.0-1.0.
func Normalize[TR num.ExtendedInteger, TG num.ExtendedInteger, TB num.ExtendedInteger, TA num.ExtendedInteger, TOut num.Float](c std.RGBAGeneric[TR, TG, TB, TA]) std.RGBAGeneric[TOut, TOut, TOut, TOut] {
	return std.RGBAGeneric[TOut, TOut, TOut, TOut]{
		R: normalize.To[TR, TOut](c.R),
		G: normalize.To[TG, TOut](c.G),
		B: normalize.To[TB, TOut](c.B),
		A: normalize.To[TA, TOut](c.A),
	}
}

// ReScale returns a std.RGBAGeneric[TR, TG, TB, TA] scaled up to [0, TChan.Max] from an input bounded in the fully closed interval [0.0, 1.0].
func ReScale[TIn num.Float, TR num.ExtendedInteger, TG num.ExtendedInteger, TB num.ExtendedInteger, TA num.ExtendedInteger](c std.RGBAGeneric[TIn, TIn, TIn, TIn]) std.RGBAGeneric[TR, TG, TB, TA] {
	return std.RGBAGeneric[TR, TG, TB, TA]{
		R: normalize.From[TIn, TR](c.R),
		G: normalize.From[TIn, TG](c.G),
		B: normalize.From[TIn, TB](c.B),
		A: normalize.From[TIn, TA](c.A),
	}
}

// Comparator returns if the two RGBA values are equal in values.
func Comparator[TR num.ExtendedInteger, TG num.ExtendedInteger, TB num.ExtendedInteger, TA num.ExtendedInteger](a std.RGBAGeneric[TR, TG, TB, TA], b std.RGBAGeneric[TR, TG, TB, TA]) bool {
	return a.R == b.R && a.G == b.G && a.B == b.B && a.A == b.A
}

// Random returns a pseudo-random std.RGBAGeneric[TR, TG, TB, TA] of the provided type using math.Random[TChan].
//
// NOTE: If requesting a floating point type, the resulting number will be bounded in the fully closed interval [0.0, 1.0]
//
// NOTE: If requesting an integer type, the resulting number will be bounded in the fully closed interval [0, TChan.Max]
func Random[TR num.ExtendedInteger, TG num.ExtendedInteger, TB num.ExtendedInteger, TA num.ExtendedInteger]() std.RGBAGeneric[TR, TG, TB, TA] {
	return std.RGBAGeneric[TR, TG, TB, TA]{
		R: std.Random[TR](),
		G: std.Random[TG](),
		B: std.Random[TB](),
		A: std.Random[TA](),
	}
}
