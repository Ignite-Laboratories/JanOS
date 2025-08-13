package pad

import (
	"github.com/ignite-laboratories/core/enum/direction"
	"github.com/ignite-laboratories/core/enum/direction/awareness"
	"github.com/ignite-laboratories/core/enum/direction/consciousness"
	"github.com/ignite-laboratories/core/enum/direction/ordinal"
	"github.com/ignite-laboratories/core/enum/direction/orthogonal"
	"github.com/ignite-laboratories/core/enum/direction/reality"
	"github.com/ignite-laboratories/core/std/bounded"
	"github.com/ignite-laboratories/core/std/bounded/pattern"
	"math"
)


func Row[T any, TSide orthogonal.XAxis[uint], TDirection orthogonal.XAxis[uint]](size uint, buffer bounded.Pattern[T], columns ...T) []T {
	// TODO implement this
	return nil
}


func Table[T any, TSide orthogonal.XYAxis[uint], TDirection orthogonal.XYAxis[uint]](size uint, buffer bounded.Pattern[T], rows ...[]T) [][]T {
	// TODO implement this
	return nil
}

func Cube[T any, TSide orthogonal.XYAxis[uint], TDirection orthogonal.XYAxis[uint]](size uint, buffer bounded.Pattern[T], tables ...[][]T) [][][]T {
	// TODO implement this
	return nil
}

func Tesseract[T any, TSide orthogonal.XYZAxis[uint], TDirection orthogonal.XYZAxis[uint]](size uint, buffer bounded.Pattern[T], cubes ...[][][]T) [][][][]T {
	// TODO implement this
	return nil
}

func Awareness[T any, TSide direction.SpaceTime[uint], TDirection direction.SpaceTime[uint]](size uint, buffer bounded.Pattern[T], tesseracts ...[][][][]T) [][][][][]T {
	// TODO implement this
	return nil
}

func Consciousness[T any, TSideA direction.Awareness[uint], TSideB awareness.Axis[uint], TDirectionA direction.Awareness[uint], TDirectionB awareness.Axis[T]](size uint, awarenesses ...[][][][][]T) [][][][][][]T {
	// TODO implement this
	return nil
}

func Reality[T any, TSideA direction.Consciousness[uint], TSideB consciousness.Axis[uint], TDirectionA direction.Consciousness[uint], TDirectionB consciousness.Axis[T]](size uint, consciousnesses ...[][][][][][]T) [][][][][][][]T {
	// TODO implement this
	return nil
}


// Operands pads the provided operands using a pattern according to the rules defined by Type TS.
//
//  TElement - Indicates the type of elements used in the operands
//  TSide - Indicates which side of the operands to pad
//  TPatternDirection - Indicates which direction to cursor through the pattern - eastbound or westbound.
//
// See cardinal.Direction, alignment.Scheme, and scheme.Type
func Operands[TElement any, TSide orthogonal.Direction[uint], TPatternDirection direction.SpaceTime[uint]](size uint, padPattern pattern.Buffer[TElement], operands ...[]TElement) [][]TElement {
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
