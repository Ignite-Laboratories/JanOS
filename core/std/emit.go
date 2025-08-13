package std

// Emit represents a pairing of a std.MovementSingleFn and std.MovementManyFn.
//
// See Axis, Emit, Movement, and Pattern
type Emit[T any] struct {
	Single Movement[T, MovementSingleFn[T]]
	Many   Movement[T, MovementManyFn[T]]
}

func NewEmit[T any](singleRelative MovementSingleFn[T], manyRelative MovementManyFn[T], singleAbsolute MovementSingleFn[T], manyAbsolute MovementManyFn[T]) Emit[T] {
	return Emit[T]{
		Single: Movement[T, MovementSingleFn[T]]{
			Relatively: singleRelative,
			Absolutely: singleAbsolute,
		},
		Many: Movement[T, MovementManyFn[T]]{
			Relatively: manyRelative,
			Absolutely: manyAbsolute,
		},
	}
}
