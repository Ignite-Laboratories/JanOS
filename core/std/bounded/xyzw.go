package bounded

import (
	"fmt"
	"github.com/ignite-laboratories/core/std"
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
// However, this type is merely a general container for three like-types and one unlike type - please
// do not feel obligated to follow 'homogeneous coordinate' rules only.
//
// NOTE: This type also provides rudimentary "swizzling."
type XYZW[T num.Primitive, TW num.Primitive] struct {
	X std.Bounded[T]
	Y std.Bounded[T]
	Z std.Bounded[T]
	W std.Bounded[TW]
}

// SetClamp sets whether the directions should clamp to their boundaries or overflow and under-flow.
func (coords XYZW[T, TW]) SetClamp(shouldClamp bool) XYZW[T, TW] {
	coords.X.Clamp = shouldClamp
	coords.Y.Clamp = shouldClamp
	coords.Z.Clamp = shouldClamp
	return coords
}

// Set sets the coordinate values.
//
// NOTE: If no w is provided, the value remains unchanged.
func (coords XYZW[T, TW]) Set(x, y, z T, w ...TW) XYZW[T, TW] {
	coords.X, _ = coords.X.Set(x)
	coords.Y, _ = coords.Y.Set(y)
	coords.Z, _ = coords.Z.Set(z)
	if len(w) > 0 {
		coords.W, _ = coords.W.Set(w[0])
	}
	return coords
}

// SetBoundaries inclusively sets the coordinate boundaries for all directions.
//
// NOTE: This means to represent 1024x768 you should use 1023x767 =)
func (coords XYZW[T, TW]) SetBoundaries(minX, maxX, minY, maxY, minZ, maxZ T, minW, maxW TW) XYZW[T, TW] {
	coords.X, _ = coords.X.SetBoundaries(minX, maxX)
	coords.Y, _ = coords.Y.SetBoundaries(minY, maxY)
	coords.Z, _ = coords.Z.SetBoundaries(minZ, maxZ)
	coords.W, _ = coords.W.SetBoundaries(minW, maxW)
	return coords
}

// SetAll first sets the boundaries for each direction, then sets their directional values.
func (coords XYZW[T, TW]) SetAll(x, y, z T, w TW, minX, maxX, minY, maxY, minZ, maxZ T, minW, maxW TW) XYZW[T, TW] {
	return coords.SetBoundaries(minX, maxX, minY, maxY, minZ, maxZ, minW, maxW).Set(x, y, z, w)
}

// SetFromNormalized sets the bounded directional values using float64 unit vectors from the [0.0, 1.0]
// range, where 0.0 maps to the coordinate space's bounded minimum and 1.0 maps to the bounded maximum.
//
// NOTE: If no w is provided, the value remains unchanged.
func (coords XYZW[T, TW]) SetFromNormalized(x, y, z float64, w ...float64) XYZW[T, TW] {
	coords.X, _ = coords.X.SetFromNormalized(x)
	coords.Y, _ = coords.Y.SetFromNormalized(y)
	coords.Z, _ = coords.Z.SetFromNormalized(z)
	if len(w) > 0 {
		coords.W, _ = coords.W.SetFromNormalized(w[0])
	}
	return coords
}

// SetFromNormalized32 sets the bounded directional values using float32 unit vectors from the [0.0, 1.0]
// range, where 0.0 maps to the coordinate space's bounded minimum and 1.0 maps to the bounded maximum.
//
// NOTE: If no w is provided, the value remains unchanged.
func (coords XYZW[T, TW]) SetFromNormalized32(x, y, z float32, w ...float32) XYZW[T, TW] {
	w64 := make([]float64, len(w))
	for i, v := range w {
		w64[i] = float64(v)
	}

	return coords.SetFromNormalized(float64(x), float64(y), float64(z), w64...)
}

// Normalize converts the bounded directional values to float64 unit vectors in the range [0.0, 1.0],
// where the coordinate space's bounded minimum maps to 0.0 and the bounded maximum maps to 1.0.
func (coords XYZW[T, TW]) Normalize() (x float64, y float64, z float64, w float64) {
	return coords.X.Normalize(), coords.Y.Normalize(), coords.Z.Normalize(), coords.W.Normalize()
}

// Normalize32 converts the bounded directional values to float32 unit vectors in the range [0.0, 1.0],
// where the coordinate space's bounded minimum maps to 0.0 and the bounded maximum maps to 1.0.
func (coords XYZW[T, TW]) Normalize32() (x float32, y float32, z float32, w float32) {
	return coords.X.Normalize32(), coords.Y.Normalize32(), coords.Z.Normalize32(), coords.W.Normalize32()
}

func (coords XYZW[T, TW]) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v)", coords.X.Value(), coords.Y.Value(), coords.Z.Value(), coords.W.Value())
}

/**
Swizzling

NOTE: This is a regular expression to find and replace swizzle functions into a one-liner if the auto formatter ever kicks in

Find -
func \((.*?)\) ([A-Z]{2,4})\(\) \((.*?)\)[ ]*\{[\n\t ]*return(.*?)[\n\t ]*\}

Replace -
func ($1) $2() ($3) { return$4 }
*/

