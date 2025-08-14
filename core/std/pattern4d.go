package std

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
// See Axis, Emit, Movement, Pattern, Pattern2D, Pattern3D, Pattern4D, Pattern5D, Pattern6D, and Pattern7D
type Pattern4D[T any] struct {
	X Axis[T]
	Y Axis[[]T]
	Z Axis[[][]T]
	W Axis[[][][]T]
}
