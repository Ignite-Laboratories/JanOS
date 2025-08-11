package std

// Pattern represents an infinitely repeating slice of elements which can be walked westward or eastward.
//
// Walking eastward will yield and then step, while walking westward will step and then yield.  This is to
// ensure that walking eastward will start at index 0 while walking westward will start at index 𝑛-1.
type Pattern[T any] struct {
	// WalkWest walks one position westward and then yields that position.
	WalkWest PatternFn[T]

	// WalkEast yields the current index position and then walks one position eastward.
	WalkEast PatternFn[T]
}

// PatternFn represents a function that walks through a sequential slice of elements.
type PatternFn[T any] func() T
