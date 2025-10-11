package given

import (
	"fmt"

	"git.ignitelabs.net/janos/core/enum/italic"
)

// Given represents a name, as well as its original cultural meaning.  This can be used as a way to assign logical names
// to entities as needed.  Bear in mind, however, that legible names like 'Johnny-5' or 'Α' and 'Ω' are FAR easier
// to track in the trace output than '4DCE9-A9' =)
//
// JanOS, by default, seeds all entities with a random cultural descriptive name. Your interpretation and meaning of its
// origin and meanings may absolutely vary. The true beauty of language is in such prismatic interpretations based entirely
// upon contextual experiences <3
//
//	tl;dr - you own your identifier, not the other way around!
//
// See Name, String, StringQuoted, and Italicize
type Given struct {
	Name        string
	Description string
	Heritage    Heritage
}

// String returns the Name as a string.
//
// See Name, String, StringQuoted, and Italicize
func (n Given) String() string {
	return fmt.Sprintf("%v", n.Name)
}

// StringQuoted returns the Name and, if present, Description with both parts wrapped in double-quotes.
//
// For Example:
//
//	With:    "Jane Doe" - "Tenacious Mystery"
//	Without: "Kurt Weller"
//
// See Name, String, StringQuoted, and Italicize
func (n Given) StringQuoted(wrapDescription ...bool) string {
	if n.Description == "" {
		return fmt.Sprintf("\"%v\"", n.Name)
	}
	if len(wrapDescription) > 0 && wrapDescription[0] {
		return fmt.Sprintf("\"%v\" - \"%v\"", n.Name, n.Description)
	}
	return fmt.Sprintf("\"%v\" - %v", n.Name, n.Description)
}

// Italicize walks the Name's runes and changes each English alphabetic character to their italicized mathematical variant,
// leaving the other characters alone.
//
// See italic.Italicize.
func (n Given) Italicize() Given {
	n.Name = italic.Italicize(n.Name)
	return n
}
