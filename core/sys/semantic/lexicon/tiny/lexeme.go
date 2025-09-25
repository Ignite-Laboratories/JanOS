package tiny

import (
	"fmt"
)

// Lexeme is the most basic unit of the 𝑡𝑖𝑛𝑦 lexicon and represents a single placeholder value in an arithmetic matrix.
//
// NOTE: All bases from Base2 to Base256 are well-defined in the lexicon, but the most common are aliased for convenience.
//
// See Lexeme, Digit, Binary, Ternary, Octal, Decimal, Hexadecimal, and Sexagesimal
type Lexeme struct {
	Placeholder uint8
	Base        Base
}

func (t Lexeme) String() string {
	return fmt.Sprintf("%x", t.Placeholder)
}
