package std

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/num"
)

// RGB is a structure for holding symmetrical red, green, and blue color channel values - as well as providing rudimentary "swizzling."
//
// NOTE: This derives from RGBGeneric, which allows asymmetric channel bit widths if desired =)
type RGB[T num.ExtendedPrimitive] RGBGeneric[T, T, T]

func (c RGB[T]) String() string {
	return fmt.Sprintf("(%v, %v, %v)", c.R, c.G, c.B)
}

/**
Swizzling
*/

func (c RGB[T]) RR() (T, T) { return c.R, c.R }
func (c RGB[T]) RG() (T, T) { return c.R, c.G }
func (c RGB[T]) RB() (T, T) { return c.R, c.B }
func (c RGB[T]) GR() (T, T) { return c.G, c.R }
func (c RGB[T]) GG() (T, T) { return c.G, c.G }
func (c RGB[T]) GB() (T, T) { return c.G, c.B }
func (c RGB[T]) BR() (T, T) { return c.B, c.R }
func (c RGB[T]) BG() (T, T) { return c.B, c.G }
func (c RGB[T]) BB() (T, T) { return c.B, c.B }

func (c RGB[T]) RRR() (T, T, T) { return c.R, c.R, c.R }
func (c RGB[T]) RRG() (T, T, T) { return c.R, c.R, c.G }
func (c RGB[T]) RRB() (T, T, T) { return c.R, c.R, c.B }
func (c RGB[T]) RGR() (T, T, T) { return c.R, c.G, c.R }
func (c RGB[T]) RGG() (T, T, T) { return c.R, c.G, c.G }
func (c RGB[T]) RGB() (T, T, T) { return c.R, c.G, c.B }
func (c RGB[T]) RBR() (T, T, T) { return c.R, c.B, c.R }
func (c RGB[T]) RBG() (T, T, T) { return c.R, c.B, c.G }
func (c RGB[T]) RBB() (T, T, T) { return c.R, c.B, c.B }
func (c RGB[T]) GRR() (T, T, T) { return c.G, c.R, c.R }
func (c RGB[T]) GRG() (T, T, T) { return c.G, c.R, c.G }
func (c RGB[T]) GRB() (T, T, T) { return c.G, c.R, c.B }
func (c RGB[T]) GGR() (T, T, T) { return c.G, c.G, c.R }
func (c RGB[T]) GGG() (T, T, T) { return c.G, c.G, c.G }
func (c RGB[T]) GGB() (T, T, T) { return c.G, c.G, c.B }
func (c RGB[T]) GBR() (T, T, T) { return c.G, c.B, c.R }
func (c RGB[T]) GBG() (T, T, T) { return c.G, c.B, c.G }
func (c RGB[T]) GBB() (T, T, T) { return c.G, c.B, c.B }
func (c RGB[T]) BRR() (T, T, T) { return c.B, c.R, c.R }
func (c RGB[T]) BRG() (T, T, T) { return c.B, c.R, c.G }
func (c RGB[T]) BRB() (T, T, T) { return c.B, c.R, c.B }
func (c RGB[T]) BGR() (T, T, T) { return c.B, c.G, c.R }
func (c RGB[T]) BGG() (T, T, T) { return c.B, c.G, c.G }
func (c RGB[T]) BGB() (T, T, T) { return c.B, c.G, c.B }
func (c RGB[T]) BBR() (T, T, T) { return c.B, c.B, c.R }
func (c RGB[T]) BBG() (T, T, T) { return c.B, c.B, c.G }
func (c RGB[T]) BBB() (T, T, T) { return c.B, c.B, c.B }

