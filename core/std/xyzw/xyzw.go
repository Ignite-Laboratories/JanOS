package xyzw

//
//import (
//	"github.com/ignite-laboratories/core/std"
//	"github.com/ignite-laboratories/core/std/normalize"
//	"github.com/ignite-laboratories/core/std/num"
//)
//
//// From creates a new instance of std.XYZW[T] with the provided values.
//func From[T num.ExtendedPrimitive](x, y, z, w T, xBound, yBound, zBound, wBound T) std.XYZW[T] {
//	return std.XYZW[T]{}.SetBoundaries(xBound, yBound, zBound, wBound).Set(x, y, z, w)
//}
//
//// FromInfinite creates a new instance of std.XYZW[T] with the provided values, setting the boundaries to the result of std.MaxValue[T].
//func FromInfinite[T num.ExtendedPrimitive](x, y, z, w T) std.XYZW[T] {
//	return std.XYZW[T]{}.SetBoundaries(T(num.MaxValue[T]()), T(num.MaxValue[T]()), T(num.MaxValue[T]()), T(num.MaxValue[T]())).Set(x, y, z, w)
//}
//
//// Random returns a pseudo-random std.XYZW[T] of the provided type using math.Random[T].
////
//// If requesting a floating point type, the resulting number will be bounded
//// in the fully closed interval [0.0, 1.0]
////
//// If requesting an integer type, the resulting number will be bounded
//// in the fully closed interval [0, n] - where n is the maximum value of
//// the provided type.
//func Random[T num.ExtendedPrimitive](xBound, yBound, zBound, wBound T) std.XYZW[T] {
//	x := num.RandomBounded[T](0, xBound)
//	y := num.RandomBounded[T](0, yBound)
//	z := num.RandomBounded[T](0, zBound)
//	w := num.RandomBounded[T](0, wBound)
//	return std.XYZW[T]{}.SetBoundaries(xBound, yBound, zBound, wBound).Set(x, y, z, w)
//}
//
//// Normalize returns an std.XYZW[TOut] ranging from 0.0-1.0.
//func Normalize[TIn num.ExtendedPrimitive, TOut num.Float](source std.XYZW[TIn]) std.XYZW[TOut] {
//	x := normalize.To[TIn, TOut](source.X.Value(), source.X.Boundary())
//	y := normalize.To[TIn, TOut](source.Y.Value(), source.Y.Boundary())
//	z := normalize.To[TIn, TOut](source.Z.Value(), source.Z.Boundary())
//	w := normalize.To[TIn, TOut](source.W.Value(), source.W.Boundary())
//	return std.XYZW[TOut]{}.SetBoundaries(TOut(source.X.Boundary()), TOut(source.Y.Boundary()), TOut(source.Z.Boundary()), TOut(source.W.Boundary())).Set(x, y, z, w)
//}
//
//// ReScale returns an std.XYZW[TOut] scaled up to [0, TIn.Max] from an input bounded in the fully closed interval [0.0, 1.0].
//func ReScale[TIn num.Float, TOut num.Integer](source std.XYZW[TIn]) std.XYZW[TOut] {
//	x := normalize.From[TIn, TOut](source.X.Value(), TOut(source.X.Boundary()))
//	y := normalize.From[TIn, TOut](source.Y.Value(), TOut(source.Y.Boundary()))
//	z := normalize.From[TIn, TOut](source.Z.Value(), TOut(source.Z.Boundary()))
//	w := normalize.From[TIn, TOut](source.W.Value(), TOut(source.W.Boundary()))
//	return std.XYZW[TOut]{}.SetBoundaries(TOut(source.X.Boundary()), TOut(source.Y.Boundary()), TOut(source.Z.Boundary()), TOut(source.W.Boundary())).Set(x, y, z, w)
//}
//
//// Comparator returns if the two std.XYZW values are equal in values.
//func Comparator[T num.ExtendedPrimitive](a std.XYZW[T], b std.XYZW[T]) bool {
//	return a.X.Value() == b.X.Value() && a.Y.Value() == b.Y.Value() && a.Z.Value() == b.Z.Value() && a.W.Value() == b.W.Value()
//}
