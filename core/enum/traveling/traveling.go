// Package traveling provides access to the Traveling enumeration.
package traveling

// Traveling represents a longitudinal or latitudinal direction.Direction of "travel".
//
// These directly relate to the cardinal directions of calculation.
//
// See Westbound, Eastbound, Northbound, Southbound, Outbound, and Inbound.
type Traveling byte

const (
	// Westbound represents a westerly direction of travel.
	//
	// See Eastbound, Northbound, Southbound, Outbound, and Inbound.
	Westbound Traveling = iota

	// Eastbound represents an easterly direction of travel.
	//
	// See Westbound, Northbound, Southbound, Outbound, and Inbound.
	Eastbound

	// Northbound represents a northerly direction of travel.
	//
	// See Westbound, Eastbound, Southbound, Outbound, and Inbound.
	Northbound

	// Southbound represents a southerly direction of travel.
	//
	// See Westbound, Eastbound, Northbound, Outbound, and Inbound.
	Southbound

	// Outbound represents an outward direction of travel.
	//
	// See Westbound, Eastbound, Northbound, Southbound, and Inbound.
	Outbound

	// Inbound represents an inward direction of travel.
	//
	// See Westbound, Eastbound, Northbound, Southbound, and Outbound.
	Inbound
)

// String prints a two (or three) character representation of the Traveling direction -
//
//	 Westbound: WB
//	Northbound: NB
//	 Eastbound: EB
//	Southbound: SB
//	   Outward: OUT
//	    Inward: IN
func (t Traveling) String() string {
	switch t {
	case Westbound:
		return "WB"
	case Northbound:
		return "NB"
	case Eastbound:
		return "EB"
	case Southbound:
		return "SB"
	case Outbound:
		return "OUT"
	case Inbound:
		return "IN"
	default:
		return "Unknown"
	}
}

// StringFull prints an uppercase full word representation of the Traveling direction.
//
// You may optionally pass true for a lowercase representation.
func (t Traveling) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	switch t {
	case Westbound:
		if lower {
			return "westbound"
		}
		return "Westbound"
	case Northbound:
		if lower {
			return "northbound"
		}
		return "Northbound"
	case Eastbound:
		if lower {
			return "eastbound"
		}
		return "Eastbound"
	case Southbound:
		if lower {
			return "southbound"
		}
		return "Southbound"
	case Outbound:
		if lower {
			return "outbound"
		}
		return "Outbound"
	case Inbound:
		if lower {
			return "inbound"
		}
		return "Inbound"
	default:
		if lower {
			return "unknown"
		}
		return "Unknown"
	}
}
