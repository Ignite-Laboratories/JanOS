package num

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

// IsNaN uses 'any' to return whether the value is an IEEE 754 floating point 'NaN' value.
func IsNaN(a any) bool {
	switch typed := a.(type) {
	case float32:
		return math.IsNaN(float64(typed))
	case float64:
		return math.IsNaN(typed)
	default:
		return false
	}
}

// IsInf uses 'any' to return whether the value is 'Inf' and whether it's negative.
func IsInf(a any) (isInf bool, negative bool) {
	switch typed := a.(type) {
	case float32:
		return IsInf(float64(typed))
	case float64:
		if math.IsInf(typed, 1) {
			return true, false
		} else if math.IsInf(typed, -1) {
			return true, true
		}
		return false, false
	default:
		return false, false
	}
}

// ToString uses strconv to format a string representation of the number in base 10.
// The output will be a decimal value and not in notation form, using strconv's 'f' format whenever possible.
//
// NOTE: This will panic if provided a complex number or a non Primitive type.
func ToString(value any) string {
	var out string
	switch typed := any(value).(type) {
	case float32:
		out = strconv.FormatFloat(float64(typed), 'f', -1, 32)
	case float64:
		out = strconv.FormatFloat(float64(typed), 'f', -1, 64)
	case uint:
		out = strconv.FormatUint(uint64(typed), 10)
	case uint8:
		out = strconv.FormatUint(uint64(typed), 10)
	case uint16:
		out = strconv.FormatUint(uint64(typed), 10)
	case uint32:
		out = strconv.FormatUint(uint64(typed), 10)
	case uint64:
		out = strconv.FormatUint(typed, 10)
	case uintptr:
		out = strconv.FormatUint(uint64(typed), 10)
	case Bit:
		out = strconv.FormatUint(uint64(typed), 10)
	case Crumb:
		out = strconv.FormatUint(uint64(typed), 10)
	case Note:
		out = strconv.FormatUint(uint64(typed), 10)
	case Nibble:
		out = strconv.FormatUint(uint64(typed), 10)
	case Flake:
		out = strconv.FormatUint(uint64(typed), 10)
	case Morsel:
		out = strconv.FormatUint(uint64(typed), 10)
	case Shred:
		out = strconv.FormatUint(uint64(typed), 10)
	case Run:
		out = strconv.FormatUint(uint64(typed), 10)
	case Scale:
		out = strconv.FormatUint(uint64(typed), 10)
	case Riff:
		out = strconv.FormatUint(uint64(typed), 10)
	case Hook:
		out = strconv.FormatUint(uint64(typed), 10)
	case int:
		out = strconv.FormatInt(int64(typed), 10)
	case int8:
		out = strconv.FormatInt(int64(typed), 10)
	case int16:
		out = strconv.FormatInt(int64(typed), 10)
	case int32:
		out = strconv.FormatInt(int64(typed), 10)
	case int64:
		out = strconv.FormatInt(typed, 10)
	case complex64, complex128:
		// TODO: Revisit complex number comparison later
		panic("cannot currently compare the order of two complex numbers")
	default:
		panic("cannot compare non-primitive types")
	}

	// Ensure the data is never '.5' or '5.'
	// strconv should never produce this, but it's still a good check to perform
	if len(out) > 0 {
		if out[0] == '.' {
			out = "0" + out
		} else if out[len(out)-1] == '.' {
			out = out + "0"
		}
	}
	return out
}

