package std

import "fmt"

// Lexeme is just an alias for fmt.Stringer and forms the base of all semantic operations.  In some cases, this might be
// a word - others, a special character, object, or even an idiom!  Ultimately, this is just a container for a single logical
// string value of arbitrary width.
//
// There are many kinds of lexemes, so you'll have to branch through their documentation to follow along.
//
// See std.Lexeme, lexeme.Base, lexeme.Function, lexeme.Lexicon, and lexeme.Special
type Lexeme interface {
	fmt.Stringer
}
