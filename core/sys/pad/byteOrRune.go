package pad

// ByteOrRune represents a delineation of 'byte' or 'rune' when converting strings to character slices.
type ByteOrRune interface {
	byte | rune
}
