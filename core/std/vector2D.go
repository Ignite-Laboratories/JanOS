package std

import "github.com/ignite-laboratories/core/std/num"

type components2D[TX num.Primitive, TY num.Primitive] struct {
	x Cursor[TX]
	y Cursor[TY]
}

// Vector2DTyped represents an asymmetric vector of seven dissimilar numeric cursors.
type Vector2DTyped[TX num.Primitive, TY num.Primitive] struct {
	components components2D[TX, TY]
}

// Vector2D represents a vector of seven like-typed numeric cursors.
type Vector2D[T num.Primitive] = Vector2DTyped[T, T]

func (v Vector2DTyped[TX, TY]) SetValues(x TX, y TY) Vector2DTyped[TX, TY] {
	_ = v.components.x.Set(x)
	_ = v.components.y.Set(y)
	return v
}

func (v Vector2DTyped[TX, TY]) SetClamp(clamp bool) Vector2DTyped[TX, TY] {
	v.components.x.Clamp = clamp
	v.components.y.Clamp = clamp
	return v
}

func (v Vector2DTyped[TX, TY]) SetBoundaries(minX, maxX TX, minY, maxY TY) Vector2DTyped[TX, TY] {
	_ = v.components.x.SetBoundaries(minX, maxX)
	_ = v.components.y.SetBoundaries(minY, maxY)
	return v
}
