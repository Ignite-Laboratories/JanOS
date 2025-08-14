package std

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
// See Axis, Emit, Movement, Pattern, Pattern2D, Pattern3D, Pattern4D, Pattern5D, Pattern6D, and Pattern7D
type Pattern5D[T any] struct {
	X Axis[T]
	Y Axis[[]T]
	Z Axis[[][]T]
	W Axis[[][][]T]
	A Axis[[][][][]T]
}
