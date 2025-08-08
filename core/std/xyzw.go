package std

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/num"
)

// XYZW is a general structure for holding homogeneous (x,y,z,w) coordinate values.
//
// In homogeneous coordinates, W acts as a scaling factor where:
//
//	W = 0 represents points at infinity (vectors)
//	W = 1 represents normal 3D points
//	Other W values represent scaled 3D points (x/w, y/w, z/w)
//
// NOTE: This type also provides rudimentary "swizzling."
type XYZW[T num.ExtendedPrimitive] struct {
	X num.Bounded[T]
	Y num.Bounded[T]
	Z num.Bounded[T]
	W float64
}

// Set sets the coordinate values.
//
// NOTE: If no w is provided, the value remains unchanged.
func (coords XYZW[T]) Set(x, y, z T, w ...float64) XYZW[T] {
	coords.X = coords.X.Set(x)
	coords.Y = coords.Y.Set(y)
	coords.Z = coords.Z.Set(z)
	if len(w) > 0 {
		coords.W = w[0]
	}
	return coords
}

// SetBoundaries inclusively sets the coordinate boundaries for all directions.
//
// NOTE: This means to represent 1024x768 you should use 1023x767 =)
func (coords XYZW[T]) SetBoundaries(minX, maxX, minY, maxY, minZ, maxZ T) XYZW[T] {
	coords.X = coords.X.SetBoundaries(minX, maxX)
	coords.Y = coords.Y.SetBoundaries(minY, maxY)
	coords.Z = coords.Z.SetBoundaries(minZ, maxZ)
	return coords
}

// SetAll first sets the boundaries for each direction, then sets their directional values.
func (coords XYZW[T]) SetAll(x, y, z T, w float64, minX, maxX, minY, maxY, minZ, maxZ T) XYZW[T] {
	return coords.SetBoundaries(minX, maxX, minY, maxY, minZ, maxZ).Set(x, y, z, w)
}

// SetFromNormalized sets the bounded directional values using float64 unit vectors from the [0.0, 1.0]
// range, where 0.0 maps to the coordinate space's bounded minimum and 1.0 maps to the bounded maximum.
//
// NOTE: If no w is provided, the value remains unchanged.
func (coords XYZW[T]) SetFromNormalized(x, y, z float64, w ...float64) XYZW[T] {
	coords.X = coords.X.SetFromNormalized(x)
	coords.Y = coords.Y.SetFromNormalized(y)
	coords.Z = coords.Z.SetFromNormalized(z)
	if len(w) > 0 {
		coords.W = w[0]
	}
	return coords
}

// SetFromNormalized32 sets the bounded directional values using float32 unit vectors from the [0.0, 1.0]
// range, where 0.0 maps to the coordinate space's bounded minimum and 1.0 maps to the bounded maximum.
//
// NOTE: If no w is provided, the value remains unchanged.
func (coords XYZW[T]) SetFromNormalized32(x, y, z float32, w ...float32) XYZW[T] {
	w64 := make([]float64, len(w))
	for i, v := range w {
		w64[i] = float64(v)
	}

	return coords.SetFromNormalized(float64(x), float64(y), float64(z), w64...)
}

// Normalize converts the bounded directional values to float64 unit vectors in the range [0.0, 1.0],
// where the coordinate space's bounded minimum maps to 0.0 and the bounded maximum maps to 1.0.
func (coords XYZW[T]) Normalize() (float64, float64, float64, float64) {
	return coords.X.Normalize(), coords.Y.Normalize(), coords.Z.Normalize(), coords.W
}

// Normalize32 converts the bounded directional values to float32 unit vectors in the range [0.0, 1.0],
// where the coordinate space's bounded minimum maps to 0.0 and the bounded maximum maps to 1.0.
func (coords XYZW[T]) Normalize32() (float32, float32, float32, float32) {
	return coords.X.Normalize32(), coords.Y.Normalize32(), coords.Z.Normalize32(), float32(coords.W)
}

func (coords XYZW[T]) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v)", coords.X, coords.Y, coords.Z, coords.W)
}

/**
Swizzling
*/

func (c XYZW[T]) XX() (T, T)             { return c.X.Value(), c.X.Value() }
func (c XYZW[T]) XY() (T, T)             { return c.X.Value(), c.Y.Value() }
func (c XYZW[T]) XZ() (T, T)             { return c.X.Value(), c.Z.Value() }
func (c XYZW[T]) XW() (T, float64)       { return c.X.Value(), c.W }
func (c XYZW[T]) YX() (T, T)             { return c.Y.Value(), c.X.Value() }
func (c XYZW[T]) YY() (T, T)             { return c.Y.Value(), c.Y.Value() }
func (c XYZW[T]) YZ() (T, T)             { return c.Y.Value(), c.Z.Value() }
func (c XYZW[T]) YW() (T, float64)       { return c.Y.Value(), c.W }
func (c XYZW[T]) ZX() (T, T)             { return c.Z.Value(), c.X.Value() }
func (c XYZW[T]) ZY() (T, T)             { return c.Z.Value(), c.Y.Value() }
func (c XYZW[T]) ZZ() (T, T)             { return c.Z.Value(), c.Z.Value() }
func (c XYZW[T]) ZW() (T, float64)       { return c.Z.Value(), c.W }
func (c XYZW[T]) WX() (float64, T)       { return c.W, c.X.Value() }
func (c XYZW[T]) WY() (float64, T)       { return c.W, c.Y.Value() }
func (c XYZW[T]) WZ() (float64, T)       { return c.W, c.Z.Value() }
func (c XYZW[T]) WW() (float64, float64) { return c.W, c.W }

