package std

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/name/format"
	"github.com/ignite-laboratories/core/std/num"
	"github.com/ignite-laboratories/core/sys/support"
)

type components2D[TX num.Primitive, TY num.Primitive] struct {
	names [2]string
	x     Cursor[TX]
	y     Cursor[TY]
}

// VectorTyped2D represents an asymmetric vector of seven dissimilar numeric cursors.
type VectorTyped2D[TX num.Primitive, TY num.Primitive] struct {
	Entity
	components components2D[TX, TY]
}

func NewVectorTyped2D[TX num.Primitive, TY num.Primitive]() VectorTyped2D[TX, TY] {
	return VectorTyped2D[TX, TY]{
		Entity: NewEntity[format.Default](),
	}
}

// Vector2D represents a vector of seven like-typed numeric cursors.
type Vector2D[T num.Primitive] = VectorTyped2D[T, T]

func NewVector2D[T num.Primitive]() Vector2D[T] {
	return Vector2D[T]{
		Entity: NewEntity[format.Default](),
	}
}

func (v VectorTyped2D[TX, TY]) String() string {
	if len(v.components.names[0]) == 0 {
		v.components.names = [2]string{"X", "Y"}
	}

	if support.AllSameTypes(TX(0), TY(0)) {
		return fmt.Sprintf("2D[%T].{%v: %v, %v: %v}(%v)", TX(0), v.components.names[0], v.components.x.ValueString(), v.components.names[1], v.components.y.ValueString(), v.GivenName.Name)
	}
	return fmt.Sprintf("2D[%T, %T].{%v: %v, %v: %v}(%v)", TX(0), TY(0), v.components.names[0], v.components.x.ValueString(), v.components.names[1], v.components.y.ValueString(), v.GivenName.Name)
}

func (v VectorTyped2D[TX, TY]) SetValues(x TX, y TY) VectorTyped2D[TX, TY] {
	_ = v.components.x.Set(x)
	_ = v.components.y.Set(y)
	return v
}

func (v VectorTyped2D[TX, TY]) SetX(x TX) VectorTyped2D[TX, TY] {
	_ = v.components.x.Set(x)
	return v
}

func (v VectorTyped2D[TX, TY]) SetY(y TY) VectorTyped2D[TX, TY] {
	_ = v.components.y.Set(y)
	return v
}

func (v VectorTyped2D[TX, TY]) SetClamp(clamp bool) VectorTyped2D[TX, TY] {
	v.components.x.Clamp = clamp
	v.components.y.Clamp = clamp
	return v
}

func (v VectorTyped2D[TX, TY]) SetBoundaries(minX, maxX TX, minY, maxY TY) VectorTyped2D[TX, TY] {
	_ = v.components.x.SetBoundaries(minX, maxX)
	_ = v.components.y.SetBoundaries(minY, maxY)
	return v
}

func (v VectorTyped2D[TX, TY]) MapNames(x, y string) VectorTyped2D[TX, TY] {
	v.components.names[0] = x
	v.components.names[1] = y
	return v
}
