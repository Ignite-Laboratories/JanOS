package std

import "github.com/ignite-laboratories/core/std/num"

type components4D[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive] struct {
	x Cursor[TX]
	y Cursor[TY]
	z Cursor[TZ]
	w Cursor[TW]
}

// Vector4DTyped represents an asymmetric vector of seven dissimilar numeric cursors.
type Vector4DTyped[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive] struct {
	components components4D[TX, TY, TZ, TW]
}

// Vector4D represents a vector of seven like-typed numeric cursors.
type Vector4D[T num.Primitive] = Vector4DTyped[T, T, T, T]

func (v Vector4DTyped[TX, TY, TZ, TW]) SetValues(x TX, y TY, z TZ, w TW) Vector4DTyped[TX, TY, TZ, TW] {
	_ = v.components.x.Set(x)
	_ = v.components.y.Set(y)
	_ = v.components.z.Set(z)
	_ = v.components.w.Set(w)
	return v
}

func (v Vector4DTyped[TX, TY, TZ, TW]) SetClamp(clamp bool) Vector4DTyped[TX, TY, TZ, TW] {
	v.components.x.Clamp = clamp
	v.components.y.Clamp = clamp
	v.components.z.Clamp = clamp
	v.components.w.Clamp = clamp
	return v
}

func (v Vector4DTyped[TX, TY, TZ, TW]) SetBoundaries(minX, maxX TX, minY, maxY TY, minZ, maxZ TZ, minW, maxW TW) Vector4DTyped[TX, TY, TZ, TW] {
	_ = v.components.x.SetBoundaries(minX, maxX)
	_ = v.components.y.SetBoundaries(minY, maxY)
	_ = v.components.z.SetBoundaries(minZ, maxZ)
	_ = v.components.w.SetBoundaries(minW, maxW)
	return v
}
