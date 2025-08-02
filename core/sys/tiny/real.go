package tiny

import (
	"fmt"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/measurement"
	"github.com/ignite-laboratories/core/std/name"
	"github.com/ignite-laboratories/core/std/phrase"
	"math"
	"math/big"
	"strings"
)

// Real represents an Operable Phrase that represents a "real number" - real numbers are held in four measurements:
//
//	  ⬐ The Sign              ⬐ The Fractional Part
//	| 1 - 1 0 1 1 0 - 1 0 0 0 1 0 1 0 1 1 - 1 1 0 | ( -22.5556̅ )
//	          ⬑ The Whole Part                ⬑ The Periodic Part
//
// All parts (except the sign) can grow to arbitrary widths to accommodate whatever size number you can imagine.  To put
// this into perspective, a 256 bit number can hold up to the value 1.1579208923731619542357098500869x10⁷⁷!  A single
// gigabyte of memory can hold up to 3.2x10⁸ 256-bit real numbers - with the most common numbers fitting in less than 64 bits.
// In essence, you've got plenty of storage in the modern age to work with even the most astronomically large of numbers =)
//
// By default, real numbers are given a maximum combined -decimal- precision bit width of 256 bits - but you may override that if desired.
// This prevents infinitely repeating (or irrational) numbers from exhausting your computer's memory, unless you specifically need it to.
//
// After every arithmetic operation, a check is performed to see if the periodic part is missing and if the fractional part
// fills the entire allotted precision - if so, the real number is deemed to be "irrational".
//
// The REASON to work with a type like this is to ensure that all arithmetic is done as -Math- intended, not within the bounds of
// tight computational memory spaces.  By knowing exactly where the decimal point is located, all reals can be aligned implicitly
// by the matrix engine without first performing what's called 'type coercion' (or defining how to switch between numeric encoding
// schemes on the fly).  The need to differentiate between floating point and integer numbers is entirely a computer science issue
// born from ancient memory requirements and NOT one that a mathematician should have to bear while exploring their theories.
//
// See Natural, Complex, Index, and Operable
type Real struct {
	// Name represents the name of this real number.  By default, numbers are given a random cultural human name to ensure that
	// it doesn't step on any of the standard variable names you'll want to provide ('a', 'x', 'y', etc...).  The names provided
	// are guaranteed to be a single word containing only letters of the English alphabet for fluent proof generation.
	Name string

	base byte

	// Precision represents the maximum combined bit-width of any part of the real number beyond the decimal place.
	//
	// NOTE: This defaults to 256 bits through the real creation functions.
	Precision uint // Defaults to 256

	// Irrational is true when the number continues on indefinitely with no observed repetition up to the defined precision.
	Irrational bool

	// Negative represents the sign of the real number - with true representing negative.
	Negative bool

	// Whole represents the whole part of the real number.
	Whole Natural

	// Fractional represents the decimal portion of the real number.
	Fractional Natural

	// Periodic represents the periodic end of the fractional portion of the real number and may or may not be present.
	Periodic Natural
}

// NewReal creates a new instance of a Real number from the provided Primitive value.
//
// NOTE: You may also set the desired precision at this point, though it defaults to 256.
func NewReal[T std.Primitive](value T, precision ...uint) Real {
	return NewRealNamed(name.Tiny().Name, value, precision...)
}

