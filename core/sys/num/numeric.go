package num

import (
	"fmt"
	"math"
	"math/big"
)

// Numeric represents a num.Advanced value bounded within the closed set [minimum, maximum].
// Additionally, all bounded types can be 'clamped' into the bounded range - meaning that
// they will not automatically overflow or underflow when they exceed the bounds.
//
// NOTE: All set operations return an error with the amount that the value overflowed or underflowed - otherwise nil.
type Numeric[T Advanced] struct {
	value       T
	minimum     T
	maximum     T
	initialized bool
	unbounded   bool

	Clamp bool
}

func (bnd *Numeric[T]) sanityCheck() {
	var zero T
	switch any(zero).(type) {
	case Natural, Realized, complex64, complex128:
		// CRITICAL: This does NOT type assert the string advanced types!!!
		// ALWAYS pass through the generic type for tiny types - never type assert and then mutate
		sanityCheckAdvanced(bnd)
	case int:
		sanityCheckPrimitive(any(bnd).(*Numeric[int]))
	case int8:
		sanityCheckPrimitive(any(bnd).(*Numeric[int8]))
	case int16:
		sanityCheckPrimitive(any(bnd).(*Numeric[int16]))
	case int32:
		sanityCheckPrimitive(any(bnd).(*Numeric[int32]))
	case int64:
		sanityCheckPrimitive(any(bnd).(*Numeric[int64]))
	case uint:
		sanityCheckPrimitive(any(bnd).(*Numeric[uint]))
	case uint8:
		sanityCheckPrimitive(any(bnd).(*Numeric[uint8]))
	case uint16:
		sanityCheckPrimitive(any(bnd).(*Numeric[uint16]))
	case uint32:
		sanityCheckPrimitive(any(bnd).(*Numeric[uint32]))
	case uint64:
		sanityCheckPrimitive(any(bnd).(*Numeric[uint64]))
	case uintptr:
		sanityCheckPrimitive(any(bnd).(*Numeric[uintptr]))
	case float32:
		sanityCheckPrimitive(any(bnd).(*Numeric[float32]))
	case float64:
		sanityCheckPrimitive(any(bnd).(*Numeric[float64]))
	default:
		panic(fmt.Errorf("unknown type %T", zero))
	}
}

func sanityCheckAdvanced[T Advanced](bnd *Numeric[T]) {
	// TODO: Implement this
}

func sanityCheckPrimitive[T Primitive](bnd *Numeric[T]) {
	if !bnd.initialized {
		if !bnd.unbounded {
			bnd.minimum = MinValue[T]()
			bnd.maximum = MaxValue[T]()
			var zero T
			if IsFloat(zero) {
				bnd.Clamp = true
			}
		}
		bnd.initialized = true
	}
}

// Unbound marks the bounded.Numeric as an 'unbounded' value - meaning set operations directly set the underlying value.
// To rebound your number, simply set them through any of the setter functions that include boundary values.
func (bnd *Numeric[T]) Unbound() {
	bnd.unbounded = true
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

/**

Boundary Access

*/

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
	var zero T
	switch any(zero).(type) {
	case Natural, Realized, complex64, complex128:
		// CRITICAL: This does NOT type assert the string advanced types!!!
		// ALWAYS pass through the generic type for tiny types - never type assert and then mutate
		return rangeAdvanced(bnd)
	case int:
		return rangePrimitive(any(bnd).(*Numeric[int]))
	case int8:
		return rangePrimitive(any(bnd).(*Numeric[int8]))
	case int16:
		return rangePrimitive(any(bnd).(*Numeric[int16]))
	case int32:
		return rangePrimitive(any(bnd).(*Numeric[int32]))
	case int64:
		return rangePrimitive(any(bnd).(*Numeric[int64]))
	case uint:
		return rangePrimitive(any(bnd).(*Numeric[uint]))
	case uint8:
		return rangePrimitive(any(bnd).(*Numeric[uint8]))
	case uint16:
		return rangePrimitive(any(bnd).(*Numeric[uint16]))
	case uint32:
		return rangePrimitive(any(bnd).(*Numeric[uint32]))
	case uint64:
		return rangePrimitive(any(bnd).(*Numeric[uint64]))
	case uintptr:
		return rangePrimitive(any(bnd).(*Numeric[uintptr]))
	case float32:
		return rangePrimitive(any(bnd).(*Numeric[float32]))
	case float64:
		return rangePrimitive(any(bnd).(*Numeric[float64]))
	default:
		panic(fmt.Errorf("unknown type %T", zero))
	}
}

