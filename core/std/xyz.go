package std

import (
	"github.com/ignite-laboratories/core/std/num"
)

// XYZ is a kind of Vector3D that provides X Y Z mappings to the underlying component vectors.
type XYZ[T num.Primitive] = XYZTyped[T, T, T]

// XYZTyped is a kind of Vector3DTyped that provides X Y Z mappings to the underlying component vectors.
type XYZTyped[TX num.Primitive, TY num.Primitive, TZ num.Primitive] Vector3DTyped[TX, TY, TZ]

func (v XYZTyped[TX, TY, TZ]) SetClamp(clamp bool) XYZTyped[TX, TY, TZ] {
	return v.SetClamp(clamp)
}

func (v XYZTyped[TX, TY, TZ]) SetBoundaries(minX, maxX TX, minY, maxY TY, minZ, maxZ TZ) XYZTyped[TX, TY, TZ] {
	return v.SetBoundaries(minX, maxX, minY, maxY, minZ, maxZ)
}

func (v XYZTyped[TX, TY, TZ]) Set(x TX, y TY, z TZ) {
	_ = v.components.x.Set(x)
	_ = v.components.y.Set(y)
	_ = v.components.z.Set(z)
}

func (v XYZTyped[TX, TY, TZ]) X() Cursor[TX] {
	return v.components.x
}

func (v XYZTyped[TX, TY, TZ]) SetX(value TX) {
	_ = v.components.x.Set(value)
}

func (v XYZTyped[TX, TY, TZ]) Y() Cursor[TY] {
	return v.components.y
}

func (v XYZTyped[TX, TY, TZ]) SetY(value TY) {
	_ = v.components.y.Set(value)
}

func (v XYZTyped[TX, TY, TZ]) Z() Cursor[TZ] {
	return v.components.z
}

func (v XYZTyped[TX, TY, TZ]) SetZ(value TZ) {
	_ = v.components.z.Set(value)
}
