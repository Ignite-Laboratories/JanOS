package std

import "github.com/ignite-laboratories/core/std/num"

// RGBAGeneric is the underlying structure of color operations.  It differs from the more common RGBA in that it supports
// asymmetric channel bit widths (as it's far more common in the modern age to work with symmetric channel widths).  This
// provides accessibility to legacy color spaces while allowing the more common RGBA type to evolve from old paradigms,
// rather than dismissing their existence entirely. =)
//
// NOTE: This type also provides rudimentary "swizzling."
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

/**
Swizzling
*/

func (c RGBAGeneric[TR, TG, TB, TA]) RR() (TR, TR) { return c.R, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) RG() (TR, TG) { return c.R, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) RB() (TR, TB) { return c.R, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) RA() (TR, TA) { return c.R, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) GR() (TG, TR) { return c.G, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) GG() (TG, TG) { return c.G, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) GB() (TG, TB) { return c.G, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) GA() (TG, TA) { return c.G, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) BR() (TB, TR) { return c.B, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) BG() (TB, TG) { return c.B, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) BB() (TB, TB) { return c.B, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) BA() (TB, TA) { return c.B, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) AR() (TA, TR) { return c.A, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) AG() (TA, TG) { return c.A, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) AB() (TA, TB) { return c.A, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) AA() (TA, TA) { return c.A, c.A }

func (c RGBAGeneric[TR, TG, TB, TA]) RRR() (TR, TR, TR) { return c.R, c.R, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) RRG() (TR, TR, TG) { return c.R, c.R, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) RRB() (TR, TR, TB) { return c.R, c.R, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) RRA() (TR, TR, TA) { return c.R, c.R, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) RGR() (TR, TG, TR) { return c.R, c.G, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) RGG() (TR, TG, TG) { return c.R, c.G, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) RGB() (TR, TG, TB) { return c.R, c.G, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) RGA() (TR, TG, TA) { return c.R, c.G, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) RBR() (TR, TB, TR) { return c.R, c.B, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) RBG() (TR, TB, TG) { return c.R, c.B, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) RBB() (TR, TB, TB) { return c.R, c.B, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) RBA() (TR, TB, TA) { return c.R, c.B, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) RAR() (TR, TA, TR) { return c.R, c.A, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) RAG() (TR, TA, TG) { return c.R, c.A, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) RAB() (TR, TA, TB) { return c.R, c.A, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) RAA() (TR, TA, TA) { return c.R, c.A, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) GRR() (TG, TR, TR) { return c.G, c.R, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) GRG() (TG, TR, TG) { return c.G, c.R, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) GRB() (TG, TR, TB) { return c.G, c.R, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) GRA() (TG, TR, TA) { return c.G, c.R, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) GGR() (TG, TG, TR) { return c.G, c.G, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) GGG() (TG, TG, TG) { return c.G, c.G, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) GGB() (TG, TG, TB) { return c.G, c.G, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) GGA() (TG, TG, TA) { return c.G, c.G, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) GBR() (TG, TB, TR) { return c.G, c.B, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) GBG() (TG, TB, TG) { return c.G, c.B, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) GBB() (TG, TB, TB) { return c.G, c.B, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) GBA() (TG, TB, TA) { return c.G, c.B, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) GAR() (TG, TA, TR) { return c.G, c.A, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) GAG() (TG, TA, TG) { return c.G, c.A, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) GAB() (TG, TA, TB) { return c.G, c.A, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) GAA() (TG, TA, TA) { return c.G, c.A, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) BRR() (TB, TR, TR) { return c.B, c.R, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) BRG() (TB, TR, TG) { return c.B, c.R, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) BRB() (TB, TR, TB) { return c.B, c.R, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) BRA() (TB, TR, TA) { return c.B, c.R, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) BGR() (TB, TG, TR) { return c.B, c.G, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) BGG() (TB, TG, TG) { return c.B, c.G, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) BGB() (TB, TG, TB) { return c.B, c.G, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) BGA() (TB, TG, TA) { return c.B, c.G, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) BBR() (TB, TB, TR) { return c.B, c.B, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) BBG() (TB, TB, TG) { return c.B, c.B, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) BBB() (TB, TB, TB) { return c.B, c.B, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) BBA() (TB, TB, TA) { return c.B, c.B, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) BAR() (TB, TA, TR) { return c.B, c.A, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) BAG() (TB, TA, TG) { return c.B, c.A, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) BAB() (TB, TA, TB) { return c.B, c.A, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) BAA() (TB, TA, TA) { return c.B, c.A, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) ARR() (TA, TR, TR) { return c.A, c.R, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) ARG() (TA, TR, TG) { return c.A, c.R, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) ARB() (TA, TR, TB) { return c.A, c.R, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) ARA() (TA, TR, TA) { return c.A, c.R, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) AGR() (TA, TG, TR) { return c.A, c.G, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) AGG() (TA, TG, TG) { return c.A, c.G, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) AGB() (TA, TG, TB) { return c.A, c.G, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) AGA() (TA, TG, TA) { return c.A, c.G, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) ABR() (TA, TB, TR) { return c.A, c.B, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) ABG() (TA, TB, TG) { return c.A, c.B, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) ABB() (TA, TB, TB) { return c.A, c.B, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) ABA() (TA, TB, TA) { return c.A, c.B, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) AAR() (TA, TA, TR) { return c.A, c.A, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) AAG() (TA, TA, TG) { return c.A, c.A, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) AAB() (TA, TA, TB) { return c.A, c.A, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) AAA() (TA, TA, TA) { return c.A, c.A, c.A }

