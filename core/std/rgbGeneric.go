package std

import "github.com/ignite-laboratories/core/std/num"

// RGBGeneric is a reduced variant of RGBGeneric that does not have an alpha channel, but all channel operations are identical.
//
// NOTE: This type also provides rudimentary "swizzling."
type RGBGeneric[TR num.ExtendedPrimitive, TG num.ExtendedPrimitive, TB num.ExtendedPrimitive] struct {
	// R is the red channel.
	R TR

	// G is the green channel.
	G TG

	// B is the blue channel.
	B TB
}

/**
Swizzling
*/

func (c RGBGeneric[TR, TG, TB]) RR() (TR, TR) { return c.R, c.R }
func (c RGBGeneric[TR, TG, TB]) RG() (TR, TG) { return c.R, c.G }
func (c RGBGeneric[TR, TG, TB]) RB() (TR, TB) { return c.R, c.B }
func (c RGBGeneric[TR, TG, TB]) RA() (TR, TA) { return c.R, c.A }
func (c RGBGeneric[TR, TG, TB]) GR() (TG, TR) { return c.G, c.R }
func (c RGBGeneric[TR, TG, TB]) GG() (TG, TG) { return c.G, c.G }
func (c RGBGeneric[TR, TG, TB]) GB() (TG, TB) { return c.G, c.B }
func (c RGBGeneric[TR, TG, TB]) GA() (TG, TA) { return c.G, c.A }
func (c RGBGeneric[TR, TG, TB]) BR() (TB, TR) { return c.B, c.R }
func (c RGBGeneric[TR, TG, TB]) BG() (TB, TG) { return c.B, c.G }
func (c RGBGeneric[TR, TG, TB]) BB() (TB, TB) { return c.B, c.B }
func (c RGBGeneric[TR, TG, TB]) BA() (TB, TA) { return c.B, c.A }
func (c RGBGeneric[TR, TG, TB]) AR() (TA, TR) { return c.A, c.R }
func (c RGBGeneric[TR, TG, TB]) AG() (TA, TG) { return c.A, c.G }
func (c RGBGeneric[TR, TG, TB]) AB() (TA, TB) { return c.A, c.B }
func (c RGBGeneric[TR, TG, TB]) AA() (TA, TA) { return c.A, c.A }

