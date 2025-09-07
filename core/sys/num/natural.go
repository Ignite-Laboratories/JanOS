package num

import (
	"core/enum/direction/ordinal"
	"core/sys/num/bases"
	"core/sys/pad"
	"core/sys/pad/scheme"
	"fmt"
	"strings"
)

type Natural struct {
	measurement Measurement
}

// ParseNatural - Creates a static natural number from an advanced operand, which includes 'string'.
//
// Primitive operands:
//
//	base₁₀ is implied and whatever base you provide is ignored.
//
// Advanced operands:
//
//	string - the operand must be encoded in the provided base value (or base₁₀ if omitted)
//	Natural - the natural is cloned and the base value is ignored (naturals do not store base)
//	Realized - the whole part of the realized number is captured and the base is ignored entirely
//	complex64 or complex128 - this will panic, as a natural number cannot describe a complex number
func ParseNatural[T Advanced](operand T, base ...uint16) Natural {

	switch raw := any(operand).(type) {
	case Natural, Realized, complex64, complex128:
		// CRITICAL: This does NOT type assert the string advanced types!!!
		// ALWAYS pass through the generic type for tiny types - never type assert and then mutate
		return parseNaturalAdvanced(operand, base...)
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		float32, float64:
		return parseNaturalString(ToString(raw)) // Primitive Go types are always base₁₀
	default:
		panic(fmt.Errorf("unknown type %T", raw))
	}
}

func parseNaturalAdvanced[T Advanced](operand T, base ...uint16) Natural {
	switch raw := any(operand).(type) {
	case string:
		bytes, _ := bases.StringToDigits(raw, PanicIfInvalidBase(base[0]), 2)
		bits := make([]Bit, len(bytes))
		for i, digit := range bytes {
			bits[i] = Bit(digit)
		}

		return Natural{NewMeasurement(bits...)}
	case Natural:
		return Natural{raw.measurement}
	case Realized:
		return Natural{raw.whole.measurement}
	default:
		panic(fmt.Errorf("cannot parse natural from type %T", raw))
	}
}

func parseNaturalString(source string, base ...uint16) Natural {
	b := PanicIfInvalidBase(base[0])

	if len(source) == 0 {
		return Natural{NewMeasurement()}
	}

	if source[0] == '~' {
		source = source[1:]
	}
	if source[0] == '-' {
		source = source[1:]
	}

	var digits []string
	if b > 16 {
		digits = strings.Split(source, " ")
	} else {
		digits = strings.Split(source, "")
	}

	for i := 0; i < len(digits); i++ {
		if digits[i] == "." {
			digits = digits[:i]
			break
		}
	}

	var whole string
	if b > 16 {
		whole = strings.Join(digits, " ")
	} else {
		whole = strings.Join(digits, "")
	}

	bytes, _ := bases.StringToDigits(whole, b, 2)
	bits := make([]Bit, len(bytes))
	for i, digit := range bytes {
		bits[i] = Bit(digit)
	}

	return Natural{NewMeasurement(bits...)}
}

// Digits returns the natural's underlying digits in the provided base, or base₁₀ if omitted.
func (n Natural) Digits(base ...uint16) []bases.Digit {
	return n.measurement.ToNaturalDigits(PanicIfInvalidBase(base...))
}

func (n Natural) String() string {
	return n.Print()
}

func (n Natural) Print(base ...uint16) string {
	str, _ := n.measurement.ToNaturalString(PanicIfInvalidBase(base[0]))
	return str
}

func (n Natural) Matrix(width uint, base ...uint16) string {
	str := n.Print(PanicIfInvalidBase(base[0]))
	return pad.String[rune](scheme.Tile, ordinal.Negative, width, str, "0")
}
