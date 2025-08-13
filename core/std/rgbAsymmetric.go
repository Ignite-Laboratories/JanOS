package std

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/num"
)

// RGBAsymmetric is the underlying structure of color operations.  It differs from the more common RGB in that it supports
// asymmetric channel bit widths - gaining the lengthier type name as it's far more common in the modern age to work with
// symmetric channel widths.
//
// This provides accessibility to legacy color spaces while allowing the more common RGB type to evolve from old paradigms,
// rather than dismissing their existence entirely. =)
//
// NOTE: This type also provides rudimentary "swizzling."
type RGBAsymmetric[TR num.Primitive, TG num.Primitive, TB num.Primitive] struct {
	// R is the red channel.
	R Bounded[TR]

	// G is the green channel.
	G Bounded[TG]

	// B is the blue channel.
	B Bounded[TB]
}

// SetClamp sets whether the color channels should clamp to their boundaries or overflow and under-flow.
func (c RGBAsymmetric[TR, TG, TB]) SetClamp(shouldClamp bool) RGBAsymmetric[TR, TG, TB] {
	c.R.Clamp = shouldClamp
	c.G.Clamp = shouldClamp
	c.B.Clamp = shouldClamp
	return c
}

// Set sets the all color channels and returns the new color.
func (c RGBAsymmetric[TR, TG, TB]) Set(r TR, g TG, b TB) RGBAsymmetric[TR, TG, TB] {
	c.R, _ = c.R.SetAll(r, 0, TR(num.MaxValue[TR]()))
	c.G, _ = c.G.SetAll(g, 0, TG(num.MaxValue[TG]()))
	c.B, _ = c.B.SetAll(b, 0, TB(num.MaxValue[TB]()))
	return c
}

// SetRed sets the red channel and returns the new color.
func (c RGBAsymmetric[TR, TG, TB]) SetRed(r TR) RGBAsymmetric[TR, TG, TB] {
	c.R, _ = c.R.Set(r)
	return c
}

// SetGreen sets the green channel and returns the new color.
func (c RGBAsymmetric[TR, TG, TB]) SetGreen(g TG) RGBAsymmetric[TR, TG, TB] {
	c.G, _ = c.G.Set(g)
	return c
}

// SetBlue sets the blue channel and returns the new color.
func (c RGBAsymmetric[TR, TG, TB]) SetBlue(b TB) RGBAsymmetric[TR, TG, TB] {
	c.B, _ = c.B.Set(b)
	return c
}

// SetFromNormalized sets the bounded directional values using float64 unit vectors from the [0.0, 1.0]
// range, where 0.0 maps to the coordinate space's bounded minimum and 1.0 maps to the bounded maximum.
func (c RGBAsymmetric[TR, TG, TB]) SetFromNormalized(r, g, b float64) RGBAsymmetric[TR, TG, TB] {
	c.R, _ = c.R.SetFromNormalized(r)
	c.G, _ = c.G.SetFromNormalized(g)
	c.B, _ = c.B.SetFromNormalized(b)
	return c
}

// SetFromNormalized32 sets the bounded directional values using float32 unit vectors from the [0.0, 1.0]
// range, where 0.0 maps to the coordinate space's bounded minimum and 1.0 maps to the bounded maximum.
func (c RGBAsymmetric[TR, TG, TB]) SetFromNormalized32(r, g, b float32) RGBAsymmetric[TR, TG, TB] {
	return c.SetFromNormalized(float64(r), float64(g), float64(b))
}

// Normalize converts the bounded directional values to float64 unit vectors in the range [0.0, 1.0],
// where the coordinate space's bounded minimum maps to 0.0 and the bounded maximum maps to 1.0.
func (c RGBAsymmetric[TR, TG, TB]) Normalize() (float64, float64, float64) {
	return c.R.Normalize(), c.G.Normalize(), c.B.Normalize()
}

