package std

import "github.com/ignite-laboratories/core/std/num"

type components7D[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive, TA num.Primitive, TB num.Primitive, TC num.Primitive] struct {
	x Cursor[TX]
	y Cursor[TY]
	z Cursor[TZ]
	w Cursor[TW]
	a Cursor[TA]
	b Cursor[TB]
	c Cursor[TC]
}

// Vector7DTyped represents an asymmetric vector of seven dissimilar numeric cursors.
type Vector7DTyped[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive, TA num.Primitive, TB num.Primitive, TC num.Primitive] struct {
	components components7D[TX, TY, TZ, TW, TA, TB, TC]
}

// Vector7D represents a vector of seven like-typed numeric cursors.
type Vector7D[T num.Primitive] = Vector7DTyped[T, T, T, T, T, T, T]

func (v Vector7DTyped[TX, TY, TZ, TW, TA, TB, TC]) SetValues(x TX, y TY, z TZ, w TW, a TA, b TB, c TC) Vector7DTyped[TX, TY, TZ, TW, TA, TB, TC] {
	_ = v.components.x.Set(x)
	_ = v.components.y.Set(y)
	_ = v.components.z.Set(z)
	_ = v.components.w.Set(w)
	_ = v.components.a.Set(a)
	_ = v.components.b.Set(b)
	_ = v.components.c.Set(c)
	return v
}

func (v Vector7DTyped[TX, TY, TZ, TW, TA, TB, TC]) SetClamp(clamp bool) Vector7DTyped[TX, TY, TZ, TW, TA, TB, TC] {
	v.components.x.Clamp = clamp
	v.components.y.Clamp = clamp
	v.components.z.Clamp = clamp
	v.components.w.Clamp = clamp
	v.components.a.Clamp = clamp
	v.components.b.Clamp = clamp
	v.components.c.Clamp = clamp
	return v
}

func (v Vector7DTyped[TX, TY, TZ, TW, TA, TB, TC]) SetBoundaries(minX, maxX TX, minY, maxY TY, minZ, maxZ TZ, minW, maxW TW, minA, maxA TA, minB, maxB TB, minC, maxC TC) Vector7DTyped[TX, TY, TZ, TW, TA, TB, TC] {
	_ = v.components.x.SetBoundaries(minX, maxX)
	_ = v.components.y.SetBoundaries(minY, maxY)
	_ = v.components.z.SetBoundaries(minZ, maxZ)
	_ = v.components.w.SetBoundaries(minW, maxW)
	_ = v.components.a.SetBoundaries(minA, maxA)
	_ = v.components.b.SetBoundaries(minB, maxB)
	_ = v.components.c.SetBoundaries(minC, maxC)
	return v
}
