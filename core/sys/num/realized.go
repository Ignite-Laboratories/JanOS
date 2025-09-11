package num

import (
	"core/sys/atlas"
	"core/sys/num/internal"
	"strings"
	"sync"
)

// A Realized number is one generated through the execution of a neural pathway.  see.RealizedNumbers
//
// To create a Realized number, please use one of the following methods:
//
// ParseRealized - Creates a static realized number.
//
// NewRealized - Creates a dynamic realized number.
type Realized struct {
	gate sync.Mutex

	Identity string

	irrational bool
	Negative   bool
	whole      Natural
	fractional Natural
	periodic   Natural

	revelation func(Realization, uint16, uint) Realization
	potential  func() bool

	precision       *uint
	_precisionStale bool
	_precisionNew   *uint

	base       uint16
	_baseStale bool
	_baseNew   uint16

	created bool
}

// ParseRealized - Creates a static realized number from an advanced operand - including 'string'.
//
// Primitive operands:
//
//	base₁₀ is implied and whatever base you provide is ignored.
//
// Advanced operands:
//
//	string - the operand must be encoded in the provided base value (or base₁₀ if omitted)
//	Natural - sets the whole part and base of the realized number (or base₁₀ if omitted)
//	Realized - sets all the parts and assigns the number to the provided base (or base₁₀ if omitted)
//
// NOTE: Parse operations do not incorporate the underlying action potential of the provided operand.
//
// For dynamic number generation, see NewRealized
func ParseRealized(operand any, base ...uint16) Realized {
	b := PanicIfInvalidBase(base[0])
	op := ToString(FilterOperands(b, operand)[0])

	if len(op) == 0 {
		return Realized{
			irrational: false,
			Negative:   false,
			whole:      NaturalZero,
			fractional: NaturalZero,
			periodic:   NaturalZero,
			base:       b,
			precision:  &atlas.Precision,
			created:    true,
		}
	}

	negative := false
	irrational := false
	if op[0] == '~' {
		irrational = true
		op = op[1:]
	}
	if op[0] == ' ' {
		op = op[1:]
	}
	if op[0] == '-' {
		negative = true
		op = op[1:]
	}
	if op[0] == ' ' {
		op = op[1:]
	}

	var digits []string
	if b > 16 {
		digits = strings.Split(op, " ")
	} else {
		digits = strings.Split(op, "")
	}

	var wholeDigits []string
	var fractionalDigits []string
	var periodicDigits []string

	isWholePart := true
	isPeriodicPart := false
	for i := 0; i < len(digits); i++ {
		if digits[i] == "." {
			isWholePart = false
			continue
		}
		if digits[i] == "‾" {
			isPeriodicPart = true
			continue
		}

		if isWholePart {
			wholeDigits = append(wholeDigits, digits[i])
		} else if isPeriodicPart {
			periodicDigits = append(periodicDigits, digits[i])
		} else {
			fractionalDigits = append(fractionalDigits, digits[i])
		}
	}

	var wholePart string
	var fractionalPart string
	var periodicPart string
	if b > 16 {
		wholePart = strings.Join(wholeDigits, " ")
		fractionalPart = strings.Join(fractionalDigits, " ")
		periodicPart = strings.Join(periodicDigits, " ")
	} else {
		wholePart = strings.Join(wholeDigits, "")
		fractionalPart = strings.Join(fractionalDigits, "")
		periodicPart = strings.Join(periodicDigits, "")
	}

	return Realized{
		irrational: irrational,
		Negative:   negative,
		whole:      ParseNatural(wholePart, b),
		fractional: ParseNatural(fractionalPart, b),
		periodic:   ParseNatural(periodicPart, b),
		precision:  &atlas.Precision,
		created:    true,
	}
}

// NewRealized - Creates a dynamic realized number, which realizes it's value from the provided action potential functions.
// see.ActionPotentials and see.RealizedNumbers
//
// For more details on the action, see Realized.SetAction()
//
// For more details on the potential, see Realized.SetPotential()
//
// NOTE: This does NOT impulse the action!
//
// NOTE: If no base is provided, base₁₀ is implied.
//
// For static number generation, see ParseRealized.
func NewRealized(action func(current Realization, base uint16, precision uint) Realization, potential func() bool, base ...uint16) Realized {
	b := PanicIfInvalidBase(base[0])

	return Realized{
		whole:      NaturalZero,
		fractional: NaturalZero,
		periodic:   NaturalZero,
		precision:  &atlas.Precision,
		revelation: action,
		potential:  potential,
		base:       b,
		created:    true,
	}
}

