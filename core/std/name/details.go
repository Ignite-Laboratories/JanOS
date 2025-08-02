package name

import (
	"fmt"
	"github.com/ignite-laboratories/core/enum/gender"
)

// Details represents the cultural and gender details behind a name.
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
