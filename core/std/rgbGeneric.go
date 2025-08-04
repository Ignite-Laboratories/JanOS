package std

import "github.com/ignite-laboratories/core/std/num"

// RGBGeneric is the underlying structure of color operations.  It differs from the more common RGB in that it supports
// asymmetric channel bit widths (as it's far more common in the modern age to work with symmetric channel widths).  This
// provides accessibility to legacy color spaces while allowing the more common RGB type to evolve from old paradigms,
// rather than dismissing their existence entirely. =)
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
func (c RGBGeneric[TR, TG, TB]) GR() (TG, TR) { return c.G, c.R }
func (c RGBGeneric[TR, TG, TB]) GG() (TG, TG) { return c.G, c.G }
func (c RGBGeneric[TR, TG, TB]) GB() (TG, TB) { return c.G, c.B }
func (c RGBGeneric[TR, TG, TB]) BR() (TB, TR) { return c.B, c.R }
func (c RGBGeneric[TR, TG, TB]) BG() (TB, TG) { return c.B, c.G }
func (c RGBGeneric[TR, TG, TB]) BB() (TB, TB) { return c.B, c.B }

func (c RGBGeneric[TR, TG, TB]) RRR() (TR, TR, TR) { return c.R, c.R, c.R }
func (c RGBGeneric[TR, TG, TB]) RRG() (TR, TR, TG) { return c.R, c.R, c.G }
func (c RGBGeneric[TR, TG, TB]) RRB() (TR, TR, TB) { return c.R, c.R, c.B }
func (c RGBGeneric[TR, TG, TB]) RGR() (TR, TG, TR) { return c.R, c.G, c.R }
func (c RGBGeneric[TR, TG, TB]) RGG() (TR, TG, TG) { return c.R, c.G, c.G }
func (c RGBGeneric[TR, TG, TB]) RGB() (TR, TG, TB) { return c.R, c.G, c.B }
func (c RGBGeneric[TR, TG, TB]) RBR() (TR, TB, TR) { return c.R, c.B, c.R }
func (c RGBGeneric[TR, TG, TB]) RBG() (TR, TB, TG) { return c.R, c.B, c.G }
func (c RGBGeneric[TR, TG, TB]) RBB() (TR, TB, TB) { return c.R, c.B, c.B }
func (c RGBGeneric[TR, TG, TB]) GRR() (TG, TR, TR) { return c.G, c.R, c.R }
func (c RGBGeneric[TR, TG, TB]) GRG() (TG, TR, TG) { return c.G, c.R, c.G }
func (c RGBGeneric[TR, TG, TB]) GRB() (TG, TR, TB) { return c.G, c.R, c.B }
func (c RGBGeneric[TR, TG, TB]) GGR() (TG, TG, TR) { return c.G, c.G, c.R }
func (c RGBGeneric[TR, TG, TB]) GGG() (TG, TG, TG) { return c.G, c.G, c.G }
func (c RGBGeneric[TR, TG, TB]) GGB() (TG, TG, TB) { return c.G, c.G, c.B }
func (c RGBGeneric[TR, TG, TB]) GBR() (TG, TB, TR) { return c.G, c.B, c.R }
func (c RGBGeneric[TR, TG, TB]) GBG() (TG, TB, TG) { return c.G, c.B, c.G }
func (c RGBGeneric[TR, TG, TB]) GBB() (TG, TB, TB) { return c.G, c.B, c.B }
func (c RGBGeneric[TR, TG, TB]) BRR() (TB, TR, TR) { return c.B, c.R, c.R }
func (c RGBGeneric[TR, TG, TB]) BRG() (TB, TR, TG) { return c.B, c.R, c.G }
func (c RGBGeneric[TR, TG, TB]) BRB() (TB, TR, TB) { return c.B, c.R, c.B }
func (c RGBGeneric[TR, TG, TB]) BGR() (TB, TG, TR) { return c.B, c.G, c.R }
func (c RGBGeneric[TR, TG, TB]) BGG() (TB, TG, TG) { return c.B, c.G, c.G }
func (c RGBGeneric[TR, TG, TB]) BGB() (TB, TG, TB) { return c.B, c.G, c.B }
func (c RGBGeneric[TR, TG, TB]) BBR() (TB, TB, TR) { return c.B, c.B, c.R }
func (c RGBGeneric[TR, TG, TB]) BBG() (TB, TB, TG) { return c.B, c.B, c.G }
func (c RGBGeneric[TR, TG, TB]) BBB() (TB, TB, TB) { return c.B, c.B, c.B }

