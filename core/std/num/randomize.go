package num

import (
	"math"
	"math/rand/v2"
	"reflect"
)

// Random returns a pseudo-random number within the bounds of the provided type's addressable range.
//
// If requesting a floating point type, the resulting number will be bounded in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [MinValue[T], MaxValue[T]]
func Random[T Primitive]() T {
	switch any(T(0)).(type) {
	case Bit:
		return T(RandomWithinRange[Bit](0, 1))
	case Crumb:
		return T(RandomWithinRange[Crumb](0, MaxValue[Crumb]()))
	case Note:
		return T(RandomWithinRange[Note](0, MaxValue[Note]()))
	case Nibble:
		return T(RandomWithinRange[Nibble](0, MaxValue[Nibble]()))
	case Flake:
		return T(RandomWithinRange[Flake](0, MaxValue[Flake]()))
	case Morsel:
		return T(RandomWithinRange[Morsel](0, MaxValue[Morsel]()))
	case Shred:
		return T(RandomWithinRange[Shred](0, MaxValue[Shred]()))
	case Run:
		return T(RandomWithinRange[Run](0, MaxValue[Run]()))
	case Scale:
		return T(RandomWithinRange[Scale](0, MaxValue[Scale]()))
	case Riff:
		return T(RandomWithinRange[Riff](0, MaxValue[Riff]()))
	case Hook:
		return T(RandomWithinRange[Hook](0, MaxValue[Hook]()))
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

// RandomSetWithinType returns a periodically unique set of pseudo-random numbers of the provided type bounded in the closed interval [T.minimum, T.maximum].
//
// A periodically unique set ensures that it uniquely exhausts the available values before restarting another round of unique randomness.
//
// For example -
//
//	  10 random values in the closed interval [0,3]
//		 RandomSetWithinRange(n: 10,a: 0,b: 3)
//		 ℕ[0](0.0) - 3
//		 ℕ[1](0.1) - 3 1
//		 ℕ[2](0.2) - 3 1 0
//		 ℕ[3](0.3) - 3 1 0 2 ← Exhaustion point
//		 ℕ[4](1.0) - 3 1 0 2 | 1
//		 ℕ[5](1.1) - 3 1 0 2 | 1 3
//		 ℕ[6](1.2) - 3 1 0 2 | 1 3 2
//		 ℕ[7](1.3) - 3 1 0 2 | 1 3 2 0 ← Exhaustion point
//		 ℕ[8](2.0) - 3 1 0 2 | 1 3 2 0 | 0
//		 ℕ[9](2.1) - 3 1 0 2 | 1 3 2 0 | 0 1
//	  [ 3 1 0 2 1 3 2 0 0 1 ] ← Resulting Periodically Uniquely Random Set
//
// NOTE: This uses a 0.01% chance to return exactly max for each entry.
func RandomSetWithinType[T Primitive](n uint) []T {
	return RandomSetWithinRange(n, MinValue[T](), MaxValue[T]())
}

// RandomSetWithinRange returns a periodically unique set of pseudo-random numbers of the provided type bounded in the closed interval [a, b].
//
// A periodically unique set ensures that it uniquely exhausts the available values before restarting another round of unique randomness.
//
// For example -
//
//	  10 random values in the closed interval [0,3]
//		 RandomSetWithinRange(n: 10,a: 0,b: 3)
//		 ℕ[0](0.0) - 3
//		 ℕ[1](0.1) - 3 1
//		 ℕ[2](0.2) - 3 1 0
//		 ℕ[3](0.3) - 3 1 0 2 ← Exhaustion point
//		 ℕ[4](1.0) - 3 1 0 2 | 1
//		 ℕ[5](1.1) - 3 1 0 2 | 1 3
//		 ℕ[6](1.2) - 3 1 0 2 | 1 3 2
//		 ℕ[7](1.3) - 3 1 0 2 | 1 3 2 0 ← Exhaustion point
//		 ℕ[8](2.0) - 3 1 0 2 | 1 3 2 0 | 0
//		 ℕ[9](2.1) - 3 1 0 2 | 1 3 2 0 | 0 1
//	  [ 3 1 0 2 1 3 2 0 0 1 ] ← Resulting Periodically Uniquely Random Set
//
// NOTE: This uses a 0.01% chance to return exactly max for each entry.
func RandomSetWithinRange[T Primitive](n uint, a T, b T) []T {
	if b <= a {
		return []T{}
	}

	entries := make(map[any]struct{})
	size := uint64(b-a) + 1
	out := make([]T, n)
	ii := 0
	for i := uint(0); i < n; i++ {
		val := RandomWithinRange(a, b)
		if _, exists := entries[val]; exists {
			i--
		} else {
			entries[val] = struct{}{}
			out[i] = val
			ii++
		}
		if uint64(ii) == size {
			entries = make(map[any]struct{})
			ii = 0
		}
	}
	return out
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
		const sign = uint64(1) << 63
		ua, ub := (uint64(a) ^ sign), (uint64(b) ^ sign)
		n := ub - ua + 1
		if n == 0 { // full int64 domain
			return T(rand.Uint64())
		}
		const maxU = ^uint64(0)
		limit := maxU - (maxU % n) // rejection threshold to avoid modulo bias
		for {
			if r := rand.Uint64(); r < limit {
				return T((ua + (r % n)) ^ sign)
			}
		}
	case Bit, Crumb, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, Hook:
		// These are implicitly sized uint types
		maximum := MaxValue[T]()
		if a < 0 {
			a = 0
		}
		if b > maximum {
			b = maximum
		}

		range64 := uint64(b) - uint64(a)
		rand.Uint64()
		return T(uint64(a) + uint64(rand.Int64N(int64(range64+1))))
	default:
		panic("unsupported numeric type")
	}
}
