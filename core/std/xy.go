package std

import (
	"github.com/ignite-laboratories/core/std/num"
)

// XY is a kind of Vector2D that provides X Y mappings to the underlying component vectors.
type XY[T num.Primitive] = XYTyped[T, T]

// XYTyped is a kind of VectorTyped2D that provides X Y mappings to the underlying component vectors.
type XYTyped[TX num.Primitive, TY num.Primitive] VectorTyped2D[TX, TY]

func (v XYTyped[TX, TY]) SetClamp(clamp bool) XYTyped[TX, TY] {
	return v.SetClamp(clamp)
}

func (v XYTyped[TX, TY]) SetBoundaries(minX, maxX TX, minY, maxY TY) XYTyped[TX, TY] {
	return v.SetBoundaries(minX, maxX, minY, maxY)
}

func (v XYTyped[TX, TY]) Set(x TX, y TY) {
	_ = v.components.x.Set(x)
	_ = v.components.y.Set(y)
}

func (v XYTyped[TX, TY]) X() TX {
	return v.components.x.Value()
}

func (v XYTyped[TX, TY]) SetX(value TX) {
	_ = v.components.x.Set(value)
}

func (v XYTyped[TX, TY]) Y() TY {
	return v.components.y.Value()
}

func (v XYTyped[TX, TY]) SetY(value TY) {
	_ = v.components.y.Set(value)
}

/**
Swizzling

NOTE: This is a regular expression to find and replace swizzle functions into a one-liner if the auto formatter ever kicks in

Find -
func \((.*?)\) ([A-Z]{2,4})\(\) \((.*?)\)[ ]*\{[\n\t ]*return(.*?)[\n\t ]*\}

Replace -
func ($1) $2() ($3) { return$4 }
*/

func (c XYTyped[TX, TY]) XX() (TX, TX) { return c.X(), c.X() }
func (c XYTyped[TX, TY]) XY() (TX, TY) { return c.X(), c.Y() }
func (c XYTyped[TX, TY]) YX() (TY, TX) { return c.Y(), c.X() }
func (c XYTyped[TX, TY]) YY() (TY, TY) { return c.Y(), c.Y() }
