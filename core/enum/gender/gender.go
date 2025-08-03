// Package gender provides access to the gender enumeration.
package gender

// Gender provides global identifiers for Male, Female, or NonBinary interpretations.
//
// NOTE: Everything is inherently 'non-binary' - these simply are contextual guides =)
type Gender byte

const (
	// Female represents the colloquially 'female' gender.
	Female Gender = iota

	// Male represents the colloquially 'male' gender.
	Male

	// NonBinary represents the colloquial 'any' gender.
	NonBinary
)
