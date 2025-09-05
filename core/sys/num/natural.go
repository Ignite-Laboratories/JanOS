package num

import (
	"core/sys/num/bases"
	"fmt"
	"math"
	"strings"
)

// Natural represents a value belonging to the set of naturally countable numbers - or all positive whole numbers, including zero.
//
// To those who think zero shouldn't be included in the set of natural numbers, I present a counter-argument:
// base₁ has only one identifier, meaning it can only "represent" zero by -not- holding a value in an observable
// location.  Subsequently, all bases are built upon determining the size of a value through "identification" - in
// binary, through zeros or ones - in decimal, through the digits 0-9.
//
// Now here's where it gets tricky: a value cannot even EXIST until it is given a place to exist within, meaning its
// existence directly implies a void which has now been filled - an identifiable "zero" state.  In fact, the very first
// identifier of all higher order bases (zero) specifically identifies this state!  Counting, itself, comes from the act of observing
// the general relativistic -presence- of anything - fingers, digits, different length squiggles, feelings - meaning to exclude
// zero attempts to redefine the very fundamental definition of identification itself: it's PERFECTLY reasonable to -naturally-
// count -zero- hairs on a magnificently bald head!
//
//	tl;dr - to count naturally involves identification, including identifying NON-existence as a countable state!
//
// I should note this entire system hinges on one fundamental flaw - this container technically holds one additional value beyond
// the 'natural' number set: nil! Technically, until a number occupies a location, that space holds a 'nil' value in all bases
// above base 1, which might consider that to be 'zero'.  When factoring this trait in, I call it the "programmatic set" of
// numbers.  I can't stop you from setting your natural phrase to it - but I can empower you with awareness of it =)
type Natural struct {
	value Measurement
	base  uint16

	Width uint
}

// ParseNatural creates a new instance of a Natural number using the provided base-encoded source string - (see
// Natural.String for the standard encoding format) - if no source base is provided, this implies the source string
// is encoded in base₁₀.
//
// NOTE: A Natural can only hold a whole positive value - anything after a decimal place will be dropped entirely,
// the `~` and `-` characters will be dropped, and any other input than a base-encoded value will panic.
func ParseNatural(source string, sourceBase ...uint16) Natural {
	b := uint16(10)
	if len(sourceBase) > 0 {
		b = sourceBase[0]
	}

	if len(source) == 0 {
		return Natural{
			value: NewMeasurement(),
			base:  b,
			Width: 0,
		}
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

	binary, _ := bases.StringToDigits(whole, b, 2)
	bits := make([]Bit, len(binary))
	for i := 0; i < len(binary); i++ {
		bits[i] = Bit(binary[i])
	}

	return Natural{
		value: NewMeasurement(bits...),
		base:  b,
		Width: uint(len(binary)),
	}
}

// NewNatural creates a new instance of a Natural number using the provided Primitive operand
//
// NOTE: ALL primitive operands in Go are base₁₀!  For parsing a Natural from a different base, please see ParseNatural
//
// NOTE: This will panic if provided a float32 or float64 value of 'Inf' or 'NaN', otherwise it will trim the fractional component.
func NewNatural[T Primitive](operand T) Natural {
	switch raw := any(operand).(type) {
	case uint, uint8, uint16, uint32, uint64, uintptr,
		int, int8, int16, int32, int64:
	case float32:
		if math.IsInf(float64(raw), 0) {
			panic(fmt.Sprintf("cannot create a Realized from a Inf valued %T", raw))
		}
		if math.IsNaN(float64(raw)) {
			panic(fmt.Sprintf("cannot create a Realized from a NaN valued %T", raw))
		}
	case float64:
		if math.IsInf(raw, 0) {
			panic(fmt.Sprintf("cannot create a Realized from a Inf valued %T", raw))
		}
		if math.IsNaN(raw) {
			panic(fmt.Sprintf("cannot create a Realized from a NaN valued %T", raw))
		}
	default:
		panic(fmt.Sprintf("cannot create a Realized from type %T", raw))
	}

	return ParseNatural(ToString(operand))
}

func (n *Natural) Digits() []byte {
	return n.value.ToNaturalDigits(n.base)
}

func (n *Natural) ChangeBase(base uint16) {
	if base < 2 || base > 256 {
		panic(fmt.Errorf("invalid base '%d' - must be between 2 and 256", base))
	}

	n.base = base
}

func (n *Natural) Base() uint16 {
	return n.base
}

// String prints a Natural in a legibly-encoded form using the below convention:
//
//	"100" ← An base₁₆ or less unsigned integer value
//
//	"0B 00 10 06" ← A base₁₇ or higher unsigned integer value [54321 in base₁₀]
//
// For base₁₇ and above, all positions are printed with a space character between, and the digits are represented
// as two-digit hexadecimal values up to base₂₅₆ - which is 𝑡𝑖𝑛𝑦's limit.
//
// See Natural.String and Realized.String
func (n Natural) String() string {
	str, _ := n.value.ToNaturalString(n.base)
	return str
}
