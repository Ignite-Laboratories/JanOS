package num

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// Numeric represents any Primitive numeric Value.
type Numeric[T Primitive] struct {
	// Value represents the current value of this Numeric Primitive.
	Value T
}

// SignedNumeric represents any primitive Signed numeric Value.
type SignedNumeric[T Signed] struct {
	// Value represents the current value of this Signed Numeric.
	Value T
}

// Primitive represents any general primitive numeric type.
//
// See Integer, Float, Bit, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Primitive interface {
	Integer | Float
}

// Signed represents any general primitive signed numeric type.
type Signed interface {
	constraints.Signed
}

// Integer represents any general primitive integer type.
//
// See Primitive, Float, Bit, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Integer interface {
	constraints.Integer | Bit | Crumb | Note | Nibble | Flake | Morsel | Shred | Run | Scale | Riff | Hook
}

// Float represents any general primitive floating-point type.
//
// See Primitive, Integer, Bit, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Float interface {
	constraints.Float
}

// Bit is a uint1 Primitive, which implicitly overflows at 2
//
//	  Name | Width | Overflow
//	   Bit |    2  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1024)
//	 Scale |   12  | 2¹² (4096)
//	  Riff |   24  | 2²⁴
//	  Hook |   48  | 2⁴⁸
//
// See Bit, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Bit byte

// SanityCheck ensures the source and provided bits are all 0 or 1.
//
// See SanityCheckWithNil
func (a Bit) SanityCheck(bits ...Bit) bool {
	for _, bit := range bits {
		if bit > Bit(MaxValue[Bit]()) {
			return false
		}
	}
	return true
}

// SanityCheckWithNil ensures the source and provided bits are all 0, 1, or 219 - which is implicitly 'nil'
//
// See SanityCheck
func (a Bit) SanityCheckWithNil(bits ...Bit) bool {
	for _, bit := range bits {
		if bit > Bit(MaxValue[Bit]()) || bit == 219 {
			return false
		}
	}
	return true
}

// BitSanityCheck ensures the provided bytes are either 0, 1, or Nil (219) - otherwise, it panics.
func BitSanityCheck(bits ...Bit) {
	for _, b := range bits {
		if b != 0 && b != 1 {
			panic(fmt.Errorf("not a bit value: %d", b))
		}
	}
}

// String converts the provided Bit to a string "1", "0", or "-" for Nil [219] and panics if the found value is anything else.
func (b Bit) String() string {
	switch b {
	case 0:
		return "0"
	case 1:
		return "1"
	case 219:
		return "-"
	default:
		panic(fmt.Errorf("not a bit value: %d", b))
	}
}

// Crumb is a uint2 Primitive, which implicitly overflows at 2²
//
//	  Name | Width | Overflow
//	   Bit |    2  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1024)
//	 Scale |   12  | 2¹² (4096)
//	  Riff |   24  | 2²⁴
//	  Hook |   48  | 2⁴⁸
//
// See Bit, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Crumb byte

func (a Crumb) SanityCheck(b Crumb) bool {
	return b <= Crumb(MaxValue[Crumb]())
}

// Note is a uint3 Primitive, which implicitly overflows at 2³
//
//	  Name | Width | Overflow
//	   Bit |    2  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1024)
//	 Scale |   12  | 2¹² (4096)
//	  Riff |   24  | 2²⁴
//	  Hook |   48  | 2⁴⁸
//
// See Bit, Crumb, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Note byte

func (a Note) SanityCheck(b Note) bool {
	return b <= Note(MaxValue[Note]())
}

// Nibble is a uint4 Primitive, which implicitly overflows at 2⁴
//
//	  Name | Width | Overflow
//	   Bit |    2  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1024)
//	 Scale |   12  | 2¹² (4096)
//	  Riff |   24  | 2²⁴
//	  Hook |   48  | 2⁴⁸
//
// See Bit, Crumb, Note, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Nibble byte

