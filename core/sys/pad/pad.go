package pad

import (
	"github.com/ignite-laboratories/core/enum/direction/orthogonal"
)

// String pads the provided side of a string with the value in 'toPad' to the provided side's magnitude.  If
// toPad is longer than a single character, this will repeatedly place the data on the provided side and then trim to length.
func String[T ByteOrRune, TSide orthogonal.LeftOrRight](size uint, source string, toPad string) string {
	var zero T
	switch any(zero).(type) {
	case byte:
		return string(Any1D[byte, TSide](size, []byte(source), []byte(toPad)))
	case rune:
		return string(Any1D[rune, TSide](size, []rune(source), []rune(toPad)))
	default:
		panic("invalid type - this function only supports byte or rune")
	}
}

func Any1D[T any, TSide orthogonal.LeftOrRight](totalWidth uint, source []T, toPad []T) []T {
	if len(toPad) == 0 {
		panic("cannot pad without data to pad with")
	}
	var side TSide
	for uint(len(source)) < totalWidth {
		switch any(side).(type) {
		case orthogonal.Left:
			source = append(toPad, source...)
			length := uint(len(source))
			if length > totalWidth {
				source = source[length-totalWidth:]
			}
		case orthogonal.Right:
			source = append(source, toPad...)
			if uint(len(source)) > totalWidth {
				source = source[:totalWidth]
			}
		default:
			panic("invalid side - this function only supports orthogonal Left or Right")
		}
	}

	return source
}
