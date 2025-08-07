// Package bounded provides access to creating num.Bounded values.
package bounded

import (
	"github.com/ignite-laboratories/core/std/num"
)

// By creates a num.Bounded[T] bound within the closed set [min, max], with min and max
// being evaluated from the provided values of a and b.
//
// NOTE: If clamp is not provided, the value will automatically overflow or underflow when
// it exceeds the bounds, otherwise it 'pins' to that boundary point.
func By[T num.ExtendedPrimitive](value, a, b T, clamp ...bool) num.Bounded[T] {
	bnd := num.Bounded[T]{
		Clamp: len(clamp) > 0 && clamp[0],
	}
	bnd.SetAll(value, a, b)
	return bnd
}

// ByType creates a num.Bounded[T] bound within the closed set [T.Min, T.Max], with min and max
// being evaluated from the provided values of a and b.
//
// NOTE: If clamp is not provided, the value will automatically overflow or underflow when
// it exceeds the bounds, otherwise it 'pins' to that boundary point.
//
// NOTE: This supports the num.ExtendedPrimitive implicitly sized types.
func ByType[T num.ExtendedPrimitive](value T, clamp ...bool) num.Bounded[T] {
	return By(value, T(num.MinValue[T]()), T(num.MaxValue[T]()), clamp...)
}
