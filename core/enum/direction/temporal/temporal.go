// Package temporal provides access to the temporal.Direction enumeration.
package temporal

// Direction represents the direction of temporal activity.
//
// Abstractly, the result of calculation (the target) is always relatively "down" (or "towards the enemy gate") no matter YOUR orientation
// in space.  Mentally this may be the direction of "gravity" while standing up and writing calculations on a whiteboard, but I think Ender
// described it best.  All binary data is oriented with the most-significant side towards the "left" (or "west").  When operating against a
// matrix, the abstract orientation of "rows" are aligned with the cardinal direction of "south".
//
// Abstract references consider your relative orientation as you float through the void of time and spatial calculation.
//
// See Direction, Past, Present, and Future
type Direction byte

const (
	// Past represents the abstract Direction of "historically" - which is the direction of reflection.
	//
	// See Direction, Past, Present, and Future
	Past Direction = iota

	// Present represents the abstract Direction of "currently" - which is the direction of experience.
	//
	// See Direction, Past, Present, and Future
	Present

	// Future represents the abstract Direction of "eminently" - which is the direction of anticipation.
	//
	// See Direction, Past, Present, and Future
	Future
)

func (d Direction) String() string {
	switch d {
	case Future:
		return "⏭"
	case Present:
		return "⏸"
	case Past:
		return "⏮"
	default:
		return "Unknown"
	}
}

// StringFull prints an uppercase full word representation of the Direction.
//
// You may optionally pass true for a lowercase representation.
func (d Direction) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	switch d {
	case Future:
		if lower {
			return "future"
		}
		return "Future"
	case Present:
		if lower {
			return "present"
		}
		return "Present"
	case Past:
		if lower {
			return "past"
		}
		return "Past"
	default:
		if lower {
			return "unknown"
		}
		return "Unknown"
	}
}