func (c RGBGeneric[TR, TG, TB]) RRR() (TR, TR, TR) { return c.R, c.R, c.R }
func (c RGBGeneric[TR, TG, TB]) RRG() (TR, TR, TG) { return c.R, c.R, c.G }
func (c RGBGeneric[TR, TG, TB]) RRB() (TR, TR, TB) { return c.R, c.R, c.B }
func (c RGBGeneric[TR, TG, TB]) RRA() (TR, TR, TA) { return c.R, c.R, c.A }
func (c RGBGeneric[TR, TG, TB]) RGR() (TR, TG, TR) { return c.R, c.G, c.R }
func (c RGBGeneric[TR, TG, TB]) RGG() (TR, TG, TG) { return c.R, c.G, c.G }
func (c RGBGeneric[TR, TG, TB]) RGB() (TR, TG, TB) { return c.R, c.G, c.B }
func (c RGBGeneric[TR, TG, TB]) RGA() (TR, TG, TA) { return c.R, c.G, c.A }
func (c RGBGeneric[TR, TG, TB]) RBR() (TR, TB, TR) { return c.R, c.B, c.R }
func (c RGBGeneric[TR, TG, TB]) RBG() (TR, TB, TG) { return c.R, c.B, c.G }
func (c RGBGeneric[TR, TG, TB]) RBB() (TR, TB, TB) { return c.R, c.B, c.B }
func (c RGBGeneric[TR, TG, TB]) RBA() (TR, TB, TA) { return c.R, c.B, c.A }
func (c RGBGeneric[TR, TG, TB]) RAR() (TR, TA, TR) { return c.R, c.A, c.R }
func (c RGBGeneric[TR, TG, TB]) RAG() (TR, TA, TG) { return c.R, c.A, c.G }
func (c RGBGeneric[TR, TG, TB]) RAB() (TR, TA, TB) { return c.R, c.A, c.B }
func (c RGBGeneric[TR, TG, TB]) RAA() (TR, TA, TA) { return c.R, c.A, c.A }
func (c RGBGeneric[TR, TG, TB]) GRR() (TG, TR, TR) { return c.G, c.R, c.R }
func (c RGBGeneric[TR, TG, TB]) GRG() (TG, TR, TG) { return c.G, c.R, c.G }
func (c RGBGeneric[TR, TG, TB]) GRB() (TG, TR, TB) { return c.G, c.R, c.B }
func (c RGBGeneric[TR, TG, TB]) GRA() (TG, TR, TA) { return c.G, c.R, c.A }
func (c RGBGeneric[TR, TG, TB]) GGR() (TG, TG, TR) { return c.G, c.G, c.R }
func (c RGBGeneric[TR, TG, TB]) GGG() (TG, TG, TG) { return c.G, c.G, c.G }
func (c RGBGeneric[TR, TG, TB]) GGB() (TG, TG, TB) { return c.G, c.G, c.B }
func (c RGBGeneric[TR, TG, TB]) GGA() (TG, TG, TA) { return c.G, c.G, c.A }
func (c RGBGeneric[TR, TG, TB]) GBR() (TG, TB, TR) { return c.G, c.B, c.R }
func (c RGBGeneric[TR, TG, TB]) GBG() (TG, TB, TG) { return c.G, c.B, c.G }
func (c RGBGeneric[TR, TG, TB]) GBB() (TG, TB, TB) { return c.G, c.B, c.B }
func (c RGBGeneric[TR, TG, TB]) GBA() (TG, TB, TA) { return c.G, c.B, c.A }
func (c RGBGeneric[TR, TG, TB]) GAR() (TG, TA, TR) { return c.G, c.A, c.R }
func (c RGBGeneric[TR, TG, TB]) GAG() (TG, TA, TG) { return c.G, c.A, c.G }
func (c RGBGeneric[TR, TG, TB]) GAB() (TG, TA, TB) { return c.G, c.A, c.B }
func (c RGBGeneric[TR, TG, TB]) GAA() (TG, TA, TA) { return c.G, c.A, c.A }
func (c RGBGeneric[TR, TG, TB]) BRR() (TB, TR, TR) { return c.B, c.R, c.R }
func (c RGBGeneric[TR, TG, TB]) BRG() (TB, TR, TG) { return c.B, c.R, c.G }
func (c RGBGeneric[TR, TG, TB]) BRB() (TB, TR, TB) { return c.B, c.R, c.B }
func (c RGBGeneric[TR, TG, TB]) BRA() (TB, TR, TA) { return c.B, c.R, c.A }
func (c RGBGeneric[TR, TG, TB]) BGR() (TB, TG, TR) { return c.B, c.G, c.R }
func (c RGBGeneric[TR, TG, TB]) BGG() (TB, TG, TG) { return c.B, c.G, c.G }
func (c RGBGeneric[TR, TG, TB]) BGB() (TB, TG, TB) { return c.B, c.G, c.B }
func (c RGBGeneric[TR, TG, TB]) BGA() (TB, TG, TA) { return c.B, c.G, c.A }
func (c RGBGeneric[TR, TG, TB]) BBR() (TB, TB, TR) { return c.B, c.B, c.R }
func (c RGBGeneric[TR, TG, TB]) BBG() (TB, TB, TG) { return c.B, c.B, c.G }
func (c RGBGeneric[TR, TG, TB]) BBB() (TB, TB, TB) { return c.B, c.B, c.B }
func (c RGBGeneric[TR, TG, TB]) BBA() (TB, TB, TA) { return c.B, c.B, c.A }
func (c RGBGeneric[TR, TG, TB]) BAR() (TB, TA, TR) { return c.B, c.A, c.R }
func (c RGBGeneric[TR, TG, TB]) BAG() (TB, TA, TG) { return c.B, c.A, c.G }
func (c RGBGeneric[TR, TG, TB]) BAB() (TB, TA, TB) { return c.B, c.A, c.B }
func (c RGBGeneric[TR, TG, TB]) BAA() (TB, TA, TA) { return c.B, c.A, c.A }
func (c RGBGeneric[TR, TG, TB]) ARR() (TA, TR, TR) { return c.A, c.R, c.R }
func (c RGBGeneric[TR, TG, TB]) ARG() (TA, TR, TG) { return c.A, c.R, c.G }
func (c RGBGeneric[TR, TG, TB]) ARB() (TA, TR, TB) { return c.A, c.R, c.B }
func (c RGBGeneric[TR, TG, TB]) ARA() (TA, TR, TA) { return c.A, c.R, c.A }
func (c RGBGeneric[TR, TG, TB]) AGR() (TA, TG, TR) { return c.A, c.G, c.R }
func (c RGBGeneric[TR, TG, TB]) AGG() (TA, TG, TG) { return c.A, c.G, c.G }
func (c RGBGeneric[TR, TG, TB]) AGB() (TA, TG, TB) { return c.A, c.G, c.B }
func (c RGBGeneric[TR, TG, TB]) AGA() (TA, TG, TA) { return c.A, c.G, c.A }
func (c RGBGeneric[TR, TG, TB]) ABR() (TA, TB, TR) { return c.A, c.B, c.R }
func (c RGBGeneric[TR, TG, TB]) ABG() (TA, TB, TG) { return c.A, c.B, c.G }
func (c RGBGeneric[TR, TG, TB]) ABB() (TA, TB, TB) { return c.A, c.B, c.B }
func (c RGBGeneric[TR, TG, TB]) ABA() (TA, TB, TA) { return c.A, c.B, c.A }
func (c RGBGeneric[TR, TG, TB]) AAR() (TA, TA, TR) { return c.A, c.A, c.R }
func (c RGBGeneric[TR, TG, TB]) AAG() (TA, TA, TG) { return c.A, c.A, c.G }
func (c RGBGeneric[TR, TG, TB]) AAB() (TA, TA, TB) { return c.A, c.A, c.B }
func (c RGBGeneric[TR, TG, TB]) AAA() (TA, TA, TA) { return c.A, c.A, c.A }
