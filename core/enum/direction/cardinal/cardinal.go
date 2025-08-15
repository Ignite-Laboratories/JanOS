// Package cardinal provides access to the cardinal.Any enumeration system.
package cardinal

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/num"
)

// Any represents any cardinal direction.
//
// NOTE: For complex intercardinal dimensions, see the Inter and InterSecondary types.
//
// See Any, Complex, Longitudinal, Latitudinal, and Axis
type Any[T num.Primitive] = std.Vector2D[T]

func New[T num.Primitive](value ...std.Vector2D[T]) Any[T] {
	if len(value) == 0 {
		return Any[T](std.Vector2D[T]{T(0), T(0)})
	}
	return T(value[0])
}

func North[T Any](value ...T) T {
	if len(value) == 0 {
		return T(1)
	}
	return T(value[0])
}

// NorthByNorthEast represents a cardinal.Any.
var NorthByNorthEast = InterSecondary[T]{North, Inter[T]{North, East}}

// NorthEast represents a cardinal.Any.
var NorthEast = Inter[T]{North, East}

// EastByNorthEast represents a cardinal.Any.
var EastByNorthEast = InterSecondary[T]{East, Inter[T]{North, East}}

// East represents a cardinal.Any.
var East Any[T]

// EastBySouthEast represents a cardinal.Any.
var EastBySouthEast = InterSecondary[T]{East, Inter[T]{South, East}}

// SouthEast represents a cardinal.Any.
var SouthEast = InterDirection{South, East}

// SouthBySouthEast represents a cardinal.Any.
var SouthBySouthEast = InterSecondary[T]{South, Inter[T]{South, East}}

// South represents a cardinal.Any.
var South Any[T]

// SouthBySouthWest represents a cardinal.Any.
var SouthBySouthWest = InterSecondary[T]{South, Inter[T]{South, West}}

// SouthWest represents a cardinal.Any.
var SouthWest = Inter[T]{South, West}

// WestBySouthWest represents a cardinal.Any.
var WestBySouthWest = InterSecondary[T]{West, Inter[T]{South, West}}

// West represents a cardinal.Any.
var West Any[T]

// WestByNorthWest represents a cardinal.Any.
var WestByNorthWest = InterSecondary[T]{West, Inter[T]{North, West}}

// NorthWest represents a cardinal.Any.
var NorthWest = Inter[T]{North, West}

// NorthByNorthWest represents a cardinal.Any.
var NorthByNorthWest = InterSecondary[T]{North, Inter[T]{North, West}}
