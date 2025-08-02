package name

import (
	"fmt"
)

// Given represents a name, as well as its original cultural meaning.
//
// Your interpretation and meaning may absolutely vary. The true beauty of language
// is in such prismatic interpretations based entirely upon contextual experiences <3
//
//	tl;dr - you own your identifier, not the other way around!
type Given struct {
	Name        string
	Description string
	Details     Details
}

func (n Given) String() string {
	if n.Description == "" {
		return fmt.Sprintf("%v", n.Name)
	}
	return fmt.Sprintf("%v - %v", n.Name, n.Description)
}
