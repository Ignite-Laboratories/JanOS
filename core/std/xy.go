package std

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/num"
)

// XY is a general structure for holding generic bounded (x,y) coordinate values.
type XY[T num.ExtendedPrimitive] struct {
	X Bounded[T]
	Y Bounded[T]
}

// Set sets the coordinate values.
func (c XY[T]) Set(x, y T) XY[T] {
	c.X = c.X.Set(x)
	c.Y = c.Y.Set(y)
	return c
}

// SetBoundaries sets the limits for the coordinates.  As coordinates are zero indexed, the boundary
// value should be the number of positions.  For example, in a 1024x768 coordinate space the boundary
// values should also be 1024x768.
func (c XY[T]) SetBoundaries(xBound, yBound T) XY[T] {
	c.X = c.X.SetBoundary(xBound)
	c.Y = c.Y.SetBoundary(yBound)
	return c
}

func (c XY[T]) SetX(x T) XY[T] {
	c.X = c.X.Set(x)
	return c
}

func (c XY[T]) SetY(y T) XY[T] {
	c.Y = c.Y.Set(y)
	return c
}

func (c XY[T]) SetXBoundary(xBound T) XY[T] {
	c.X.SetBoundary(xBound)
	return c
}

func (c XY[T]) SetYBoundary(yBound T) XY[T] {
	c.Y.SetBoundary(yBound)
	return c
}

func (c XY[T]) String() string {
	return fmt.Sprintf("(%v, %v)", c.X, c.Y)
}
