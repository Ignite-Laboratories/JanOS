// Package pad provides access to the padding Side enumeration.
package pad

// Scheme represents how to apply the padding operation to a Side of the operands.
//
// See Truncating, Tiling, Repeating, and Matching
type Scheme interface {
	Tiling | Repeating | Matching | Truncating
}

// Truncating indicates the right operand will be truncated to match the width of the left, or else padded.
//
// See Tiling, Repeating, Matching, and Scheme
type Truncating byte

// Tiling indicates to walk every bit of the right operand in the provided direction of travel.  This means the
// operand's bits will be read left→to→right while being emitted in the direction of travel.
//
// See Truncating, Repeating, Matching, and Scheme
type Tiling byte

// Repeating indicates to repeatedly place a copy of the right operand in the provided direction of travel.  This means
// the operand's bits will be grouped in left→to→right significant order regardless of the direction of travel.
//
// See Truncating, Tiling, Matching, and Scheme
type Repeating byte

// Matching indicates that the smaller operand will grow to match the size of the larger operand.
//
// See Truncating, Tiling, Repeating, and Scheme
type Matching byte
