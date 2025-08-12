// Package ordinal provides access to the ordinal.Direction enumeration.
package ordinal

import "github.com/ignite-laboratories/core/std/num"

// Direction represents the logical directional order of elements.
//
// Abstractly, the result of calculation (the target) is always relatively "down" (or "towards the enemy gate") no matter YOUR orientation
// in space.  Mentally this may be the direction of "gravity" while standing up and writing calculations on a whiteboard, but I think Ender
// described it best.  All binary data is oriented with the most-significant side towards the "left" (or "west").  When operating against a
// matrix, the abstract orientation of "rows" are aligned with the cardinal direction of "south".
//
// Abstract references consider your relative orientation as you float through the void of time and spatial calculation.
//
// See Direction, Before, Current, and After
type Direction[T num.Primitive] interface {
	Before[T] | Current[T] | After[T]
}

// Before represents the ordinal Direction of "i-1"
//
// See Direction, Before, Current, and After
type Before[T num.Primitive] num.Numeric[T]

func (_ Before[T]) String() string {
	return "←"
}

func (_ Before[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "before"
	}
	return "Before"
}

// Current represents the ordinal Direction of "i"
//
// See Direction, Before, Current, and After
type Current[T num.Primitive] num.Numeric[T]

func (_ Current[T]) String() string {
	return "X"
}

func (_ Current[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "current"
	}
	return "Current"
}

// After represents the ordinal Direction of "i+1"
//
// See Direction, Before, Current, and After
type After[T num.Primitive] num.Numeric[T]

func (_ After[T]) String() string {
	return "→"
}

func (_ After[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "after"
	}
	return "After"
}
