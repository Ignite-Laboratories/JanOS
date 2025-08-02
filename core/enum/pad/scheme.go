// Package pad provides access to the padding Scheme enumeration.
package pad

// Scheme represents a scheme of how to align operands relative to each other.  Alignment operations are applied using
// the cardinal directions East, West, North, and South.  In addition, if using a pattern, the direction.Direction of
// travel.Travel also matters.
//
// NOTE: You can only longitudinally align variable width binary types - such as slices of bytes or bits, or measurements or phrases.
// These will panic if you attempt to "pad a single byte to the west," for instance, as it's a static-width element.  Latitudinal
// directions can take in static width binary operands, as they simply introduce new operands to the North and South sides of the matrix.
//
// See WestSide, EastSide, NorthSide, SouthSide, and ToMiddle.
type Scheme byte

const (
	// WestSide will pad the West side of the smaller operands to the width of the largest.
	WestSide Scheme = iota

	// EastSide will pad the East side of the smaller operands to the width of the largest.
	EastSide

	// NorthSide will pad the North side of the matrix to the desired height.
	NorthSide

	// SouthSide will pad the South side of the matrix to the desired height.
	SouthSide

	// ToMiddle will equally pad both sides of the smaller operands to the size of the largest, biased towards the West or North.
	ToMiddle
)

// String prints a two-character representation of the Scheme -
//
//	 WestSide: WS
//	 EastSide: ES
//	NorthSide: NS
//	SouthSide: SS
func (s Scheme) String() string {
	switch s {
	case WestSide:
		return "WS"
	case EastSide:
		return "ES"
	case NorthSide:
		return "NS"
	case SouthSide:
		return "SS"
	default:
		return "Unknown"
	}
}

// StringFull prints an uppercase full two-word representation of the Scheme.
//
// You may optionally pass true for a lowercase representation.
func (s Scheme) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	switch s {
	case WestSide:
		if lower {
			return "west side"
		}
		return "West Side"
	case EastSide:
		if lower {
			return "east side"
		}
		return "East Side"
	case NorthSide:
		if lower {
			return "north side"
		}
		return "North Side"
	case SouthSide:
		if lower {
			return "south side"
		}
		return "South Side"
	default:
		if lower {
			return "unknown"
		}
		return "Unknown"
	}
}