func rangeAdvanced[T Advanced](bnd *Numeric[T]) uint64 {
	// TODO: implement this
	return 0
}

func rangePrimitive[T Primitive](bnd *Numeric[T]) uint64 {
	bnd.sanityCheck()

	var zero T
	if IsFloat(zero) {
		return math.MaxUint64
	}

	return uint64(bnd.maximum-bnd.minimum) + 1
}

/**

Basic Arithmetic

*/

// Random sets the value to a random number within the bounded range.  For unbounded numbers, this will use [ MinValue, MaxValue ]
//
// NOTE: tiny.Natural and tiny.Realized do not have an upper boundary - thus, this method will panic if unbounded.
func (bnd *Numeric[T]) Random() T {
	var zero T
	switch any(zero).(type) {
	case Natural, Realized, complex64, complex128:
		// CRITICAL: This does NOT type assert the string advanced types!!!
		// ALWAYS pass through the generic type for tiny types - never type assert and then mutate
		return randomAdvanced[T](bnd)
	case int:
		return any(randomPrimitive[int](any(bnd).(*Numeric[int]))).(T)
	case int8:
		return any(randomPrimitive[int8](any(bnd).(*Numeric[int8]))).(T)
	case int16:
		return any(randomPrimitive[int16](any(bnd).(*Numeric[int16]))).(T)
	case int32:
		return any(randomPrimitive[int32](any(bnd).(*Numeric[int32]))).(T)
	case int64:
		return any(randomPrimitive[int64](any(bnd).(*Numeric[int64]))).(T)
	case uint:
		return any(randomPrimitive[uint](any(bnd).(*Numeric[uint]))).(T)
	case uint8:
		return any(randomPrimitive[uint8](any(bnd).(*Numeric[uint8]))).(T)
	case uint16:
		return any(randomPrimitive[uint16](any(bnd).(*Numeric[uint16]))).(T)
	case uint32:
		return any(randomPrimitive[uint32](any(bnd).(*Numeric[uint32]))).(T)
	case uint64:
		return any(randomPrimitive[uint64](any(bnd).(*Numeric[uint64]))).(T)
	case uintptr:
		return any(randomPrimitive[uintptr](any(bnd).(*Numeric[uintptr]))).(T)
	case float32:
		return any(randomPrimitive[float32](any(bnd).(*Numeric[float32]))).(T)
	case float64:
		return any(randomPrimitive[float64](any(bnd).(*Numeric[float64]))).(T)
	default:
		panic(fmt.Errorf("unknown type %T", zero))
	}
}

func randomAdvanced[T Advanced](bnd *Numeric[T]) T {
	var zero T
	// TODO: Implement this
	return zero
}

func randomPrimitive[T Primitive](bnd *Numeric[T]) T {
	if bnd.unbounded {
		return Random[T]()
	}
	return RandomWithinRange(bnd.minimum, bnd.maximum)
}

// Increment adds 1 or the provided amount to the bound value.
//
// NOTE: If you provide a negative number, this will 'decrement'
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd *Numeric[T]) Increment(amount ...T) Breach {
	var zero T
	switch any(zero).(type) {
	case Natural, Realized, complex64, complex128:
		// CRITICAL: This does NOT type assert the string advanced types!!!
		// ALWAYS pass through the generic type for tiny types - never type assert and then mutate
		return incrementAdvanced(bnd, amount...)
	case int:
		return incrementPrimitive(any(bnd).(*Numeric[int]), any(amount).([]int)...)
	case int8:
		return incrementPrimitive(any(bnd).(*Numeric[int8]), any(amount).([]int8)...)
	case int16:
		return incrementPrimitive(any(bnd).(*Numeric[int16]), any(amount).([]int16)...)
	case int32:
		return incrementPrimitive(any(bnd).(*Numeric[int32]), any(amount).([]int32)...)
	case int64:
		return incrementPrimitive(any(bnd).(*Numeric[int64]), any(amount).([]int64)...)
	case uint:
		return incrementPrimitive(any(bnd).(*Numeric[uint]), any(amount).([]uint)...)
	case uint8:
		return incrementPrimitive(any(bnd).(*Numeric[uint8]), any(amount).([]uint8)...)
	case uint16:
		return incrementPrimitive(any(bnd).(*Numeric[uint16]), any(amount).([]uint16)...)
	case uint32:
		return incrementPrimitive(any(bnd).(*Numeric[uint32]), any(amount).([]uint32)...)
	case uint64:
		return incrementPrimitive(any(bnd).(*Numeric[uint64]), any(amount).([]uint64)...)
	case uintptr:
		return incrementPrimitive(any(bnd).(*Numeric[uintptr]), any(amount).([]uintptr)...)
	case float32:
		return incrementPrimitive(any(bnd).(*Numeric[float32]), any(amount).([]float32)...)
	case float64:
		return incrementPrimitive(any(bnd).(*Numeric[float64]), any(amount).([]float64)...)
	default:
		panic(fmt.Errorf("unknown type %T", zero))
	}
}

