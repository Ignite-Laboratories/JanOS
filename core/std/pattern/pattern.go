// Package pattern provides access to creating patterns From data.
package pattern

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/bounded"
	"github.com/ignite-laboratories/core/std/num"
)

// Any represents any general pattern type, regardless of dimensionality.
//
// NOTE: For advanced pattern generation and predefined patterns, see the 'std/pattern' package.
//
// See Any, UpTo2D, UpTo3D, and UpTo4D
type Any[T any] interface {
	std.Pattern[T] | std.Pattern2D[T] | std.Pattern3D[T] | std.Pattern4D[T]
}

// UpTo2D represents patterns up to 2 dimensions wide.
//
// See Any, UpTo2D, UpTo3D, and UpTo4D
type UpTo2D[T any] interface {
	std.Pattern[T] | std.Pattern2D[T]
}

// UpTo3D represents patterns up to 3 dimensions wide.
//
// See Any, UpTo2D, UpTo3D, and UpTo4D
type UpTo3D[T any] interface {
	std.Pattern[T] | std.Pattern2D[T] | std.Pattern3D[T]
}

// UpTo4D represents patterns up to 4 dimensions wide.
//
// See Any, UpTo2D, UpTo3D, and UpTo4D
type UpTo4D[T any] interface {
	std.Pattern[T] | std.Pattern2D[T] | std.Pattern3D[T] | std.Pattern4D[T]
}

// NilAny returns a pattern which always yields a nil interface value (type any).
func NilAny() std.Pattern[any] {
	fn := func() any {
		return nil
	}
	cursor := bounded.By[uint](0, 0, 0)
	return std.NewPattern[any](&cursor, fn, fn, nil)
}

// Nil returns a pattern which always yields a nil pointer of type *T.
func Nil[T any]() std.Pattern[*T] {
	return Zero[*T]()
}

// Zero returns a pattern which always yields the zero value of T.
func Zero[T any]() std.Pattern[T] {
	var zero T
	fn := func() T {
		return zero
	}
	cursor := bounded.By[uint](0, 0, 0)
	return std.NewPattern[T](&cursor, fn, fn, zero)
}

// One is a pattern of a numeric one.
func One[T num.Primitive]() std.Pattern[T] {
	return From[T](1)
}

// ZeroOne returns a pattern which always yields a numeric '01'.
//
// NOTE: Patterns are stored as you would read them (left→to→right) but are evaluated in a direction of travel.
func ZeroOne[T num.Primitive]() std.Pattern[T] {
	return From[T](0, 1)
}

// OneZero returns a pattern which always yields a numeric '10'.
//
// NOTE: Patterns are stored as you would read them (left→to→right) but are evaluated in a direction of travel.
func OneZero[T num.Primitive]() std.Pattern[T] {
	return From[T](1, 0)
}

// From creates a new std.Pattern which can infinitely walk through the provided data either westbound or eastbound.
//
// NOTE: This will create a single element 'zero' instance pattern of T if provided no data.
func From[T any](data ...T) std.Pattern[T] {
	if len(data) == 0 {
		var zero T
		data = append(data, zero)
	}

	c := bounded.By[uint](0, 0, uint(len(data)-1))
	walkEast := func() T {
		out := data[c.Value()]
		c.IncrementPtr()
		return out
	}
	walkWest := func() T {
		c.DecrementPtr()
		return data[c.Value()]
	}

	return std.NewPattern[T](&c, walkEast, walkWest, data...)
}
