package pad

import (
	"github.com/ignite-laboratories/core/enum/direction/orthogonal"
	"github.com/ignite-laboratories/core/std/num"
	"github.com/ignite-laboratories/core/sys/pad/scheme"
	"golang.org/x/exp/slices"
)

// Scheme represents how to apply padding information against the operand.  You may either have it reflected, tiled, or
// randomized.
//
// Padding operations can be applied to ANY dimension - but each has a 'left' and 'right' side, represented by indices
// '0' and '𝑛', respectively.  When applying the data, one may want to 'tile' it in the order it's presented or 'reflect'
// it before padding.  This allows padding operations on the 'left' side to walk in the negative direction of travel, if
// desired.  If you wish to 'interleave' the padding information with the source data, provide a direction of orthogonal.Static.
//
// For example:
//
//	  directional "side" ⬎           result width ⬎     ⬐ source data       ⬐ output
//	pad.String[orthogonal.Left, scheme.Reverse]  (10, "11111", "ABC") // BACBA11111  (CBA CBA CBA CBA)
//	pad.String[orthogonal.Left, scheme.Tile]     (10, "11111", "ABC") // BCABC11111  (ABC ABC ABC ABC)
//	pad.String[orthogonal.Left, scheme.Randomize](10, "11111", "ABC") // BABCA11111  (BAC CAB CBA BCA)
//	                      padding scheme ⬏              pattern ⬏          implied pattern ⬏
//
//	pad.String[orthogonal.Right, scheme.Reverse]  (10, "11111", "ABC") // 11111CBACB  (CBA CBA CBA CBA)
//	pad.String[orthogonal.Right, scheme.Tile]     (10, "11111", "ABC") // 11111ABCAB  (ABC ABC ABC ABC)
//	pad.String[orthogonal.Right, scheme.Randomize](10, "11111", "ABC") // 11111BACBC  (BAC BCA ACB CAB)
//
//	pad.String[orthogonal.Static, scheme.Reverse]  (10, "11111", "ABC") // 1B1A1C1B1A (CBA CBA CBA CBA)
//	pad.String[orthogonal.Static, scheme.Tile]     (10, "11111", "ABC") // A1B1C1A1B1 (ABC ABC ABC ABC)
//	pad.String[orthogonal.Static, scheme.Randomize](10, "11111", "ABC") // 1A1B1A1C1B (BAC CBA CAB ACB)
//
// The process of reflecting pad data like this is called 'symmetric padding.'
//
// NOTE: All random operations in JanOS are 'periodically random'.  This means that the set of available values
// will be exhausted before another round of randomness begins, ensuring you never get repetition within one
// periodic cycle of randomness.
type Scheme interface {
	scheme.Reverse | scheme.Tile | scheme.Randomize
}

// String pads the provided side of a string with the value in toPad to totalWidth, according to the provided Scheme and direction.
func String[T orthogonal.LeftOrRight, TScheme Scheme](totalWidth uint, source string, toPad string) string {
	if len(toPad) == 0 {
		panic("cannot pad without data to pad with")
	}

	return string(Any1D[rune, T, TScheme](totalWidth, []rune(source), func() []rune { return []rune(toPad) }))
}

// FixedPattern pads the provided side of the source slice with the value in toPad to totalWidth, according to the provided Scheme and direction.
func FixedPattern[T any, TSide orthogonal.LeftOrRight, TScheme Scheme](totalWidth uint, source []T, toPad []T) []T {
	if len(toPad) == 0 {
		panic("cannot pad without data to pad with")
	}

	return Any1D[T, TSide, TScheme](totalWidth, source, func() []T { return toPad })
}

// Any1D pads the provided side of the source data using a function that provides at least one element to pad the data with.
// The padding information can be applied in multiple different ways - see Scheme.
//
// Truncation -
//
// If padding the left side, the data is 'pinned' to the right and the left elements are trimmed
//
// If padding the right side, the data is 'pinned' to the left and the right elements are trimmed
//
// If padding statically, the data is 'pinned' to the middle and both sides are trimmed as equally possible
func Any1D[T any, TSide orthogonal.LeftOrRight, TScheme Scheme](totalWidth uint, source []T, padFn func() []T) []T {
	fn := func() []T {
		var s TScheme
		switch any(s).(type) {
		case scheme.Reverse:
			toPad := slices.Clone(padFn())
			slices.Reverse(toPad)
			if len(toPad) == 0 {
				panic("cannot pad without data to pad with")
			}
			return toPad
		case scheme.Tile:
			toPad := slices.Clone(padFn())
			if len(toPad) == 0 {
				panic("cannot pad without data to pad with")
			}
			return toPad
		case scheme.Randomize:
			toPad := slices.Clone(padFn())
			toPad = num.RandomizeSet(toPad)
			if len(toPad) == 0 {
				panic("cannot pad without data to pad with")
			}
			return toPad
		default:
			panic("invalid scheme - this function only supports scheme.Reverse, scheme.Tile, or scheme.Randomize")
		}
	}

	width := int(totalWidth)
	var side TSide
	switch any(side).(type) {
	case orthogonal.Left:
		if width < len(source) {
			// Truncate by 'pinning' to the right and trimming the left
			delta := len(source) - width
			return source[delta:]
		}

		for len(source) < width {
			toPad := fn()
			source = append(toPad, source...)
		}
		source = source[len(source)-width:]
	case orthogonal.Right:
		if width < len(source) {
			// Truncate by 'pinning' to the left and trimming the right
			return source[:width]
		}

		for len(source) < width {
			toPad := fn()
			source = append(source, toPad...)
		}
		source = source[:width]
	case orthogonal.Static:
		if width < len(source) {
			// Truncate y 'pinning' to the middle and trimming as equally possible on both sides
			delta := len(source) - width
			left := delta / 2
			right := delta - left
			return source[left : len(source)-right]
		}
	default:
		panic("invalid side - this function only supports orthogonal Left or Right")
	}

	return source
}

func interleave[T any](source []T, pad []T) []T {

}