// Normalize32 converts the bounded directional values to float32 unit vectors in the range [0.0, 1.0],
// where the coordinate space's bounded minimum maps to 0.0 and the bounded maximum maps to 1.0.
func (c RGBAsymmetric[TR, TG, TB]) Normalize32() (float32, float32, float32) {
	return c.R.Normalize32(), c.G.Normalize32(), c.B.Normalize32()
}

func (c RGBAsymmetric[TR, TG, TB]) String() string {
	var rZero TR
	var gZero TG
	var bZero TB
	return fmt.Sprintf("rgba[%T, %T, %T](%v, %v, %v)", rZero, gZero, bZero, c.R.Value(), c.G.Value(), c.B.Value())
}

/**
Swizzling

NOTE: This is a regular expression to find and replace swizzle functions into a one-liner if the auto formatter ever kicks in

Find -
func \((.*?)\) ([A-Z]{2,4})\(\) \((.*?)\)[ ]*\{[\n\t ]*return(.*?)[\n\t ]*\}

Replace -
func ($1) $2() ($3) { return$4 }
*/

func (c RGBAsymmetric[TR, TG, TB]) RR() (TR, TR) { return c.R.Value(), c.R.Value() }
func (c RGBAsymmetric[TR, TG, TB]) RG() (TR, TG) { return c.R.Value(), c.G.Value() }
func (c RGBAsymmetric[TR, TG, TB]) RB() (TR, TB) { return c.R.Value(), c.B.Value() }
func (c RGBAsymmetric[TR, TG, TB]) GR() (TG, TR) { return c.G.Value(), c.R.Value() }
func (c RGBAsymmetric[TR, TG, TB]) GG() (TG, TG) { return c.G.Value(), c.G.Value() }
func (c RGBAsymmetric[TR, TG, TB]) GB() (TG, TB) { return c.G.Value(), c.B.Value() }
func (c RGBAsymmetric[TR, TG, TB]) BR() (TB, TR) { return c.B.Value(), c.R.Value() }
func (c RGBAsymmetric[TR, TG, TB]) BG() (TB, TG) { return c.B.Value(), c.G.Value() }
func (c RGBAsymmetric[TR, TG, TB]) BB() (TB, TB) { return c.B.Value(), c.B.Value() }

