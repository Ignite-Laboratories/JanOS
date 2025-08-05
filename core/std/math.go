package std

import (
	"github.com/ignite-laboratories/core/std/num"
	"math"
	"math/rand"
	"reflect"
)

// ImplicitOverflow performs any implicit type overflow operations on num.ExtendedPrimitive types.
func ImplicitOverflow[T num.ExtendedInteger](value T) T {
	var zero T
	switch any(zero).(type) {
	case num.Crumb, num.Note, num.Nibble, num.Flake, num.Morsel, num.Shred, num.Run, num.Scale, num.Riff, num.Hook:
		overflow := MaxValue[T]() + 1
		return T(int(value) % int(overflow))
	case float32, float64:
		return value
	}
	return value
}

// MaxValue returns the maximum whole integer value of the provided type.
//
// NOTE: This will panic for non-integer types.
func MaxValue[T num.ExtendedPrimitive]() uint64 {
	switch any(T(0)).(type) {
	case num.Crumb:
		return 1<<2 - 1
	case num.Note:
		return 1<<3 - 1
	case num.Nibble:
		return 1<<4 - 1
	case num.Flake:
		return 1<<5 - 1
	case num.Morsel:
		return 1<<6 - 1
	case num.Shred:
		return 1<<7 - 1
	case num.Run:
		return 1<<10 - 1
	case num.Scale:
		return 1<<12 - 1
	case num.Riff:
		return 1<<24 - 1
	case num.Hook:
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

// Random returns a non-negative pseudo-random number of the provided type.
//
// If requesting a floating point type, the resulting number will be bounded
// in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [0, n] - where n is the maximum value of
// the provided type.
func Random[T num.ExtendedPrimitive]() T {
	switch any(T(0)).(type) {
	case num.Crumb:
		return T(RandomBounded[num.Crumb](0, num.Crumb(MaxValue[num.Crumb]())))
	case num.Note:
		return T(RandomBounded[num.Note](0, num.Note(MaxValue[num.Note]())))
	case num.Nibble:
		return T(RandomBounded[num.Nibble](0, num.Nibble(MaxValue[num.Nibble]())))
	case num.Flake:
		return T(RandomBounded[num.Flake](0, num.Flake(MaxValue[num.Flake]())))
	case num.Morsel:
		return T(RandomBounded[num.Morsel](0, num.Morsel(MaxValue[num.Morsel]())))
	case num.Shred:
		return T(RandomBounded[num.Shred](0, num.Shred(MaxValue[num.Shred]())))
	case num.Run:
		return T(RandomBounded[num.Run](0, num.Run(MaxValue[num.Run]())))
	case num.Scale:
		return T(RandomBounded[num.Scale](0, num.Scale(MaxValue[num.Scale]())))
	case num.Riff:
		return T(RandomBounded[num.Riff](0, num.Riff(MaxValue[num.Riff]())))
	case num.Hook:
		return T(RandomBounded[num.Hook](0, num.Hook(MaxValue[num.Hook]())))
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
type RandomNumberGeneratorFunc[T num.ExtendedPrimitive] func(a T, b T) T

var generators = make(map[reflect.Type]any)
var generatorsNil = make(map[reflect.Type]bool)

// DefineRandomGenerator sets the global random number generator for the provided type.
func DefineRandomGenerator[T num.ExtendedPrimitive](generator RandomNumberGeneratorFunc[T]) {
	// Get the type of T using a nil pointer to T
	t := reflect.TypeOf((*T)(nil)).Elem()
	generators[t] = generator
	generatorsNil[t] = generator == nil
}

// RandomBounded returns a pseudo-random number of the provided type bounded in the provided closed interval [a, b].
//
// NOTE: This uses a 0.01% chance to return exactly max.
func RandomBounded[T num.ExtendedPrimitive](a T, b T) T {
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
	case int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint,
		num.Crumb, num.Note, num.Nibble, num.Flake, num.Morsel, num.Shred, num.Run, num.Scale, num.Riff, num.Hook:
		range64 := uint64(b) - uint64(a)
		return T(uint64(a) + uint64(rand.Int63n(int64(range64+1))))
	default:
		panic("unsupported numeric type")
	}
}
