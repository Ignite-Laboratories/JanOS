// Package gender provides access to the gender enumeration.
package gender

// Gender provides global identifiers for Male, Female, or NonBinary interpretations.
//
// NOTE: Everything is inherently 'non-binary' - these simply are contextual guides =)
//
// See Gender, Female, Male, and NonBinary
type Gender byte

const (
	// Female represents the colloquially 'female' gender.
	//
	// See Gender, Female, Male, and NonBinary
	Female Gender = iota

	// Male represents the colloquially 'male' gender.
	//
	// See Gender, Female, Male, and NonBinary
	Male

	// NonBinary represents the colloquial 'any' gender.
	//
	// See Gender, Female, Male, and NonBinary
	NonBinary
)
