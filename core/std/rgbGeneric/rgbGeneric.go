// Package rgbGeneric provides access to all std.RGBGeneric procedures.
package rgbGeneric

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

// From constructs a std.RGBGeneric[TR, TG, TB] from individual red, green, blue, and alpha components of their own type.
//
// NOTE: If you provide a sub-byte size, each channel's value will be modulo-d against 2ⁿ, with 𝑛 being the sub-byte bit-width.
func From[TR num.ExtendedInteger, TG num.ExtendedInteger, TB num.ExtendedInteger](r TR, g TG, b TB) std.RGBGeneric[TR, TG, TB] {
	return std.RGBGeneric[TR, TG, TB]{
		R: overflowValue(r),
		G: overflowValue(g),
		B: overflowValue(b),
	}
}

// Normalize returns a std.RGBGeneric[TOut] ranging from 0.0-1.0.
func Normalize[TR num.ExtendedInteger, TG num.ExtendedInteger, TB num.ExtendedInteger, TOut num.Float](c std.RGBGeneric[TR, TG, TB]) std.RGBGeneric[TOut, TOut, TOut] {
	return std.RGBGeneric[TOut, TOut, TOut]{
		R: normalize.To[TR, TOut](c.R),
		G: normalize.To[TG, TOut](c.G),
		B: normalize.To[TB, TOut](c.B),
	}
}

// ReScale returns a std.RGBGeneric[TR, TG, TB] scaled up to [0, TChan.Max] from an input bounded in the fully closed interval [0.0, 1.0].
func ReScale[TIn num.Float, TR num.ExtendedInteger, TG num.ExtendedInteger, TB num.ExtendedInteger](c std.RGBGeneric[TIn, TIn, TIn]) std.RGBGeneric[TR, TG, TB] {
	return std.RGBGeneric[TR, TG, TB]{
		R: normalize.From[TIn, TR](c.R),
		G: normalize.From[TIn, TG](c.G),
		B: normalize.From[TIn, TB](c.B),
	}
}

// Comparator returns if the two RGB values are equal in values.
func Comparator[TR num.ExtendedInteger, TG num.ExtendedInteger, TB num.ExtendedInteger](a std.RGBGeneric[TR, TG, TB], b std.RGBGeneric[TR, TG, TB]) bool {
	return a.R == b.R && a.G == b.G && a.B == b.B
}

// Random returns a pseudo-random std.RGBGeneric[TR, TG, TB] of the provided type using math.Random[TChan].
//
// NOTE: If requesting a floating point type, the resulting number will be bounded in the fully closed interval [0.0, 1.0]
//
// NOTE: If requesting an integer type, the resulting number will be bounded in the fully closed interval [0, TChan.Max]
func Random[TR num.ExtendedInteger, TG num.ExtendedInteger, TB num.ExtendedInteger]() std.RGBGeneric[TR, TG, TB] {
	return std.RGBGeneric[TR, TG, TB]{
		R: std.Random[TR](),
		G: std.Random[TG](),
		B: std.Random[TB](),
	}
}
