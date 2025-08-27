package bounded

import (
	"core/sys/num"
	"errors"
	"fmt"
	"math"
	"math/big"
)

// Numeric represents a bounded numeric value bound within the closed set [minimum, maximum].
// Additionally, all bounded types can be 'clamped' into the bounded range - meaning that
// they will not automatically overflow or underflow when they exceed the bounds.
//
// NOTE: All set operations return an error with the amount that the value overflowed or underflowed - otherwise nil.
type Numeric[T num.Primitive] struct {
	value       T
	minimum     T
	maximum     T
	initialized bool
	Clamp       bool
}

// NewNumber creates a new instance of Numeric[T].
//
// NOTE: While you can call this directly, the convention is to use the 'std/bounded' package.
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func NewNumber[T num.Primitive](value, minimum, maximum T, clamp ...bool) (Numeric[T], error) {
	c := len(clamp) > 0 && clamp[0]
	b := Numeric[T]{
		value:       T(0),
		minimum:     minimum,
		maximum:     maximum,
		initialized: true,
		Clamp:       c,
	}
	err := b.Set(value)
	return b, err
}

func (bnd *Numeric[T]) sanityCheck() {
	if !bnd.initialized {
		bnd.minimum = num.MinValue[T]()
		bnd.maximum = num.MaxValue[T]()
		bnd.initialized = true
	}
}

// Value returns the currently held Numeric value.
func (bnd *Numeric[T]) Value() T {
	bnd.sanityCheck()

	return bnd.value
}

// ValueAsAny returns the currently held Numeric value as an 'any' type, satisfying INumeric.
func (bnd *Numeric[T]) ValueAsAny() any {
	bnd.sanityCheck()

	return any(bnd.Value())
}

// Minimum returns the current minimum boundary.
func (bnd *Numeric[T]) Minimum() T {
	bnd.sanityCheck()

	return bnd.minimum
}

// MinimumAsAny returns the current minimum boundary as an 'any' type, satisfying INumeric.
func (bnd *Numeric[T]) MinimumAsAny() any {
	bnd.sanityCheck()

	return any(bnd.Minimum())
}

// Maximum returns the current maximum boundary.
func (bnd *Numeric[T]) Maximum() T {
	bnd.sanityCheck()

	return bnd.maximum
}

// MaximumAsAny returns the current maximum boundary as an 'any' type, satisfying INumeric.
func (bnd *Numeric[T]) MaximumAsAny() any {
	bnd.sanityCheck()

	return any(bnd.Maximum())
}

// Range returns the bounded range -
//
//	(maximum - minimum) + 1
func (bnd *Numeric[T]) Range() uint64 {
	bnd.sanityCheck()

	if num.IsFloat[T]() {
		return math.MaxUint64
	}

	return uint64(bnd.maximum-bnd.minimum) + 1
}

// Increment adds 1 or the provided amount to the bound value.
//
// NOTE: If you provide a negative number, this will 'decrement'
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd *Numeric[T]) Increment(amount ...T) error {
	bnd.sanityCheck()

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
func (bnd *Numeric[T]) Decrement(amount ...T) error {
	bnd.sanityCheck()

	i := T(1)
	if len(amount) > 0 {
		i = amount[0]
	}
	return bnd.Set(bnd.value - i)
}

// AddOrSubtract adds or subtracts the provided amount to the bound value.
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd *Numeric[T]) AddOrSubtract(amount T) error {
	bnd.sanityCheck()

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
func (bnd *Numeric[T]) SetAll(value, minimum, maximum T, clamp ...bool) error {
	bnd.sanityCheck()

	c := len(clamp) > 0 && clamp[0]
	if minimum > maximum {
		minimum, maximum = maximum, minimum
	}
	bnd.minimum = minimum
	bnd.maximum = maximum
	bnd.Clamp = c
	return bnd.Set(value)
}

// SetBoundariesFromType sets the boundaries to the implied limits of the bounded type before calling Set(current value).
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd *Numeric[T]) SetBoundariesFromType() error {
	bnd.sanityCheck()

	bnd.minimum = 0
	bnd.maximum = num.MaxValue[T]()
	return bnd.Set(bnd.value)
}

