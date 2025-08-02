package math

import (
	"math"
	"math/rand"
	"reflect"
)

// MaxIntegerValue returns the maximum whole integer value of the provided type.
func MaxIntegerValue[T Numeric]() uint64 {
	switch any(T(0)).(type) {
	case float64:
		panic("unsupported numeric type - float64's maximum cannot be represented by a uint64")
	case float32:
		panic("unsupported numeric type - float32's maximum cannot be represented by a uint64")
	case int8:
		return math.MaxInt8
	case uint8:
		return uint64(math.MaxUint8)
	case int16:
		return uint64(math.MaxInt16)
	case uint16:
		return uint64(math.MaxUint16)
	case int32:
		return uint64(math.MaxInt32)
	case uint32:
		return uint64(math.MaxUint32)
	case int64:
		return uint64(math.MaxInt64)
	case uint64:
		return math.MaxUint64
	case int:
		return math.MaxInt
	case uint:
		return math.MaxUint
	default:
		panic("unsupported numeric type")
	}
}

// NormalizeToFloat64 returns a normalized value of the provided type in the range [0.0, 1.0].
func NormalizeToFloat64[T Numeric](value T) float64 {
	return float64(value) / float64(MaxIntegerValue[T]())
}

// NormalizeToFloat32 returns a normalized value of the provided type in the range [0.0, 1.0].
func NormalizeToFloat32[T Numeric](value T) float32 {
	return float32(value) / float32(MaxIntegerValue[T]())
}

// ScaleFloat64ToType returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleFloat64ToType[T Numeric](value float64) T {
	if value < 0.0 || value > 1.0 {
		panic("value must be in range [0.0, 1.0]")
	}
	return T(value * float64(MaxIntegerValue[T]()))
}

// ScaleFloat32ToType returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleFloat32ToType[T Numeric](value float32) T {
	if value < 0.0 || value > 1.0 {
		panic("value must be in range [0.0, 1.0]")
	}
	return T(value * float32(MaxIntegerValue[T]()))
}

// NumberRange creates a std.Tuple with A representing the start and B representing the stop.
func NumberRange[T Numeric](start T, stop T) Tuple[T] {
	return Tuple[T]{
		A: start,
		B: stop,
	}
}

// RandomNumber returns a non-negative pseudo-random number of the provided type.
//
// If requesting a floating point type, the resulting number will be bounded
// in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [0, n] - where n is the maximum value of
// the provided type.
func RandomNumber[T Numeric]() T {
	switch any(T(0)).(type) {
	case float32:
		return T(RandomNumberRange[float32](NumberRange[float32](0.0, 1.0)))
	case float64:
		return T(RandomNumberRange[float64](NumberRange[float64](0.0, 1.0)))
	case int8:
		return T(RandomNumberRange[int8](NumberRange[int8](math.MinInt8, math.MaxInt8)))
	case uint8:
		return T(RandomNumberRange[uint8](NumberRange[uint8](0, math.MaxUint8)))
	case int16:
		return T(RandomNumberRange[int16](NumberRange[int16](math.MinInt16, math.MaxInt16)))
	case uint16:
		return T(RandomNumberRange[uint16](NumberRange[uint16](0, math.MaxUint16)))
	case int32:
		return T(RandomNumberRange[int32](NumberRange[int32](math.MinInt32, math.MaxInt32)))
	case uint32:
		return T(RandomNumberRange[uint32](NumberRange[uint32](0, math.MaxUint32)))
	case int64:
		return T(RandomNumberRange[int64](NumberRange[int64](math.MinInt64, math.MaxInt64)))
	case int:
		return T(RandomNumberRange[int](NumberRange[int](math.MinInt, math.MaxInt)))
	case uint64:
		return T(RandomNumberRange[uint64](NumberRange[uint64](0, math.MaxUint64)))
	case uint:
		return T(RandomNumberRange[uint](NumberRange[uint](0, math.MaxUint)))
	default:
		panic("unsupported numeric type")
	}
}

// RandomNumberGeneratorFunc is a function that returns a random number of the provided type.
type RandomNumberGeneratorFunc[T Numeric] func(Tuple[T]) T

var generators = make(map[reflect.Type]any)
var generatorsNil = make(map[reflect.Type]bool)

// SetRandomNumberGenerator sets the global random number generator for the provided type.
func SetRandomNumberGenerator[T Numeric](generator RandomNumberGeneratorFunc[T]) {
	// Get the type of T using a nil pointer to T
	t := reflect.TypeOf((*T)(nil)).Elem()
	generators[t] = generator
	generatorsNil[t] = generator == nil
}

// RandomNumberRange returns a pseudo-random number of the provided type bounded in the closed interval [min, max].
//
// NOTE: This uses a 0.01% chance to return exactly max.
func RandomNumberRange[T Numeric](r Tuple[T]) T {
	// Get the type of T
	t := reflect.TypeOf((*T)(nil)).Elem()

	// Check if we have a generator for this type
	if gen, ok := generators[t]; ok {
		if !generatorsNil[t] {
			return gen.(RandomNumberGeneratorFunc[T])(r)
		}
	}

	if r.A >= r.B {
		return r.A
	}
	switch any(T(0)).(type) {
	case float32, float64:
		// 0.1% chance to return exactly max
		if rand.Float64() < 0.001 {
			return r.B
		}
		return T(float64(r.A) + (float64(r.B)-float64(r.A))*rand.Float64())
	case int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint:
		range64 := uint64(r.B) - uint64(r.A)
		return T(uint64(r.A) + uint64(rand.Int63n(int64(range64+1))))
	default:
		panic("unsupported numeric type")
	}
}
