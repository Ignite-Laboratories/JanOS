package pad

import (
	"github.com/ignite-laboratories/core/enum/direction"
	"github.com/ignite-laboratories/core/enum/direction/ordinal"
	"github.com/ignite-laboratories/core/enum/direction/orthogonal"
	"github.com/ignite-laboratories/core/std"
	pattern2 "github.com/ignite-laboratories/core/std/pattern"
	"math"
)

// Row pads the row of 1D data on the provided TSide by driving a std.Cursor through the same dimension of the pattern buffer in TDirection until reaching the requested `size`.
func Row[T any, TSide orthogonal.XAxis[uint], TDirection ordinal.Axis[uint]](size uint, buffer std.Pattern[T], data ...T) []T {
	// TODO implement this
	return nil
}

// Plane pads the plane of 2D data on the provided TSide by driving a std.Cursor through the same dimension of the pattern buffer in TDirection until reaching the requested `size`.
func Plane[T any, TSide orthogonal.XYAxis[uint], TDirection ordinal.Axis[uint]](size uint, buffer std.Pattern2D[T], data ...[]T) [][]T {
	// TODO implement this
	return nil
}

// Cube pads the cube of 3D data on the provided TSide by driving a std.Cursor through the same dimension of the pattern buffer in TDirection until reaching the requested `size`.
func Cube[T any, TSide orthogonal.XYZAxis[uint], TDirection ordinal.Axis[uint]](size uint, buffer std.Pattern3D[T], data ...[][]T) [][][]T {
	// TODO implement this
	return nil
}

// Tesseract pads the tesseract of 4D data on the provided TSide by driving a std.Cursor through the same dimension of the pattern buffer in TDirection until reaching the requested `size`.
func Tesseract[T any, TSide direction.SpaceTime[uint], TDirection ordinal.Axis[uint]](size uint, buffer std.Pattern4D[T], data ...[][][]T) [][][][]T {
	// TODO implement this
	return nil
}

// Awareness pads the awareness of 5D data on the provided TSide by driving a std.Cursor through the same dimension of the pattern buffer in TDirection until reaching the requested `size`.
func Awareness[T any, TSide direction.Awareness[uint], TDirection ordinal.Axis[uint]](size uint, buffer std.Pattern5D[T], data ...[][][][]T) [][][][][]T {
	// TODO implement this
	return nil
}

// Consciousness pads the consciousness of 6D data on the provided TSide by driving a std.Cursor through the same dimension of the pattern buffer in TDirection until reaching the requested `size`.
func Consciousness[T any, TSide direction.Consciousness[uint], TDirection ordinal.Axis[uint]](size uint, buffer std.Pattern6D[T], data ...[][][][][]T) [][][][][][]T {
	// TODO implement this
	return nil
}

// Reality pads the reality of 7D data on the provided TSide by driving a std.Cursor through the same dimension of the pattern buffer in TDirection until reaching the requested `size`.
func Reality[T any, TSide direction.Reality[uint], TDirection ordinal.Axis[uint]](size uint, buffer std.Pattern7D[T], data ...[][][][][][]T) [][][][][][][]T {
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
func Operands[TElement any, TSide orthogonal.Direction[uint], TPatternDirection direction.SpaceTime[uint]](size uint, padPattern pattern2.Buffer[TElement], operands ...[]TElement) [][]TElement {
	p := pattern2.Zero[TElement]()
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
