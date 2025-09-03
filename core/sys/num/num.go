package num

import (
	"core/enum/direction/ordinal"
	"core/sys/pad"
	"core/sys/pad/scheme"
	"fmt"
	"strconv"
)

const strNaN = "NaN"
const strInf = "Inf"

// Advanced represents any Numeric-compatible type, including Primitive types.  This interface acts as a bridge
// between traditional and raster-based computational arithmetic.  If you accept this interface level, you
// must be able to parse 'any' input into a logical number.
type Advanced interface {
	Primitive | complex64 | complex128 | Natural | Real
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
	Real |
		int | int8 | int16 | int32 | int64 |
		float32 | float64 | complex64 | complex128
}

// Integer represents any Numeric-compatible type that's an integer.
type Integer interface {
	Natural |
		int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		uintptr
}

// Float represents any Numeric-compatible type that supports floating-point preceision.
type Float interface {
	Real |
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

func ToStringAligned(operands ...any) []string {
	if !IsNumeric(operands...) {
		panic("cannot align non Numeric-compatible types")
	}

	if IsComplex(operands...) {
		panic("cannot align complex numbers")
	}

	if IsNaN(operands...) {
		panic("cannot align NaN")
	}

	sign := 1
	whole := 2
	fractional := 4

	matrix := make([][]string, len(operands))
	var widestWhole int
	var widestFractional int

	for i, v := range operands {
		inf, _ := IsInf(v)
		if inf {
			panic("cannot align Inf")
		}
		str := ToString(v)
		parts := decimalPattern.FindStringSubmatch(str)
		if parts == nil {
			panic("unknown input type for alignment")
		}

		// NOTE: This ensures we don't get a '-0' in the output
		if allZerosPattern.MatchString(str) && len(parts[sign]) > 0 {
			parts[sign] = ""
		}
		if len(parts[sign]) == 0 {
			parts[sign] = " "
		}

		matrix[i] = parts
		if len(parts[whole]) > widestWhole {
			widestWhole = len(parts[whole])
		}
		if len(parts[fractional]) > widestFractional {
			widestFractional = len(parts[fractional])
		}
	}

	out := make([]string, len(operands))
	for i, v := range matrix {
		matrix[i][whole] = pad.String[rune](scheme.Tile, ordinal.Negative, uint(widestWhole), v[whole], "0")
		matrix[i][fractional] = pad.String[rune](scheme.Tile, ordinal.Positive, uint(widestFractional), v[fractional], "0")
		out[i] = matrix[i][sign] + matrix[i][whole] + "." + matrix[i][fractional]
	}
	return out
}

// ToString uses strconv to format a string representation of the number in base 10.
// The output will be a decimal value and not in notation form, using strconv's 'f' format whenever possible.
//
// NOTE: This will panic if provided a non Numeric-compatible type.
func ToString(value any) string {
	var out string
	switch typed := any(value).(type) {
	case Natural:
		return typed.String()
	case Real:
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
