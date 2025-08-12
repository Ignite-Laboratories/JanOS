// Package shape provides access to the Any enumeration.
package shape

import "github.com/ignite-laboratories/core/std/num"

type Any[T num.Primitive] interface {
	XY[T] | XYZ[T]
}

type XY[T num.Primitive] interface {
	Square[T] | Circle[T] | Ellipse[T] | Star[T]
}

type XYZ[T num.Primitive] interface {
	Cube[T] | Sphere[T] | Ellipsoid[T] | Star[T]
}
