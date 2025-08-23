// Package ordinal provides access to the ordinal.Direction enumeration.
package ordinal

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
type Direction interface {
	Negative | Current | Positive
}

// Negative represents the ordinal Direction of "i-1"
//
// See direction.Any, Direction, Negative, Current, and Positive
type Negative byte

func (_ Negative) String() string {
	return "←"
}

func (_ Negative) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "negative"
	}
	return "Negative"
}

// Current represents the ordinal Direction of "i"
//
// See direction.Any, Direction, Negative, Current, and Positive
type Current byte

func (_ Current) String() string {
	return "X"
}

func (_ Current) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "current"
	}
	return "Current"
}

// Positive represents the ordinal Direction of "i+1"
//
// See direction.Any, Direction, Negative, Current, and Positive
type Positive byte

func (_ Positive) String() string {
	return "→"
}

func (_ Positive) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "positive"
	}
	return "Positive"
}
