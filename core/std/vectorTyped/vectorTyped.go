package vectorTyped

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/num"
)

// From2D creates a new std.Vector2DTyped[T] with the provided component values and bounded in the fully closed interval [0, max[TComponent]].
func From2D[TX num.Primitive, TY num.Primitive](x TX, y TY, clamp ...bool) std.Vector2DTyped[TX, TY] {
	return From2DBounded(x, y, 0, num.MaxValue[TX](), 0, num.MaxValue[TY](), clamp...)
}

// From2DBounded creates a new std.Vector2DTyped[T] with the provided component values and bounded in the fully closed interval [minComponent, maxComponent].
func From2DBounded[TX num.Primitive, TY num.Primitive](x TX, y TY, minX, maxX TX, minY, maxY TY, clamp ...bool) std.Vector2DTyped[TX, TY] {
	return std.Vector2DTyped[TX, TY]{}.SetClamp(len(clamp) > 0 && clamp[0]).SetBoundaries(minX, maxX, minY, maxY).SetValues(x, y)
}

// From3D creates a new std.Vector3DTyped[T] with the provided component values and bounded in the fully closed interval [0, max[TComponent]].
func From3D[TX num.Primitive, TY num.Primitive, TZ num.Primitive](x TX, y TY, z TZ, clamp ...bool) std.Vector3DTyped[TX, TY, TZ] {
	return From3DBounded(x, y, z, 0, num.MaxValue[TX](), 0, num.MaxValue[TY](), 0, num.MaxValue[TZ](), clamp...)
}

// From3DBounded creates a new std.Vector3DTyped[T] with the provided component values and bounded in the fully closed interval [minComponent, maxComponent].
func From3DBounded[TX num.Primitive, TY num.Primitive, TZ num.Primitive](x TX, y TY, z TZ, minX, maxX TX, minY, maxY TY, minZ, maxZ TZ, clamp ...bool) std.Vector3DTyped[TX, TY, TZ] {
	return std.Vector3DTyped[TX, TY, TZ]{}.SetClamp(len(clamp) > 0 && clamp[0]).SetBoundaries(minX, maxX, minY, maxY, minZ, maxZ).SetValues(x, y, z)
}

// From4D creates a new std.Vector4DTyped[T] with the provided component values and bounded in the fully closed interval [0, max[TComponent]].
func From4D[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive](x TX, y TY, z TZ, w TW, clamp ...bool) std.Vector4DTyped[TX, TY, TZ, TW] {
	return From4DBounded(x, y, z, w, 0, num.MaxValue[TX](), 0, num.MaxValue[TY](), 0, num.MaxValue[TZ](), 0, num.MaxValue[TW](), clamp...)
}

// From4DBounded creates a new std.Vector4DTyped[T] with the provided component values and bounded in the fully closed interval [minComponent, maxComponent].
func From4DBounded[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive](x TX, y TY, z TZ, w TW, minX, maxX TX, minY, maxY TY, minZ, maxZ TZ, minW, maxW TW, clamp ...bool) std.Vector4DTyped[TX, TY, TZ, TW] {
	return std.Vector4DTyped[TX, TY, TZ, TW]{}.SetClamp(len(clamp) > 0 && clamp[0]).SetBoundaries(minX, maxX, minY, maxY, minZ, maxZ, minW, maxW).SetValues(x, y, z, w)
}

// From5D creates a new std.Vector5DTyped[T] with the provided component values and bounded in the fully closed interval [0, max[TComponent]].
func From5D[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive, TA num.Primitive](x TX, y TY, z TZ, w TW, a TA, clamp ...bool) std.Vector5DTyped[TX, TY, TZ, TW, TA] {
	return From5DBounded(x, y, z, w, a, 0, num.MaxValue[TX](), 0, num.MaxValue[TY](), 0, num.MaxValue[TZ](), 0, num.MaxValue[TW](), 0, num.MaxValue[TA](), clamp...)
}

