// Package traffic provides access to the traffic.Direction enumeration.
package traffic

import "github.com/ignite-laboratories/core/std/num"

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
type Direction[T num.Primitive] interface {
	Inbound[T] | Outbound[T] | Bidirectional[T]
}

// Inbound represents the abstract Direction of "receiving" - which is the direction of listening.
//
// See Direction, Inbound, Outbound, and Bidirectional
type Inbound[T num.Primitive] num.Numeric[T]

func (_ Inbound[T]) String() string {
	return "⇤"
}

func (_ Inbound[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "inbound"
	}
	return "Inbound"
}

// Outbound represents the abstract Direction of "transmitting" - which is the direction of talking.
//
// See Direction, Inbound, Outbound, and Bidirectional
type Outbound[T num.Primitive] num.Numeric[T]

func (_ Outbound[T]) String() string {
	return "↦"
}

func (_ Outbound[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "outbound"
	}
	return "Outbound"
}

// Bidirectional represents the abstract Direction of "communication" - which is the direction of discourse.
//
// See Direction, Inbound, Outbound, and Bidirectional
type Bidirectional[T num.Primitive] num.Numeric[T]

func (_ Bidirectional[T]) String() string {
	return "⇹"
}

func (_ Bidirectional[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "bidirectional"
	}
	return "Bidirectional"
}
