package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
)

// Difference returns a-b.
func Difference[TValue core.Numeric](a TValue, b TValue) TValue {
	return a - b
}

// Delta returns b-a.
func Delta[TValue core.Numeric](a TValue, b TValue) TValue {
	return b - a
}

// Change functions are called when a dimension's current point value changes.
type Change[TValue any] func(ctx core.Context, old std.Data[TValue], current std.Data[TValue])

// PointCalculation functions calculate a contextual value.
type PointCalculation[T any] func(core.Context) T

// Integral functions take in a set of contextual values and calculate a result.
//
// They are also provided with a cache pointer that can hold values between activations.
type Integral[TIn any, TOut any, TCache any] func(core.Context, *TCache, []TIn) TOut

// Blend functions take in many dissimilar-typed values and output a result.
type Blend[TOut any] func(...any) TOut

// Operate functions take in two numeric values and output a result.
type Operate[TValue core.Numeric] func(TValue, TValue) TValue

// Comparator functions should return true if the two provided values are "equal."
type Comparator[TValue any] func(a TValue, b TValue) bool
