package bounded

import "github.com/ignite-laboratories/core/std"

// Pattern represents an infinitely repeating pattern buffer of elements which can be traversed along its only axis.
//
// The first three dimensions typically are spatial.
//
//	Typical Dimensional Layout
//	    X | Spatial Width
//
// NOTE: For pattern generation and predefined patterns, see the 'std/bounded/pattern' package.
//
// See Axis, Emit, and Movement
type Pattern[T any] struct {
	Data []T
	X    Axis[T]
}

func NewPattern[T any](x Emit[T], data ...T) Pattern[T] {
	return Pattern[T]{
		Data: data,
		X: Axis[T]{
			Emit:   x,
			Cursor: std.NewCursorDefault(),
		},
	}
}