// NewRealNamed creates a new instance of a named Real number from the provided Primitive value and name.
//
// NOTE: You may also set the desired precision at this point, though it defaults to 256.
func NewRealNamed[T std.Primitive](name string, value T, precision ...uint) Real {
	p := uint(256)
	if len(precision) > 0 {
		p = precision[0]
	}

	out := Real{
		Name:       name,
		Precision:  p,
		Whole:      NaturalFrom(0),
		Fractional: NaturalFrom(0),
		Periodic:   NaturalFrom(0), // *pushes glasses up nose* - "technically," all numbers fractionally end in infinitely repeating zeros =)
	}

	switch operand := any(value).(type) {
	case *big.Int:
		out.Negative = operand.Sign() < 0
		out.Whole = Natural{
			Measurement: measurement.OfString(operand.Text(2)),
		}
	case *big.Float:
		out.Negative = operand.Sign() < 0
		operand = new(big.Float).Abs(operand)

		// NOTE: We're already working with a big type, so don't be afraid to leverage it's power =)

		entire := operand.Text('f', int(out.Precision))
		pointPos := strings.Index(entire, ".")

		if pointPos <= 0 {
			// No decimal place - this is a whole number

			whole, _ := new(big.Int).SetString(entire, 10)
			out.Whole = Natural{
				Measurement: measurement.OfString(whole.Text(2)),
			}
		} else {
			// A decimal place was found - this is a fractional number

			whole, _ := new(big.Int).SetString(entire[:pointPos], 10)
			out.Whole = Natural{
				Measurement: measurement.OfString(whole.Text(2)),
			}

			fractional, _ := new(big.Int).SetString(entire[pointPos+1:], 10)
			out.Fractional = Natural{
				Measurement: measurement.OfString(fractional.Text(2)),
			}
		}
	case uint8, uint16, uint32, uint64, uint, uintptr:
		var v uint
		switch u := operand.(type) {
		case uint8:
			v = uint(u)
		case uint16:
			v = uint(u)
		case uint32:
			v = uint(u)
		case uint64:
			v = uint(u)
		case uint:
			v = u
		case uintptr:
			v = uint(u)
		}

		out.Whole = NaturalFrom(v)
	case int8, int16, int32, int64, int:
		var v int
		switch i := operand.(type) {
		case int8:
			v = int(i)
		case int16:
			v = int(i)
		case int32:
			v = int(i)
		case int64:
			v = int(i)
		case int:
			v = i
		}

		out.Negative = v < 0
		out.Whole = NaturalFrom(uint(v))
	case float32, float64:
		var v float64
		switch f := operand.(type) {
		case float32:
			v = float64(f)
		case float64:
			v = f
		}

		if math.IsNaN(v) {
			panic(fmt.Errorf("cannot create real from a 'NaN' valued float"))
		}
		if math.IsInf(v, 0) {
			panic(fmt.Errorf("cannot create real from an 'Inf' valued float"))
		}

		// Hand this off to big.Float, as they have EXCELLENT and robust precision guarantees.
		out = NewRealNamed(name, big.NewFloat(v), precision...)
	case complex64:
		if imag(operand) != 0 {
			panic(fmt.Errorf("cannot create real from a complex number with a non-zero imaginary part - [%v]", operand))
		}

		// Hand this off as a float32 for big.Float to process
		out = NewRealNamed(name, real(operand), precision...)
	case complex128:
		if imag(operand) != 0 {
			panic(fmt.Errorf("cannot create real from a complex number with a non-zero imaginary part - [%v]", operand))
		}

		// Hand this off as a float64 for big.Float to process
		out = NewRealNamed(name, real(operand), precision...)
	case bool:
		if operand {
			out.Whole = NaturalFrom(1)
		} else {
			out.Whole = NaturalFrom(0)
		}
	default:
		panic(fmt.Errorf("cannot create real from primitive type '%T'", operand))
	}

	// Lastly, perform checks for rationality and infinite repetition
	return out.cleanup()
}

// SetBase sets the base this number should be observed in.  While the underlying information (and all arithmetic) is done in
// binary, the representation of the number is still handled in an appropriate base.  This allows us to capture the periodic
// portion of the fractional part of the real number, since the periodic portion can transiently appear across bases.
//
// When converting between string and numeric form in different bases the string should comply with the below table -
//
//	 Base | Symbol Set
//	 0-1  | Unsupported
//	 2-10 | 0-9
//	11-36 | 0-9, then A-Z
//	37-60 | 0-9, then A-Z, then the lowercase Greek alphabet
//	 61+  | Unsupported
//
// For reference here's the lowercase Greek alphabet, in order -
//
//	α β γ δ ε ζ η θ ι κ λ μ ν ξ ο π ρ σ τ υ φ χ ψ ω
//	              NOTE: Medial Sigma ⬏
func (a Real) SetBase(base byte) Real {
	a.base = base
	a.sanityCheck()
	return a
}

// Text converts the Real to a string of its currently set base value.
//
// See SetBase() for a table of how the string will be encoded.
func (a Real) Text() string {
	// TODO: Implement Real.Text()
	// The periodic part should be interleaved with the overline character in unicode 0̅1̅2̅3̅4̅5̅6̅7̅8̅9̅
	return ""
}

/**
Utilities
*/

// sanityCheck is a function to perform general sanity checks against the real number, such as for an invalid base value.
func (a Real) sanityCheck() Real {
	if a.base < 2 {
		panic(fmt.Errorf("base %d is unsupported", a.base))
	}
	if a.base > 60 {
		panic(fmt.Errorf("base %d is unsupported", a.base))
	}
	return a
}

// cleanup expands the real number to full precision and then checks for irrationality or periodicity in
// the fractional component before rolling up those conditions into the appropriate measurements.
func (a Real) cleanup() Real {
	// TODO: Implement a periodicity and irrationality check
	return a
}

// BitWidth returns the overall bit width of this Real in its smallest representation, up to its decimal precision width.
//
// NOTE: This includes the sign as a single Bit value.
func (a Real) BitWidth() uint {
	// TODO: Handle the periodic bit width appropriately
	return a.Whole.BitWidth() + a.Fractional.BitWidth() + a.Periodic.BitWidth()
}

func (a Real) AsPhrase() std.Phrase[any] {
	sign := measurement.OfBits(0)
	if a.Negative {
		sign = measurement.OfBits(1)
	}

	// TODO: Handle the periodic bit width appropriately
	return phrase.Of[any](sign, a.Whole.Measurement, a.Fractional.Measurement, a.Periodic.Measurement)
}
