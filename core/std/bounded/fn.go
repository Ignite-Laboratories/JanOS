package bounded

// MovementSingleFn functions should be given the data from which to move through, a target number, and return a single element.
//
// See Axis, Emit, and Movement
type MovementSingleFn[T any] func(n int) T

// MovementManyFn functions should be given the data from which to move through, a target number, and return many elements.
//
// See Axis, Emit, and Movement
type MovementManyFn[T any] func(n int) []T

// MovementFn represents any generic kind of movement function.
//
// See Axis, Emit, and Movement
type MovementFn[T any] interface {
	MovementSingleFn[T] | MovementManyFn[T]
}
