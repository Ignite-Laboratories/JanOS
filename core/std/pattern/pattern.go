// Package pattern provides access to creating patterns From data.
package pattern

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/bounded"
	"github.com/ignite-laboratories/core/std/internal"
	"github.com/ignite-laboratories/core/std/num"
)

// NilAny returns a pattern which always yields a nil interface value (type any).
func NilAny() std.Pattern[any] {
	fns := func(int) any {
		return nil
	}
	fnm := func(int) []any {
		return []any{nil}
	}
	return internal.NewPattern[any](std.NewEmit(fns, fnm, fns, fnm), nil)
}

// Nil returns a pattern which always yields nil.
func Nil[T any]() std.Pattern[*T] {
	return Zero[*T]()
}

// Zero returns a pattern which always yields the zero value of T.
func Zero[T any]() std.Pattern[T] {
	var zero T
	fns := func(int) T {
		return zero
	}
	fnm := func(int) []T {
		return []T{zero}
	}
	return internal.NewPattern[T](std.NewEmit(fns, fnm, fns, fnm), zero)
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

	return internal.NewPattern[T](&c, walkEast, walkWest, walkTo, yieldEast, yieldWest, yieldTo, data...)
}