func (a Nibble) SanityCheck(b Nibble) bool {
	return b <= Nibble(MaxValue[Nibble]())
}

// Flake is a uint5 Primitive, which implicitly overflows at 2⁵
//
//	  Name | Width | Overflow
//	   Bit |    2  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1024)
//	 Scale |   12  | 2¹² (4096)
//	  Riff |   24  | 2²⁴
//	  Hook |   48  | 2⁴⁸
//
// See Bit, Crumb, Note, Nibble, Morsel, Shred, Run, Scale, Riff, and Hook
type Flake byte

func (a Flake) SanityCheck(b Flake) bool {
	return b <= Flake(MaxValue[Flake]())
}

// Morsel is a uint6 Primitive, which implicitly overflows at 2⁶
//
//	  Name | Width | Overflow
//	   Bit |    2  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1024)
//	 Scale |   12  | 2¹² (4096)
//	  Riff |   24  | 2²⁴
//	  Hook |   48  | 2⁴⁸
//
// See Bit, Crumb, Note, Nibble, Flake, Shred, Run, Scale, Riff, and Hook
type Morsel byte

func (a Morsel) SanityCheck(b Morsel) bool {
	return b <= Morsel(MaxValue[Morsel]())
}

// Shred is a uint7 Primitive, which implicitly overflows at 2⁷
//
//	  Name | Width | Overflow
//	   Bit |    2  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1024)
//	 Scale |   12  | 2¹² (4096)
//	  Riff |   24  | 2²⁴
//	  Hook |   48  | 2⁴⁸
//
// See Bit, Crumb, Note, Nibble, Flake, Morsel, Run, Scale, Riff, and Hook
type Shred byte

func (a Shred) SanityCheck(b Shred) bool {
	return b <= Shred(MaxValue[Shred]())
}

// Run is a uint10 Primitive, which implicitly overflows at 2¹⁰
//
//	  Name | Width | Overflow
//	   Bit |    2  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1024)
//	 Scale |   12  | 2¹² (4096)
//	  Riff |   24  | 2²⁴
//	  Hook |   48  | 2⁴⁸
//
// See Bit, Crumb, Note, Nibble, Flake, Morsel, Shred, Scale, Riff, and Hook
type Run uint

func (a Run) SanityCheck(b Run) bool {
	return b <= Run(MaxValue[Run]())
}

// Scale is a uint12 Primitive, which implicitly overflows at 2¹²
//
//	  Name | Width | Overflow
//	   Bit |    2  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1024)
//	 Scale |   12  | 2¹² (4096)
//	  Riff |   24  | 2²⁴
//	  Hook |   48  | 2⁴⁸
//
// See Bit, Crumb, Note, Nibble, Flake, Morsel, Shred, Run, Riff, and Hook
type Scale uint

func (a Scale) SanityCheck(b Scale) bool {
	return b <= Scale(MaxValue[Scale]())
}

// Riff is a uint24 Primitive, which implicitly overflows at 2²⁴
//
//	  Name | Width | Overflow
//	   Bit |    2  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1024)
//	 Scale |   12  | 2¹² (4096)
//	  Riff |   24  | 2²⁴
//	  Hook |   48  | 2⁴⁸
//
// See Bit, Crumb, Note, Nibble, Flake, Morsel, Shred, Run, Scale, and Hook
type Riff uint

func (a Riff) SanityCheck(b Riff) bool {
	return b <= Riff(MaxValue[Riff]())
}

// Hook is a uint48 Primitive, which implicitly overflows at 2⁴⁸
//
//	  Name | Width | Overflow
//	   Bit |    2  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1024)
//	 Scale |   12  | 2¹² (4096)
//	  Riff |   24  | 2²⁴
//	  Hook |   48  | 2⁴⁸
//
// See Bit, Crumb, Note, Nibble, Flake, Morsel, Shred, Run, Scale, and Riff
type Hook uint

func (a Hook) SanityCheck(b Hook) bool {
	return b <= Hook(MaxValue[Hook]())
}
