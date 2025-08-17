package std

import (
	"github.com/ignite-laboratories/core/std/name/format"
	"github.com/ignite-laboratories/core/std/num"
)

// RGB is a kind of Vector3D that provides R G B mappings to the underlying component vectors.
type RGB[T num.Primitive] = RGBTyped[T, T, T]

func NewRGB[T num.Primitive]() RGB[T] {
	maximum := num.MaxValue[T]()
	c := RGB[T]{}
	c.vector.Entity = NewEntity[format.Default]()
	c.vector = c.vector.SetBoundaries(0, maximum, 0, maximum, 0, maximum)
	c.vector = c.vector.MapNames("R", "G", "B")
	return c
}

// RGBTyped is a kind of VectorTyped3D that provides R G B mappings to the underlying component vectors.
type RGBTyped[TR num.Primitive, TG num.Primitive, TB num.Primitive] struct {
	vector VectorTyped3D[TR, TG, TB]
}

func NewRGBTyped[TR num.Primitive, TG num.Primitive, TB num.Primitive]() RGBTyped[TR, TG, TB] {
	c := RGBTyped[TR, TG, TB]{}
	c.vector.Entity = NewEntity[format.Default]()
	return c
}

func (v RGBTyped[TR, TG, TB]) String() string {
	return v.vector.String()
}

func (v RGBTyped[TR, TG, TB]) SetClamp(clamp bool) RGBTyped[TR, TG, TB] {
	v.vector = v.vector.SetClamp(clamp)
	return v
}

func (v RGBTyped[TR, TG, TB]) SetBoundaries(minR, maxR TR, minG, maxG TG, minB, maxB TB) RGBTyped[TR, TG, TB] {
	v.vector = v.vector.SetBoundaries(minR, maxR, minG, maxG, minB, maxB)
	return v
}

func (v RGBTyped[TR, TG, TB]) Set(r TR, g TG, b TB) RGBTyped[TR, TG, TB] {
	v.vector = v.vector.SetValues(r, g, b)
	return v
}

func (v RGBTyped[TR, TG, TB]) R() TR {
	return v.vector.components.x.Value()
}

func (v RGBTyped[TR, TG, TB]) SetR(value TR) RGBTyped[TR, TG, TB] {
	v.vector = v.vector.SetX(value)
	return v
}

func (v RGBTyped[TR, TG, TB]) G() TG {
	return v.vector.components.y.Value()
}

func (v RGBTyped[TR, TG, TB]) SetG(value TG) RGBTyped[TR, TG, TB] {
	v.vector = v.vector.SetY(value)
	return v
}

func (v RGBTyped[TR, TG, TB]) B() TB {
	return v.vector.components.z.Value()
}

func (v RGBTyped[TR, TG, TB]) SetB(value TB) RGBTyped[TR, TG, TB] {
	v.vector = v.vector.SetZ(value)
	return v
}

/**
Swizzling

NOTE: This is a regular expression to find and replace swizzle functions into a one-liner if the auto formatter ever kicks in

Find -
func \((.*?)\) ([A-Z]{2,4})\(\) \((.*?)\)[ ]*\{[\n\t ]*return(.*?)[\n\t ]*\}

Replace -
func ($1) $2() ($3) { return$4 }
*/

func (c RGBTyped[TR, TG, TB]) RR() (TR, TR) { return c.R(), c.R() }
func (c RGBTyped[TR, TG, TB]) RG() (TR, TG) { return c.R(), c.G() }
func (c RGBTyped[TR, TG, TB]) RB() (TR, TB) { return c.R(), c.B() }
func (c RGBTyped[TR, TG, TB]) GR() (TG, TR) { return c.G(), c.R() }
func (c RGBTyped[TR, TG, TB]) GG() (TG, TG) { return c.G(), c.G() }
func (c RGBTyped[TR, TG, TB]) GB() (TG, TB) { return c.G(), c.B() }
func (c RGBTyped[TR, TG, TB]) BR() (TB, TR) { return c.B(), c.R() }
func (c RGBTyped[TR, TG, TB]) BG() (TB, TG) { return c.B(), c.G() }
func (c RGBTyped[TR, TG, TB]) BB() (TB, TB) { return c.B(), c.B() }

