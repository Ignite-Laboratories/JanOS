// Package pattern provides access to creating patterns From data.
package pattern

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/bounded"
	"github.com/ignite-laboratories/core/std/num"
)

// NilAny returns a pattern which always yields a nil interface value (type any).
func NilAny() std.Pattern[any] {
	fn := func() any {
		return nil
	}
	return std.NewPattern[any](fn, fn, nil)
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
	return std.NewPattern[T](fn, fn, zero)
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

	b := bounded.By[int](0, 0, len(data)-1)
	walkEast := func() T {
		out := data[b.Value()]
		b.IncrementPtr()
		return out
	}
	walkWest := func() T {
		b.DecrementPtr()
		return data[b.Value()]
	}

	return std.NewPattern[T](walkEast, walkWest, data...)
}
