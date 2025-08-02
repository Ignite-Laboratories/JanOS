package std

import (
	"fmt"
	"github.com/ignite-laboratories/core/math"
)

// XYZ is a general structure for holding generic (x,y,z) coordinate values.
type XYZ[T math.Numeric] struct {
	X T
	Y T
	Z T
}

func (c XYZ[T]) String() string {
	return fmt.Sprintf("(%v, %v, %v)", c.X, c.Y, c.Z)
}

/**
Swizzling
*/

func (c XYZ[T]) XX() (T, T) { return c.X, c.X }
func (c XYZ[T]) XY() (T, T) { return c.X, c.Y }
func (c XYZ[T]) XZ() (T, T) { return c.X, c.Z }
func (c XYZ[T]) YX() (T, T) { return c.Y, c.X }
func (c XYZ[T]) YY() (T, T) { return c.Y, c.Y }
func (c XYZ[T]) YZ() (T, T) { return c.Y, c.Z }
func (c XYZ[T]) ZX() (T, T) { return c.Z, c.X }
func (c XYZ[T]) ZY() (T, T) { return c.Z, c.Y }
func (c XYZ[T]) ZZ() (T, T) { return c.Z, c.Z }

func (c XYZ[T]) XXX() (T, T, T) { return c.X, c.X, c.X }
func (c XYZ[T]) XXY() (T, T, T) { return c.X, c.X, c.Y }
func (c XYZ[T]) XXZ() (T, T, T) { return c.X, c.X, c.Z }
func (c XYZ[T]) XYX() (T, T, T) { return c.X, c.Y, c.X }
func (c XYZ[T]) XYY() (T, T, T) { return c.X, c.Y, c.Y }
func (c XYZ[T]) XYZ() (T, T, T) { return c.X, c.Y, c.Z }
func (c XYZ[T]) XZX() (T, T, T) { return c.X, c.Z, c.X }
func (c XYZ[T]) XZY() (T, T, T) { return c.X, c.Z, c.Y }
func (c XYZ[T]) XZZ() (T, T, T) { return c.X, c.Z, c.Z }
func (c XYZ[T]) YXX() (T, T, T) { return c.Y, c.X, c.X }
func (c XYZ[T]) YXY() (T, T, T) { return c.Y, c.X, c.Y }
func (c XYZ[T]) YXZ() (T, T, T) { return c.Y, c.X, c.Z }
func (c XYZ[T]) YYX() (T, T, T) { return c.Y, c.Y, c.X }
func (c XYZ[T]) YYY() (T, T, T) { return c.Y, c.Y, c.Y }
func (c XYZ[T]) YYZ() (T, T, T) { return c.Y, c.Y, c.Z }
func (c XYZ[T]) YZX() (T, T, T) { return c.Y, c.Z, c.X }
func (c XYZ[T]) YZY() (T, T, T) { return c.Y, c.Z, c.Y }
func (c XYZ[T]) YZZ() (T, T, T) { return c.Y, c.Z, c.Z }
func (c XYZ[T]) ZXX() (T, T, T) { return c.Z, c.X, c.X }
func (c XYZ[T]) ZXY() (T, T, T) { return c.Z, c.X, c.Y }
func (c XYZ[T]) ZXZ() (T, T, T) { return c.Z, c.X, c.Z }
func (c XYZ[T]) ZYX() (T, T, T) { return c.Z, c.Y, c.X }
func (c XYZ[T]) ZYY() (T, T, T) { return c.Z, c.Y, c.Y }
func (c XYZ[T]) ZYZ() (T, T, T) { return c.Z, c.Y, c.Z }
func (c XYZ[T]) ZZX() (T, T, T) { return c.Z, c.Z, c.X }
func (c XYZ[T]) ZZY() (T, T, T) { return c.Z, c.Z, c.Y }
func (c XYZ[T]) ZZZ() (T, T, T) { return c.Z, c.Z, c.Z }