// SetBoundaries sets the boundaries before calling Set(current value).
//
// NOTE: The boundary parameters are evaluated to ensure the lower bound is always the 'minimum'
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd *Numeric[T]) SetBoundaries(minimum, maximum T) error {
	bnd.sanityCheck()

	if minimum > maximum {
		minimum, maximum = maximum, minimum
	}
	bnd.minimum = minimum
	bnd.maximum = maximum
	return bnd.Set(bnd.value)
}

// SetBoundariesUsingAny casts the provided values to T and then calls SetBoundariesAny.
//
// NOTE: In addition to the error from SetBoundariesAny, this will yield an error if the provided values are not of type T.
func (bnd *Numeric[T]) SetBoundariesUsingAny(minimum, maximum any) error {
	if _, ok := minimum.(T); !ok {
		return fmt.Errorf("value 'a' was not of type %T", T(0))
	}
	if _, ok := maximum.(T); !ok {
		return fmt.Errorf("value 'b' was not of type %T", T(0))
	}
	return bnd.SetBoundaries(minimum.(T), minimum.(T))
}

// Normalize converts the Numeric value to a float64 unit vector in the interval [0.0, 1.0],
// by linearly mapping the value from its bounded interval's [minimum, maximum]. A value equal
// to minimum maps to 0.0, a value equal to maximum maps to 1.0, and values in between
// are linearly interpolated.
func (bnd *Numeric[T]) Normalize() float64 {
	bnd.sanityCheck()

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

// Normalize32 converts the Numeric value to a float32 unit vector in the range [0.0, 1.0],
// where the bounded minimum maps to 0.0 and the bounded maximum maps to 1.0.
func (bnd *Numeric[T]) Normalize32() float32 {
	bnd.sanityCheck()

	return float32(bnd.Normalize())
}

// SetFromNormalized sets the bounded value using a float64 unit vector from the [0.0, 1.0]
// range, where 0.0 maps to the bounded minimum and 1.0 maps to the bounded maximum.
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd *Numeric[T]) SetFromNormalized(normalized float64) error {
	bnd.sanityCheck()

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
func (bnd *Numeric[T]) SetFromNormalized32(normalized float32) error {
	bnd.sanityCheck()

	return bnd.SetFromNormalized(float64(normalized))
}

// SetUsingAny casts the provided value to T and then calls Set.
//
// NOTE: In addition to the error from Set, this will yield an error if the provided value is not of type T.
func (bnd *Numeric[T]) SetUsingAny(value any) error {
	if _, ok := value.(T); !ok {
		return fmt.Errorf("value was not of type %T", T(0))
	}
	return bnd.Set(value.(T))
}

// Set sets the bounded value.  If clamp is true, the value is "clamped" to the closed set [minimum, maximum] - otherwise,
// the value overflows and underflows.  The amount that the value over and underflows is returned as an error, returning nil
// when the assigned value was within the closed interval.
func (bnd *Numeric[T]) Set(value T) error {
	bnd.sanityCheck()

	var err error

	if bnd.Clamp {
		if value > bnd.maximum {
			overflow := value - bnd.maximum
			value = bnd.maximum
			err = errors.New(num.String[T](overflow))
		} else if value < bnd.minimum {
			underflow := bnd.minimum - value
			value = bnd.minimum
			err = errors.New("-" + num.String[T](underflow))
		}
	} else {
		if value > bnd.maximum {
			overflow := value - bnd.maximum
			err = errors.New(num.String[T](overflow))
		} else if value < bnd.minimum {
			underflow := bnd.minimum - value
			err = errors.New("-" + num.String[T](underflow))
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

		mod := T(diff)
		if distance > 0 {
			mod = T(diff % distance)
		}
		value = bnd.minimum + mod
	}

	bnd.value = value
	return err
}

// String returns the value as a numeric string.
func (bnd Numeric[T]) String() string {
	bnd.sanityCheck()

	var zero T
	switch any(zero).(type) {
	case float32, float64:
		return fmt.Sprintf("%f", bnd.Value())
	default:
		return fmt.Sprintf("%d", bnd.Value())
	}
}
