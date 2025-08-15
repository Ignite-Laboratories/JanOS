package vector

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/num"
	"github.com/ignite-laboratories/core/std/vectorTyped"
)

// From2D creates a new std.Vector2D[T] with the provided component values and bounded in the fully closed interval [0, max[TComponent]].
func From2D[T num.Primitive](x, y T, clamp ...bool) std.Vector2D[T] {
	return vectorTyped.From2DBounded(x, y, 0, num.MaxValue[T](), 0, num.MaxValue[T](), clamp...)
}

// From2DBounded creates a new std.Vector2D[T] with the provided component values and bounded in the fully closed interval [minComponent, maxComponent].
func From2DBounded[T num.Primitive](x, y, minX, maxX, minY, maxY T, clamp ...bool) std.Vector2D[T] {
	return std.Vector2D[T](vectorTyped.From2DBounded(x, y, minX, maxX, minY, maxY, clamp...))
}

// From3D creates a new std.Vector3D[T] with the provided component values and bounded in the fully closed interval [0, max[TComponent]].
func From3D[T num.Primitive](x, y, z T, clamp ...bool) std.Vector3D[T] {
	return vectorTyped.From3DBounded(x, y, z, 0, num.MaxValue[T](), 0, num.MaxValue[T](), 0, num.MaxValue[T](), clamp...)
}

// From3DBounded creates a new std.Vector3D[T] with the provided component values and bounded in the fully closed interval [minComponent, maxComponent].
func From3DBounded[T num.Primitive](x, y, z, minX, maxX, minY, maxY, minZ, maxZ T, clamp ...bool) std.Vector3D[T] {
	return std.Vector3D[T](vectorTyped.From3DBounded(x, y, z, minX, maxX, minY, maxY, minZ, maxZ, clamp...))
}

// From4D creates a new std.Vector4D[T] with the provided component values and bounded in the fully closed interval [0, max[TComponent]].
func From4D[T num.Primitive](x, y, z, w T, clamp ...bool) std.Vector4D[T] {
	return vectorTyped.From4DBounded(x, y, z, w, 0, num.MaxValue[T](), 0, num.MaxValue[T](), 0, num.MaxValue[T](), 0, num.MaxValue[T](), clamp...)
}

// From4DBounded creates a new std.Vector4D[T] with the provided component values and bounded in the fully closed interval [minComponent, maxComponent].
func From4DBounded[T num.Primitive](x, y, z, w, minX, maxX, minY, maxY, minZ, maxZ, minW, maxW T, clamp ...bool) std.Vector4D[T] {
	return std.Vector4D[T](vectorTyped.From4DBounded(x, y, z, w, minX, maxX, minY, maxY, minZ, maxZ, minW, maxW, clamp...))
}

// From5D creates a new std.Vector5D[T] with the provided component values and bounded in the fully closed interval [0, max[TComponent]].
func From5D[T num.Primitive](x, y, z, w, a T, clamp ...bool) std.Vector5D[T] {
	return vectorTyped.From5DBounded(x, y, z, w, a, 0, num.MaxValue[T](), 0, num.MaxValue[T](), 0, num.MaxValue[T](), 0, num.MaxValue[T](), 0, num.MaxValue[T](), clamp...)
}

// From5DBounded creates a new std.Vector5D[T] with the provided component values and bounded in the fully closed interval [minComponent, maxComponent].
func From5DBounded[T num.Primitive](x, y, z, w, a, minX, maxX, minY, maxY, minZ, maxZ, minW, maxW, minA, maxA T, clamp ...bool) std.Vector5D[T] {
	return std.Vector5D[T](vectorTyped.From5DBounded(x, y, z, w, a, minX, maxX, minY, maxY, minZ, maxZ, minW, maxW, minA, maxA, clamp...))
}

// From6D creates a new std.Vector6D[T] with the provided component values and bounded in the fully closed interval [0, max[TComponent]].
func From6D[T num.Primitive](x, y, z, w, a, b T, clamp ...bool) std.Vector6D[T] {
	return vectorTyped.From6DBounded(x, y, z, w, a, b, 0, num.MaxValue[T](), 0, num.MaxValue[T](), 0, num.MaxValue[T](), 0, num.MaxValue[T](), 0, num.MaxValue[T](), 0, num.MaxValue[T](), clamp...)
}

// From6DBounded creates a new std.Vector6D[T] with the provided component values and bounded in the fully closed interval [minComponent, maxComponent].
func From6DBounded[T num.Primitive](x, y, z, w, a, b, minX, maxX, minY, maxY, minZ, maxZ, minW, maxW, minA, maxA, minB, maxB T, clamp ...bool) std.Vector6D[T] {
	return std.Vector6D[T](vectorTyped.From6DBounded(x, y, z, w, a, b, minX, maxX, minY, maxY, minZ, maxZ, minW, maxW, minA, maxA, minB, maxB, clamp...))
}

// From7D creates a new std.Vector7D[T] with the provided component values and bounded in the fully closed interval [0, max[TComponent]].
func From7D[T num.Primitive](x, y, z, w, a, b, c T, clamp ...bool) std.Vector7D[T] {
	return vectorTyped.From7DBounded(x, y, z, w, a, b, c, 0, num.MaxValue[T](), 0, num.MaxValue[T](), 0, num.MaxValue[T](), 0, num.MaxValue[T](), 0, num.MaxValue[T](), 0, num.MaxValue[T](), 0, num.MaxValue[T](), clamp...)
}

// From7DBounded creates a new std.Vector7D[T] with the provided component values and bounded in the fully closed interval [minComponent, maxComponent].
func From7DBounded[T num.Primitive](x, y, z, w, a, b, c, minX, maxX, minY, maxY, minZ, maxZ, minW, maxW, minA, maxA, minB, maxB, minC, maxC T, clamp ...bool) std.Vector7D[T] {
	return std.Vector7D[T](vectorTyped.From7DBounded(x, y, z, w, a, b, c, minX, maxX, minY, maxY, minZ, maxZ, minW, maxW, minA, maxA, minB, maxB, minC, maxC, clamp...))
}
