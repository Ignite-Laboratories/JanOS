// Package pad provides access to the padding Side enumeration.
package pad

// Side represents which side of one operand to reference while padding the other.
//
// See West, East, and Mid
type Side byte

const (
	// West will pad the West side of the smaller operands to the width of the largest.
	//
	// See East, Mid, and Side
	West Side = iota

	// East will pad the East side of the smaller operands to the width of the largest.
	//
	// See West, Mid, and Side
	East

	// Mid will equally pad both sides of the smaller operands to the size of the largest, biased towards the West.
	//
	// See West, East, and Side
	Mid
)

// String prints a one-character representation of the Side -
//
//	West: W
//	East: E
//	 Mid: M
func (s Side) String() string {
	switch s {
	case West:
		return "W"
	case East:
		return "E"
	case Mid:
		return "M"
	default:
		return "Unknown"
	}
}

// StringFull prints an uppercase full two-word representation of the Side.
//
// You may optionally pass true for a lowercase representation.
func (s Side) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	switch s {
	case West:
		if lower {
			return "west side"
		}
		return "West Side"
	case East:
		if lower {
			return "east side"
		}
		return "East Side"
	case Mid:
		if lower {
			return "middle"
		}
		return "Middle"
	default:
		if lower {
			return "unknown"
		}
		return "Unknown"
	}
}
