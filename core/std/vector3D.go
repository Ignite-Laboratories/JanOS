package std

import "github.com/ignite-laboratories/core/std/num"

type components3D[TX num.Primitive, TY num.Primitive, TZ num.Primitive] struct {
	x Cursor[TX]
	y Cursor[TY]
	z Cursor[TZ]
}

// Vector3DTyped represents an asymmetric vector of seven dissimilar numeric cursors.
type Vector3DTyped[TX num.Primitive, TY num.Primitive, TZ num.Primitive] struct {
	components components3D[TX, TY, TZ]
}

// Vector3D represents a vector of seven like-typed numeric cursors.
type Vector3D[T num.Primitive] = Vector3DTyped[T, T, T]

func (v Vector3DTyped[TX, TY, TZ]) SetValues(x TX, y TY, z TZ) Vector3DTyped[TX, TY, TZ] {
	_ = v.components.x.Set(x)
	_ = v.components.y.Set(y)
	_ = v.components.z.Set(z)
	return v
}

func (v Vector3DTyped[TX, TY, TZ]) SetClamp(clamp bool) Vector3DTyped[TX, TY, TZ] {
	v.components.x.Clamp = clamp
	v.components.y.Clamp = clamp
	v.components.z.Clamp = clamp
	return v
}

func (v Vector3DTyped[TX, TY, TZ]) SetBoundaries(minX, maxX TX, minY, maxY TY, minZ, maxZ TZ) Vector3DTyped[TX, TY, TZ] {
	_ = v.components.x.SetBoundaries(minX, maxX)
	_ = v.components.y.SetBoundaries(minY, maxY)
	_ = v.components.z.SetBoundaries(minZ, maxZ)
	return v
}
