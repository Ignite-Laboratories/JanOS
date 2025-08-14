package internal

import "github.com/ignite-laboratories/core/std"

/**
Nil Any
*/

func New7DNilAnyPattern() std.Pattern7D[any] {
	single1D := func(int) any { return nil }
	many1D := func(int) []any { return []any{nil} }
	emit1D := newEmitSimple[any](single1D, many1D)

	single2D := func(int) []any { return nil }
	many2D := func(int) [][]any { return [][]any{nil} }
	emit2D := newEmitSimple[[]any](single2D, many2D)

	single3D := func(int) [][]any { return nil }
	many3D := func(int) [][][]any { return [][][]any{nil} }
	emit3D := newEmitSimple[[][]any](single3D, many3D)

	single4D := func(int) [][][]any { return nil }
	many4D := func(int) [][][][]any { return [][][][]any{nil} }
	emit4D := newEmitSimple[[][][]any](single4D, many4D)

	single5D := func(int) [][][][]any { return nil }
	many5D := func(int) [][][][][]any { return [][][][][]any{nil} }
	emit5D := newEmitSimple[[][][][]any](single5D, many5D)

	single6D := func(int) [][][][][]any { return nil }
	many6D := func(int) [][][][][][]any { return [][][][][][]any{nil} }
	emit6D := newEmitSimple[[][][][][]any](single6D, many6D)

	single7D := func(int) [][][][][][]any { return nil }
	many7D := func(int) [][][][][][][]any { return [][][][][][][]any{nil} }
	emit7D := newEmitSimple[[][][][][][]any](single7D, many7D)

	return newPattern7d[any](emit1D, emit2D, emit3D, emit4D, emit5D, emit6D, emit7D)
}

/**
Zero
*/

func New7DZeroPattern[T any]() std.Pattern7D[T] {
	var zero1D T
	single1D := func(int) T { return zero1D }
	many1D := func(int) []T { return []T{zero1D} }
	emit1D := newEmitSimple[T](single1D, many1D)

	var zero2D []T
	single2D := func(int) []T { return zero2D }
	many2D := func(int) [][]T { return [][]T{zero2D} }
	emit2D := newEmitSimple[[]T](single2D, many2D)

	var zero3D [][]T
	single3D := func(int) [][]T { return zero3D }
	many3D := func(int) [][][]T { return [][][]T{zero3D} }
	emit3D := newEmitSimple[[][]T](single3D, many3D)

	var zero4D [][][]T
	single4D := func(int) [][][]T { return zero4D }
	many4D := func(int) [][][][]T { return [][][][]T{zero4D} }
	emit4D := newEmitSimple[[][][]T](single4D, many4D)

	var zero5D [][][][]T
	single5D := func(int) [][][][]T { return zero5D }
	many5D := func(int) [][][][][]T { return [][][][][]T{zero5D} }
	emit5D := newEmitSimple[[][][][]T](single5D, many5D)

	var zero6D [][][][][]T
	single6D := func(int) [][][][][]T { return zero6D }
	many6D := func(int) [][][][][][]T { return [][][][][][]T{zero6D} }
	emit6D := newEmitSimple[[][][][][]T](single6D, many6D)

	var zero7D [][][][][][]T
	single7D := func(int) [][][][][][]T { return zero7D }
	many7D := func(int) [][][][][][][]T { return [][][][][][][]T{zero7D} }
	emit7D := newEmitSimple[[][][][][][]T](single7D, many7D)

	return newPattern7d[T](emit1D, emit2D, emit3D, emit4D, emit5D, emit6D, emit7D)
}

/**
Helpers
*/

// newEmitSimple is the same as NewEmit but shares the logic between Relatively and Absolutely for both single and many.
func newEmitSimple[T any](single std.MovementSingleFn[T], many std.MovementManyFn[T]) std.Emit[T] {
	return std.Emit[T]{
		Single: std.Movement[T, std.MovementSingleFn[T]]{
			Relatively: single,
			Absolutely: single,
		},
		Many: std.Movement[T, std.MovementManyFn[T]]{
			Relatively: many,
			Absolutely: many,
		},
	}
}

func newPattern7d[T any](x std.Emit[T], y std.Emit[[]T], z std.Emit[[][]T], w std.Emit[[][][]T], a std.Emit[[][][][]T], b std.Emit[[][][][][]T], c std.Emit[[][][][][][]T]) std.Pattern7D[T] {
	return std.Pattern7D[T]{
		X: std.Axis[T]{
			Emit:   x,
			Cursor: std.NewCursorDefault(),
		},
		Y: std.Axis[[]T]{
			Emit:   y,
			Cursor: std.NewCursorDefault(),
		},
		Z: std.Axis[[][]T]{
			Emit:   z,
			Cursor: std.NewCursorDefault(),
		},
		W: std.Axis[[][][]T]{
			Emit:   w,
			Cursor: std.NewCursorDefault(),
		},
		A: std.Axis[[][][][]T]{
			Emit:   a,
			Cursor: std.NewCursorDefault(),
		},
		B: std.Axis[[][][][][]T]{
			Emit:   b,
			Cursor: std.NewCursorDefault(),
		},
		C: std.Axis[[][][][][][]T]{
			Emit:   c,
			Cursor: std.NewCursorDefault(),
		},
	}
}
