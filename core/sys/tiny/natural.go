package tiny

import (
	"github.com/ignite-laboratories/core/internal"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/measurement"
)

// Natural represents a kind of Measurement with a value belonging to the set of naturally countable numbers - or all
// positive whole numbers, including zero.
//
// To those who think zero shouldn't be included in the set of natural numbers, I present a counter-argument:
// Base 1 has only one identifier, meaning it can only "represent" zero by -not- holding a value in an observable
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
//
// See Real, Complex, Index, and Operable
type Natural struct {
	std.Measurement[any]
}

// Text converts the Natural to a string of the provided base, encoded to the specification defined by Real.SetBase()
func (a Natural) Text(base byte) string {
	// TODO: Implement Natural.Text(base)
	return ""
}

// NaturalFrom takes a Measurement of the provided unsigned integer value as a Natural number.
func NaturalFrom(value uint) Natural {
	return Natural{
		Measurement: measurement.OfBytes(internal.Measure(value)[0]...),
	}
}

// FromString creates a new Natural measurement that represents the provided base-encoded string.
//
// NOTE: The input string must be encoded as expected by Real.SetBase()
func FromString(base byte, value string) Natural {
	// TODO: Implement this
	panic("unsupported")
}
