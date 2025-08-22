package pad

import (
	"github.com/ignite-laboratories/core/enum/direction/orthogonal"
	"github.com/ignite-laboratories/core/sys/pad/scheme"
	"golang.org/x/exp/slices"
)

// Scheme represents how to apply padding information against the operand, either using scheme.Reflect or scheme.Tile.
//
// Padding operations can be applied to ANY dimension - but each has a 'left' and 'right' side, represented by indices
// '0' and '𝑛', respectively.  When applying the data, one may want to 'tile' it in the order it's presented or 'reflect'
// it before padding.  This allows padding operations on the 'left' side to walk in the negative direction of travel, if
// desired.
//
// For example:
//
//	                                                            padded result ⬎              ⬐ implied pattern
//		pad.String[orthogonal.Left, scheme.Reflect]  (10, "11111", "ABC") → "BACBA11111" // CBA CBA CBA CBA |
//		pad.String[orthogonal.Left, scheme.Tile]     (10, "11111", "ABC") → "BCABC11111" // ABC ABC ABC ABC |
//		pad.String[orthogonal.Right, scheme.Reflect] (10, "11111", "ABC") → "11111CBACB" // | CBA CBA CBA CBA
//		pad.String[orthogonal.Right, scheme.Tile]    (10, "11111", "ABC") → "11111ABCAB" // | ACB ACB ACB ACB
//
// The process of reflecting pad data like this is called 'symmetric padding.'
type Scheme interface {
	scheme.Reflect | scheme.Tile
}

// String pads the provided side of a string with the value in toPad to totalWidth, according to the provided Scheme.
func String[T orthogonal.LeftOrRight, TScheme Scheme](totalWidth uint, source string, toPad string) string {
	if len(toPad) == 0 {
		panic("cannot pad without data to pad with")
	}

	return string(Any1D[rune, T, TScheme](totalWidth, []rune(source), func() []rune { return []rune(toPad) }))
}

// Any1D pads the provided side of the source data using a function that provides at least one element to pad the data with.
// The padding information can be applied in multiple different ways - see Scheme.
func Any1D[T any, TSide orthogonal.LeftOrRight, TScheme Scheme](totalWidth uint, source []T, padFn func() []T) []T {
	fn := func() []T {
		var s TScheme
		switch any(s).(type) {
		case scheme.Reflect:
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
		default:
			panic("invalid scheme - this function only supports scheme.Reflect or scheme.Tile")
		}
	}

	width := int(totalWidth)
	var side TSide
	switch any(side).(type) {
	case orthogonal.Left:
		if width < len(source) {
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
			return source[:width]
		}

		for len(source) < width {
			toPad := fn()
			source = append(source, toPad...)
		}
		source = source[:width]
	default:
		panic("invalid side - this function only supports orthogonal Left or Right")
	}

	return source
}
