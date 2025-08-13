// Package reality provides access to the reality.Direction enumeration.
package reality

import "github.com/ignite-laboratories/core/std/num"

// Direction represents the 7th dimensional axis of collaborative perspectives which form stable universes.  A higher-order dimension
// still can only traverse forwards or backwards; however, to reference which logical direction you wish to traverse, you
// still require a unique term.
//
// For reality, the ordinal directions are 'Chaos' for negatively, 'Coherence' for current, and 'Stability' for positively.  This is
// because we should always seek coherence in the present moment above stability.
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
// See direction.Any, Axis, Chaos, Coherence, and Stability
type Direction[T num.Primitive] interface {
	Chaos[T] | Coherence[T] | Stability[T]
}

// Axis represents the axis of traversal across the dimension of reality, as such it does not include the present moment of Coherence.
//
// See direction.Any, Axis, Chaos, Coherence, and Stability
type Axis[T num.Primitive] interface {
	Chaos[T] | Stability[T]
}

// Chaos represents the ordinal Axis of "i-1"
//
// See direction.Any, Axis, Chaos, Coherence, and Stability
type Chaos[T num.Primitive] num.Numeric[T]

func (_ Chaos[T]) String() string {
	return "←"
}

func (_ Chaos[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "chaos"
	}
	return "Chaos"
}

// Coherence represents the ordinal Axis of "i"
//
// See direction.Any, Axis, Chaos, Coherence, and Stability
type Coherence[T num.Primitive] num.Numeric[T]

func (_ Coherence[T]) String() string {
	return "X"
}

func (_ Coherence[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "coherence"
	}
	return "Coherence"
}

// Stability represents the ordinal Axis of "i+1"
//
// See direction.Any, Axis, Chaos, Coherence, and Stability
type Stability[T num.Primitive] num.Numeric[T]

func (_ Stability[T]) String() string {
	return "→"
}

func (_ Stability[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "stability"
	}
	return "Stability"
}
