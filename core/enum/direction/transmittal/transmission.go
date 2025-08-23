// Package transmittal provides access to the transmittal.Direction enumeration.
package transmittal

// Direction represents the general transmission of information between two abstract entities.
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
// See direction.Any, Direction, Inbound, Outbound, and Bidirectional
type Direction interface {
	Inbound | Outbound | Bidirectional
}

// Inbound represents the abstract Direction of "receiving" - which is the direction of listening.
//
// See direction.Any, Direction, Inbound, Outbound, and Bidirectional
type Inbound byte

func (_ Inbound) String() string {
	return "⇤"
}

func (_ Inbound) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "inbound"
	}
	return "Inbound"
}

// Outbound represents the abstract Direction of "transmitting" - which is the direction of talking.
//
// See direction.Any, Direction, Inbound, Outbound, and Bidirectional
type Outbound byte

func (_ Outbound) String() string {
	return "↦"
}

func (_ Outbound) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "outbound"
	}
	return "Outbound"
}

// Bidirectional represents the abstract Direction of "communication" - which is the direction of discourse.
//
// See direction.Any, Direction, Inbound, Outbound, and Bidirectional
type Bidirectional byte

func (_ Bidirectional) String() string {
	return "⇹"
}

func (_ Bidirectional) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "bidirectional"
	}
	return "Bidirectional"
}
