// Package orthogonal provides access to the orthogonal.Direction enumeration.
package orthogonal

// Direction represents a single orthogonal direction.
//
// Abstractly, the result of calculation (the target) is always relatively "down" (or "towards the enemy gate") no matter YOUR orientation
// in space.  Mentally this may be the direction of "gravity" while standing up and writing calculations on a whiteboard, but I think Ender
// described it best.  All binary data is oriented with the most-significant side towards the "left" (or "west").  When operating against a
// matrix, the abstract orientation of "rows" are aligned with the cardinal direction of "south".
//
// Abstract references consider your relative orientation as you float through the void of time and spatial calculation.
//
// See Direction, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type Direction interface {
	Left | Right | Up | Down | In | Out
}

// Axis represents an axis of an orthogonal direction.
type Axis interface {
	XAxis | YAxis | ZAxis
}

// XAxis represents the Left ↔ Right axis.
type XAxis interface {
	Left | Right
}

// YAxis represents the Up ↔ Down axis.
type YAxis interface {
	Up | Down
}

// ZAxis represents the In ↔ Out axis.
type ZAxis interface {
	In | Out
}

// In represents the orthogonal Direction "towards the viewport."
//
// See Direction, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type In byte

func (_ In) String() string {
	return "⤓"
}

func (_ In) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "in"
	}
	return "In"
}

// Out represents the orthogonal Direction "away from the viewport."
//
// See Direction, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type Out byte

func (_ Out) String() string {
	return "↥"
}

func (_ Out) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "out"
	}
	return "Out"
}

// Up represents the orthogonal Direction "towards the top of the viewport."
//
// See Direction, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type Up byte

func (_ Up) String() string {
	return "↑"
}

func (_ Up) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "up"
	}
	return "Up"
}

// Down represents the orthogonal Direction "towards the bottom of the viewport."
//
// See Direction, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type Down byte

func (_ Down) String() string {
	return "↓"
}

func (_ Down) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "down"
	}
	return "Down"
}

// Left represents the orthogonal Direction "towards the left of the viewport."
//
// See Direction, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type Left byte

func (_ Left) String() string {
	return "←"
}

func (_ Left) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "left"
	}
	return "Left"
}

// Right represents the orthogonal Direction "towards the right of the viewport."
//
// See Direction, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type Right byte

func (_ Right) String() string {
	return "→"
}

func (_ Right) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "right"
	}
	return "Right"
}
