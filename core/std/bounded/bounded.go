// Package bounded provides access to creating num.Bounded values.
package bounded

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/num"
)

// By creates a std.Bounded[T] bound within the closed set [min, max], with min and max
// being evaluated from the provided values of a and b.
//
// NOTE: If clamp is not provided, the value will automatically overflow or underflow when
// it exceeds the bounds, otherwise it 'pins' to that boundary point.
func By[T num.Primitive](value, a, b T, clamp ...bool) std.Bounded[T] {
	return std.NewBounded[T](value, a, b, clamp...)
}

// ByType creates a std.Bounded[T] bound within the closed set [T.Min, T.Max], with min and max
// being evaluated from the provided values of a and b.
//
// NOTE: If clamp is not provided, the value will automatically overflow or underflow when
// it exceeds the bounds, otherwise it 'pins' to that boundary point.
//
// NOTE: This supports the num.Primitive implicitly sized types.
func ByType[T num.Primitive](value T, clamp ...bool) std.Bounded[T] {
	return By(value, T(num.MinValue[T]()), T(num.MaxValue[T]()), clamp...)
}

// Random seeds a random std.Bounded[T] value within the closed set [0, T.max], with T.max representing
// the implied 'maximum' of the type.
//
// NOTE: If clamp is not provided, the value will automatically overflow or underflow when
// it exceeds the bounds, otherwise it 'pins' to that boundary point.
func Random[T num.Primitive](clamp ...bool) std.Bounded[T] {
	return RandomSubset(0, T(num.MaxValue[T]()), clamp...)
}

// RandomSubset seeds a random std.Bounded[T] value within the closed set [min, max], with min and max
// being evaluated from the provided values of a and b.
//
// NOTE: If clamp is not provided, the value will automatically overflow or underflow when
// it exceeds the bounds, otherwise it 'pins' to that boundary point.
func RandomSubset[T num.Primitive](minimum, maximum T, clamp ...bool) std.Bounded[T] {
	random := num.RandomWithinRange(minimum, maximum)
	return By[T](random, minimum, maximum, clamp...)
}

// ScaleToType normalizes the std.Bounded[TIn] value to a unit vector and then returns a new std.Bounded[TOut] it.
func ScaleToType[TIn num.Primitive, TOut num.Primitive](value std.Bounded[TIn]) std.Bounded[TOut] {
	normalizedPos := value.Normalize()
	return ByType[TOut](0).SetFromNormalized(normalizedPos)
}
