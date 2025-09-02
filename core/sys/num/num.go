package num

import (
	"core/sys/num/tiny"
	"fmt"
	"strconv"
)

const strNaN = "NaN"
const strInf = "Inf"

// Advanced represents any Numeric-compatible type, including Primitive types.  This interface acts as a bridge
// between traditional and raster-based computational arithmetic.  If you accept this interface level, you
// must be able to parse 'any' input into a logical number.
type Advanced interface {
	Primitive | complex64 | complex128 | tiny.Placeholder | tiny.Natural | tiny.Real
}

// Primitive represents any general primitive Numeric-compatible type.  These retain the standard mathematical operators,
// whereas Advanced numerics describe their own methods for standard arithmetic operations.
type Primitive interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		uintptr |

		float32 | float64
}

// Signed represents any Numeric-compatible type that's signed.
type Signed interface {
	tiny.Real |
		int | int8 | int16 | int32 | int64 |
		float32 | float64 | complex64 | complex128
}

// Integer represents any Numeric-compatible type that's an integer.
type Integer interface {
	tiny.Placeholder | tiny.Natural |
		int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		uintptr
}

// Float represents any Numeric-compatible type that supports floating-point preceision.
type Float interface {
	tiny.Real |
		float32 | float64 | complex64 | complex128
}

// Complex represents any complex number.
type Complex interface {
	complex64 | complex128
	// TODO: Add the tiny complex types when they are ready
}

// TypeAssert will type assert an 'any' type to its underlying provided Primitive.
//
// NOTE: This will panic if given a non-primitive value
func TypeAssert[TOut Primitive](value any) TOut {
	switch typed := value.(type) {
	case uint, uint8, uint16, uint32, uint64, uintptr,
		int, int8, int16, int32, int64,
		float32, float64,
		complex64, complex128:
		return typed.(TOut)
	}
	panic("cannot cast non-primitive types")
}

// ToString uses strconv to format a string representation of the number in base 10.
// The output will be a decimal value and not in notation form, using strconv's 'f' format whenever possible.
//
// NOTE: This will panic if provided a non Numeric-compatible type.
func ToString(value any) string {
	var out string
	switch typed := any(value).(type) {
	case tiny.Placeholder:
		return typed.String()
	case tiny.Natural:
		return typed.String()
	case tiny.Real:
		return typed.String()
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
		return fmt.Sprintf("%v", typed)
	default:
		panic("cannot compare non-numeric types")
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
