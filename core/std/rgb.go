package std

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/num"
)

// RGB is a structure for holding symmetrical red, green, and blue color channel values - as well as providing rudimentary "swizzling."
//
// NOTE: This derives from RGBAsymmetric, which allows asymmetric channel bit widths if desired =)
type RGB[T num.Primitive] RGBAsymmetric[T, T, T]

// SetClamp sets whether the color channels should clamp to their boundaries or overflow and under-flow.
func (c RGB[T]) SetClamp(shouldClamp bool) RGB[T] {
	c.R.Clamp = shouldClamp
	c.G.Clamp = shouldClamp
	c.B.Clamp = shouldClamp
	return c
}

// Set sets the all color channels and returns the new color.
func (c RGB[T]) Set(r, g, b T) RGB[T] {
	c.R, _ = c.R.SetAll(r, 0, T(num.MaxValue[T]()))
	c.G, _ = c.G.SetAll(g, 0, T(num.MaxValue[T]()))
	c.B, _ = c.B.SetAll(b, 0, T(num.MaxValue[T]()))
	return c
}

// SetRed sets the red channel and returns the new color.
func (c RGB[T]) SetRed(r T) RGB[T] {
	c.R, _ = c.R.Set(r)
	return c
}

// SetGreen sets the green channel and returns the new color.
func (c RGB[T]) SetGreen(g T) RGB[T] {
	c.G, _ = c.G.Set(g)
	return c
}

// SetBlue sets the blue channel and returns the new color.
func (c RGB[T]) SetBlue(b T) RGB[T] {
	c.B, _ = c.B.Set(b)
	return c
}

// SetFromNormalized sets the bounded directional values using float64 unit vectors from the [0.0, 1.0]
// range, where 0.0 maps to the coordinate space's bounded minimum and 1.0 maps to the bounded maximum.
func (c RGB[T]) SetFromNormalized(r, g, b float64) RGB[T] {
	c.R, _ = c.R.SetFromNormalized(r)
	c.G, _ = c.G.SetFromNormalized(g)
	c.B, _ = c.B.SetFromNormalized(b)
	return c
}

// SetFromNormalized32 sets the bounded directional values using float32 unit vectors from the [0.0, 1.0]
// range, where 0.0 maps to the coordinate space's bounded minimum and 1.0 maps to the bounded maximum.
func (c RGB[T]) SetFromNormalized32(r, g, b float32) RGB[T] {
	return c.SetFromNormalized(float64(r), float64(g), float64(b))
}

// Normalize converts the bounded directional values to float64 unit vectors in the range [0.0, 1.0],
// where the coordinate space's bounded minimum maps to 0.0 and the bounded maximum maps to 1.0.
func (c RGB[T]) Normalize() (float64, float64, float64) {
	return c.R.Normalize(), c.G.Normalize(), c.B.Normalize()
}

// Normalize32 converts the bounded directional values to float32 unit vectors in the range [0.0, 1.0],
// where the coordinate space's bounded minimum maps to 0.0 and the bounded maximum maps to 1.0.
func (c RGB[T]) Normalize32() (float32, float32, float32) {
	return c.R.Normalize32(), c.G.Normalize32(), c.B.Normalize32()
}

func (c RGB[T]) String() string {
	var zero T
	return fmt.Sprintf("rgb[%T](%v, %v, %v)", zero, c.R.Value(), c.G.Value(), c.B.Value())
}

/**
Swizzling

NOTE: This is a regular expression to find and replace swizzle functions into a one-liner if the auto formatter ever kicks in

Find -
func \((.*?)\) ([A-Z]{2,4})\(\) \((.*?)\)[ ]*\{[\n\t ]*return(.*?)[\n\t ]*\}

Replace -
func ($1) $2() ($3) { return$4 }
*/

