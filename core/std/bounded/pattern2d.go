package bounded

import "github.com/ignite-laboratories/core/std"

// Pattern2D represents a 2D "Matrix" of pattern data.
//
// The first three dimensions typically are spatial.
//
//	Typical Dimensional Layout
//	    X | Spatial Width
//	    Y | Spatial Height
//
// NOTE: For pattern generation and predefined patterns, see the 'std/bounded/pattern' package.
//
// See Axis, Emit, and Movement
type Pattern2D[T any] struct {
	Data [][]T
	X    Axis[T]
	Y    Axis[[]T]
}

func NewPattern2D[T any](x Emit[T], y Emit[[]T], data ...[]T) Pattern2D[T] {
	return Pattern2D[T]{
		Data: data,
		X: Axis[T]{
			Emit:   x,
			Cursor: std.NewCursorDefault(),
		},
		Y: Axis[[]T]{
			Emit:   y,
			Cursor: std.NewCursorDefault(),
		},
	}
}
