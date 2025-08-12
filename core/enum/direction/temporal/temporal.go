// Package temporal provides access to the temporal.Direction enumeration.
package temporal

import "github.com/ignite-laboratories/core/std/num"

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
type Direction[T num.Primitive] interface {
	Past[T] | Present[T] | Future[T]
}

// Past represents the abstract Direction of "historically" - which is the direction of reflection.
//
// See Direction, Past, Present, and Future
type Past[T num.Primitive] num.Numeric[T]

func (_ Past[T]) String() string {
	return "⏮"
}

func (_ Past[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "past"
	}
	return "Past"
}

// Present represents the abstract Direction of "currently" - which is the direction of experience.
//
// See Direction, Past, Present, and Future
type Present[T num.Primitive] num.Numeric[T]

func (_ Present[T]) String() string {
	return "⏸"
}

func (_ Present[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "present"
	}
	return "Present"
}

// Future represents the abstract Direction of "eminently" - which is the direction of anticipation.
//
// See Direction, Past, Present, and Future
type Future[T num.Primitive] num.Numeric[T]

func (_ Future[T]) String() string {
	return "⏭"
}

func (_ Future[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "future"
	}
	return "Future"
}
