package num

import (
	"core/enum/transcendental"
	"fmt"
	"math"
	"strings"
)

// A Realized number is a real number consisting of a Whole Natural part and a Fractional Natural part.
// This differs from a "real" number in that the periodic aspect of the fractional component is also tracked,
// allowing infinite reconstruction of a 'periodic' value - thus, the number technically exists in memory as
// a "realizable value" and only becomes a "real number" in the context of an arithmetic operation.
//
// This also carries into its ability to track irrational numbers.  Rather than storing the irrational result,
// an anonymous closure is made over the irrational operation which can be called on demand.  This allows the system
// to track an irrational result of any placeholder width, rather than tracking the formulae that generated the
// irrational condition.
type Realized struct {
	Negative       bool
	Whole          Natural
	Fractional     Natural
	PeriodicWidth  uint
	Transcendental transcendental.Transcendental
	Irrational     func(uint)
	base           uint16
}

// ParseRealized creates a new instance of a Realized number using the provided base-encoded source string - (see
// Realized.String for the standard encoding format) - if no source base is provided, this implies the source string
// is encoded in base₁₀.
//
// NOTE: Transcendentals -
//
// If you wish to parse a transcendental number, you can - but this simply short-circuits to the Transcendental methods.
// For instance, "-π" will yield you a negative instance of π from the cached constants.  For the STRING constants themselves,
// please see the transcendental.Transcendental enumeration package.
//
// NOTE: Irrationals -
//
// Irrational numbers can ONLY be identified during an arithmetic operation, as their infinitely repeating quality
// must be OBSERVED.  Thus, the '~' character is entirely ignored during parsing and the number is treated at face value.
func ParseRealized(source string, sourceBase ...uint16) Realized {
	b := uint16(10)
	if len(sourceBase) > 0 {
		b = sourceBase[0]
	}

	if len(source) == 0 {
		return Realized{
			Negative:   false,
			Whole:      ParseNatural("0"),
			Fractional: ParseNatural("0"),
			base:       b,
		}
	}

	negative := false
	if source[0] == '~' {
		source = source[1:]
	}
	if source[0] == '-' {
		source = source[1:]
		negative = true
	}

	t := transcendental.IsIdentifier(source)
	if t != transcendental.Non {
		switch t {
		case transcendental.Pi:
			r := Transcendental.Pi(b)
			r.Negative = negative
			return r
		case transcendental.E:
			r := Transcendental.E(b)
			r.Negative = negative
			return r
		}
	}

	var digits []string
	if b > 16 {
		digits = strings.Split(source, " ")
	} else {
		digits = strings.Split(source, "")
	}

	var wholeDigits []string
	var fractionalDigits []string

	isWholePart := true
	for i := 0; i < len(digits); i++ {
		if digits[i] == "." {
			isWholePart = false
		}

		if isWholePart {
			wholeDigits = append(wholeDigits, digits[i])
		} else {
			fractionalDigits = append(fractionalDigits, digits[i])
		}
	}

	var wholePart string
	var fractionalPart string
	if b > 16 {
		wholePart = strings.Join(wholeDigits, " ")
		fractionalPart = strings.Join(fractionalDigits, " ")
	} else {
		wholePart = strings.Join(wholeDigits, "")
		fractionalPart = strings.Join(fractionalDigits, "")
	}

	r := Realized{
		Negative:   negative,
		Whole:      ParseNatural(wholePart, b),
		Fractional: ParseNatural(fractionalPart, b),
		base:       b,
	}
	if t == transcendental.Non {
		r.Transcendental = Transcendental.Is(r)
	} else {
		r.Transcendental = t
	}
	return r
}

// NewRealized creates a new instance of a Realized number using the provided Primitive operand
//
// NOTE: ALL primitive operands in Go are base₁₀!  For parsing a real from a different base, please see ParseRealized
//
// NOTE: This will panic if provided a float32 or float64 value of 'Inf' or 'NaN'.
func NewRealized[T Primitive](operand T) Realized {
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

	return ParseRealized(ToString(operand))
}

