package bounded

import "github.com/ignite-laboratories/core/std"

// Pattern7D represents a 7D "Reality" of pattern data.
//
// The seventh dimension provides access to a multiverse of conscious awareness across space and time.
//
//	Typical Dimensional Layout
//	    X | Spatial Width
//	    Y | Spatial Height
//	    Z | Spatial Depth
//	    W | Time
//	    A | Context
//	    B | Perspective
//	    C | Reality
//
// NOTE: For pattern generation and predefined patterns, see the 'std/bounded/pattern' package.
//
// See Axis, Emit, and Movement
type Pattern7D[T any] struct {
	Data [][][][][][][]T
	X    Axis[T]
	Y    Axis[[]T]
	Z    Axis[[][]T]
	W    Axis[[][][]T]
	A    Axis[[][][][]T]
	B    Axis[[][][][][]T]
	C    Axis[[][][][][][]T]
}

func NewPattern7D[T any](x Emit[T], y Emit[[]T], z Emit[[][]T], w Emit[[][][]T], a Emit[[][][][]T], b Emit[[][][][][]T], c Emit[[][][][][][]T], data ...[][][][][][]T) Pattern7D[T] {
	return Pattern7D[T]{
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
		B: Axis[[][][][][]T]{
			Emit:   b,
			Cursor: std.NewCursorDefault(),
		},
		C: Axis[[][][][][][]T]{
			Emit:   c,
			Cursor: std.NewCursorDefault(),
		},
	}
}
