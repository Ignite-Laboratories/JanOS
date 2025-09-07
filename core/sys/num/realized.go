package num

import (
	"core/enum/transcendental"
	"core/sys/atlas"
	"core/sys/num/bases"
	"fmt"
	"math"
	"strings"
	"sync"
)

// A Realized number is one generated through the execution of a neural pathway.  see.RealizedNumbers
//
// ParseRealized - Creates a static realized number.
//
// NewRealized - Creates a dynamic realized number.
type Realized struct {
	gate sync.Mutex

	transcendental transcendental.Number

	irrational    bool
	Negative      bool
	whole         Natural
	fractional    Natural
	periodicWidth uint

	revelation func(Realization) Realization
	potential  func() bool

	precision *uint
	base      uint16

	created bool
}

// ParseRealized - Creates a static realized number from an advanced operand, which includes 'string'.
//
// Primitive operands:
//
//	base₁₀ is implied and whatever base you provide is ignored.
//
// Advanced operands:
//
//	string - the operand must be encoded in the provided base value (or base₁₀ if omitted)
//	Natural - sets the whole part and base of the realized number (or base₁₀ if omitted)
//	Realized - sets the whole and fractional part and changes it to the provided base (or base₁₀ if omitted)
//	complex64 or complex128 - this will panic, as a realized number cannot describe a complex number
//
// NOTE: Parse operations do not incorporate the underlying action potential of the provided operand.
//
// For dynamic number generation, see NewRealized
func ParseRealized[T Advanced](operand T, base ...uint16) Realized {
	b := PanicIfInvalidBase(base[0])

	switch raw := any(operand).(type) {
	case uint, uint8, uint16, uint32, uint64, uintptr,
		int, int8, int16, int32, int64,
		Natural, Realized, string:
	case float32:
		if math.IsInf(float64(raw), 0) {
			panic(fmt.Sprintf("cannot create a Realized from an Inf valued %T", raw))
		}
		if math.IsNaN(float64(raw)) {
			panic(fmt.Sprintf("cannot create a Realized from an NaN valued %T", raw))
		}
	case float64:
		if math.IsInf(raw, 0) {
			panic(fmt.Sprintf("cannot create a Realized from an Inf valued %T", raw))
		}
		if math.IsNaN(raw) {
			panic(fmt.Sprintf("cannot create a Realized from an NaN valued %T", raw))
		}
	default:
		panic(fmt.Sprintf("cannot create a Realized from type %T", raw))
	}

	return parseRealized(ToString(operand), b)
}

func parseRealized(source string, base uint16) Realized {
	if len(source) == 0 {
		return Realized{
			irrational: false,
			Negative:   false,
			whole:      ParseNatural("0"),
			fractional: ParseNatural("0"),
			base:       base,
			precision:  &atlas.Precision,
			created:    true,
		}
	}

	negative := false
	irrational := false
	if source[0] == '~' {
		irrational = true
		source = source[1:]
	}
	if source[0] == ' ' {
		source = source[1:]
	}
	if source[0] == '-' {
		negative = true
		source = source[1:]
	}
	if source[0] == ' ' {
		source = source[1:]
	}

	t := transcendental.IsIdentifier(string(source[0]))
	if t != transcendental.Non {
		r := Transcendental.From(base, t)
		r.Negative = negative
		r.created = true
		return r
	}

	var digits []string
	if base > 16 {
		digits = strings.Split(source, " ")
	} else {
		digits = strings.Split(source, "")
	}

	var wholeDigits []string
	var fractionalDigits []string

	isWholePart := true
	startCount := false
	periodicWidth := uint(0)
	for i := 0; i < len(digits); i++ {
		if digits[i] == "." {
			isWholePart = false
			continue
		}
		if digits[i] == "‾" {
			startCount = true
			continue
		}

		if startCount {
			periodicWidth++
		}
		if isWholePart {
			wholeDigits = append(wholeDigits, digits[i])
		} else {
			fractionalDigits = append(fractionalDigits, digits[i])
		}
	}

	var wholePart string
	var fractionalPart string
	if base > 16 {
		wholePart = strings.Join(wholeDigits, " ")
		fractionalPart = strings.Join(fractionalDigits, " ")
	} else {
		wholePart = strings.Join(wholeDigits, "")
		fractionalPart = strings.Join(fractionalDigits, "")
	}

	return Realized{
		irrational:    irrational,
		Negative:      negative,
		whole:         ParseNatural(wholePart, base),
		fractional:    ParseNatural(fractionalPart, base),
		periodicWidth: periodicWidth,
		precision:     &atlas.Precision,
		created:       true,
	}
}

// NewRealized - Creates a dynamic realized number, which realizes it's value from the provided action potential functions.
// see.ActionPotentials and see.RealizedNumbers
//
// NOTE: This does NOT impulse the action!
//
// NOTE: If no base is provided, base₁₀ is implied.
//
// For static number generation, see ParseRealized.
func NewRealized(action func(Realization) Realization, potential func() bool, base ...uint16) Realized {
	b := PanicIfInvalidBase(base[0])
	return Realized{
		whole:         ParseNatural("0", b),
		fractional:    ParseNatural("0", b),
		periodicWidth: 0,
		precision:     &atlas.Precision,
		revelation:    action,
		potential:     potential,
		base:          b,
		created:       true,
	}
}

func (r *Realized) sanityCheck() {
	if !r.created {
		panic("this realized was not created through a constructor")
	}
}

