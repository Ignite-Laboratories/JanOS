package std

import "github.com/ignite-laboratories/core/std/num"

type components6D[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive, TA num.Primitive, TB num.Primitive] struct {
	x Cursor[TX]
	y Cursor[TY]
	z Cursor[TZ]
	w Cursor[TW]
	a Cursor[TA]
	b Cursor[TB]
}

// Vector6DTyped represents an asymmetric vector of seven dissimilar numeric cursors.
type Vector6DTyped[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive, TA num.Primitive, TB num.Primitive] struct {
	components components6D[TX, TY, TZ, TW, TA, TB]
}

// Vector6D represents a vector of seven like-typed numeric cursors.
type Vector6D[T num.Primitive] = Vector6DTyped[T, T, T, T, T, T]

func (v Vector6DTyped[TX, TY, TZ, TW, TA, TB]) SetValues(x TX, y TY, z TZ, w TW, a TA, b TB) Vector6DTyped[TX, TY, TZ, TW, TA, TB] {
	_ = v.components.x.Set(x)
	_ = v.components.y.Set(y)
	_ = v.components.z.Set(z)
	_ = v.components.w.Set(w)
	_ = v.components.a.Set(a)
	_ = v.components.b.Set(b)
	return v
}

func (v Vector6DTyped[TX, TY, TZ, TW, TA, TB]) SetClamp(clamp bool) Vector6DTyped[TX, TY, TZ, TW, TA, TB] {
	v.components.x.Clamp = clamp
	v.components.y.Clamp = clamp
	v.components.z.Clamp = clamp
	v.components.w.Clamp = clamp
	v.components.a.Clamp = clamp
	v.components.b.Clamp = clamp
	return v
}

func (v Vector6DTyped[TX, TY, TZ, TW, TA, TB]) SetBoundaries(minX, maxX TX, minY, maxY TY, minZ, maxZ TZ, minW, maxW TW, minA, maxA TA, minB, maxB TB) Vector6DTyped[TX, TY, TZ, TW, TA, TB] {
	_ = v.components.x.SetBoundaries(minX, maxX)
	_ = v.components.y.SetBoundaries(minY, maxY)
	_ = v.components.z.SetBoundaries(minZ, maxZ)
	_ = v.components.w.SetBoundaries(minW, maxW)
	_ = v.components.a.SetBoundaries(minA, maxA)
	_ = v.components.b.SetBoundaries(minB, maxB)
	return v
}
