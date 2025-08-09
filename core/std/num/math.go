package num

import (
	"math"
	"math/rand"
	"reflect"
)

// IsSubByte checks if the provided type is a sub-byte num.Primitive type.
func IsSubByte[T Primitive]() bool {
	var zero T
	switch any(zero).(type) {
	case Crumb, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, Hook:
		return true
	}
	return false
}

// ImplicitOverflow performs any implicit type overflow operations on num.Primitive types.
func ImplicitOverflow[T Primitive](value T) T {
	var zero T
	switch any(zero).(type) {
	case Crumb, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, Hook:
		overflow := MaxValue[T]() + 1
		return T(int(value) % int(overflow))
	}
	return value
}

// IsSigned returns whether the provided type is a signed type or not.
func IsSigned[T Primitive]() bool {
	var zero T
	switch any(zero).(type) {
	case uint8, uint16, uint32, uint64, uint, Crumb, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, Hook:
		return false
	}
	return true
}

// MaxValue returns the maximum whole integer value of the provided type.
//
// NOTE: This will panic for non-integer types.
func MaxValue[T Primitive]() uint64 {
	switch any(T(0)).(type) {
	case Crumb:
		return 1<<2 - 1
	case Note:
		return 1<<3 - 1
	case Nibble:
		return 1<<4 - 1
	case Flake:
		return 1<<5 - 1
	case Morsel:
		return 1<<6 - 1
	case Shred:
		return 1<<7 - 1
	case Run:
		return 1<<10 - 1
	case Scale:
		return 1<<12 - 1
	case Riff:
		return 1<<24 - 1
	case Hook:
		return 1<<48 - 1
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
		panic("cannot provide the maximum value of a non-integer type.")
	}
}

// MinValue returns the minimum whole integer value of the provided type.
//
// NOTE: This will panic for non-integer types.
func MinValue[T Primitive]() int64 {
	switch any(T(0)).(type) {
	case Crumb, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, Hook,
		uint8, uint16, uint32, uint64, uint:
		return 0
	case int8:
		return math.MinInt8
	case int16:
		return math.MinInt16
	case int32:
		return math.MinInt32
	case int64:
		return math.MinInt64
	case int:
		return math.MinInt
	default:
		panic("cannot provide the minimum value of a non-integer type.")
	}
}

// Random returns a non-negative pseudo-random number of the provided type.
//
// If requesting a floating point type, the resulting number will be bounded
// in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [0, n] - where n is the maximum value of
// the provided type, including the implicit size of the extended primitives.
func Random[T Primitive]() T {
	switch any(T(0)).(type) {
	case Crumb:
		return T(RandomWithinRange[Crumb](0, Crumb(MaxValue[Crumb]())))
	case Note:
		return T(RandomWithinRange[Note](0, Note(MaxValue[Note]())))
	case Nibble:
		return T(RandomWithinRange[Nibble](0, Nibble(MaxValue[Nibble]())))
	case Flake:
		return T(RandomWithinRange[Flake](0, Flake(MaxValue[Flake]())))
	case Morsel:
		return T(RandomWithinRange[Morsel](0, Morsel(MaxValue[Morsel]())))
	case Shred:
		return T(RandomWithinRange[Shred](0, Shred(MaxValue[Shred]())))
	case Run:
		return T(RandomWithinRange[Run](0, Run(MaxValue[Run]())))
	case Scale:
		return T(RandomWithinRange[Scale](0, Scale(MaxValue[Scale]())))
	case Riff:
		return T(RandomWithinRange[Riff](0, Riff(MaxValue[Riff]())))
	case Hook:
		return T(RandomWithinRange[Hook](0, Hook(MaxValue[Hook]())))
	case float32:
		return T(RandomWithinRange[float32](0.0, 1.0))
	case float64:
		return T(RandomWithinRange[float64](0.0, 1.0))
	case int8:
		return T(RandomWithinRange[int8](math.MinInt8, math.MaxInt8))
	case uint8:
		return T(RandomWithinRange[uint8](0, math.MaxUint8))
	case int16:
		return T(RandomWithinRange[int16](math.MinInt16, math.MaxInt16))
	case uint16:
		return T(RandomWithinRange[uint16](0, math.MaxUint16))
	case int32:
		return T(RandomWithinRange[int32](math.MinInt32, math.MaxInt32))
	case uint32:
		return T(RandomWithinRange[uint32](0, math.MaxUint32))
	case int64:
		return T(RandomWithinRange[int64](math.MinInt64, math.MaxInt64))
	case int:
		return T(RandomWithinRange[int](math.MinInt, math.MaxInt))
	case uint64:
		return T(RandomWithinRange[uint64](0, math.MaxUint64))
	case uint:
		return T(RandomWithinRange[uint](0, math.MaxUint))
	default:
		panic("unsupported numeric type")
	}
}

// RandomNumberGeneratorFunc is a function that should return a random number of the defined type bounded within the closed interval of [a, b].
type RandomNumberGeneratorFunc[T Primitive] func(a T, b T) T

var generators = make(map[reflect.Type]any)
var generatorsNil = make(map[reflect.Type]bool)

// DefineRandomGenerator sets the global random number generator for the provided type.
func DefineRandomGenerator[T Primitive](generator RandomNumberGeneratorFunc[T]) {
	// Get the type of T using a nil pointer to T
	t := reflect.TypeOf((*T)(nil)).Elem()
	generators[t] = generator
	generatorsNil[t] = generator == nil
}

// RandomWithinRange returns a pseudo-random number of the provided type bounded in the provided closed interval [a, b].
//
// NOTE: This uses a 0.01% chance to return exactly max.
func RandomWithinRange[T Primitive](a T, b T) T {
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
	case Crumb, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, Hook:
		// These are implicitly sized uint types
		if a < 0 || b > T(MaxValue[T]()) {
			panic("cannot provide a random number exceeding the implicit bounds of the type.")
		}

		range64 := uint64(b) - uint64(a)
		return T(uint64(a) + uint64(rand.Int63n(int64(range64+1))))
	default:
		panic("unsupported numeric type")
	}
}
