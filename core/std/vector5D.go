package std

import "github.com/ignite-laboratories/core/std/num"

type components5D[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive, TA num.Primitive] struct {
	x Cursor[TX]
	y Cursor[TY]
	z Cursor[TZ]
	w Cursor[TW]
	a Cursor[TA]
}

// Vector5DTyped represents an asymmetric vector of seven dissimilar numeric cursors.
type Vector5DTyped[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive, TA num.Primitive] struct {
	components components5D[TX, TY, TZ, TW, TA]
}

// Vector5D represents a vector of seven like-typed numeric cursors.
type Vector5D[T num.Primitive] = Vector5DTyped[T, T, T, T, T]

func (v Vector5DTyped[TX, TY, TZ, TW, TA]) SetValues(x TX, y TY, z TZ, w TW, a TA) Vector5DTyped[TX, TY, TZ, TW, TA] {
	_ = v.components.x.Set(x)
	_ = v.components.y.Set(y)
	_ = v.components.z.Set(z)
	_ = v.components.w.Set(w)
	_ = v.components.a.Set(a)
	return v
}

func (v Vector5DTyped[TX, TY, TZ, TW, TA]) SetClamp(clamp bool) Vector5DTyped[TX, TY, TZ, TW, TA] {
	v.components.x.Clamp = clamp
	v.components.y.Clamp = clamp
	v.components.z.Clamp = clamp
	v.components.w.Clamp = clamp
	v.components.a.Clamp = clamp
	return v
}

func (v Vector5DTyped[TX, TY, TZ, TW, TA]) SetBoundaries(minX, maxX TX, minY, maxY TY, minZ, maxZ TZ, minW, maxW TW, minA, maxA TA) Vector5DTyped[TX, TY, TZ, TW, TA] {
	_ = v.components.x.SetBoundaries(minX, maxX)
	_ = v.components.y.SetBoundaries(minY, maxY)
	_ = v.components.z.SetBoundaries(minZ, maxZ)
	_ = v.components.w.SetBoundaries(minW, maxW)
	_ = v.components.a.SetBoundaries(minA, maxA)
	return v
}