func (c XYZW[T]) XXX() (T, T, T)                   { return c.X.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T]) XXY() (T, T, T)                   { return c.X.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T]) XXZ() (T, T, T)                   { return c.X.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T]) XXW() (T, T, float64)             { return c.X.Value(), c.X.Value(), c.W }
func (c XYZW[T]) XYX() (T, T, T)                   { return c.X.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T]) XYY() (T, T, T)                   { return c.X.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T]) XYZ() (T, T, T)                   { return c.X.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T]) XYW() (T, T, float64)             { return c.X.Value(), c.Y.Value(), c.W }
func (c XYZW[T]) XZX() (T, T, T)                   { return c.X.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T]) XZY() (T, T, T)                   { return c.X.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T]) XZZ() (T, T, T)                   { return c.X.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T]) XZW() (T, T, float64)             { return c.X.Value(), c.Z.Value(), c.W }
func (c XYZW[T]) XWX() (T, float64, T)             { return c.X.Value(), c.W, c.X.Value() }
func (c XYZW[T]) XWY() (T, float64, T)             { return c.X.Value(), c.W, c.Y.Value() }
func (c XYZW[T]) XWZ() (T, float64, T)             { return c.X.Value(), c.W, c.Z.Value() }
func (c XYZW[T]) XWW() (T, float64, float64)       { return c.X.Value(), c.W, c.W }
func (c XYZW[T]) YXX() (T, T, T)                   { return c.Y.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T]) YXY() (T, T, T)                   { return c.Y.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T]) YXZ() (T, T, T)                   { return c.Y.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T]) YXW() (T, T, float64)             { return c.Y.Value(), c.X.Value(), c.W }
func (c XYZW[T]) YYX() (T, T, T)                   { return c.Y.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T]) YYY() (T, T, T)                   { return c.Y.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T]) YYZ() (T, T, T)                   { return c.Y.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T]) YYW() (T, T, float64)             { return c.Y.Value(), c.Y.Value(), c.W }
func (c XYZW[T]) YZX() (T, T, T)                   { return c.Y.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T]) YZY() (T, T, T)                   { return c.Y.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T]) YZZ() (T, T, T)                   { return c.Y.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T]) YZW() (T, T, float64)             { return c.Y.Value(), c.Z.Value(), c.W }
func (c XYZW[T]) YWX() (T, float64, T)             { return c.Y.Value(), c.W, c.X.Value() }
func (c XYZW[T]) YWY() (T, float64, T)             { return c.Y.Value(), c.W, c.Y.Value() }
func (c XYZW[T]) YWZ() (T, float64, T)             { return c.Y.Value(), c.W, c.Z.Value() }
func (c XYZW[T]) YWW() (T, float64, float64)       { return c.Y.Value(), c.W, c.W }
func (c XYZW[T]) ZXX() (T, T, T)                   { return c.Z.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T]) ZXY() (T, T, T)                   { return c.Z.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T]) ZXZ() (T, T, T)                   { return c.Z.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T]) ZXW() (T, T, float64)             { return c.Z.Value(), c.X.Value(), c.W }
func (c XYZW[T]) ZYX() (T, T, T)                   { return c.Z.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T]) ZYY() (T, T, T)                   { return c.Z.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T]) ZYZ() (T, T, T)                   { return c.Z.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T]) ZYW() (T, T, float64)             { return c.Z.Value(), c.Y.Value(), c.W }
func (c XYZW[T]) ZZX() (T, T, T)                   { return c.Z.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T]) ZZY() (T, T, T)                   { return c.Z.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T]) ZZZ() (T, T, T)                   { return c.Z.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T]) ZZW() (T, T, float64)             { return c.Z.Value(), c.Z.Value(), c.W }
func (c XYZW[T]) ZWX() (T, float64, T)             { return c.Z.Value(), c.W, c.X.Value() }
func (c XYZW[T]) ZWY() (T, float64, T)             { return c.Z.Value(), c.W, c.Y.Value() }
func (c XYZW[T]) ZWZ() (T, float64, T)             { return c.Z.Value(), c.W, c.Z.Value() }
func (c XYZW[T]) ZWW() (T, float64, float64)       { return c.Z.Value(), c.W, c.W }
func (c XYZW[T]) WXX() (float64, T, T)             { return c.W, c.X.Value(), c.X.Value() }
func (c XYZW[T]) WXY() (float64, T, T)             { return c.W, c.X.Value(), c.Y.Value() }
func (c XYZW[T]) WXZ() (float64, T, T)             { return c.W, c.X.Value(), c.Z.Value() }
func (c XYZW[T]) WXW() (float64, T, float64)       { return c.W, c.X.Value(), c.W }
func (c XYZW[T]) WYX() (float64, T, T)             { return c.W, c.Y.Value(), c.X.Value() }
func (c XYZW[T]) WYY() (float64, T, T)             { return c.W, c.Y.Value(), c.Y.Value() }
func (c XYZW[T]) WYZ() (float64, T, T)             { return c.W, c.Y.Value(), c.Z.Value() }
func (c XYZW[T]) WYW() (float64, T, float64)       { return c.W, c.Y.Value(), c.W }
func (c XYZW[T]) WZX() (float64, T, T)             { return c.W, c.Z.Value(), c.X.Value() }
func (c XYZW[T]) WZY() (float64, T, T)             { return c.W, c.Z.Value(), c.Y.Value() }
func (c XYZW[T]) WZZ() (float64, T, T)             { return c.W, c.Z.Value(), c.Z.Value() }
func (c XYZW[T]) WZW() (float64, T, float64)       { return c.W, c.Z.Value(), c.W }
func (c XYZW[T]) WWX() (float64, float64, T)       { return c.W, c.W, c.X.Value() }
func (c XYZW[T]) WWY() (float64, float64, T)       { return c.W, c.W, c.Y.Value() }
func (c XYZW[T]) WWZ() (float64, float64, T)       { return c.W, c.W, c.Z.Value() }
func (c XYZW[T]) WWW() (float64, float64, float64) { return c.W, c.W, c.W }