func (c RGB[T]) RR() (T, T) { return c.R.Value(), c.R.Value() }
func (c RGB[T]) RG() (T, T) { return c.R.Value(), c.G.Value() }
func (c RGB[T]) RB() (T, T) { return c.R.Value(), c.B.Value() }
func (c RGB[T]) GR() (T, T) { return c.G.Value(), c.R.Value() }
func (c RGB[T]) GG() (T, T) { return c.G.Value(), c.G.Value() }
func (c RGB[T]) GB() (T, T) { return c.G.Value(), c.B.Value() }
func (c RGB[T]) BR() (T, T) { return c.B.Value(), c.R.Value() }
func (c RGB[T]) BG() (T, T) { return c.B.Value(), c.G.Value() }
func (c RGB[T]) BB() (T, T) { return c.B.Value(), c.B.Value() }

func (c RGB[T]) RRR() (T, T, T) { return c.R.Value(), c.R.Value(), c.R.Value() }
func (c RGB[T]) RRG() (T, T, T) { return c.R.Value(), c.R.Value(), c.G.Value() }
func (c RGB[T]) RRB() (T, T, T) { return c.R.Value(), c.R.Value(), c.B.Value() }
func (c RGB[T]) RGR() (T, T, T) { return c.R.Value(), c.G.Value(), c.R.Value() }
func (c RGB[T]) RGG() (T, T, T) { return c.R.Value(), c.G.Value(), c.G.Value() }
func (c RGB[T]) RGB() (T, T, T) { return c.R.Value(), c.G.Value(), c.B.Value() }
func (c RGB[T]) RBR() (T, T, T) { return c.R.Value(), c.B.Value(), c.R.Value() }
func (c RGB[T]) RBG() (T, T, T) { return c.R.Value(), c.B.Value(), c.G.Value() }
func (c RGB[T]) RBB() (T, T, T) { return c.R.Value(), c.B.Value(), c.B.Value() }
func (c RGB[T]) GRR() (T, T, T) { return c.G.Value(), c.R.Value(), c.R.Value() }
func (c RGB[T]) GRG() (T, T, T) { return c.G.Value(), c.R.Value(), c.G.Value() }
func (c RGB[T]) GRB() (T, T, T) { return c.G.Value(), c.R.Value(), c.B.Value() }
func (c RGB[T]) GGR() (T, T, T) { return c.G.Value(), c.G.Value(), c.R.Value() }
func (c RGB[T]) GGG() (T, T, T) { return c.G.Value(), c.G.Value(), c.G.Value() }
func (c RGB[T]) GGB() (T, T, T) { return c.G.Value(), c.G.Value(), c.B.Value() }
func (c RGB[T]) GBR() (T, T, T) { return c.G.Value(), c.B.Value(), c.R.Value() }
func (c RGB[T]) GBG() (T, T, T) { return c.G.Value(), c.B.Value(), c.G.Value() }
func (c RGB[T]) GBB() (T, T, T) { return c.G.Value(), c.B.Value(), c.B.Value() }
func (c RGB[T]) BRR() (T, T, T) { return c.B.Value(), c.R.Value(), c.R.Value() }
func (c RGB[T]) BRG() (T, T, T) { return c.B.Value(), c.R.Value(), c.G.Value() }
func (c RGB[T]) BRB() (T, T, T) { return c.B.Value(), c.R.Value(), c.B.Value() }
func (c RGB[T]) BGR() (T, T, T) { return c.B.Value(), c.G.Value(), c.R.Value() }
func (c RGB[T]) BGG() (T, T, T) { return c.B.Value(), c.G.Value(), c.G.Value() }
func (c RGB[T]) BGB() (T, T, T) { return c.B.Value(), c.G.Value(), c.B.Value() }
func (c RGB[T]) BBR() (T, T, T) { return c.B.Value(), c.B.Value(), c.R.Value() }
func (c RGB[T]) BBG() (T, T, T) { return c.B.Value(), c.B.Value(), c.G.Value() }
func (c RGB[T]) BBB() (T, T, T) { return c.B.Value(), c.B.Value(), c.B.Value() }

