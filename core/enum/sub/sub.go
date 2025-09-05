package sub

import "core/sys/num"

// A SubByte is an implied integer of a non-standard bit width.
//
// NOTE: These originally only defined sub-byte unsigned integers, but evolved to include a few extras between a byte and uint64.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
type SubByte uint

// A SubByte is an implied integer of a non-standard bit width.
//
// NOTE: These originally only defined sub-byte unsigned integers, but evolved to include a few extras between a byte and uint64.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
const BitMax SubByte = 1

// A SubByte is an implied integer of a non-standard bit width.
//
// NOTE: These originally only defined sub-byte unsigned integers, but evolved to include a few extras between a byte and uint64.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
const CrumbMax SubByte = 1<<2 - 1

// A SubByte is an implied integer of a non-standard bit width.
//
// NOTE: These originally only defined sub-byte unsigned integers, but evolved to include a few extras between a byte and uint64.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
const NoteMax SubByte = 1<<3 - 1

// A SubByte is an implied integer of a non-standard bit width.
//
// NOTE: These originally only defined sub-byte unsigned integers, but evolved to include a few extras between a byte and uint64.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
const NibbleMax SubByte = 1<<4 - 1

// A SubByte is an implied integer of a non-standard bit width.
//
// NOTE: These originally only defined sub-byte unsigned integers, but evolved to include a few extras between a byte and uint64.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
const FlakeMax SubByte = 1<<5 - 1

// A SubByte is an implied integer of a non-standard bit width.
//
// NOTE: These originally only defined sub-byte unsigned integers, but evolved to include a few extras between a byte and uint64.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
const MorselMax SubByte = 1<<6 - 1

// A SubByte is an implied integer of a non-standard bit width.
//
// NOTE: These originally only defined sub-byte unsigned integers, but evolved to include a few extras between a byte and uint64.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
const ShredMax SubByte = 1<<7 - 1

// A SubByte is an implied integer of a non-standard bit width.
//
// NOTE: These originally only defined sub-byte unsigned integers, but evolved to include a few extras between a byte and uint64.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
const ByteMax SubByte = 1<<8 - 1

// A SubByte is an implied integer of a non-standard bit width.
//
// NOTE: These originally only defined sub-byte unsigned integers, but evolved to include a few extras between a byte and uint64.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
const RunMax SubByte = 1<<10 - 1

// A SubByte is an implied integer of a non-standard bit width.
//
// NOTE: These originally only defined sub-byte unsigned integers, but evolved to include a few extras between a byte and uint64.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
const ScaleMax SubByte = 1<<12 - 1

// A SubByte is an implied integer of a non-standard bit width.
//
// NOTE: These originally only defined sub-byte unsigned integers, but evolved to include a few extras between a byte and uint64.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
const RiffMax SubByte = 1<<24 - 1

// A SubByte is an implied integer of a non-standard bit width.
//
// NOTE: These originally only defined sub-byte unsigned integers, but evolved to include a few extras between a byte and uint64.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
const HookMax SubByte = 1<<48 - 1

// NewBit returns a num.Numeric[uint] bounded in the closed interval [0, 1]
//
// NOTE: If no value is provided, this returns a random value in the bounded range.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
func NewBit(value ...uint) num.Numeric[uint] { return newSubByte(BitMax, value...) }

// NewCrumb returns a num.Numeric[uint] bounded in the closed interval [0, 2² - 1]
//
// NOTE: If no value is provided, this returns a random value in the bounded range.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
func NewCrumb(value ...uint) num.Numeric[uint] { return newSubByte(CrumbMax, value...) }

// NewNote returns a num.Numeric[uint] bounded in the closed interval [0, 2³ - 1]
//
// NOTE: If no value is provided, this returns a random value in the bounded range.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
func NewNote(value ...uint) num.Numeric[uint] { return newSubByte(NoteMax, value...) }

// NewNibble returns a num.Numeric[uint] bounded in the closed interval [0, 2⁴ - 1]
//
// NOTE: If no value is provided, this returns a random value in the bounded range.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
func NewNibble(value ...uint) num.Numeric[uint] { return newSubByte(NibbleMax, value...) }

// NewFlake returns a num.Numeric[uint] bounded in the closed interval [0, 2⁵ - 1]
//
// NOTE: If no value is provided, this returns a random value in the bounded range.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
func NewFlake(value ...uint) num.Numeric[uint] { return newSubByte(FlakeMax, value...) }

// NewMorsel returns a num.Numeric[uint] bounded in the closed interval [0, 2⁶ - 1]
//
// NOTE: If no value is provided, this returns a random value in the bounded range.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
func NewMorsel(value ...uint) num.Numeric[uint] { return newSubByte(MorselMax, value...) }

// NewShred returns a num.Numeric[uint] bounded in the closed interval [0, 2⁷ - 1]
//
// NOTE: If no value is provided, this returns a random value in the bounded range.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
func NewShred(value ...uint) num.Numeric[uint] { return newSubByte(ShredMax, value...) }

// NewByte returns an unbounded num.Numeric[byte].
//
// NOTE: If no value is provided, this returns a random byte.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
func NewByte(value ...byte) num.Numeric[byte] {
	var v byte
	if len(value) > 0 {
		v = value[0]
	} else {
		v = num.RandomWithinRange[byte](0, byte(ByteMax))
	}
	n, _ := num.NewNumericBounded[byte](v, 0, byte(ByteMax))
	return n
}

// NewRun returns a num.Numeric[uint] bounded in the closed interval [0, 2¹⁰ - 1]
//
// NOTE: If no value is provided, this returns a random value in the bounded range.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
func NewRun(value ...uint) num.Numeric[uint] { return newSubByte(RunMax, value...) }

// NewScale returns a num.Numeric[uint] bounded in the closed interval [0, 2¹² - 1]
//
// NOTE: If no value is provided, this returns a random value in the bounded range.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
func NewScale(value ...uint) num.Numeric[uint] { return newSubByte(ScaleMax, value...) }

// NewRiff returns a num.Numeric[uint] bounded in the closed interval [0, 2²⁴ - 1]
//
// NOTE: If no value is provided, this returns a random value in the bounded range.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
func NewRiff(value ...uint) num.Numeric[uint] { return newSubByte(RiffMax, value...) }

// NewHook returns a num.Numeric[uint] bounded in the closed interval [0, 2⁴⁸ - 1]
//
// NOTE: If no value is provided, this returns a random value in the bounded range.
//
// "The Extended Unsigned Integer Types"
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
// See NewBit, NewCrumb, NewNote, NewNibble, NewFlake, NewMorsel, NewShred, NewByte, NewRun, NewScale, NewRiff, and NewHook
func NewHook(value ...uint) num.Numeric[uint] { return newSubByte(HookMax, value...) }

func newSubByte(maximum SubByte, value ...uint) num.Numeric[uint] {
	var v uint
	if len(value) > 0 {
		v = value[0]
	} else {
		v = num.RandomWithinRange[uint](0, uint(maximum))
	}
	n, _ := num.NewNumericBounded[uint](v, 0, uint(maximum))
	return n
}
