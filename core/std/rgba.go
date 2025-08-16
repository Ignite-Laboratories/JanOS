package std

import (
	"github.com/ignite-laboratories/core/std/num"
)

// RGBA is a kind of Vector4D that provides R G B A mappings to the underlying component vectors.
type RGBA[T num.Primitive] = RGBATyped[T, T, T, T]

// RGBATyped is a kind of VectorTyped4D that provides R G B A mappings to the underlying component vectors.
type RGBATyped[TR num.Primitive, TG num.Primitive, TB num.Primitive, TA num.Primitive] VectorTyped4D[TR, TG, TB, TA]

func (v RGBATyped[TR, TG, TB, TA]) SetClamp(clamp bool) RGBATyped[TR, TG, TB, TA] {
	return v.SetClamp(clamp)
}

func (v RGBATyped[TR, TG, TB, TA]) SetBoundaries(minR, maxR TR, minG, maxG TG, minB, maxB TB, minA, maxA TA) RGBATyped[TR, TG, TB, TA] {
	return v.SetBoundaries(minR, maxR, minG, maxG, minB, maxB, minA, maxA)
}

func (v RGBATyped[TR, TG, TB, TA]) Set(r TR, g TG, b TB, a TA) {
	_ = v.components.x.Set(r)
	_ = v.components.y.Set(g)
	_ = v.components.z.Set(b)
	_ = v.components.w.Set(a)
}

func (v RGBATyped[TR, TG, TB, TA]) R() Cursor[TR] {
	return v.components.x
}

func (v RGBATyped[TR, TG, TB, TA]) SetR(value TR) {
	_ = v.components.x.Set(value)
}

func (v RGBATyped[TR, TG, TB, TA]) G() Cursor[TG] {
	return v.components.y
}

func (v RGBATyped[TR, TG, TB, TA]) SetG(value TG) {
	_ = v.components.y.Set(value)
}

func (v RGBATyped[TR, TG, TB, TA]) B() Cursor[TB] {
	return v.components.z
}

func (v RGBATyped[TR, TG, TB, TA]) SetB(value TB) {
	_ = v.components.z.Set(value)
}

func (v RGBATyped[TR, TG, TB, TA]) A() Cursor[TA] {
	return v.components.w
}

func (v RGBATyped[TR, TG, TB, TA]) SetA(value TA) {
	_ = v.components.w.Set(value)
}