func (c RGB[T]) RRRR() (T, T, T, T) { return c.R.Value(), c.R.Value(), c.R.Value(), c.R.Value() }
func (c RGB[T]) RRRG() (T, T, T, T) { return c.R.Value(), c.R.Value(), c.R.Value(), c.G.Value() }
func (c RGB[T]) RRRB() (T, T, T, T) { return c.R.Value(), c.R.Value(), c.R.Value(), c.B.Value() }
func (c RGB[T]) RRGR() (T, T, T, T) { return c.R.Value(), c.R.Value(), c.G.Value(), c.R.Value() }
func (c RGB[T]) RRGG() (T, T, T, T) { return c.R.Value(), c.R.Value(), c.G.Value(), c.G.Value() }
func (c RGB[T]) RRGB() (T, T, T, T) { return c.R.Value(), c.R.Value(), c.G.Value(), c.B.Value() }
func (c RGB[T]) RRBR() (T, T, T, T) { return c.R.Value(), c.R.Value(), c.B.Value(), c.R.Value() }
func (c RGB[T]) RRBG() (T, T, T, T) { return c.R.Value(), c.R.Value(), c.B.Value(), c.G.Value() }
func (c RGB[T]) RRBB() (T, T, T, T) { return c.R.Value(), c.R.Value(), c.B.Value(), c.B.Value() }
func (c RGB[T]) RGRR() (T, T, T, T) { return c.R.Value(), c.G.Value(), c.R.Value(), c.R.Value() }
func (c RGB[T]) RGRG() (T, T, T, T) { return c.R.Value(), c.G.Value(), c.R.Value(), c.G.Value() }
func (c RGB[T]) RGRB() (T, T, T, T) { return c.R.Value(), c.G.Value(), c.R.Value(), c.B.Value() }
func (c RGB[T]) RGGR() (T, T, T, T) { return c.R.Value(), c.G.Value(), c.G.Value(), c.R.Value() }
func (c RGB[T]) RGGG() (T, T, T, T) { return c.R.Value(), c.G.Value(), c.G.Value(), c.G.Value() }
func (c RGB[T]) RGGB() (T, T, T, T) { return c.R.Value(), c.G.Value(), c.G.Value(), c.B.Value() }
func (c RGB[T]) RGBR() (T, T, T, T) { return c.R.Value(), c.G.Value(), c.B.Value(), c.R.Value() }
func (c RGB[T]) RGBG() (T, T, T, T) { return c.R.Value(), c.G.Value(), c.B.Value(), c.G.Value() }
func (c RGB[T]) RGBB() (T, T, T, T) { return c.R.Value(), c.G.Value(), c.B.Value(), c.B.Value() }
func (c RGB[T]) RBRR() (T, T, T, T) { return c.R.Value(), c.B.Value(), c.R.Value(), c.R.Value() }
func (c RGB[T]) RBRG() (T, T, T, T) { return c.R.Value(), c.B.Value(), c.R.Value(), c.G.Value() }
func (c RGB[T]) RBRB() (T, T, T, T) { return c.R.Value(), c.B.Value(), c.R.Value(), c.B.Value() }
func (c RGB[T]) RBGR() (T, T, T, T) { return c.R.Value(), c.B.Value(), c.G.Value(), c.R.Value() }
func (c RGB[T]) RBGG() (T, T, T, T) { return c.R.Value(), c.B.Value(), c.G.Value(), c.G.Value() }
func (c RGB[T]) RBGB() (T, T, T, T) { return c.R.Value(), c.B.Value(), c.G.Value(), c.B.Value() }
func (c RGB[T]) RBBR() (T, T, T, T) { return c.R.Value(), c.B.Value(), c.B.Value(), c.R.Value() }
func (c RGB[T]) RBBG() (T, T, T, T) { return c.R.Value(), c.B.Value(), c.B.Value(), c.G.Value() }
func (c RGB[T]) RBBB() (T, T, T, T) { return c.R.Value(), c.B.Value(), c.B.Value(), c.B.Value() }
func (c RGB[T]) GRRR() (T, T, T, T) { return c.G.Value(), c.R.Value(), c.R.Value(), c.R.Value() }
func (c RGB[T]) GRRG() (T, T, T, T) { return c.G.Value(), c.R.Value(), c.R.Value(), c.G.Value() }
func (c RGB[T]) GRRB() (T, T, T, T) { return c.G.Value(), c.R.Value(), c.R.Value(), c.B.Value() }
func (c RGB[T]) GRGR() (T, T, T, T) { return c.G.Value(), c.R.Value(), c.G.Value(), c.R.Value() }
func (c RGB[T]) GRGG() (T, T, T, T) { return c.G.Value(), c.R.Value(), c.G.Value(), c.G.Value() }
func (c RGB[T]) GRGB() (T, T, T, T) { return c.G.Value(), c.R.Value(), c.G.Value(), c.B.Value() }
func (c RGB[T]) GRBR() (T, T, T, T) { return c.G.Value(), c.R.Value(), c.B.Value(), c.R.Value() }
func (c RGB[T]) GRBG() (T, T, T, T) { return c.G.Value(), c.R.Value(), c.B.Value(), c.G.Value() }
func (c RGB[T]) GRBB() (T, T, T, T) { return c.G.Value(), c.R.Value(), c.B.Value(), c.B.Value() }
func (c RGB[T]) GGRR() (T, T, T, T) { return c.G.Value(), c.G.Value(), c.R.Value(), c.R.Value() }
func (c RGB[T]) GGRG() (T, T, T, T) { return c.G.Value(), c.G.Value(), c.R.Value(), c.G.Value() }
func (c RGB[T]) GGRB() (T, T, T, T) { return c.G.Value(), c.G.Value(), c.R.Value(), c.B.Value() }
func (c RGB[T]) GGGR() (T, T, T, T) { return c.G.Value(), c.G.Value(), c.G.Value(), c.R.Value() }
func (c RGB[T]) GGGG() (T, T, T, T) { return c.G.Value(), c.G.Value(), c.G.Value(), c.G.Value() }
func (c RGB[T]) GGGB() (T, T, T, T) { return c.G.Value(), c.G.Value(), c.G.Value(), c.B.Value() }
func (c RGB[T]) GGBR() (T, T, T, T) { return c.G.Value(), c.G.Value(), c.B.Value(), c.R.Value() }
func (c RGB[T]) GGBG() (T, T, T, T) { return c.G.Value(), c.G.Value(), c.B.Value(), c.G.Value() }
func (c RGB[T]) GGBB() (T, T, T, T) { return c.G.Value(), c.G.Value(), c.B.Value(), c.B.Value() }
func (c RGB[T]) GBRR() (T, T, T, T) { return c.G.Value(), c.B.Value(), c.R.Value(), c.R.Value() }
func (c RGB[T]) GBRG() (T, T, T, T) { return c.G.Value(), c.B.Value(), c.R.Value(), c.G.Value() }
func (c RGB[T]) GBRB() (T, T, T, T) { return c.G.Value(), c.B.Value(), c.R.Value(), c.B.Value() }
func (c RGB[T]) GBGR() (T, T, T, T) { return c.G.Value(), c.B.Value(), c.G.Value(), c.R.Value() }
func (c RGB[T]) GBGG() (T, T, T, T) { return c.G.Value(), c.B.Value(), c.G.Value(), c.G.Value() }
func (c RGB[T]) GBGB() (T, T, T, T) { return c.G.Value(), c.B.Value(), c.G.Value(), c.B.Value() }
func (c RGB[T]) GBBR() (T, T, T, T) { return c.G.Value(), c.B.Value(), c.B.Value(), c.R.Value() }
func (c RGB[T]) GBBG() (T, T, T, T) { return c.G.Value(), c.B.Value(), c.B.Value(), c.G.Value() }
func (c RGB[T]) GBBB() (T, T, T, T) { return c.G.Value(), c.B.Value(), c.B.Value(), c.B.Value() }
func (c RGB[T]) BRRR() (T, T, T, T) { return c.B.Value(), c.R.Value(), c.R.Value(), c.R.Value() }
func (c RGB[T]) BRRG() (T, T, T, T) { return c.B.Value(), c.R.Value(), c.R.Value(), c.G.Value() }
func (c RGB[T]) BRRB() (T, T, T, T) { return c.B.Value(), c.R.Value(), c.R.Value(), c.B.Value() }
func (c RGB[T]) BRGR() (T, T, T, T) { return c.B.Value(), c.R.Value(), c.G.Value(), c.R.Value() }
func (c RGB[T]) BRGG() (T, T, T, T) { return c.B.Value(), c.R.Value(), c.G.Value(), c.G.Value() }
func (c RGB[T]) BRGB() (T, T, T, T) { return c.B.Value(), c.R.Value(), c.G.Value(), c.B.Value() }
func (c RGB[T]) BRBR() (T, T, T, T) { return c.B.Value(), c.R.Value(), c.B.Value(), c.R.Value() }
func (c RGB[T]) BRBG() (T, T, T, T) { return c.B.Value(), c.R.Value(), c.B.Value(), c.G.Value() }
func (c RGB[T]) BRBB() (T, T, T, T) { return c.B.Value(), c.R.Value(), c.B.Value(), c.B.Value() }
func (c RGB[T]) BGRR() (T, T, T, T) { return c.B.Value(), c.G.Value(), c.R.Value(), c.R.Value() }
func (c RGB[T]) BGRG() (T, T, T, T) { return c.B.Value(), c.G.Value(), c.R.Value(), c.G.Value() }
func (c RGB[T]) BGRB() (T, T, T, T) { return c.B.Value(), c.G.Value(), c.R.Value(), c.B.Value() }
func (c RGB[T]) BGGR() (T, T, T, T) { return c.B.Value(), c.G.Value(), c.G.Value(), c.R.Value() }
func (c RGB[T]) BGGG() (T, T, T, T) { return c.B.Value(), c.G.Value(), c.G.Value(), c.G.Value() }
func (c RGB[T]) BGGB() (T, T, T, T) { return c.B.Value(), c.G.Value(), c.G.Value(), c.B.Value() }
func (c RGB[T]) BGBR() (T, T, T, T) { return c.B.Value(), c.G.Value(), c.B.Value(), c.R.Value() }
func (c RGB[T]) BGBG() (T, T, T, T) { return c.B.Value(), c.G.Value(), c.B.Value(), c.G.Value() }
func (c RGB[T]) BGBB() (T, T, T, T) { return c.B.Value(), c.G.Value(), c.B.Value(), c.B.Value() }
func (c RGB[T]) BBRR() (T, T, T, T) { return c.B.Value(), c.B.Value(), c.R.Value(), c.R.Value() }
func (c RGB[T]) BBRG() (T, T, T, T) { return c.B.Value(), c.B.Value(), c.R.Value(), c.G.Value() }
func (c RGB[T]) BBRB() (T, T, T, T) { return c.B.Value(), c.B.Value(), c.R.Value(), c.B.Value() }
func (c RGB[T]) BBGR() (T, T, T, T) { return c.B.Value(), c.B.Value(), c.G.Value(), c.R.Value() }
func (c RGB[T]) BBGG() (T, T, T, T) { return c.B.Value(), c.B.Value(), c.G.Value(), c.G.Value() }
func (c RGB[T]) BBGB() (T, T, T, T) { return c.B.Value(), c.B.Value(), c.G.Value(), c.B.Value() }
func (c RGB[T]) BBBR() (T, T, T, T) { return c.B.Value(), c.B.Value(), c.B.Value(), c.R.Value() }
func (c RGB[T]) BBBG() (T, T, T, T) { return c.B.Value(), c.B.Value(), c.B.Value(), c.G.Value() }
func (c RGB[T]) BBBB() (T, T, T, T) { return c.B.Value(), c.B.Value(), c.B.Value(), c.B.Value() }
