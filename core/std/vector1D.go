package std

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/name/format"
	"github.com/ignite-laboratories/core/std/num"
)

type components1D[TX num.Primitive] struct {
	names [1]string
	x     Cursor[TX]
}

// VectorTyped1D represents an asymmetric vector of seven dissimilar numeric cursors.
type VectorTyped1D[TX num.Primitive] struct {
	Entity
	components components1D[TX]
}

func NewVectorTyped1D[TX num.Primitive]() VectorTyped1D[TX] {
	return VectorTyped1D[TX]{
		Entity: NewEntity[format.Default](),
	}
}

// Vector1D represents a vector of seven like-typed numeric cursors.
type Vector1D[T num.Primitive] = VectorTyped1D[T]

func NewVector1D[T num.Primitive]() Vector1D[T] {
	return Vector1D[T]{
		Entity: NewEntity[format.Default](),
	}
}

func (v VectorTyped1D[TX]) String() string {
	if len(v.components.names[0]) == 0 {
		v.components.names = [1]string{"X"}
	}
	return fmt.Sprintf("1D[%T].{%v: %v}(%v)", TX(0), v.components.names[0], v.components.x.ValueString(), v.GivenName.Name)
}

func (v VectorTyped1D[TX]) SetValues(x TX) VectorTyped1D[TX] {
	_ = v.components.x.Set(x)
	return v
}

func (v VectorTyped1D[TX]) GetX() TX {
	return v.components.x.Value()
}

func (v VectorTyped1D[TX]) SetX(x TX) VectorTyped1D[TX] {
	_ = v.components.x.Set(x)
	return v
}

func (v VectorTyped1D[TX]) SetClamp(clamp bool) VectorTyped1D[TX] {
	v.components.x.Clamp = clamp
	return v
}

func (v VectorTyped1D[TX]) SetBoundaries(minX, maxX TX) VectorTyped1D[TX] {
	_ = v.components.x.SetBoundaries(minX, maxX)
	return v
}

func (v VectorTyped1D[TX]) MapNames(x string) VectorTyped1D[TX] {
	v.components.names[0] = x
	return v
}