func (c XYZW[T, TW]) XX() (T, T)   { return c.X.Value(), c.X.Value() }
func (c XYZW[T, TW]) XY() (T, T)   { return c.X.Value(), c.Y.Value() }
func (c XYZW[T, TW]) XZ() (T, T)   { return c.X.Value(), c.Z.Value() }
func (c XYZW[T, TW]) XW() (T, TW)  { return c.X.Value(), c.W.Value() }
func (c XYZW[T, TW]) YX() (T, T)   { return c.Y.Value(), c.X.Value() }
func (c XYZW[T, TW]) YY() (T, T)   { return c.Y.Value(), c.Y.Value() }
func (c XYZW[T, TW]) YZ() (T, T)   { return c.Y.Value(), c.Z.Value() }
func (c XYZW[T, TW]) YW() (T, TW)  { return c.Y.Value(), c.W.Value() }
func (c XYZW[T, TW]) ZX() (T, T)   { return c.Z.Value(), c.X.Value() }
func (c XYZW[T, TW]) ZY() (T, T)   { return c.Z.Value(), c.Y.Value() }
func (c XYZW[T, TW]) ZZ() (T, T)   { return c.Z.Value(), c.Z.Value() }
func (c XYZW[T, TW]) ZW() (T, TW)  { return c.Z.Value(), c.W.Value() }
func (c XYZW[T, TW]) WX() (TW, T)  { return c.W.Value(), c.X.Value() }
func (c XYZW[T, TW]) WY() (TW, T)  { return c.W.Value(), c.Y.Value() }
func (c XYZW[T, TW]) WZ() (TW, T)  { return c.W.Value(), c.Z.Value() }
func (c XYZW[T, TW]) WW() (TW, TW) { return c.W.Value(), c.W.Value() }

func (c XYZW[T, TW]) XXX() (T, T, T)    { return c.X.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T, TW]) XXY() (T, T, T)    { return c.X.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T, TW]) XXZ() (T, T, T)    { return c.X.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T, TW]) XXW() (T, T, TW)   { return c.X.Value(), c.X.Value(), c.W.Value() }
func (c XYZW[T, TW]) XYX() (T, T, T)    { return c.X.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T, TW]) XYY() (T, T, T)    { return c.X.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T, TW]) XYZ() (T, T, T)    { return c.X.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T, TW]) XYW() (T, T, TW)   { return c.X.Value(), c.Y.Value(), c.W.Value() }
func (c XYZW[T, TW]) XZX() (T, T, T)    { return c.X.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T, TW]) XZY() (T, T, T)    { return c.X.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T, TW]) XZZ() (T, T, T)    { return c.X.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T, TW]) XZW() (T, T, TW)   { return c.X.Value(), c.Z.Value(), c.W.Value() }
func (c XYZW[T, TW]) XWX() (T, TW, T)   { return c.X.Value(), c.W.Value(), c.X.Value() }
func (c XYZW[T, TW]) XWY() (T, TW, T)   { return c.X.Value(), c.W.Value(), c.Y.Value() }
func (c XYZW[T, TW]) XWZ() (T, TW, T)   { return c.X.Value(), c.W.Value(), c.Z.Value() }
func (c XYZW[T, TW]) XWW() (T, TW, TW)  { return c.X.Value(), c.W.Value(), c.W.Value() }
func (c XYZW[T, TW]) YXX() (T, T, T)    { return c.Y.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T, TW]) YXY() (T, T, T)    { return c.Y.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T, TW]) YXZ() (T, T, T)    { return c.Y.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T, TW]) YXW() (T, T, TW)   { return c.Y.Value(), c.X.Value(), c.W.Value() }
func (c XYZW[T, TW]) YYX() (T, T, T)    { return c.Y.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T, TW]) YYY() (T, T, T)    { return c.Y.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T, TW]) YYZ() (T, T, T)    { return c.Y.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T, TW]) YYW() (T, T, TW)   { return c.Y.Value(), c.Y.Value(), c.W.Value() }
func (c XYZW[T, TW]) YZX() (T, T, T)    { return c.Y.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T, TW]) YZY() (T, T, T)    { return c.Y.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T, TW]) YZZ() (T, T, T)    { return c.Y.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T, TW]) YZW() (T, T, TW)   { return c.Y.Value(), c.Z.Value(), c.W.Value() }
func (c XYZW[T, TW]) YWX() (T, TW, T)   { return c.Y.Value(), c.W.Value(), c.X.Value() }
func (c XYZW[T, TW]) YWY() (T, TW, T)   { return c.Y.Value(), c.W.Value(), c.Y.Value() }
func (c XYZW[T, TW]) YWZ() (T, TW, T)   { return c.Y.Value(), c.W.Value(), c.Z.Value() }
func (c XYZW[T, TW]) YWW() (T, TW, TW)  { return c.Y.Value(), c.W.Value(), c.W.Value() }
func (c XYZW[T, TW]) ZXX() (T, T, T)    { return c.Z.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T, TW]) ZXY() (T, T, T)    { return c.Z.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T, TW]) ZXZ() (T, T, T)    { return c.Z.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T, TW]) ZXW() (T, T, TW)   { return c.Z.Value(), c.X.Value(), c.W.Value() }
func (c XYZW[T, TW]) ZYX() (T, T, T)    { return c.Z.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T, TW]) ZYY() (T, T, T)    { return c.Z.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T, TW]) ZYZ() (T, T, T)    { return c.Z.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T, TW]) ZYW() (T, T, TW)   { return c.Z.Value(), c.Y.Value(), c.W.Value() }
func (c XYZW[T, TW]) ZZX() (T, T, T)    { return c.Z.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T, TW]) ZZY() (T, T, T)    { return c.Z.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T, TW]) ZZZ() (T, T, T)    { return c.Z.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T, TW]) ZZW() (T, T, TW)   { return c.Z.Value(), c.Z.Value(), c.W.Value() }
func (c XYZW[T, TW]) ZWX() (T, TW, T)   { return c.Z.Value(), c.W.Value(), c.X.Value() }
func (c XYZW[T, TW]) ZWY() (T, TW, T)   { return c.Z.Value(), c.W.Value(), c.Y.Value() }
func (c XYZW[T, TW]) ZWZ() (T, TW, T)   { return c.Z.Value(), c.W.Value(), c.Z.Value() }
func (c XYZW[T, TW]) ZWW() (T, TW, TW)  { return c.Z.Value(), c.W.Value(), c.W.Value() }
func (c XYZW[T, TW]) WXX() (TW, T, T)   { return c.W.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T, TW]) WXY() (TW, T, T)   { return c.W.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T, TW]) WXZ() (TW, T, T)   { return c.W.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T, TW]) WXW() (TW, T, TW)  { return c.W.Value(), c.X.Value(), c.W.Value() }
func (c XYZW[T, TW]) WYX() (TW, T, T)   { return c.W.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T, TW]) WYY() (TW, T, T)   { return c.W.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T, TW]) WYZ() (TW, T, T)   { return c.W.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T, TW]) WYW() (TW, T, TW)  { return c.W.Value(), c.Y.Value(), c.W.Value() }
func (c XYZW[T, TW]) WZX() (TW, T, T)   { return c.W.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T, TW]) WZY() (TW, T, T)   { return c.W.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T, TW]) WZZ() (TW, T, T)   { return c.W.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T, TW]) WZW() (TW, T, TW)  { return c.W.Value(), c.Z.Value(), c.W.Value() }
func (c XYZW[T, TW]) WWX() (TW, TW, T)  { return c.W.Value(), c.W.Value(), c.X.Value() }
func (c XYZW[T, TW]) WWY() (TW, TW, T)  { return c.W.Value(), c.W.Value(), c.Y.Value() }
func (c XYZW[T, TW]) WWZ() (TW, TW, T)  { return c.W.Value(), c.W.Value(), c.Z.Value() }
func (c XYZW[T, TW]) WWW() (TW, TW, TW) { return c.W.Value(), c.W.Value(), c.W.Value() }

