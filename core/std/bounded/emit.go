package bounded

// Emit represents a pairing of a MovementSingleFn and MovementManyFn.
//
// See Axis, Emit, and Movement
type Emit[T any] struct {
	Single Movement[T, MovementSingleFn[T]]
	Many   Movement[T, MovementManyFn[T]]
}
