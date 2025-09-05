// Package endianness provides access to the Endianness enumeration.
package endian

// Endianness indicates the logical -byte- ordering of sequential bytes.  All binary data has a most significant side,
// where the binary placeholder has the highest relative value, as well as a least significant side.  The individual BITS
// of a byte are colloquially manipulated in most→to→least significant order, but multiple BYTES worth of information may
// be stored in least←to←most significant order while retaining the individual BIT order of each byte. (Say that ten times fast!)
//
// There are two types of endianness -
//
// Big, where the most significant bytes come first - or "raw" binary:
//
//	| Most Sig. Byte  |   Middle Byte   | Least Sig. Byte |
//	| 0 1 0 0 1 1 0 1 | 0 0 1 0 1 1 0 0 | 0 0 0 1 0 1 1 0 | (5,057,558)
//	|        4D       |        2C       |        16       |
//
// Little, where the least significant bytes come first - used by x86, AMD64, ARM, and the general world over:
//
//	| Least Sig. Byte |   Middle Byte   |  Most Sig. Byte |
//	| 0 0 0 1 0 1 1 0 | 0 0 1 0 1 1 0 0 | 0 1 0 0 1 1 0 1 | (5,057,558)
//	|        16       |        2C       |        4D       |
//	         ⬑  The byte's internal bits remain in most→to→least order
//
// NOTE: While SOME obscure hardware might store BITS in least←to←most significant order, Go's shift operators (<< and >>) are
// guaranteed by the language specification to always operate in most→to→least significant order. This, in turn, means that bit
// operations in tiny will -also- work with bits in most→to→least significant order - regardless of the underlying architecture's
// bit storage order. Because of this, when reading raw memory, only byte ordering needs to be handled explicitly.
//
// NOTE: Some protocols, like UART and SPI, traditionally transmit in -BITWISE- little endian order, so you may also need to reverse
// bits within bytes when interfacing with such protocols!
//
// See Endianness, Little, and Big.
type Endianness byte

const (
	// Little indicates that bytes are handled in least←to←most significant order and is used by x86, AMD64, ARM, and the general
	// world over.
	//
	// See Endianness, Little, and Big.
	Little Endianness = iota

	// Big indicates that bytes are handled in most→to→least significant order and represents "raw" binary - it's often favored by network protocols.
	//
	// See Endianness, Little, and Big.
	Big
)

// String prints an uppercase one-word representation of the Endianness.
func (e Endianness) String() string {
	switch e {
	case Little:
		return "LittleEndian"
	case Big:
		return "BigEndian"
	default:
		return "Unknown"
	}
}

// StringFull prints an uppercase full two-word representation of the Endianness.
//
// You may optionally pass true for a lowercase representation.
func (e Endianness) StringFull(lowercase ...bool) string {
	lower := len(lowercase) > 0 && lowercase[0]
	switch e {
	case Little:
		if lower {
			return "little endian"
		}
		return "Little Endian"
	case Big:
		if lower {
			return "big endian"
		}
		return "Big Endian"
	default:
		if lower {
			return "unknown"
		}
		return "Unknown"
	}
}
