package std

import "fmt"

// Lexeme is the base level of all semantic operations.  In some cases, this might be a word - others, a special
// character, or even a phrase!  Ultimately, this is just a container for a single logical string value of arbitrary width.
//
// There are many kinds of lexemes, so you'll have to branch through their documentation to follow along.
//
// See Lexeme, lexeme.Base, lexeme.Function
type Lexeme interface {
	fmt.Stringer
}
