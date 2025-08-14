package pattern5D

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/internal"
	"github.com/ignite-laboratories/core/std/movement"
	"github.com/ignite-laboratories/core/std/pattern"
)

// UpTo5D represents patterns up to 5 dimensions wide.
//
// NOTE: This is an alias for pattern.UpTo5D
type UpTo5D[T any] = pattern.UpTo5D[T]

// Nil returns a pattern which always yields nil.
func Nil[T any]() std.Pattern5D[*T] {
	return Default[*T]()
}

// NilAny returns a pattern which always yields a nil interface value (type any).
func NilAny() std.Pattern5D[any] {
	seven := internal.New7DNilAnyPattern()
	return std.Pattern5D[any]{
		X: seven.X,
		Y: seven.Y,
		Z: seven.Z,
		W: seven.W,
		A: seven.A,
	}
}

// Default returns a pattern which always yields the default value of T (type *T).
func Default[T any]() std.Pattern5D[T] {
	seven := internal.New7DZeroPattern[T]()
	return std.Pattern5D[T]{
		X: seven.X,
		Y: seven.Y,
		Z: seven.Z,
		W: seven.W,
		A: seven.A,
	}
}

// From creates a new std.Pattern which can infinitely walk through the provided data along an axis of travel.
//
// NOTE: This will create a pattern holding a single element 'default' instance of T if provided no data.
func From[T any](data ...[][][][]T) std.Pattern5D[T] {
	if len(data) == 0 {
		var zero [][][][]T
		data = append(data, zero)
	}

	x, y, z, w, a := movement.Function5D[T](data...)
	return std.Pattern5D[T]{
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
	}
}
