package std

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/num"
)

// XYZ is a general structure for holding generic (x,y,z) coordinate values.
//
// NOTE: This type also provides rudimentary "swizzling."
type XYZ[T num.ExtendedPrimitive] struct {
	X Bounded[T]
	Y Bounded[T]
	Z Bounded[T]
}

func (c XYZ[T]) Set(x, y, z T) XYZ[T] {
	c.X = c.X.Set(x)
	c.Y = c.Y.Set(y)
	c.Z = c.Z.Set(z)
	return c
}

func (c XYZ[T]) SetBoundaries(xBound, yBound, zBound T) XYZ[T] {
	c.X = c.X.SetBoundary(xBound)
	c.Y = c.Y.SetBoundary(yBound)
	c.Z = c.Z.SetBoundary(zBound)
	return c
}

func (c XYZ[T]) SetX(x T) XYZ[T] {
	c.X = c.X.Set(x)
	return c
}

func (c XYZ[T]) SetY(y T) XYZ[T] {
	c.Y = c.Y.Set(y)
	return c
}

func (c XYZ[T]) SetZ(z T) XYZ[T] {
	c.Z = c.Z.Set(z)
	return c
}

func (c XYZ[T]) SetXBoundary(xBound T) XYZ[T] {
	c.X.SetBoundary(xBound)
	return c
}

func (c XYZ[T]) SetYBoundary(yBound T) XYZ[T] {
	c.Y.SetBoundary(yBound)
	return c
}

func (c XYZ[T]) SetZBoundary(zBound T) XYZ[T] {
	c.Z.SetBoundary(zBound)
	return c
}

func (c XYZ[T]) String() string {
	return fmt.Sprintf("(%v, %v, %v)", c.X.Value(), c.Y.Value(), c.Z.Value())
}

/**
Swizzling
*/

func (c XYZ[T]) XX() (T, T) { return c.X.Value(), c.X.Value() }
func (c XYZ[T]) XY() (T, T) { return c.X.Value(), c.Y.Value() }
func (c XYZ[T]) XZ() (T, T) { return c.X.Value(), c.Z.Value() }
func (c XYZ[T]) YX() (T, T) { return c.Y.Value(), c.X.Value() }
func (c XYZ[T]) YY() (T, T) { return c.Y.Value(), c.Y.Value() }
func (c XYZ[T]) YZ() (T, T) { return c.Y.Value(), c.Z.Value() }
func (c XYZ[T]) ZX() (T, T) { return c.Z.Value(), c.X.Value() }
func (c XYZ[T]) ZY() (T, T) { return c.Z.Value(), c.Y.Value() }
func (c XYZ[T]) ZZ() (T, T) { return c.Z.Value(), c.Z.Value() }

func (c XYZ[T]) XXX() (T, T, T) { return c.X.Value(), c.X.Value(), c.X.Value() }
func (c XYZ[T]) XXY() (T, T, T) { return c.X.Value(), c.X.Value(), c.Y.Value() }
func (c XYZ[T]) XXZ() (T, T, T) { return c.X.Value(), c.X.Value(), c.Z.Value() }
func (c XYZ[T]) XYX() (T, T, T) { return c.X.Value(), c.Y.Value(), c.X.Value() }
func (c XYZ[T]) XYY() (T, T, T) { return c.X.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZ[T]) XYZ() (T, T, T) { return c.X.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZ[T]) XZX() (T, T, T) { return c.X.Value(), c.Z.Value(), c.X.Value() }
func (c XYZ[T]) XZY() (T, T, T) { return c.X.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZ[T]) XZZ() (T, T, T) { return c.X.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZ[T]) YXX() (T, T, T) { return c.Y.Value(), c.X.Value(), c.X.Value() }
func (c XYZ[T]) YXY() (T, T, T) { return c.Y.Value(), c.X.Value(), c.Y.Value() }
func (c XYZ[T]) YXZ() (T, T, T) { return c.Y.Value(), c.X.Value(), c.Z.Value() }
func (c XYZ[T]) YYX() (T, T, T) { return c.Y.Value(), c.Y.Value(), c.X.Value() }
func (c XYZ[T]) YYY() (T, T, T) { return c.Y.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZ[T]) YYZ() (T, T, T) { return c.Y.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZ[T]) YZX() (T, T, T) { return c.Y.Value(), c.Z.Value(), c.X.Value() }
func (c XYZ[T]) YZY() (T, T, T) { return c.Y.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZ[T]) YZZ() (T, T, T) { return c.Y.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZ[T]) ZXX() (T, T, T) { return c.Z.Value(), c.X.Value(), c.X.Value() }
func (c XYZ[T]) ZXY() (T, T, T) { return c.Z.Value(), c.X.Value(), c.Y.Value() }
func (c XYZ[T]) ZXZ() (T, T, T) { return c.Z.Value(), c.X.Value(), c.Z.Value() }
func (c XYZ[T]) ZYX() (T, T, T) { return c.Z.Value(), c.Y.Value(), c.X.Value() }
func (c XYZ[T]) ZYY() (T, T, T) { return c.Z.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZ[T]) ZYZ() (T, T, T) { return c.Z.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZ[T]) ZZX() (T, T, T) { return c.Z.Value(), c.Z.Value(), c.X.Value() }
func (c XYZ[T]) ZZY() (T, T, T) { return c.Z.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZ[T]) ZZZ() (T, T, T) { return c.Z.Value(), c.Z.Value(), c.Z.Value() }
