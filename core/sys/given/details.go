package given

import (
	"core/enum/gender"
	"fmt"
)

// Details represents the cultural origin and the culture's implied gender - this is purely for informational purposes.
type Details struct {
	Origin string
	Gender gender.Gender
}

func (d Details) String() string {
	switch d.Gender {
	case gender.Male:
		return fmt.Sprintf("%v - Male", d.Origin)
	case gender.Female:
		return fmt.Sprintf("%v - Female", d.Origin)
	case gender.NonBinary:
		return fmt.Sprintf("%v - Non-binary", d.Origin)
	default:
		return fmt.Sprintf("%v", d.Origin)
	}
}
