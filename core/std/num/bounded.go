package num

import (
	"fmt"
	"math/big"
)

// Bounded represents a numeric value bound within the closed set [minimum, maximum].
// Additionally, all bounded types can be 'clamped' into the bounded range - meaning that
// they will not automatically overflow or underflow when they exceed the bounds.
type Bounded[T ExtendedPrimitive] struct {
	value   T
	minimum T
	maximum T
	Clamp   bool
}
type BoundedByType[T ExtendedPrimitive] struct {
	value T
	clamp bool
}

// Value returns the currently held Bounded value.
func (bnd *Bounded[T]) Value() T {
	return bnd.value
}

// Minimum returns the current minimum boundary.
func (bnd *Bounded[T]) Minimum() T {
	return bnd.minimum
}

// Maximum returns the current maximum boundary.
func (bnd *Bounded[T]) Maximum() T {
	return bnd.maximum
}

// SetAll sets the value and boundaries all in one operation, preventing multiple calls to Set().
//
// NOTE: The boundary parameters are evaluated to ensure the lower bound is always the 'minimum'
func (bnd *Bounded[T]) SetAll(value, a, b T) *Bounded[T] {
	if a > b {
		a, b = b, a
	}
	bnd.minimum = a
	bnd.maximum = b
	return bnd.Set(value)
}

// SetBoundaries sets the boundaries before calling Set(current value).
//
// NOTE: The boundary parameters are evaluated to ensure the lower bound is always the 'minimum'
func (bnd *Bounded[T]) SetBoundaries(a, b T) *Bounded[T] {
	if a > b {
		a, b = b, a
	}
	bnd.minimum = a
	bnd.maximum = b
	return bnd.Set(bnd.value)
}

// Set sets the value of Bounded and automatically handles when the value exceeds the boundaries.
func (bnd *Bounded[T]) Set(value T) *Bounded[T] {
	signed := IsSigned[T]()
	if bnd.Clamp {
		if value > bnd.maximum {
			value = bnd.maximum
		} else if value < bnd.minimum {
			value = bnd.minimum
		}
	} else {
		// NOTE: The maximum distance of a primitive type will ALWAYS be a uint64, which is very nice =)
		distance := uint64(bnd.maximum - bnd.minimum)
		// NOTE: This circumvents conversion to a float64 to use Math.Abs()
		if distance < 0 {
			distance = -distance
		}

		// NOTE: We bump to big here since the underflow logic requires negative values, but the type might restrict that
		var diff uint64
		if signed {
			m := new(big.Int).SetInt64(int64(bnd.minimum))
			v := new(big.Int).SetInt64(int64(value))
			r := new(big.Int).SetUint64(distance)

			d := new(big.Int).Sub(v, m)
			if d.Sign() < 0 {
				d = new(big.Int).Add(d, r)
			}

			diff = d.Uint64()
		} else {
			m := new(big.Int).SetUint64(uint64(bnd.minimum))
			v := new(big.Int).SetUint64(uint64(value))
			r := new(big.Int).SetUint64(distance)

			d := new(big.Int).Sub(v, m)
			if d.Sign() < 0 {
				d = new(big.Int).Add(d, r)
			}

			diff = d.Uint64()
		}

		mod := T(diff % distance)
		value = bnd.minimum + mod
	}

	bnd.value = value
	return bnd
}

func (bnd Bounded[T]) String() string {
	return fmt.Sprintf("%v", bnd.value)
}
