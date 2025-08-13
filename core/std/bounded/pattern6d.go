package bounded

import "github.com/ignite-laboratories/core/std"

// Pattern6D represents a 6D "Consciousness" of pattern data.
//
// The sixth dimension typically provides access to perspectives of contextual spatial events across time.
//
//	Typical Dimensional Layout
//	    X | Spatial Width
//	    Y | Spatial Height
//	    Z | Spatial Depth
//	    W | Time
//	    A | Context
//	    B | Perspective
//
// NOTE: For pattern generation and predefined patterns, see the 'std/bounded/pattern' package.
//
// See Axis, Emit, and Movement
type Pattern6D[T any] struct {
	Data [][][][][][]T
	X    Axis[T]
	Y    Axis[[]T]
	Z    Axis[[][]T]
	W    Axis[[][][]T]
	A    Axis[[][][][]T]
	B    Axis[[][][][][]T]
}

func NewPattern6D[T any](x Emit[T], y Emit[[]T], z Emit[[][]T], w Emit[[][][]T], a Emit[[][][][]T], b Emit[[][][][][]T], data ...[][][][][]T) Pattern6D[T] {
	return Pattern6D[T]{
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
	}
}