// SetAction sets the current revelation action, while SetPotential sets its potential function - see.ActionPotentials
//
// A realized number's action should take in the current realization, the base to produce a result in, and the placeholder
// precision with which to calculate the result to.  It should output a result realization and an optional identity value.
//
// For example, if building a realization of π, you might choose to reveal the identity "π" during activation.
func (r *Realized) SetAction(action func(current Realization, base uint16, precision uint) Realization) {
	r.revelation = action
}

// SetPotential sets the current revelation potential, while SetAction sets its activation action.
//
// A potential is tested before a revelation can be fired to determine if the revelation should even
// take place yet.  This should yield 'true' when a revelation should occur - see.ActionPotentials
func (r *Realized) SetPotential(potential func() bool) {
	r.potential = potential
}

func (r *Realized) sanityCheck(base ...uint16) uint16 {
	if !r.created {
		panic("this realized was not created through a constructor")
	}
	return PanicIfInvalidBase(base...)
}

func (r *Realized) realize() {
	r.sanityCheck()

	// Check for any changes to precision or base - last one in wins

	if r._baseStale {
		r.base = r._baseNew
		r._baseStale = false
	}
	if r._precisionStale {
		r.precision = r._precisionNew
		r._precisionStale = false
	}

	if r.revelation == nil {
		return
	}

	// Self-realization! =)

	whole, fractional, periodic := r.Digits()
	self := r.revelation(Realization{
		Irrational: r.irrational,
		Negative:   r.Negative,
		Whole:      whole,
		Fractional: fractional,
		Periodic:   periodic,
	}, r.base, *r.precision)

	// If the user indicates a periodic width but DIDN'T trim their fractional component, that's OKAY!
	// We should allow that, as they are NOT expected to understand the inner workings of 𝑡𝑖𝑛𝑦 =)

	for func() bool {
		if len(self.Fractional) > len(self.Periodic) {
			last := self.Fractional[len(self.Fractional)-len(self.Periodic):]

			match := true
			for i, d := range last {
				if self.Periodic[i] != d {
					match = false
					break
				}
			}
			if match {
				self.Fractional = self.Fractional[:len(self.Fractional)-len(self.Periodic)]
				return true
			}
		}
		return false
	}() {
		// ...
	}

	r.irrational = self.Irrational
	r.Negative = self.Negative
	r.whole = ParseNatural(self.Whole, r.base)
	r.fractional = ParseNatural(self.Fractional, r.base)
	r.periodic = ParseNatural(self.Periodic, r.base)
}

func (r *Realized) Digits() (whole []byte, fractional []byte, periodic []byte) {
	r.sanityCheck()

	whole = r.whole.Digits(r.base)
	fractional = r.fractional.Digits(r.base)
	periodic = r.periodic.Digits(r.base)

	return whole, fractional, periodic
}

// Width returns the number of calculated placeholders in the whole and fractional components.
//
// NOTE: For irrational or periodic values, this will return the stored precision for the fractional component.
func (r *Realized) Width() (whole uint, fractional uint) {
	r.sanityCheck()

	if r.irrational || len(r.periodic.String()) > 0 {
		return uint(len(r.whole.Digits())), *r.precision
	}
	return uint(len(r.whole.Digits())), uint(len(r.fractional.Digits()))
}

