package pattern6D

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/internal"
	"github.com/ignite-laboratories/core/std/movement"
	"github.com/ignite-laboratories/core/std/pattern"
)

// UpTo6D represents patterns up to 6 dimensions wide.
//
// NOTE: This is an alias for pattern.UpTo6D
type UpTo6D[T any] = pattern.UpTo6D[T]

// Nil returns a pattern which always yields nil.
func Nil[T any]() std.Pattern6D[*T] {
	return Default[*T]()
}

// NilAny returns a pattern which always yields a nil interface value (type any).
func NilAny() std.Pattern6D[any] {
	seven := internal.New7DNilAnyPattern()
	return std.Pattern6D[any]{
		X: seven.X,
		Y: seven.Y,
		Z: seven.Z,
		W: seven.W,
		A: seven.A,
		B: seven.B,
	}
}

// Default returns a pattern which always yields the default value of T (type *T).
func Default[T any]() std.Pattern6D[T] {
	seven := internal.New7DZeroPattern[T]()
	return std.Pattern6D[T]{
		X: seven.X,
		Y: seven.Y,
		Z: seven.Z,
		W: seven.W,
		A: seven.A,
		B: seven.B,
	}
}

// From creates a new std.Pattern which can infinitely walk through the provided data along an axis of travel.
//
// NOTE: This will create a pattern holding a single element 'default' instance of T if provided no data.
func From[T any](data ...[][][][][]T) std.Pattern6D[T] {
	if len(data) == 0 {
		var zero [][][][][]T
		data = append(data, zero)
	}

	x, y, z, w, a, b := movement.Function6D[T](data...)
	return std.Pattern6D[T]{
		X: std.Axis[T]{
			Emit:   x,
			Cursor: std.NewCursorDefault(),
		},
		Y: std.Axis[[]T]{
			Emit:   y,
			Cursor: std.NewCursorDefault(),
		},
		Z: std.Axis[[][]T]{
			Emit:   z,
			Cursor: std.NewCursorDefault(),
		},
		W: std.Axis[[][][]T]{
			Emit:   w,
			Cursor: std.NewCursorDefault(),
		},
		A: std.Axis[[][][][]T]{
			Emit:   a,
			Cursor: std.NewCursorDefault(),
		},
		B: std.Axis[[][][][][]T]{
			Emit:   b,
			Cursor: std.NewCursorDefault(),
		},
	}
}