func (c RGBTyped[TR, TG, TB]) RRR() (TR, TR, TR) { return c.R(), c.R(), c.R() }
func (c RGBTyped[TR, TG, TB]) RRG() (TR, TR, TG) { return c.R(), c.R(), c.G() }
func (c RGBTyped[TR, TG, TB]) RRB() (TR, TR, TB) { return c.R(), c.R(), c.B() }
func (c RGBTyped[TR, TG, TB]) RGR() (TR, TG, TR) { return c.R(), c.G(), c.R() }
func (c RGBTyped[TR, TG, TB]) RGG() (TR, TG, TG) { return c.R(), c.G(), c.G() }
func (c RGBTyped[TR, TG, TB]) RGB() (TR, TG, TB) { return c.R(), c.G(), c.B() }
func (c RGBTyped[TR, TG, TB]) RBR() (TR, TB, TR) { return c.R(), c.B(), c.R() }
func (c RGBTyped[TR, TG, TB]) RBG() (TR, TB, TG) { return c.R(), c.B(), c.G() }
func (c RGBTyped[TR, TG, TB]) RBB() (TR, TB, TB) { return c.R(), c.B(), c.B() }
func (c RGBTyped[TR, TG, TB]) GRR() (TG, TR, TR) { return c.G(), c.R(), c.R() }
func (c RGBTyped[TR, TG, TB]) GRG() (TG, TR, TG) { return c.G(), c.R(), c.G() }
func (c RGBTyped[TR, TG, TB]) GRB() (TG, TR, TB) { return c.G(), c.R(), c.B() }
func (c RGBTyped[TR, TG, TB]) GGR() (TG, TG, TR) { return c.G(), c.G(), c.R() }
func (c RGBTyped[TR, TG, TB]) GGG() (TG, TG, TG) { return c.G(), c.G(), c.G() }
func (c RGBTyped[TR, TG, TB]) GGB() (TG, TG, TB) { return c.G(), c.G(), c.B() }
func (c RGBTyped[TR, TG, TB]) GBR() (TG, TB, TR) { return c.G(), c.B(), c.R() }
func (c RGBTyped[TR, TG, TB]) GBG() (TG, TB, TG) { return c.G(), c.B(), c.G() }
func (c RGBTyped[TR, TG, TB]) GBB() (TG, TB, TB) { return c.G(), c.B(), c.B() }
func (c RGBTyped[TR, TG, TB]) BRR() (TB, TR, TR) { return c.B(), c.R(), c.R() }
func (c RGBTyped[TR, TG, TB]) BRG() (TB, TR, TG) { return c.B(), c.R(), c.G() }
func (c RGBTyped[TR, TG, TB]) BRB() (TB, TR, TB) { return c.B(), c.R(), c.B() }
func (c RGBTyped[TR, TG, TB]) BGR() (TB, TG, TR) { return c.B(), c.G(), c.R() }
func (c RGBTyped[TR, TG, TB]) BGG() (TB, TG, TG) { return c.B(), c.G(), c.G() }
func (c RGBTyped[TR, TG, TB]) BGB() (TB, TG, TB) { return c.B(), c.G(), c.B() }
func (c RGBTyped[TR, TG, TB]) BBR() (TB, TB, TR) { return c.B(), c.B(), c.R() }
func (c RGBTyped[TR, TG, TB]) BBG() (TB, TB, TG) { return c.B(), c.B(), c.G() }
func (c RGBTyped[TR, TG, TB]) BBB() (TB, TB, TB) { return c.B(), c.B(), c.B() }

