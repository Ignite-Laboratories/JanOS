package std

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
// See Axis, Emit, Movement, Pattern, Pattern2D, Pattern3D, Pattern4D, Pattern5D, Pattern6D, and Pattern7D
type Pattern2D[T any] struct {
	Data [][]T
	X    Axis[T]
	Y    Axis[[]T]
}
