package std

// Pattern represents an immutable infinitely repeating slice of elements which can be walked westward or eastward.
//
// Walking eastward will yield and then step, while walking westward will step and then yield.  This is to
// ensure that walking eastward will start at index 0 while walking westward will start at index 𝑛-1.
//
// NOTE: While you can call NewPattern directly, pattern generation exists within the 'std/pattern' package.
type Pattern[T any] struct {
	walkWest PatternFn[T]
	walkEast PatternFn[T]
	data     []T
}

// PatternFn represents a function that walks to the next position in a Pattern[T]
type PatternFn[T any] func() T

// NewPattern creates a new immutable instance of Pattern[T].
//
// NOTE: For pre-defined pattern generation, please see the 'std/pattern' package.
func NewPattern[T any](walkEast, walkWest PatternFn[T], data ...T) Pattern[T] {
	return Pattern[T]{
		walkWest: walkWest,
		walkEast: walkEast,
		data:     data,
	}
}

// GetData gets the underlying data which drives this pattern.
func (p Pattern[T]) GetData() []T {
	return p.data
}

// WalkWest walks one position westward and then yields that position.
func (p Pattern[T]) WalkWest() T {
	return p.walkWest()
}

// WalkEast yields the current index position and then walks one position eastward.
func (p Pattern[T]) WalkEast() T {
	return p.walkEast()
}
