package std

import "github.com/ignite-laboratories/core/std/num"

// RGBGeneric is a reduced variant of RGBAGeneric that does not have an alpha channel, but all channel operations are identical.
type RGBGeneric[TR num.ExtendedPrimitive, TG num.ExtendedPrimitive, TB num.ExtendedPrimitive] struct {
	// R is the red channel.
	R TR

	// G is the green channel.
	G TG

	// B is the blue channel.
	B TB
}
