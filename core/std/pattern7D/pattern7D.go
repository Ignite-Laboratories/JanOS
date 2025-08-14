package pattern7D

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/internal"
	"github.com/ignite-laboratories/core/std/movement"
	"github.com/ignite-laboratories/core/std/pattern"
)

// UpTo7D represents patterns up to 7 dimensions wide.
//
// NOTE: This is an alias for pattern.UpTo7D
type UpTo7D[T any] = pattern.UpTo7D[T]

// Nil returns a pattern which always yields nil (type *T).
func Nil[T any]() std.Pattern7D[*T] {
	return Default[*T]()
}

// NilAny returns a pattern which always yields a nil interface value (type any).
func NilAny() std.Pattern7D[any] {
	seven := internal.New7DNilAnyPattern()
	return std.Pattern7D[any]{
		X: seven.X,
		Y: seven.Y,
		Z: seven.Z,
		W: seven.W,
		A: seven.A,
		B: seven.B,
		C: seven.C,
	}
}

// Default returns a pattern which always yields the default value of T (type *T).
func Default[T any]() std.Pattern7D[T] {
	seven := internal.New7DZeroPattern[T]()
	return std.Pattern7D[T]{
		X: seven.X,
		Y: seven.Y,
		Z: seven.Z,
		W: seven.W,
		A: seven.A,
		B: seven.B,
		C: seven.C,
	}
}

// From creates a new std.Pattern which can infinitely walk through the provided data along an axis of travel.
//
// NOTE: This will create a pattern holding a single element 'default' instance of T if provided no data.
func From[T any](data ...[][][][][][]T) std.Pattern7D[T] {
	if len(data) == 0 {
		var zero [][][][][][]T
		data = append(data, zero)
	}

	x, y, z, w, a, b, c := movement.Function7D[T](data...)
	return std.Pattern7D[T]{
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
		C: std.Axis[[][][][][][]T]{
			Emit:   c,
			Cursor: std.NewCursorDefault(),
		},
	}
}
