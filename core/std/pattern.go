package std

// Pattern represents an infinitely repeating pattern buffer of elements which can be traversed along its only axis.
//
// The first three dimensions typically are spatial.
//
//	Typical Dimensional Layout
//	    X | Spatial Width
//
// NOTE: For pattern generation and predefined patterns, see the 'std/bounded/pattern' package.
//
// See Axis, Emit, Movement, Pattern, Pattern2D, Pattern3D, Pattern4D, Pattern5D, Pattern6D, and Pattern7D
type Pattern[T any] struct {
	Data []T
	X    Axis[T]
}
