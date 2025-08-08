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

// Set sets the coordinate values.
func (coords XYZ[T]) Set(x, y, z T) XYZ[T] {
	coords.X = coords.X.Set(x)
	coords.Y = coords.Y.Set(y)
	coords.Z = coords.Z.Set(z)
	return coords
}

// SetBoundaries inclusively sets the coordinate boundaries for all directions.
//
// NOTE: This means to represent 1024x768 you should use 1023x767 =)
func (coords XYZ[T]) SetBoundaries(minX, maxX, minY, maxY, minZ, maxZ T) XYZ[T] {
	coords.X = coords.X.SetBoundaries(minX, maxX)
	coords.Y = coords.Y.SetBoundaries(minY, maxY)
	coords.Z = coords.Z.SetBoundaries(minZ, maxZ)
	return coords
}

// SetAll first sets the boundaries for each direction, then sets their directional values.
func (coords XYZ[T]) SetAll(x, y, z, minX, maxX, minY, maxY, minZ, maxZ T) XYZ[T] {
	return coords.SetBoundaries(minX, maxX, minY, maxY, minZ, maxZ).Set(x, y, z)
}

// SetFromNormalized sets the bounded directional values using float64 unit vectors from the [0.0, 1.0]
// range, where 0.0 maps to the coordinate space's bounded minimum and 1.0 maps to the bounded maximum.
func (coords XYZ[T]) SetFromNormalized(x, y, z float64) XYZ[T] {
	coords.X = coords.X.SetFromNormalized(x)
	coords.Y = coords.Y.SetFromNormalized(y)
	coords.Z = coords.Z.SetFromNormalized(z)
	return coords
}

// SetFromNormalized32 sets the bounded directional values using float32 unit vectors from the [0.0, 1.0]
// range, where 0.0 maps to the coordinate space's bounded minimum and 1.0 maps to the bounded maximum.
func (coords XYZ[T]) SetFromNormalized32(x, y, z float32) XYZ[T] {
	return coords.SetFromNormalized(float64(x), float64(y), float64(z))
}

// Normalize converts the bounded directional values to float64 unit vectors in the range [0.0, 1.0],
// where the coordinate space's bounded minimum maps to 0.0 and the bounded maximum maps to 1.0.
func (coords XYZ[T]) Normalize() (x float64, y float64, z float64) {
	return coords.X.Normalize(), coords.Y.Normalize(), coords.Z.Normalize()
}

// Normalize32 converts the bounded directional values to float32 unit vectors in the range [0.0, 1.0],
// where the coordinate space's bounded minimum maps to 0.0 and the bounded maximum maps to 1.0.
func (coords XYZ[T]) Normalize32() (x float32, y float32, z float32) {
	return coords.X.Normalize32(), coords.Y.Normalize32(), coords.Z.Normalize32()
}

func (coords XYZ[T]) String() string {
	return fmt.Sprintf("(%v, %v, %v)", coords.X, coords.Y, coords.Z)
}

/**
Swizzling

NOTE: This is a regular expression to find and replace swizzle functions into a one-liner if the auto formatter ever kicks in

Find -
func \((.*?)\) ([A-Z]{2,4})\(\) \((.*?)\)[ ]*\{[\n\t ]*return(.*?)[\n\t ]*\}

Replace -
func ($1) $2() ($3) { return$4 }
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
