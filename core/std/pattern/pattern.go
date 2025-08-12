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
	fn := func(i uint) any {
		return nil
	}
	fns := func(i uint) []any {
		return []any{nil}
	}
	cursor, _ := bounded.By[uint](0, 0, 0)
	return std.NewPattern[any](&cursor, fn, fn, fn, fns, fns, fns, nil)
}

// Nil returns a pattern which always yields nil.
func Nil[T any]() std.Pattern[*T] {
	return Zero[*T]()
}

// Zero returns a pattern which always yields the zero value of T.
func Zero[T any]() std.Pattern[T] {
	var zero T
	fn := func(i uint) T {
		return zero
	}
	fns := func(i uint) []T {
		return []T{zero}
	}
	cursor, _ := bounded.By[uint](0, 0, 0)
	return std.NewPattern[T](&cursor, fn, fn, fn, fns, fns, fns, zero)
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

	c, _ := bounded.By[uint](0, 0, uint(len(data)-1))
	walkEast := func(i uint) T {
		out := data[c.Value()]
		_ = c.IncrementPtr(i)
		return out
	}
	walkWest := func(i uint) T {
		_ = c.DecrementPtr(i)
		return data[c.Value()]
	}
	walkTo := func(i uint) T {
		_ = c.SetPtr(i)
		return data[c.Value()]
	}
	yieldEast := func(i uint) []T {
		out := make([]T, i)
		for j := uint(0); j < i; j++ {
			out[j] = walkEast(1)
		}
		return out
	}
	yieldWest := func(i uint) []T {
		out := make([]T, i)
		for j := uint(0); j < i; j++ {
			out[j] = walkWest(1)
		}
		return out
	}
	yieldTo := func(i uint) []T {
		if i > c.Value() {
			return yieldEast(i - c.Value())
		} else if i < c.Value() {
			return yieldWest(c.Value() - i)
		}
		return []T{data[c.Value()]}
	}

	return std.NewPattern[T](&c, walkEast, walkWest, walkTo, yieldEast, yieldWest, yieldTo, data...)
}
