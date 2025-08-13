package std

// Axis represents a pairing of a std.Cursor and a way to Emit functional Movement.
//
// See Axis, Emit, Movement, and Pattern
type Axis[T any] struct {
	Emit   Emit[T]
	Cursor *Cursor[uint]
}
