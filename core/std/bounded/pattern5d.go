package bounded

import "github.com/ignite-laboratories/core/std"

// Pattern5D represents a 5D "Awareness" of pattern data.
//
// The fifth dimension typically provides access to spatial context across time.
//
//	Typical Dimensional Layout
//	    X | Spatial Width
//	    Y | Spatial Height
//	    Z | Spatial Depth
//	    W | Time
//	    A | Context
//
// NOTE: For pattern generation and predefined patterns, see the 'std/bounded/pattern' package.
//
// See Axis, Emit, and Movement
type Pattern5D[T any] struct {
	Data [][][][][]T
	X    Axis[T]
	Y    Axis[[]T]
	Z    Axis[[][]T]
	W    Axis[[][][]T]
	A    Axis[[][][][]T]
}

func NewPattern5D[T any](x Emit[T], y Emit[[]T], z Emit[[][]T], w Emit[[][][]T], a Emit[[][][][]T], data ...[][][][]T) Pattern5D[T] {
	return Pattern5D[T]{
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
		A: Axis[[][][][]T]{
			Emit:   a,
			Cursor: std.NewCursorDefault(),
		},
	}
}
