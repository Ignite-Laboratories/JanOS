package pattern2D

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/internal"
	"github.com/ignite-laboratories/core/std/movement"
	"github.com/ignite-laboratories/core/std/pattern"
)

// UpTo2D represents patterns up to 2 dimensions wide.
//
// NOTE: This is an alias for pattern.UpTo2D
type UpTo2D[T any] = pattern.UpTo2D[T]

// Nil returns a pattern which always yields nil.
func Nil[T any]() std.Pattern2D[*T] {
	return Default[*T]()
}

// NilAny returns a pattern which always yields a nil interface value (type any).
func NilAny() std.Pattern2D[any] {
	seven := internal.New7DNilAnyPattern()
	return std.Pattern2D[any]{
		X: seven.X,
		Y: seven.Y,
	}
}

// Default returns a pattern which always yields the default value of T (type *T).
func Default[T any]() std.Pattern2D[T] {
	seven := internal.New7DZeroPattern[T]()
	return std.Pattern2D[T]{
		X: seven.X,
		Y: seven.Y,
	}
}

// From creates a new std.Pattern which can infinitely walk through the provided data along an axis of travel.
//
// NOTE: This will create a pattern holding a single element 'default' instance of T if provided no data.
func From[T any](data ...[]T) std.Pattern2D[T] {
	if len(data) == 0 {
		var zero []T
		data = append(data, zero)
	}

	x, y := movement.Function2D[T](data...)
	return std.Pattern2D[T]{
		X: std.Axis[T]{
			Emit:   x,
			Cursor: std.NewCursorDefault(),
		},
		Y: std.Axis[[]T]{
			Emit:   y,
			Cursor: std.NewCursorDefault(),
		},
	}
}
