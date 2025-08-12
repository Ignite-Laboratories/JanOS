package pad

import (
	"github.com/ignite-laboratories/core/enum/direction/cardinal"
	"github.com/ignite-laboratories/core/enum/direction/orthogonal"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/pattern"
	"math"
)

// Operands pads the provided operands using a pattern according to the rules defined by Type TS.
//
//  TElement - Indicates the type of elements used in the operands
//  TSide - Indicates which side of the operands to pad
//  TPatternDirection - Indicates which direction to cursor through the pattern - eastbound or westbound.
//
// See cardinal.Direction, alignment.Scheme, and scheme.Type
func Operands[TElement any, TSide orthogonal.Direction[uint], TPatternDirection cardinal.Longitudinal[uint]](size uint, padPattern std.Pattern[TElement], operands ...[]TElement) [][]TElement {
	p := pattern.Zero[TElement]()
	if len(padPattern) > 0 {
		p = padPattern[0]
	}

	lengthLeft := len(left)
	lengthRight := len(right)
	length := lengthLeft + lengthRight
	lMinusR := lengthLeft-lengthRight
	rMinusL := lengthRight-lengthLeft

	switch any(TS{}).(type) {
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
		padding := make([]TElement, -lMinusR)
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
