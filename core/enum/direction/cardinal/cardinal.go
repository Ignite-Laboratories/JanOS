// Package cardinal provides access to the cardinal.Direction enumeration.
package cardinal

import (
	"github.com/ignite-laboratories/core/std/num"
)

// Direction represents map-oriented spatial directions.
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
// See direction.Any, Direction, Longitudinal, Latitudinal, North, West, South, and East
type Direction[T num.Primitive] interface {
	Longitudinal[T] | Latitudinal[T]
}

// Longitudinal represents only the cardinal axis of East and West.
//
// See direction.Any, Direction, Longitudinal, Latitudinal, North, West, South, and East
type Longitudinal[T num.Primitive] interface {
	East[T] | West[T]
}

// Latitudinal represents only the cardinal axis of North and South.
//
// See direction.Any, Direction, Longitudinal, Latitudinal, North, West, South, and East
type Latitudinal[T num.Primitive] interface {
	North[T] | South[T]
}

// North represents the cardinal Direction "up" - which is the direction of accumulation.
//
// See direction.Any, Direction, Longitudinal, Latitudinal, North, West, South, and East
type North[T num.Primitive] num.Numeric[T]

func (_ North[T]) String() string {
	return "N"
}

func (_ North[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "north"
	}
	return "North"
}

// East represents the cardinal Direction "right" - which is the direction of reduction.
//
// See direction.Any, Direction, Longitudinal, Latitudinal, North, West, South, and East
type East[T num.Primitive] num.Numeric[T]

func (_ East[T]) String() string {
	return "E"
}

func (_ East[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "east"
	}
	return "East"
}

// South represents the cardinal Direction "down" - which is the target of all calculation
//
// See direction.Any, Direction, Longitudinal, Latitudinal, North, West, South, and East
type South[T num.Primitive] num.Numeric[T]

func (_ South[T]) String() string {
	return "S"
}

func (_ South[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "south"
	}
	return "South"
}

// West represents the cardinal Direction "left" - which is the direction of scale.
//
// See direction.Any, Direction, Longitudinal, Latitudinal, North, West, South, and East
type West[T num.Primitive] num.Numeric[T]

func (_ West[T]) String() string {
	return "W"
}

func (_ West[T]) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	if lower {
		return "west"
	}
	return "West"
}
