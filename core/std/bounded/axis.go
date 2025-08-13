package bounded

import "github.com/ignite-laboratories/core/std"

// Axis represents a pairing of a std.Cursor and a way to Emit functional Movement.
//
// See Axis, Emit, and Movement
type Axis[T any] struct {
	Emit   Emit[T]
	Cursor *std.Cursor[uint]
}
