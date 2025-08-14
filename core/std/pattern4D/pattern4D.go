package pattern4D

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/internal"
	"github.com/ignite-laboratories/core/std/movement"
	"github.com/ignite-laboratories/core/std/pattern"
)

// UpTo4D represents patterns up to 4 dimensions wide.
//
// NOTE: This is an alias for pattern.UpTo4D
type UpTo4D[T any] = pattern.UpTo4D[T]

// Nil returns a pattern which always yields nil.
func Nil[T any]() std.Pattern4D[*T] {
	return Default[*T]()
}

// NilAny returns a pattern which always yields a nil interface value (type any).
func NilAny() std.Pattern4D[any] {
	seven := internal.New7DNilAnyPattern()
	return std.Pattern4D[any]{
		X: seven.X,
		Y: seven.Y,
		Z: seven.Z,
		W: seven.W,
	}
}

// Default returns a pattern which always yields the default value of T (type *T).
func Default[T any]() std.Pattern4D[T] {
	seven := internal.New7DZeroPattern[T]()
	return std.Pattern4D[T]{
		X: seven.X,
		Y: seven.Y,
		Z: seven.Z,
		W: seven.W,
	}
}

// From creates a new std.Pattern which can infinitely walk through the provided data along an axis of travel.
//
// NOTE: This will create a pattern holding a single element 'default' instance of T if provided no data.
func From[T any](data ...[][][]T) std.Pattern4D[T] {
	if len(data) == 0 {
		var zero [][][]T
		data = append(data, zero)
	}

	x, y, z, w := movement.Function4D[T](data...)
	return std.Pattern4D[T]{
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
	}
}