func (r *Realized) realize() {
	r.sanityCheck()

	// Self-realization! =)

	self := r.revelation(Realization{
		Irrational:    r.irrational,
		Negative:      r.Negative,
		Whole:         r.whole,
		Fractional:    r.fractional,
		periodicWidth: r.periodicWidth,
	})

	r.irrational = self.Irrational
	r.Negative = self.Negative
	r.whole = self.Whole
	r.fractional = self.Fractional
	r.periodicWidth = self.periodicWidth
}

func (r *Realized) Digits() (whole []bases.Digit, fractional []bases.Digit) {
	r.sanityCheck()

	// 0 - Get the digits
	whole = r.whole.Digits()
	fractional = r.fractional.Digits()

	// 1 - Check if it's periodic
	if r.periodicWidth > 0 {

		// 2 - Run the digits out to precision width

		periodic := fractional[len(fractional)-int(r.periodicWidth):]
		delta := int(*r.precision) - len(fractional)
		toAppend := make([]bases.Digit, delta)

		ii := 0
		for i := 0; i < delta; i++ {
			toAppend[i] = periodic[ii]

			ii++
			if ii >= len(periodic) {
				ii = 0
			}
		}

		fractional = append(fractional, toAppend...)

		// 3 - Backtrack to perform a round operation

		keepRounding := false
		for i := len(fractional) - 1; i >= 0; i-- {
			fractional[i] = fractional[i].Increment(r.base)
			if fractional[i] > 0 {
				break
			}
			if i == 0 {
				keepRounding = true
			}
		}

		if keepRounding {
			for i := len(whole) - 1; i >= 0; i-- {
				whole[i] = whole[i].Increment(r.base)
				if whole[i] > 0 {
					break
				}
			}
		}
	}
	return whole, fractional
}

// Width returns the number of placeholders the whole and fractional parts require.
func (r *Realized) Width() (whole uint, fractional uint) {
	r.sanityCheck()

	if r.irrational || r.periodicWidth > 0 {
		return uint(len(r.whole.Digits())), *r.precision
	}
	return uint(len(r.whole.Digits())), uint(len(r.fractional.Digits()))
}

// Impulse tests the potential and then sparks the Realized number's neural pathway.
func (r *Realized) Impulse() {
	r.sanityCheck()

	if r.potential() {
		r.gate.Lock()
		defer r.gate.Unlock()

		r.realize()
	}
}

// Reveal tests the potential and then sparks the neural pathway before revealing the Realized number in a single
// lock operation.  In this case, the neurological response is to Print the realized number.
func (r *Realized) Reveal() string {
	r.sanityCheck()

	// NOTE: This intentionally locks for the entire operation and cannot be replaced with a call to Impulse()
	if r.potential() {
		r.gate.Lock()
		defer r.gate.Unlock()

		r.realize()
	}
	return r.print()
}

// Precision "sets and/or gets" the precision of the Realized number.  If no precision is provided, this simply returns
// the stored value - otherwise, this will set the precision AND call Impulse (as the realized number must be re-realized).
//
// NOTE: precision is a reference value!
//
// In a neural architecture, calls like this do NOT guarantee you will see a change in the output =)
//
// For most operations, the pathways aren't gated beyond a call to when.Always - but, in a live system, the value
// may not have another "revelation" until time has passed, or a condition has been met (for instance).
//
// see.ActionPotentials, see.Neuron, and see.RealizedNumbers.
func (r *Realized) Precision(precision ...*uint) uint {
	r.sanityCheck()

	if len(precision) > 0 {
		r.precision = precision[0]
		r.Impulse()
	}
	return *r.precision
}

// Base "sets and/or gets" the base of the Realized number.  If no base is provided, this simply returns
// the stored value - otherwise, this will set the base AND call Impulse (as the realized number must be re-realized).
//
// NOTE:
//
// In a neural architecture, calls like this do NOT guarantee you will see a change in the output =)
//
// For most operations, the pathways aren't gated beyond a call to when.Always - but, in a live system, the value
// may not have another "revelation" until time has passed, or a condition has been met (for instance).
//
// see.ActionPotentials, see.Neuron, and see.RealizedNumbers.
func (r *Realized) Base(base ...uint16) uint16 {
	r.sanityCheck()

	if len(base) > 0 {
		r.base = PanicIfInvalidBase(base[0])
		r.Impulse()
	}
	return r.base
}

// String - see.PrintingNumbers
func (r *Realized) String() string {
	r.sanityCheck()

	// NOTE: These lock to ensure another thread doesn't mutate the whole and fractional parts mid-print.
	r.gate.Lock()
	defer r.gate.Unlock()
	return r.print()
}

// Print - see.PrintingNumbers
func (r *Realized) Print(base ...uint16) string {
	r.sanityCheck()

	// NOTE: These lock to ensure another thread doesn't mutate the whole and fractional parts mid-print.
	r.gate.Lock()
	defer r.gate.Unlock()
	return r.print(base...)
}

// Matrix - see.PrintingNumbers
func (r *Realized) Matrix(whole, fractional uint) string {
	r.sanityCheck()

	// NOTE: These lock to ensure another thread doesn't mutate the whole and fractional parts mid-print.
	r.gate.Lock()
	defer r.gate.Unlock()
	return ""
}

// print is a non-locked printing function.
func (r *Realized) print(base ...uint16) string {
	r.sanityCheck()

}