// From5DBounded creates a new std.Vector5DTyped[T] with the provided component values and bounded in the fully closed interval [minComponent, maxComponent].
func From5DBounded[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive, TA num.Primitive](x TX, y TY, z TZ, w TW, a TA, minX, maxX TX, minY, maxY TY, minZ, maxZ TZ, minW, maxW TW, minA, maxA TA, clamp ...bool) std.Vector5DTyped[TX, TY, TZ, TW, TA] {
	return std.Vector5DTyped[TX, TY, TZ, TW, TA]{}.SetClamp(len(clamp) > 0 && clamp[0]).SetBoundaries(minX, maxX, minY, maxY, minZ, maxZ, minW, maxW, minA, maxA).SetValues(x, y, z, w, a)
}

// From6D creates a new std.Vector6DTyped[T] with the provided component values and bounded in the fully closed interval [0, max[TComponent]].
func From6D[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive, TA num.Primitive, TB num.Primitive](x TX, y TY, z TZ, w TW, a TA, b TB, clamp ...bool) std.Vector6DTyped[TX, TY, TZ, TW, TA, TB] {
	return From6DBounded(x, y, z, w, a, b, 0, num.MaxValue[TX](), 0, num.MaxValue[TY](), 0, num.MaxValue[TZ](), 0, num.MaxValue[TW](), 0, num.MaxValue[TA](), 0, num.MaxValue[TB](), clamp...)
}

// From6DBounded creates a new std.Vector6DTyped[T] with the provided component values and bounded in the fully closed interval [minComponent, maxComponent].
func From6DBounded[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive, TA num.Primitive, TB num.Primitive](x TX, y TY, z TZ, w TW, a TA, b TB, minX, maxX TX, minY, maxY TY, minZ, maxZ TZ, minW, maxW TW, minA, maxA TA, minB, maxB TB, clamp ...bool) std.Vector6DTyped[TX, TY, TZ, TW, TA, TB] {
	return std.Vector6DTyped[TX, TY, TZ, TW, TA, TB]{}.SetClamp(len(clamp) > 0 && clamp[0]).SetBoundaries(minX, maxX, minY, maxY, minZ, maxZ, minW, maxW, minA, maxA, minB, maxB).SetValues(x, y, z, w, a, b)
}

// From7D creates a new std.Vector7DTyped[T] with the provided component values and bounded in the fully closed interval [0, max[TComponent]].
func From7D[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive, TA num.Primitive, TB num.Primitive, TC num.Primitive](x TX, y TY, z TZ, w TW, a TA, b TB, c TC, clamp ...bool) std.Vector7DTyped[TX, TY, TZ, TW, TA, TB, TC] {
	return From7DBounded(x, y, z, w, a, b, c, 0, num.MaxValue[TX](), 0, num.MaxValue[TY](), 0, num.MaxValue[TZ](), 0, num.MaxValue[TW](), 0, num.MaxValue[TA](), 0, num.MaxValue[TB](), 0, num.MaxValue[TC](), clamp...)
}

// From7DBounded creates a new std.Vector7DTyped[T] with the provided component values and bounded in the fully closed interval [minComponent, maxComponent].
func From7DBounded[TX num.Primitive, TY num.Primitive, TZ num.Primitive, TW num.Primitive, TA num.Primitive, TB num.Primitive, TC num.Primitive](x TX, y TY, z TZ, w TW, a TA, b TB, c TC, minX, maxX TX, minY, maxY TY, minZ, maxZ TZ, minW, maxW TW, minA, maxA TA, minB, maxB TB, minC, maxC TC, clamp ...bool) std.Vector7DTyped[TX, TY, TZ, TW, TA, TB, TC] {
	return std.Vector7DTyped[TX, TY, TZ, TW, TA, TB, TC]{}.SetClamp(len(clamp) > 0 && clamp[0]).SetBoundaries(minX, maxX, minY, maxY, minZ, maxZ, minW, maxW, minA, maxA, minB, maxB, minC, maxC).SetValues(x, y, z, w, a, b, c)
}
