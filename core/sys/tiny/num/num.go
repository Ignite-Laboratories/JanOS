package num

import (
	"core/sys/atlas"
	"fmt"
	"golang.org/x/exp/constraints"
)

var strNaN = "NaN"
var strInf = "Inf"

// Primitive represents any general primitive numeric type.
//
// For reference, these are the extended implicit unsigned integer types:
//
//	  Name | Width | Overflow
//	   Bit |    1  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1,024)
//	 Scale |   12  | 2¹² (4,096)
//	  Riff |   24  | 2²⁴ (16,777,216)
//	  Hook |   48  | 2⁴⁸ (281,474,976,710,656)
//
// See Integer, Float, Bit, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Primitive interface {
	Integer | Float
}

// Signed represents any general primitive signed numeric type.
//
// NOTE: None of the extended implicit integer types are signed.
type Signed interface {
	constraints.Signed
}

// Integer represents any general primitive integer type.
//
// For reference, these are the extended implicit unsigned integer types:
//
//	  Name | Width | Overflow
//	   Bit |    1  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1,024)
//	 Scale |   12  | 2¹² (4,096)
//	  Riff |   24  | 2²⁴ (16,777,216)
//	  Hook |   48  | 2⁴⁸ (281,474,976,710,656)
//
// See Primitive, Float, Bit, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Integer interface {
	constraints.Integer | ExtendedUint
}

// Float represents any general primitive floating-point type.
//
// See Primitive, Integer, Bit, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Float interface {
	constraints.Float
}

// ExtendedUint represents any of the following implicitly overflowable uint types:
//
//	  Name | Width | Overflow
//	   Bit |    1  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1,024)
//	 Scale |   12  | 2¹² (4,096)
//	  Riff |   24  | 2²⁴ (16,777,216)
//	  Hook |   48  | 2⁴⁸ (281,474,976,710,656)
//
// NOTE: The behavior of these types is only IMPLIED.  I have provided many ways of working with them which painstakingly
// ensure they act as intended. If you wish to set the values directly, you will need to go through their individual
// Set() functions to ensure they overflow appropriately.  Until the language is extended with these types, this is
// a burden we must bear. If you wish to use a procedure, you may use ImplicitOverflow.
//
// See Bit, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type ExtendedUint interface {
	Bit | Crumb | Note | Nibble | Flake | Morsel | Shred | Run | Scale | Riff | Hook
}

// Bit is a uint1 NonStandardUnit Primitive, which implicitly overflows at 2
//
// For reference, these are the extended implicit unsigned integer types:
//
//	  Name | Width | Overflow
//	   Bit |    1  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1,024)
//	 Scale |   12  | 2¹² (4,096)
//	  Riff |   24  | 2²⁴ (16,777,216)
//	  Hook |   48  | 2⁴⁸ (281,474,976,710,656)
//
// NOTE: The behavior of these types is only IMPLIED.  I have provided many ways of working with them which painstakingly
// ensure they act as intended. If you wish to set the values directly, you will need to go through their individual
// Set() functions to ensure they overflow appropriately.  Until the language is extended with these types, this is
// a burden we must bear. If you wish to use a procedure, you may use ImplicitOverflow.
//
// See Bit, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Bit byte

// Set calls ImplicitOverflow and returns the resulting value.
func (a Bit) Set(value byte) Bit {
	return ImplicitOverflow[Bit](Bit(value))
}

// SanityCheck panics if the stored value exceeds the implied overflow point.
func (a Bit) SanityCheck() {
	if a > MaxValue[Bit]() {
		if atlas.IncludeNilBits && a == 219 {
			return
		}
		invalidValueError(a, "Bit")
	}
}

