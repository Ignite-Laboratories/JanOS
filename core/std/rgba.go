package std

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/num"
)

// RGBA is a structure for holding red, green, and blue color values.
type RGBA[T num.ExtendedPrimitive] struct {
	// R is the red channel.
	R T

	// G is the green channel.
	G T

	// B is the blue channel.
	B T

	// A is the alpha channel.
	A T
}

func (c RGBA[T]) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v)", c.R, c.G, c.B, c.A)
}

/**
Swizzling
*/

func (c RGBA[T]) RR() (T, T) { return c.R, c.R }
func (c RGBA[T]) RG() (T, T) { return c.R, c.G }
func (c RGBA[T]) RB() (T, T) { return c.R, c.B }
func (c RGBA[T]) RA() (T, T) { return c.R, c.A }
func (c RGBA[T]) GR() (T, T) { return c.G, c.R }
func (c RGBA[T]) GG() (T, T) { return c.G, c.G }
func (c RGBA[T]) GB() (T, T) { return c.G, c.B }
func (c RGBA[T]) GA() (T, T) { return c.G, c.A }
func (c RGBA[T]) BR() (T, T) { return c.B, c.R }
func (c RGBA[T]) BG() (T, T) { return c.B, c.G }
func (c RGBA[T]) BB() (T, T) { return c.B, c.B }
func (c RGBA[T]) BA() (T, T) { return c.B, c.A }
func (c RGBA[T]) AR() (T, T) { return c.A, c.R }
func (c RGBA[T]) AG() (T, T) { return c.A, c.G }
func (c RGBA[T]) AB() (T, T) { return c.A, c.B }
func (c RGBA[T]) AA() (T, T) { return c.A, c.A }

