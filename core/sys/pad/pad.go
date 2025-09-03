package pad

import (
	"core/enum/direction/ordinal"
	"core/sys/pad/pattern"
	"core/sys/pad/scheme"
	"core/sys/support"
	"slices"
)

// ByteOrRune represents either a byte or a rune.
type ByteOrRune interface {
	byte | rune
}

// String pads the provided side of a string with the value in toPad to totalWidth, according to the provided Scheme and direction.
// If you'd like the data to always be added and the existing data 'pushed' out, pass 'true' to the roll parameter.
//
// NOTE: This will either pad as a byte or a rune, which you must provide.
func String[T ByteOrRune](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source string, toPad string, roll ...bool) string {
	if len(toPad) == 0 {
		panic("cannot pad without data to pad with")
	}

	switch any(T(0)).(type) {
	case byte:
		return string(UsingPatternOLD(patternScheme, side, totalWidth, []byte(source), pattern.Fixed([]byte(toPad)...), roll...))
	case rune:
		return string(UsingPatternOLD(patternScheme, side, totalWidth, []rune(source), pattern.Fixed([]rune(toPad)...), roll...))
	default:
		panic("invalid type - this function only supports byte or rune")
	}
}

func RollOnto[T any](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, count uint, source []T, pattern func(uint) T) []T {
	// TODO: Truncate off the desired positions and then using a pattern scheme inject them in
	return source
}

func UsingPattern[T any](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []T, pattern func(uint) T) []T {
	source = slices.Clone(source)
	out := make([]T, totalWidth)
	slots := make([]byte, totalWidth)

	pattern(0)
	pattern(1)
	pattern(2)

	return source

	switch side {
	case ordinal.Negative:
		srcLen := len(source)
		ii := 0
		for i := len(out) - 1; i > 0; i-- {
			out[i] = source[srcLen-ii]
			slots[i] = 1
			ii++
		}

		out = fillNegatively(patternScheme, side, totalWidth, out, slots, pattern)
	case ordinal.Static:
	case ordinal.Positive:
		for i := 0; i < int(totalWidth); i++ {
			out[i] = source[i]
			slots[i] = 1
		}

		out = fillNegatively(patternScheme, side, totalWidth, out, slots, pattern)
	default:
		panic("invalid ordinal side")
	}

	//switch patternScheme {
	//case scheme.Tile:
	//case scheme.Reverse:
	//case scheme.Shuffle:
	//case scheme.ReflectInward:
	//case scheme.ReflectOutward:
	//default:
	//	panic("invalid pattern scheme")
	//}
	return out
}

func fillNegatively[T any](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []T, slots []byte, pattern func(uint) T) []T {
	return source
}

func fillStatically[T any](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []T, slots []byte, patternFill bool, pattern func(uint) T) []T {
	return source
}

func fillPositively[T any](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []T, slots []byte, pattern func(uint) T) []T {
	return source
}