func (c XYZW[T, TW]) XXXX() (T, T, T, T)   { return c.X.Value(), c.X.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T, TW]) XXXY() (T, T, T, T)   { return c.X.Value(), c.X.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T, TW]) XXXZ() (T, T, T, T)   { return c.X.Value(), c.X.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T, TW]) XXXW() (T, T, T, TW)  { return c.X.Value(), c.X.Value(), c.X.Value(), c.W.Value() }
func (c XYZW[T, TW]) XXYX() (T, T, T, T)   { return c.X.Value(), c.X.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T, TW]) XXYY() (T, T, T, T)   { return c.X.Value(), c.X.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T, TW]) XXYZ() (T, T, T, T)   { return c.X.Value(), c.X.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T, TW]) XXYW() (T, T, T, TW)  { return c.X.Value(), c.X.Value(), c.Y.Value(), c.W.Value() }
func (c XYZW[T, TW]) XXZX() (T, T, T, T)   { return c.X.Value(), c.X.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T, TW]) XXZY() (T, T, T, T)   { return c.X.Value(), c.X.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T, TW]) XXZZ() (T, T, T, T)   { return c.X.Value(), c.X.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T, TW]) XXZW() (T, T, T, TW)  { return c.X.Value(), c.X.Value(), c.Z.Value(), c.W.Value() }
func (c XYZW[T, TW]) XXWX() (T, T, TW, T)  { return c.X.Value(), c.X.Value(), c.W.Value(), c.X.Value() }
func (c XYZW[T, TW]) XXWY() (T, T, TW, T)  { return c.X.Value(), c.X.Value(), c.W.Value(), c.Y.Value() }
func (c XYZW[T, TW]) XXWZ() (T, T, TW, T)  { return c.X.Value(), c.X.Value(), c.W.Value(), c.Z.Value() }
func (c XYZW[T, TW]) XXWW() (T, T, TW, TW) { return c.X.Value(), c.X.Value(), c.W.Value(), c.W.Value() }
func (c XYZW[T, TW]) XYXX() (T, T, T, T)   { return c.X.Value(), c.Y.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T, TW]) XYXY() (T, T, T, T)   { return c.X.Value(), c.Y.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T, TW]) XYXZ() (T, T, T, T)   { return c.X.Value(), c.Y.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T, TW]) XYXW() (T, T, T, TW)  { return c.X.Value(), c.Y.Value(), c.X.Value(), c.W.Value() }
func (c XYZW[T, TW]) XYYX() (T, T, T, T)   { return c.X.Value(), c.Y.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T, TW]) XYYY() (T, T, T, T)   { return c.X.Value(), c.Y.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T, TW]) XYYZ() (T, T, T, T)   { return c.X.Value(), c.Y.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T, TW]) XYYW() (T, T, T, TW)  { return c.X.Value(), c.Y.Value(), c.Y.Value(), c.W.Value() }
func (c XYZW[T, TW]) XYZX() (T, T, T, T)   { return c.X.Value(), c.Y.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T, TW]) XYZY() (T, T, T, T)   { return c.X.Value(), c.Y.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T, TW]) XYZZ() (T, T, T, T)   { return c.X.Value(), c.Y.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T, TW]) XYZW() (T, T, T, TW)  { return c.X.Value(), c.Y.Value(), c.Z.Value(), c.W.Value() }
func (c XYZW[T, TW]) XYWX() (T, T, TW, T)  { return c.X.Value(), c.Y.Value(), c.W.Value(), c.X.Value() }
func (c XYZW[T, TW]) XYWY() (T, T, TW, T)  { return c.X.Value(), c.Y.Value(), c.W.Value(), c.Y.Value() }
func (c XYZW[T, TW]) XYWZ() (T, T, TW, T)  { return c.X.Value(), c.Y.Value(), c.W.Value(), c.Z.Value() }
func (c XYZW[T, TW]) XYWW() (T, T, TW, TW) { return c.X.Value(), c.Y.Value(), c.W.Value(), c.W.Value() }
func (c XYZW[T, TW]) XZXX() (T, T, T, T)   { return c.X.Value(), c.Z.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T, TW]) XZXY() (T, T, T, T)   { return c.X.Value(), c.Z.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T, TW]) XZXZ() (T, T, T, T)   { return c.X.Value(), c.Z.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T, TW]) XZXW() (T, T, T, TW)  { return c.X.Value(), c.Z.Value(), c.X.Value(), c.W.Value() }
func (c XYZW[T, TW]) XZYX() (T, T, T, T)   { return c.X.Value(), c.Z.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T, TW]) XZYY() (T, T, T, T)   { return c.X.Value(), c.Z.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T, TW]) XZYZ() (T, T, T, T)   { return c.X.Value(), c.Z.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T, TW]) XZYW() (T, T, T, TW)  { return c.X.Value(), c.Z.Value(), c.Y.Value(), c.W.Value() }
func (c XYZW[T, TW]) XZZX() (T, T, T, T)   { return c.X.Value(), c.Z.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T, TW]) XZZY() (T, T, T, T)   { return c.X.Value(), c.Z.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T, TW]) XZZZ() (T, T, T, T)   { return c.X.Value(), c.Z.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T, TW]) XZZW() (T, T, T, TW)  { return c.X.Value(), c.Z.Value(), c.Z.Value(), c.W.Value() }
func (c XYZW[T, TW]) XZWX() (T, T, TW, T)  { return c.X.Value(), c.Z.Value(), c.W.Value(), c.X.Value() }
func (c XYZW[T, TW]) XZWY() (T, T, TW, T)  { return c.X.Value(), c.Z.Value(), c.W.Value(), c.Y.Value() }
func (c XYZW[T, TW]) XZWZ() (T, T, TW, T)  { return c.X.Value(), c.Z.Value(), c.W.Value(), c.Z.Value() }
func (c XYZW[T, TW]) XZWW() (T, T, TW, TW) { return c.X.Value(), c.Z.Value(), c.W.Value(), c.W.Value() }
func (c XYZW[T, TW]) XWXX() (T, TW, T, T)  { return c.X.Value(), c.W.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T, TW]) XWXY() (T, TW, T, T)  { return c.X.Value(), c.W.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T, TW]) XWXZ() (T, TW, T, T)  { return c.X.Value(), c.W.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T, TW]) XWXW() (T, TW, T, TW) { return c.X.Value(), c.W.Value(), c.X.Value(), c.W.Value() }
func (c XYZW[T, TW]) XWYX() (T, TW, T, T)  { return c.X.Value(), c.W.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T, TW]) XWYY() (T, TW, T, T)  { return c.X.Value(), c.W.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T, TW]) XWYZ() (T, TW, T, T)  { return c.X.Value(), c.W.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T, TW]) XWYW() (T, TW, T, TW) { return c.X.Value(), c.W.Value(), c.Y.Value(), c.W.Value() }
func (c XYZW[T, TW]) XWZX() (T, TW, T, T)  { return c.X.Value(), c.W.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T, TW]) XWZY() (T, TW, T, T)  { return c.X.Value(), c.W.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T, TW]) XWZZ() (T, TW, T, T)  { return c.X.Value(), c.W.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T, TW]) XWZW() (T, TW, T, TW) { return c.X.Value(), c.W.Value(), c.Z.Value(), c.W.Value() }
func (c XYZW[T, TW]) XWWX() (T, TW, TW, T) { return c.X.Value(), c.W.Value(), c.W.Value(), c.X.Value() }
func (c XYZW[T, TW]) XWWY() (T, TW, TW, T) { return c.X.Value(), c.W.Value(), c.W.Value(), c.Y.Value() }
func (c XYZW[T, TW]) XWWZ() (T, TW, TW, T) { return c.X.Value(), c.W.Value(), c.W.Value(), c.Z.Value() }
func (c XYZW[T, TW]) XWWW() (T, TW, TW, TW) {
	return c.X.Value(), c.W.Value(), c.W.Value(), c.W.Value()
}
func (c XYZW[T, TW]) YXXX() (T, T, T, T)   { return c.Y.Value(), c.X.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T, TW]) YXXY() (T, T, T, T)   { return c.Y.Value(), c.X.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T, TW]) YXXZ() (T, T, T, T)   { return c.Y.Value(), c.X.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T, TW]) YXXW() (T, T, T, TW)  { return c.Y.Value(), c.X.Value(), c.X.Value(), c.W.Value() }
func (c XYZW[T, TW]) YXYX() (T, T, T, T)   { return c.Y.Value(), c.X.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T, TW]) YXYY() (T, T, T, T)   { return c.Y.Value(), c.X.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T, TW]) YXYZ() (T, T, T, T)   { return c.Y.Value(), c.X.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T, TW]) YXYW() (T, T, T, TW)  { return c.Y.Value(), c.X.Value(), c.Y.Value(), c.W.Value() }
func (c XYZW[T, TW]) YXZX() (T, T, T, T)   { return c.Y.Value(), c.X.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T, TW]) YXZY() (T, T, T, T)   { return c.Y.Value(), c.X.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T, TW]) YXZZ() (T, T, T, T)   { return c.Y.Value(), c.X.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T, TW]) YXZW() (T, T, T, TW)  { return c.Y.Value(), c.X.Value(), c.Z.Value(), c.W.Value() }
func (c XYZW[T, TW]) YXWX() (T, T, TW, T)  { return c.Y.Value(), c.X.Value(), c.W.Value(), c.X.Value() }
func (c XYZW[T, TW]) YXWY() (T, T, TW, T)  { return c.Y.Value(), c.X.Value(), c.W.Value(), c.Y.Value() }
func (c XYZW[T, TW]) YXWZ() (T, T, TW, T)  { return c.Y.Value(), c.X.Value(), c.W.Value(), c.Z.Value() }
func (c XYZW[T, TW]) YXWW() (T, T, TW, TW) { return c.Y.Value(), c.X.Value(), c.W.Value(), c.W.Value() }
func (c XYZW[T, TW]) YYXX() (T, T, T, T)   { return c.Y.Value(), c.Y.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T, TW]) YYXY() (T, T, T, T)   { return c.Y.Value(), c.Y.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T, TW]) YYXZ() (T, T, T, T)   { return c.Y.Value(), c.Y.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T, TW]) YYXW() (T, T, T, TW)  { return c.Y.Value(), c.Y.Value(), c.X.Value(), c.W.Value() }
func (c XYZW[T, TW]) YYYX() (T, T, T, T)   { return c.Y.Value(), c.Y.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T, TW]) YYYY() (T, T, T, T)   { return c.Y.Value(), c.Y.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T, TW]) YYYZ() (T, T, T, T)   { return c.Y.Value(), c.Y.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T, TW]) YYYW() (T, T, T, TW)  { return c.Y.Value(), c.Y.Value(), c.Y.Value(), c.W.Value() }
func (c XYZW[T, TW]) YYZX() (T, T, T, T)   { return c.Y.Value(), c.Y.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T, TW]) YYZY() (T, T, T, T)   { return c.Y.Value(), c.Y.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T, TW]) YYZZ() (T, T, T, T)   { return c.Y.Value(), c.Y.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T, TW]) YYZW() (T, T, T, TW)  { return c.Y.Value(), c.Y.Value(), c.Z.Value(), c.W.Value() }
func (c XYZW[T, TW]) YYWX() (T, T, TW, T)  { return c.Y.Value(), c.Y.Value(), c.W.Value(), c.X.Value() }
func (c XYZW[T, TW]) YYWY() (T, T, TW, T)  { return c.Y.Value(), c.Y.Value(), c.W.Value(), c.Y.Value() }
func (c XYZW[T, TW]) YYWZ() (T, T, TW, T)  { return c.Y.Value(), c.Y.Value(), c.W.Value(), c.Z.Value() }
func (c XYZW[T, TW]) YYWW() (T, T, TW, TW) { return c.Y.Value(), c.Y.Value(), c.W.Value(), c.W.Value() }
func (c XYZW[T, TW]) YZXX() (T, T, T, T)   { return c.Y.Value(), c.Z.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T, TW]) YZXY() (T, T, T, T)   { return c.Y.Value(), c.Z.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T, TW]) YZXZ() (T, T, T, T)   { return c.Y.Value(), c.Z.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T, TW]) YZXW() (T, T, T, TW)  { return c.Y.Value(), c.Z.Value(), c.X.Value(), c.W.Value() }
func (c XYZW[T, TW]) YZYX() (T, T, T, T)   { return c.Y.Value(), c.Z.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T, TW]) YZYY() (T, T, T, T)   { return c.Y.Value(), c.Z.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T, TW]) YZYZ() (T, T, T, T)   { return c.Y.Value(), c.Z.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T, TW]) YZYW() (T, T, T, TW)  { return c.Y.Value(), c.Z.Value(), c.Y.Value(), c.W.Value() }
func (c XYZW[T, TW]) YZZX() (T, T, T, T)   { return c.Y.Value(), c.Z.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T, TW]) YZZY() (T, T, T, T)   { return c.Y.Value(), c.Z.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T, TW]) YZZZ() (T, T, T, T)   { return c.Y.Value(), c.Z.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T, TW]) YZZW() (T, T, T, TW)  { return c.Y.Value(), c.Z.Value(), c.Z.Value(), c.W.Value() }
func (c XYZW[T, TW]) YZWX() (T, T, TW, T)  { return c.Y.Value(), c.Z.Value(), c.W.Value(), c.X.Value() }
func (c XYZW[T, TW]) YZWY() (T, T, TW, T)  { return c.Y.Value(), c.Z.Value(), c.W.Value(), c.Y.Value() }
func (c XYZW[T, TW]) YZWZ() (T, T, TW, T)  { return c.Y.Value(), c.Z.Value(), c.W.Value(), c.Z.Value() }
func (c XYZW[T, TW]) YZWW() (T, T, TW, TW) { return c.Y.Value(), c.Z.Value(), c.W.Value(), c.W.Value() }
func (c XYZW[T, TW]) YWXX() (T, TW, T, T)  { return c.Y.Value(), c.W.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T, TW]) YWXY() (T, TW, T, T)  { return c.Y.Value(), c.W.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T, TW]) YWXZ() (T, TW, T, T)  { return c.Y.Value(), c.W.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T, TW]) YWXW() (T, TW, T, TW) { return c.Y.Value(), c.W.Value(), c.X.Value(), c.W.Value() }
func (c XYZW[T, TW]) YWYX() (T, TW, T, T)  { return c.Y.Value(), c.W.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T, TW]) YWYY() (T, TW, T, T)  { return c.Y.Value(), c.W.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T, TW]) YWYZ() (T, TW, T, T)  { return c.Y.Value(), c.W.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T, TW]) YWYW() (T, TW, T, TW) { return c.Y.Value(), c.W.Value(), c.Y.Value(), c.W.Value() }
func (c XYZW[T, TW]) YWZX() (T, TW, T, T)  { return c.Y.Value(), c.W.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T, TW]) YWZY() (T, TW, T, T)  { return c.Y.Value(), c.W.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T, TW]) YWZZ() (T, TW, T, T)  { return c.Y.Value(), c.W.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T, TW]) YWZW() (T, TW, T, TW) { return c.Y.Value(), c.W.Value(), c.Z.Value(), c.W.Value() }
func (c XYZW[T, TW]) YWWX() (T, TW, TW, T) { return c.Y.Value(), c.W.Value(), c.W.Value(), c.X.Value() }
func (c XYZW[T, TW]) YWWY() (T, TW, TW, T) { return c.Y.Value(), c.W.Value(), c.W.Value(), c.Y.Value() }
func (c XYZW[T, TW]) YWWZ() (T, TW, TW, T) { return c.Y.Value(), c.W.Value(), c.W.Value(), c.Z.Value() }
func (c XYZW[T, TW]) YWWW() (T, TW, TW, TW) {
	return c.Y.Value(), c.W.Value(), c.W.Value(), c.W.Value()
}
func (c XYZW[T, TW]) ZXXX() (T, T, T, T)   { return c.Z.Value(), c.X.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T, TW]) ZXXY() (T, T, T, T)   { return c.Z.Value(), c.X.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T, TW]) ZXXZ() (T, T, T, T)   { return c.Z.Value(), c.X.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T, TW]) ZXXW() (T, T, T, TW)  { return c.Z.Value(), c.X.Value(), c.X.Value(), c.W.Value() }
func (c XYZW[T, TW]) ZXYX() (T, T, T, T)   { return c.Z.Value(), c.X.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T, TW]) ZXYY() (T, T, T, T)   { return c.Z.Value(), c.X.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T, TW]) ZXYZ() (T, T, T, T)   { return c.Z.Value(), c.X.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T, TW]) ZXYW() (T, T, T, TW)  { return c.Z.Value(), c.X.Value(), c.Y.Value(), c.W.Value() }
func (c XYZW[T, TW]) ZXZX() (T, T, T, T)   { return c.Z.Value(), c.X.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T, TW]) ZXZY() (T, T, T, T)   { return c.Z.Value(), c.X.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T, TW]) ZXZZ() (T, T, T, T)   { return c.Z.Value(), c.X.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T, TW]) ZXZW() (T, T, T, TW)  { return c.Z.Value(), c.X.Value(), c.Z.Value(), c.W.Value() }
func (c XYZW[T, TW]) ZXWX() (T, T, TW, T)  { return c.Z.Value(), c.X.Value(), c.W.Value(), c.X.Value() }
func (c XYZW[T, TW]) ZXWY() (T, T, TW, T)  { return c.Z.Value(), c.X.Value(), c.W.Value(), c.Y.Value() }
func (c XYZW[T, TW]) ZXWZ() (T, T, TW, T)  { return c.Z.Value(), c.X.Value(), c.W.Value(), c.Z.Value() }
func (c XYZW[T, TW]) ZXWW() (T, T, TW, TW) { return c.Z.Value(), c.X.Value(), c.W.Value(), c.W.Value() }
func (c XYZW[T, TW]) ZYXX() (T, T, T, T)   { return c.Z.Value(), c.Y.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T, TW]) ZYXY() (T, T, T, T)   { return c.Z.Value(), c.Y.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T, TW]) ZYXZ() (T, T, T, T)   { return c.Z.Value(), c.Y.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T, TW]) ZYXW() (T, T, T, TW)  { return c.Z.Value(), c.Y.Value(), c.X.Value(), c.W.Value() }
func (c XYZW[T, TW]) ZYYX() (T, T, T, T)   { return c.Z.Value(), c.Y.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T, TW]) ZYYY() (T, T, T, T)   { return c.Z.Value(), c.Y.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T, TW]) ZYYZ() (T, T, T, T)   { return c.Z.Value(), c.Y.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T, TW]) ZYYW() (T, T, T, TW)  { return c.Z.Value(), c.Y.Value(), c.Y.Value(), c.W.Value() }
func (c XYZW[T, TW]) ZYZX() (T, T, T, T)   { return c.Z.Value(), c.Y.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T, TW]) ZYZY() (T, T, T, T)   { return c.Z.Value(), c.Y.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T, TW]) ZYZZ() (T, T, T, T)   { return c.Z.Value(), c.Y.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T, TW]) ZYZW() (T, T, T, TW)  { return c.Z.Value(), c.Y.Value(), c.Z.Value(), c.W.Value() }
func (c XYZW[T, TW]) ZYWX() (T, T, TW, T)  { return c.Z.Value(), c.Y.Value(), c.W.Value(), c.X.Value() }
func (c XYZW[T, TW]) ZYWY() (T, T, TW, T)  { return c.Z.Value(), c.Y.Value(), c.W.Value(), c.Y.Value() }
func (c XYZW[T, TW]) ZYWZ() (T, T, TW, T)  { return c.Z.Value(), c.Y.Value(), c.W.Value(), c.Z.Value() }
func (c XYZW[T, TW]) ZYWW() (T, T, TW, TW) { return c.Z.Value(), c.Y.Value(), c.W.Value(), c.W.Value() }
func (c XYZW[T, TW]) ZZXX() (T, T, T, T)   { return c.Z.Value(), c.Z.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T, TW]) ZZXY() (T, T, T, T)   { return c.Z.Value(), c.Z.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T, TW]) ZZXZ() (T, T, T, T)   { return c.Z.Value(), c.Z.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T, TW]) ZZXW() (T, T, T, TW)  { return c.Z.Value(), c.Z.Value(), c.X.Value(), c.W.Value() }
func (c XYZW[T, TW]) ZZYX() (T, T, T, T)   { return c.Z.Value(), c.Z.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T, TW]) ZZYY() (T, T, T, T)   { return c.Z.Value(), c.Z.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T, TW]) ZZYZ() (T, T, T, T)   { return c.Z.Value(), c.Z.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T, TW]) ZZYW() (T, T, T, TW)  { return c.Z.Value(), c.Z.Value(), c.Y.Value(), c.W.Value() }
func (c XYZW[T, TW]) ZZZX() (T, T, T, T)   { return c.Z.Value(), c.Z.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T, TW]) ZZZY() (T, T, T, T)   { return c.Z.Value(), c.Z.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T, TW]) ZZZZ() (T, T, T, T)   { return c.Z.Value(), c.Z.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T, TW]) ZZZW() (T, T, T, TW)  { return c.Z.Value(), c.Z.Value(), c.Z.Value(), c.W.Value() }
func (c XYZW[T, TW]) ZZWX() (T, T, TW, T)  { return c.Z.Value(), c.Z.Value(), c.W.Value(), c.X.Value() }
func (c XYZW[T, TW]) ZZWY() (T, T, TW, T)  { return c.Z.Value(), c.Z.Value(), c.W.Value(), c.Y.Value() }
func (c XYZW[T, TW]) ZZWZ() (T, T, TW, T)  { return c.Z.Value(), c.Z.Value(), c.W.Value(), c.Z.Value() }
func (c XYZW[T, TW]) ZZWW() (T, T, TW, TW) { return c.Z.Value(), c.Z.Value(), c.W.Value(), c.W.Value() }
func (c XYZW[T, TW]) ZWXX() (T, TW, T, T)  { return c.Z.Value(), c.W.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T, TW]) ZWXY() (T, TW, T, T)  { return c.Z.Value(), c.W.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T, TW]) ZWXZ() (T, TW, T, T)  { return c.Z.Value(), c.W.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T, TW]) ZWXW() (T, TW, T, TW) { return c.Z.Value(), c.W.Value(), c.X.Value(), c.W.Value() }
func (c XYZW[T, TW]) ZWYX() (T, TW, T, T)  { return c.Z.Value(), c.W.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T, TW]) ZWYY() (T, TW, T, T)  { return c.Z.Value(), c.W.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T, TW]) ZWYZ() (T, TW, T, T)  { return c.Z.Value(), c.W.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T, TW]) ZWYW() (T, TW, T, TW) { return c.Z.Value(), c.W.Value(), c.Y.Value(), c.W.Value() }
func (c XYZW[T, TW]) ZWZX() (T, TW, T, T)  { return c.Z.Value(), c.W.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T, TW]) ZWZY() (T, TW, T, T)  { return c.Z.Value(), c.W.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T, TW]) ZWZZ() (T, TW, T, T)  { return c.Z.Value(), c.W.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T, TW]) ZWZW() (T, TW, T, TW) { return c.Z.Value(), c.W.Value(), c.Z.Value(), c.W.Value() }
func (c XYZW[T, TW]) ZWWX() (T, TW, TW, T) { return c.Z.Value(), c.W.Value(), c.W.Value(), c.X.Value() }
func (c XYZW[T, TW]) ZWWY() (T, TW, TW, T) { return c.Z.Value(), c.W.Value(), c.W.Value(), c.Y.Value() }
func (c XYZW[T, TW]) ZWWZ() (T, TW, TW, T) { return c.Z.Value(), c.W.Value(), c.W.Value(), c.Z.Value() }
func (c XYZW[T, TW]) ZWWW() (T, TW, TW, TW) {
	return c.Z.Value(), c.W.Value(), c.W.Value(), c.W.Value()
}
func (c XYZW[T, TW]) WXXX() (TW, T, T, T)  { return c.W.Value(), c.X.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T, TW]) WXXY() (TW, T, T, T)  { return c.W.Value(), c.X.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T, TW]) WXXZ() (TW, T, T, T)  { return c.W.Value(), c.X.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T, TW]) WXXW() (TW, T, T, TW) { return c.W.Value(), c.X.Value(), c.X.Value(), c.W.Value() }
func (c XYZW[T, TW]) WXYX() (TW, T, T, T)  { return c.W.Value(), c.X.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T, TW]) WXYY() (TW, T, T, T)  { return c.W.Value(), c.X.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T, TW]) WXYZ() (TW, T, T, T)  { return c.W.Value(), c.X.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T, TW]) WXYW() (TW, T, T, TW) { return c.W.Value(), c.X.Value(), c.Y.Value(), c.W.Value() }
func (c XYZW[T, TW]) WXZX() (TW, T, T, T)  { return c.W.Value(), c.X.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T, TW]) WXZY() (TW, T, T, T)  { return c.W.Value(), c.X.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T, TW]) WXZZ() (TW, T, T, T)  { return c.W.Value(), c.X.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T, TW]) WXZW() (TW, T, T, TW) { return c.W.Value(), c.X.Value(), c.Z.Value(), c.W.Value() }
func (c XYZW[T, TW]) WXWX() (TW, T, TW, T) { return c.W.Value(), c.X.Value(), c.W.Value(), c.X.Value() }
func (c XYZW[T, TW]) WXWY() (TW, T, TW, T) { return c.W.Value(), c.X.Value(), c.W.Value(), c.Y.Value() }
func (c XYZW[T, TW]) WXWZ() (TW, T, TW, T) { return c.W.Value(), c.X.Value(), c.W.Value(), c.Z.Value() }
func (c XYZW[T, TW]) WXWW() (TW, T, TW, TW) {
	return c.W.Value(), c.X.Value(), c.W.Value(), c.W.Value()
}
func (c XYZW[T, TW]) WYXX() (TW, T, T, T)  { return c.W.Value(), c.Y.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T, TW]) WYXY() (TW, T, T, T)  { return c.W.Value(), c.Y.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T, TW]) WYXZ() (TW, T, T, T)  { return c.W.Value(), c.Y.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T, TW]) WYXW() (TW, T, T, TW) { return c.W.Value(), c.Y.Value(), c.X.Value(), c.W.Value() }
func (c XYZW[T, TW]) WYYX() (TW, T, T, T)  { return c.W.Value(), c.Y.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T, TW]) WYYY() (TW, T, T, T)  { return c.W.Value(), c.Y.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T, TW]) WYYZ() (TW, T, T, T)  { return c.W.Value(), c.Y.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T, TW]) WYYW() (TW, T, T, TW) { return c.W.Value(), c.Y.Value(), c.Y.Value(), c.W.Value() }
func (c XYZW[T, TW]) WYZX() (TW, T, T, T)  { return c.W.Value(), c.Y.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T, TW]) WYZY() (TW, T, T, T)  { return c.W.Value(), c.Y.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T, TW]) WYZZ() (TW, T, T, T)  { return c.W.Value(), c.Y.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T, TW]) WYZW() (TW, T, T, TW) { return c.W.Value(), c.Y.Value(), c.Z.Value(), c.W.Value() }
func (c XYZW[T, TW]) WYWX() (TW, T, TW, T) { return c.W.Value(), c.Y.Value(), c.W.Value(), c.X.Value() }
func (c XYZW[T, TW]) WYWY() (TW, T, TW, T) { return c.W.Value(), c.Y.Value(), c.W.Value(), c.Y.Value() }
func (c XYZW[T, TW]) WYWZ() (TW, T, TW, T) { return c.W.Value(), c.Y.Value(), c.W.Value(), c.Z.Value() }
func (c XYZW[T, TW]) WYWW() (TW, T, TW, TW) {
	return c.W.Value(), c.Y.Value(), c.W.Value(), c.W.Value()
}
func (c XYZW[T, TW]) WZXX() (TW, T, T, T)  { return c.W.Value(), c.Z.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T, TW]) WZXY() (TW, T, T, T)  { return c.W.Value(), c.Z.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T, TW]) WZXZ() (TW, T, T, T)  { return c.W.Value(), c.Z.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T, TW]) WZXW() (TW, T, T, TW) { return c.W.Value(), c.Z.Value(), c.X.Value(), c.W.Value() }
func (c XYZW[T, TW]) WZYX() (TW, T, T, T)  { return c.W.Value(), c.Z.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T, TW]) WZYY() (TW, T, T, T)  { return c.W.Value(), c.Z.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T, TW]) WZYZ() (TW, T, T, T)  { return c.W.Value(), c.Z.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T, TW]) WZYW() (TW, T, T, TW) { return c.W.Value(), c.Z.Value(), c.Y.Value(), c.W.Value() }
func (c XYZW[T, TW]) WZZX() (TW, T, T, T)  { return c.W.Value(), c.Z.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T, TW]) WZZY() (TW, T, T, T)  { return c.W.Value(), c.Z.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T, TW]) WZZZ() (TW, T, T, T)  { return c.W.Value(), c.Z.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T, TW]) WZZW() (TW, T, T, TW) { return c.W.Value(), c.Z.Value(), c.Z.Value(), c.W.Value() }
func (c XYZW[T, TW]) WZWX() (TW, T, TW, T) { return c.W.Value(), c.Z.Value(), c.W.Value(), c.X.Value() }
func (c XYZW[T, TW]) WZWY() (TW, T, TW, T) { return c.W.Value(), c.Z.Value(), c.W.Value(), c.Y.Value() }
func (c XYZW[T, TW]) WZWZ() (TW, T, TW, T) { return c.W.Value(), c.Z.Value(), c.W.Value(), c.Z.Value() }
func (c XYZW[T, TW]) WZWW() (TW, T, TW, TW) {
	return c.W.Value(), c.Z.Value(), c.W.Value(), c.W.Value()
}
func (c XYZW[T, TW]) WWXX() (TW, TW, T, T) { return c.W.Value(), c.W.Value(), c.X.Value(), c.X.Value() }
func (c XYZW[T, TW]) WWXY() (TW, TW, T, T) { return c.W.Value(), c.W.Value(), c.X.Value(), c.Y.Value() }
func (c XYZW[T, TW]) WWXZ() (TW, TW, T, T) { return c.W.Value(), c.W.Value(), c.X.Value(), c.Z.Value() }
func (c XYZW[T, TW]) WWXW() (TW, TW, T, TW) {
	return c.W.Value(), c.W.Value(), c.X.Value(), c.W.Value()
}
func (c XYZW[T, TW]) WWYX() (TW, TW, T, T) { return c.W.Value(), c.W.Value(), c.Y.Value(), c.X.Value() }
func (c XYZW[T, TW]) WWYY() (TW, TW, T, T) { return c.W.Value(), c.W.Value(), c.Y.Value(), c.Y.Value() }
func (c XYZW[T, TW]) WWYZ() (TW, TW, T, T) { return c.W.Value(), c.W.Value(), c.Y.Value(), c.Z.Value() }
func (c XYZW[T, TW]) WWYW() (TW, TW, T, TW) {
	return c.W.Value(), c.W.Value(), c.Y.Value(), c.W.Value()
}
func (c XYZW[T, TW]) WWZX() (TW, TW, T, T) { return c.W.Value(), c.W.Value(), c.Z.Value(), c.X.Value() }
func (c XYZW[T, TW]) WWZY() (TW, TW, T, T) { return c.W.Value(), c.W.Value(), c.Z.Value(), c.Y.Value() }
func (c XYZW[T, TW]) WWZZ() (TW, TW, T, T) { return c.W.Value(), c.W.Value(), c.Z.Value(), c.Z.Value() }
func (c XYZW[T, TW]) WWZW() (TW, TW, T, TW) {
	return c.W.Value(), c.W.Value(), c.Z.Value(), c.W.Value()
}
func (c XYZW[T, TW]) WWWX() (TW, TW, TW, T) {
	return c.W.Value(), c.W.Value(), c.W.Value(), c.X.Value()
}
func (c XYZW[T, TW]) WWWY() (TW, TW, TW, T) {
	return c.W.Value(), c.W.Value(), c.W.Value(), c.Y.Value()
}
func (c XYZW[T, TW]) WWWZ() (TW, TW, TW, T) {
	return c.W.Value(), c.W.Value(), c.W.Value(), c.Z.Value()
}
func (c XYZW[T, TW]) WWWW() (TW, TW, TW, TW) {
	return c.W.Value(), c.W.Value(), c.W.Value(), c.W.Value()
}
