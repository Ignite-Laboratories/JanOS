// Package pattern provides access to creating patterns From data.
package pattern

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/internal"
	"github.com/ignite-laboratories/core/std/movement"
	"github.com/ignite-laboratories/core/std/num"
)

// Nil returns a pattern which always yields nil.
func Nil[T any]() std.Pattern[*T] {
	return Default[*T]()
}

// NilAny returns a pattern which always yields a nil interface value (type any).
func NilAny() std.Pattern[any] {
	seven := internal.New7DNilAnyPattern()
	return std.Pattern[any]{
		X: seven.X,
	}
}

// Default returns a pattern which always yields the default value of T (type *T).
func Default[T any]() std.Pattern[T] {
	seven := internal.New7DZeroPattern[T]()
	return std.Pattern[T]{
		X: seven.X,
	}
}

// Zero is a pattern of a numeric 0.
func Zero[T num.Primitive]() std.Pattern[T] {
	return From[T](0)
}

// One is a pattern of a numeric one.
func One[T num.Primitive]() std.Pattern[T] {
	return From[T](1)
}

// ZeroOne is a pattern of a numeric sequence of '01'.
//
// NOTE: Patterns are stored as you would read them (left→to→right) but are evaluated along an axis of travel.
func ZeroOne[T num.Primitive]() std.Pattern[T] {
	return From[T](0, 1)
}

// OneZero is a pattern of a numeric sequence of '10'.
//
// NOTE: Patterns are stored as you would read them (left→to→right) but are evaluated along an axis of travel.
func OneZero[T num.Primitive]() std.Pattern[T] {
	return From[T](1, 0)
}

// From creates a new std.Pattern which can infinitely walk through the provided data along an axis of travel.
//
// NOTE: This will create a pattern holding a single element 'default' instance of T if provided no data.
func From[T any](data ...T) std.Pattern[T] {
	if len(data) == 0 {
		var zero T
		data = append(data, zero)
	}

	x := movement.Function1D[T](data...)
	return std.Pattern[T]{
		X: std.Axis[T]{
			Emit:   x,
			Cursor: std.NewCursorDefault(),
		},
	}
}
