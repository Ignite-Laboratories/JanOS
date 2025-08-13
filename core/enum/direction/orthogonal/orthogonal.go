// Package orthogonal provides access to the orthogonal.Direction enumeration.
package orthogonal

import (
	"github.com/ignite-laboratories/core/std/num"
)

// Direction represents a single orthogonal direction.
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
// See direction.Any, Direction, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type Direction[T num.Primitive] interface {
	Left[T] | Right[T] | Up[T] | Down[T] | In[T] | Out[T]
}

// Axis represents an axis of an orthogonal direction.
//
// See direction.Any, Direction, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type Axis[T num.Primitive] interface {
	XAxis[T] | YAxis[T] | ZAxis[T]
}

// XAxis represents the Left ↔ Right axis.
//
// See direction.Any, Direction, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type XAxis[T num.Primitive] interface {
	Left[T] | Right[T]
}

// YAxis represents the Up ↕ Down axis.
//
// See direction.Any, Direction, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type YAxis[T num.Primitive] interface {
	Up[T] | Down[T]
}

// XYAxis represents the Up ↕ Down and Left ↔ Right axes.
type XYAxis[T num.Primitive] interface {
	XAxis[T] | YAxis[T]
}

// XYZAxis represents the Up ↕ Down, Left ↔ Right, and In ⇌ Out axes.
type XYZAxis[T num.Primitive] interface {
	XAxis[T] | YAxis[T] | ZAxis[T]
}

// ZAxis represents the In ↔ Out axis.
//
// See direction.Any, Direction, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type ZAxis[T num.Primitive] interface {
	In[T] | Out[T]
}

// In represents the orthogonal Direction "negatively along the Z axis perpendicular to the orthogonal XY plane."
//
// See direction.Any, Direction, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type In[T num.Primitive] num.Numeric[T]

func (_ In[T]) String() string {
	return "⤓"
}

func (_ In[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "in"
	}
	return "In"
}

// Out represents the orthogonal Direction "positively along the Z axis perpendicular to the orthogonal XY plane."
//
// See direction.Any, Direction, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type Out[T num.Primitive] num.Numeric[T]

func (_ Out[T]) String() string {
	return "↥"
}

func (_ Out[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "out"
	}
	return "Out"
}

// Up represents the orthogonal Direction "towards the top of the orthogonal XY plane."
//
// See direction.Any, Direction, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type Up[T num.Primitive] num.Numeric[T]

func (_ Up[T]) String() string {
	return "↑"
}

func (_ Up[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "up"
	}
	return "Up"
}

// Down represents the orthogonal Direction "towards the bottom of the orthogonal XY plane."
//
// See direction.Any, Direction, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type Down[T num.Primitive] num.Numeric[T]

func (_ Down[T]) String() string {
	return "↓"
}

func (_ Down[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "down"
	}
	return "Down"
}

// Left represents the orthogonal Direction "towards the left of the orthogonal XY plane."
//
// See direction.Any, Direction, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type Left[T num.Primitive] num.Numeric[T]

func (_ Left[T]) String() string {
	return "←"
}

func (_ Left[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "left"
	}
	return "Left"
}

// Right represents the orthogonal Direction "towards the right of the orthogonal XY plane."
//
// See direction.Any, Direction, Axis, XAxis, YAxis, ZAxis, In, Out, Up, Up, Down, Down, Left, Right, B, A, Start
type Right[T num.Primitive] num.Numeric[T]

func (_ Right[T]) String() string {
	return "→"
}

func (_ Right[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "right"
	}
	return "Right"
}