// Impulse tests the potential and then sparks the Realized number's neural pathway.
func (r *Realized) Impulse() {
	r.sanityCheck()

	if r.potential != nil && r.potential() {
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
	if r.potential != nil && r.potential() {
		r.gate.Lock()
		defer r.gate.Unlock()

		r.realize()
	}
	return r.print(r.base)
}

// Precision "sets and/or gets" the precision of the Realized number.  If no precision is provided, this simply returns
// the stored value - otherwise, this will set the precision AND call Impulse (as the realized number must be re-realized).
//
// NOTE: precision is a reference value!
//
// NOTE: This is a neural architecture - so setting the value does NOT guarantee it has actually picked up the change, yet.
//
// see.ActionPotentials, see.Neuron, and see.RealizedNumbers.
func (r *Realized) Precision(precision ...*uint) uint {
	r.sanityCheck()

	if len(precision) > 0 {
		r._precisionNew = precision[0]
		r._precisionStale = true
		r.Impulse()
	}
	return *r.precision
}

// Base "sets and/or gets" the base of the Realized number.  If no base is provided, this simply returns
// the stored value - otherwise, this will set the base AND call Impulse (as the realized number must be re-realized).
//
// NOTE: The periodic nature of the number
//
// NOTE: This is a neural architecture - so setting the value does NOT guarantee it has actually picked up the change, yet.
//
// see.ActionPotentials, see.Neuron, and see.RealizedNumbers.
func (r *Realized) Base(base ...uint16) uint16 {
	r.sanityCheck()

	if len(base) > 0 {
		r._baseNew = PanicIfInvalidBase(base[0])
		r._baseStale = true
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
	return r.print(-1, true, r.base)
}

// Print - see.PrintingNumbers
//
// To print your value to whatever precision it's currently calculated out to, please use a fractionalWidth of '-1'.
// Otherwise, fractionalWidth will round the fractional part of your number early, or right pad it with zeros to width.
func (r *Realized) Print(fractionalWidth int, base ...uint16) string {
	b := r.sanityCheck(base...)

	// NOTE: These lock to ensure another thread doesn't mutate the whole and fractional parts mid-print.
	r.gate.Lock()
	defer r.gate.Unlock()
	return r.print(fractionalWidth, false, b)
}

// Matrix - see.PrintingNumbers
//
// When "matrixing" operands together, their whole parts are each left-padded with zeros to the widest operand's
// whole-part width.  The fractional part can either follow the same logic (using a fractionalWidth of '-1') or
// be explicitly defined.  As with Print operations, if setting a fractionalWidth other than -1, the fractional
// component will either be rounded early or right-padded with zeros to the desired width.
func (r *Realized) Matrix(fractionalWidth int, operands ...any) string {
	r.sanityCheck()

	// NOTE: These lock to ensure another thread doesn't mutate the whole and fractional parts mid-print.
	r.gate.Lock()
	defer r.gate.Unlock()

	w, f, p := r.Digits()

	if whole < uint(len(w)) {
		// Trim the whole part down, as they want a smaller matrix
		delta := uint(len(w)) - whole
		w = w[delta:]
	} else {
		// Pad the left of the whole part with 0s
		delta := uint(len(w)) - whole
		prepend := make([]byte, delta)
		for i := uint(0); i < delta; i++ {
			prepend[i] = byte(0)
		}
		w = append(prepend, w...)
	}

	if fractional < uint(len(f)) {
		// Trim the fractional part down and return

		return
	}
	// Otherwise, pad with zeros or periodic

	wStr := make([]string, len(w))
	for i, d := range w {
		wStr[i] = internal.PrintDigit(d)
	}

	fStr := make([]string, len(f))
	for i, d := range f {
		fStr[i] = internal.PrintDigit(d)
	}

	pStr := make([]string, len(p))
	for i, d := range p {
		pStr[i] = internal.PrintDigit(d)
	}
	return ""
}

// print is a non-locked printing function.
func (r *Realized) print(fractionalWidth int, truncateIrrationals bool, base uint16) string {
	var prefix []string
	if r.irrational {
		prefix = append(prefix, "~")
	}
	if r.Negative {
		prefix = append(prefix, "-")
	}

	whole, fractional, periodic := r.Digits()

	wholeStr := make([]string, len(whole))
	for i, d := range whole {
		wholeStr[i] = internal.PrintDigit(d)
	}

	fractionalStr := make([]string, len(fractional))
	for i, d := range fractional {
		fractionalStr[i] = internal.PrintDigit(d)
	}

	periodicStr := make([]string, len(periodic))
	for i, d := range periodic {
		periodicStr[i] = internal.PrintDigit(d)
	}

	components := append(prefix, wholeStr...)
	if len(fractionalStr) > 0 || len(periodicStr) > 0 {
		components = append(components, ".")
		components = append(components, fractionalStr...)

		if len(periodicStr) > 0 {
			components = append(components, "‾")
			components = append(components, periodicStr...)
		}
	}

	if base > 16 {
		return strings.Join(components, " ")
	}
	return strings.Join(components, "")
}
