package std

import (
	"github.com/ignite-laboratories/core/std/num"
)

// XYZ is a kind of Vector3D that provides X Y Z mappings to the underlying component vectors.
type XYZ[T num.Primitive] = XYZTyped[T, T, T]

// XYZTyped is a kind of VectorTyped3D that provides X Y Z mappings to the underlying component vectors.
type XYZTyped[TX num.Primitive, TY num.Primitive, TZ num.Primitive] VectorTyped3D[TX, TY, TZ]

func (v XYZTyped[TX, TY, TZ]) SetClamp(clamp bool) XYZTyped[TX, TY, TZ] {
	return v.SetClamp(clamp)
}

func (v XYZTyped[TX, TY, TZ]) SetBoundaries(minX, maxX TX, minY, maxY TY, minZ, maxZ TZ) XYZTyped[TX, TY, TZ] {
	return v.SetBoundaries(minX, maxX, minY, maxY, minZ, maxZ)
}

func (v XYZTyped[TX, TY, TZ]) Set(x TX, y TY, z TZ) {
	_ = v.components.x.Set(x)
	_ = v.components.y.Set(y)
	_ = v.components.z.Set(z)
}

func (v XYZTyped[TX, TY, TZ]) X() TX {
	return v.components.x.Value()
}

func (v XYZTyped[TX, TY, TZ]) SetX(value TX) {
	_ = v.components.x.Set(value)
}

func (v XYZTyped[TX, TY, TZ]) Y() TY {
	return v.components.y.Value()
}

func (v XYZTyped[TX, TY, TZ]) SetY(value TY) {
	_ = v.components.y.Set(value)
}

func (v XYZTyped[TX, TY, TZ]) Z() TZ {
	return v.components.z.Value()
}

func (v XYZTyped[TX, TY, TZ]) SetZ(value TZ) {
	_ = v.components.z.Set(value)
}

/**
Swizzling

NOTE: This is a regular expression to find and replace swizzle functions into a one-liner if the auto formatter ever kicks in

Find -
func \((.*?)\) ([A-Z]{2,4})\(\) \((.*?)\)[ ]*\{[\n\t ]*return(.*?)[\n\t ]*\}

Replace -
func ($1) $2() ($3) { return$4 }
*/

func (c XYZTyped[TX, TY, TZ]) XX() (TX, TX) { return c.X(), c.X() }
func (c XYZTyped[TX, TY, TZ]) XY() (TX, TY) { return c.X(), c.Y() }
func (c XYZTyped[TX, TY, TZ]) XZ() (TX, TZ) { return c.X(), c.Z() }
func (c XYZTyped[TX, TY, TZ]) YX() (TY, TX) { return c.Y(), c.X() }
func (c XYZTyped[TX, TY, TZ]) YY() (TY, TY) { return c.Y(), c.Y() }
func (c XYZTyped[TX, TY, TZ]) YZ() (TY, TZ) { return c.Y(), c.Z() }
func (c XYZTyped[TX, TY, TZ]) ZX() (TZ, TX) { return c.Z(), c.X() }
func (c XYZTyped[TX, TY, TZ]) ZY() (TZ, TY) { return c.Z(), c.Y() }
func (c XYZTyped[TX, TY, TZ]) ZZ() (TZ, TZ) { return c.Z(), c.Z() }

func (c XYZTyped[TX, TY, TZ]) XXX() (TX, TX, TX) { return c.X(), c.X(), c.X() }
func (c XYZTyped[TX, TY, TZ]) XXY() (TX, TX, TY) { return c.X(), c.X(), c.Y() }
func (c XYZTyped[TX, TY, TZ]) XXZ() (TX, TX, TZ) { return c.X(), c.X(), c.Z() }
func (c XYZTyped[TX, TY, TZ]) XYX() (TX, TY, TX) { return c.X(), c.Y(), c.X() }
func (c XYZTyped[TX, TY, TZ]) XYY() (TX, TY, TY) { return c.X(), c.Y(), c.Y() }
func (c XYZTyped[TX, TY, TZ]) XYZ() (TX, TY, TZ) { return c.X(), c.Y(), c.Z() }
func (c XYZTyped[TX, TY, TZ]) XZX() (TX, TZ, TX) { return c.X(), c.Z(), c.X() }
func (c XYZTyped[TX, TY, TZ]) XZY() (TX, TZ, TY) { return c.X(), c.Z(), c.Y() }
func (c XYZTyped[TX, TY, TZ]) XZZ() (TX, TZ, TZ) { return c.X(), c.Z(), c.Z() }
func (c XYZTyped[TX, TY, TZ]) YXX() (TY, TX, TX) { return c.Y(), c.X(), c.X() }
func (c XYZTyped[TX, TY, TZ]) YXY() (TY, TX, TY) { return c.Y(), c.X(), c.Y() }
func (c XYZTyped[TX, TY, TZ]) YXZ() (TY, TX, TZ) { return c.Y(), c.X(), c.Z() }
func (c XYZTyped[TX, TY, TZ]) YYX() (TY, TY, TX) { return c.Y(), c.Y(), c.X() }
func (c XYZTyped[TX, TY, TZ]) YYY() (TY, TY, TY) { return c.Y(), c.Y(), c.Y() }
func (c XYZTyped[TX, TY, TZ]) YYZ() (TY, TY, TZ) { return c.Y(), c.Y(), c.Z() }
func (c XYZTyped[TX, TY, TZ]) YZX() (TY, TZ, TX) { return c.Y(), c.Z(), c.X() }
func (c XYZTyped[TX, TY, TZ]) YZY() (TY, TZ, TY) { return c.Y(), c.Z(), c.Y() }
func (c XYZTyped[TX, TY, TZ]) YZZ() (TY, TZ, TZ) { return c.Y(), c.Z(), c.Z() }
func (c XYZTyped[TX, TY, TZ]) ZXX() (TZ, TX, TX) { return c.Z(), c.X(), c.X() }
func (c XYZTyped[TX, TY, TZ]) ZXY() (TZ, TX, TY) { return c.Z(), c.X(), c.Y() }
func (c XYZTyped[TX, TY, TZ]) ZXZ() (TZ, TX, TZ) { return c.Z(), c.X(), c.Z() }
func (c XYZTyped[TX, TY, TZ]) ZYX() (TZ, TY, TX) { return c.Z(), c.Y(), c.X() }
func (c XYZTyped[TX, TY, TZ]) ZYY() (TZ, TY, TY) { return c.Z(), c.Y(), c.Y() }
func (c XYZTyped[TX, TY, TZ]) ZYZ() (TZ, TY, TZ) { return c.Z(), c.Y(), c.Z() }
func (c XYZTyped[TX, TY, TZ]) ZZX() (TZ, TZ, TX) { return c.Z(), c.Z(), c.X() }
func (c XYZTyped[TX, TY, TZ]) ZZY() (TZ, TZ, TY) { return c.Z(), c.Z(), c.Y() }
func (c XYZTyped[TX, TY, TZ]) ZZZ() (TZ, TZ, TZ) { return c.Z(), c.Z(), c.Z() }