// Cast will convert an 'any' type to the provided Primitive.
//
// NOTE: If the output type is unsigned and the input value is signed, this will underflow
//
// NOTE: This will panic if given a non-primitive value
func Cast[TOut Primitive](value any) TOut {
	switch typed := value.(type) {
	case float32:
		return ImplicitOverflow(TOut(typed))
	case float64:
		return ImplicitOverflow(TOut(typed))
	case uint:
		return ImplicitOverflow(TOut(typed))
	case uint8:
		return ImplicitOverflow(TOut(typed))
	case uint16:
		return ImplicitOverflow(TOut(typed))
	case uint32:
		return ImplicitOverflow(TOut(typed))
	case uint64:
		return ImplicitOverflow(TOut(typed))
	case uintptr:
		return ImplicitOverflow(TOut(typed))
	case Bit:
		return ImplicitOverflow(TOut(typed))
	case Crumb:
		return ImplicitOverflow(TOut(typed))
	case Note:
		return ImplicitOverflow(TOut(typed))
	case Nibble:
		return ImplicitOverflow(TOut(typed))
	case Flake:
		return ImplicitOverflow(TOut(typed))
	case Morsel:
		return ImplicitOverflow(TOut(typed))
	case Shred:
		return ImplicitOverflow(TOut(typed))
	case Run:
		return ImplicitOverflow(TOut(typed))
	case Scale:
		return ImplicitOverflow(TOut(typed))
	case Riff:
		return ImplicitOverflow(TOut(typed))
	case Hook:
		return ImplicitOverflow(TOut(typed))
	case int:
		return ImplicitOverflow(TOut(typed))
	case int8:
		return ImplicitOverflow(TOut(typed))
	case int16:
		return ImplicitOverflow(TOut(typed))
	case int32:
		return ImplicitOverflow(TOut(typed))
	case int64:
		return ImplicitOverflow(TOut(typed))
	default:
		panic("cannot cast non-primitive types")
	}
}

// IsPrimitive returns whether the provided values are all Primitive runtime objects.
//
// NOTE: If no values are provided, this still checks the type.
func IsPrimitive(values ...any) bool {
	if len(values) == 0 {
		return false
	}

	for _, v := range values {
		switch v.(type) {
		case int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64, uintptr,
			float32, float64,
			Bit, Crumb, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, Hook:
		default:
			return false
		}
	}
	return true
}

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
//
// NOTE: If no values are provided, this still checks the type.
func IsSubByte[T any](values ...T) bool {
	var zero T
	values = append(values, zero)
	allSub := true
	for _, v := range values {
		switch any(v).(type) {
		case Bit, Crumb, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, Hook:
		default:
			allSub = false
		}
	}
	return allSub
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

// IsInteger returns whether the provided Primitive type is an integer.
//
// NOTE: If no values are provided, this still checks the type.
func IsInteger[T any](values ...T) bool {
	var zero T
	values = append(values, zero)
	for _, v := range values {
		if IsFloat(v) || IsComplex(v) {
			return false
		}
	}
	return true
}

// IsComplex returns whether the provided Primitive type is a complex value.
//
// NOTE: If no values are provided, this still checks the type.
func IsComplex[T any](values ...T) bool {
	var zero T
	values = append(values, zero)
	for _, v := range values {
		switch any(v).(type) {
		case complex64, complex128:
		default:
			return false
		}
	}
	return true
}

// IsFloat returns whether the provided Primitive type is a floating point value.
//
// NOTE: If no values are provided, this still checks the type.
func IsFloat[T any](values ...T) bool {
	var zero T
	values = append(values, zero)
	for _, v := range values {
		switch any(v).(type) {
		case float32, float64:
		default:
			return false
		}
	}
	return true
}

// IsSigned returns whether the provided Primitive is a signed type or not.
//
// NOTE: If no values are provided, this still checks the type.
func IsSigned[T any](values ...T) bool {
	var zero T
	values = append(values, zero)
	allSigned := true
	for _, v := range values {
		switch any(v).(type) {
		case uint8, uint16, uint32, uint64, uint, Bit, Crumb, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, Hook:
			allSigned = false
		default:
		}
	}
	return allSigned
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

func boundaryValue[T Primitive](min bool) T {
	// NOTE: This may appear daunting, but I assure you there is a method to every mad decision in this switch branch.

	// Modify at your own risk =)
	var out uint64
	var zero T
	switch any(zero).(type) {
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
