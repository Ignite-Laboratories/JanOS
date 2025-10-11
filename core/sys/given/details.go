package given

import (
	"fmt"

	"git.ignitelabs.net/janos/core/enum/gender"
)

// Heritage represents the cultural origin and the culture's implied gender - this is purely for informational purposes.
type Heritage struct {
	Origin string
	Gender gender.Gender
}

func (d Heritage) String() string {
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