func (r *Realized) Base() uint16 {
	return r.base
}

func (r *Realized) ChangeBase(base uint16) {
	r.Whole.ChangeBase(base)
	r.Fractional.ChangeBase(base)
	r.Transcendental = Transcendental.Is(*r)
}

// String prints a Realized in a legibly-encoded form using the below convention:
//
//	 "123"         ← An integer value
//	"-123"         ← An negative integer value
//
//	 "123.45"      ← A floating point value
//	"-123.45"      ← A negative floating point value
//
//	 "123.‾4"      ← A periodic value, broken with an "overscore"
//	"-123.‾4"      ← A negative periodic value, broken with an "overscore"
//
//	 "123.‾456"    ← A "wide" periodic value
//	"-123.‾456"    ← A "wide" negative periodic value
//
//	 "123.45‾678"  ← An mixed-repeating periodic value
//	"-123.45‾678"  ← A negative mixed-repeating periodic value
//
//	  "~1.7320508" ← An irrational value to atlas.PrecisionMinimum digits (default 7) [√3]
//	 "~-1.7320508" ← A negative irrational value to atlas.PrecisionMinimum digits (default 7) [-√3]
//
//	   "π"         ← Pi
//	   "ℯ"         ← Euler's number
//
// NOTE: A transcendental.Transcendental will appear automatically when the real number matches its value - and can be prefixed with a negative sign.
//
// Irrational values are always prefixed with a `~` character to indicate they are visibly truncated during String operations
// while retaining their atlas.Precision placeholder-width for calculation.  The minimum number of fractional placeholder positions
// 𝑡𝑖𝑛𝑦 will print out (to keep the output "reasonable" to read) is defined by atlas.PrecisionMinimum.
//
// For base₁₇ and above, all positions are printed with a space character between, and the digits are represented
// as two-digit hexadecimal values up to base₂₅₆ - which is 𝑡𝑖𝑛𝑦's limit.
//
//	"- 02 08 . 0B 00 10 06" ← "-42.54321" in base₁₇
//
// See Natural.String and Realized.String
func (r Realized) String() string {
	// 0 - Check if it's negative
	sign := ""
	if r.Negative {
		sign = "-"
	}

	// 1 - Check if it's a transcendental
	if r.Transcendental != transcendental.Non {
		if r.base > 16 {
			return sign + " " + string(r.Transcendental)
		}
		return sign + string(r.Transcendental)
	}

	// 2 - Check if it's irrational
	irrational := ""
	if r.Irrational != nil {
		irrational = "~"
	}

	// 3 - Set up a way to space out higher-order bases
	join := func(a ...string) string {
		if r.base > 16 {
			return strings.Join(a, " ")
		}
		return strings.Join(a, "")
	}

	frac := r.Fractional.String()

	// 4 - Check if it's fractional or periodic
	if len(frac) > 0 || r.PeriodicWidth > 0 {
		// 5 - Check if it's periodic
		if r.PeriodicWidth > 0 {
			// NOTE: The periodic component is an implied continuation of a single instance present in the fractional component.
			offset := len(r.Fractional.String()) - int(r.PeriodicWidth)
			f := r.Fractional.String()[:offset]
			p := r.Fractional.String()[offset:]

			// 6 - Print
			return join(irrational, sign, r.Whole.String(), ".", f, "‾", p)
		}
		return join(irrational, sign, r.Whole.String(), ".", r.Fractional.String())
	}
	return join(irrational, sign, r.Whole.String())
}

// Print pads the result of string to a "row" format for use in a calculation matrix.  This simply pads the left
// side of the whole part and the right side of the fractional part with '0's to width and puts a '0' or '1' in
// the first column to indicate if the value is signed (1).  The periodic "overscore" character is not printed.
// Instead, the number is synthesized to the desired placeholder precision and then rounded.
func (r Realized) Print(whole, fractional uint, precision ...uint) string {

}
