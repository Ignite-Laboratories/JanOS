package num

import (
	"fmt"
	"math"
)

// A Real number is one of arbitrary placeholder width.  It holds a sign, plus three Natural numbers - Whole,
// Fractional, and Periodic.  The fractional component holds one instance of the periodic pattern, while the
// periodic value retains which part of the fractional component was periodic.  This allows a Real number to
// synthesize out to any placeholder width "on-demand" without loss of precision.
//
// Periodicity is evaluated during every arithmetic operation up to the global atlas.Precision placeholder width.
type Real struct {
	Negative   bool
	Whole      Natural
	Fractional Natural
	Periodic   Natural
	base       byte
}

func NewReal(operand any, base ...byte) Real {
	b := byte(10)
	if len(base) > 0 {
		b = base[0]
		if b < 2 {
			panic(fmt.Sprintf("cannot create Real from base %d", base))
		}
	}

	switch raw := operand.(type) {
	case Natural, Real,
		uint, uint8, uint16, uint32, uint64, uintptr,
		int, int8, int16, int32, int64:
	case float32:
		if math.IsInf(float64(raw), 0) {
			panic(fmt.Sprintf("cannot create a Real from a Inf valued %T", raw))
		}
		if math.IsNaN(float64(raw)) {
			panic(fmt.Sprintf("cannot create a Real from a NaN valued %T", raw))
		}
	case float64:
		if math.IsInf(raw, 0) {
			panic(fmt.Sprintf("cannot create a Real from a Inf valued %T", raw))
		}
		if math.IsNaN(raw) {
			panic(fmt.Sprintf("cannot create a Real from a NaN valued %T", raw))
		}
	default:
		panic(fmt.Sprintf("cannot create a Real from type %T", raw))
	}

	str := ToString(operand)
	parts := decimalPattern.FindStringSubmatch(str)
	if parts == nil {
		panic("cannot create Real: unknown input type")
	}

	negative := len(parts[1]) > 0 && parts[1][0] == '-'
	whole := newNatural(parts[2], b)
	fractional := newNatural(parts[4], b)

	return Real{
		Negative:   negative,
		Whole:      whole,
		Fractional: fractional,
		base:       b,
	}
}

func (r Real) Base() byte {
	return r.base
}

func (r Real) String() string {
	n := ""
	if r.Negative {
		n = "-"
	}
	if len(r.Fractional.digits) > 0 || len(r.Periodic.digits) > 0 {
		if len(r.Periodic.digits) > 0 {
			// The periodic component is an implied continuation of a single instance present in the fractional component.
			p := r.Periodic.String()
			f := r.Fractional.String()[:len(r.Fractional.String())-len(p)]

			return fmt.Sprintf("%s%s.%s[%s]", n, r.Whole, f, p)
		}
		return fmt.Sprintf("%s%s.%s", n, r.Whole, r.Fractional)
	}
	return fmt.Sprintf("%s%s", n, r.Whole)
}