func (c RGBAsymmetric[TR, TG, TB]) RRR() (TR, TR, TR) { return c.R.Value(), c.R.Value(), c.R.Value() }
func (c RGBAsymmetric[TR, TG, TB]) RRG() (TR, TR, TG) { return c.R.Value(), c.R.Value(), c.G.Value() }
func (c RGBAsymmetric[TR, TG, TB]) RRB() (TR, TR, TB) { return c.R.Value(), c.R.Value(), c.B.Value() }
func (c RGBAsymmetric[TR, TG, TB]) RGR() (TR, TG, TR) { return c.R.Value(), c.G.Value(), c.R.Value() }
func (c RGBAsymmetric[TR, TG, TB]) RGG() (TR, TG, TG) { return c.R.Value(), c.G.Value(), c.G.Value() }
func (c RGBAsymmetric[TR, TG, TB]) RGB() (TR, TG, TB) { return c.R.Value(), c.G.Value(), c.B.Value() }
func (c RGBAsymmetric[TR, TG, TB]) RBR() (TR, TB, TR) { return c.R.Value(), c.B.Value(), c.R.Value() }
func (c RGBAsymmetric[TR, TG, TB]) RBG() (TR, TB, TG) { return c.R.Value(), c.B.Value(), c.G.Value() }
func (c RGBAsymmetric[TR, TG, TB]) RBB() (TR, TB, TB) { return c.R.Value(), c.B.Value(), c.B.Value() }
func (c RGBAsymmetric[TR, TG, TB]) GRR() (TG, TR, TR) { return c.G.Value(), c.R.Value(), c.R.Value() }
func (c RGBAsymmetric[TR, TG, TB]) GRG() (TG, TR, TG) { return c.G.Value(), c.R.Value(), c.G.Value() }
func (c RGBAsymmetric[TR, TG, TB]) GRB() (TG, TR, TB) { return c.G.Value(), c.R.Value(), c.B.Value() }
func (c RGBAsymmetric[TR, TG, TB]) GGR() (TG, TG, TR) { return c.G.Value(), c.G.Value(), c.R.Value() }
func (c RGBAsymmetric[TR, TG, TB]) GGG() (TG, TG, TG) { return c.G.Value(), c.G.Value(), c.G.Value() }
func (c RGBAsymmetric[TR, TG, TB]) GGB() (TG, TG, TB) { return c.G.Value(), c.G.Value(), c.B.Value() }
func (c RGBAsymmetric[TR, TG, TB]) GBR() (TG, TB, TR) { return c.G.Value(), c.B.Value(), c.R.Value() }
func (c RGBAsymmetric[TR, TG, TB]) GBG() (TG, TB, TG) { return c.G.Value(), c.B.Value(), c.G.Value() }
func (c RGBAsymmetric[TR, TG, TB]) GBB() (TG, TB, TB) { return c.G.Value(), c.B.Value(), c.B.Value() }
func (c RGBAsymmetric[TR, TG, TB]) BRR() (TB, TR, TR) { return c.B.Value(), c.R.Value(), c.R.Value() }
func (c RGBAsymmetric[TR, TG, TB]) BRG() (TB, TR, TG) { return c.B.Value(), c.R.Value(), c.G.Value() }
func (c RGBAsymmetric[TR, TG, TB]) BRB() (TB, TR, TB) { return c.B.Value(), c.R.Value(), c.B.Value() }
func (c RGBAsymmetric[TR, TG, TB]) BGR() (TB, TG, TR) { return c.B.Value(), c.G.Value(), c.R.Value() }
func (c RGBAsymmetric[TR, TG, TB]) BGG() (TB, TG, TG) { return c.B.Value(), c.G.Value(), c.G.Value() }
func (c RGBAsymmetric[TR, TG, TB]) BGB() (TB, TG, TB) { return c.B.Value(), c.G.Value(), c.B.Value() }
func (c RGBAsymmetric[TR, TG, TB]) BBR() (TB, TB, TR) { return c.B.Value(), c.B.Value(), c.R.Value() }
func (c RGBAsymmetric[TR, TG, TB]) BBG() (TB, TB, TG) { return c.B.Value(), c.B.Value(), c.G.Value() }
func (c RGBAsymmetric[TR, TG, TB]) BBB() (TB, TB, TB) { return c.B.Value(), c.B.Value(), c.B.Value() }

