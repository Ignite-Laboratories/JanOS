package std

import "strings"

// Natural represents a base-explicit Digit slice representing a value belonging to the set of naturally
// countable numbers - or all positive whole numbers, including zero.
//
// To those who think zero shouldn't be included in the set of natural numbers, I present a counter-argument:
// base 1 has only one identifier, meaning it can only "represent" zero by -not- holding a value in an observable
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
// numbers.  I can't stop you from emptying your natural's digits and creating a 'nil' state - but I can empower you with awareness of it =)
//
// See Real, Complex, Index, and Operable
type Natural struct {
	Digits []Digit
	Base   byte
}

// String returns the string representation of this natural number in its current base.  As bases above 16 begin to
// double up the hexadecimal characters to span the remainder of the address space, digits above base 16 are spaced
// with a single whitespace character between them.
func (a Natural) String() string {
	str := ""

	if a.Base > 16 {
		for _, d := range a.Digits {
			str += " " + d.String()
		}
		str = strings.TrimSpace(str)
	} else {
		for _, d := range a.Digits {
			str += d.String()
		}
	}
	return str
}
