// Package awareness provides access to the awareness.Direction enumeration.
package awareness

// Direction represents the 5th dimensional axis of contextual awareness.  A higher-order dimension still can only traverse forwards
// or backwards; however, to reference which logical direction you wish to traverse, you still require a unique term.
//
// For awareness, the ordinal directions are 'Nascent' for negatively, 'Naive' for current, and 'Mature' for positively.  This is
// because we should always remain naive to discovery at the present moment.
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
// See direction.Any, Nascent, Naive, and Mature
type Direction interface {
	Nascent | Naive | Mature
}

// Nascent represents the ordinal Axis of "i-1"
//
// See direction.Any, Nascent, Naive, and Mature
type Nascent byte

func (_ Nascent) String() string {
	return "←"
}

func (_ Nascent) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "nascent"
	}
	return "Nascent"
}

// Naive represents the ordinal Axis of "i"
//
// See direction.Any, Nascent, Naive, and Mature
type Naive byte

func (_ Naive) String() string {
	return "X"
}

func (_ Naive) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "naive"
	}
	return "Naive"
}

// Mature represents the ordinal Axis of "i+1"
//
// See direction.Any, Nascent, Naive, and Mature
type Mature byte

func (_ Mature) String() string {
	return "→"
}

func (_ Mature) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "mature"
	}
	return "Mature"
}
