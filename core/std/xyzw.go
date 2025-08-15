package std

import (
	"github.com/ignite-laboratories/core/std/num"
)

// XYZW is a kind of Vector4D that provides X Y Z W mappings to the underlying component vectors.
type XYZW[T num.Primitive] = XYZWTyped[T, T, T, T]

// XYZWTyped is a kind of Vector4DTyped that provides X Y Z W mappings to the underlying component vectors.
type XYZWTyped[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive] Vector4DTyped[TX, TY, TZ, TW]

func (v XYZWTyped[TX, TY, TZ, TW]) SetClamp(clamp bool) XYZWTyped[TX, TY, TZ, TW] {
	return v.SetClamp(clamp)
}

func (v XYZWTyped[TX, TY, TZ, TW]) SetBoundaries(minX, maxX TX, minY, maxY TY, minZ, maxZ TZ, minW, maxW TW) XYZWTyped[TX, TY, TZ, TW] {
	return v.SetBoundaries(minX, maxX, minY, maxY, minZ, maxZ, minW, maxW)
}

func (v XYZWTyped[TX, TY, TZ, TW]) Set(x TX, y TY, z TZ, w TW) {
	_ = v.components.x.Set(x)
	_ = v.components.y.Set(y)
	_ = v.components.z.Set(z)
	_ = v.components.w.Set(w)
}

func (v XYZWTyped[TX, TY, TZ, TW]) X() Cursor[TX] {
	return v.components.x
}

func (v XYZWTyped[TX, TY, TZ, TW]) SetX(value TX) {
	_ = v.components.x.Set(value)
}

func (v XYZWTyped[TX, TY, TZ, TW]) Y() Cursor[TY] {
	return v.components.y
}

func (v XYZWTyped[TX, TY, TZ, TW]) SetY(value TY) {
	_ = v.components.y.Set(value)
}

func (v XYZWTyped[TX, TY, TZ, TW]) Z() Cursor[TZ] {
	return v.components.z
}

func (v XYZWTyped[TX, TY, TZ, TW]) SetZ(value TZ) {
	_ = v.components.z.Set(value)
}

func (v XYZWTyped[TX, TY, TZ, TW]) W() Cursor[TW] {
	return v.components.w
}

func (v XYZWTyped[TX, TY, TZ, TW]) SetW(value TW) {
	_ = v.components.w.Set(value)
}
