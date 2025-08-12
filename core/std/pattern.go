package std

// Pattern represents an immutable infinitely repeating slice of elements which can be traversed westward or eastward.
//
// "Going" eastward will yield and then step, while "going" westward will step and then yield.  This is to
// ensure that going eastward will start at index 0 while going westward will start at index 𝑛-1.  Alternatively,
// you may "yield" in a direction to inclusively return all the elements found along the way.
//
// NOTE: For advanced pattern generation and predefined patterns, see the 'std/pattern' package.
//
// See PatternBuffer, Pattern, Pattern2D, Pattern3D, Pattern4D, and PatternFn.
type Pattern[T any] struct {
	goWest    PatternFn[T]
	goEast    PatternFn[T]
	goTo      PatternFn[T]
	yieldWest PatternYieldFn[T]
	yieldEast PatternYieldFn[T]
	yieldTo   PatternYieldFn[T]
	data      []T
	Cursor    *Bounded[uint]
}

// Pattern2D represents a 2D "Matrix" of pattern data.
//
// See PatternBuffer, Pattern, Pattern2D, Pattern3D, Pattern4D, and PatternFn.
type Pattern2D[T any] struct {
	X Pattern[T]
	Y Pattern[T]
}

// Pattern3D represents a 3D "Voxel" of pattern data.
//
// See PatternBuffer, Pattern, Pattern2D, Pattern3D, Pattern4D, and PatternFn.
type Pattern3D[T any] struct {
	X Pattern[T]
	Y Pattern[T]
	Z Pattern[T]
}

// Pattern4D represents a 4D "Cadence" of pattern data.
//
// See PatternBuffer, Pattern, Pattern2D, Pattern3D, Pattern4D, and PatternFn.
type Pattern4D[T any] struct {
	X Pattern[T]
	Y Pattern[T]
	Z Pattern[T]
	W Pattern[T]
}

// PatternFn represents a function that goes to a position in a Pattern[T]
//
// See PatternBuffer, Pattern, Pattern2D, Pattern3D, Pattern4D, and PatternFn.
type PatternFn[T any] func(i uint) T

// PatternYieldFn represents a function that yields the elements founds while going to a position in a Pattern[T]
//
// See PatternBuffer, Pattern, Pattern2D, Pattern3D, Pattern4D, and PatternFn.
type PatternYieldFn[T any] func(i uint) []T

// NewPattern creates a new immutable instance of Pattern[T].
//
// NOTE: For advanced pattern generation and predefined patterns, see the 'std/pattern' package.
func NewPattern[T any](cursor *Bounded[uint], goEast, goWest, goTo PatternFn[T], yieldEast, yieldWest, yieldTo PatternYieldFn[T], data ...T) Pattern[T] {
	return Pattern[T]{
		goWest: goWest,
		goEast: goEast,
		goTo:   goTo,
		data:   data,
		Cursor: cursor,
	}
}

// GetData gets the underlying data which drives this pattern.
func (p Pattern[T]) GetData() []T {
	return p.data
}

// GoTo attempts to walk to the provided index and yields the resulting position.
func (p Pattern[T]) GoTo(i uint) T {
	return p.goTo(i)
}

// GoWest walks westward a fixed number of steps and then yields that position.
func (p Pattern[T]) GoWest(i uint) T {
	return p.goWest(i)
}

// GoEast yields the current index position and then walks a fixed number of steps eastward.
func (p Pattern[T]) GoEast(i uint) T {
	return p.goEast(i)
}

// YieldTo attempts to walk to the provided index and yields the elements found along the way - stopping at the boundary.
func (p Pattern[T]) YieldTo(i uint) []T {
	return p.yieldTo(i)
}

// YieldWest walks westward a fixed number of steps and yields the elements found along the way.
func (p Pattern[T]) YieldWest(i uint) []T {
	return p.yieldWest(i)
}

// YieldEast walks eastward a fixed number of steps and yields the elements found along the way.
func (p Pattern[T]) YieldEast(i uint) []T {
	return p.yieldEast(i)
}
