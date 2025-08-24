// Package orthogonal provides access to the orthogonal.Direction enumeration.
package orthogonal

// Direction is an "enumeration" that represents an orthographic directional side.
//
// All dimensions can be distilled down to an infinitely repeating number line which can be traversed only in binary directions -
// but, as you layer these dimensions on top of each other, they orthographically align relative to one another.  The terminology
// used to describe this is entirely dependent upon context, and as such I've provided a robust set of general abstract dimensions
// from which to describe this mechanic in code.  It truly does NOT matter which you use, as long as the called method knows
// how to talk in THAT language. =)
//
// For all directional enumerations, your intuition is probably spot on in how it's intended to be used - this is -my- current logic behind them:
//
// Abstractly, the result of calculation (the target) is always relatively "down" (or "towards the enemies gate") no matter YOUR orientation
// in space.  Mentally this may be the direction of "gravity" while standing up and writing calculations on a whiteboard, but I think Ender
// described it best.  All binary data is oriented with the most-significant side towards the "left" (or "west").  When operating against a
// matrix, you walk "latitudinally" between rows along the Y axis and "longitudinally" between columns on the X axis.  Against a voxel or 3D
// structure, you'd walk negatively "in" or positively "out" along the Z axis.
//
// The idea is these abstract directions consider your relative orientation as you float through the void of time and spatial calculation.
//
// The orthographic enumeration has seven ultimate directions:
//
//	Static (0), Left (-1), Right (1), Up (1), Down (-1), In (-1), and Out (1)
//
// You have four interfaces for constraining these directions to subsets of this interface -
//
//  0. Direction - Static (0), Left (-1), Right (1), Up (1), Down (-1), In (-1), and Out (1)
//  1. Pole - LeftOrRight (-1 and 1), UpOrDown (1 and -1), and InOrOut (-1 and 1) ← Excludes Static
//  2. Axis - XAxis (-1, 0, 1), YAxis (-1, 0, 1), and ZAxis (-1, 0, 1) ← Includes Static
//  3. Volume - XYPlane, XZPlane, and YZPlane ← Includes Static
//
// If you wish to describe Static movement along a particular axis, three aliases for Static are provided: XStatic, YStatic, and ZStatic
//
// NOTE: These represent traversal constraints applied to motion, thus they are -relative- directions.  This is why there is no
// 'origin' or 'home', as that implies an absolute position - instead, 'Static' acts as a stand-in for 'no movement'.
//
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
type Direction interface {
	Left | Right | Up | Down | In | Out | Static
}

// LeftOrRight represents the orthogonally polar 'sides' of the XAxis.
//
// NOTE: If you wish to include the Static direction, please see XAxis.
//
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
type LeftOrRight interface {
	Left | Right
}

// UpOrDown represents the orthogonally polar 'sides' of the YAxis.
//
// NOTE: If you wish to include the Static direction, please see YAxis.
//
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
type UpOrDown interface {
	Up | Down
}

// InOrOut represents the orthogonally polar 'sides' of the ZAxis.
//
// NOTE: If you wish to include the Static direction, please see ZAxis.
//
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
type InOrOut interface {
	In | Out
}

// Pole represents the orthogonally polar 'sides' of LeftOrRight, UpOrDown, and InOrOut - inherently excluding Static.
//
// NOTE: If you wish to include the Static direction, please see Axis.
//
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
type Pole interface {
	LeftOrRight | UpOrDown | InOrOut
}

// Axis represents an orthogonal 'axis' of XAxis, YAxis, and ZAxis - including Static.
//
// NOTE: If you do not wish to include the Static direction, please see Pole.
//
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
type Axis interface {
	XAxis | YAxis | ZAxis
}

// XAxis represents the Left ↔ Right axis, including Static.
//
// NOTE: If you do not wish to include the Static direction, please see LeftOrRight.
//
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
type XAxis interface {
	LeftOrRight | Static
}

// YAxis represents the Up ↕ Down axis, including Static.
//
// NOTE: If you do not wish to include the Static direction, please see UpOrDown.
//
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
type YAxis interface {
	UpOrDown | Static
}

// ZAxis represents the In ↔ Out axis, including Static.
//
// NOTE: If you do not wish to include the Static direction, please see InOrOut.
//
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
type ZAxis interface {
	InOrOut | Static
}

// XYPlane represents the Up ↕ Down and Left ↔ Right plane, including Static.
//
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
type XYPlane interface {
	XAxis | YAxis
}

// XZPlane represents the Up ↕ Down and Left ↔ Right plane, including Static.
//
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
type XZPlane interface {
	XAxis | ZAxis
}

// YZPlane represents the Up ↕ Down and In ⇌ Out plane, including Static.
//
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
type YZPlane interface {
	YAxis | ZAxis
}

// Volume represents the Up ↕ Down, Left ↔ Right, and In ⇌ Out axes, including Static.
//
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
type Volume interface {
	XAxis | YAxis | ZAxis
}

// XStatic represents the orthogonal Direction of 'no movement' (0) along the X axis.
//
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
type XStatic = Static

// YStatic represents the orthogonal Direction of 'no movement' (0) along the Y axis.
//
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
type YStatic = Static

// ZStatic represents the orthogonal Direction of 'no movement' (0) along the Z axis.
//
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
type ZStatic = Static

// Static represents the orthogonal Direction of 'no movement' (0).
//
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
type Static byte

func (_ Static) String() string {
	return "ℴ"
}

func (_ Static) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "origin"
	}
	return "Static"
}

// In represents the orthogonal Direction "negatively along the Z axis perpendicular to the orthogonal XY plane."
//
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
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
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
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
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
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
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
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
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
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
// See direction.Spatial, Direction, Pole, Axis, Volume, In, Out, Up, Up, Down, Down, Left, Right, Left, Right, B, A, Start
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
