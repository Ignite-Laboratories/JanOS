// Package padding provides access to the padding Scheme enumeration.
package padding

// Scheme represents the scheme of how to apply a padding operation against a collection of elements.
//
// See Scheme, Left, Right, Middle, Truncate, TruncateLeftOperand, TruncateRightOperand, WalkEast, WalkWest, RepeatEast, and RepeatWest
type Scheme interface {
	Left | Right | Middle | Truncate | TruncateLeftOperand | TruncateRightOperand | WalkEast | WalkWest | RepeatEast | RepeatWest
}

// Left indicates that the smaller operand will be padded on its left side to match the width of the larger operand.
//
// See Scheme, Left, Right, Middle, Truncate, TruncateLeftOperand, TruncateRightOperand, WalkEast, WalkWest, RepeatEast, and RepeatWest
type Left struct{}

// Right indicates that the smaller operand will be padded on its right side to match the width of the larger operand.
//
// See Scheme, Left, Right, Middle, Truncate, TruncateLeftOperand, TruncateRightOperand, WalkEast, WalkWest, RepeatEast, and RepeatWest
type Right struct{}

// Middle indicates that the smaller operand will be padded on both sides to match the width of the larger operand, biased smaller on the left in uneven splits.
//
// See Scheme, Left, Right, Middle, Truncate, TruncateLeftOperand, TruncateRightOperand, WalkEast, WalkWest, RepeatEast, and RepeatWest
type Middle struct{}

// Truncate indicates the widest operand will be truncated to match the width of the smallest.
//
// See Scheme, Left, Right, Middle, Truncate, TruncateLeftOperand, TruncateRightOperand, WalkEast, WalkWest, RepeatEast, and RepeatWest
type Truncate struct{}

// TruncateLeftOperand indicates the left operand will be truncated to match the width of the right, or else padded.
//
// See Scheme, Left, Right, Middle, Truncate, TruncateLeftOperand, TruncateRightOperand, WalkEast, WalkWest, RepeatEast, and RepeatWest
type TruncateLeftOperand struct{}

// TruncateRightOperand indicates the right operand will be truncated to match the width of the left, or else padded.
//
// See Scheme, Left, Right, Middle, Truncate, TruncateLeftOperand, TruncateRightOperand, WalkEast, WalkWest, RepeatEast, and RepeatWest
type TruncateRightOperand struct{}

// WalkEast indicates to infinitely walk the individual pattern elements in an easterly direction.
//
// See Scheme, Left, Right, Middle, Truncate, TruncateLeftOperand, TruncateRightOperand, WalkEast, WalkWest, RepeatEast, and RepeatWest
type WalkEast struct{}

// WalkWest indicates to infinitely walk the individual pattern elements in a westerly direction.
//
// See Scheme, Left, Right, Middle, Truncate, TruncateLeftOperand, TruncateRightOperand, WalkEast, WalkWest, RepeatEast, and RepeatWest
type WalkWest struct{}

// RepeatEast indicates to repeatedly emit the entirety of the pattern data in an easterly direction, leaving the pattern's element order intact.
//
// See Scheme, Left, Right, Middle, Truncate, TruncateLeftOperand, TruncateRightOperand, WalkEast, WalkWest, RepeatEast, and RepeatWest
type RepeatEast struct{}

// RepeatWest indicates to repeatedly emit the entirety of the pattern data in a westerly direction, leaving the pattern's element order intact.
//
// See Scheme, Left, Right, Middle, Truncate, TruncateLeftOperand, TruncateRightOperand, WalkEast, WalkWest, RepeatEast, and RepeatWest
type RepeatWest struct{}
