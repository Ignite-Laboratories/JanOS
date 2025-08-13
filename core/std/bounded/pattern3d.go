package bounded

import "github.com/ignite-laboratories/core/std"

// Pattern3D represents a 3D "Voxel" of pattern data.
//
// The first three dimensions typically are spatial.
//
//	Typical Dimensional Layout
//	    X | Spatial Width
//	    Y | Spatial Height
//	    Z | Spatial Depth
//
// NOTE: For pattern generation and predefined patterns, see the 'std/bounded/pattern' package.
//
// See Axis, Emit, and Movement
type Pattern3D[T any] struct {
	Data [][][]T
	X    Axis[T]
	Y    Axis[[]T]
	Z    Axis[[][]T]
}

func NewPattern3D[T any](x Emit[T], y Emit[[]T], z Emit[[][]T], data ...[][]T) Pattern3D[T] {
	return Pattern3D[T]{
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
	}
}