func (c RGBA[T]) RRR() (T, T, T) { return c.R, c.R, c.R }
func (c RGBA[T]) RRG() (T, T, T) { return c.R, c.R, c.G }
func (c RGBA[T]) RRB() (T, T, T) { return c.R, c.R, c.B }
func (c RGBA[T]) RRA() (T, T, T) { return c.R, c.R, c.A }
func (c RGBA[T]) RGR() (T, T, T) { return c.R, c.G, c.R }
func (c RGBA[T]) RGG() (T, T, T) { return c.R, c.G, c.G }
func (c RGBA[T]) RGB() (T, T, T) { return c.R, c.G, c.B }
func (c RGBA[T]) RGA() (T, T, T) { return c.R, c.G, c.A }
func (c RGBA[T]) RBR() (T, T, T) { return c.R, c.B, c.R }
func (c RGBA[T]) RBG() (T, T, T) { return c.R, c.B, c.G }
func (c RGBA[T]) RBB() (T, T, T) { return c.R, c.B, c.B }
func (c RGBA[T]) RBA() (T, T, T) { return c.R, c.B, c.A }
func (c RGBA[T]) RAR() (T, T, T) { return c.R, c.A, c.R }
func (c RGBA[T]) RAG() (T, T, T) { return c.R, c.A, c.G }
func (c RGBA[T]) RAB() (T, T, T) { return c.R, c.A, c.B }
func (c RGBA[T]) RAA() (T, T, T) { return c.R, c.A, c.A }
func (c RGBA[T]) GRR() (T, T, T) { return c.G, c.R, c.R }
func (c RGBA[T]) GRG() (T, T, T) { return c.G, c.R, c.G }
func (c RGBA[T]) GRB() (T, T, T) { return c.G, c.R, c.B }
func (c RGBA[T]) GRA() (T, T, T) { return c.G, c.R, c.A }
func (c RGBA[T]) GGR() (T, T, T) { return c.G, c.G, c.R }
func (c RGBA[T]) GGG() (T, T, T) { return c.G, c.G, c.G }
func (c RGBA[T]) GGB() (T, T, T) { return c.G, c.G, c.B }
func (c RGBA[T]) GGA() (T, T, T) { return c.G, c.G, c.A }
func (c RGBA[T]) GBR() (T, T, T) { return c.G, c.B, c.R }
func (c RGBA[T]) GBG() (T, T, T) { return c.G, c.B, c.G }
func (c RGBA[T]) GBB() (T, T, T) { return c.G, c.B, c.B }
func (c RGBA[T]) GBA() (T, T, T) { return c.G, c.B, c.A }
func (c RGBA[T]) GAR() (T, T, T) { return c.G, c.A, c.R }
func (c RGBA[T]) GAG() (T, T, T) { return c.G, c.A, c.G }
func (c RGBA[T]) GAB() (T, T, T) { return c.G, c.A, c.B }
func (c RGBA[T]) GAA() (T, T, T) { return c.G, c.A, c.A }
func (c RGBA[T]) BRR() (T, T, T) { return c.B, c.R, c.R }
func (c RGBA[T]) BRG() (T, T, T) { return c.B, c.R, c.G }
func (c RGBA[T]) BRB() (T, T, T) { return c.B, c.R, c.B }
func (c RGBA[T]) BRA() (T, T, T) { return c.B, c.R, c.A }
func (c RGBA[T]) BGR() (T, T, T) { return c.B, c.G, c.R }
func (c RGBA[T]) BGG() (T, T, T) { return c.B, c.G, c.G }
func (c RGBA[T]) BGB() (T, T, T) { return c.B, c.G, c.B }
func (c RGBA[T]) BGA() (T, T, T) { return c.B, c.G, c.A }
func (c RGBA[T]) BBR() (T, T, T) { return c.B, c.B, c.R }
func (c RGBA[T]) BBG() (T, T, T) { return c.B, c.B, c.G }
func (c RGBA[T]) BBB() (T, T, T) { return c.B, c.B, c.B }
func (c RGBA[T]) BBA() (T, T, T) { return c.B, c.B, c.A }
func (c RGBA[T]) BAR() (T, T, T) { return c.B, c.A, c.R }
func (c RGBA[T]) BAG() (T, T, T) { return c.B, c.A, c.G }
func (c RGBA[T]) BAB() (T, T, T) { return c.B, c.A, c.B }
func (c RGBA[T]) BAA() (T, T, T) { return c.B, c.A, c.A }
func (c RGBA[T]) ARR() (T, T, T) { return c.A, c.R, c.R }
func (c RGBA[T]) ARG() (T, T, T) { return c.A, c.R, c.G }
func (c RGBA[T]) ARB() (T, T, T) { return c.A, c.R, c.B }
func (c RGBA[T]) ARA() (T, T, T) { return c.A, c.R, c.A }
func (c RGBA[T]) AGR() (T, T, T) { return c.A, c.G, c.R }
func (c RGBA[T]) AGG() (T, T, T) { return c.A, c.G, c.G }
func (c RGBA[T]) AGB() (T, T, T) { return c.A, c.G, c.B }
func (c RGBA[T]) AGA() (T, T, T) { return c.A, c.G, c.A }
func (c RGBA[T]) ABR() (T, T, T) { return c.A, c.B, c.R }
func (c RGBA[T]) ABG() (T, T, T) { return c.A, c.B, c.G }
func (c RGBA[T]) ABB() (T, T, T) { return c.A, c.B, c.B }
func (c RGBA[T]) ABA() (T, T, T) { return c.A, c.B, c.A }
func (c RGBA[T]) AAR() (T, T, T) { return c.A, c.A, c.R }
func (c RGBA[T]) AAG() (T, T, T) { return c.A, c.A, c.G }
func (c RGBA[T]) AAB() (T, T, T) { return c.A, c.A, c.B }
func (c RGBA[T]) AAA() (T, T, T) { return c.A, c.A, c.A }