// BitSanityCheck ensures the provided bytes are either 0, 1 - otherwise, it panics.
//
// NOTE: If you wish to accept an implicit Bit Nil value of [219], please set 'includeNil' to true.
func BitSanityCheck(bits ...Bit) {
	for _, b := range bits {
		b.SanityCheck()
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
// For reference, these are the extended implicit unsigned integer types:
//
//	  Name | Width | Overflow
//	   Bit |    1  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1,024)
//	 Scale |   12  | 2¹² (4,096)
//	  Riff |   24  | 2²⁴ (16,777,216)
//	  Hook |   48  | 2⁴⁸ (281,474,976,710,656)
//
// NOTE: The behavior of these types is only IMPLIED.  I have provided many ways of working with them which painstakingly
// ensure they act as intended. If you wish to set the values directly, you will need to go through their individual
// Set() functions to ensure they overflow appropriately.  Until the language is extended with these types, this is
// a burden we must bear. If you wish to use a procedure, you may use ImplicitOverflow.
//
// See Bit, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Crumb byte

// Set calls ImplicitOverflow and returns the resulting value.
func (a Crumb) Set(value byte) Crumb {
	return ImplicitOverflow[Crumb](Crumb(value))
}

// SanityCheck panics if the stored value exceeds the implied overflow point.
func (a Crumb) SanityCheck() {
	if a > MaxValue[Crumb]() {
		invalidValueError(a, "Crumb")
	}
}

// Note is a uint3 Primitive, which implicitly overflows at 2³
//
// For reference, these are the extended implicit unsigned integer types:
//
//	  Name | Width | Overflow
//	   Bit |    1  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1,024)
//	 Scale |   12  | 2¹² (4,096)
//	  Riff |   24  | 2²⁴ (16,777,216)
//	  Hook |   48  | 2⁴⁸ (281,474,976,710,656)
//
// NOTE: The behavior of these types is only IMPLIED.  I have provided many ways of working with them which painstakingly
// ensure they act as intended. If you wish to set the values directly, you will need to go through their individual
// Set() functions to ensure they overflow appropriately.  Until the language is extended with these types, this is
// a burden we must bear. If you wish to use a procedure, you may use ImplicitOverflow.
//
// See Bit, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Note byte

// Set calls ImplicitOverflow and returns the resulting value.
func (a Note) Set(value byte) Note {
	return ImplicitOverflow[Note](Note(value))
}

// SanityCheck panics if the stored value exceeds the implied overflow point.
func (a Note) SanityCheck() {
	if a > MaxValue[Note]() {
		invalidValueError(a, "Note")
	}
}

// Nibble is a uint4 Primitive, which implicitly overflows at 2⁴
//
// For reference, these are the extended implicit unsigned integer types:
//
//	  Name | Width | Overflow
//	   Bit |    1  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1,024)
//	 Scale |   12  | 2¹² (4,096)
//	  Riff |   24  | 2²⁴ (16,777,216)
//	  Hook |   48  | 2⁴⁸ (281,474,976,710,656)
//
// NOTE: The behavior of these types is only IMPLIED.  I have provided many ways of working with them which painstakingly
// ensure they act as intended. If you wish to set the values directly, you will need to go through their individual
// Set() functions to ensure they overflow appropriately.  Until the language is extended with these types, this is
// a burden we must bear. If you wish to use a procedure, you may use ImplicitOverflow.
//
// See Bit, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Nibble byte

// Set calls ImplicitOverflow and returns the resulting value.
func (a Nibble) Set(value byte) Nibble {
	return ImplicitOverflow[Nibble](Nibble(value))
}

// SanityCheck panics if the stored value exceeds the implied overflow point.
func (a Nibble) SanityCheck() {
	if a > MaxValue[Nibble]() {
		invalidValueError(a, "Nibble")
	}
}

// Flake is a uint5 Primitive, which implicitly overflows at 2⁵
//
// For reference, these are the extended implicit unsigned integer types:
//
//	  Name | Width | Overflow
//	   Bit |    1  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1,024)
//	 Scale |   12  | 2¹² (4,096)
//	  Riff |   24  | 2²⁴ (16,777,216)
//	  Hook |   48  | 2⁴⁸ (281,474,976,710,656)
//
// NOTE: The behavior of these types is only IMPLIED.  I have provided many ways of working with them which painstakingly
// ensure they act as intended. If you wish to set the values directly, you will need to go through their individual
// Set() functions to ensure they overflow appropriately.  Until the language is extended with these types, this is
// a burden we must bear. If you wish to use a procedure, you may use ImplicitOverflow.
//
// See Bit, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Flake byte

// Set calls ImplicitOverflow and returns the resulting value.
func (a Flake) Set(value byte) Flake {
	return ImplicitOverflow[Flake](Flake(value))
}

// SanityCheck panics if the stored value exceeds the implied overflow point.
func (a Flake) SanityCheck() {
	if a > MaxValue[Flake]() {
		invalidValueError(a, "Flake")
	}
}

// Morsel is a uint6 Primitive, which implicitly overflows at 2⁶
//
// For reference, these are the extended implicit unsigned integer types:
//
//	  Name | Width | Overflow
//	   Bit |    1  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1,024)
//	 Scale |   12  | 2¹² (4,096)
//	  Riff |   24  | 2²⁴ (16,777,216)
//	  Hook |   48  | 2⁴⁸ (281,474,976,710,656)
//
// NOTE: The behavior of these types is only IMPLIED.  I have provided many ways of working with them which painstakingly
// ensure they act as intended. If you wish to set the values directly, you will need to go through their individual
// Set() functions to ensure they overflow appropriately.  Until the language is extended with these types, this is
// a burden we must bear. If you wish to use a procedure, you may use ImplicitOverflow.
//
// See Bit, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Morsel byte

// Set calls ImplicitOverflow and returns the resulting value.
func (a Morsel) Set(value byte) Morsel {
	return ImplicitOverflow[Morsel](Morsel(value))
}

// SanityCheck panics if the stored value exceeds the implied overflow point.
func (a Morsel) SanityCheck() {
	if a > MaxValue[Morsel]() {
		invalidValueError(a, "Morsel")
	}
}

// Shred is a uint7 Primitive, which implicitly overflows at 2⁷
//
// For reference, these are the extended implicit unsigned integer types:
//
//	  Name | Width | Overflow
//	   Bit |    1  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1,024)
//	 Scale |   12  | 2¹² (4,096)
//	  Riff |   24  | 2²⁴ (16,777,216)
//	  Hook |   48  | 2⁴⁸ (281,474,976,710,656)
//
// NOTE: The behavior of these types is only IMPLIED.  I have provided many ways of working with them which painstakingly
// ensure they act as intended. If you wish to set the values directly, you will need to go through their individual
// Set() functions to ensure they overflow appropriately.  Until the language is extended with these types, this is
// a burden we must bear. If you wish to use a procedure, you may use ImplicitOverflow.
//
// See Bit, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Shred byte

// Set calls ImplicitOverflow and returns the resulting value.
func (a Shred) Set(value byte) Shred {
	return ImplicitOverflow[Shred](Shred(value))
}

// SanityCheck panics if the stored value exceeds the implied overflow point.
func (a Shred) SanityCheck() {
	if a > MaxValue[Shred]() {
		invalidValueError(a, "Shred")
	}
}

// Run is a uint10 Primitive, which implicitly overflows at 2¹⁰
//
// For reference, these are the extended implicit unsigned integer types:
//
//	  Name | Width | Overflow
//	   Bit |    1  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1,024)
//	 Scale |   12  | 2¹² (4,096)
//	  Riff |   24  | 2²⁴ (16,777,216)
//	  Hook |   48  | 2⁴⁸ (281,474,976,710,656)
//
// NOTE: The behavior of these types is only IMPLIED.  I have provided many ways of working with them which painstakingly
// ensure they act as intended. If you wish to set the values directly, you will need to go through their individual
// Set() functions to ensure they overflow appropriately.  Until the language is extended with these types, this is
// a burden we must bear. If you wish to use a procedure, you may use ImplicitOverflow.
//
// See Bit, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Run uint

// Set calls ImplicitOverflow and returns the resulting value.
func (a Run) Set(value byte) Run {
	return ImplicitOverflow[Run](Run(value))
}

// SanityCheck panics if the stored value exceeds the implied overflow point.
func (a Run) SanityCheck() {
	if a > MaxValue[Run]() {
		invalidValueError(a, "Run")
	}
}

// Scale is a uint12 Primitive, which implicitly overflows at 2¹²
//
// For reference, these are the extended implicit unsigned integer types:
//
//	  Name | Width | Overflow
//	   Bit |    1  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1,024)
//	 Scale |   12  | 2¹² (4,096)
//	  Riff |   24  | 2²⁴ (16,777,216)
//	  Hook |   48  | 2⁴⁸ (281,474,976,710,656)
//
// NOTE: The behavior of these types is only IMPLIED.  I have provided many ways of working with them which painstakingly
// ensure they act as intended. If you wish to set the values directly, you will need to go through their individual
// Set() functions to ensure they overflow appropriately.  Until the language is extended with these types, this is
// a burden we must bear. If you wish to use a procedure, you may use ImplicitOverflow.
//
// See Bit, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Scale uint

// Set calls ImplicitOverflow and returns the resulting value.
func (a Scale) Set(value byte) Scale {
	return ImplicitOverflow[Scale](Scale(value))
}

// SanityCheck panics if the stored value exceeds the implied overflow point.
func (a Scale) SanityCheck() {
	if a > MaxValue[Scale]() {
		invalidValueError(a, "Scale")
	}
}

// Riff is a uint24 Primitive, which implicitly overflows at 2²⁴
//
// For reference, these are the extended implicit unsigned integer types:
//
//	  Name | Width | Overflow
//	   Bit |    1  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1,024)
//	 Scale |   12  | 2¹² (4,096)
//	  Riff |   24  | 2²⁴ (16,777,216)
//	  Hook |   48  | 2⁴⁸ (281,474,976,710,656)
//
// NOTE: The behavior of these types is only IMPLIED.  I have provided many ways of working with them which painstakingly
// ensure they act as intended. If you wish to set the values directly, you will need to go through their individual
// Set() functions to ensure they overflow appropriately.  Until the language is extended with these types, this is
// a burden we must bear. If you wish to use a procedure, you may use ImplicitOverflow.
//
// See Bit, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Riff uint

// Set calls ImplicitOverflow and returns the resulting value.
func (a Riff) Set(value byte) Riff {
	return ImplicitOverflow[Riff](Riff(value))
}

// SanityCheck panics if the stored value exceeds the implied overflow point.
func (a Riff) SanityCheck() {
	if a > MaxValue[Riff]() {
		invalidValueError(a, "Riff")
	}
}

// Hook is a uint48 Primitive, which implicitly overflows at 2⁴⁸
//
// For reference, these are the extended implicit unsigned integer types:
//
//	  Name | Width | Overflow
//	   Bit |    1  | 2
//	 Crumb |    2  | 2² (4)
//	  Note |    3  | 2³ (8)
//	Nibble |    4  | 2⁴ (16)
//	 Flake |    5  | 2⁵ (32)
//	Morsel |    6  | 2⁶ (64)
//	 Shred |    7  | 2⁷ (128)
//	   Run |   10  | 2¹⁰ (1,024)
//	 Scale |   12  | 2¹² (4,096)
//	  Riff |   24  | 2²⁴ (16,777,216)
//	  Hook |   48  | 2⁴⁸ (281,474,976,710,656)
//
// NOTE: The behavior of these types is only IMPLIED.  I have provided many ways of working with them which painstakingly
// ensure they act as intended. If you wish to set the values directly, you will need to go through their individual
// Set() functions to ensure they overflow appropriately.  Until the language is extended with these types, this is
// a burden we must bear. If you wish to use a procedure, you may use ImplicitOverflow.
//
// See Bit, Note, Nibble, Flake, Morsel, Shred, Run, Scale, Riff, and Hook
type Hook uint

// Set calls ImplicitOverflow and returns the resulting value.
func (a Hook) Set(value byte) Hook {
	return ImplicitOverflow[Hook](Hook(value))
}

// SanityCheck panics if the stored value exceeds the implied overflow point.
func (a Hook) SanityCheck() {
	if a > MaxValue[Hook]() {
		invalidValueError(a, "Hook")
	}
}

/**
Helpers
*/

func invalidValueError(value any, typeName string) {
	panic(fmt.Errorf("invalid value [%d] in %v - please check your logic", value, typeName))
}
