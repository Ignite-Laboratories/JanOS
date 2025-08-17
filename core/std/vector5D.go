package std

import (
	"fmt"
	format "github.com/ignite-laboratories/core/std/name/format"
	"github.com/ignite-laboratories/core/std/num"
	"github.com/ignite-laboratories/core/sys/support"
)

type components5D[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive, TA num.Primitive] struct {
	names [5]string
	x     Bounded[TX]
	y     Bounded[TY]
	z     Bounded[TZ]
	w     Bounded[TW]
	a     Bounded[TA]
}

// VectorTyped5D represents an asymmetric vector of seven dissimilar numeric cursors.
type VectorTyped5D[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive, TA num.Primitive] struct {
	Entity
	components components5D[TX, TY, TZ, TW, TA]
}

func NewVectorTyped5D[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive, TA num.Primitive]() VectorTyped5D[TX, TY, TZ, TW, TA] {
	return VectorTyped5D[TX, TY, TZ, TW, TA]{
		Entity: NewEntity[format.Default](),
	}
}

func (v VectorTyped5D[TX, TY, TZ, TW, TA]) String() string {
	if len(v.components.names[0]) == 0 {
		v.components.names = [5]string{"X", "Y", "Z", "W", "A"}
	}

	if support.AllSameTypes(TX(0), TY(0), TZ(0), TW(0), TA(0)) {
		return fmt.Sprintf("5D[%T].{%v: %v, %v: %v, %v: %v, %v: %v, %v: %v}(%v)", TX(0), v.components.names[0], v.components.x.ValueString(), v.components.names[1], v.components.y.ValueString(), v.components.names[2], v.components.z.ValueString(), v.components.names[3], v.components.w.ValueString(), v.components.names[4], v.components.a.ValueString(), v.GivenName.Name)
	}
	return fmt.Sprintf("5D[%T, %T, %T, %T, %T].{%v: %v, %v: %v, %v: %v, %v: %v, %v: %v}(%v)", TX(0), TY(0), TZ(0), TW(0), TA(0), v.components.names[0], v.components.x.ValueString(), v.components.names[1], v.components.y.ValueString(), v.components.names[2], v.components.z.ValueString(), v.components.names[3], v.components.w.ValueString(), v.components.names[4], v.components.a.ValueString(), v.GivenName.Name)
}

// Vector5D represents a vector of seven like-typed numeric cursors.
type Vector5D[T num.Primitive] = VectorTyped5D[T, T, T, T, T]

func NewVector5D[T num.Primitive]() Vector5D[T] {
	return Vector5D[T]{
		Entity: NewEntity[format.Default](),
	}
}

func (v VectorTyped5D[TX, TY, TZ, TW, TA]) SetValues(x TX, y TY, z TZ, w TW, a TA) VectorTyped5D[TX, TY, TZ, TW, TA] {
	_ = v.components.x.Set(x)
	_ = v.components.y.Set(y)
	_ = v.components.z.Set(z)
	_ = v.components.w.Set(w)
	_ = v.components.a.Set(a)
	return v
}

func (v VectorTyped5D[TX, TY, TZ, TW, TA]) GetX() TX {
	return v.components.x.Value()
}

func (v VectorTyped5D[TX, TY, TZ, TW, TA]) GetY() TY {
	return v.components.y.Value()
}

func (v VectorTyped5D[TX, TY, TZ, TW, TA]) GetZ() TZ {
	return v.components.z.Value()
}

func (v VectorTyped5D[TX, TY, TZ, TW, TA]) GetW() TW {
	return v.components.w.Value()
}

func (v VectorTyped5D[TX, TY, TZ, TW, TA]) GetA() TA {
	return v.components.a.Value()
}

func (v VectorTyped5D[TX, TY, TZ, TW, TA]) SetX(x TX) VectorTyped5D[TX, TY, TZ, TW, TA] {
	_ = v.components.x.Set(x)
	return v
}

func (v VectorTyped5D[TX, TY, TZ, TW, TA]) SetY(y TY) VectorTyped5D[TX, TY, TZ, TW, TA] {
	_ = v.components.y.Set(y)
	return v
}

func (v VectorTyped5D[TX, TY, TZ, TW, TA]) SetZ(z TZ) VectorTyped5D[TX, TY, TZ, TW, TA] {
	_ = v.components.z.Set(z)
	return v
}

func (v VectorTyped5D[TX, TY, TZ, TW, TA]) SetW(w TW) VectorTyped5D[TX, TY, TZ, TW, TA] {
	_ = v.components.w.Set(w)
	return v
}

func (v VectorTyped5D[TX, TY, TZ, TW, TA]) SetA(a TA) VectorTyped5D[TX, TY, TZ, TW, TA] {
	_ = v.components.a.Set(a)
	return v
}

func (v VectorTyped5D[TX, TY, TZ, TW, TA]) SetClamp(clamp bool) VectorTyped5D[TX, TY, TZ, TW, TA] {
	v.components.x.Clamp = clamp
	v.components.y.Clamp = clamp
	v.components.z.Clamp = clamp
	v.components.w.Clamp = clamp
	v.components.a.Clamp = clamp
	return v
}

func (v VectorTyped5D[TX, TY, TZ, TW, TA]) SetBoundaries(minX, maxX TX, minY, maxY TY, minZ, maxZ TZ, minW, maxW TW, minA, maxA TA) VectorTyped5D[TX, TY, TZ, TW, TA] {
	_ = v.components.x.SetBoundaries(minX, maxX)
	_ = v.components.y.SetBoundaries(minY, maxY)
	_ = v.components.z.SetBoundaries(minZ, maxZ)
	_ = v.components.w.SetBoundaries(minW, maxW)
	_ = v.components.a.SetBoundaries(minA, maxA)
	return v
}

func (v VectorTyped5D[TX, TY, TZ, TW, TA]) MapNames(x, y, z, w, a string) VectorTyped5D[TX, TY, TZ, TW, TA] {
	v.components.names[0] = x
	v.components.names[1] = y
	v.components.names[2] = z
	v.components.names[3] = w
	v.components.names[4] = a
	return v
}