// UsingPattern pads the provided side of the source data using a pattern function.  A pattern function yields a single padding element
// per call - allowing it to be dynamically generated.  If you have static data to apply, you may use FixedPatternFn to create a standard
// pattern function.  The resulting pattern elements can be manipulated in multiple different ways - see Scheme.
//
// Padding operates in two modes: pad and roll.  In the default 'pad mode', data is simply grown until it reaches totalWidth and then returned.  In
// rolling mode, exactly one pattern element is always added to the desired side while trimming the opposite - creating a "rolling buffer".
//
// If padding a 'static' side, the elements are interleaved into the existing data through distribution across the desired width.  When 'rolling',
// the pattern elements are added halfway into the data and both sides equally trimmed to length.
//
// Truncation Logic -
//
// NOTE: All dimensions have a 'left', 'static', and 'right' side.  See ordinal.Direction
//
//	If padding the left side, the resulting data is 'pinned' to the right and the left elements get trimmed.
//	If rolling onto the left side, the resulting data is 'pinned' to the left and the right elements get trimmed.
//
//	If padding the right side, the resulting data is 'pinned' to the left and the right elements get trimmed.
//	If rolling onto the right side, the resulting data is 'pinned' to the right and the left elements get trimmed.
//
//	If padding or rolling statically, the resulting data is 'pinned' to the middle and both sides trimmed as equally as possible.
//
// NOTE: This will panic if the padFn does not return the requested number of elements.
func UsingPatternOLD[T any](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []T, pattern func() T, rolling ...bool) []T {
	roll := len(rolling) > 0 && rolling[0]

	// Step 0 - bail out early if nothing should be done
	if !roll && totalWidth == uint(len(source)) {
		return source
	}

	// Step 1 - grab the padding elements and apply the padding scheme
	fn := func(n int) []T {
		if n <= 0 {
			return []T{}
		}

		toPad := make([]T, n)
		for i := 0; i < n; i++ {
			toPad[i] = pattern()
		}

		switch patternScheme {
		case scheme.Reverse:
			slices.Reverse(toPad)
			return toPad
		case scheme.Tile:
			return toPad
		case scheme.Shuffle:
			return support.ShuffleSet(toPad)
		case scheme.ReflectInward:
			out := make([]T, len(toPad))
			ltr := 0
			rtl := len(toPad) - 1
			toggler := false
			for i := 0; i < n; i++ {
				if !toggler {
					out[ltr] = toPad[i]
					ltr++
				} else {
					out[rtl] = toPad[i]
					rtl--
				}

				toggler = !toggler
			}
			return out
		case scheme.ReflectOutward:
			out := make([]T, len(toPad))
			midToLeft := (len(toPad) / 2) - 1
			midToRight := len(toPad) / 2

			if midToLeft < 0 {
				midToLeft = 0
			}

			toggler := false
			for i := 0; i < n; i++ {
				if !toggler {
					out[midToLeft] = toPad[i]
					midToLeft--

					if midToLeft < 0 {
						midToLeft = len(toPad) - 1
					}
				} else {
					out[midToRight] = toPad[i]
					midToRight++
				}

				toggler = !toggler
			}
			return out
		default:
			panic("invalid scheme - this function only supports scheme.Reverse, scheme.Tile, or scheme.Shuffle")
		}
	}

	// Step 3 - pad the source data
	width := int(totalWidth)
	switch side {
	case ordinal.Negative:
		count := width - len(source)
		if roll {
			count = 1
		}

		if count > 0 {
			source = append(fn(count), source...)
		}

		if roll {
			// Truncate the opposite side we are adding data to
			source = source[:width]
		} else {
			// Truncate by 'pinning' to the right and trimming the left
			source = source[len(source)-width:]
		}
	case ordinal.Positive:
		count := width - len(source)
		if roll {
			count = 1
		}

		source = append(source, fn(count)...)

		if roll {
			// Truncate the opposite side we are adding data to
			source = source[len(source)-width:]
		} else {
			// Truncate by 'pinning' to the left and trimming the right
			source = source[:width]
		}
	case ordinal.Static:
		count := width - len(source)
		opWidth := width
		if roll && count == 0 {
			opWidth++
			count = 1
		}

		if count > 0 {
			toPad := fn(count)

			if len(source) < len(toPad) {
				source, toPad = toPad, source
			}

			out := make([]T, opWidth)
			filled := make(map[int]struct{})
			spacing := float64(opWidth) / float64(len(toPad)+1)

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
			for i := 0; i < opWidth; i++ {
				if _, ok := filled[i]; ok {
					continue
				}
				out[i] = source[ii]
				ii++
			}
			source = out
		}

		// Truncate by 'pinning' to the middle and trimming as equally possible on both sides
		delta := len(source) - width
		left := delta / 2
		right := delta - left
		source = source[left : len(source)-right]
	default:
		panic("invalid side - this function only supports the ordinal values 'Negative', 'Positive', or 'Static'")
	}

	return source
}