func incrementAdvanced[T Advanced](bnd *Numeric[T], amount ...T) Breach {
	// TODO: Implement this
	return ""
}

func incrementPrimitive[T Primitive](bnd *Numeric[T], amount ...T) Breach {
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
func (bnd *Numeric[T]) Decrement(amount ...T) Breach {
	var zero T
	switch any(zero).(type) {
	case Natural, Realized, complex64, complex128:
		// CRITICAL: This does NOT type assert the string advanced types!!!
		// ALWAYS pass through the generic type for tiny types - never type assert and then mutate
		return decrementAdvanced(bnd, amount...)
	case int:
		return decrementPrimitive(any(bnd).(*Numeric[int]), any(amount).([]int)...)
	case int8:
		return decrementPrimitive(any(bnd).(*Numeric[int8]), any(amount).([]int8)...)
	case int16:
		return decrementPrimitive(any(bnd).(*Numeric[int16]), any(amount).([]int16)...)
	case int32:
		return decrementPrimitive(any(bnd).(*Numeric[int32]), any(amount).([]int32)...)
	case int64:
		return decrementPrimitive(any(bnd).(*Numeric[int64]), any(amount).([]int64)...)
	case uint:
		return decrementPrimitive(any(bnd).(*Numeric[uint]), any(amount).([]uint)...)
	case uint8:
		return decrementPrimitive(any(bnd).(*Numeric[uint8]), any(amount).([]uint8)...)
	case uint16:
		return decrementPrimitive(any(bnd).(*Numeric[uint16]), any(amount).([]uint16)...)
	case uint32:
		return decrementPrimitive(any(bnd).(*Numeric[uint32]), any(amount).([]uint32)...)
	case uint64:
		return decrementPrimitive(any(bnd).(*Numeric[uint64]), any(amount).([]uint64)...)
	case uintptr:
		return decrementPrimitive(any(bnd).(*Numeric[uintptr]), any(amount).([]uintptr)...)
	case float32:
		return decrementPrimitive(any(bnd).(*Numeric[float32]), any(amount).([]float32)...)
	case float64:
		return decrementPrimitive(any(bnd).(*Numeric[float64]), any(amount).([]float64)...)
	default:
		panic(fmt.Errorf("unknown type %T", zero))
	}
}

func decrementAdvanced[T Advanced](bnd *Numeric[T], amount ...T) Breach {
	// TODO: Implement this
	return ""
}

func decrementPrimitive[T Primitive](bnd *Numeric[T], amount ...T) Breach {
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
func (bnd *Numeric[T]) AddOrSubtract(amount T) Breach {
	var zero T
	switch any(zero).(type) {
	case Natural, Realized, complex64, complex128:
		// CRITICAL: This does NOT type assert the string advanced types!!!
		// ALWAYS pass through the generic type for tiny types - never type assert and then mutate
		return addOrSubtractAdvanced(bnd, amount)
	case int:
		return addOrSubtractPrimitive(any(bnd).(*Numeric[int]), any(amount).(int))
	case int8:
		return addOrSubtractPrimitive(any(bnd).(*Numeric[int8]), any(amount).(int8))
	case int16:
		return addOrSubtractPrimitive(any(bnd).(*Numeric[int16]), any(amount).(int16))
	case int32:
		return addOrSubtractPrimitive(any(bnd).(*Numeric[int32]), any(amount).(int32))
	case int64:
		return addOrSubtractPrimitive(any(bnd).(*Numeric[int64]), any(amount).(int64))
	case uint:
		return addOrSubtractPrimitive(any(bnd).(*Numeric[uint]), any(amount).(uint))
	case uint8:
		return addOrSubtractPrimitive(any(bnd).(*Numeric[uint8]), any(amount).(uint8))
	case uint16:
		return addOrSubtractPrimitive(any(bnd).(*Numeric[uint16]), any(amount).(uint16))
	case uint32:
		return addOrSubtractPrimitive(any(bnd).(*Numeric[uint32]), any(amount).(uint32))
	case uint64:
		return addOrSubtractPrimitive(any(bnd).(*Numeric[uint64]), any(amount).(uint64))
	case uintptr:
		return addOrSubtractPrimitive(any(bnd).(*Numeric[uintptr]), any(amount).(uintptr))
	case float32:
		return addOrSubtractPrimitive(any(bnd).(*Numeric[float32]), any(amount).(float32))
	case float64:
		return addOrSubtractPrimitive(any(bnd).(*Numeric[float64]), any(amount).(float64))
	default:
		panic(fmt.Errorf("unknown type %T", zero))
	}
}

func addOrSubtractAdvanced[T Advanced](bnd *Numeric[T], amount T) Breach {
	// TODO: Implement this
	return ""
}

func addOrSubtractPrimitive[T Primitive](bnd *Numeric[T], amount T) Breach {
	bnd.sanityCheck()

	if Compare(amount, 0) == -1 {
		amount = -amount
		return bnd.Decrement(T(amount))
	}
	return bnd.Increment(T(amount))
}

/**

SetAll

*/

// SetAll sets the value and boundaries all in one operation, preventing multiple calls to Set().
//
// NOTE: The boundary parameters are evaluated to ensure the lower bound is always the 'minimum'
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd *Numeric[T]) SetAll(value, minimum, maximum T, clamp ...bool) Breach {
	bnd.sanityCheck()

	bnd.unbounded = false

	c := len(clamp) > 0 && clamp[0]
	if Compare(minimum, maximum) == 1 {
		minimum, maximum = maximum, minimum
	}
	bnd.minimum = minimum
	bnd.maximum = maximum
	bnd.Clamp = c
	return bnd.Set(value)
}

/**

SetBoundaries

*/

// SetBoundariesUsingAny casts the provided values to T and then calls SetBoundariesAny.
//
// NOTE: In addition to the error from SetBoundariesAny, this will yield an error if the provided values are not of type T.
func (bnd *Numeric[T]) SetBoundariesUsingAny(minimum, maximum any) (Breach, error) {
	var zero T
	if _, ok := minimum.(T); !ok {
		return "", fmt.Errorf("value 'a' was not of type %T", zero)
	}
	if _, ok := maximum.(T); !ok {
		return "", fmt.Errorf("value 'b' was not of type %T", zero)
	}

	bnd.unbounded = false

	return bnd.SetBoundaries(minimum.(T), minimum.(T)), nil
}

// SetBoundaries sets the boundaries before calling Set(current value).
//
// NOTE: The boundary parameters are evaluated to ensure the lower bound is always the 'minimum'
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd *Numeric[T]) SetBoundaries(minimum, maximum T) Breach {
	bnd.sanityCheck()

	bnd.unbounded = false

	if Compare(minimum, maximum) == 1 {
		minimum, maximum = maximum, minimum
	}
	bnd.minimum = minimum
	bnd.maximum = maximum
	return bnd.Set(bnd.value)
}

