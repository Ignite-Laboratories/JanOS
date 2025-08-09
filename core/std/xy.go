package std

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/num"
)

// XY is a general structure for holding generic bounded (x,y) coordinate values.
type XY[T num.Primitive] struct {
	X Bounded[T]
	Y Bounded[T]
}

// SetClamp sets whether the directions should clamp to their boundaries or overflow and under-flow.
func (coords XY[T]) SetClamp(shouldClamp bool) XY[T] {
	coords.X.Clamp = shouldClamp
	coords.Y.Clamp = shouldClamp
	return coords
}

// Set sets the coordinate values.
func (coords XY[T]) Set(x, y T) XY[T] {
	coords.X = coords.X.Set(x)
	coords.Y = coords.Y.Set(y)
	return coords
}

// SetBoundaries inclusively sets the coordinate boundaries for all directions.
//
// NOTE: This means to represent 1024x768 you should use 1023x767 =)
func (coords XY[T]) SetBoundaries(minX, maxX, minY, maxY T) XY[T] {
	coords.X = coords.X.SetBoundaries(minX, maxX)
	coords.Y = coords.Y.SetBoundaries(minY, maxY)
	return coords
}

// SetAll first sets the boundaries for each direction, then sets their directional values.
func (coords XY[T]) SetAll(x, y, minX, maxX, minY, maxY T) XY[T] {
	return coords.SetBoundaries(minX, maxX, minY, maxY).Set(x, y)
}

// SetFromNormalized sets the bounded directional values using float64 unit vectors from the [0.0, 1.0]
// range, where 0.0 maps to the coordinate space's bounded minimum and 1.0 maps to the bounded maximum.
func (coords XY[T]) SetFromNormalized(x, y float64) XY[T] {
	coords.X = coords.X.SetFromNormalized(x)
	coords.Y = coords.Y.SetFromNormalized(y)
	return coords
}

// SetFromNormalized32 sets the bounded directional values using float32 unit vectors from the [0.0, 1.0]
// range, where 0.0 maps to the coordinate space's bounded minimum and 1.0 maps to the bounded maximum.
func (coords XY[T]) SetFromNormalized32(x, y float32) XY[T] {
	return coords.SetFromNormalized(float64(x), float64(y))
}

// Normalize converts the bounded directional values to float64 unit vectors in the range [0.0, 1.0],
// where the coordinate space's bounded minimum maps to 0.0 and the bounded maximum maps to 1.0.
func (coords XY[T]) Normalize() (x float64, y float64) {
	return coords.X.Normalize(), coords.Y.Normalize()
}

// Normalize32 converts the bounded directional values to float32 unit vectors in the range [0.0, 1.0],
// where the coordinate space's bounded minimum maps to 0.0 and the bounded maximum maps to 1.0.
func (coords XY[T]) Normalize32() (x float32, y float32) {
	return coords.X.Normalize32(), coords.Y.Normalize32()
}

func (coords XY[T]) String() string {
	return fmt.Sprintf("(%v, %v)", coords.X.Value(), coords.Y.Value())
}
