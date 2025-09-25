package scheme

// Scheme represents how to apply pattern information against an operand.  You may either have it reflected, tiled, or
// randomized.
//
// Padding operations can be applied to ANY dimension - but each has a 'left' and 'right' side, represented by indices
// '0' and '𝑛', respectively.  When applying a pattern, one may want to 'tile' it in the order it's presented or 'reverse'
// it before padding.  This allows padding operations on the 'left' side to walk in the negative direction of travel, if
// desired.  If you wish to 'interleave' the padding information with the source data, provide a direction of orthogonal.Static.
//
// For example:
//
//	  directional "side" ⬎        result width ⬎      ⬐ source data       ⬐ output
//	pad.String[orthogonal.Left, scheme.Reverse](10, "11111", "ABC") // BACBA11111  (CBA CBA CBA CBA)
//	pad.String[orthogonal.Left, scheme.Tile]   (10, "11111", "ABC") // BCABC11111  (ABC ABC ABC ABC)
//	pad.String[orthogonal.Left, scheme.Shuffle](10, "11111", "ABC") // BABCA11111  (BAC CAB CBA BCA)
//	                     padding scheme ⬏            pattern ⬏          implied pattern ⬏
//
//	pad.String[orthogonal.Right, scheme.Reverse](10, "11111", "ABC") // 11111CBACB  (CBA CBA CBA CBA)
//	pad.String[orthogonal.Right, scheme.Tile]   (10, "11111", "ABC") // 11111ABCAB  (ABC ABC ABC ABC)
//	pad.String[orthogonal.Right, scheme.Shuffle](10, "11111", "ABC") // 11111BACBC  (BAC BCA ACB CAB)
//
//	pad.String[orthogonal.Static, scheme.Reverse](10, "11111", "ABC") // 1B1A1C1B1A (CBA CBA CBA CBA)
//	pad.String[orthogonal.Static, scheme.Tile]   (10, "11111", "ABC") // A1B1C1A1B1 (ABC ABC ABC ABC)
//	pad.String[orthogonal.Static, scheme.Shuffle](10, "11111", "ABC") // 1A1B1A1C1B (BAC CBA CAB ACB)
//
// In addition, you may use scheme.ReflectInward and scheme.ReflectOutward.  Please see their documentation for details.
//
// The process of reflecting pad data using these schemes is what I call 'symmetrical padding.'
//
// NOTE: Random operations in JanOS are written as 'periodically random'.  This means that the entire range of possible
// random values will be exhausted before another round of randomness begins, ensuring you never get repetition within
// one contextual 'cycle' of randomness.  If the possible range is larger than the number of output random values, this
// means the numbers will never repeat.
type Scheme byte

const (
	// Reverse indicates that the pattern elements should be reversed before application to the source data.
	//
	// See Reverse, Tile, Shuffle, ReflectInward, and ReflectOutward
	Reverse Scheme = iota

	// Tile indicates that the pattern elements should be cyclically repeated as-is while applying it to the source data.
	//
	// See Reverse, Tile, Shuffle, ReflectInward, and ReflectOutward
	Tile

	// Shuffle indicates that the pattern elements should be randomly re-ordered before application to the source data.
	//
	// See Reverse, Tile, Shuffle, ReflectInward, and ReflectOutward
	Shuffle

	// ReflectInward indicates that the pattern elements should be treated as walking centrally inward.
	//
	// This means the elements are ordered from ABCDE to ACEDB, where elements are walked pairwise inward.
	//
	// If statically interleaving '11111' with 'ABCDE' to 10 elements wide - A 1 C 1 E 1 D 1 B 1
	//
	// If padding the left of '11111' with 'ABCDE' to 10 elements wide - A C E D B 1 1 1 1 1
	//
	// If padding the right of '11111' with 'ABCDE' to 10 elements wide - 1 1 1 1 1 A C E D B
	//
	// See Reverse, Tile, Shuffle, ReflectInward, and ReflectOutward
	ReflectInward

	// ReflectOutward indicates that the pattern elements should be treated as walking centrally outward.
	//
	// This means the elements are ordered from ABCDE to DBACE, where A is placed in the middle before
	// elements are walked pairwise outward.
	//
	// If statically interleaving '11111' with 'ABCDE' to 10 elements wide - D 1 B 1 A 1 C 1 E 1
	//
	// If padding the left of '11111' with 'ABCDE' to 10 elements wide - D B A C E 1 1 1 1 1
	//
	// If padding the right of '11111' with 'ABCDE' to 10 elements wide - 1 1 1 1 1 D B A C E
	//
	// See Reverse, Tile, Shuffle, ReflectInward, and ReflectOutward
	ReflectOutward
)
