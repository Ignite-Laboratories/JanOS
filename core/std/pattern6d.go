package std

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
// See Axis, Emit, Movement, Pattern, Pattern2D, Pattern3D, Pattern4D, Pattern5D, Pattern6D, and Pattern7D
type Pattern6D[T any] struct {
	Data [][][][][][]T
	X    Axis[T]
	Y    Axis[[]T]
	Z    Axis[[][]T]
	W    Axis[[][][]T]
	A    Axis[[][][][]T]
	B    Axis[[][][][][]T]
}
