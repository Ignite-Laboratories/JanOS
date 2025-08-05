package xyz

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/normalize"
	"github.com/ignite-laboratories/core/std/num"
)

// From creates a new instance of std.XYZ[T] with the provided values.
func From[T num.ExtendedPrimitive](x, y, z T, xBound, yBound, zBound T) std.XYZ[T] {
	return std.XYZ[T]{}.SetBoundaries(xBound, yBound, zBound).Set(x, y, z)
}

// FromInfinite creates a new instance of std.XYZ[T] with the provided values, setting the boundaries to the result of std.MaxValue[T].
func FromInfinite[T num.ExtendedPrimitive](x, y, z T) std.XYZ[T] {
	return std.XYZ[T]{}.SetBoundaries(T(std.MaxValue[T]()), T(std.MaxValue[T]()), T(std.MaxValue[T]())).Set(x, y, z)
}

// Random returns a pseudo-random std.XYZ[T] of the provided type using math.Random[T].
//
// If requesting a floating point type, the resulting number will be bounded
// in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [0, n] - where n is the maximum value of
// the provided type.
func Random[T num.ExtendedPrimitive](xBound, yBound, zBound T) std.XYZ[T] {
	x := std.RandomBounded[T](0, xBound)
	y := std.RandomBounded[T](0, yBound)
	z := std.RandomBounded[T](0, zBound)
	return std.XYZ[T]{}.SetBoundaries(xBound, yBound, zBound).Set(x, y, z)
}

// Normalize returns an std.XYZ[TOut] ranging from 0.0-1.0.
func Normalize[TIn num.ExtendedPrimitive, TOut num.Float](source std.XYZ[TIn]) std.XYZ[TOut] {
	x := normalize.To[TIn, TOut](source.X.Value(), source.X.Boundary())
	y := normalize.To[TIn, TOut](source.Y.Value(), source.Y.Boundary())
	z := normalize.To[TIn, TOut](source.Z.Value(), source.Z.Boundary())
	return std.XYZ[TOut]{}.SetBoundaries(TOut(source.X.Boundary()), TOut(source.Y.Boundary()), TOut(source.Z.Boundary())).Set(x, y, z)
}

// ReScale returns an std.XYZ[TOut] scaled up to [0, TIn.Max] from an input bounded in the fully closed interval [0.0, 1.0].
func ReScale[TIn num.Float, TOut num.Integer](source std.XYZ[TIn]) std.XYZ[TOut] {
	x := normalize.From[TIn, TOut](source.X.Value(), TOut(source.X.Boundary()))
	y := normalize.From[TIn, TOut](source.Y.Value(), TOut(source.Y.Boundary()))
	z := normalize.From[TIn, TOut](source.Z.Value(), TOut(source.Z.Boundary()))
	return std.XYZ[TOut]{}.SetBoundaries(TOut(source.X.Boundary()), TOut(source.Y.Boundary()), TOut(source.Z.Boundary())).Set(x, y, z)
}

// Comparator returns if the two std.XYZ values are equal in values.
func Comparator[T num.ExtendedPrimitive](a std.XYZ[T], b std.XYZ[T]) bool {
	return a.X.Value() == b.X.Value() && a.Y.Value() == b.Y.Value() && a.Z.Value() == b.Z.Value()
}
