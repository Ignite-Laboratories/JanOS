package num

import (
	"golang.org/x/exp/constraints"
)

// Primitive represents any general primitive numeric type.
//
// See Integer, Float, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Primitive interface {
	Integer | Float
}

// Integer represents any general integer type.
//
// See Primitive, Float, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Integer interface {
	constraints.Integer | Crumb | Note | Nibble | Flake | Morsel | Shred | Run | Scale | Riff | Hook
}

// Float represents any general floating-point type.
//
// See Primitive, Integer, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Float interface {
	constraints.Float
}

// Crumb is a uint2 Primitive, which implicitly overflows at 2²
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
// See Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Crumb byte

// Note is a uint3 Primitive, which implicitly overflows at 2³
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
// See Crumb, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Note byte

// Nibble is a uint4 Primitive, which implicitly overflows at 2⁴
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
// See Crumb, Note, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Nibble byte

// Flake is a uint5 Primitive, which implicitly overflows at 2⁵
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
// See Crumb, Note, Nibble, Morsel, Shred, Run, Scale, Riff, and Hook
type Flake byte

// Morsel is a uint6 Primitive, which implicitly overflows at 2⁶
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
// See Crumb, Note, Nibble, Flake, Shred, Run, Scale, Riff, and Hook
type Morsel byte

// Shred is a uint7 Primitive, which implicitly overflows at 2⁷
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
// See Crumb, Note, Nibble, Flake, Morsel, Run, Scale, Riff, and Hook
type Shred byte

// Run is a uint10 Primitive, which implicitly overflows at 2¹⁰
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
// See Crumb, Note, Nibble, Flake, Morsel, Shred, Scale, Riff, and Hook
type Run uint

// Scale is a uint12 Primitive, which implicitly overflows at 2¹²
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
// See Crumb, Note, Nibble, Flake, Morsel, Shred, Run, Riff, and Hook
type Scale uint

// Riff is a uint24 Primitive, which implicitly overflows at 2²⁴
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
// See Crumb, Note, Nibble, Flake, Morsel, Shred, Run, Scale, and Hook
type Riff uint

// Hook is a uint48 Primitive, which implicitly overflows at 2⁴⁸
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
// See Crumb, Note, Nibble, Flake, Morsel, Shred, Run, Scale, and Riff
type Hook uint