func (c RGBA[T]) RRRR() (T, T, T, T) { return c.R, c.R, c.R, c.R }
func (c RGBA[T]) RRRG() (T, T, T, T) { return c.R, c.R, c.R, c.G }
func (c RGBA[T]) RRRB() (T, T, T, T) { return c.R, c.R, c.R, c.B }
func (c RGBA[T]) RRRA() (T, T, T, T) { return c.R, c.R, c.R, c.A }
func (c RGBA[T]) RRGR() (T, T, T, T) { return c.R, c.R, c.G, c.R }
func (c RGBA[T]) RRGG() (T, T, T, T) { return c.R, c.R, c.G, c.G }
func (c RGBA[T]) RRGB() (T, T, T, T) { return c.R, c.R, c.G, c.B }
func (c RGBA[T]) RRGA() (T, T, T, T) { return c.R, c.R, c.G, c.A }
func (c RGBA[T]) RRBR() (T, T, T, T) { return c.R, c.R, c.B, c.R }
func (c RGBA[T]) RRBG() (T, T, T, T) { return c.R, c.R, c.B, c.G }
func (c RGBA[T]) RRBB() (T, T, T, T) { return c.R, c.R, c.B, c.B }
func (c RGBA[T]) RRBA() (T, T, T, T) { return c.R, c.R, c.B, c.A }
func (c RGBA[T]) RRAR() (T, T, T, T) { return c.R, c.R, c.A, c.R }
func (c RGBA[T]) RRAG() (T, T, T, T) { return c.R, c.R, c.A, c.G }
func (c RGBA[T]) RRAB() (T, T, T, T) { return c.R, c.R, c.A, c.B }
func (c RGBA[T]) RRAA() (T, T, T, T) { return c.R, c.R, c.A, c.A }
func (c RGBA[T]) RGRR() (T, T, T, T) { return c.R, c.G, c.R, c.R }
func (c RGBA[T]) RGRG() (T, T, T, T) { return c.R, c.G, c.R, c.G }
func (c RGBA[T]) RGRB() (T, T, T, T) { return c.R, c.G, c.R, c.B }
func (c RGBA[T]) RGRA() (T, T, T, T) { return c.R, c.G, c.R, c.A }
func (c RGBA[T]) RGGR() (T, T, T, T) { return c.R, c.G, c.G, c.R }
func (c RGBA[T]) RGGG() (T, T, T, T) { return c.R, c.G, c.G, c.G }
func (c RGBA[T]) RGGB() (T, T, T, T) { return c.R, c.G, c.G, c.B }
func (c RGBA[T]) RGGA() (T, T, T, T) { return c.R, c.G, c.G, c.A }
func (c RGBA[T]) RGBR() (T, T, T, T) { return c.R, c.G, c.B, c.R }
func (c RGBA[T]) RGBG() (T, T, T, T) { return c.R, c.G, c.B, c.G }
func (c RGBA[T]) RGBB() (T, T, T, T) { return c.R, c.G, c.B, c.B }
func (c RGBA[T]) RGBA() (T, T, T, T) { return c.R, c.G, c.B, c.A }
func (c RGBA[T]) RGAR() (T, T, T, T) { return c.R, c.G, c.A, c.R }
func (c RGBA[T]) RGAG() (T, T, T, T) { return c.R, c.G, c.A, c.G }
func (c RGBA[T]) RGAB() (T, T, T, T) { return c.R, c.G, c.A, c.B }
func (c RGBA[T]) RGAA() (T, T, T, T) { return c.R, c.G, c.A, c.A }
func (c RGBA[T]) RBRR() (T, T, T, T) { return c.R, c.B, c.R, c.R }
func (c RGBA[T]) RBRG() (T, T, T, T) { return c.R, c.B, c.R, c.G }
func (c RGBA[T]) RBRB() (T, T, T, T) { return c.R, c.B, c.R, c.B }
func (c RGBA[T]) RBRA() (T, T, T, T) { return c.R, c.B, c.R, c.A }
func (c RGBA[T]) RBGR() (T, T, T, T) { return c.R, c.B, c.G, c.R }
func (c RGBA[T]) RBGG() (T, T, T, T) { return c.R, c.B, c.G, c.G }
func (c RGBA[T]) RBGB() (T, T, T, T) { return c.R, c.B, c.G, c.B }
func (c RGBA[T]) RBGA() (T, T, T, T) { return c.R, c.B, c.G, c.A }
func (c RGBA[T]) RBBR() (T, T, T, T) { return c.R, c.B, c.B, c.R }
func (c RGBA[T]) RBBG() (T, T, T, T) { return c.R, c.B, c.B, c.G }
func (c RGBA[T]) RBBB() (T, T, T, T) { return c.R, c.B, c.B, c.B }
func (c RGBA[T]) RBBA() (T, T, T, T) { return c.R, c.B, c.B, c.A }
func (c RGBA[T]) RBAR() (T, T, T, T) { return c.R, c.B, c.A, c.R }
func (c RGBA[T]) RBAG() (T, T, T, T) { return c.R, c.B, c.A, c.G }
func (c RGBA[T]) RBAB() (T, T, T, T) { return c.R, c.B, c.A, c.B }
func (c RGBA[T]) RBAA() (T, T, T, T) { return c.R, c.B, c.A, c.A }
func (c RGBA[T]) RARR() (T, T, T, T) { return c.R, c.A, c.R, c.R }
func (c RGBA[T]) RARG() (T, T, T, T) { return c.R, c.A, c.R, c.G }
func (c RGBA[T]) RARB() (T, T, T, T) { return c.R, c.A, c.R, c.B }
func (c RGBA[T]) RARA() (T, T, T, T) { return c.R, c.A, c.R, c.A }
func (c RGBA[T]) RAGR() (T, T, T, T) { return c.R, c.A, c.G, c.R }
func (c RGBA[T]) RAGG() (T, T, T, T) { return c.R, c.A, c.G, c.G }
func (c RGBA[T]) RAGB() (T, T, T, T) { return c.R, c.A, c.G, c.B }
func (c RGBA[T]) RAGA() (T, T, T, T) { return c.R, c.A, c.G, c.A }
func (c RGBA[T]) RABR() (T, T, T, T) { return c.R, c.A, c.B, c.R }
func (c RGBA[T]) RABG() (T, T, T, T) { return c.R, c.A, c.B, c.G }
func (c RGBA[T]) RABB() (T, T, T, T) { return c.R, c.A, c.B, c.B }
func (c RGBA[T]) RABA() (T, T, T, T) { return c.R, c.A, c.B, c.A }
func (c RGBA[T]) RAAR() (T, T, T, T) { return c.R, c.A, c.A, c.R }
func (c RGBA[T]) RAAG() (T, T, T, T) { return c.R, c.A, c.A, c.G }
func (c RGBA[T]) RAAB() (T, T, T, T) { return c.R, c.A, c.A, c.B }
func (c RGBA[T]) RAAA() (T, T, T, T) { return c.R, c.A, c.A, c.A }
func (c RGBA[T]) GRRR() (T, T, T, T) { return c.G, c.R, c.R, c.R }
func (c RGBA[T]) GRRG() (T, T, T, T) { return c.G, c.R, c.R, c.G }
func (c RGBA[T]) GRRB() (T, T, T, T) { return c.G, c.R, c.R, c.B }
func (c RGBA[T]) GRRA() (T, T, T, T) { return c.G, c.R, c.R, c.A }
func (c RGBA[T]) GRGR() (T, T, T, T) { return c.G, c.R, c.G, c.R }
func (c RGBA[T]) GRGG() (T, T, T, T) { return c.G, c.R, c.G, c.G }
func (c RGBA[T]) GRGB() (T, T, T, T) { return c.G, c.R, c.G, c.B }
func (c RGBA[T]) GRGA() (T, T, T, T) { return c.G, c.R, c.G, c.A }
func (c RGBA[T]) GRBR() (T, T, T, T) { return c.G, c.R, c.B, c.R }
func (c RGBA[T]) GRBG() (T, T, T, T) { return c.G, c.R, c.B, c.G }
func (c RGBA[T]) GRBB() (T, T, T, T) { return c.G, c.R, c.B, c.B }
func (c RGBA[T]) GRBA() (T, T, T, T) { return c.G, c.R, c.B, c.A }
func (c RGBA[T]) GRAR() (T, T, T, T) { return c.G, c.R, c.A, c.R }
func (c RGBA[T]) GRAG() (T, T, T, T) { return c.G, c.R, c.A, c.G }
func (c RGBA[T]) GRAB() (T, T, T, T) { return c.G, c.R, c.A, c.B }
func (c RGBA[T]) GRAA() (T, T, T, T) { return c.G, c.R, c.A, c.A }
func (c RGBA[T]) GGRR() (T, T, T, T) { return c.G, c.G, c.R, c.R }
func (c RGBA[T]) GGRG() (T, T, T, T) { return c.G, c.G, c.R, c.G }
func (c RGBA[T]) GGRB() (T, T, T, T) { return c.G, c.G, c.R, c.B }
func (c RGBA[T]) GGRA() (T, T, T, T) { return c.G, c.G, c.R, c.A }
func (c RGBA[T]) GGGR() (T, T, T, T) { return c.G, c.G, c.G, c.R }
func (c RGBA[T]) GGGG() (T, T, T, T) { return c.G, c.G, c.G, c.G }
func (c RGBA[T]) GGGB() (T, T, T, T) { return c.G, c.G, c.G, c.B }
func (c RGBA[T]) GGGA() (T, T, T, T) { return c.G, c.G, c.G, c.A }
func (c RGBA[T]) GGBR() (T, T, T, T) { return c.G, c.G, c.B, c.R }
func (c RGBA[T]) GGBG() (T, T, T, T) { return c.G, c.G, c.B, c.G }
func (c RGBA[T]) GGBB() (T, T, T, T) { return c.G, c.G, c.B, c.B }
func (c RGBA[T]) GGBA() (T, T, T, T) { return c.G, c.G, c.B, c.A }
func (c RGBA[T]) GGAR() (T, T, T, T) { return c.G, c.G, c.A, c.R }
func (c RGBA[T]) GGAG() (T, T, T, T) { return c.G, c.G, c.A, c.G }
func (c RGBA[T]) GGAB() (T, T, T, T) { return c.G, c.G, c.A, c.B }
func (c RGBA[T]) GGAA() (T, T, T, T) { return c.G, c.G, c.A, c.A }
func (c RGBA[T]) GBRR() (T, T, T, T) { return c.G, c.B, c.R, c.R }
func (c RGBA[T]) GBRG() (T, T, T, T) { return c.G, c.B, c.R, c.G }
func (c RGBA[T]) GBRB() (T, T, T, T) { return c.G, c.B, c.R, c.B }
func (c RGBA[T]) GBRA() (T, T, T, T) { return c.G, c.B, c.R, c.A }
func (c RGBA[T]) GBGR() (T, T, T, T) { return c.G, c.B, c.G, c.R }
func (c RGBA[T]) GBGG() (T, T, T, T) { return c.G, c.B, c.G, c.G }
func (c RGBA[T]) GBGB() (T, T, T, T) { return c.G, c.B, c.G, c.B }
func (c RGBA[T]) GBGA() (T, T, T, T) { return c.G, c.B, c.G, c.A }
func (c RGBA[T]) GBBR() (T, T, T, T) { return c.G, c.B, c.B, c.R }
func (c RGBA[T]) GBBG() (T, T, T, T) { return c.G, c.B, c.B, c.G }
func (c RGBA[T]) GBBB() (T, T, T, T) { return c.G, c.B, c.B, c.B }
func (c RGBA[T]) GBBA() (T, T, T, T) { return c.G, c.B, c.B, c.A }
func (c RGBA[T]) GBAR() (T, T, T, T) { return c.G, c.B, c.A, c.R }
func (c RGBA[T]) GBAG() (T, T, T, T) { return c.G, c.B, c.A, c.G }
func (c RGBA[T]) GBAB() (T, T, T, T) { return c.G, c.B, c.A, c.B }
func (c RGBA[T]) GBAA() (T, T, T, T) { return c.G, c.B, c.A, c.A }
func (c RGBA[T]) GARR() (T, T, T, T) { return c.G, c.A, c.R, c.R }
func (c RGBA[T]) GARG() (T, T, T, T) { return c.G, c.A, c.R, c.G }
func (c RGBA[T]) GARB() (T, T, T, T) { return c.G, c.A, c.R, c.B }
func (c RGBA[T]) GARA() (T, T, T, T) { return c.G, c.A, c.R, c.A }
func (c RGBA[T]) GAGR() (T, T, T, T) { return c.G, c.A, c.G, c.R }
func (c RGBA[T]) GAGG() (T, T, T, T) { return c.G, c.A, c.G, c.G }
func (c RGBA[T]) GAGB() (T, T, T, T) { return c.G, c.A, c.G, c.B }
func (c RGBA[T]) GAGA() (T, T, T, T) { return c.G, c.A, c.G, c.A }
func (c RGBA[T]) GABR() (T, T, T, T) { return c.G, c.A, c.B, c.R }
func (c RGBA[T]) GABG() (T, T, T, T) { return c.G, c.A, c.B, c.G }
func (c RGBA[T]) GABB() (T, T, T, T) { return c.G, c.A, c.B, c.B }
func (c RGBA[T]) GABA() (T, T, T, T) { return c.G, c.A, c.B, c.A }
func (c RGBA[T]) GAAR() (T, T, T, T) { return c.G, c.A, c.A, c.R }
func (c RGBA[T]) GAAG() (T, T, T, T) { return c.G, c.A, c.A, c.G }
func (c RGBA[T]) GAAB() (T, T, T, T) { return c.G, c.A, c.A, c.B }
func (c RGBA[T]) GAAA() (T, T, T, T) { return c.G, c.A, c.A, c.A }
func (c RGBA[T]) BRRR() (T, T, T, T) { return c.B, c.R, c.R, c.R }
func (c RGBA[T]) BRRG() (T, T, T, T) { return c.B, c.R, c.R, c.G }
func (c RGBA[T]) BRRB() (T, T, T, T) { return c.B, c.R, c.R, c.B }
func (c RGBA[T]) BRRA() (T, T, T, T) { return c.B, c.R, c.R, c.A }
func (c RGBA[T]) BRGR() (T, T, T, T) { return c.B, c.R, c.G, c.R }
func (c RGBA[T]) BRGG() (T, T, T, T) { return c.B, c.R, c.G, c.G }
func (c RGBA[T]) BRGB() (T, T, T, T) { return c.B, c.R, c.G, c.B }
func (c RGBA[T]) BRGA() (T, T, T, T) { return c.B, c.R, c.G, c.A }
func (c RGBA[T]) BRBR() (T, T, T, T) { return c.B, c.R, c.B, c.R }
func (c RGBA[T]) BRBG() (T, T, T, T) { return c.B, c.R, c.B, c.G }
func (c RGBA[T]) BRBB() (T, T, T, T) { return c.B, c.R, c.B, c.B }
func (c RGBA[T]) BRBA() (T, T, T, T) { return c.B, c.R, c.B, c.A }
func (c RGBA[T]) BRAR() (T, T, T, T) { return c.B, c.R, c.A, c.R }
func (c RGBA[T]) BRAG() (T, T, T, T) { return c.B, c.R, c.A, c.G }
func (c RGBA[T]) BRAB() (T, T, T, T) { return c.B, c.R, c.A, c.B }
func (c RGBA[T]) BRAA() (T, T, T, T) { return c.B, c.R, c.A, c.A }
func (c RGBA[T]) BGRR() (T, T, T, T) { return c.B, c.G, c.R, c.R }
func (c RGBA[T]) BGRG() (T, T, T, T) { return c.B, c.G, c.R, c.G }
func (c RGBA[T]) BGRB() (T, T, T, T) { return c.B, c.G, c.R, c.B }
func (c RGBA[T]) BGRA() (T, T, T, T) { return c.B, c.G, c.R, c.A }
func (c RGBA[T]) BGGR() (T, T, T, T) { return c.B, c.G, c.G, c.R }
func (c RGBA[T]) BGGG() (T, T, T, T) { return c.B, c.G, c.G, c.G }
func (c RGBA[T]) BGGB() (T, T, T, T) { return c.B, c.G, c.G, c.B }
func (c RGBA[T]) BGGA() (T, T, T, T) { return c.B, c.G, c.G, c.A }
func (c RGBA[T]) BGBR() (T, T, T, T) { return c.B, c.G, c.B, c.R }
func (c RGBA[T]) BGBG() (T, T, T, T) { return c.B, c.G, c.B, c.G }
func (c RGBA[T]) BGBB() (T, T, T, T) { return c.B, c.G, c.B, c.B }
func (c RGBA[T]) BGBA() (T, T, T, T) { return c.B, c.G, c.B, c.A }
func (c RGBA[T]) BGAR() (T, T, T, T) { return c.B, c.G, c.A, c.R }
func (c RGBA[T]) BGAG() (T, T, T, T) { return c.B, c.G, c.A, c.G }
func (c RGBA[T]) BGAB() (T, T, T, T) { return c.B, c.G, c.A, c.B }
func (c RGBA[T]) BGAA() (T, T, T, T) { return c.B, c.G, c.A, c.A }
func (c RGBA[T]) BBRR() (T, T, T, T) { return c.B, c.B, c.R, c.R }
func (c RGBA[T]) BBRG() (T, T, T, T) { return c.B, c.B, c.R, c.G }
func (c RGBA[T]) BBRB() (T, T, T, T) { return c.B, c.B, c.R, c.B }
func (c RGBA[T]) BBRA() (T, T, T, T) { return c.B, c.B, c.R, c.A }
func (c RGBA[T]) BBGR() (T, T, T, T) { return c.B, c.B, c.G, c.R }
func (c RGBA[T]) BBGG() (T, T, T, T) { return c.B, c.B, c.G, c.G }
func (c RGBA[T]) BBGB() (T, T, T, T) { return c.B, c.B, c.G, c.B }
func (c RGBA[T]) BBGA() (T, T, T, T) { return c.B, c.B, c.G, c.A }
func (c RGBA[T]) BBBR() (T, T, T, T) { return c.B, c.B, c.B, c.R }
func (c RGBA[T]) BBBG() (T, T, T, T) { return c.B, c.B, c.B, c.G }
func (c RGBA[T]) BBBB() (T, T, T, T) { return c.B, c.B, c.B, c.B }
func (c RGBA[T]) BBBA() (T, T, T, T) { return c.B, c.B, c.B, c.A }
func (c RGBA[T]) BBAR() (T, T, T, T) { return c.B, c.B, c.A, c.R }
func (c RGBA[T]) BBAG() (T, T, T, T) { return c.B, c.B, c.A, c.G }
func (c RGBA[T]) BBAB() (T, T, T, T) { return c.B, c.B, c.A, c.B }
func (c RGBA[T]) BBAA() (T, T, T, T) { return c.B, c.B, c.A, c.A }
func (c RGBA[T]) BARR() (T, T, T, T) { return c.B, c.A, c.R, c.R }
func (c RGBA[T]) BARG() (T, T, T, T) { return c.B, c.A, c.R, c.G }
func (c RGBA[T]) BARB() (T, T, T, T) { return c.B, c.A, c.R, c.B }
func (c RGBA[T]) BARA() (T, T, T, T) { return c.B, c.A, c.R, c.A }
func (c RGBA[T]) BAGR() (T, T, T, T) { return c.B, c.A, c.G, c.R }
func (c RGBA[T]) BAGG() (T, T, T, T) { return c.B, c.A, c.G, c.G }
func (c RGBA[T]) BAGB() (T, T, T, T) { return c.B, c.A, c.G, c.B }
func (c RGBA[T]) BAGA() (T, T, T, T) { return c.B, c.A, c.G, c.A }
func (c RGBA[T]) BABR() (T, T, T, T) { return c.B, c.A, c.B, c.R }
func (c RGBA[T]) BABG() (T, T, T, T) { return c.B, c.A, c.B, c.G }
func (c RGBA[T]) BABB() (T, T, T, T) { return c.B, c.A, c.B, c.B }
func (c RGBA[T]) BABA() (T, T, T, T) { return c.B, c.A, c.B, c.A }
func (c RGBA[T]) BAAR() (T, T, T, T) { return c.B, c.A, c.A, c.R }
func (c RGBA[T]) BAAG() (T, T, T, T) { return c.B, c.A, c.A, c.G }
func (c RGBA[T]) BAAB() (T, T, T, T) { return c.B, c.A, c.A, c.B }
func (c RGBA[T]) BAAA() (T, T, T, T) { return c.B, c.A, c.A, c.A }
func (c RGBA[T]) ARRR() (T, T, T, T) { return c.A, c.R, c.R, c.R }
func (c RGBA[T]) ARRG() (T, T, T, T) { return c.A, c.R, c.R, c.G }
func (c RGBA[T]) ARRB() (T, T, T, T) { return c.A, c.R, c.R, c.B }
func (c RGBA[T]) ARRA() (T, T, T, T) { return c.A, c.R, c.R, c.A }
func (c RGBA[T]) ARGR() (T, T, T, T) { return c.A, c.R, c.G, c.R }
func (c RGBA[T]) ARGG() (T, T, T, T) { return c.A, c.R, c.G, c.G }
func (c RGBA[T]) ARGB() (T, T, T, T) { return c.A, c.R, c.G, c.B }
func (c RGBA[T]) ARGA() (T, T, T, T) { return c.A, c.R, c.G, c.A }
func (c RGBA[T]) ARBR() (T, T, T, T) { return c.A, c.R, c.B, c.R }
func (c RGBA[T]) ARBG() (T, T, T, T) { return c.A, c.R, c.B, c.G }
func (c RGBA[T]) ARBB() (T, T, T, T) { return c.A, c.R, c.B, c.B }
func (c RGBA[T]) ARBA() (T, T, T, T) { return c.A, c.R, c.B, c.A }
func (c RGBA[T]) ARAR() (T, T, T, T) { return c.A, c.R, c.A, c.R }
func (c RGBA[T]) ARAG() (T, T, T, T) { return c.A, c.R, c.A, c.G }
func (c RGBA[T]) ARAB() (T, T, T, T) { return c.A, c.R, c.A, c.B }
func (c RGBA[T]) ARAA() (T, T, T, T) { return c.A, c.R, c.A, c.A }
func (c RGBA[T]) AGRR() (T, T, T, T) { return c.A, c.G, c.R, c.R }
func (c RGBA[T]) AGRG() (T, T, T, T) { return c.A, c.G, c.R, c.G }
func (c RGBA[T]) AGRB() (T, T, T, T) { return c.A, c.G, c.R, c.B }
func (c RGBA[T]) AGRA() (T, T, T, T) { return c.A, c.G, c.R, c.A }
func (c RGBA[T]) AGGR() (T, T, T, T) { return c.A, c.G, c.G, c.R }
func (c RGBA[T]) AGGG() (T, T, T, T) { return c.A, c.G, c.G, c.G }
func (c RGBA[T]) AGGB() (T, T, T, T) { return c.A, c.G, c.G, c.B }
func (c RGBA[T]) AGGA() (T, T, T, T) { return c.A, c.G, c.G, c.A }
func (c RGBA[T]) AGBR() (T, T, T, T) { return c.A, c.G, c.B, c.R }
func (c RGBA[T]) AGBG() (T, T, T, T) { return c.A, c.G, c.B, c.G }
func (c RGBA[T]) AGBB() (T, T, T, T) { return c.A, c.G, c.B, c.B }
func (c RGBA[T]) AGBA() (T, T, T, T) { return c.A, c.G, c.B, c.A }
func (c RGBA[T]) AGAR() (T, T, T, T) { return c.A, c.G, c.A, c.R }
func (c RGBA[T]) AGAG() (T, T, T, T) { return c.A, c.G, c.A, c.G }
func (c RGBA[T]) AGAB() (T, T, T, T) { return c.A, c.G, c.A, c.B }
func (c RGBA[T]) AGAA() (T, T, T, T) { return c.A, c.G, c.A, c.A }
func (c RGBA[T]) ABRR() (T, T, T, T) { return c.A, c.B, c.R, c.R }
func (c RGBA[T]) ABRG() (T, T, T, T) { return c.A, c.B, c.R, c.G }
func (c RGBA[T]) ABRB() (T, T, T, T) { return c.A, c.B, c.R, c.B }
func (c RGBA[T]) ABRA() (T, T, T, T) { return c.A, c.B, c.R, c.A }
func (c RGBA[T]) ABGR() (T, T, T, T) { return c.A, c.B, c.G, c.R }
func (c RGBA[T]) ABGG() (T, T, T, T) { return c.A, c.B, c.G, c.G }
func (c RGBA[T]) ABGB() (T, T, T, T) { return c.A, c.B, c.G, c.B }
func (c RGBA[T]) ABGA() (T, T, T, T) { return c.A, c.B, c.G, c.A }
func (c RGBA[T]) ABBR() (T, T, T, T) { return c.A, c.B, c.B, c.R }
func (c RGBA[T]) ABBG() (T, T, T, T) { return c.A, c.B, c.B, c.G }
func (c RGBA[T]) ABBB() (T, T, T, T) { return c.A, c.B, c.B, c.B }
func (c RGBA[T]) ABBA() (T, T, T, T) { return c.A, c.B, c.B, c.A }
func (c RGBA[T]) ABAR() (T, T, T, T) { return c.A, c.B, c.A, c.R }
func (c RGBA[T]) ABAG() (T, T, T, T) { return c.A, c.B, c.A, c.G }
func (c RGBA[T]) ABAB() (T, T, T, T) { return c.A, c.B, c.A, c.B }
func (c RGBA[T]) ABAA() (T, T, T, T) { return c.A, c.B, c.A, c.A }
func (c RGBA[T]) AARR() (T, T, T, T) { return c.A, c.A, c.R, c.R }
func (c RGBA[T]) AARG() (T, T, T, T) { return c.A, c.A, c.R, c.G }
func (c RGBA[T]) AARB() (T, T, T, T) { return c.A, c.A, c.R, c.B }
func (c RGBA[T]) AARA() (T, T, T, T) { return c.A, c.A, c.R, c.A }
func (c RGBA[T]) AAGR() (T, T, T, T) { return c.A, c.A, c.G, c.R }
func (c RGBA[T]) AAGG() (T, T, T, T) { return c.A, c.A, c.G, c.G }
func (c RGBA[T]) AAGB() (T, T, T, T) { return c.A, c.A, c.G, c.B }
func (c RGBA[T]) AAGA() (T, T, T, T) { return c.A, c.A, c.G, c.A }
func (c RGBA[T]) AABR() (T, T, T, T) { return c.A, c.A, c.B, c.R }
func (c RGBA[T]) AABG() (T, T, T, T) { return c.A, c.A, c.B, c.G }
func (c RGBA[T]) AABB() (T, T, T, T) { return c.A, c.A, c.B, c.B }
func (c RGBA[T]) AABA() (T, T, T, T) { return c.A, c.A, c.B, c.A }
func (c RGBA[T]) AAAR() (T, T, T, T) { return c.A, c.A, c.A, c.R }
func (c RGBA[T]) AAAG() (T, T, T, T) { return c.A, c.A, c.A, c.G }
func (c RGBA[T]) AAAB() (T, T, T, T) { return c.A, c.A, c.A, c.B }
func (c RGBA[T]) AAAA() (T, T, T, T) { return c.A, c.A, c.A, c.A }