func (c XYZW[T]) XXXX() (T, T, T, T)                         { return c.X.Value(), c.X.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T]) XXXY() (T, T, T, T)                         { return c.X.Value(), c.X.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T]) XXXZ() (T, T, T, T)                         { return c.X.Value(), c.X.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T]) XXXW() (T, T, T, float64)                   { return c.X.Value(), c.X.Value(), c.X.Value(), c.W }
func (c XYZW[T]) XXYX() (T, T, T, T)                         { return c.X.Value(), c.X.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T]) XXYY() (T, T, T, T)                         { return c.X.Value(), c.X.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T]) XXYZ() (T, T, T, T)                         { return c.X.Value(), c.X.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T]) XXYW() (T, T, T, float64)                   { return c.X.Value(), c.X.Value(), c.Y.Value(), c.W }
func (c XYZW[T]) XXZX() (T, T, T, T)                         { return c.X.Value(), c.X.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T]) XXZY() (T, T, T, T)                         { return c.X.Value(), c.X.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T]) XXZZ() (T, T, T, T)                         { return c.X.Value(), c.X.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T]) XXZW() (T, T, T, float64)                   { return c.X.Value(), c.X.Value(), c.Z.Value(), c.W }
func (c XYZW[T]) XXWX() (T, T, float64, T)                   { return c.X.Value(), c.X.Value(), c.W, c.X.Value() }
func (c XYZW[T]) XXWY() (T, T, float64, T)                   { return c.X.Value(), c.X.Value(), c.W, c.Y.Value() }
func (c XYZW[T]) XXWZ() (T, T, float64, T)                   { return c.X.Value(), c.X.Value(), c.W, c.Z.Value() }
func (c XYZW[T]) XXWW() (T, T, float64, float64)             { return c.X.Value(), c.X.Value(), c.W, c.W }
func (c XYZW[T]) XYXX() (T, T, T, T)                         { return c.X.Value(), c.Y.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T]) XYXY() (T, T, T, T)                         { return c.X.Value(), c.Y.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T]) XYXZ() (T, T, T, T)                         { return c.X.Value(), c.Y.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T]) XYXW() (T, T, T, float64)                   { return c.X.Value(), c.Y.Value(), c.X.Value(), c.W }
func (c XYZW[T]) XYYX() (T, T, T, T)                         { return c.X.Value(), c.Y.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T]) XYYY() (T, T, T, T)                         { return c.X.Value(), c.Y.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T]) XYYZ() (T, T, T, T)                         { return c.X.Value(), c.Y.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T]) XYYW() (T, T, T, float64)                   { return c.X.Value(), c.Y.Value(), c.Y.Value(), c.W }
func (c XYZW[T]) XYZX() (T, T, T, T)                         { return c.X.Value(), c.Y.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T]) XYZY() (T, T, T, T)                         { return c.X.Value(), c.Y.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T]) XYZZ() (T, T, T, T)                         { return c.X.Value(), c.Y.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T]) XYZW() (T, T, T, float64)                   { return c.X.Value(), c.Y.Value(), c.Z.Value(), c.W }
func (c XYZW[T]) XYWX() (T, T, float64, T)                   { return c.X.Value(), c.Y.Value(), c.W, c.X.Value() }
func (c XYZW[T]) XYWY() (T, T, float64, T)                   { return c.X.Value(), c.Y.Value(), c.W, c.Y.Value() }
func (c XYZW[T]) XYWZ() (T, T, float64, T)                   { return c.X.Value(), c.Y.Value(), c.W, c.Z.Value() }
func (c XYZW[T]) XYWW() (T, T, float64, float64)             { return c.X.Value(), c.Y.Value(), c.W, c.W }
func (c XYZW[T]) XZXX() (T, T, T, T)                         { return c.X.Value(), c.Z.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T]) XZXY() (T, T, T, T)                         { return c.X.Value(), c.Z.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T]) XZXZ() (T, T, T, T)                         { return c.X.Value(), c.Z.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T]) XZXW() (T, T, T, float64)                   { return c.X.Value(), c.Z.Value(), c.X.Value(), c.W }
func (c XYZW[T]) XZYX() (T, T, T, T)                         { return c.X.Value(), c.Z.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T]) XZYY() (T, T, T, T)                         { return c.X.Value(), c.Z.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T]) XZYZ() (T, T, T, T)                         { return c.X.Value(), c.Z.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T]) XZYW() (T, T, T, float64)                   { return c.X.Value(), c.Z.Value(), c.Y.Value(), c.W }
func (c XYZW[T]) XZZX() (T, T, T, T)                         { return c.X.Value(), c.Z.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T]) XZZY() (T, T, T, T)                         { return c.X.Value(), c.Z.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T]) XZZZ() (T, T, T, T)                         { return c.X.Value(), c.Z.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T]) XZZW() (T, T, T, float64)                   { return c.X.Value(), c.Z.Value(), c.Z.Value(), c.W }
func (c XYZW[T]) XZWX() (T, T, float64, T)                   { return c.X.Value(), c.Z.Value(), c.W, c.X.Value() }
func (c XYZW[T]) XZWY() (T, T, float64, T)                   { return c.X.Value(), c.Z.Value(), c.W, c.Y.Value() }
func (c XYZW[T]) XZWZ() (T, T, float64, T)                   { return c.X.Value(), c.Z.Value(), c.W, c.Z.Value() }
func (c XYZW[T]) XZWW() (T, T, float64, float64)             { return c.X.Value(), c.Z.Value(), c.W, c.W }
func (c XYZW[T]) XWXX() (T, float64, T, T)                   { return c.X.Value(), c.W, c.X.Value(), c.X.Value() }
func (c XYZW[T]) XWXY() (T, float64, T, T)                   { return c.X.Value(), c.W, c.X.Value(), c.Y.Value() }
func (c XYZW[T]) XWXZ() (T, float64, T, T)                   { return c.X.Value(), c.W, c.X.Value(), c.Z.Value() }
func (c XYZW[T]) XWXW() (T, float64, T, float64)             { return c.X.Value(), c.W, c.X.Value(), c.W }
func (c XYZW[T]) XWYX() (T, float64, T, T)                   { return c.X.Value(), c.W, c.Y.Value(), c.X.Value() }
func (c XYZW[T]) XWYY() (T, float64, T, T)                   { return c.X.Value(), c.W, c.Y.Value(), c.Y.Value() }
func (c XYZW[T]) XWYZ() (T, float64, T, T)                   { return c.X.Value(), c.W, c.Y.Value(), c.Z.Value() }
func (c XYZW[T]) XWYW() (T, float64, T, float64)             { return c.X.Value(), c.W, c.Y.Value(), c.W }
func (c XYZW[T]) XWZX() (T, float64, T, T)                   { return c.X.Value(), c.W, c.Z.Value(), c.X.Value() }
func (c XYZW[T]) XWZY() (T, float64, T, T)                   { return c.X.Value(), c.W, c.Z.Value(), c.Y.Value() }
func (c XYZW[T]) XWZZ() (T, float64, T, T)                   { return c.X.Value(), c.W, c.Z.Value(), c.Z.Value() }
func (c XYZW[T]) XWZW() (T, float64, T, float64)             { return c.X.Value(), c.W, c.Z.Value(), c.W }
func (c XYZW[T]) XWWX() (T, float64, float64, T)             { return c.X.Value(), c.W, c.W, c.X.Value() }
func (c XYZW[T]) XWWY() (T, float64, float64, T)             { return c.X.Value(), c.W, c.W, c.Y.Value() }
func (c XYZW[T]) XWWZ() (T, float64, float64, T)             { return c.X.Value(), c.W, c.W, c.Z.Value() }
func (c XYZW[T]) XWWW() (T, float64, float64, float64)       { return c.X.Value(), c.W, c.W, c.W }
func (c XYZW[T]) YXXX() (T, T, T, T)                         { return c.Y.Value(), c.X.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T]) YXXY() (T, T, T, T)                         { return c.Y.Value(), c.X.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T]) YXXZ() (T, T, T, T)                         { return c.Y.Value(), c.X.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T]) YXXW() (T, T, T, float64)                   { return c.Y.Value(), c.X.Value(), c.X.Value(), c.W }
func (c XYZW[T]) YXYX() (T, T, T, T)                         { return c.Y.Value(), c.X.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T]) YXYY() (T, T, T, T)                         { return c.Y.Value(), c.X.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T]) YXYZ() (T, T, T, T)                         { return c.Y.Value(), c.X.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T]) YXYW() (T, T, T, float64)                   { return c.Y.Value(), c.X.Value(), c.Y.Value(), c.W }
func (c XYZW[T]) YXZX() (T, T, T, T)                         { return c.Y.Value(), c.X.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T]) YXZY() (T, T, T, T)                         { return c.Y.Value(), c.X.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T]) YXZZ() (T, T, T, T)                         { return c.Y.Value(), c.X.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T]) YXZW() (T, T, T, float64)                   { return c.Y.Value(), c.X.Value(), c.Z.Value(), c.W }
func (c XYZW[T]) YXWX() (T, T, float64, T)                   { return c.Y.Value(), c.X.Value(), c.W, c.X.Value() }
func (c XYZW[T]) YXWY() (T, T, float64, T)                   { return c.Y.Value(), c.X.Value(), c.W, c.Y.Value() }
func (c XYZW[T]) YXWZ() (T, T, float64, T)                   { return c.Y.Value(), c.X.Value(), c.W, c.Z.Value() }
func (c XYZW[T]) YXWW() (T, T, float64, float64)             { return c.Y.Value(), c.X.Value(), c.W, c.W }
func (c XYZW[T]) YYXX() (T, T, T, T)                         { return c.Y.Value(), c.Y.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T]) YYXY() (T, T, T, T)                         { return c.Y.Value(), c.Y.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T]) YYXZ() (T, T, T, T)                         { return c.Y.Value(), c.Y.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T]) YYXW() (T, T, T, float64)                   { return c.Y.Value(), c.Y.Value(), c.X.Value(), c.W }
func (c XYZW[T]) YYYX() (T, T, T, T)                         { return c.Y.Value(), c.Y.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T]) YYYY() (T, T, T, T)                         { return c.Y.Value(), c.Y.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T]) YYYZ() (T, T, T, T)                         { return c.Y.Value(), c.Y.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T]) YYYW() (T, T, T, float64)                   { return c.Y.Value(), c.Y.Value(), c.Y.Value(), c.W }
func (c XYZW[T]) YYZX() (T, T, T, T)                         { return c.Y.Value(), c.Y.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T]) YYZY() (T, T, T, T)                         { return c.Y.Value(), c.Y.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T]) YYZZ() (T, T, T, T)                         { return c.Y.Value(), c.Y.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T]) YYZW() (T, T, T, float64)                   { return c.Y.Value(), c.Y.Value(), c.Z.Value(), c.W }
func (c XYZW[T]) YYWX() (T, T, float64, T)                   { return c.Y.Value(), c.Y.Value(), c.W, c.X.Value() }
func (c XYZW[T]) YYWY() (T, T, float64, T)                   { return c.Y.Value(), c.Y.Value(), c.W, c.Y.Value() }
func (c XYZW[T]) YYWZ() (T, T, float64, T)                   { return c.Y.Value(), c.Y.Value(), c.W, c.Z.Value() }
func (c XYZW[T]) YYWW() (T, T, float64, float64)             { return c.Y.Value(), c.Y.Value(), c.W, c.W }
func (c XYZW[T]) YZXX() (T, T, T, T)                         { return c.Y.Value(), c.Z.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T]) YZXY() (T, T, T, T)                         { return c.Y.Value(), c.Z.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T]) YZXZ() (T, T, T, T)                         { return c.Y.Value(), c.Z.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T]) YZXW() (T, T, T, float64)                   { return c.Y.Value(), c.Z.Value(), c.X.Value(), c.W }
func (c XYZW[T]) YZYX() (T, T, T, T)                         { return c.Y.Value(), c.Z.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T]) YZYY() (T, T, T, T)                         { return c.Y.Value(), c.Z.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T]) YZYZ() (T, T, T, T)                         { return c.Y.Value(), c.Z.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T]) YZYW() (T, T, T, float64)                   { return c.Y.Value(), c.Z.Value(), c.Y.Value(), c.W }
func (c XYZW[T]) YZZX() (T, T, T, T)                         { return c.Y.Value(), c.Z.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T]) YZZY() (T, T, T, T)                         { return c.Y.Value(), c.Z.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T]) YZZZ() (T, T, T, T)                         { return c.Y.Value(), c.Z.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T]) YZZW() (T, T, T, float64)                   { return c.Y.Value(), c.Z.Value(), c.Z.Value(), c.W }
func (c XYZW[T]) YZWX() (T, T, float64, T)                   { return c.Y.Value(), c.Z.Value(), c.W, c.X.Value() }
func (c XYZW[T]) YZWY() (T, T, float64, T)                   { return c.Y.Value(), c.Z.Value(), c.W, c.Y.Value() }
func (c XYZW[T]) YZWZ() (T, T, float64, T)                   { return c.Y.Value(), c.Z.Value(), c.W, c.Z.Value() }
func (c XYZW[T]) YZWW() (T, T, float64, float64)             { return c.Y.Value(), c.Z.Value(), c.W, c.W }
func (c XYZW[T]) YWXX() (T, float64, T, T)                   { return c.Y.Value(), c.W, c.X.Value(), c.X.Value() }
func (c XYZW[T]) YWXY() (T, float64, T, T)                   { return c.Y.Value(), c.W, c.X.Value(), c.Y.Value() }
func (c XYZW[T]) YWXZ() (T, float64, T, T)                   { return c.Y.Value(), c.W, c.X.Value(), c.Z.Value() }
func (c XYZW[T]) YWXW() (T, float64, T, float64)             { return c.Y.Value(), c.W, c.X.Value(), c.W }
func (c XYZW[T]) YWYX() (T, float64, T, T)                   { return c.Y.Value(), c.W, c.Y.Value(), c.X.Value() }
func (c XYZW[T]) YWYY() (T, float64, T, T)                   { return c.Y.Value(), c.W, c.Y.Value(), c.Y.Value() }
func (c XYZW[T]) YWYZ() (T, float64, T, T)                   { return c.Y.Value(), c.W, c.Y.Value(), c.Z.Value() }
func (c XYZW[T]) YWYW() (T, float64, T, float64)             { return c.Y.Value(), c.W, c.Y.Value(), c.W }
func (c XYZW[T]) YWZX() (T, float64, T, T)                   { return c.Y.Value(), c.W, c.Z.Value(), c.X.Value() }
func (c XYZW[T]) YWZY() (T, float64, T, T)                   { return c.Y.Value(), c.W, c.Z.Value(), c.Y.Value() }
func (c XYZW[T]) YWZZ() (T, float64, T, T)                   { return c.Y.Value(), c.W, c.Z.Value(), c.Z.Value() }
func (c XYZW[T]) YWZW() (T, float64, T, float64)             { return c.Y.Value(), c.W, c.Z.Value(), c.W }
func (c XYZW[T]) YWWX() (T, float64, float64, T)             { return c.Y.Value(), c.W, c.W, c.X.Value() }
func (c XYZW[T]) YWWY() (T, float64, float64, T)             { return c.Y.Value(), c.W, c.W, c.Y.Value() }
func (c XYZW[T]) YWWZ() (T, float64, float64, T)             { return c.Y.Value(), c.W, c.W, c.Z.Value() }
func (c XYZW[T]) YWWW() (T, float64, float64, float64)       { return c.Y.Value(), c.W, c.W, c.W }
func (c XYZW[T]) ZXXX() (T, T, T, T)                         { return c.Z.Value(), c.X.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T]) ZXXY() (T, T, T, T)                         { return c.Z.Value(), c.X.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T]) ZXXZ() (T, T, T, T)                         { return c.Z.Value(), c.X.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T]) ZXXW() (T, T, T, float64)                   { return c.Z.Value(), c.X.Value(), c.X.Value(), c.W }
func (c XYZW[T]) ZXYX() (T, T, T, T)                         { return c.Z.Value(), c.X.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T]) ZXYY() (T, T, T, T)                         { return c.Z.Value(), c.X.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T]) ZXYZ() (T, T, T, T)                         { return c.Z.Value(), c.X.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T]) ZXYW() (T, T, T, float64)                   { return c.Z.Value(), c.X.Value(), c.Y.Value(), c.W }
func (c XYZW[T]) ZXZX() (T, T, T, T)                         { return c.Z.Value(), c.X.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T]) ZXZY() (T, T, T, T)                         { return c.Z.Value(), c.X.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T]) ZXZZ() (T, T, T, T)                         { return c.Z.Value(), c.X.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T]) ZXZW() (T, T, T, float64)                   { return c.Z.Value(), c.X.Value(), c.Z.Value(), c.W }
func (c XYZW[T]) ZXWX() (T, T, float64, T)                   { return c.Z.Value(), c.X.Value(), c.W, c.X.Value() }
func (c XYZW[T]) ZXWY() (T, T, float64, T)                   { return c.Z.Value(), c.X.Value(), c.W, c.Y.Value() }
func (c XYZW[T]) ZXWZ() (T, T, float64, T)                   { return c.Z.Value(), c.X.Value(), c.W, c.Z.Value() }
func (c XYZW[T]) ZXWW() (T, T, float64, float64)             { return c.Z.Value(), c.X.Value(), c.W, c.W }
func (c XYZW[T]) ZYXX() (T, T, T, T)                         { return c.Z.Value(), c.Y.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T]) ZYXY() (T, T, T, T)                         { return c.Z.Value(), c.Y.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T]) ZYXZ() (T, T, T, T)                         { return c.Z.Value(), c.Y.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T]) ZYXW() (T, T, T, float64)                   { return c.Z.Value(), c.Y.Value(), c.X.Value(), c.W }
func (c XYZW[T]) ZYYX() (T, T, T, T)                         { return c.Z.Value(), c.Y.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T]) ZYYY() (T, T, T, T)                         { return c.Z.Value(), c.Y.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T]) ZYYZ() (T, T, T, T)                         { return c.Z.Value(), c.Y.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T]) ZYYW() (T, T, T, float64)                   { return c.Z.Value(), c.Y.Value(), c.Y.Value(), c.W }
func (c XYZW[T]) ZYZX() (T, T, T, T)                         { return c.Z.Value(), c.Y.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T]) ZYZY() (T, T, T, T)                         { return c.Z.Value(), c.Y.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T]) ZYZZ() (T, T, T, T)                         { return c.Z.Value(), c.Y.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T]) ZYZW() (T, T, T, float64)                   { return c.Z.Value(), c.Y.Value(), c.Z.Value(), c.W }
func (c XYZW[T]) ZYWX() (T, T, float64, T)                   { return c.Z.Value(), c.Y.Value(), c.W, c.X.Value() }
func (c XYZW[T]) ZYWY() (T, T, float64, T)                   { return c.Z.Value(), c.Y.Value(), c.W, c.Y.Value() }
func (c XYZW[T]) ZYWZ() (T, T, float64, T)                   { return c.Z.Value(), c.Y.Value(), c.W, c.Z.Value() }
func (c XYZW[T]) ZYWW() (T, T, float64, float64)             { return c.Z.Value(), c.Y.Value(), c.W, c.W }
func (c XYZW[T]) ZZXX() (T, T, T, T)                         { return c.Z.Value(), c.Z.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T]) ZZXY() (T, T, T, T)                         { return c.Z.Value(), c.Z.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T]) ZZXZ() (T, T, T, T)                         { return c.Z.Value(), c.Z.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T]) ZZXW() (T, T, T, float64)                   { return c.Z.Value(), c.Z.Value(), c.X.Value(), c.W }
func (c XYZW[T]) ZZYX() (T, T, T, T)                         { return c.Z.Value(), c.Z.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T]) ZZYY() (T, T, T, T)                         { return c.Z.Value(), c.Z.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T]) ZZYZ() (T, T, T, T)                         { return c.Z.Value(), c.Z.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T]) ZZYW() (T, T, T, float64)                   { return c.Z.Value(), c.Z.Value(), c.Y.Value(), c.W }
func (c XYZW[T]) ZZZX() (T, T, T, T)                         { return c.Z.Value(), c.Z.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T]) ZZZY() (T, T, T, T)                         { return c.Z.Value(), c.Z.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T]) ZZZZ() (T, T, T, T)                         { return c.Z.Value(), c.Z.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T]) ZZZW() (T, T, T, float64)                   { return c.Z.Value(), c.Z.Value(), c.Z.Value(), c.W }
func (c XYZW[T]) ZZWX() (T, T, float64, T)                   { return c.Z.Value(), c.Z.Value(), c.W, c.X.Value() }
func (c XYZW[T]) ZZWY() (T, T, float64, T)                   { return c.Z.Value(), c.Z.Value(), c.W, c.Y.Value() }
func (c XYZW[T]) ZZWZ() (T, T, float64, T)                   { return c.Z.Value(), c.Z.Value(), c.W, c.Z.Value() }
func (c XYZW[T]) ZZWW() (T, T, float64, float64)             { return c.Z.Value(), c.Z.Value(), c.W, c.W }
func (c XYZW[T]) ZWXX() (T, float64, T, T)                   { return c.Z.Value(), c.W, c.X.Value(), c.X.Value() }
func (c XYZW[T]) ZWXY() (T, float64, T, T)                   { return c.Z.Value(), c.W, c.X.Value(), c.Y.Value() }
func (c XYZW[T]) ZWXZ() (T, float64, T, T)                   { return c.Z.Value(), c.W, c.X.Value(), c.Z.Value() }
func (c XYZW[T]) ZWXW() (T, float64, T, float64)             { return c.Z.Value(), c.W, c.X.Value(), c.W }
func (c XYZW[T]) ZWYX() (T, float64, T, T)                   { return c.Z.Value(), c.W, c.Y.Value(), c.X.Value() }
func (c XYZW[T]) ZWYY() (T, float64, T, T)                   { return c.Z.Value(), c.W, c.Y.Value(), c.Y.Value() }
func (c XYZW[T]) ZWYZ() (T, float64, T, T)                   { return c.Z.Value(), c.W, c.Y.Value(), c.Z.Value() }
func (c XYZW[T]) ZWYW() (T, float64, T, float64)             { return c.Z.Value(), c.W, c.Y.Value(), c.W }
func (c XYZW[T]) ZWZX() (T, float64, T, T)                   { return c.Z.Value(), c.W, c.Z.Value(), c.X.Value() }
func (c XYZW[T]) ZWZY() (T, float64, T, T)                   { return c.Z.Value(), c.W, c.Z.Value(), c.Y.Value() }
func (c XYZW[T]) ZWZZ() (T, float64, T, T)                   { return c.Z.Value(), c.W, c.Z.Value(), c.Z.Value() }
func (c XYZW[T]) ZWZW() (T, float64, T, float64)             { return c.Z.Value(), c.W, c.Z.Value(), c.W }
func (c XYZW[T]) ZWWX() (T, float64, float64, T)             { return c.Z.Value(), c.W, c.W, c.X.Value() }
func (c XYZW[T]) ZWWY() (T, float64, float64, T)             { return c.Z.Value(), c.W, c.W, c.Y.Value() }
func (c XYZW[T]) ZWWZ() (T, float64, float64, T)             { return c.Z.Value(), c.W, c.W, c.Z.Value() }
func (c XYZW[T]) ZWWW() (T, float64, float64, float64)       { return c.Z.Value(), c.W, c.W, c.W }
func (c XYZW[T]) WXXX() (float64, T, T, T)                   { return c.W, c.X.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T]) WXXY() (float64, T, T, T)                   { return c.W, c.X.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T]) WXXZ() (float64, T, T, T)                   { return c.W, c.X.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T]) WXXW() (float64, T, T, float64)             { return c.W, c.X.Value(), c.X.Value(), c.W }
func (c XYZW[T]) WXYX() (float64, T, T, T)                   { return c.W, c.X.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T]) WXYY() (float64, T, T, T)                   { return c.W, c.X.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T]) WXYZ() (float64, T, T, T)                   { return c.W, c.X.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T]) WXYW() (float64, T, T, float64)             { return c.W, c.X.Value(), c.Y.Value(), c.W }
func (c XYZW[T]) WXZX() (float64, T, T, T)                   { return c.W, c.X.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T]) WXZY() (float64, T, T, T)                   { return c.W, c.X.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T]) WXZZ() (float64, T, T, T)                   { return c.W, c.X.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T]) WXZW() (float64, T, T, float64)             { return c.W, c.X.Value(), c.Z.Value(), c.W }
func (c XYZW[T]) WXWX() (float64, T, float64, T)             { return c.W, c.X.Value(), c.W, c.X.Value() }
func (c XYZW[T]) WXWY() (float64, T, float64, T)             { return c.W, c.X.Value(), c.W, c.Y.Value() }
func (c XYZW[T]) WXWZ() (float64, T, float64, T)             { return c.W, c.X.Value(), c.W, c.Z.Value() }
func (c XYZW[T]) WXWW() (float64, T, float64, float64)       { return c.W, c.X.Value(), c.W, c.W }
func (c XYZW[T]) WYXX() (float64, T, T, T)                   { return c.W, c.Y.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T]) WYXY() (float64, T, T, T)                   { return c.W, c.Y.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T]) WYXZ() (float64, T, T, T)                   { return c.W, c.Y.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T]) WYXW() (float64, T, T, float64)             { return c.W, c.Y.Value(), c.X.Value(), c.W }
func (c XYZW[T]) WYYX() (float64, T, T, T)                   { return c.W, c.Y.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T]) WYYY() (float64, T, T, T)                   { return c.W, c.Y.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T]) WYYZ() (float64, T, T, T)                   { return c.W, c.Y.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T]) WYYW() (float64, T, T, float64)             { return c.W, c.Y.Value(), c.Y.Value(), c.W }
func (c XYZW[T]) WYZX() (float64, T, T, T)                   { return c.W, c.Y.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T]) WYZY() (float64, T, T, T)                   { return c.W, c.Y.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T]) WYZZ() (float64, T, T, T)                   { return c.W, c.Y.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T]) WYZW() (float64, T, T, float64)             { return c.W, c.Y.Value(), c.Z.Value(), c.W }
func (c XYZW[T]) WYWX() (float64, T, float64, T)             { return c.W, c.Y.Value(), c.W, c.X.Value() }
func (c XYZW[T]) WYWY() (float64, T, float64, T)             { return c.W, c.Y.Value(), c.W, c.Y.Value() }
func (c XYZW[T]) WYWZ() (float64, T, float64, T)             { return c.W, c.Y.Value(), c.W, c.Z.Value() }
func (c XYZW[T]) WYWW() (float64, T, float64, float64)       { return c.W, c.Y.Value(), c.W, c.W }
func (c XYZW[T]) WZXX() (float64, T, T, T)                   { return c.W, c.Z.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T]) WZXY() (float64, T, T, T)                   { return c.W, c.Z.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T]) WZXZ() (float64, T, T, T)                   { return c.W, c.Z.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T]) WZXW() (float64, T, T, float64)             { return c.W, c.Z.Value(), c.X.Value(), c.W }
func (c XYZW[T]) WZYX() (float64, T, T, T)                   { return c.W, c.Z.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T]) WZYY() (float64, T, T, T)                   { return c.W, c.Z.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T]) WZYZ() (float64, T, T, T)                   { return c.W, c.Z.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T]) WZYW() (float64, T, T, float64)             { return c.W, c.Z.Value(), c.Y.Value(), c.W }
func (c XYZW[T]) WZZX() (float64, T, T, T)                   { return c.W, c.Z.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T]) WZZY() (float64, T, T, T)                   { return c.W, c.Z.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T]) WZZZ() (float64, T, T, T)                   { return c.W, c.Z.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T]) WZZW() (float64, T, T, float64)             { return c.W, c.Z.Value(), c.Z.Value(), c.W }
func (c XYZW[T]) WZWX() (float64, T, float64, T)             { return c.W, c.Z.Value(), c.W, c.X.Value() }
func (c XYZW[T]) WZWY() (float64, T, float64, T)             { return c.W, c.Z.Value(), c.W, c.Y.Value() }
func (c XYZW[T]) WZWZ() (float64, T, float64, T)             { return c.W, c.Z.Value(), c.W, c.Z.Value() }
func (c XYZW[T]) WZWW() (float64, T, float64, float64)       { return c.W, c.Z.Value(), c.W, c.W }
func (c XYZW[T]) WWXX() (float64, float64, T, T)             { return c.W, c.W, c.X.Value(), c.X.Value() }
func (c XYZW[T]) WWXY() (float64, float64, T, T)             { return c.W, c.W, c.X.Value(), c.Y.Value() }
func (c XYZW[T]) WWXZ() (float64, float64, T, T)             { return c.W, c.W, c.X.Value(), c.Z.Value() }
func (c XYZW[T]) WWXW() (float64, float64, T, float64)       { return c.W, c.W, c.X.Value(), c.W }
func (c XYZW[T]) WWYX() (float64, float64, T, T)             { return c.W, c.W, c.Y.Value(), c.X.Value() }
func (c XYZW[T]) WWYY() (float64, float64, T, T)             { return c.W, c.W, c.Y.Value(), c.Y.Value() }
func (c XYZW[T]) WWYZ() (float64, float64, T, T)             { return c.W, c.W, c.Y.Value(), c.Z.Value() }
func (c XYZW[T]) WWYW() (float64, float64, T, float64)       { return c.W, c.W, c.Y.Value(), c.W }
func (c XYZW[T]) WWZX() (float64, float64, T, T)             { return c.W, c.W, c.Z.Value(), c.X.Value() }
func (c XYZW[T]) WWZY() (float64, float64, T, T)             { return c.W, c.W, c.Z.Value(), c.Y.Value() }
func (c XYZW[T]) WWZZ() (float64, float64, T, T)             { return c.W, c.W, c.Z.Value(), c.Z.Value() }
func (c XYZW[T]) WWZW() (float64, float64, T, float64)       { return c.W, c.W, c.Z.Value(), c.W }
func (c XYZW[T]) WWWX() (float64, float64, float64, T)       { return c.W, c.W, c.W, c.X.Value() }
func (c XYZW[T]) WWWY() (float64, float64, float64, T)       { return c.W, c.W, c.W, c.Y.Value() }
func (c XYZW[T]) WWWZ() (float64, float64, float64, T)       { return c.W, c.W, c.W, c.Z.Value() }
func (c XYZW[T]) WWWW() (float64, float64, float64, float64) { return c.W, c.W, c.W, c.W }
