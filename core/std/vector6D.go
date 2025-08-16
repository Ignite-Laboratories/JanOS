package std

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/name/format"
	"github.com/ignite-laboratories/core/std/num"
	"github.com/ignite-laboratories/core/sys/support"
)

type components6D[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive, TA num.Primitive, TB num.Primitive] struct {
	names [6]string
	x     Cursor[TX]
	y     Cursor[TY]
	z     Cursor[TZ]
	w     Cursor[TW]
	a     Cursor[TA]
	b     Cursor[TB]
}

// VectorTyped6D represents an asymmetric vector of seven dissimilar numeric cursors.
type VectorTyped6D[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive, TA num.Primitive, TB num.Primitive] struct {
	Entity
	components components6D[TX, TY, TZ, TW, TA, TB]
}

func NewVectorTyped6D[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive, TA num.Primitive, TB num.Primitive]() VectorTyped6D[TX, TY, TZ, TW, TA, TB] {
	return VectorTyped6D[TX, TY, TZ, TW, TA, TB]{
		Entity: NewEntity[format.Default](),
	}
}

func (v VectorTyped6D[TX, TY, TZ, TW, TA, TB]) String() string {
	if len(v.components.names[0]) == 0 {
		v.components.names = [6]string{"X", "Y", "Z", "W", "A", "B"}
	}

	if support.AllSameTypes(TX(0), TY(0), TZ(0), TW(0), TA(0), TB(0)) {
		return fmt.Sprintf("6D[%T].{%v: %v, %v: %v, %v: %v, %v: %v, %v: %v, %v: %v}(%v)", TX(0), v.components.names[0], v.components.x.ValueString(), v.components.names[1], v.components.y.ValueString(), v.components.names[2], v.components.z.ValueString(), v.components.names[3], v.components.w.ValueString(), v.components.names[4], v.components.a.ValueString(), v.components.names[5], v.components.b.ValueString(), v.GivenName.Name)
	}
	return fmt.Sprintf("6D[%T, %T, %T, %T, %T, %T].{%v: %v, %v: %v, %v: %v, %v: %v, %v: %v, %v: %v}(%v)", TX(0), TY(0), TZ(0), TW(0), TA(0), TB(0), v.components.names[0], v.components.x.ValueString(), v.components.names[1], v.components.y.ValueString(), v.components.names[2], v.components.z.ValueString(), v.components.names[3], v.components.w.ValueString(), v.components.names[4], v.components.a.ValueString(), v.components.names[5], v.components.b.ValueString(), v.GivenName.Name)
}

// Vector6D represents a vector of seven like-typed numeric cursors.
type Vector6D[T num.Primitive] = VectorTyped6D[T, T, T, T, T, T]

func NewVector6D[T num.Primitive]() Vector6D[T] {
	return Vector6D[T]{
		Entity: NewEntity[format.Default](),
	}
}

func (v VectorTyped6D[TX, TY, TZ, TW, TA, TB]) SetValues(x TX, y TY, z TZ, w TW, a TA, b TB) VectorTyped6D[TX, TY, TZ, TW, TA, TB] {
	_ = v.components.x.Set(x)
	_ = v.components.y.Set(y)
	_ = v.components.z.Set(z)
	_ = v.components.w.Set(w)
	_ = v.components.a.Set(a)
	_ = v.components.b.Set(b)
	return v
}

func (v VectorTyped6D[TX, TY, TZ, TW, TA, TB]) SetX(x TX) VectorTyped6D[TX, TY, TZ, TW, TA, TB] {
	_ = v.components.x.Set(x)
	return v
}

func (v VectorTyped6D[TX, TY, TZ, TW, TA, TB]) SetY(y TY) VectorTyped6D[TX, TY, TZ, TW, TA, TB] {
	_ = v.components.y.Set(y)
	return v
}

func (v VectorTyped6D[TX, TY, TZ, TW, TA, TB]) SetZ(z TZ) VectorTyped6D[TX, TY, TZ, TW, TA, TB] {
	_ = v.components.z.Set(z)
	return v
}

func (v VectorTyped6D[TX, TY, TZ, TW, TA, TB]) SetW(w TW) VectorTyped6D[TX, TY, TZ, TW, TA, TB] {
	_ = v.components.w.Set(w)
	return v
}

func (v VectorTyped6D[TX, TY, TZ, TW, TA, TB]) SetA(a TA) VectorTyped6D[TX, TY, TZ, TW, TA, TB] {
	_ = v.components.a.Set(a)
	return v
}

func (v VectorTyped6D[TX, TY, TZ, TW, TA, TB]) SetB(b TB) VectorTyped6D[TX, TY, TZ, TW, TA, TB] {
	_ = v.components.b.Set(b)
	return v
}

func (v VectorTyped6D[TX, TY, TZ, TW, TA, TB]) SetClamp(clamp bool) VectorTyped6D[TX, TY, TZ, TW, TA, TB] {
	v.components.x.Clamp = clamp
	v.components.y.Clamp = clamp
	v.components.z.Clamp = clamp
	v.components.w.Clamp = clamp
	v.components.a.Clamp = clamp
	v.components.b.Clamp = clamp
	return v
}

func (v VectorTyped6D[TX, TY, TZ, TW, TA, TB]) SetBoundaries(minX, maxX TX, minY, maxY TY, minZ, maxZ TZ, minW, maxW TW, minA, maxA TA, minB, maxB TB) VectorTyped6D[TX, TY, TZ, TW, TA, TB] {
	_ = v.components.x.SetBoundaries(minX, maxX)
	_ = v.components.y.SetBoundaries(minY, maxY)
	_ = v.components.z.SetBoundaries(minZ, maxZ)
	_ = v.components.w.SetBoundaries(minW, maxW)
	_ = v.components.a.SetBoundaries(minA, maxA)
	_ = v.components.b.SetBoundaries(minB, maxB)
	return v
}

func (v VectorTyped6D[TX, TY, TZ, TW, TA, TB]) MapNames(x, y, z, w, a, b string) VectorTyped6D[TX, TY, TZ, TW, TA, TB] {
	v.components.names[0] = x
	v.components.names[1] = y
	v.components.names[2] = z
	v.components.names[3] = w
	v.components.names[4] = a
	v.components.names[5] = b
	return v
}
