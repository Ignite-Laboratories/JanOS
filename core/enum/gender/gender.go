// Package gender provides access to the gender enumeration.
package gender

// Gender provides global identifiers for Male, Female, or NonBinary interpretations.
//
// NOTE: Everything is inherently 'non-binary' - these simply are contextual guides =)
type Gender int

const (
	Female Gender = iota
	Male
	NonBinary
)
