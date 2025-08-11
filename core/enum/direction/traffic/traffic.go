// Package traffic provides access to the traffic.Direction enumeration.
package traffic

// Direction represents the general flow of interaction between two abstract entities.
//
// Abstractly, the result of calculation (the target) is always relatively "down" (or "towards the enemy gate") no matter YOUR orientation
// in space.  Mentally this may be the direction of "gravity" while standing up and writing calculations on a whiteboard, but I think Ender
// described it best.  All binary data is oriented with the most-significant side towards the "left" (or "west").  When operating against a
// matrix, the abstract orientation of "rows" are aligned with the cardinal direction of "south".
//
// Abstract references consider your relative orientation as you float through the void of time and spatial calculation.
//
// See Direction, Inbound, Outbound, and Bidirectional
type Direction byte

const (
	// Inbound represents the abstract Direction of "receiving" - which is the direction of listening.
	//
	// See Direction, Inbound, Outbound, and Bidirectional
	Inbound Direction = iota

	// Outbound represents the abstract Direction of "transmitting" - which is the direction of talking.
	//
	// See Direction, Inbound, Outbound, and Bidirectional
	Outbound

	// Bidirectional represents the abstract Direction of "communication" - which is the direction of discourse.
	//
	// See Direction, Inbound, Outbound, and Bidirectional
	Bidirectional
)

func (d Direction) String() string {
	switch d {
	case Inbound:
		return "⇤"
	case Outbound:
		return "↦"
	case Bidirectional:
		return "⇹"
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
	case Inbound:
		if lower {
			return "inbound"
		}
		return "Inbound"
	case Outbound:
		if lower {
			return "outbound"
		}
		return "Outbound"
	case Bidirectional:
		if lower {
			return "bidirectional"
		}
		return "Bidirectional"
	default:
		if lower {
			return "unknown"
		}
		return "Unknown"
	}
}
