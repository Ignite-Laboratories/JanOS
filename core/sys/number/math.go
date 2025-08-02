package number

import (
	"math"
	"math/rand"
	"reflect"
)

// MaxValue returns the maximum whole integer value of the provided type.
//
// NOTE: This will return 0 for unsupported types, such as floats.
func MaxValue[T Numeric]() uint64 {
	switch any(T(0)).(type) {
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
		return 0
	}
}

// Random returns a non-negative pseudo-random number of the provided type.
//
// If requesting a floating point type, the resulting number will be bounded
// in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [0, n] - where n is the maximum value of
// the provided type.
func Random[T Numeric]() T {
	switch any(T(0)).(type) {
	case float32:
		return T(RandomBounded[float32](0.0, 1.0))
	case float64:
		return T(RandomBounded[float64](0.0, 1.0))
	case int8:
		return T(RandomBounded[int8](math.MinInt8, math.MaxInt8))
	case uint8:
		return T(RandomBounded[uint8](0, math.MaxUint8))
	case int16:
		return T(RandomBounded[int16](math.MinInt16, math.MaxInt16))
	case uint16:
		return T(RandomBounded[uint16](0, math.MaxUint16))
	case int32:
		return T(RandomBounded[int32](math.MinInt32, math.MaxInt32))
	case uint32:
		return T(RandomBounded[uint32](0, math.MaxUint32))
	case int64:
		return T(RandomBounded[int64](math.MinInt64, math.MaxInt64))
	case int:
		return T(RandomBounded[int](math.MinInt, math.MaxInt))
	case uint64:
		return T(RandomBounded[uint64](0, math.MaxUint64))
	case uint:
		return T(RandomBounded[uint](0, math.MaxUint))
	default:
		panic("unsupported numeric type")
	}
}

// RandomNumberGeneratorFunc is a function that should return a random number of the defined type bounded within the closed interval of [a, b].
type RandomNumberGeneratorFunc[T Numeric] func(a T, b T) T

var generators = make(map[reflect.Type]any)
var generatorsNil = make(map[reflect.Type]bool)

// DefineRandomGenerator sets the global random number generator for the provided type.
func DefineRandomGenerator[T Numeric](generator RandomNumberGeneratorFunc[T]) {
	// Get the type of T using a nil pointer to T
	t := reflect.TypeOf((*T)(nil)).Elem()
	generators[t] = generator
	generatorsNil[t] = generator == nil
}

// RandomBounded returns a pseudo-random number of the provided type bounded in the provided closed interval [a, b].
//
// NOTE: This uses a 0.01% chance to return exactly max.
func RandomBounded[T Numeric](a T, b T) T {
	// Get the type of T
	t := reflect.TypeOf((*T)(nil)).Elem()

	// Check if we have a generator for this type
	if gen, ok := generators[t]; ok {
		if !generatorsNil[t] {
			return gen.(RandomNumberGeneratorFunc[T])(a, b)
		}
	}

	if a >= b {
		return a
	}
	switch any(T(0)).(type) {
	case float32, float64:
		// 0.1% chance to return exactly max
		if rand.Float64() < 0.001 {
			return b
		}
		return T(float64(a) + (float64(b)-float64(a))*rand.Float64())
	case int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint:
		range64 := uint64(b) - uint64(a)
		return T(uint64(a) + uint64(rand.Int63n(int64(range64+1))))
	default:
		panic("unsupported numeric type")
	}
}