func (c RGBGeneric[TR, TG, TB]) RRRR() (TR, TR, TR, TR) { return c.R, c.R, c.R, c.R }
func (c RGBGeneric[TR, TG, TB]) RRRG() (TR, TR, TR, TG) { return c.R, c.R, c.R, c.G }
func (c RGBGeneric[TR, TG, TB]) RRRB() (TR, TR, TR, TB) { return c.R, c.R, c.R, c.B }
func (c RGBGeneric[TR, TG, TB]) RRGR() (TR, TR, TG, TR) { return c.R, c.R, c.G, c.R }
func (c RGBGeneric[TR, TG, TB]) RRGG() (TR, TR, TG, TG) { return c.R, c.R, c.G, c.G }
func (c RGBGeneric[TR, TG, TB]) RRGB() (TR, TR, TG, TB) { return c.R, c.R, c.G, c.B }
func (c RGBGeneric[TR, TG, TB]) RRBR() (TR, TR, TB, TR) { return c.R, c.R, c.B, c.R }
func (c RGBGeneric[TR, TG, TB]) RRBG() (TR, TR, TB, TG) { return c.R, c.R, c.B, c.G }
func (c RGBGeneric[TR, TG, TB]) RRBB() (TR, TR, TB, TB) { return c.R, c.R, c.B, c.B }
func (c RGBGeneric[TR, TG, TB]) RGRR() (TR, TG, TR, TR) { return c.R, c.G, c.R, c.R }
func (c RGBGeneric[TR, TG, TB]) RGRG() (TR, TG, TR, TG) { return c.R, c.G, c.R, c.G }
func (c RGBGeneric[TR, TG, TB]) RGRB() (TR, TG, TR, TB) { return c.R, c.G, c.R, c.B }
func (c RGBGeneric[TR, TG, TB]) RGGR() (TR, TG, TG, TR) { return c.R, c.G, c.G, c.R }
func (c RGBGeneric[TR, TG, TB]) RGGG() (TR, TG, TG, TG) { return c.R, c.G, c.G, c.G }
func (c RGBGeneric[TR, TG, TB]) RGGB() (TR, TG, TG, TB) { return c.R, c.G, c.G, c.B }
func (c RGBGeneric[TR, TG, TB]) RGBR() (TR, TG, TB, TR) { return c.R, c.G, c.B, c.R }
func (c RGBGeneric[TR, TG, TB]) RGBG() (TR, TG, TB, TG) { return c.R, c.G, c.B, c.G }
func (c RGBGeneric[TR, TG, TB]) RGBB() (TR, TG, TB, TB) { return c.R, c.G, c.B, c.B }
func (c RGBGeneric[TR, TG, TB]) RBRR() (TR, TB, TR, TR) { return c.R, c.B, c.R, c.R }
func (c RGBGeneric[TR, TG, TB]) RBRG() (TR, TB, TR, TG) { return c.R, c.B, c.R, c.G }
func (c RGBGeneric[TR, TG, TB]) RBRB() (TR, TB, TR, TB) { return c.R, c.B, c.R, c.B }
func (c RGBGeneric[TR, TG, TB]) RBGR() (TR, TB, TG, TR) { return c.R, c.B, c.G, c.R }
func (c RGBGeneric[TR, TG, TB]) RBGG() (TR, TB, TG, TG) { return c.R, c.B, c.G, c.G }
func (c RGBGeneric[TR, TG, TB]) RBGB() (TR, TB, TG, TB) { return c.R, c.B, c.G, c.B }
func (c RGBGeneric[TR, TG, TB]) RBBR() (TR, TB, TB, TR) { return c.R, c.B, c.B, c.R }
func (c RGBGeneric[TR, TG, TB]) RBBG() (TR, TB, TB, TG) { return c.R, c.B, c.B, c.G }
func (c RGBGeneric[TR, TG, TB]) RBBB() (TR, TB, TB, TB) { return c.R, c.B, c.B, c.B }
func (c RGBGeneric[TR, TG, TB]) GRRR() (TG, TR, TR, TR) { return c.G, c.R, c.R, c.R }
func (c RGBGeneric[TR, TG, TB]) GRRG() (TG, TR, TR, TG) { return c.G, c.R, c.R, c.G }
func (c RGBGeneric[TR, TG, TB]) GRRB() (TG, TR, TR, TB) { return c.G, c.R, c.R, c.B }
func (c RGBGeneric[TR, TG, TB]) GRGR() (TG, TR, TG, TR) { return c.G, c.R, c.G, c.R }
func (c RGBGeneric[TR, TG, TB]) GRGG() (TG, TR, TG, TG) { return c.G, c.R, c.G, c.G }
func (c RGBGeneric[TR, TG, TB]) GRGB() (TG, TR, TG, TB) { return c.G, c.R, c.G, c.B }
func (c RGBGeneric[TR, TG, TB]) GRBR() (TG, TR, TB, TR) { return c.G, c.R, c.B, c.R }
func (c RGBGeneric[TR, TG, TB]) GRBG() (TG, TR, TB, TG) { return c.G, c.R, c.B, c.G }
func (c RGBGeneric[TR, TG, TB]) GRBB() (TG, TR, TB, TB) { return c.G, c.R, c.B, c.B }
func (c RGBGeneric[TR, TG, TB]) GGRR() (TG, TG, TR, TR) { return c.G, c.G, c.R, c.R }
func (c RGBGeneric[TR, TG, TB]) GGRG() (TG, TG, TR, TG) { return c.G, c.G, c.R, c.G }
func (c RGBGeneric[TR, TG, TB]) GGRB() (TG, TG, TR, TB) { return c.G, c.G, c.R, c.B }
func (c RGBGeneric[TR, TG, TB]) GGGR() (TG, TG, TG, TR) { return c.G, c.G, c.G, c.R }
func (c RGBGeneric[TR, TG, TB]) GGGG() (TG, TG, TG, TG) { return c.G, c.G, c.G, c.G }
func (c RGBGeneric[TR, TG, TB]) GGGB() (TG, TG, TG, TB) { return c.G, c.G, c.G, c.B }
func (c RGBGeneric[TR, TG, TB]) GGBR() (TG, TG, TB, TR) { return c.G, c.G, c.B, c.R }
func (c RGBGeneric[TR, TG, TB]) GGBG() (TG, TG, TB, TG) { return c.G, c.G, c.B, c.G }
func (c RGBGeneric[TR, TG, TB]) GGBB() (TG, TG, TB, TB) { return c.G, c.G, c.B, c.B }
func (c RGBGeneric[TR, TG, TB]) GBRR() (TG, TB, TR, TR) { return c.G, c.B, c.R, c.R }
func (c RGBGeneric[TR, TG, TB]) GBRG() (TG, TB, TR, TG) { return c.G, c.B, c.R, c.G }
func (c RGBGeneric[TR, TG, TB]) GBRB() (TG, TB, TR, TB) { return c.G, c.B, c.R, c.B }
func (c RGBGeneric[TR, TG, TB]) GBGR() (TG, TB, TG, TR) { return c.G, c.B, c.G, c.R }
func (c RGBGeneric[TR, TG, TB]) GBGG() (TG, TB, TG, TG) { return c.G, c.B, c.G, c.G }
func (c RGBGeneric[TR, TG, TB]) GBGB() (TG, TB, TG, TB) { return c.G, c.B, c.G, c.B }
func (c RGBGeneric[TR, TG, TB]) GBBR() (TG, TB, TB, TR) { return c.G, c.B, c.B, c.R }
func (c RGBGeneric[TR, TG, TB]) GBBG() (TG, TB, TB, TG) { return c.G, c.B, c.B, c.G }
func (c RGBGeneric[TR, TG, TB]) GBBB() (TG, TB, TB, TB) { return c.G, c.B, c.B, c.B }
func (c RGBGeneric[TR, TG, TB]) BRRR() (TB, TR, TR, TR) { return c.B, c.R, c.R, c.R }
func (c RGBGeneric[TR, TG, TB]) BRRG() (TB, TR, TR, TG) { return c.B, c.R, c.R, c.G }
func (c RGBGeneric[TR, TG, TB]) BRRB() (TB, TR, TR, TB) { return c.B, c.R, c.R, c.B }
func (c RGBGeneric[TR, TG, TB]) BRGR() (TB, TR, TG, TR) { return c.B, c.R, c.G, c.R }
func (c RGBGeneric[TR, TG, TB]) BRGG() (TB, TR, TG, TG) { return c.B, c.R, c.G, c.G }
func (c RGBGeneric[TR, TG, TB]) BRGB() (TB, TR, TG, TB) { return c.B, c.R, c.G, c.B }
func (c RGBGeneric[TR, TG, TB]) BRBR() (TB, TR, TB, TR) { return c.B, c.R, c.B, c.R }
func (c RGBGeneric[TR, TG, TB]) BRBG() (TB, TR, TB, TG) { return c.B, c.R, c.B, c.G }
func (c RGBGeneric[TR, TG, TB]) BRBB() (TB, TR, TB, TB) { return c.B, c.R, c.B, c.B }
func (c RGBGeneric[TR, TG, TB]) BGRR() (TB, TG, TR, TR) { return c.B, c.G, c.R, c.R }
func (c RGBGeneric[TR, TG, TB]) BGRG() (TB, TG, TR, TG) { return c.B, c.G, c.R, c.G }
func (c RGBGeneric[TR, TG, TB]) BGRB() (TB, TG, TR, TB) { return c.B, c.G, c.R, c.B }
func (c RGBGeneric[TR, TG, TB]) BGGR() (TB, TG, TG, TR) { return c.B, c.G, c.G, c.R }
func (c RGBGeneric[TR, TG, TB]) BGGG() (TB, TG, TG, TG) { return c.B, c.G, c.G, c.G }
func (c RGBGeneric[TR, TG, TB]) BGGB() (TB, TG, TG, TB) { return c.B, c.G, c.G, c.B }
func (c RGBGeneric[TR, TG, TB]) BGBR() (TB, TG, TB, TR) { return c.B, c.G, c.B, c.R }
func (c RGBGeneric[TR, TG, TB]) BGBG() (TB, TG, TB, TG) { return c.B, c.G, c.B, c.G }
func (c RGBGeneric[TR, TG, TB]) BGBB() (TB, TG, TB, TB) { return c.B, c.G, c.B, c.B }
func (c RGBGeneric[TR, TG, TB]) BBRR() (TB, TB, TR, TR) { return c.B, c.B, c.R, c.R }
func (c RGBGeneric[TR, TG, TB]) BBRG() (TB, TB, TR, TG) { return c.B, c.B, c.R, c.G }
func (c RGBGeneric[TR, TG, TB]) BBRB() (TB, TB, TR, TB) { return c.B, c.B, c.R, c.B }
func (c RGBGeneric[TR, TG, TB]) BBGR() (TB, TB, TG, TR) { return c.B, c.B, c.G, c.R }
func (c RGBGeneric[TR, TG, TB]) BBGG() (TB, TB, TG, TG) { return c.B, c.B, c.G, c.G }
func (c RGBGeneric[TR, TG, TB]) BBGB() (TB, TB, TG, TB) { return c.B, c.B, c.G, c.B }
func (c RGBGeneric[TR, TG, TB]) BBBR() (TB, TB, TB, TR) { return c.B, c.B, c.B, c.R }
func (c RGBGeneric[TR, TG, TB]) BBBG() (TB, TB, TB, TG) { return c.B, c.B, c.B, c.G }
func (c RGBGeneric[TR, TG, TB]) BBBB() (TB, TB, TB, TB) { return c.B, c.B, c.B, c.B }
