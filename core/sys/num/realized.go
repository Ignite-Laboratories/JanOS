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

// NewRealized creates a new instance of a Realized number using the provided operand, then converts it to the provided base.
//
// NOTE: If no base is provided, it implies 'base-10'
func NewRealized(operand any, base ...uint16) Realized {
	b := uint16(10)
	if len(base) > 0 {
		b = base[0]
		if b < 2 {
			panic(fmt.Sprintf("cannot create Realized from base %d", base))
		}
	}

	switch raw := operand.(type) {
	case string:
	case Natural, Realized,
		uint, uint8, uint16, uint32, uint64, uintptr,
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

	str := ToString(operand)
	parts := decimalPattern.FindStringSubmatch(str)
	if parts == nil {
		panic("cannot create Realized: unknown input type")
	}

	negative := len(parts[1]) > 0 && parts[1][0] == '-'
	fractional := NewNatural(parts[4], b)
	whole := NewNatural(parts[2], b)

	return Realized{
		Negative:   negative,
		Whole:      whole,
		Fractional: fractional,
		base:       b,
	}
}

func (r *Realized) Base() uint16 {
	return r.base
}

func (r *Realized) ChangeBase(base uint16) {
	r.Whole.ChangeBase(base)
	r.Fractional.ChangeBase(base)
	r.evaluate()
}

// evaluate looks at the currently stored real number and determines if it's transcendental.
func (r *Realized) evaluate() {
	r.Transcendental = IsTranscendental(*r)
}

// String prints a Realized in a legibly-encoded form using the below convention:
//
//	 "123"         ← An integer value
//	"-123"         ← An negative integer value
//
//	 "123.45"      ← A floating point value
//	"-123.45"      ← A negative floating point value
//
//	 "123.45‾678"  ← A periodic value
//	"-123.45‾678"  ← A negative periodic value
//
//	  "~1.7320508" ← An irrational value to atlas.PrecisionMinimum digits (default 7) [√3]
//	 "~-1.7320508" ← A negative irrational value to atlas.PrecisionMinimum digits (default 7) [-√3]
//	   "ℯ"         ← Euler's number
//	   "π"         ← Pi
//
// NOTE: A transcendental.Transcendental will appear automatically when the real number matches its value
//
// Irrational values are always prefixed with a `~` character to indicate they are visibly truncated during String operations
// while retaining their atlas.Precision placeholder-width for calculation.  The minimum number of fractional placeholder positions
// 𝑡𝑖𝑛𝑦 will print out (to keep the output "reasonable" to read) is defined by atlas.PrecisionMinimum.
//
// For base₁₇ and above, all positions are printed with a space character between and the digits are represented
// as two-digit hexadecimal values up to base₂₅₆ - which is 𝑡𝑖𝑛𝑦's limit.
//
//	"~ - 02 0A . 09 06 0F 06 02 0E 0F 0D 01 0D 07 05 0C 10 00 04 09 02 04 0D" ← "~-44.5533" in base₁₇
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
