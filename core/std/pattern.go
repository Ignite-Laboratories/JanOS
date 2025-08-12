package std

// Pattern represents an immutable infinitely repeating slice of elements which can be walked westward or eastward.
//
// Walking eastward will yield and then step, while walking westward will step and then yield.  This is to
// ensure that walking eastward will start at index 0 while walking westward will start at index 𝑛-1.  All patterns
// are cursored using a Bounded[uint64] - while you CAN directly change the boundaries, please note that you are
// walking into uncharted territory!  I'd love to see what you come up with =)
//
// NOTE: For advanced pattern generation and predefined patterns, see the 'std/pattern' package.
//
// See PatternBuffer, Pattern, Pattern2D, Pattern3D, Pattern4D, and PatternFn.
type Pattern[T any] struct {
	walkWest PatternFn[T]
	walkEast PatternFn[T]
	walkTo   PatternFn[T]
	data     []T
	Cursor   *Bounded[uint]
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

// PatternFn represents a function that walks to a position in a Pattern[T]
//
// See PatternBuffer, Pattern, Pattern2D, Pattern3D, Pattern4D, and PatternFn.
type PatternFn[T any] func(i uint) T

// NewPattern creates a new immutable instance of Pattern[T].
//
// NOTE: For advanced pattern generation and predefined patterns, see the 'std/pattern' package.
func NewPattern[T any](cursor *Bounded[uint], walkEast, walkWest, walkTo PatternFn[T], data ...T) Pattern[T] {
	return Pattern[T]{
		walkWest: walkWest,
		walkEast: walkEast,
		walkTo:   walkTo,
		data:     data,
		Cursor:   cursor,
	}
}

// GetData gets the underlying data which drives this pattern.
func (p Pattern[T]) GetData() []T {
	return p.data
}

// WalkTo attempts to walk to the provided index and yields the resulting position.
func (p Pattern[T]) WalkTo(i uint) T {
	return p.walkTo(i)
}

// WalkWest walks westward a fixed number of steps and then yields that position.
func (p Pattern[T]) WalkWest(i uint) T {
	return p.walkWest(i)
}

// WalkEast yields the current index position and then walks a fixed number of steps eastward.
func (p Pattern[T]) WalkEast(i uint) T {
	return p.walkEast(i)
}
