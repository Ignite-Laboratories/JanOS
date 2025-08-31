// Package bounded provides access to creating bounded.Numeric types.
package bounded

import "core/sys/num"

// By creates a std.Numeric[T] bound within the closed set [min, max], with min and max
// being evaluated from the provided values of a and b.
//
// NOTE: If clamp is not provided, the value will automatically overflow or underflow when
// it exceeds the bounds, otherwise it 'pins' to that boundary point.
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func By[T num.Primitive](value, a, b T, clamp ...bool) (Numeric[T], error) {
	return NewNumber[T](value, a, b, clamp...)
}

// ByType creates a std.Numeric[T] bound within the closed set [T.Min, T.Max], with min and max
// being evaluated from the provided values of a and b.
//
// NOTE: If clamp is not provided, the value will automatically overflow or underflow when
// it exceeds the bounds, otherwise it 'pins' to that boundary point.
//
// NOTE: This supports the num.Primitive implicitly sized types.
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func ByType[T num.Primitive](value T, clamp ...bool) (Numeric[T], error) {
	return By(value, num.MinValue[T](), num.MaxValue[T](), clamp...)
}

// Random seeds a random std.Numeric[T] value within the closed set [0, T.max], with T.max representing
// the implied 'maximum' of the type.
//
// NOTE: If clamp is not provided, the value will automatically overflow or underflow when
// it exceeds the bounds, otherwise it 'pins' to that boundary point.
func Random[T num.Primitive](clamp ...bool) Numeric[T] {
	return RandomSubset(0, num.MaxValue[T](), clamp...)
}

// RandomSubset seeds a random std.Numeric[T] value within the closed set [min, max], with min and max
// being evaluated from the provided values of a and b.
//
// NOTE: If clamp is not provided, the value will automatically overflow or underflow when
// it exceeds the bounds, otherwise it 'pins' to that boundary point.
func RandomSubset[T num.Primitive](minimum, maximum T, clamp ...bool) Numeric[T] {
	random := num.RandomWithinRange(minimum, maximum)
	b, _ := By[T](random, minimum, maximum, clamp...)
	return b
}

// ScaleToType normalizes the std.Numeric[TIn] value to a unit vector and then returns a new std.Numeric[TOut] it.
func ScaleToType[TIn num.Primitive, TOut num.Primitive](value Numeric[TIn]) Numeric[TOut] {
	normalizedPos := value.Normalize()
	b, _ := ByType[TOut](0)
	_ = b.SetFromNormalized(normalizedPos)
	return b
}
