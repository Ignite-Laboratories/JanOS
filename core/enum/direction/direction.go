// noinspection GoStructureTemplates
// Package direction provides access to the Direction enumeration.
package direction

// Direction represents general directionality and includes both cardinal and abstract reference points in time and space.
//
// Abstractly, the result of calculation (the target) is always relatively "down" (or "towards the enemies gate") no matter YOUR orientation
// in space.  Mentally this may be the direction of "gravity" while standing up and writing calculations on a whiteboard, but I think Ender
// described it best.  All binary data is oriented with the most-significant side towards the "left" (or "west").  When operating against a
// matrix, the abstract orientation of "down" is aligned with the cardinal direction of "south".
//
// Abstract references consider your relative orientation as you float through the void of time and spatial calculation.
//
// The cardinal directions each have an explicitly defined purpose relative to a fixed point of orientation (the result of calculation) -
//
//	  South - Calculation
//	  North - Accumulation
//	   West - Scale
//	   East - Reduction
//	 Future - Anticipation
//	Present - Experience
//	   Past - Reflection
//
// In addition, you may also require traveling.Traveling in a particular orientation.
//
// See South, West, North, East, Future, Present, Past, Forward, Backward, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
type Direction byte

const (
	// South represents the cardinal Direction "down" - which is the -target- of all calculation.
	//
	// See West, North, and East.
	//
	// See also Future, Present, Past, Forward, Backward, Up, Down, Left, and Right.
	South Direction = iota

	// West represents the cardinal Direction "left" - which is the direction of scale.
	//
	// See South, North, and East.
	//
	// See also Future, Present, Past, Forward, Backward, Up, Down, Left, and Right.
	West

	// North represents the cardinal Direction "up" - which is the direction of accumulation.
	//
	// See South, West, and East.
	//
	// See also Future, Present, Past, Forward, Backward, Up, Down, Left, and Right.
	North

	// East represents the cardinal Direction "right" - which is the direction of reduction.
	//
	// See South, West, and North.
	//
	// See also Future, Present, Past, Forward, Backward, Up, Down, Left, and Right.
	East

	// Future represents the abstract temporal Direction of "eminently" - which is the direction of anticipation.
	//
	// See Present and Past.
	//
	// See also South, West, North, East, Forward, Backward, Up, Down, Left, and Right
	Future

	// Present represents the abstract temporal Direction of "currently" - which is the direction of experience.
	//
	// See Future and Past.
	//
	// See also South, West, North, East, Forward, Backward, Up, Down, Left, and Right
	Present

	// Past represents the abstract temporal Direction of "historically" - which is the direction of reflection.
	//
	// See Future and Present.
	//
	// See also South, West, North, East, Forward, Backward, Up, Down, Left, and Right
	Past

	// Up represents the abstract Direction of presently relative "up."
	//
	// See Down, Left, Right, Forward, and Backward.
	//
	// See also South, West, North, East, Future, Present, and Past.
	Up // Up Down Down Left Right A B Start

	// Down represents the abstract Direction of presently relative "down."
	//
	// See Up, Left, Right, Forward, and Backward.
	//
	// See also South, West, North, East, Future, Present, and Past.
	Down // Down Left Right A B Start

	// Left represents the abstract Direction of presently relative "left."
	//
	// See Up, Down, Right, Forward, and Backward.
	//
	// See also South, West, North, East, Future, Present, and Past.
	Left

	// Right represents the abstract Direction of presently relative "right."
	//
	// See Up, Down, Left, Forward, and Backward.
	//
	// See also South, West, North, East, Future, Present, and Past.
	Right // A B Start

	// Forward represents the abstract Direction of presently relative "forward."
	//
	// See Up, Down, Left, Right, and Backward.
	//
	// See also South, West, North, East, Future, Present, and Past.
	Forward

	// Backward represents the abstract Direction of presently relative "backward."
	//
	// See Up, Down, Left, Right, and Forward.
	//
	// See also South, West, North, East, Future, Present, and Past.
	Backward
)

// String prints a single-character representation of the Direction -
//
//	   South: S
//	    West: W
//	   North: N
//	    East: E
//
//	  Future: ⏭
//	 Present: ⏸
//	    Past: ⏮
//
//	      Up: ↑
//	    Down: ↓
//	    Left: ←
//	   Right: →
//	 Forward: ↷
//	Backward: ↶
func (d Direction) String() string {
	switch d {
	case South:
		return "S"
	case West:
		return "W"
	case North:
		return "N"
	case East:
		return "E"
	case Future:
		return "⏭"
	case Present:
		return "⏸"
	case Past:
		return "⏮"
	case Up:
		return "↑"
	case Down:
		return "↓"
	case Left:
		return "←"
	case Right:
		return "→"
	case Forward:
		return "↷"
	case Backward:
		return "↶"
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
	case South:
		if lower {
			return "south"
		}
		return "South"
	case West:
		if lower {
			return "west"
		}
		return "West"
	case North:
		if lower {
			return "north"
		}
		return "North"
	case East:
		if lower {
			return "east"
		}
		return "East"
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
	case Forward:
		if lower {
			return "forward"
		}
		return "Forward"
	case Backward:
		if lower {
			return "backward"
		}
		return "Backward"
	default:
		if lower {
			return "unknown"
		}
		return "Unknown"
	}
}
