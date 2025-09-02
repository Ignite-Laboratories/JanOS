package num

import (
	"core/sys/num/tiny"
	"math"
)

// IsNumeric returns whether the provided values are all Numeric-compatible types.
func IsNumeric(values ...any) bool {
	if len(values) == 0 {
		return false
	}

	for _, v := range values {
		switch v.(type) {
		case tiny.Placeholder, tiny.Natural, tiny.Real,
			uint, uint8, uint16, uint32, uint64, uintptr,
			int, int8, int16, int32, int64,
			float32, float64,
			complex64, complex128:
		default:
			return false
		}
	}
	return true
}

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

// IsInteger returns whether the provided Primitive type is an integer type.
//
// NOTE: tiny.Real is not an integer type, though it can hold integer values.
func IsInteger(values ...any) bool {
	if len(values) == 0 {
		return false
	}
	for _, v := range values {
		switch v.(type) {
		case tiny.Placeholder, tiny.Natural,
			int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64, uintptr:
		default:
			return false
		}
	}
	return true
}

// IsComplex returns whether the provided Primitive type is a complex value.
func IsComplex(values ...any) bool {
	if len(values) == 0 {
		return false
	}
	for _, v := range values {
		switch v.(type) {
		case complex64, complex128:
		default:
			return false
		}
	}
	return true
}

// IsFloat returns whether the provided Primitive type is a floating point value.
//
// NOTE: tiny.Real is a floating point type.
func IsFloat(values ...any) bool {
	if len(values) == 0 {
		return false
	}
	for _, v := range values {
		switch v.(type) {
		case tiny.Real, float32, float64:
		default:
			return false
		}
	}
	return true
}

// IsSigned returns whether the provided Primitive is a signable type or not.
//
// NOTE: tiny.Real is a signable type.
func IsSigned(values ...any) bool {
	if len(values) == 0 {
		return false
	}
	for _, v := range values {
		switch v.(type) {
		case tiny.Real, int, int8, int16, int32, int64, float32, float64, complex64, complex128:
		default:
			return false
		}
	}
	return true
}