func (c RGBAGeneric[TR, TG, TB, TA]) RRRR() (TR, TR, TR, TR) { return c.R, c.R, c.R, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) RRRG() (TR, TR, TR, TG) { return c.R, c.R, c.R, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) RRRB() (TR, TR, TR, TB) { return c.R, c.R, c.R, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) RRRA() (TR, TR, TR, TA) { return c.R, c.R, c.R, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) RRGR() (TR, TR, TG, TR) { return c.R, c.R, c.G, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) RRGG() (TR, TR, TG, TG) { return c.R, c.R, c.G, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) RRGB() (TR, TR, TG, TB) { return c.R, c.R, c.G, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) RRGA() (TR, TR, TG, TA) { return c.R, c.R, c.G, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) RRBR() (TR, TR, TB, TR) { return c.R, c.R, c.B, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) RRBG() (TR, TR, TB, TG) { return c.R, c.R, c.B, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) RRBB() (TR, TR, TB, TB) { return c.R, c.R, c.B, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) RRBA() (TR, TR, TB, TA) { return c.R, c.R, c.B, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) RRAR() (TR, TR, TA, TR) { return c.R, c.R, c.A, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) RRAG() (TR, TR, TA, TG) { return c.R, c.R, c.A, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) RRAB() (TR, TR, TA, TB) { return c.R, c.R, c.A, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) RRAA() (TR, TR, TA, TA) { return c.R, c.R, c.A, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) RGRR() (TR, TG, TR, TR) { return c.R, c.G, c.R, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) RGRG() (TR, TG, TR, TG) { return c.R, c.G, c.R, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) RGRB() (TR, TG, TR, TB) { return c.R, c.G, c.R, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) RGRA() (TR, TG, TR, TA) { return c.R, c.G, c.R, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) RGGR() (TR, TG, TG, TR) { return c.R, c.G, c.G, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) RGGG() (TR, TG, TG, TG) { return c.R, c.G, c.G, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) RGGB() (TR, TG, TG, TB) { return c.R, c.G, c.G, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) RGGA() (TR, TG, TG, TA) { return c.R, c.G, c.G, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) RGBR() (TR, TG, TB, TR) { return c.R, c.G, c.B, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) RGBG() (TR, TG, TB, TG) { return c.R, c.G, c.B, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) RGBB() (TR, TG, TB, TB) { return c.R, c.G, c.B, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) RGBA() (TR, TG, TB, TA) { return c.R, c.G, c.B, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) RGAR() (TR, TG, TA, TR) { return c.R, c.G, c.A, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) RGAG() (TR, TG, TA, TG) { return c.R, c.G, c.A, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) RGAB() (TR, TG, TA, TB) { return c.R, c.G, c.A, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) RGAA() (TR, TG, TA, TA) { return c.R, c.G, c.A, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) RBRR() (TR, TB, TR, TR) { return c.R, c.B, c.R, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) RBRG() (TR, TB, TR, TG) { return c.R, c.B, c.R, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) RBRB() (TR, TB, TR, TB) { return c.R, c.B, c.R, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) RBRA() (TR, TB, TR, TA) { return c.R, c.B, c.R, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) RBGR() (TR, TB, TG, TR) { return c.R, c.B, c.G, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) RBGG() (TR, TB, TG, TG) { return c.R, c.B, c.G, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) RBGB() (TR, TB, TG, TB) { return c.R, c.B, c.G, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) RBGA() (TR, TB, TG, TA) { return c.R, c.B, c.G, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) RBBR() (TR, TB, TB, TR) { return c.R, c.B, c.B, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) RBBG() (TR, TB, TB, TG) { return c.R, c.B, c.B, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) RBBB() (TR, TB, TB, TB) { return c.R, c.B, c.B, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) RBBA() (TR, TB, TB, TA) { return c.R, c.B, c.B, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) RBAR() (TR, TB, TA, TR) { return c.R, c.B, c.A, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) RBAG() (TR, TB, TA, TG) { return c.R, c.B, c.A, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) RBAB() (TR, TB, TA, TB) { return c.R, c.B, c.A, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) RBAA() (TR, TB, TA, TA) { return c.R, c.B, c.A, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) RARR() (TR, TA, TR, TR) { return c.R, c.A, c.R, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) RARG() (TR, TA, TR, TG) { return c.R, c.A, c.R, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) RARB() (TR, TA, TR, TB) { return c.R, c.A, c.R, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) RARA() (TR, TA, TR, TA) { return c.R, c.A, c.R, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) RAGR() (TR, TA, TG, TR) { return c.R, c.A, c.G, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) RAGG() (TR, TA, TG, TG) { return c.R, c.A, c.G, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) RAGB() (TR, TA, TG, TB) { return c.R, c.A, c.G, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) RAGA() (TR, TA, TG, TA) { return c.R, c.A, c.G, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) RABR() (TR, TA, TB, TR) { return c.R, c.A, c.B, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) RABG() (TR, TA, TB, TG) { return c.R, c.A, c.B, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) RABB() (TR, TA, TB, TB) { return c.R, c.A, c.B, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) RABA() (TR, TA, TB, TA) { return c.R, c.A, c.B, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) RAAR() (TR, TA, TA, TR) { return c.R, c.A, c.A, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) RAAG() (TR, TA, TA, TG) { return c.R, c.A, c.A, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) RAAB() (TR, TA, TA, TB) { return c.R, c.A, c.A, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) RAAA() (TR, TA, TA, TA) { return c.R, c.A, c.A, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) GRRR() (TG, TR, TR, TR) { return c.G, c.R, c.R, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) GRRG() (TG, TR, TR, TG) { return c.G, c.R, c.R, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) GRRB() (TG, TR, TR, TB) { return c.G, c.R, c.R, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) GRRA() (TG, TR, TR, TA) { return c.G, c.R, c.R, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) GRGR() (TG, TR, TG, TR) { return c.G, c.R, c.G, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) GRGG() (TG, TR, TG, TG) { return c.G, c.R, c.G, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) GRGB() (TG, TR, TG, TB) { return c.G, c.R, c.G, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) GRGA() (TG, TR, TG, TA) { return c.G, c.R, c.G, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) GRBR() (TG, TR, TB, TR) { return c.G, c.R, c.B, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) GRBG() (TG, TR, TB, TG) { return c.G, c.R, c.B, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) GRBB() (TG, TR, TB, TB) { return c.G, c.R, c.B, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) GRBA() (TG, TR, TB, TA) { return c.G, c.R, c.B, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) GRAR() (TG, TR, TA, TR) { return c.G, c.R, c.A, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) GRAG() (TG, TR, TA, TG) { return c.G, c.R, c.A, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) GRAB() (TG, TR, TA, TB) { return c.G, c.R, c.A, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) GRAA() (TG, TR, TA, TA) { return c.G, c.R, c.A, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) GGRR() (TG, TG, TR, TR) { return c.G, c.G, c.R, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) GGRG() (TG, TG, TR, TG) { return c.G, c.G, c.R, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) GGRB() (TG, TG, TR, TB) { return c.G, c.G, c.R, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) GGRA() (TG, TG, TR, TA) { return c.G, c.G, c.R, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) GGGR() (TG, TG, TG, TR) { return c.G, c.G, c.G, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) GGGG() (TG, TG, TG, TG) { return c.G, c.G, c.G, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) GGGB() (TG, TG, TG, TB) { return c.G, c.G, c.G, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) GGGA() (TG, TG, TG, TA) { return c.G, c.G, c.G, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) GGBR() (TG, TG, TB, TR) { return c.G, c.G, c.B, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) GGBG() (TG, TG, TB, TG) { return c.G, c.G, c.B, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) GGBB() (TG, TG, TB, TB) { return c.G, c.G, c.B, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) GGBA() (TG, TG, TB, TA) { return c.G, c.G, c.B, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) GGAR() (TG, TG, TA, TR) { return c.G, c.G, c.A, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) GGAG() (TG, TG, TA, TG) { return c.G, c.G, c.A, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) GGAB() (TG, TG, TA, TB) { return c.G, c.G, c.A, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) GGAA() (TG, TG, TA, TA) { return c.G, c.G, c.A, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) GBRR() (TG, TB, TR, TR) { return c.G, c.B, c.R, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) GBRG() (TG, TB, TR, TG) { return c.G, c.B, c.R, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) GBRB() (TG, TB, TR, TB) { return c.G, c.B, c.R, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) GBRA() (TG, TB, TR, TA) { return c.G, c.B, c.R, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) GBGR() (TG, TB, TG, TR) { return c.G, c.B, c.G, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) GBGG() (TG, TB, TG, TG) { return c.G, c.B, c.G, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) GBGB() (TG, TB, TG, TB) { return c.G, c.B, c.G, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) GBGA() (TG, TB, TG, TA) { return c.G, c.B, c.G, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) GBBR() (TG, TB, TB, TR) { return c.G, c.B, c.B, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) GBBG() (TG, TB, TB, TG) { return c.G, c.B, c.B, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) GBBB() (TG, TB, TB, TB) { return c.G, c.B, c.B, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) GBBA() (TG, TB, TB, TA) { return c.G, c.B, c.B, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) GBAR() (TG, TB, TA, TR) { return c.G, c.B, c.A, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) GBAG() (TG, TB, TA, TG) { return c.G, c.B, c.A, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) GBAB() (TG, TB, TA, TB) { return c.G, c.B, c.A, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) GBAA() (TG, TB, TA, TA) { return c.G, c.B, c.A, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) GARR() (TG, TA, TR, TR) { return c.G, c.A, c.R, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) GARG() (TG, TA, TR, TG) { return c.G, c.A, c.R, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) GARB() (TG, TA, TR, TB) { return c.G, c.A, c.R, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) GARA() (TG, TA, TR, TA) { return c.G, c.A, c.R, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) GAGR() (TG, TA, TG, TR) { return c.G, c.A, c.G, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) GAGG() (TG, TA, TG, TG) { return c.G, c.A, c.G, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) GAGB() (TG, TA, TG, TB) { return c.G, c.A, c.G, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) GAGA() (TG, TA, TG, TA) { return c.G, c.A, c.G, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) GABR() (TG, TA, TB, TR) { return c.G, c.A, c.B, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) GABG() (TG, TA, TB, TG) { return c.G, c.A, c.B, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) GABB() (TG, TA, TB, TB) { return c.G, c.A, c.B, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) GABA() (TG, TA, TB, TA) { return c.G, c.A, c.B, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) GAAR() (TG, TA, TA, TR) { return c.G, c.A, c.A, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) GAAG() (TG, TA, TA, TG) { return c.G, c.A, c.A, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) GAAB() (TG, TA, TA, TB) { return c.G, c.A, c.A, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) GAAA() (TG, TA, TA, TA) { return c.G, c.A, c.A, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) BRRR() (TB, TR, TR, TR) { return c.B, c.R, c.R, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) BRRG() (TB, TR, TR, TG) { return c.B, c.R, c.R, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) BRRB() (TB, TR, TR, TB) { return c.B, c.R, c.R, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) BRRA() (TB, TR, TR, TA) { return c.B, c.R, c.R, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) BRGR() (TB, TR, TG, TR) { return c.B, c.R, c.G, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) BRGG() (TB, TR, TG, TG) { return c.B, c.R, c.G, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) BRGB() (TB, TR, TG, TB) { return c.B, c.R, c.G, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) BRGA() (TB, TR, TG, TA) { return c.B, c.R, c.G, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) BRBR() (TB, TR, TB, TR) { return c.B, c.R, c.B, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) BRBG() (TB, TR, TB, TG) { return c.B, c.R, c.B, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) BRBB() (TB, TR, TB, TB) { return c.B, c.R, c.B, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) BRBA() (TB, TR, TB, TA) { return c.B, c.R, c.B, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) BRAR() (TB, TR, TA, TR) { return c.B, c.R, c.A, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) BRAG() (TB, TR, TA, TG) { return c.B, c.R, c.A, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) BRAB() (TB, TR, TA, TB) { return c.B, c.R, c.A, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) BRAA() (TB, TR, TA, TA) { return c.B, c.R, c.A, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) BGRR() (TB, TG, TR, TR) { return c.B, c.G, c.R, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) BGRG() (TB, TG, TR, TG) { return c.B, c.G, c.R, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) BGRB() (TB, TG, TR, TB) { return c.B, c.G, c.R, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) BGRA() (TB, TG, TR, TA) { return c.B, c.G, c.R, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) BGGR() (TB, TG, TG, TR) { return c.B, c.G, c.G, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) BGGG() (TB, TG, TG, TG) { return c.B, c.G, c.G, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) BGGB() (TB, TG, TG, TB) { return c.B, c.G, c.G, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) BGGA() (TB, TG, TG, TA) { return c.B, c.G, c.G, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) BGBR() (TB, TG, TB, TR) { return c.B, c.G, c.B, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) BGBG() (TB, TG, TB, TG) { return c.B, c.G, c.B, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) BGBB() (TB, TG, TB, TB) { return c.B, c.G, c.B, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) BGBA() (TB, TG, TB, TA) { return c.B, c.G, c.B, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) BGAR() (TB, TG, TA, TR) { return c.B, c.G, c.A, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) BGAG() (TB, TG, TA, TG) { return c.B, c.G, c.A, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) BGAB() (TB, TG, TA, TB) { return c.B, c.G, c.A, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) BGAA() (TB, TG, TA, TA) { return c.B, c.G, c.A, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) BBRR() (TB, TB, TR, TR) { return c.B, c.B, c.R, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) BBRG() (TB, TB, TR, TG) { return c.B, c.B, c.R, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) BBRB() (TB, TB, TR, TB) { return c.B, c.B, c.R, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) BBRA() (TB, TB, TR, TA) { return c.B, c.B, c.R, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) BBGR() (TB, TB, TG, TR) { return c.B, c.B, c.G, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) BBGG() (TB, TB, TG, TG) { return c.B, c.B, c.G, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) BBGB() (TB, TB, TG, TB) { return c.B, c.B, c.G, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) BBGA() (TB, TB, TG, TA) { return c.B, c.B, c.G, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) BBBR() (TB, TB, TB, TR) { return c.B, c.B, c.B, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) BBBG() (TB, TB, TB, TG) { return c.B, c.B, c.B, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) BBBB() (TB, TB, TB, TB) { return c.B, c.B, c.B, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) BBBA() (TB, TB, TB, TA) { return c.B, c.B, c.B, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) BBAR() (TB, TB, TA, TR) { return c.B, c.B, c.A, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) BBAG() (TB, TB, TA, TG) { return c.B, c.B, c.A, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) BBAB() (TB, TB, TA, TB) { return c.B, c.B, c.A, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) BBAA() (TB, TB, TA, TA) { return c.B, c.B, c.A, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) BARR() (TB, TA, TR, TR) { return c.B, c.A, c.R, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) BARG() (TB, TA, TR, TG) { return c.B, c.A, c.R, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) BARB() (TB, TA, TR, TB) { return c.B, c.A, c.R, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) BARA() (TB, TA, TR, TA) { return c.B, c.A, c.R, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) BAGR() (TB, TA, TG, TR) { return c.B, c.A, c.G, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) BAGG() (TB, TA, TG, TG) { return c.B, c.A, c.G, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) BAGB() (TB, TA, TG, TB) { return c.B, c.A, c.G, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) BAGA() (TB, TA, TG, TA) { return c.B, c.A, c.G, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) BABR() (TB, TA, TB, TR) { return c.B, c.A, c.B, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) BABG() (TB, TA, TB, TG) { return c.B, c.A, c.B, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) BABB() (TB, TA, TB, TB) { return c.B, c.A, c.B, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) BABA() (TB, TA, TB, TA) { return c.B, c.A, c.B, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) BAAR() (TB, TA, TA, TR) { return c.B, c.A, c.A, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) BAAG() (TB, TA, TA, TG) { return c.B, c.A, c.A, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) BAAB() (TB, TA, TA, TB) { return c.B, c.A, c.A, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) BAAA() (TB, TA, TA, TA) { return c.B, c.A, c.A, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) ARRR() (TA, TR, TR, TR) { return c.A, c.R, c.R, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) ARRG() (TA, TR, TR, TG) { return c.A, c.R, c.R, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) ARRB() (TA, TR, TR, TB) { return c.A, c.R, c.R, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) ARRA() (TA, TR, TR, TA) { return c.A, c.R, c.R, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) ARGR() (TA, TR, TG, TR) { return c.A, c.R, c.G, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) ARGG() (TA, TR, TG, TG) { return c.A, c.R, c.G, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) ARGB() (TA, TR, TG, TB) { return c.A, c.R, c.G, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) ARGA() (TA, TR, TG, TA) { return c.A, c.R, c.G, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) ARBR() (TA, TR, TB, TR) { return c.A, c.R, c.B, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) ARBG() (TA, TR, TB, TG) { return c.A, c.R, c.B, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) ARBB() (TA, TR, TB, TB) { return c.A, c.R, c.B, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) ARBA() (TA, TR, TB, TA) { return c.A, c.R, c.B, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) ARAR() (TA, TR, TA, TR) { return c.A, c.R, c.A, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) ARAG() (TA, TR, TA, TG) { return c.A, c.R, c.A, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) ARAB() (TA, TR, TA, TB) { return c.A, c.R, c.A, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) ARAA() (TA, TR, TA, TA) { return c.A, c.R, c.A, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) AGRR() (TA, TG, TR, TR) { return c.A, c.G, c.R, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) AGRG() (TA, TG, TR, TG) { return c.A, c.G, c.R, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) AGRB() (TA, TG, TR, TB) { return c.A, c.G, c.R, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) AGRA() (TA, TG, TR, TA) { return c.A, c.G, c.R, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) AGGR() (TA, TG, TG, TR) { return c.A, c.G, c.G, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) AGGG() (TA, TG, TG, TG) { return c.A, c.G, c.G, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) AGGB() (TA, TG, TG, TB) { return c.A, c.G, c.G, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) AGGA() (TA, TG, TG, TA) { return c.A, c.G, c.G, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) AGBR() (TA, TG, TB, TR) { return c.A, c.G, c.B, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) AGBG() (TA, TG, TB, TG) { return c.A, c.G, c.B, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) AGBB() (TA, TG, TB, TB) { return c.A, c.G, c.B, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) AGBA() (TA, TG, TB, TA) { return c.A, c.G, c.B, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) AGAR() (TA, TG, TA, TR) { return c.A, c.G, c.A, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) AGAG() (TA, TG, TA, TG) { return c.A, c.G, c.A, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) AGAB() (TA, TG, TA, TB) { return c.A, c.G, c.A, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) AGAA() (TA, TG, TA, TA) { return c.A, c.G, c.A, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) ABRR() (TA, TB, TR, TR) { return c.A, c.B, c.R, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) ABRG() (TA, TB, TR, TG) { return c.A, c.B, c.R, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) ABRB() (TA, TB, TR, TB) { return c.A, c.B, c.R, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) ABRA() (TA, TB, TR, TA) { return c.A, c.B, c.R, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) ABGR() (TA, TB, TG, TR) { return c.A, c.B, c.G, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) ABGG() (TA, TB, TG, TG) { return c.A, c.B, c.G, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) ABGB() (TA, TB, TG, TB) { return c.A, c.B, c.G, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) ABGA() (TA, TB, TG, TA) { return c.A, c.B, c.G, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) ABBR() (TA, TB, TB, TR) { return c.A, c.B, c.B, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) ABBG() (TA, TB, TB, TG) { return c.A, c.B, c.B, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) ABBB() (TA, TB, TB, TB) { return c.A, c.B, c.B, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) ABBA() (TA, TB, TB, TA) { return c.A, c.B, c.B, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) ABAR() (TA, TB, TA, TR) { return c.A, c.B, c.A, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) ABAG() (TA, TB, TA, TG) { return c.A, c.B, c.A, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) ABAB() (TA, TB, TA, TB) { return c.A, c.B, c.A, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) ABAA() (TA, TB, TA, TA) { return c.A, c.B, c.A, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) AARR() (TA, TA, TR, TR) { return c.A, c.A, c.R, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) AARG() (TA, TA, TR, TG) { return c.A, c.A, c.R, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) AARB() (TA, TA, TR, TB) { return c.A, c.A, c.R, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) AARA() (TA, TA, TR, TA) { return c.A, c.A, c.R, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) AAGR() (TA, TA, TG, TR) { return c.A, c.A, c.G, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) AAGG() (TA, TA, TG, TG) { return c.A, c.A, c.G, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) AAGB() (TA, TA, TG, TB) { return c.A, c.A, c.G, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) AAGA() (TA, TA, TG, TA) { return c.A, c.A, c.G, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) AABR() (TA, TA, TB, TR) { return c.A, c.A, c.B, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) AABG() (TA, TA, TB, TG) { return c.A, c.A, c.B, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) AABB() (TA, TA, TB, TB) { return c.A, c.A, c.B, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) AABA() (TA, TA, TB, TA) { return c.A, c.A, c.B, c.A }
func (c RGBAGeneric[TR, TG, TB, TA]) AAAR() (TA, TA, TA, TR) { return c.A, c.A, c.A, c.R }
func (c RGBAGeneric[TR, TG, TB, TA]) AAAG() (TA, TA, TA, TG) { return c.A, c.A, c.A, c.G }
func (c RGBAGeneric[TR, TG, TB, TA]) AAAB() (TA, TA, TA, TB) { return c.A, c.A, c.A, c.B }
func (c RGBAGeneric[TR, TG, TB, TA]) AAAA() (TA, TA, TA, TA) { return c.A, c.A, c.A, c.A }
