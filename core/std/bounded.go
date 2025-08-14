package std

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/num"
	"math"
	"math/big"
)

// Bounded represents a numeric value bound within the closed set [minimum, maximum].
// Additionally, all bounded types can be 'clamped' into the bounded range - meaning that
// they will not automatically overflow or underflow when they exceed the bounds.
type Bounded[T num.Primitive] struct {
	value   T
	minimum T
	maximum T
	Clamp   bool
}

// NewBounded creates a new instance of Bounded[T].
//
// NOTE: While you can call this directly, the convention is to use the 'std/bounded' package.
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func NewBounded[T num.Primitive](value, minimum, maximum T, clamp ...bool) (Bounded[T], error) {
	c := len(clamp) > 0 && clamp[0]
	return Bounded[T]{
		value:   T(0),
		minimum: minimum,
		maximum: maximum,
		Clamp:   c,
	}.Set(value)
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

// Increment adds 1 or the provided amount to the bound value.
//
// NOTE: If you provide a negative number, this will 'decrement'
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd Bounded[T]) Increment(amount ...T) (Bounded[T], error) {
	i := T(1)
	if len(amount) > 0 {
		i = amount[0]
	}
	return bnd.Set(bnd.value + i)
}

// Decrement subtracts 1 or the provided amount from the bound value.
//
// NOTE: If you provide a negative number, this will 'increment'
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd Bounded[T]) Decrement(amount ...T) (Bounded[T], error) {
	i := T(1)
	if len(amount) > 0 {
		i = amount[0]
	}
	return bnd.Set(bnd.value - i)
}

// AddOrSubtract adds or subtracts the provided amount to the bound value.
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd Bounded[T]) AddOrSubtract(amount T) (Bounded[T], error) {
	if amount < 0 {
		amount = -amount
		return bnd.Decrement(T(amount))
	}
	return bnd.Increment(T(amount))
}

// SetAll sets the value and boundaries all in one operation, preventing multiple calls to Set().
//
// NOTE: The boundary parameters are evaluated to ensure the lower bound is always the 'minimum'
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd Bounded[T]) SetAll(value, a, b T, clamp ...bool) (Bounded[T], error) {
	c := len(clamp) > 0 && clamp[0]
	if a > b {
		a, b = b, a
	}
	bnd.minimum = a
	bnd.maximum = b
	bnd.Clamp = c
	return bnd.Set(value)
}

// SetBoundariesFromType sets the boundaries to the implied limits of the bounded type before calling Set(current value).
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd Bounded[T]) SetBoundariesFromType() (Bounded[T], error) {
	bnd.minimum = 0
	bnd.maximum = T(num.MaxValue[T]())
	return bnd.Set(bnd.value)
}

// SetBoundaries sets the boundaries before calling Set(current value).
//
// NOTE: The boundary parameters are evaluated to ensure the lower bound is always the 'minimum'
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd Bounded[T]) SetBoundaries(a, b T) (Bounded[T], error) {
	if a > b {
		a, b = b, a
	}
	bnd.minimum = a
	bnd.maximum = b
	return bnd.Set(bnd.value)
}

// Normalize converts the Bounded value to a float64 unit vector in the interval [0.0, 1.0],
// by linearly mapping the value from its bounded interval's [minimum, maximum]. A value equal
// to minimum maps to 0.0, a value equal to maximum maps to 1.0, and values in between
// are linearly interpolated.
func (bnd Bounded[T]) Normalize() float64 {
	numerator := uint64(bnd.value - bnd.minimum)
	denominator := uint64(bnd.maximum - bnd.minimum)

	// If the bounded range cannot be represented in a float64, bump to big.Float
	// NOTE: float64 only maintains full precision up to 2⁵³

	if numerator > (1<<53) || denominator > (1<<53) {
		n := new(big.Float).SetUint64(numerator)
		den := new(big.Float).SetUint64(denominator)
		result, _ := new(big.Float).Quo(n, den).Float64()
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
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd Bounded[T]) SetFromNormalized(normalized float64) (Bounded[T], error) {
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
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd Bounded[T]) SetFromNormalized32(normalized float32) (Bounded[T], error) {
	return bnd.SetFromNormalized(float64(normalized))
}

// Set sets the value of Bounded and automatically handles when the value exceeds the bounds.
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd Bounded[T]) Set(value T) (Bounded[T], error) {
	var err error

	if bnd.Clamp {
		if value > bnd.maximum {
			value = bnd.maximum
			err = fmt.Errorf("over")
		} else if value < bnd.minimum {
			value = bnd.minimum
			err = fmt.Errorf("under")
		}
	} else {
		if value > bnd.maximum {
			err = fmt.Errorf("over")
		} else if value < bnd.minimum {
			err = fmt.Errorf("under")
		}

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
	return bnd, err
}
