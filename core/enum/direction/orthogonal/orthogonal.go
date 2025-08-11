// Package orthogonal provides access to the orthogonal.Direction enumeration.
package orthogonal

// Direction represents orthogonal spatial directions.
//
// Abstractly, the result of calculation (the target) is always relatively "down" (or "towards the enemy gate") no matter YOUR orientation
// in space.  Mentally this may be the direction of "gravity" while standing up and writing calculations on a whiteboard, but I think Ender
// described it best.  All binary data is oriented with the most-significant side towards the "left" (or "west").  When operating against a
// matrix, the abstract orientation of "rows" are aligned with the cardinal direction of "south".
//
// Abstract references consider your relative orientation as you float through the void of time and spatial calculation.
//
// See Direction, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type Direction byte

const (
	// In represents the orthogonal Direction "towards the viewport."
	//
	// See Direction, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
	In Direction = iota

	// Out represents the orthogonal Direction "away from the viewport."
	//
	// See Direction, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
	Out

	// Up represents the orthogonal Direction "towards the top of the viewport."
	//
	// See Direction, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
	Up

	// Down represents the orthogonal Direction "towards the bottom of the viewport."
	//
	// See Direction, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
	Down

	// Left represents the orthogonal Direction "towards the left of the viewport."
	//
	// See Direction, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
	Left

	// Right represents the orthogonal Direction "towards the right of the viewport."
	//
	// See Direction, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
	Right
)

func (d Direction) String() string {
	switch d {
	case Up:
		return "↑"
	case Down:
		return "↓"
	case Left:
		return "←"
	case Right:
		return "→"
	case In:
		return "⤓"
	case Out:
		return "↥"
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
	case Up:
		if lower {
			return "up"
		}
		return "Up"
	case Down:
		if lower {
			return "down"
		}
		return "Down"
	case Left:
		if lower {
			return "left"
		}
		return "Left"
	case Right:
		if lower {
			return "right"
		}
		return "Right"
	case In:
		if lower {
			return "in"
		}
		return "In"
	case Out:
		if lower {
			return "out"
		}
		return "Out"
	default:
		if lower {
			return "unknown"
		}
		return "Unknown"
	}
}
