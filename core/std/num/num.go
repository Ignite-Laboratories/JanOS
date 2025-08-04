package num

import (
	"golang.org/x/exp/constraints"
)

// Primitive represents any primitive Go integer or floating-point type.
//
// See Float, Integer, ExtendedInteger, and ExtendedPrimitive
type Primitive interface {
	Integer | Float
}

// ExtendedInteger represents any Integer or ExtendedPrimitive types.
//
// See Primitive, Float, Integer, and ExtendedPrimitive
type ExtendedInteger interface {
	Integer | ExtendedPrimitive
}

// Integer represents any primitive Go integer type.
//
// See Primitive, Float, ExtendedInteger, and ExtendedPrimitive
type Integer interface {
	constraints.Integer
}

// Float represents any primitive Go floating-point type.
//
// See Primitive, Integer, ExtendedInteger, and ExtendedPrimitive
type Float interface {
	constraints.Float
}

// ExtendedPrimitive extends the Primitive Go numeric types with several unsigned integers which implicitly overflow
// at their implied bit width:
//
//	  Name | Width | Overflow
//	 Crumb |    2  |    2²
//	  Note |    3  |    2³
//	Nibble |    4  |    2⁴
//	 Flake |    5  |    2⁵
//	Morsel |    6  |    2⁶
//	 Shred |    7  |    2⁷
//	   Run |   10  |    2¹⁰
//	 Scale |   12  |    2¹²
//	  Riff |   24  |    2²⁴
//	  Hook |   48  |    2⁴⁸
//
// See Primitive, Float, Integer, ExtendedInteger, Crumb, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type ExtendedPrimitive interface {
	Primitive | Crumb | Note | Nibble | Flake | Morsel | Shred | Run | Scale | Riff | Hook
}

// Crumb is an uint2 ExtendedPrimitive, which implicitly overflows at 2²
//
// See Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Crumb byte

// Note is an uint3 ExtendedPrimitive, which implicitly overflows at 2³
//
// See Crumb, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Note byte

// Nibble is an uint4 ExtendedPrimitive, which implicitly overflows at 2⁴
//
// See Crumb, Note, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Nibble byte

// Flake is an uint5 ExtendedPrimitive, which implicitly overflows at 2⁵
//
// See Crumb, Note, Nibble, Morsel, Shred, Run, Scale, Riff, and Hook
type Flake byte

// Morsel is an uint6 ExtendedPrimitive, which implicitly overflows at 2⁶
//
// See Crumb, Note, Nibble, Flake, Shred, Run, Scale, Riff, and Hook
type Morsel byte

// Shred is an uint7 ExtendedPrimitive, which implicitly overflows at 2⁷
//
// See Crumb, Note, Nibble, Flake, Morsel, Run, Scale, Riff, and Hook
type Shred byte

// Run is an uint10 ExtendedPrimitive, which implicitly overflows at 2¹⁰
//
// See Crumb, Note, Nibble, Flake, Morsel, Shred, Scale, Riff, and Hook
type Run uint

// Scale is an uint12 ExtendedPrimitive, which implicitly overflows at 2¹²
//
// See Crumb, Note, Nibble, Flake, Morsel, Shred, Run, Riff, and Hook
type Scale uint

// Riff is an uint24 ExtendedPrimitive, which implicitly overflows at 2²⁴
//
// See Crumb, Note, Nibble, Flake, Morsel, Shred, Run, Scale, and Hook
type Riff uint

// Hook is an uint48 ExtendedPrimitive, which implicitly overflows at 2⁴⁸
//
// See Crumb, Note, Nibble, Flake, Morsel, Shred, Run, Scale, and Riff
type Hook uint