func (c RGB[T]) RRRR() (T, T, T, T) { return c.R, c.R, c.R, c.R }
func (c RGB[T]) RRRG() (T, T, T, T) { return c.R, c.R, c.R, c.G }
func (c RGB[T]) RRRB() (T, T, T, T) { return c.R, c.R, c.R, c.B }
func (c RGB[T]) RRGR() (T, T, T, T) { return c.R, c.R, c.G, c.R }
func (c RGB[T]) RRGG() (T, T, T, T) { return c.R, c.R, c.G, c.G }
func (c RGB[T]) RRGB() (T, T, T, T) { return c.R, c.R, c.G, c.B }
func (c RGB[T]) RRBR() (T, T, T, T) { return c.R, c.R, c.B, c.R }
func (c RGB[T]) RRBG() (T, T, T, T) { return c.R, c.R, c.B, c.G }
func (c RGB[T]) RRBB() (T, T, T, T) { return c.R, c.R, c.B, c.B }
func (c RGB[T]) RGRR() (T, T, T, T) { return c.R, c.G, c.R, c.R }
func (c RGB[T]) RGRG() (T, T, T, T) { return c.R, c.G, c.R, c.G }
func (c RGB[T]) RGRB() (T, T, T, T) { return c.R, c.G, c.R, c.B }
func (c RGB[T]) RGGR() (T, T, T, T) { return c.R, c.G, c.G, c.R }
func (c RGB[T]) RGGG() (T, T, T, T) { return c.R, c.G, c.G, c.G }
func (c RGB[T]) RGGB() (T, T, T, T) { return c.R, c.G, c.G, c.B }
func (c RGB[T]) RGBR() (T, T, T, T) { return c.R, c.G, c.B, c.R }
func (c RGB[T]) RGBG() (T, T, T, T) { return c.R, c.G, c.B, c.G }
func (c RGB[T]) RGBB() (T, T, T, T) { return c.R, c.G, c.B, c.B }
func (c RGB[T]) RBRR() (T, T, T, T) { return c.R, c.B, c.R, c.R }
func (c RGB[T]) RBRG() (T, T, T, T) { return c.R, c.B, c.R, c.G }
func (c RGB[T]) RBRB() (T, T, T, T) { return c.R, c.B, c.R, c.B }
func (c RGB[T]) RBGR() (T, T, T, T) { return c.R, c.B, c.G, c.R }
func (c RGB[T]) RBGG() (T, T, T, T) { return c.R, c.B, c.G, c.G }
func (c RGB[T]) RBGB() (T, T, T, T) { return c.R, c.B, c.G, c.B }
func (c RGB[T]) RBBR() (T, T, T, T) { return c.R, c.B, c.B, c.R }
func (c RGB[T]) RBBG() (T, T, T, T) { return c.R, c.B, c.B, c.G }
func (c RGB[T]) RBBB() (T, T, T, T) { return c.R, c.B, c.B, c.B }
func (c RGB[T]) GRRR() (T, T, T, T) { return c.G, c.R, c.R, c.R }
func (c RGB[T]) GRRG() (T, T, T, T) { return c.G, c.R, c.R, c.G }
func (c RGB[T]) GRRB() (T, T, T, T) { return c.G, c.R, c.R, c.B }
func (c RGB[T]) GRGR() (T, T, T, T) { return c.G, c.R, c.G, c.R }
func (c RGB[T]) GRGG() (T, T, T, T) { return c.G, c.R, c.G, c.G }
func (c RGB[T]) GRGB() (T, T, T, T) { return c.G, c.R, c.G, c.B }
func (c RGB[T]) GRBR() (T, T, T, T) { return c.G, c.R, c.B, c.R }
func (c RGB[T]) GRBG() (T, T, T, T) { return c.G, c.R, c.B, c.G }
func (c RGB[T]) GRBB() (T, T, T, T) { return c.G, c.R, c.B, c.B }
func (c RGB[T]) GGRR() (T, T, T, T) { return c.G, c.G, c.R, c.R }
func (c RGB[T]) GGRG() (T, T, T, T) { return c.G, c.G, c.R, c.G }
func (c RGB[T]) GGRB() (T, T, T, T) { return c.G, c.G, c.R, c.B }
func (c RGB[T]) GGGR() (T, T, T, T) { return c.G, c.G, c.G, c.R }
func (c RGB[T]) GGGG() (T, T, T, T) { return c.G, c.G, c.G, c.G }
func (c RGB[T]) GGGB() (T, T, T, T) { return c.G, c.G, c.G, c.B }
func (c RGB[T]) GGBR() (T, T, T, T) { return c.G, c.G, c.B, c.R }
func (c RGB[T]) GGBG() (T, T, T, T) { return c.G, c.G, c.B, c.G }
func (c RGB[T]) GGBB() (T, T, T, T) { return c.G, c.G, c.B, c.B }
func (c RGB[T]) GBRR() (T, T, T, T) { return c.G, c.B, c.R, c.R }
func (c RGB[T]) GBRG() (T, T, T, T) { return c.G, c.B, c.R, c.G }
func (c RGB[T]) GBRB() (T, T, T, T) { return c.G, c.B, c.R, c.B }
func (c RGB[T]) GBGR() (T, T, T, T) { return c.G, c.B, c.G, c.R }
func (c RGB[T]) GBGG() (T, T, T, T) { return c.G, c.B, c.G, c.G }
func (c RGB[T]) GBGB() (T, T, T, T) { return c.G, c.B, c.G, c.B }
func (c RGB[T]) GBBR() (T, T, T, T) { return c.G, c.B, c.B, c.R }
func (c RGB[T]) GBBG() (T, T, T, T) { return c.G, c.B, c.B, c.G }
func (c RGB[T]) GBBB() (T, T, T, T) { return c.G, c.B, c.B, c.B }
func (c RGB[T]) BRRR() (T, T, T, T) { return c.B, c.R, c.R, c.R }
func (c RGB[T]) BRRG() (T, T, T, T) { return c.B, c.R, c.R, c.G }
func (c RGB[T]) BRRB() (T, T, T, T) { return c.B, c.R, c.R, c.B }
func (c RGB[T]) BRGR() (T, T, T, T) { return c.B, c.R, c.G, c.R }
func (c RGB[T]) BRGG() (T, T, T, T) { return c.B, c.R, c.G, c.G }
func (c RGB[T]) BRGB() (T, T, T, T) { return c.B, c.R, c.G, c.B }
func (c RGB[T]) BRBR() (T, T, T, T) { return c.B, c.R, c.B, c.R }
func (c RGB[T]) BRBG() (T, T, T, T) { return c.B, c.R, c.B, c.G }
func (c RGB[T]) BRBB() (T, T, T, T) { return c.B, c.R, c.B, c.B }
func (c RGB[T]) BGRR() (T, T, T, T) { return c.B, c.G, c.R, c.R }
func (c RGB[T]) BGRG() (T, T, T, T) { return c.B, c.G, c.R, c.G }
func (c RGB[T]) BGRB() (T, T, T, T) { return c.B, c.G, c.R, c.B }
func (c RGB[T]) BGGR() (T, T, T, T) { return c.B, c.G, c.G, c.R }
func (c RGB[T]) BGGG() (T, T, T, T) { return c.B, c.G, c.G, c.G }
func (c RGB[T]) BGGB() (T, T, T, T) { return c.B, c.G, c.G, c.B }
func (c RGB[T]) BGBR() (T, T, T, T) { return c.B, c.G, c.B, c.R }
func (c RGB[T]) BGBG() (T, T, T, T) { return c.B, c.G, c.B, c.G }
func (c RGB[T]) BGBB() (T, T, T, T) { return c.B, c.G, c.B, c.B }
func (c RGB[T]) BBRR() (T, T, T, T) { return c.B, c.B, c.R, c.R }
func (c RGB[T]) BBRG() (T, T, T, T) { return c.B, c.B, c.R, c.G }
func (c RGB[T]) BBRB() (T, T, T, T) { return c.B, c.B, c.R, c.B }
func (c RGB[T]) BBGR() (T, T, T, T) { return c.B, c.B, c.G, c.R }
func (c RGB[T]) BBGG() (T, T, T, T) { return c.B, c.B, c.G, c.G }
func (c RGB[T]) BBGB() (T, T, T, T) { return c.B, c.B, c.G, c.B }
func (c RGB[T]) BBBR() (T, T, T, T) { return c.B, c.B, c.B, c.R }
func (c RGB[T]) BBBG() (T, T, T, T) { return c.B, c.B, c.B, c.G }
func (c RGB[T]) BBBB() (T, T, T, T) { return c.B, c.B, c.B, c.B }
