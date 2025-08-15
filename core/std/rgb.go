package std

import (
	"github.com/ignite-laboratories/core/std/num"
)

// RGB is a kind of Vector3D that provides R G B mappings to the underlying component vectors.
type RGB[T num.Primitive] = RGBTyped[T, T, T]

// RGBTyped is a kind of Vector3DTyped that provides R G B mappings to the underlying component vectors.
type RGBTyped[TR num.Primitive, TG num.Primitive, TB num.Primitive] Vector3DTyped[TR, TG, TB]

func (v RGBTyped[TR, TG, TB]) SetClamp(clamp bool) RGBTyped[TR, TG, TB] {
	return v.SetClamp(clamp)
}

func (v RGBTyped[TR, TG, TB]) SetBoundaries(minR, maxR TR, minG, maxG TG, minB, maxB TB) RGBTyped[TR, TG, TB] {
	return v.SetBoundaries(minR, maxR, minG, maxG, minB, maxB)
}

func (v RGBTyped[TR, TG, TB]) Set(r TR, g TG, b TB) {
	_ = v.components.x.Set(r)
	_ = v.components.y.Set(g)
	_ = v.components.z.Set(b)
}

func (v RGBTyped[TR, TG, TB]) R() Cursor[TR] {
	return v.components.x
}

func (v RGBTyped[TR, TG, TB]) SetR(value TR) {
	_ = v.components.x.Set(value)
}

func (v RGBTyped[TR, TG, TB]) G() Cursor[TG] {
	return v.components.y
}

func (v RGBTyped[TR, TG, TB]) SetG(value TG) {
	_ = v.components.y.Set(value)
}

func (v RGBTyped[TR, TG, TB]) B() Cursor[TB] {
	return v.components.z
}

func (v RGBTyped[TR, TG, TB]) SetB(value TB) {
	_ = v.components.z.Set(value)
}
