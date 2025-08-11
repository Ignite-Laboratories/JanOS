// Package cardinal provides access to the cardinal.Direction enumeration.
package cardinal

// Direction represents map-oriented spatial directions.
//
// Abstractly, the result of calculation (the target) is always relatively "down" (or "towards the enemy gate") no matter YOUR orientation
// in space.  Mentally this may be the direction of "gravity" while standing up and writing calculations on a whiteboard, but I think Ender
// described it best.  All binary data is oriented with the most-significant side towards the "left" (or "west").  When operating against a
// matrix, the abstract orientation of "rows" are aligned with the cardinal direction of "south".
//
// Abstract references consider your relative orientation as you float through the void of time and spatial calculation.
//
// See Direction, North, West, South, and East
type Direction byte

const (
	// North represents the cardinal Direction "up" - which is the direction of accumulation.
	//
	// See Direction, North, West, South, and East
	North Direction = iota

	// East represents the cardinal Direction "right" - which is the direction of reduction.
	//
	// See Direction, North, West, South, and East
	East

	// South represents the cardinal Direction "down" - which is the target of all calculation
	//
	// See Direction, North, West, South, and East
	South

	// West represents the cardinal Direction "left" - which is the direction of scale.
	//
	// See Direction, North, West, South, and East
	West
)

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
	default:
		if lower {
			return "unknown"
		}
		return "Unknown"
	}
}
