package std

import (
	"github.com/ignite-laboratories/core/std/num"
	"math"
	"math/big"
)

// Bounded represents a numeric value bound within the closed set [minimum, maximum].
// Additionally, all bounded types can be 'clamped' into the bounded range - meaning that
// they will not automatically overflow or underflow when they exceed the bounds.
type Bounded[T num.ExtendedPrimitive] struct {
	value   T
	minimum T
	maximum T
	Clamp   bool
}
type BoundedByType[T num.ExtendedPrimitive] struct {
	value T
	clamp bool
}

// Value returns the currently held Bounded value.
func (bnd Bounded[T]) Value() T {
	return bnd.value
}

// Minimum returns the current minimum boundary.
func (bnd Bounded[T]) Minimum() T {
	return bnd.minimum
}

// Maximum returns the current maximum boundary.
func (bnd Bounded[T]) Maximum() T {
	return bnd.maximum
}

// SetAll sets the value and boundaries all in one operation, preventing multiple calls to Set().
//
// NOTE: The boundary parameters are evaluated to ensure the lower bound is always the 'minimum'
func (bnd Bounded[T]) SetAll(value, a, b T) Bounded[T] {
	if a > b {
		a, b = b, a
	}
	bnd.minimum = a
	bnd.maximum = b
	return bnd.Set(value)
}

// SetBoundariesFromType sets the boundaries to the implied limits of the bounded type before calling Set(current value).
func (bnd Bounded[T]) SetBoundariesFromType() Bounded[T] {
	bnd.minimum = 0
	bnd.maximum = T(num.MaxValue[T]())
	return bnd.Set(bnd.value)
}

// SetBoundaries sets the boundaries before calling Set(current value).
//
// NOTE: The boundary parameters are evaluated to ensure the lower bound is always the 'minimum'
func (bnd Bounded[T]) SetBoundaries(a, b T) Bounded[T] {
	if a > b {
		a, b = b, a
	}
	bnd.minimum = a
	bnd.maximum = b
	return bnd.Set(bnd.value)
}

// Normalize converts the Bounded value to a float64 unit vector in the range [0.0, 1.0],
// where the bounded minimum maps to 0.0 and the bounded maximum maps to 1.0.
func (bnd Bounded[T]) Normalize() float64 {
	numerator := uint64(bnd.value - bnd.minimum)
	denominator := uint64(bnd.maximum - bnd.minimum)

	// If the bounded range cannot be represented in a float64, bump to big.Float
	// NOTE: float64 only maintains full precision up to 2⁵³

	if numerator > (1<<53) || denominator > (1<<53) {
		num := new(big.Float).SetUint64(numerator)
		den := new(big.Float).SetUint64(denominator)
		result, _ := new(big.Float).Quo(num, den).Float64()
		return result
	}

	return float64(numerator) / float64(denominator)
}

// Normalize32 converts the Bounded value to a float32 unit vector in the range [0.0, 1.0],
// where the bounded minimum maps to 0.0 and the bounded maximum maps to 1.0.
func (bnd Bounded[T]) Normalize32() float32 {
	return float32(bnd.Normalize())
}

// SetFromNormalized sets the bounded value using a float64 unit vector from the [0.0, 1.0]
// range, where 0.0 maps to the bounded minimum and 1.0 maps to the bounded maximum.
func (bnd Bounded[T]) SetFromNormalized(normalized float64) Bounded[T] {
	distance := uint64(bnd.maximum - bnd.minimum)

	// If the bounded range cannot be represented by a float64, bump to big.Float
	// NOTE: float64 only maintains full precision up to 2⁵³

	if distance > (1 << 53) {
		normalizedBig := new(big.Float).SetFloat64(normalized)
		distanceBig := new(big.Float).SetUint64(distance)
		result := new(big.Float).Mul(normalizedBig, distanceBig)

		// Add minimum after multiplication
		minimumBig := new(big.Float).SetInt64(int64(bnd.minimum))
		result.Add(result, minimumBig)

		// Convert back to integer
		val, _ := result.Int64()
		return bnd.Set(T(val))
	}

	scaled := normalized * float64(distance)
	return bnd.Set(T(scaled) + bnd.minimum)
}

// SetFromNormalized32 sets the bounded value using a float32 unit vector from the [0.0, 1.0]
// range, where 0.0 maps to the bounded minimum and 1.0 maps to the bounded maximum.
func (bnd Bounded[T]) SetFromNormalized32(normalized float32) Bounded[T] {
	return bnd.SetFromNormalized(float64(normalized))
}

// Set sets the value of Bounded and automatically handles when the value exceeds the boundaries.
func (bnd Bounded[T]) Set(value T) Bounded[T] {
	if bnd.Clamp {
		if value > bnd.maximum {
			value = bnd.maximum
		} else if value < bnd.minimum {
			value = bnd.minimum
		}
	} else {
		// NOTE: The maximum distance of a primitive type will ALWAYS be a uint64, which is very nice =)
		distance := uint64(bnd.maximum-bnd.minimum) + 1
		// NOTE: This circumvents conversion to a float64 when using Math.Abs()
		if distance < 0 {
			distance = -distance
		}

		// Check if the distance (or any of the stored values) exceeds an int64 - requiring big.Int
		needsBig := distance > uint64(math.MaxInt64)
		if !needsBig {
			switch any(value).(type) {
			case uint64, uint:
				needsBig = uint64(value) > uint64(math.MaxInt64) ||
					uint64(bnd.minimum) > uint64(math.MaxInt64) ||
					uint64(bnd.maximum) > uint64(math.MaxInt64)
			}
		}

		var diff uint64
		if needsBig {
			m := new(big.Int).SetUint64(uint64(bnd.minimum))
			v := new(big.Int).SetUint64(uint64(value))
			r := new(big.Int).SetUint64(distance)

			d := new(big.Int).Sub(v, m)
			if d.Sign() < 0 {
				d = new(big.Int).Add(d, r)
			}

			diff = d.Uint64()
		} else {
			d := int64(value - bnd.minimum)
			if d < 0 {
				d += int64(distance)
			}
			diff = uint64(d)
		}

		mod := T(diff % distance)
		value = bnd.minimum + mod
	}

	bnd.value = value
	return bnd
}
