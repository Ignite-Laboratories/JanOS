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

func (v RGBTyped[TR, TG, TB]) R() Cursor[TR] {
	return v.vector.components.x
}

func (v RGBTyped[TR, TG, TB]) SetR(value TR) RGBTyped[TR, TG, TB] {
	v.vector = v.vector.SetX(value)
	return v
}

func (v RGBTyped[TR, TG, TB]) G() Cursor[TG] {
	return v.vector.components.y
}

func (v RGBTyped[TR, TG, TB]) SetG(value TG) RGBTyped[TR, TG, TB] {
	v.vector = v.vector.SetY(value)
	return v
}

func (v RGBTyped[TR, TG, TB]) B() Cursor[TB] {
	return v.vector.components.z
}

func (v RGBTyped[TR, TG, TB]) SetB(value TB) RGBTyped[TR, TG, TB] {
	v.vector = v.vector.SetZ(value)
	return v
}
