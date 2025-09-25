package num

import (
	"fmt"
)

// NewNumericBounded creates a new instance of Numeric[T] bounded in the closed interval [minimum, maximum] and returns the initial set operation's Breach.
//
// See Numeric, NewNumeric, and NewNumericBounded
func NewNumericBounded[T Advanced](value, minimum, maximum T, clamp ...bool) (Numeric[T], Breach) {
	c := len(clamp) > 0 && clamp[0]
	var zero T
	if len(clamp) == 0 && IsFloat(zero) {
		c = true
	}
	b := Numeric[T]{
		minimum:     minimum,
		maximum:     maximum,
		initialized: true,
		unbounded:   false,
		Clamp:       c,
	}
	err := b.Set(value)
	return b, err
}

// NewNumeric creates a new instance of Numeric[T] which does not intercept the provided types boundaries.
//
// See Numeric, NewNumeric, and NewNumericBounded
func NewNumeric[T Advanced](value T) Numeric[T] {
	switch raw := any(value).(type) {
	case Natural, Realized, complex64, complex128:
		// CRITICAL: This does NOT type assert the string advanced types!!!
		// ALWAYS pass through the generic type for tiny types - never type assert and then mutate
		return newNumericAdvanced(value)
	case int:
		return any(newNumericPrimitive(raw)).(Numeric[T])
	case int8:
		return any(newNumericPrimitive(raw)).(Numeric[T])
	case int16:
		return any(newNumericPrimitive(raw)).(Numeric[T])
	case int32:
		return any(newNumericPrimitive(raw)).(Numeric[T])
	case int64:
		return any(newNumericPrimitive(raw)).(Numeric[T])
	case uint:
		return any(newNumericPrimitive(raw)).(Numeric[T])
	case uint8:
		return any(newNumericPrimitive(raw)).(Numeric[T])
	case uint16:
		return any(newNumericPrimitive(raw)).(Numeric[T])
	case uint32:
		return any(newNumericPrimitive(raw)).(Numeric[T])
	case uint64:
		return any(newNumericPrimitive(raw)).(Numeric[T])
	case uintptr:
		return any(newNumericPrimitive(raw)).(Numeric[T])
	case float32:
		return any(newNumericPrimitive(raw)).(Numeric[T])
	case float64:
		return any(newNumericPrimitive(raw)).(Numeric[T])
	default:
		panic(fmt.Errorf("unknown type %T", value))
	}
}

func newNumericAdvanced[T Advanced](value T) Numeric[T] {
	// TODO: Implement this
	panic("not implemented yet")
	return Numeric[T]{}
}

func newNumericPrimitive[T Primitive](value T) Numeric[T] {
	b := Numeric[T]{
		minimum:     MinValue[T](),
		maximum:     MaxValue[T](),
		initialized: true,
		unbounded:   true,
		Clamp:       false,
	}
	_ = b.Set(value)
	return b
}
