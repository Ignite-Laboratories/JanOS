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
//	pad.String[orthogonal.Left, scheme.Reverse](10, "11111", "ABC") // BACBA11111  (CBA CBA CBA CBA)
//	pad.String[orthogonal.Left, scheme.Tile]   (10, "11111", "ABC") // BCABC11111  (ABC ABC ABC ABC)
//	pad.String[orthogonal.Left, scheme.Shuffle](10, "11111", "ABC") // BABCA11111  (BAC CAB CBA BCA)
//	                      padding scheme ⬏              pattern ⬏          implied pattern ⬏
//
//	pad.String[orthogonal.Right, scheme.Reverse](10, "11111", "ABC") // 11111CBACB  (CBA CBA CBA CBA)
//	pad.String[orthogonal.Right, scheme.Tile]   (10, "11111", "ABC") // 11111ABCAB  (ABC ABC ABC ABC)
//	pad.String[orthogonal.Right, scheme.Shuffle](10, "11111", "ABC") // 11111BACBC  (BAC BCA ACB CAB)
//
//	pad.String[orthogonal.Static, scheme.Reverse](10, "11111", "ABC") // 1B1A1C1B1A (CBA CBA CBA CBA)
//	pad.String[orthogonal.Static, scheme.Tile]   (10, "11111", "ABC") // A1B1C1A1B1 (ABC ABC ABC ABC)
//	pad.String[orthogonal.Static, scheme.Shuffle](10, "11111", "ABC") // 1A1B1A1C1B (BAC CBA CAB ACB)
//
// The process of reflecting pad data like this is called 'symmetric padding.'
//
// NOTE: All random operations in JanOS are 'periodically random'.  This means that the set of available values
// will be exhausted before another round of randomness begins, ensuring you never get repetition within one
// periodic cycle of randomness.
type Scheme interface {
	scheme.Reverse | scheme.Tile | scheme.Shuffle
}

// String pads the provided side of a string with the value in toPad to totalWidth, according to the provided Scheme and direction.
//
// NOTE: If you'd like the data to always be added and the existing data 'pushed' out, pass 'true' to the roll parameter.
func String[T orthogonal.LeftOrRight, TScheme Scheme](totalWidth uint, source string, toPad string, roll ...bool) string {
	if len(toPad) == 0 {
		panic("cannot pad without data to pad with")
	}

	return string(AnyDimension[rune, T, TScheme](totalWidth, []rune(source), FixedPatternFn([]rune(toPad)), roll...))
}

// FixedPatternFn creates a fixed pattern function that returns the statically provided elements.
func FixedPatternFn[T any](toPad []T) func() T {
	if len(toPad) == 0 {
		panic("cannot pad without data to pad with")
	}
	out := slices.Clone(toPad)
	i := 0
	return func() T {
		element := out[i]
		i++
		if i >= len(out) {
			i = 0
		}
		return element
	}
}

// AnyDimension pads the provided side of the source data using a function that provides one element at a time to pad the data with.
// The padding information can be applied in multiple different ways - see Scheme.
//
// Padding operates in two modes: pad and roll.  In pad mode, the data is grown until it reaches totalWidth and then returned.  In rolling
// mode, exactly one element from the padFn is ALWAYS applied while "pushing" out existing data - creating a "rolling buffer".  If performing
// a rolling pad operation against a static direction, the output data is added halfway into the data and both sides are trimmed.
//
// Truncation -
//
// If padding the left side, the data is 'pinned' to the right and the left elements are trimmed
//
// If padding the right side, the data is 'pinned' to the left and the right elements are trimmed
//
// If padding statically, the data is 'pinned' to the middle and both sides trimmed as equally possible.  If 'rolling', the left is always trimmed.
//
// NOTE: This will panic if the padFn does not return the requested number of elements.
func AnyDimension[T any, TSide orthogonal.LeftOrRight, TScheme Scheme](totalWidth uint, source []T, padFn func() T, rolling ...bool) []T {
	roll := len(rolling) > 0 && rolling[0]

	// Step 0 - bail out early if nothing should be done
	if !roll && totalWidth == uint(len(source)) {
		return source
	}

	// Step 1 - grab the padding elements and apply the padding scheme
	fn := func(n uint) []T {
		count := totalWidth - uint(len(source))
		toPad := make([]T, count)
		for i := uint(0); i < count; i++ {
			toPad[i] = padFn()
		}

		var s TScheme
		switch any(s).(type) {
		case scheme.Reverse:
			slices.Reverse(toPad)
			return toPad
		case scheme.Tile:
			for i := uint(0); i < count; i++ {
				toPad[i] = padFn()
			}
			return toPad
		case scheme.Shuffle:
			for i := uint(0); i < count; i++ {
				toPad[i] = padFn()
			}
			toPad = num.ShuffleSet(toPad)
			return toPad
		default:
			panic("invalid scheme - this function only supports scheme.Reverse, scheme.Tile, or scheme.Shuffle")
		}
	}

	// Step 3 - pad the source data
	width := int(totalWidth)
	var side TSide
	switch any(side).(type) {
	case orthogonal.Left:
		count := width - len(source)
		if roll {
			count = 1
		}

		if count > 0 {
			source = append(fn(uint(count)), source...)
		}

		if roll {
			// Truncate the opposite side we are adding data to
			source = source[:width]
		} else {
			// Truncate by 'pinning' to the right and trimming the left
			source = source[len(source)-width:]
		}
	case orthogonal.Right:
		count := width - len(source)
		if roll {
			count = 1
		}

		if count > 0 {
			source = append(source, fn(uint(count))...)
		}

		if roll {
			// Truncate the opposite side we are adding data to
			source = source[len(source)-width:]
		} else {
			// Truncate by 'pinning' to the left and trimming the right
			source = source[:width]
		}
	case orthogonal.Static:
		count := width - len(source)
		if roll {
			count = 1
		}

		if count > 0 {
			toPad := fn(uint(count))
			out := make([]T, width)
			filled := make(map[int]struct{})
			spacing := float64(width) / float64(len(toPad)+1)

			offset := 1
			if spacing < 1 {
				offset = 0
				spacing = 1
			}
			for i, v := range toPad {
				pos := int(float64(i+offset) * spacing)
				out[pos] = v
				filled[pos] = struct{}{}
			}

			ii := 0
			for i := 0; i < width; i++ {
				if _, ok := filled[i]; ok {
					continue
				}
				out[i] = source[ii]
				ii++
			}
		}

		// Truncate by 'pinning' to the middle and trimming as equally possible on both sides
		delta := len(source) - width
		left := delta / 2
		right := delta - left
		source = source[left : len(source)-right]
	default:
		panic("invalid side - this function only supports orthogonal Left or Right")
	}

	return source
}
