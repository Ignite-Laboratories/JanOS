package num

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
)

// String prints the provided value using the formatting placeholder %f for floating point values and %d for integer values.
func String[T Primitive](v T) string {
	var zero T
	switch any(zero).(type) {
	case float32, float64:
		return fmt.Sprintf("%f", v)
	default:
		return fmt.Sprintf("%d", v)
	}
}

// IsSubByte checks if the provided type is a sub-byte num.Primitive type.
func IsSubByte[T Primitive]() bool {
	var zero T
	switch any(zero).(type) {
	case Bit, Crumb, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, Hook:
		return true
	}
	return false
}

// ImplicitOverflow performs any implicit type overflow operations on num.Primitive types.
func ImplicitOverflow[T Primitive](value T) T {
	var zero T
	switch any(zero).(type) {
	case Bit, Crumb, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, Hook:
		overflow := MaxValue[T]() + 1
		return T(int(value) % int(overflow))
	}
	return value
}

// IsFloat returns whether the provided type is a floating point value or not.
func IsFloat[T Primitive]() bool {
	var zero T
	switch any(zero).(type) {
	case float32, float64:
		return true
	}
	return false
}

// IsSigned returns whether the provided type is a signed type or not.
func IsSigned[T Primitive]() bool {
	var zero T
	switch any(zero).(type) {
	case uint8, uint16, uint32, uint64, uint, Bit, Crumb, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, Hook:
		return false
	}
	return true
}

// MaxValue returns an instance of T set to the maximum value of the provided Primitive type.
func MaxValue[T Primitive]() T {
	return boundaryValue[T](false)
}

// MinValue returns an instance of T set to the minimum value of the provided Primitive type.
//
// For unsigned types, this returns 0.
//
// For signed integer types, this returns math.MinIntX()
//
// For floating point types, this returns -math.MaxFloatXX()
func MinValue[T Primitive]() T {
	return boundaryValue[T](true)
}

// Random returns a non-negative pseudo-random number of the provided type.
//
// If requesting a floating point type, the resulting number will be bounded
// in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [0, MaxValue[T]]
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
		return T(uint64(a) + uint64(rand.Int63n(int64(range64+1))))
	default:
		panic("unsupported numeric type")
	}
}

func boundaryValue[T Primitive](min bool) T {
	// NOTE: This may appear daunting, but I assure you there is a method to every mad decision in this switch branch.

	// Modify at your own risk =)
	var out uint64
	switch any(T(0)).(type) {
	case Bit:
		if min {
			return 0
		}
		out = 1
	case Crumb:
		if min {
			return 0
		}
		out = 1<<2 - 1
	case Note:
		if min {
			return 0
		}
		out = 1<<3 - 1
	case Nibble:
		if min {
			return 0
		}
		out = 1<<4 - 1
	case Flake:
		if min {
			return 0
		}
		out = 1<<5 - 1
	case Morsel:
		if min {
			return 0
		}
		out = 1<<6 - 1
	case Shred:
		if min {
			return 0
		}
		out = 1<<7 - 1
	case Run:
		if min {
			return 0
		}
		out = 1<<10 - 1
	case Scale:
		if min {
			return 0
		}
		out = 1<<12 - 1
	case Riff:
		if min {
			return 0
		}
		out = 1<<24 - 1
	case Hook:
		if min {
			return 0
		}
		out = 1<<48 - 1
	case int8:
		if min {
			v := math.MinInt8
			return T(v)
		}
		out = math.MaxInt8
	case uint8:
		if min {
			return 0
		}
		out = math.MaxUint8
	case int16:
		if min {
			v := math.MinInt16
			return T(v)
		}
		out = math.MaxInt16
	case uint16:
		if min {
			return 0
		}
		out = math.MaxUint16
	case int32:
		if min {
			v := math.MinInt32
			return T(v)
		}
		out = math.MaxInt32
	case uint32:
		if min {
			return 0
		}
		out = math.MaxUint32
	case int64:
		if min {
			v := math.MinInt64
			return T(v)
		}
		out = math.MaxInt64
	case uint64:
		if min {
			return 0
		}
		out = math.MaxUint64
	case int:
		if min {
			v := math.MinInt
			return T(v)
		}
		out = math.MaxInt
	case uint:
		if min {
			return 0
		}
		out = math.MaxUint
	case float64:
		if min {
			v := -math.MaxFloat64
			return T(v)
		}
		v := math.MaxFloat64
		return T(v)
	case float32:
		if min {
			v := -math.MaxFloat32
			return T(v)
		}
		v := math.MaxFloat32
		return T(v)
	default:
		switch reflect.TypeOf(T(0)).Kind() {
		case reflect.Float64:
			if min {
				v := -math.MaxFloat64
				return T(v)
			}
			v := math.MaxFloat64
			return T(v)
		case reflect.Float32:
			if min {
				v := -math.MaxFloat32
				return T(v)
			}
			v := math.MaxFloat32
			return T(v)
		case reflect.Int8:
			if min {
				v := math.MinInt8
				return T(v)
			}
			out = math.MaxInt8
		case reflect.Int16:
			if min {
				v := math.MinInt16
				return T(v)
			}
			out = math.MaxInt16
		case reflect.Int32:
			if min {
				v := math.MinInt32
				return T(v)
			}
			out = math.MaxInt32
		case reflect.Int64:
			if min {
				v := math.MinInt64
				return T(v)
			}
			out = math.MaxInt64
		case reflect.Int:
			if min {
				v := math.MinInt
				return T(v)
			}
			out = math.MaxInt
		case reflect.Uint8:
			if min {
				return 0
			}
			out = math.MaxUint8
		case reflect.Uint16:
			if min {
				return 0
			}
			out = math.MaxUint16
		case reflect.Uint32:
			if min {
				return 0
			}
			out = math.MaxUint32
		case reflect.Uint64:
			if min {
				return 0
			}
			out = math.MaxUint64
		case reflect.Uint:
			if min {
				return 0
			}
			out = math.MaxUint
		default:
			var zero T
			panic(fmt.Errorf("unknown type %T", zero))
		}
	}
	return T(out)
}
