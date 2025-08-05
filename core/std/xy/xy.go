package xy

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/normalize"
	"github.com/ignite-laboratories/core/std/num"
)

// From creates a new instance of std.XY[T] with the provided values.
func From[T num.ExtendedPrimitive](x, y T, xBound, yBound T) std.XY[T] {
	return std.XY[T]{}.SetBoundaries(xBound, yBound).Set(x, y)
}

// FromInfinite creates a new instance of std.XY[T] with the provided values, setting the boundaries to the result of std.MaxValue[T].
func FromInfinite[T num.ExtendedPrimitive](x, y T) std.XY[T] {
	return std.XY[T]{}.SetBoundaries(T(std.MaxValue[T]()), T(std.MaxValue[T]())).Set(x, y)
}

// Random returns a pseudo-random std.XY[T] of the provided type using math.Random[T].
//
// If requesting a floating point type, the resulting number will be bounded
// in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [0, n] - where n is the maximum value of
// the provided type.
func Random[T num.ExtendedPrimitive](xBound, yBound T) std.XY[T] {
	x := std.RandomBounded[T](0, xBound)
	y := std.RandomBounded[T](0, yBound)
	return std.XY[T]{}.SetBoundaries(xBound, yBound).Set(x, y)
}

// Normalize returns an std.XY[TOut] ranging from 0.0-1.0.
func Normalize[TIn num.ExtendedPrimitive, TOut num.Float](source std.XY[TIn]) std.XY[TOut] {
	x := normalize.To[TIn, TOut](source.X.Value(), source.X.Boundary())
	y := normalize.To[TIn, TOut](source.Y.Value(), source.Y.Boundary())
	return std.XY[TOut]{}.SetBoundaries(TOut(source.X.Boundary()), TOut(source.Y.Boundary())).Set(x, y)
}

// ReScale returns an std.XY[TOut] scaled up to [0, TIn.Max] from an input bounded in the fully closed interval [0.0, 1.0].
func ReScale[TIn num.Float, TOut num.Integer](source std.XY[TIn]) std.XY[TOut] {
	x := normalize.From[TIn, TOut](source.X.Value(), TOut(source.X.Boundary()))
	y := normalize.From[TIn, TOut](source.Y.Value(), TOut(source.Y.Boundary()))
	return std.XY[TOut]{}.SetBoundaries(TOut(source.X.Boundary()), TOut(source.Y.Boundary())).Set(x, y)
}

// Comparator returns if the two std.XY values are equal in values.
func Comparator[T num.ExtendedPrimitive](a std.XY[T], b std.XY[T]) bool {
	return a.X.Value() == b.X.Value() && a.Y.Value() == b.Y.Value()
}
