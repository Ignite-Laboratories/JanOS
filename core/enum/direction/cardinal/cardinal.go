// Package cardinal provides access to the cardinal.Direction enumeration.
package cardinal

import "github.com/ignite-laboratories/core/std/num"

// Direction represents map-oriented spatial directions.
//
// Abstractly, the result of calculation (the target) is always relatively "down" (or "towards the enemy gate") no matter YOUR orientation
// in space.  Mentally this may be the direction of "gravity" while standing up and writing calculations on a whiteboard, but I think Ender
// described it best.  All binary data is oriented with the most-significant side towards the "left" (or "west").  When operating against a
// matrix, the abstract orientation of "rows" are aligned with the cardinal direction of "south".
//
// Abstract references consider your relative orientation as you float through the void of time and spatial calculation.
//
// See Direction, North, West, South, and East
type Direction[T num.Primitive] interface {
	Longitudinal[T] | Latitudinal[T]
}

// Longitudinal represents only the cardinal directions of East and West.
type Longitudinal[T num.Primitive] interface {
	East[T] | West[T]
}

// Latitudinal represents only the cardinal directions of North and South.
type Latitudinal[T num.Primitive] interface {
	North[T] | South[T]
}

// North represents the cardinal Direction "up" - which is the direction of accumulation.
//
// See Direction, North, West, South, and East
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
// See Direction, North, West, South, and East
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
// See Direction, North, West, South, and East
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
// See Direction, North, West, South, and East
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
