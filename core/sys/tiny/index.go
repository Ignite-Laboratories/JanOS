package tiny

import "github.com/ignite-laboratories/core/std"

// Index represents an implicitly fixed-width phrase of raw binary information.
//
// See Natural, Real, Complex, and Operable
type Index struct {
	// Name represents the name of this index.  By default, indexes are given a random cultural name to ensure that
	// it doesn't step on any of the standard variable names ('a', 'x', etc...) you'll want to provide.  The names provided
	// are guaranteed to be a single word containing only letters of the English alphabet for fluent proof generation.
	Name string

	std.Phrase
	Width uint
	Flow  int
}

// TODO: Make indexes support a signed "flow" value that can be reset on demand, indicating how many times it over or under-flowed

func (a Index) AsPhrase() std.Phrase {
	// TODO: Convert an Index to a Phrase
	return std.Phrase{}
}
