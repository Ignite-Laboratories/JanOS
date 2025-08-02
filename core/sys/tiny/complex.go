package tiny

// Complex represents an Operable Phrase which holds two Real numbers - a "real" part, and an "imaginary" part.
//
// See Natural, Real, Index, and Operable
type Complex struct {
	// Name represents the name of this complex number.  By default, numbers are given a random cultural name to ensure that
	// it doesn't step on any of the standard variable names ('a', 'x', etc...) you'll want to provide.  The names provided
	// are guaranteed to be a single word containing only letters of the English alphabet for fluent proof generation.
	Name string

	Real      real.Real
	Imaginary real.Real
}
