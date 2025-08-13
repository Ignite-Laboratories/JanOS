// Package ordinal provides access to the ordinal.Direction enumeration.
package ordinal

import "github.com/ignite-laboratories/core/std/num"

// Direction represents the logical directional order of elements.
//
// All dimensions can be distilled down to an infinitely repeating number line which can be traversed in binary directions -
// but, as you layer these dimensions on top of each other, they orthographically align relative to one another.  The terminology
// used to describe this is entirely dependent upon context, and as such I've provided a robust set of general abstract dimensions
// from which to describe this mechanic in code.  It truly does NOT matter which you use, as long as the called method knows
// how to talk in THAT language. =)
//
// Abstractly, the result of calculation (the target) is always relatively "down" (or "towards the enemy gate") no matter YOUR orientation
// in space.  Mentally this may be the direction of "gravity" while standing up and writing calculations on a whiteboard, but I think Ender
// described it best.  All binary data is oriented with the most-significant side towards the "left" (or "west").  When operating against a
// matrix, you walk "latitudinally" between rows along the Y axis and "longitudinally" between columns along the X axis.  Against a voxel,
// you'd walk negatively "in" or positively "out" along the Z axis.
//
// Abstract references consider your relative orientation as you float through the void of time and spatial calculation.
//
// See direction.Any, Direction, Negative, Current, and Positive
type Direction[T num.Primitive] interface {
	Negative[T] | Current[T] | Positive[T]
}

// Axis represents the axis of traversal across an ordered set, as such it does not include the Current element.
//
// See direction.Any, Direction, Negative, Current, and Positive
type Axis[T num.Primitive] interface {
	Negative[T] | Positive[T]
}

// Negative represents the ordinal Direction of "i-1"
//
// See direction.Any, Direction, Negative, Current, and Positive
type Negative[T num.Primitive] num.Numeric[T]

func (_ Negative[T]) String() string {
	return "←"
}

func (_ Negative[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "negative"
	}
	return "Negative"
}

// Current represents the ordinal Direction of "i"
//
// See direction.Any, Direction, Negative, Current, and Positive
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

// Positive represents the ordinal Direction of "i+1"
//
// See direction.Any, Direction, Negative, Current, and Positive
type Positive[T num.Primitive] num.Numeric[T]

func (_ Positive[T]) String() string {
	return "→"
}

func (_ Positive[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "positive"
	}
	return "Positive"
}
