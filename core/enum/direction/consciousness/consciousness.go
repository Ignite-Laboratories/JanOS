// Package consciousness provides access to the Direction.Axis enumeration.
package consciousness

// Direction represents the 6th dimensional axis of conscious perspective.  A higher-order dimension still can only traverse forwards
// or backwards; however, to reference which logical direction you wish to traverse, you still require a unique term.
//
// For consciousness, the ordinal directions are 'Ignorant' for negatively, 'Emergent' for current, and 'Aware' for positively.  This is
// because we should always be emergently aware of other perspectives in the present moment.
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
// See direction.Any, Ignorant, Emergent, and Aware
type Direction interface {
	Ignorant | Emergent | Aware
}

// Ignorant represents the ordinal Axis of "i-1"
//
// See direction.Any, Ignorant, Emergent, and Aware
type Ignorant byte

func (_ Ignorant) String() string {
	return "←"
}

func (_ Ignorant) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "ignorant"
	}
	return "Ignorant"
}

// Emergent represents the ordinal Axis of "i"
//
// See direction.Any, Ignorant, Emergent, and Aware
type Emergent byte

func (_ Emergent) String() string {
	return "X"
}

func (_ Emergent) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "emergent"
	}
	return "Emergent"
}

// Aware represents the ordinal Axis of "i+1"
//
// See direction.Any, Ignorant, Emergent, and Aware
type Aware byte

func (_ Aware) String() string {
	return "→"
}

func (_ Aware) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "aware"
	}
	return "Aware"
}
