package std

import (
	"github.com/ignite-laboratories/core/std/num"
)

type VectorUpTo7D[T num.Primitive] interface {
	Vector1D[T] | Vector2D[T] | Vector3D[T] | Vector4D[T] | Vector5D[T] | Vector6D[T] | Vector7D[T]
}

type VectorUpTo6D[T num.Primitive] interface {
	Vector1D[T] | Vector2D[T] | Vector3D[T] | Vector4D[T] | Vector5D[T] | Vector6D[T]
}

type VectorUpTo5D[T num.Primitive] interface {
	Vector1D[T] | Vector2D[T] | Vector3D[T] | Vector4D[T] | Vector5D[T]
}

type VectorUpTo4D[T num.Primitive] interface {
	Vector1D[T] | Vector2D[T] | Vector3D[T] | Vector4D[T]
}

type VectorUpTo3D[T num.Primitive] interface {
	Vector1D[T] | Vector2D[T] | Vector3D[T]
}

type VectorUpTo2D[T num.Primitive] interface {
	Vector1D[T] | Vector2D[T]
}
