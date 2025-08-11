// Package pad provides access to the padding Side enumeration.
package pad

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/pattern"
)

// Scheme represents how to apply the padding operation to the operands.
type Scheme[T any] interface {
	tiling | repeating | matching | truncating[T]
}

// TruncateToLeft indicates the right operand will be truncated to match the width of the left, or else padded.
//
// NOTE: If no pattern is provided, pattern.Zero[T] will be used.
func TruncateToLeft[T any](padPattern ...std.Pattern[T]) truncating[T] {
	p := pattern.Zero[T]()
	if len(padPattern) > 0 {
		p = padPattern[0]
	}

	return truncating[T]{
		toRight: false,
		pattern: p,
	}
}

// TruncateToRight indicates the left operand will be truncated to match the width of the right, or else padded.
//
// NOTE: If no pattern is provided, pattern.Zero[T] will be used.
func TruncateToRight[T any](padPattern ...std.Pattern[T]) truncating[T] {
	p := pattern.Zero[T]()
	if len(padPattern) > 0 {
		p = padPattern[0]
	}

	return truncating[T]{
		toRight: false,
		pattern: p,
	}
}

type truncating[T any] struct {
	toRight bool
	pattern std.Pattern[T]
}

// tiling indicates to walk every bit of the right operand in the provided direction of travel.  This means the
// operand's bits will be read left→to→right while being emitted in the direction of travel.
//
// See truncating, repeating, matching, and Scheme
type tiling struct {
}

// repeating indicates to repeatedly place a copy of the right operand in the provided direction of travel.  This means
// the operand's bits will be grouped in left→to→right significant order regardless of the direction of travel.
//
// See truncating, tiling, matching, and Scheme
type repeating struct {
}

// matching indicates that the smaller operand will grow to match the size of the larger operand.
//
// See truncating, tiling, repeating, and Scheme
type matching struct {
}
