package std

import "github.com/ignite-laboratories/core/std/num"

// RGBAGeneric is the underlying structure of color operations.  It differs from RGBA in that it supports asymmetric
// channel bit widths (as it's far more common in the modern age to work with symmetric channel widths).  This provides
// accessibility to legacy color spaces while allowing the more common RGBA type to evolve from old paradigms, rather
// than dismissing their existence entirely. =)
type RGBAGeneric[TR num.ExtendedPrimitive, TG num.ExtendedPrimitive, TB num.ExtendedPrimitive, TA num.ExtendedPrimitive] struct {
	// R is the red channel.
	R TR

	// G is the green channel.
	G TG

	// B is the blue channel.
	B TB

	// A is the alpha channel.
	A TA
}
