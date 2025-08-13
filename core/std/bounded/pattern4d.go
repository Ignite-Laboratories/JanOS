package bounded

import (
	"github.com/ignite-laboratories/core/std"
)

// Pattern4D represents a 4D "Tesseract" of temporal pattern data.
//
// The fourth dimension typically provides access to the passing of time.
//
//	Typical Dimensional Layout
//	    X | Spatial Width
//	    Y | Spatial Height
//	    Z | Spatial Depth
//	    W | Time
//
// NOTE: For pattern generation and predefined patterns, see the 'std/bounded/pattern' package.
//
// See Axis, Emit, and Movement
type Pattern4D[T any] struct {
	Data [][][][]T
	X    Axis[T]
	Y    Axis[[]T]
	Z    Axis[[][]T]
	W    Axis[[][][]T]
}

func NewPattern4D[T any](x Emit[T], y Emit[[]T], z Emit[[][]T], w Emit[[][][]T], data ...[][][]T) Pattern4D[T] {
	return Pattern4D[T]{
		Data: data,
		X: Axis[T]{
			Emit:   x,
			Cursor: std.NewCursorDefault(),
		},
		Y: Axis[[]T]{
			Emit:   y,
			Cursor: std.NewCursorDefault(),
		},
		Z: Axis[[][]T]{
			Emit:   z,
			Cursor: std.NewCursorDefault(),
		},
		W: Axis[[][][]T]{
			Emit:   w,
			Cursor: std.NewCursorDefault(),
		},
	}
}
