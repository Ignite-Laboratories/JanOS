// Package orthogonal provides access to the orthogonal.Direction enumeration.
package orthogonal

// Direction represents a single orthogonal side.
//
// All dimensions can be distilled down to an infinitely repeating number line which can be traversed in binary directions -
// but, as you layer these dimensions on top of each other, they orthographically align relative to one another.  The terminology
// used to describe this is entirely dependent upon context, and as such I've provided a robust set of general abstract dimensions
// from which to describe this mechanic in code.  It truly does NOT matter which you use, as long as the called method knows
// how to talk in THAT language. =)
//
// Abstractly, the result of calculation (the target) is always relatively "down" (or "towards the enemy gate") no matter YOUR orientation
// in space.  Mentally this may be the direction of "gravity" while standing up and writing calculations on a whiteboard, but I think Ender
// described it best.  All binary data is oriented with the most-significant side towards the "left" (or "west").  When operating against a
// matrix, you walk "latitudinally" between rows along the Y axis and "longitudinally" between columns along the X axis.  Against a voxel,
// you'd walk negatively "in" or positively "out" along the Z axis.
//
// Abstract references consider your relative orientation as you float through the void of time and spatial calculation.
//
// See direction.Any, Direction, LeftOrRight, UpOrDown, InOrOut, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type Direction interface {
	Left | Right | Up | Down | In | Out
}

// LeftOrRight represents the orthogonal 'sides' of the XAxis.
//
// See direction.Any, Direction, LeftOrRight, UpOrDown, InOrOut, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type LeftOrRight interface {
	Left | Right
}

// UpOrDown represents the orthogonal 'sides' of the YAxis.
//
// See direction.Any, Direction, LeftOrRight, UpOrDown, InOrOut, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type UpOrDown interface {
	Up | Down
}

// InOrOut represents the orthogonal 'sides' of the ZAxis.
//
// See direction.Any, Direction, LeftOrRight, UpOrDown, InOrOut, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type InOrOut interface {
	In | Out
}

// Axis represents an axis of an orthogonal direction.
//
// See direction.Any, Direction, LeftOrRight, UpOrDown, InOrOut, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type Axis interface {
	XAxis | YAxis | ZAxis
}

// XAxis represents the Left ↔ Right axis.
//
// See direction.Any, Direction, LeftOrRight, UpOrDown, InOrOut, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type XAxis interface {
	Left | Right
}

// YAxis represents the Up ↕ Down axis.
//
// See direction.Any, Direction, LeftOrRight, UpOrDown, InOrOut, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type YAxis interface {
	Up | Down
}

// XYAxis represents the Up ↕ Down and Left ↔ Right axes.
type XYAxis interface {
	XAxis | YAxis
}

// XYZAxis represents the Up ↕ Down, Left ↔ Right, and In ⇌ Out axes.
type XYZAxis interface {
	XAxis | YAxis | ZAxis
}

// ZAxis represents the In ↔ Out axis.
//
// See direction.Any, Direction, LeftOrRight, UpOrDown, InOrOut, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type ZAxis interface {
	In | Out
}

// In represents the orthogonal Direction "negatively along the Z axis perpendicular to the orthogonal XY plane."
//
// See direction.Any, Direction, LeftOrRight, UpOrDown, InOrOut, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
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

// Out represents the orthogonal Direction "positively along the Z axis perpendicular to the orthogonal XY plane."
//
// See direction.Any, Direction, LeftOrRight, UpOrDown, InOrOut, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
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

// Up represents the orthogonal Direction "towards the top of the orthogonal XY plane."
//
// See direction.Any, Direction, LeftOrRight, UpOrDown, InOrOut, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
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

// Down represents the orthogonal Direction "towards the bottom of the orthogonal XY plane."
//
// See direction.Any, Direction, LeftOrRight, UpOrDown, InOrOut, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
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

// Left represents the orthogonal Direction "towards the left of the orthogonal XY plane."
//
// See direction.Any, Direction, LeftOrRight, UpOrDown, InOrOut, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
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

// Right represents the orthogonal Direction "towards the right of the orthogonal XY plane."
//
// See direction.Any, Direction, LeftOrRight, UpOrDown, InOrOut, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
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
