package num

import (
	"core/enum/direction/ordinal"
	"core/sys/pad"
	"core/sys/pad/scheme"
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
func ParseNatural(operand any, base ...uint16) Natural {
	b := PanicIfInvalidBase(base[0])
	op := ToString(FilterOperands(b, operand)[0])

	if len(op) == 0 {
		return Natural{NewMeasurement()}
	}

	if op[0] == '~' {
		op = op[1:]
	}
	if op[0] == '-' {
		op = op[1:]
	}

	var digits []string
	if b > 16 {
		digits = strings.Split(op, " ")
	} else {
		digits = strings.Split(op, "")
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

	bytes, _ := Base.StringToDigits(whole, b, 2)
	bits := make([]Bit, len(bytes))
	for i, digit := range bytes {
		bits[i] = Bit(digit)
	}

	return Natural{NewMeasurement(bits...)}
}

// Digits returns the natural's underlying digits in the provided base, or base₁₀ if omitted.
func (n Natural) Digits(base ...uint16) []byte {
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
