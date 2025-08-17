package std

import (
	"fmt"
	format "github.com/ignite-laboratories/core/std/name/format"
	"github.com/ignite-laboratories/core/std/num"
	"github.com/ignite-laboratories/core/sys/support"
)

type components4D[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive] struct {
	names [4]string
	x     Bounded[TX]
	y     Bounded[TY]
	z     Bounded[TZ]
	w     Bounded[TW]
}

// VectorTyped4D represents an asymmetric vector of seven dissimilar numeric cursors.
type VectorTyped4D[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive] struct {
	Entity
	components components4D[TX, TY, TZ, TW]
}

func NewVectorTyped4D[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive]() VectorTyped4D[TX, TY, TZ, TW] {
	return VectorTyped4D[TX, TY, TZ, TW]{
		Entity: NewEntity[format.Default](),
	}
}

func (v VectorTyped4D[TX, TY, TZ, TW]) String() string {
	if len(v.components.names[0]) == 0 {
		v.components.names = [4]string{"X", "Y", "Z", "W"}
	}

	if support.AllSameTypes(TX(0), TY(0), TZ(0), TW(0)) {
		return fmt.Sprintf("4D[%T].{%v: %v, %v: %v, %v: %v, %v: %v}(%v)", TX(0), v.components.names[0], v.components.x.ValueString(), v.components.names[1], v.components.y.ValueString(), v.components.names[2], v.components.z.ValueString(), v.components.names[3], v.components.w.ValueString(), v.GivenName.Name)
	}
	return fmt.Sprintf("4D[%T, %T, %T, %T].{%v: %v, %v: %v, %v: %v, %v: %v}(%v)", TX(0), TY(0), TZ(0), TW(0), v.components.names[0], v.components.x.ValueString(), v.components.names[1], v.components.y.ValueString(), v.components.names[2], v.components.z.ValueString(), v.components.names[3], v.components.w.ValueString(), v.GivenName.Name)
}

// Vector4D represents a vector of seven like-typed numeric cursors.
type Vector4D[T num.Primitive] = VectorTyped4D[T, T, T, T]

func NewVector4D[T num.Primitive]() Vector4D[T] {
	return Vector4D[T]{
		Entity: NewEntity[format.Default](),
	}
}

func (v VectorTyped4D[TX, TY, TZ, TW]) SetValues(x TX, y TY, z TZ, w TW) VectorTyped4D[TX, TY, TZ, TW] {
	_ = v.components.x.Set(x)
	_ = v.components.y.Set(y)
	_ = v.components.z.Set(z)
	_ = v.components.w.Set(w)
	return v
}

func (v VectorTyped4D[TX, TY, TZ, TW]) GetX() TX {
	return v.components.x.Value()
}

func (v VectorTyped4D[TX, TY, TZ, TW]) GetY() TY {
	return v.components.y.Value()
}

func (v VectorTyped4D[TX, TY, TZ, TW]) GetZ() TZ {
	return v.components.z.Value()
}

func (v VectorTyped4D[TX, TY, TZ, TW]) GetW() TW {
	return v.components.w.Value()
}

func (v VectorTyped4D[TX, TY, TZ, TW]) SetX(x TX) VectorTyped4D[TX, TY, TZ, TW] {
	_ = v.components.x.Set(x)
	return v
}

func (v VectorTyped4D[TX, TY, TZ, TW]) SetY(y TY) VectorTyped4D[TX, TY, TZ, TW] {
	_ = v.components.y.Set(y)
	return v
}

func (v VectorTyped4D[TX, TY, TZ, TW]) SetZ(z TZ) VectorTyped4D[TX, TY, TZ, TW] {
	_ = v.components.z.Set(z)
	return v
}

func (v VectorTyped4D[TX, TY, TZ, TW]) SetW(w TW) VectorTyped4D[TX, TY, TZ, TW] {
	_ = v.components.w.Set(w)
	return v
}

func (v VectorTyped4D[TX, TY, TZ, TW]) SetClamp(clamp bool) VectorTyped4D[TX, TY, TZ, TW] {
	v.components.x.Clamp = clamp
	v.components.y.Clamp = clamp
	v.components.z.Clamp = clamp
	v.components.w.Clamp = clamp
	return v
}

func (v VectorTyped4D[TX, TY, TZ, TW]) SetBoundaries(minX, maxX TX, minY, maxY TY, minZ, maxZ TZ, minW, maxW TW) VectorTyped4D[TX, TY, TZ, TW] {
	_ = v.components.x.SetBoundaries(minX, maxX)
	_ = v.components.y.SetBoundaries(minY, maxY)
	_ = v.components.z.SetBoundaries(minZ, maxZ)
	_ = v.components.w.SetBoundaries(minW, maxW)
	return v
}

func (v VectorTyped4D[TX, TY, TZ, TW]) MapNames(x, y, z, w string) VectorTyped4D[TX, TY, TZ, TW] {
	v.components.names[0] = x
	v.components.names[1] = y
	v.components.names[2] = z
	v.components.names[3] = w
	return v
}
