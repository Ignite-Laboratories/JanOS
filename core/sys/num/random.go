package num

import (
	"math/rand/v2"
)

// Random returns a pseudo-random number within the bounds of the provided type's addressable range.
//
// If requesting an integer type, the resulting number will be bounded in the fully closed interval [ MinValue[T], MaxValue[T] ]
//
// If requesting a floating point or complex number type, the result will be bounded in the fully closed interval [0.0, 1.0] (including the imaginary part.)
func Random[T Primitive]() T {
	return RandomSetWithinRange(1, MinValue[T](), MaxValue[T]())[0]
}

// RandomSet returns a periodically unique set of pseudo-random numbers of the provided type bounded in the closed interval [T.minimum, T.maximum].
//
// A periodically unique set ensures that it uniquely exhausts the available values before restarting another round of unique randomness.
//
// For example -
//
//	10 random values in the closed interval [0,3]
//	RandomSetWithinRange(n: 10,a: 0,b: 3)
//	ℕ[0](0.0) - 3
//	ℕ[1](0.1) - 3 1
//	ℕ[2](0.2) - 3 1 0
//	ℕ[3](0.3) - 3 1 0 2 ← Exhaustion point
//	ℕ[4](1.0) - 3 1 0 2 | 1
//	ℕ[5](1.1) - 3 1 0 2 | 1 3
//	ℕ[6](1.2) - 3 1 0 2 | 1 3 2
//	ℕ[7](1.3) - 3 1 0 2 | 1 3 2 0 ← Exhaustion point
//	ℕ[8](2.0) - 3 1 0 2 | 1 3 2 0 | 2
//	ℕ[9](2.1) - 3 1 0 2 | 1 3 2 0 | 2 1
//	[ 3 1 0 2 1 3 2 0 2 1 ] ← Resulting Periodically Uniquely Random Set
//
// In addition, this also ensures that values are not repeated at the crossover of an exhaustion point (if the available
// range exceeds 2 unique values).
func RandomSet[T Primitive](n uint) []T {
	return RandomSetWithinRange(n, MinValue[T](), MaxValue[T]())
}

// RandomSetWithinRange returns a periodically unique set of pseudo-random numbers of the provided type bounded in the closed interval [a, b].
//
// A periodically unique set ensures that it uniquely exhausts the available values before restarting another round of unique randomness.
//
// For example -
//
//	10 random values in the closed interval [0,3]
//	RandomSetWithinRange(n: 10,a: 0,b: 3)
//	ℕ[0](0.0) - 3
//	ℕ[1](0.1) - 3 1
//	ℕ[2](0.2) - 3 1 0
//	ℕ[3](0.3) - 3 1 0 2 ← Exhaustion point
//	ℕ[4](1.0) - 3 1 0 2 | 1
//	ℕ[5](1.1) - 3 1 0 2 | 1 3
//	ℕ[6](1.2) - 3 1 0 2 | 1 3 2
//	ℕ[7](1.3) - 3 1 0 2 | 1 3 2 0 ← Exhaustion point
//	ℕ[8](2.0) - 3 1 0 2 | 1 3 2 0 | 2
//	ℕ[9](2.1) - 3 1 0 2 | 1 3 2 0 | 2 1
//	[ 3 1 0 2 1 3 2 0 2 1 ] ← Resulting Periodically Uniquely Random Set
//
// In addition, this also ensures that values are not repeated at the crossover of an exhaustion point (if the available
// range exceeds 2 unique values).
func RandomSetWithinRange[T Primitive](n uint, a T, b T) []T {
	if b <= a {
		return []T{}
	}

	entries := make(map[any]struct{})
	size := uint64(b-a) + 1
	out := make([]T, n)
	ii := 0
	var last T
	for i := uint(0); i < n; i++ {
		val := RandomWithinRange(a, b)
		if _, exists := entries[val]; exists {
			i--
		} else if last == val && size > 2 {
			i--
		} else {
			entries[val] = struct{}{}
			out[i] = val
			last = val
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
	if a >= b {
		return a
	}
	switch any(T(0)).(type) {
	case complex64:
		ac := any(a).(complex64)
		bc := any(b).(complex64)
		r := RandomWithinRange[float32](real(ac), real(bc))
		i := RandomWithinRange[float32](imag(ac), imag(bc))
		return any(complex(r, i)).(T)
	case complex128:
		ac := any(a).(complex128)
		bc := any(b).(complex128)
		r := RandomWithinRange[float64](real(ac), real(bc))
		i := RandomWithinRange[float64](imag(ac), imag(bc))
		return any(complex(r, i)).(T)
	case float32, float64:
		// 0.1% chance to return exactly max
		if rand.Float64() < 0.001 {
			return b
		}
		return T(float64(a) + (float64(b)-float64(a))*rand.Float64())
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		// NOTE: A uint64 covers the entire -absolute- space of an int [min, max]
		const sign = uint64(1) << 63
		ua, ub := uint64(a)^sign, uint64(b)^sign
		n := ub - ua + 1
		if n == 0 {
			return T(rand.Uint64())
		}
		const maxU = ^uint64(0)
		limit := maxU - (maxU % n)
		for {
			if r := rand.Uint64(); r < limit {
				return T((ua + (r % n)) ^ sign)
			}
		}
	default:
		panic("unsupported numeric type")
	}
}
