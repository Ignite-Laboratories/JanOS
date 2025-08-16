// Package cursor provides access to creating std.Cursor values.
package cursor

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/num"
)

// By creates a std.Cursor[T] bound within the closed set [min, max], with min and max
// being evaluated from the provided values of a and b.
//
// NOTE: If clamp is not provided, the value will automatically overflow or underflow when
// it exceeds the bounds, otherwise it 'pins' to that boundary point.
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func By[T num.Primitive](value, a, b T, clamp ...bool) (std.Cursor[T], error) {
	return std.NewCursor[T](value, a, b, clamp...)
}

// ByType creates a std.Cursor[T] bound within the closed set [T.Min, T.Max], with min and max
// being evaluated from the provided values of a and b.
//
// NOTE: If clamp is not provided, the value will automatically overflow or underflow when
// it exceeds the bounds, otherwise it 'pins' to that boundary point.
//
// NOTE: This supports the num.Primitive implicitly sized types.
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func ByType[T num.Primitive](value T, clamp ...bool) (std.Cursor[T], error) {
	return By(value, num.MinValue[T](), num.MaxValue[T](), clamp...)
}

// Random seeds a random std.Cursor[T] value within the closed set [0, T.max], with T.max representing
// the implied 'maximum' of the type.
//
// NOTE: If clamp is not provided, the value will automatically overflow or underflow when
// it exceeds the bounds, otherwise it 'pins' to that boundary point.
func Random[T num.Primitive](clamp ...bool) std.Cursor[T] {
	return RandomSubset(0, num.MaxValue[T](), clamp...)
}

// RandomSubset seeds a random std.Cursor[T] value within the closed set [min, max], with min and max
// being evaluated from the provided values of a and b.
//
// NOTE: If clamp is not provided, the value will automatically overflow or underflow when
// it exceeds the bounds, otherwise it 'pins' to that boundary point.
func RandomSubset[T num.Primitive](minimum, maximum T, clamp ...bool) std.Cursor[T] {
	random := num.RandomWithinRange(minimum, maximum)
	b, _ := By[T](random, minimum, maximum, clamp...)
	return b
}

// ScaleToType normalizes the std.Cursor[TIn] value to a unit vector and then returns a new std.Cursor[TOut] it.
func ScaleToType[TIn num.Primitive, TOut num.Primitive](value std.Cursor[TIn]) std.Cursor[TOut] {
	normalizedPos := value.Normalize()
	b, _ := ByType[TOut](0)
	_ = b.SetFromNormalized(normalizedPos)
	return b
}
