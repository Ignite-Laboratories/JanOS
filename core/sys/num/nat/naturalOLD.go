package nat

import (
	"core/sys/num"
	"fmt"
	"strings"
)

// Natural represents a value belonging to the set of naturally countable numbers - or all positive whole numbers, including zero.
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
type Natural struct {
	digits []num.Numeric[byte]
	base   uint16
}

// ParseNatural takes an input string in the provided base and converts it to a new Natural number.
//
// NOTE: If no base is provided, the input is assumed to be base₁₀
func ParseNatural(input string, base ...uint16) num.Realized {
	// TODO: Implement this
	panic("cannot yet parse strings into naturals")
}

// NewNatural creates a new instance of a Natural number from the provided base₁₀ input string, then converts it to the desired base.
// A base₁₀ natural input string may only contain the characters [0-9] or [-].
//
// NOTE: If no base is provided, base₁₀ is implied.
func NewNatural(input string, base ...uint16) Natural {
	b := uint16(10)
	if len(base) > 0 {
		b = base[0]
	}

	n := Natural{
		digits: make([]num.Numeric[byte], len(input)),
		base:   b,
	}
	for i := len(input) - 1; i >= 0; i-- {
		var bnd num.Numeric[byte]
		switch input[i] {
		case '-':
			panic("natural numbers cannot be negative")
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			bnd, _ = num.NewNumericBounded[byte](input[i], 0, 10)
		default:
			panic("invalid character")
		}
		n.digits[i] = bnd
	}

	n.ChangeBase(b)
	return n
}

func (n *Natural) ChangeBase(base uint16) {
	n.base = base
	// TODO: Convert to the desired base
	panic("cannot change bases yet")
}

func (n *Natural) Width() uint {
	return uint(len(n.digits))
}

func (n *Natural) Base() uint16 {
	return n.base
}

func (n Natural) String() string {
	out := make([]string, len(n.digits))
	for i, d := range n.digits {
		out[i] = fmt.Sprintf("%x", d.Value())
	}
	if n.base > 16 {
		return strings.Join(out, " ")
	}
	return strings.Join(out, "")
}
