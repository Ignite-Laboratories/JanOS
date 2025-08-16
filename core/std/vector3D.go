package std

import (
	"fmt"
	format "github.com/ignite-laboratories/core/std/name/format"
	"github.com/ignite-laboratories/core/std/num"
	"github.com/ignite-laboratories/core/sys/support"
)

type components3D[TX num.Primitive, TY num.Primitive, TZ num.Primitive] struct {
	names [3]string
	x     Cursor[TX]
	y     Cursor[TY]
	z     Cursor[TZ]
}

// VectorTyped3D represents an asymmetric vector of seven dissimilar numeric cursors.
type VectorTyped3D[TX num.Primitive, TY num.Primitive, TZ num.Primitive] struct {
	Entity
	components components3D[TX, TY, TZ]
}

func NewVectorTyped3D[TX num.Primitive, TY num.Primitive, TZ num.Primitive]() VectorTyped3D[TX, TY, TZ] {
	return VectorTyped3D[TX, TY, TZ]{
		Entity: NewEntity[format.Default](),
	}
}

// Vector3D represents a vector of seven like-typed numeric cursors.
type Vector3D[T num.Primitive] = VectorTyped3D[T, T, T]

func NewVector3D[T num.Primitive]() Vector3D[T] {
	return Vector3D[T]{
		Entity: NewEntity[format.Default](),
	}
}

func (v VectorTyped3D[TX, TY, TZ]) String() string {
	if len(v.components.names[0]) == 0 {
		v.components.names = [3]string{"X", "Y", "Z"}
	}

	if support.AllSameTypes(TX(0), TY(0), TZ(0)) {
		return fmt.Sprintf("3D[%T].{%v: %v, %v: %v, %v: %v}(%v)", TX(0), v.components.names[0], v.components.x.ValueString(), v.components.names[1], v.components.y.ValueString(), v.components.names[2], v.components.z.ValueString(), v.GivenName.Name)
	}
	return fmt.Sprintf("3D[%T, %T, %T].{%v: %v, %v: %v, %v: %v}(%v)", TX(0), TY(0), TZ(0), v.components.names[0], v.components.x.ValueString(), v.components.names[1], v.components.y.ValueString(), v.components.names[2], v.components.z.ValueString(), v.GivenName.Name)
}

func (v VectorTyped3D[TX, TY, TZ]) SetValues(x TX, y TY, z TZ) VectorTyped3D[TX, TY, TZ] {
	_ = v.components.x.Set(x)
	_ = v.components.y.Set(y)
	_ = v.components.z.Set(z)
	return v
}

func (v VectorTyped3D[TX, TY, TZ]) SetX(x TX) VectorTyped3D[TX, TY, TZ] {
	_ = v.components.x.Set(x)
	return v
}

func (v VectorTyped3D[TX, TY, TZ]) SetY(y TY) VectorTyped3D[TX, TY, TZ] {
	_ = v.components.y.Set(y)
	return v
}

func (v VectorTyped3D[TX, TY, TZ]) SetZ(z TZ) VectorTyped3D[TX, TY, TZ] {
	_ = v.components.z.Set(z)
	return v
}

func (v VectorTyped3D[TX, TY, TZ]) SetClamp(clamp bool) VectorTyped3D[TX, TY, TZ] {
	v.components.x.Clamp = clamp
	v.components.y.Clamp = clamp
	v.components.z.Clamp = clamp
	return v
}

func (v VectorTyped3D[TX, TY, TZ]) SetBoundaries(minX, maxX TX, minY, maxY TY, minZ, maxZ TZ) VectorTyped3D[TX, TY, TZ] {
	_ = v.components.x.SetBoundaries(minX, maxX)
	_ = v.components.y.SetBoundaries(minY, maxY)
	_ = v.components.z.SetBoundaries(minZ, maxZ)
	return v
}

func (v VectorTyped3D[TX, TY, TZ]) MapNames(x, y, z string) VectorTyped3D[TX, TY, TZ] {
	v.components.names[0] = x
	v.components.names[1] = y
	v.components.names[2] = z
	return v
}