/**

Normalize

*/

// Normalize converts the Numeric value to a float64 unit vector in the interval [0.0, 1.0],
// by linearly mapping the value from its bounded interval's [minimum, maximum]. A value equal
// to minimum maps to 0.0, a value equal to maximum maps to 1.0, and values in between
// are linearly interpolated.
func (bnd *Numeric[T]) Normalize() (float64, error) {
	var zero T
	switch any(zero).(type) {
	case Natural, Realized, complex64, complex128:
		// CRITICAL: This does NOT type assert the string advanced types!!!
		// ALWAYS pass through the generic type for tiny types - never type assert and then mutate
		return normalizeAdvanced(bnd)
	case int:
		return normalizePrimitive(any(bnd).(*Numeric[int]))
	case int8:
		return normalizePrimitive(any(bnd).(*Numeric[int8]))
	case int16:
		return normalizePrimitive(any(bnd).(*Numeric[int16]))
	case int32:
		return normalizePrimitive(any(bnd).(*Numeric[int32]))
	case int64:
		return normalizePrimitive(any(bnd).(*Numeric[int64]))
	case uint:
		return normalizePrimitive(any(bnd).(*Numeric[uint]))
	case uint8:
		return normalizePrimitive(any(bnd).(*Numeric[uint8]))
	case uint16:
		return normalizePrimitive(any(bnd).(*Numeric[uint16]))
	case uint32:
		return normalizePrimitive(any(bnd).(*Numeric[uint32]))
	case uint64:
		return normalizePrimitive(any(bnd).(*Numeric[uint64]))
	case uintptr:
		return normalizePrimitive(any(bnd).(*Numeric[uintptr]))
	case float32:
		return normalizePrimitive(any(bnd).(*Numeric[float32]))
	case float64:
		return normalizePrimitive(any(bnd).(*Numeric[float64]))
	default:
		panic(fmt.Errorf("unknown type %T", zero))
	}
}