func (c RGBAsymmetric[TR, TG, TB]) RRRR() (TR, TR, TR, TR) {
	return c.R.Value(), c.R.Value(), c.R.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RRRG() (TR, TR, TR, TG) {
	return c.R.Value(), c.R.Value(), c.R.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RRRB() (TR, TR, TR, TB) {
	return c.R.Value(), c.R.Value(), c.R.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RRGR() (TR, TR, TG, TR) {
	return c.R.Value(), c.R.Value(), c.G.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RRGG() (TR, TR, TG, TG) {
	return c.R.Value(), c.R.Value(), c.G.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RRGB() (TR, TR, TG, TB) {
	return c.R.Value(), c.R.Value(), c.G.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RRBR() (TR, TR, TB, TR) {
	return c.R.Value(), c.R.Value(), c.B.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RRBG() (TR, TR, TB, TG) {
	return c.R.Value(), c.R.Value(), c.B.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RRBB() (TR, TR, TB, TB) {
	return c.R.Value(), c.R.Value(), c.B.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RGRR() (TR, TG, TR, TR) {
	return c.R.Value(), c.G.Value(), c.R.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RGRG() (TR, TG, TR, TG) {
	return c.R.Value(), c.G.Value(), c.R.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RGRB() (TR, TG, TR, TB) {
	return c.R.Value(), c.G.Value(), c.R.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RGGR() (TR, TG, TG, TR) {
	return c.R.Value(), c.G.Value(), c.G.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RGGG() (TR, TG, TG, TG) {
	return c.R.Value(), c.G.Value(), c.G.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RGGB() (TR, TG, TG, TB) {
	return c.R.Value(), c.G.Value(), c.G.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RGBR() (TR, TG, TB, TR) {
	return c.R.Value(), c.G.Value(), c.B.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RGBG() (TR, TG, TB, TG) {
	return c.R.Value(), c.G.Value(), c.B.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RGBB() (TR, TG, TB, TB) {
	return c.R.Value(), c.G.Value(), c.B.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RBRR() (TR, TB, TR, TR) {
	return c.R.Value(), c.B.Value(), c.R.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RBRG() (TR, TB, TR, TG) {
	return c.R.Value(), c.B.Value(), c.R.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RBRB() (TR, TB, TR, TB) {
	return c.R.Value(), c.B.Value(), c.R.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RBGR() (TR, TB, TG, TR) {
	return c.R.Value(), c.B.Value(), c.G.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RBGG() (TR, TB, TG, TG) {
	return c.R.Value(), c.B.Value(), c.G.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RBGB() (TR, TB, TG, TB) {
	return c.R.Value(), c.B.Value(), c.G.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RBBR() (TR, TB, TB, TR) {
	return c.R.Value(), c.B.Value(), c.B.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RBBG() (TR, TB, TB, TG) {
	return c.R.Value(), c.B.Value(), c.B.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) RBBB() (TR, TB, TB, TB) {
	return c.R.Value(), c.B.Value(), c.B.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GRRR() (TG, TR, TR, TR) {
	return c.G.Value(), c.R.Value(), c.R.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GRRG() (TG, TR, TR, TG) {
	return c.G.Value(), c.R.Value(), c.R.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GRRB() (TG, TR, TR, TB) {
	return c.G.Value(), c.R.Value(), c.R.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GRGR() (TG, TR, TG, TR) {
	return c.G.Value(), c.R.Value(), c.G.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GRGG() (TG, TR, TG, TG) {
	return c.G.Value(), c.R.Value(), c.G.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GRGB() (TG, TR, TG, TB) {
	return c.G.Value(), c.R.Value(), c.G.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GRBR() (TG, TR, TB, TR) {
	return c.G.Value(), c.R.Value(), c.B.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GRBG() (TG, TR, TB, TG) {
	return c.G.Value(), c.R.Value(), c.B.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GRBB() (TG, TR, TB, TB) {
	return c.G.Value(), c.R.Value(), c.B.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GGRR() (TG, TG, TR, TR) {
	return c.G.Value(), c.G.Value(), c.R.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GGRG() (TG, TG, TR, TG) {
	return c.G.Value(), c.G.Value(), c.R.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GGRB() (TG, TG, TR, TB) {
	return c.G.Value(), c.G.Value(), c.R.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GGGR() (TG, TG, TG, TR) {
	return c.G.Value(), c.G.Value(), c.G.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GGGG() (TG, TG, TG, TG) {
	return c.G.Value(), c.G.Value(), c.G.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GGGB() (TG, TG, TG, TB) {
	return c.G.Value(), c.G.Value(), c.G.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GGBR() (TG, TG, TB, TR) {
	return c.G.Value(), c.G.Value(), c.B.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GGBG() (TG, TG, TB, TG) {
	return c.G.Value(), c.G.Value(), c.B.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GGBB() (TG, TG, TB, TB) {
	return c.G.Value(), c.G.Value(), c.B.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GBRR() (TG, TB, TR, TR) {
	return c.G.Value(), c.B.Value(), c.R.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GBRG() (TG, TB, TR, TG) {
	return c.G.Value(), c.B.Value(), c.R.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GBRB() (TG, TB, TR, TB) {
	return c.G.Value(), c.B.Value(), c.R.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GBGR() (TG, TB, TG, TR) {
	return c.G.Value(), c.B.Value(), c.G.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GBGG() (TG, TB, TG, TG) {
	return c.G.Value(), c.B.Value(), c.G.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GBGB() (TG, TB, TG, TB) {
	return c.G.Value(), c.B.Value(), c.G.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GBBR() (TG, TB, TB, TR) {
	return c.G.Value(), c.B.Value(), c.B.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GBBG() (TG, TB, TB, TG) {
	return c.G.Value(), c.B.Value(), c.B.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) GBBB() (TG, TB, TB, TB) {
	return c.G.Value(), c.B.Value(), c.B.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BRRR() (TB, TR, TR, TR) {
	return c.B.Value(), c.R.Value(), c.R.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BRRG() (TB, TR, TR, TG) {
	return c.B.Value(), c.R.Value(), c.R.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BRRB() (TB, TR, TR, TB) {
	return c.B.Value(), c.R.Value(), c.R.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BRGR() (TB, TR, TG, TR) {
	return c.B.Value(), c.R.Value(), c.G.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BRGG() (TB, TR, TG, TG) {
	return c.B.Value(), c.R.Value(), c.G.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BRGB() (TB, TR, TG, TB) {
	return c.B.Value(), c.R.Value(), c.G.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BRBR() (TB, TR, TB, TR) {
	return c.B.Value(), c.R.Value(), c.B.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BRBG() (TB, TR, TB, TG) {
	return c.B.Value(), c.R.Value(), c.B.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BRBB() (TB, TR, TB, TB) {
	return c.B.Value(), c.R.Value(), c.B.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BGRR() (TB, TG, TR, TR) {
	return c.B.Value(), c.G.Value(), c.R.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BGRG() (TB, TG, TR, TG) {
	return c.B.Value(), c.G.Value(), c.R.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BGRB() (TB, TG, TR, TB) {
	return c.B.Value(), c.G.Value(), c.R.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BGGR() (TB, TG, TG, TR) {
	return c.B.Value(), c.G.Value(), c.G.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BGGG() (TB, TG, TG, TG) {
	return c.B.Value(), c.G.Value(), c.G.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BGGB() (TB, TG, TG, TB) {
	return c.B.Value(), c.G.Value(), c.G.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BGBR() (TB, TG, TB, TR) {
	return c.B.Value(), c.G.Value(), c.B.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BGBG() (TB, TG, TB, TG) {
	return c.B.Value(), c.G.Value(), c.B.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BGBB() (TB, TG, TB, TB) {
	return c.B.Value(), c.G.Value(), c.B.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BBRR() (TB, TB, TR, TR) {
	return c.B.Value(), c.B.Value(), c.R.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BBRG() (TB, TB, TR, TG) {
	return c.B.Value(), c.B.Value(), c.R.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BBRB() (TB, TB, TR, TB) {
	return c.B.Value(), c.B.Value(), c.R.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BBGR() (TB, TB, TG, TR) {
	return c.B.Value(), c.B.Value(), c.G.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BBGG() (TB, TB, TG, TG) {
	return c.B.Value(), c.B.Value(), c.G.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BBGB() (TB, TB, TG, TB) {
	return c.B.Value(), c.B.Value(), c.G.Value(), c.B.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BBBR() (TB, TB, TB, TR) {
	return c.B.Value(), c.B.Value(), c.B.Value(), c.R.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BBBG() (TB, TB, TB, TG) {
	return c.B.Value(), c.B.Value(), c.B.Value(), c.G.Value()
}
func (c RGBAsymmetric[TR, TG, TB]) BBBB() (TB, TB, TB, TB) {
	return c.B.Value(), c.B.Value(), c.B.Value(), c.B.Value()
}