func (c RGBTyped[TR, TG, TB]) RRRR() (TR, TR, TR, TR) { return c.R(), c.R(), c.R(), c.R() }
func (c RGBTyped[TR, TG, TB]) RRRG() (TR, TR, TR, TG) { return c.R(), c.R(), c.R(), c.G() }
func (c RGBTyped[TR, TG, TB]) RRRB() (TR, TR, TR, TB) { return c.R(), c.R(), c.R(), c.B() }
func (c RGBTyped[TR, TG, TB]) RRGR() (TR, TR, TG, TR) { return c.R(), c.R(), c.G(), c.R() }
func (c RGBTyped[TR, TG, TB]) RRGG() (TR, TR, TG, TG) { return c.R(), c.R(), c.G(), c.G() }
func (c RGBTyped[TR, TG, TB]) RRGB() (TR, TR, TG, TB) { return c.R(), c.R(), c.G(), c.B() }
func (c RGBTyped[TR, TG, TB]) RRBR() (TR, TR, TB, TR) { return c.R(), c.R(), c.B(), c.R() }
func (c RGBTyped[TR, TG, TB]) RRBG() (TR, TR, TB, TG) { return c.R(), c.R(), c.B(), c.G() }
func (c RGBTyped[TR, TG, TB]) RRBB() (TR, TR, TB, TB) { return c.R(), c.R(), c.B(), c.B() }
func (c RGBTyped[TR, TG, TB]) RGRR() (TR, TG, TR, TR) { return c.R(), c.G(), c.R(), c.R() }
func (c RGBTyped[TR, TG, TB]) RGRG() (TR, TG, TR, TG) { return c.R(), c.G(), c.R(), c.G() }
func (c RGBTyped[TR, TG, TB]) RGRB() (TR, TG, TR, TB) { return c.R(), c.G(), c.R(), c.B() }
func (c RGBTyped[TR, TG, TB]) RGGR() (TR, TG, TG, TR) { return c.R(), c.G(), c.G(), c.R() }
func (c RGBTyped[TR, TG, TB]) RGGG() (TR, TG, TG, TG) { return c.R(), c.G(), c.G(), c.G() }
func (c RGBTyped[TR, TG, TB]) RGGB() (TR, TG, TG, TB) { return c.R(), c.G(), c.G(), c.B() }
func (c RGBTyped[TR, TG, TB]) RGBR() (TR, TG, TB, TR) { return c.R(), c.G(), c.B(), c.R() }
func (c RGBTyped[TR, TG, TB]) RGBG() (TR, TG, TB, TG) { return c.R(), c.G(), c.B(), c.G() }
func (c RGBTyped[TR, TG, TB]) RGBB() (TR, TG, TB, TB) { return c.R(), c.G(), c.B(), c.B() }
func (c RGBTyped[TR, TG, TB]) RBRR() (TR, TB, TR, TR) { return c.R(), c.B(), c.R(), c.R() }
func (c RGBTyped[TR, TG, TB]) RBRG() (TR, TB, TR, TG) { return c.R(), c.B(), c.R(), c.G() }
func (c RGBTyped[TR, TG, TB]) RBRB() (TR, TB, TR, TB) { return c.R(), c.B(), c.R(), c.B() }
func (c RGBTyped[TR, TG, TB]) RBGR() (TR, TB, TG, TR) { return c.R(), c.B(), c.G(), c.R() }
func (c RGBTyped[TR, TG, TB]) RBGG() (TR, TB, TG, TG) { return c.R(), c.B(), c.G(), c.G() }
func (c RGBTyped[TR, TG, TB]) RBGB() (TR, TB, TG, TB) { return c.R(), c.B(), c.G(), c.B() }
func (c RGBTyped[TR, TG, TB]) RBBR() (TR, TB, TB, TR) { return c.R(), c.B(), c.B(), c.R() }
func (c RGBTyped[TR, TG, TB]) RBBG() (TR, TB, TB, TG) { return c.R(), c.B(), c.B(), c.G() }
func (c RGBTyped[TR, TG, TB]) RBBB() (TR, TB, TB, TB) { return c.R(), c.B(), c.B(), c.B() }
func (c RGBTyped[TR, TG, TB]) GRRR() (TG, TR, TR, TR) { return c.G(), c.R(), c.R(), c.R() }
func (c RGBTyped[TR, TG, TB]) GRRG() (TG, TR, TR, TG) { return c.G(), c.R(), c.R(), c.G() }
func (c RGBTyped[TR, TG, TB]) GRRB() (TG, TR, TR, TB) { return c.G(), c.R(), c.R(), c.B() }
func (c RGBTyped[TR, TG, TB]) GRGR() (TG, TR, TG, TR) { return c.G(), c.R(), c.G(), c.R() }
func (c RGBTyped[TR, TG, TB]) GRGG() (TG, TR, TG, TG) { return c.G(), c.R(), c.G(), c.G() }
func (c RGBTyped[TR, TG, TB]) GRGB() (TG, TR, TG, TB) { return c.G(), c.R(), c.G(), c.B() }
func (c RGBTyped[TR, TG, TB]) GRBR() (TG, TR, TB, TR) { return c.G(), c.R(), c.B(), c.R() }
func (c RGBTyped[TR, TG, TB]) GRBG() (TG, TR, TB, TG) { return c.G(), c.R(), c.B(), c.G() }
func (c RGBTyped[TR, TG, TB]) GRBB() (TG, TR, TB, TB) { return c.G(), c.R(), c.B(), c.B() }
func (c RGBTyped[TR, TG, TB]) GGRR() (TG, TG, TR, TR) { return c.G(), c.G(), c.R(), c.R() }
func (c RGBTyped[TR, TG, TB]) GGRG() (TG, TG, TR, TG) { return c.G(), c.G(), c.R(), c.G() }
func (c RGBTyped[TR, TG, TB]) GGRB() (TG, TG, TR, TB) { return c.G(), c.G(), c.R(), c.B() }
func (c RGBTyped[TR, TG, TB]) GGGR() (TG, TG, TG, TR) { return c.G(), c.G(), c.G(), c.R() }
func (c RGBTyped[TR, TG, TB]) GGGG() (TG, TG, TG, TG) { return c.G(), c.G(), c.G(), c.G() }
func (c RGBTyped[TR, TG, TB]) GGGB() (TG, TG, TG, TB) { return c.G(), c.G(), c.G(), c.B() }
func (c RGBTyped[TR, TG, TB]) GGBR() (TG, TG, TB, TR) { return c.G(), c.G(), c.B(), c.R() }
func (c RGBTyped[TR, TG, TB]) GGBG() (TG, TG, TB, TG) { return c.G(), c.G(), c.B(), c.G() }
func (c RGBTyped[TR, TG, TB]) GGBB() (TG, TG, TB, TB) { return c.G(), c.G(), c.B(), c.B() }
func (c RGBTyped[TR, TG, TB]) GBRR() (TG, TB, TR, TR) { return c.G(), c.B(), c.R(), c.R() }
func (c RGBTyped[TR, TG, TB]) GBRG() (TG, TB, TR, TG) { return c.G(), c.B(), c.R(), c.G() }
func (c RGBTyped[TR, TG, TB]) GBRB() (TG, TB, TR, TB) { return c.G(), c.B(), c.R(), c.B() }
func (c RGBTyped[TR, TG, TB]) GBGR() (TG, TB, TG, TR) { return c.G(), c.B(), c.G(), c.R() }
func (c RGBTyped[TR, TG, TB]) GBGG() (TG, TB, TG, TG) { return c.G(), c.B(), c.G(), c.G() }
func (c RGBTyped[TR, TG, TB]) GBGB() (TG, TB, TG, TB) { return c.G(), c.B(), c.G(), c.B() }
func (c RGBTyped[TR, TG, TB]) GBBR() (TG, TB, TB, TR) { return c.G(), c.B(), c.B(), c.R() }
func (c RGBTyped[TR, TG, TB]) GBBG() (TG, TB, TB, TG) { return c.G(), c.B(), c.B(), c.G() }
func (c RGBTyped[TR, TG, TB]) GBBB() (TG, TB, TB, TB) { return c.G(), c.B(), c.B(), c.B() }
func (c RGBTyped[TR, TG, TB]) BRRR() (TB, TR, TR, TR) { return c.B(), c.R(), c.R(), c.R() }
func (c RGBTyped[TR, TG, TB]) BRRG() (TB, TR, TR, TG) { return c.B(), c.R(), c.R(), c.G() }
func (c RGBTyped[TR, TG, TB]) BRRB() (TB, TR, TR, TB) { return c.B(), c.R(), c.R(), c.B() }
func (c RGBTyped[TR, TG, TB]) BRGR() (TB, TR, TG, TR) { return c.B(), c.R(), c.G(), c.R() }
func (c RGBTyped[TR, TG, TB]) BRGG() (TB, TR, TG, TG) { return c.B(), c.R(), c.G(), c.G() }
func (c RGBTyped[TR, TG, TB]) BRGB() (TB, TR, TG, TB) { return c.B(), c.R(), c.G(), c.B() }
func (c RGBTyped[TR, TG, TB]) BRBR() (TB, TR, TB, TR) { return c.B(), c.R(), c.B(), c.R() }
func (c RGBTyped[TR, TG, TB]) BRBG() (TB, TR, TB, TG) { return c.B(), c.R(), c.B(), c.G() }
func (c RGBTyped[TR, TG, TB]) BRBB() (TB, TR, TB, TB) { return c.B(), c.R(), c.B(), c.B() }
func (c RGBTyped[TR, TG, TB]) BGRR() (TB, TG, TR, TR) { return c.B(), c.G(), c.R(), c.R() }
func (c RGBTyped[TR, TG, TB]) BGRG() (TB, TG, TR, TG) { return c.B(), c.G(), c.R(), c.G() }
func (c RGBTyped[TR, TG, TB]) BGRB() (TB, TG, TR, TB) { return c.B(), c.G(), c.R(), c.B() }
func (c RGBTyped[TR, TG, TB]) BGGR() (TB, TG, TG, TR) { return c.B(), c.G(), c.G(), c.R() }
func (c RGBTyped[TR, TG, TB]) BGGG() (TB, TG, TG, TG) { return c.B(), c.G(), c.G(), c.G() }
func (c RGBTyped[TR, TG, TB]) BGGB() (TB, TG, TG, TB) { return c.B(), c.G(), c.G(), c.B() }
func (c RGBTyped[TR, TG, TB]) BGBR() (TB, TG, TB, TR) { return c.B(), c.G(), c.B(), c.R() }
func (c RGBTyped[TR, TG, TB]) BGBG() (TB, TG, TB, TG) { return c.B(), c.G(), c.B(), c.G() }
func (c RGBTyped[TR, TG, TB]) BGBB() (TB, TG, TB, TB) { return c.B(), c.G(), c.B(), c.B() }
func (c RGBTyped[TR, TG, TB]) BBRR() (TB, TB, TR, TR) { return c.B(), c.B(), c.R(), c.R() }
func (c RGBTyped[TR, TG, TB]) BBRG() (TB, TB, TR, TG) { return c.B(), c.B(), c.R(), c.G() }
func (c RGBTyped[TR, TG, TB]) BBRB() (TB, TB, TR, TB) { return c.B(), c.B(), c.R(), c.B() }
func (c RGBTyped[TR, TG, TB]) BBGR() (TB, TB, TG, TR) { return c.B(), c.B(), c.G(), c.R() }
func (c RGBTyped[TR, TG, TB]) BBGG() (TB, TB, TG, TG) { return c.B(), c.B(), c.G(), c.G() }
func (c RGBTyped[TR, TG, TB]) BBGB() (TB, TB, TG, TB) { return c.B(), c.B(), c.G(), c.B() }
func (c RGBTyped[TR, TG, TB]) BBBR() (TB, TB, TB, TR) { return c.B(), c.B(), c.B(), c.R() }
func (c RGBTyped[TR, TG, TB]) BBBG() (TB, TB, TB, TG) { return c.B(), c.B(), c.B(), c.G() }
func (c RGBTyped[TR, TG, TB]) BBBB() (TB, TB, TB, TB) { return c.B(), c.B(), c.B(), c.B() }