func normalizeAdvanced[T Advanced](bnd *Numeric[T]) (float64, error) {
	// TODO: Implement this
	return 0, nil
}

func normalizePrimitive[T Primitive](bnd *Numeric[T]) (float64, error) {
	bnd.sanityCheck()

	if bnd.unbounded {
		return 0, fmt.Errorf("cannot normalize an unbounded space")
	}

	numerator := uint64(bnd.value - bnd.minimum)
	denominator := uint64(bnd.maximum - bnd.minimum)

	// If the bounded range cannot be represented in a float64, bump to big.Float
	// NOTE: float64 only maintains full precision up to 2⁵³

	if numerator > (1<<53) || denominator > (1<<53) {
		n := new(big.Float).SetUint64(numerator)
		den := new(big.Float).SetUint64(denominator)
		result, _ := new(big.Float).Quo(n, den).Float64()
		return result, nil
	}

	return float64(numerator) / float64(denominator), nil
}

// Normalize32 converts the Numeric value to a float32 unit vector in the range [0.0, 1.0],
// where the bounded minimum maps to 0.0 and the bounded maximum maps to 1.0.
func (bnd *Numeric[T]) Normalize32() (float32, error) {
	bnd.sanityCheck()

	norm, err := bnd.Normalize()
	return float32(norm), err
}

/**

SetFromNormalized

*/

// SetFromNormalized32 sets the bounded value using a float32 unit vector from the [0.0, 1.0]
// range, where 0.0 maps to the bounded minimum and 1.0 maps to the bounded maximum.
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd *Numeric[T]) SetFromNormalized32(normalized float32) (Breach, error) {
	bnd.sanityCheck()

	return bnd.SetFromNormalized(float64(normalized))
}

// SetFromNormalized sets the bounded value using a float64 unit vector from the [0.0, 1.0]
// range, where 0.0 maps to the bounded minimum and 1.0 maps to the bounded maximum.
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd *Numeric[T]) SetFromNormalized(normalized float64) (Breach, error) {
	var zero T
	switch any(zero).(type) {
	case Natural, Realized, complex64, complex128:
		// CRITICAL: This does NOT type assert the string advanced types!!!
		// ALWAYS pass through the generic type for tiny types - never type assert and then mutate
		return setFromNormalizedAdvanced(bnd, normalized)
	case int:
		return setFromNormalizedPrimitive(any(bnd).(*Numeric[int]), normalized)
	case int8:
		return setFromNormalizedPrimitive(any(bnd).(*Numeric[int8]), normalized)
	case int16:
		return setFromNormalizedPrimitive(any(bnd).(*Numeric[int16]), normalized)
	case int32:
		return setFromNormalizedPrimitive(any(bnd).(*Numeric[int32]), normalized)
	case int64:
		return setFromNormalizedPrimitive(any(bnd).(*Numeric[int64]), normalized)
	case uint:
		return setFromNormalizedPrimitive(any(bnd).(*Numeric[uint]), normalized)
	case uint8:
		return setFromNormalizedPrimitive(any(bnd).(*Numeric[uint8]), normalized)
	case uint16:
		return setFromNormalizedPrimitive(any(bnd).(*Numeric[uint16]), normalized)
	case uint32:
		return setFromNormalizedPrimitive(any(bnd).(*Numeric[uint32]), normalized)
	case uint64:
		return setFromNormalizedPrimitive(any(bnd).(*Numeric[uint64]), normalized)
	case uintptr:
		return setFromNormalizedPrimitive(any(bnd).(*Numeric[uintptr]), normalized)
	case float32:
		return setFromNormalizedPrimitive(any(bnd).(*Numeric[float32]), normalized)
	case float64:
		return setFromNormalizedPrimitive(any(bnd).(*Numeric[float64]), normalized)
	default:
		panic(fmt.Errorf("unknown type %T", normalized))
	}
}

func setFromNormalizedAdvanced[T Advanced](bnd *Numeric[T], normalized float64) (Breach, error) {
	// TODO: Implement this
	return "", nil
}

