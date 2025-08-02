// Package relatively provides a linguistic shorthand for describing the intention of relative logical order.
//
// Relatively represents the abstract logical relationship of two entities, 𝑎 and 𝑏.
//
// Rather than imbuing 'size', 'value', or 'position', relativity aims to describe that '𝑎' has
// a logical relationship with '𝑏' that's understood contextually by the caller.  Whether
// in an ordered list, comparing physical dimensions, or relational timing - this provides
// a common language for describing the relationship between both entities.
//
// These terms have been very carefully chosen for their linguistic fluidity in code, while
// maintaining the existing convention of representing -1, 0, 1.
//
// With this, I present the two perspective operators:
//
//	𝑎 ⇝ 𝑏    "a's perspective of its abstract relationship to b"
//	𝑎 ⇜ 𝑏    "b's perspective of its abstract relationship to a"
//
// For example -
//
//	 let 𝑎 = anything
//	 let 𝑏 = anything
//	 let 𝑎𝑏 = 𝑎 ⇝ 𝑏
//	 let 𝑏𝑎 = 𝑎 ⇜ 𝑏
//	       ...
//		if 𝑎𝑏 == relatively.Aligned { ... } // Is 𝑎 relatively aligned with 𝑏?
//	 if 𝑎𝑏 == relatively.Before { ... }  // Is 𝑎 relatively before 𝑏?
//	 if 𝑏𝑎 > relatively.Aligned { ... }  // Has 𝑏 crossed beyond relative alignment to 𝑎?
//	 if 𝑏 != relatively.After { ... }   // Has 𝑏 not yet crossed a threshold?
//	       etc...
//
// See Relatively, Before, Aligned, and After
package relatively
