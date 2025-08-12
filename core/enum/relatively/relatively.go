// Package relatively provides access to the relatively.Relative enumeration.
package relatively

// Relative represents the abstract logical relationship of two entities, 𝑎 and 𝑏.
//
// Rather than imbuing 'size', 'value', or 'position', Relatively aims to describe that '𝑎' has
// a logical relationship with '𝑏' that's understood contextually by the caller.  Whether
// in an ordered list, comparing physical dimensions, or relational timing - this provides
// a common language for describing the relationship between both entities.
//
// These terms have been very carefully chosen for their linguistic fluidity in code, while
// maintaining the existing convention of representing -1, 0, and 1.
//
// For perspective, I use the following operators:
//
//	𝑎 ⇝ 𝑏 - a's perspective of its abstract relationship to b
//	𝑎 ⇜ 𝑏 - b's perspective of its abstract relationship to a
//
// For example -
//
//	let 𝑎 = anything
//	let 𝑏 = anything
//	let 𝑎𝑏 = 𝑎 ⇝ 𝑏
//	let 𝑏𝑎 = 𝑎 ⇜ 𝑏
//	      ...
//	if 𝑎𝑏 == relatively.Aligned { ... } // Is 𝑎 relatively aligned with 𝑏?
//	if 𝑎𝑏 == relatively.Before { ... }  // Is 𝑎 relatively before 𝑏?
//	if 𝑏𝑎 > relatively.Aligned { ... }  // Has 𝑏 crossed beyond relative alignment to 𝑎?
//	if 𝑏 != relatively.After { ... }    // Has 𝑏 not yet crossed a threshold?
//	      etc...
//
// See Relative, Before, Aligned, and After
type Relative int8

const (
	// Before represents the ordinal Relative of "i-1"
	//
	// See Relative, Before, Aligned, and After
	Before Relative = -1

	// Aligned represents the ordinal Relative of "i"
	//
	// See Relative, Before, Aligned, and After
	Aligned Relative = 0

	// After represents the ordinal Relative of "i+1"
	//
	// See Relative, Before, Aligned, and After
	After Relative = 1
)

func (d Relative) String() string {
	switch d {
	case Before:
		return "←"
	case Aligned:
		return "X"
	case After:
		return "→"
	default:
		return "Unknown"
	}
}

// StringFull prints an uppercase full word representation of the Relative.
//
// You may optionally pass true for a lowercase representation.
func (d Relative) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	switch d {
	case Before:
		if lower {
			return "before"
		}
		return "Before"
	case Aligned:
		if lower {
			return "aligned"
		}
		return "Aligned"
	case After:
		if lower {
			return "after"
		}
		return "After"
	default:
		if lower {
			return "unknown"
		}
		return "Unknown"
	}
}
