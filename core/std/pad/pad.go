package pad

import (
	"github.com/ignite-laboratories/core/enum/direction"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/pattern"
	"math"
)

// Operands pads the provided operands using a pattern according to the rules defined by Scheme T.
//
// NOTE: If no direction is provided, direction.East is used.
// NOTE: This will panic if provided a direction other than direction.East or direction.West
func Operands[T Scheme](left []T, right []T, padPattern std.Pattern[T], direction ...direction.Direction) (l []T, r []T) {
	p := pattern.Zero[T]()
	if len(padPattern) > 0 {
		p = padPattern[0]
	}

	lengthLeft := len(left)
	lengthRight := len(right)
	length := lengthLeft + lengthRight
	lMinusR := lengthLeft-lengthRight
	rMinusL := lengthRight-lengthLeft

	switch any(T{}).(type) {
	case Left:
	case Right:
	case Middle:
	case Truncate:
		smallest := int(math.Min(float64(lengthLeft), float64(lengthRight)))
		return left[:smallest], right[:smallest]
	case TruncateLeftOperand:
		if lMinusR >= 0 {
			return left[:lengthRight], right
		}
		// TODO: pad
		padding := make([]T, -lMinusR)
		for i := 0; i < -lMinusR; i++ {
			padding[i] = p.
		}
	case TruncateRightOperand:
		if rMinusL >= 0 {
			return left, right[:lengthRight]
		}
		// TODO: pad
	case WalkEast:
	case WalkWest:
	case RepeatEast:
	case RepeatWest:
	default:
		panic("invalid pad scheme")
	}
	return left, right
}
