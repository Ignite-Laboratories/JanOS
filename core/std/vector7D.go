package std

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/name/format"
	"github.com/ignite-laboratories/core/std/num"
	"github.com/ignite-laboratories/core/sys/support"
)

type components7D[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive, TA num.Primitive, TB num.Primitive, TC num.Primitive] struct {
	names [7]string
	x     Bounded[TX]
	y     Bounded[TY]
	z     Bounded[TZ]
	w     Bounded[TW]
	a     Bounded[TA]
	b     Bounded[TB]
	c     Bounded[TC]
}

// VectorTyped7D represents an asymmetric vector of seven dissimilar numeric cursors.
type VectorTyped7D[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive, TA num.Primitive, TB num.Primitive, TC num.Primitive] struct {
	Entity
	components components7D[TX, TY, TZ, TW, TA, TB, TC]
}

func NewVectorTyped7D[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive, TA num.Primitive, TB num.Primitive, TC num.Primitive]() VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC] {
	return VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC]{
		Entity: NewEntity[format.Default](),
	}
}

func (v VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC]) String() string {
	if len(v.components.names[0]) == 0 {
		v.components.names = [7]string{"X", "Y", "Z", "W", "A", "B", "C"}
	}

	if support.AllSameTypes(TX(0), TY(0), TZ(0), TW(0), TA(0), TB(0), TC(0)) {
		return fmt.Sprintf("7D[%T].{%v: %v, %v: %v, %v: %v, %v: %v, %v: %v, %v: %v, %v: %v}(%v)", TX(0), v.components.names[0], v.components.x.ValueString(), v.components.names[1], v.components.y.ValueString(), v.components.names[2], v.components.z.ValueString(), v.components.names[3], v.components.w.ValueString(), v.components.names[4], v.components.a.ValueString(), v.components.names[5], v.components.b.ValueString(), v.components.names[6], v.components.c.ValueString(), v.GivenName.Name)
	}
	return fmt.Sprintf("7D[%T, %T, %T, %T, %T, %T, %T].{%v: %v, %v: %v, %v: %v, %v: %v, %v: %v, %v: %v, %v: %v}(%v)", TX(0), TY(0), TZ(0), TW(0), TA(0), TB(0), TC(0), v.components.names[0], v.components.x.ValueString(), v.components.names[1], v.components.y.ValueString(), v.components.names[2], v.components.z.ValueString(), v.components.names[3], v.components.w.ValueString(), v.components.names[4], v.components.a.ValueString(), v.components.names[5], v.components.b.ValueString(), v.components.names[6], v.components.c.ValueString(), v.GivenName.Name)
}

// Vector7D represents a vector of seven like-typed numeric cursors.
type Vector7D[T num.Primitive] = VectorTyped7D[T, T, T, T, T, T, T]

func NewVector7D[T num.Primitive]() Vector7D[T] {
	return Vector7D[T]{
		Entity: NewEntity[format.Default](),
	}
}

func (v VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC]) SetValues(x TX, y TY, z TZ, w TW, a TA, b TB, c TC) VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC] {
	_ = v.components.x.Set(x)
	_ = v.components.y.Set(y)
	_ = v.components.z.Set(z)
	_ = v.components.w.Set(w)
	_ = v.components.a.Set(a)
	_ = v.components.b.Set(b)
	_ = v.components.c.Set(c)
	return v
}

func (v VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC]) GetX() TX {
	return v.components.x.Value()
}

func (v VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC]) GetY() TY {
	return v.components.y.Value()
}

func (v VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC]) GetZ() TZ {
	return v.components.z.Value()
}

func (v VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC]) GetW() TW {
	return v.components.w.Value()
}

func (v VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC]) GetA() TA {
	return v.components.a.Value()
}

func (v VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC]) GetB() TB {
	return v.components.b.Value()
}

func (v VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC]) GetC() TC {
	return v.components.c.Value()
}

func (v VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC]) SetX(x TX) VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC] {
	_ = v.components.x.Set(x)
	return v
}

func (v VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC]) SetY(y TY) VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC] {
	_ = v.components.y.Set(y)
	return v
}

func (v VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC]) SetZ(z TZ) VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC] {
	_ = v.components.z.Set(z)
	return v
}

func (v VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC]) SetW(w TW) VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC] {
	_ = v.components.w.Set(w)
	return v
}

func (v VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC]) SetA(a TA) VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC] {
	_ = v.components.a.Set(a)
	return v
}

func (v VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC]) SetB(b TB) VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC] {
	_ = v.components.b.Set(b)
	return v
}

func (v VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC]) SetC(c TC) VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC] {
	_ = v.components.c.Set(c)
	return v
}

func (v VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC]) SetClamp(clamp bool) VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC] {
	v.components.x.Clamp = clamp
	v.components.y.Clamp = clamp
	v.components.z.Clamp = clamp
	v.components.w.Clamp = clamp
	v.components.a.Clamp = clamp
	v.components.b.Clamp = clamp
	v.components.c.Clamp = clamp
	return v
}

func (v VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC]) SetBoundaries(minX, maxX TX, minY, maxY TY, minZ, maxZ TZ, minW, maxW TW, minA, maxA TA, minB, maxB TB, minC, maxC TC) VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC] {
	_ = v.components.x.SetBoundaries(minX, maxX)
	_ = v.components.y.SetBoundaries(minY, maxY)
	_ = v.components.z.SetBoundaries(minZ, maxZ)
	_ = v.components.w.SetBoundaries(minW, maxW)
	_ = v.components.a.SetBoundaries(minA, maxA)
	_ = v.components.b.SetBoundaries(minB, maxB)
	_ = v.components.c.SetBoundaries(minC, maxC)
	return v
}

func (v VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC]) MapNames(x, y, z, w, a, b, c string) VectorTyped7D[TX, TY, TZ, TW, TA, TB, TC] {
	v.components.names[0] = x
	v.components.names[1] = y
	v.components.names[2] = z
	v.components.names[3] = w
	v.components.names[4] = a
	v.components.names[5] = b
	v.components.names[6] = c
	return v
}