func setFromNormalizedPrimitive[T Primitive](bnd *Numeric[T], normalized float64) (Breach, error) {
	bnd.sanityCheck()

	if bnd.unbounded {
		return "", fmt.Errorf("cannot normalize an unbounded space")
	}

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
		return bnd.Set(T(val)), nil
	}

	scaled := normalized * float64(distance)
	return bnd.Set(T(scaled) + bnd.minimum), nil
}

/**

Set

*/

// SetUsingAny casts the provided value to T and then calls Set.
//
// NOTE: In addition to the error from Set, this will yield an error if the provided value is not of type T.
func (bnd *Numeric[T]) SetUsingAny(value any) (Breach, error) {
	if _, ok := value.(T); !ok {
		var zero T
		return "", fmt.Errorf("value was not of type %T", zero)
	}
	return bnd.Set(value.(T)), nil
}

// Set sets the bounded value.  If clamp is true, the value is "clamped" to the closed set [minimum, maximum] - otherwise,
// the value overflows and underflows.  The amount that the value over and underflows is returned as a Breach string, which is
// empty if the value did not breach the boundaries.  If the number is unbounded, this simply sets the value and returns an
// empty breach.
func (bnd *Numeric[T]) Set(value T) Breach {
	switch raw := any(value).(type) {
	case Natural, Realized, complex64, complex128:
		// CRITICAL: This does NOT type assert the string advanced types!!!
		// ALWAYS pass through the generic type for tiny types - never type assert and then mutate
		return setAdvanced(bnd, value)
	case int:
		return setPrimitive(any(bnd).(*Numeric[int]), raw)
	case int8:
		return setPrimitive(any(bnd).(*Numeric[int8]), raw)
	case int16:
		return setPrimitive(any(bnd).(*Numeric[int16]), raw)
	case int32:
		return setPrimitive(any(bnd).(*Numeric[int32]), raw)
	case int64:
		return setPrimitive(any(bnd).(*Numeric[int64]), raw)
	case uint:
		return setPrimitive(any(bnd).(*Numeric[uint]), raw)
	case uint8:
		return setPrimitive(any(bnd).(*Numeric[uint8]), raw)
	case uint16:
		return setPrimitive(any(bnd).(*Numeric[uint16]), raw)
	case uint32:
		return setPrimitive(any(bnd).(*Numeric[uint32]), raw)
	case uint64:
		return setPrimitive(any(bnd).(*Numeric[uint64]), raw)
	case uintptr:
		return setPrimitive(any(bnd).(*Numeric[uintptr]), raw)
	case float32:
		return setPrimitive(any(bnd).(*Numeric[float32]), raw)
	case float64:
		return setPrimitive(any(bnd).(*Numeric[float64]), raw)
	default:
		panic(fmt.Errorf("unknown type %T", raw))
	}
}

func setAdvanced[T Advanced](bnd *Numeric[T], value T) Breach {
	// TODO: Implement this
	return ""
}

func setPrimitive[T Primitive](bnd *Numeric[T], value T) Breach {
	bnd.sanityCheck()

	if bnd.unbounded {
		bnd.value = value
		return ""
	}

	var breach Breach

	if bnd.Clamp {
		if value > bnd.maximum {
			overflow := value - bnd.maximum
			value = bnd.maximum
			breach = Breach(ToString(overflow))
		} else if value < bnd.minimum {
			underflow := bnd.minimum - value
			value = bnd.minimum
			breach = Breach("-" + ToString(underflow))
		}
	} else {
		if value > bnd.maximum {
			overflow := value - bnd.maximum
			breach = Breach(ToString(overflow))

			var zero T
			if IsFloat(zero) {
				r := bnd.maximum - bnd.minimum
				for value > bnd.maximum && overflow > 0 {
					overflow = T(math.Mod(float64(overflow), float64(r)))
					if overflow == 0 {
						value = bnd.minimum
					} else {
						value = bnd.minimum + overflow
					}
				}
				bnd.value = value
				return breach
			}
		} else if value < bnd.minimum {
			underflow := value - bnd.minimum
			breach = Breach(ToString(underflow))

			var zero T
			if IsFloat(zero) {
				r := bnd.maximum - bnd.minimum
				for value < bnd.minimum && underflow < 0 {
					underflow = T(math.Mod(float64(underflow), float64(r)))
					if underflow == 0 {
						value = bnd.maximum
					} else {
						value = bnd.maximum + underflow
					}
				}
				bnd.value = value
				return breach
			}
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
			case uint64, uint, uintptr:
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
	return breach
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
