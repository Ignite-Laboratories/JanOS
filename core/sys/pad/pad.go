package pad

import (
	"github.com/ignite-laboratories/core/enum/direction/orthogonal"
)

// String pads the provided side of a string with the value in 'toPad' to the provided side's magnitude.  If
// toPad is longer than a single character, this will repeatedly place the data on the provided side and then trim to length.
func String[TSide orthogonal.LeftOrRight](size uint, source string, toPad string) string {
	if len(toPad) == 0 {
		panic("cannot pad without data to pad with")
	}

	var side TSide
	for uint(len(source)) < size {
		switch any(side).(type) {
		case orthogonal.Left:
			source = toPad + source
			length := uint(len(source))
			if length > size {
				source = source[length-size:]
			}
		case orthogonal.Right:
			source = source + toPad
			if uint(len(source)) > size {
				source = source[:size]
			}
		default:
			panic("invalid side - this function only supports orthogonal Left or Right")
		}
	}

	return source
}
