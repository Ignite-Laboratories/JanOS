package std

import (
	"github.com/ignite-laboratories/core/enum/endian"
	"github.com/ignite-laboratories/core/internal"
	"strings"
)

// Measurement is a variable-width slice of bits and is used to efficiently store them in operating memory.
// As Go inherently requires at least 8 bits to store custom types, storing each bit individually
// would need 8 times the size of every bit - thus, the measurement was born.  For higher-order measurement
// procedures, see the 'measurement' package directly.
//
// The type represents what the underlying information was typed as when it was measured - which for arbitrary
// binary information is simply 'any'.
//
// NOTE: ALL measurements are processed in the traditional endian.Big form - however, at the time of measurement we
// ALSO capture the original endianness of the stored value.  It's ENTIRELY informational and can be ignored - but
// it's still quite interesting if you care to investigate =)
type Measurement[T any] struct {
	// Endianness indicates the endian.Endianness of the data as it was originally stored before being measured in standard endian.Big form.
	//
	// NOTE: Don't get too caught up here - it's purely informational and has absolutely no bearing on tiny operations.
	endian.Endianness

	// Bytes holds complete byte data.
	Bytes []byte

	// Bits holds any remaining bits.
	Bits []Bit
}

/**
From Functions
*/

// newMeasurementFrom creates a new Measurement of the provided Bit slice in the order they are provided.
//
// NOTE: These are convenience methods - for the full gamut of Measurement features, see the measurement package.
func newMeasurementFrom[T any](bytes []byte, bits ...Bit) Measurement[T] {
	if bytes == nil {
		bytes = []byte{}
	}

	BitSanityCheck(bits...)
	return Measurement[T]{
		Bytes:      bytes,
		Bits:       bits,
		Endianness: endian.Big,
	}.RollUp()
}

// newMeasurementOfBits creates a new Measurement of the provided Bit slice.
//
// NOTE: These are convenience methods - for the full gamut of Measurement features, see the measurement package.
func newMeasurementOfBits[T any](bits ...Bit) Measurement[T] {
	return newMeasurementFrom[T](nil, bits...)
}

// newMeasurementFromBytes creates a new Measurement of the provided byte slice.
//
// NOTE: These are convenience methods - for the full gamut of Measurement features, see the measurement package.
func newMeasurementFromBytes[T any](bytes ...byte) Measurement[T] {
	return newMeasurementFrom[T](bytes)
}

/**
Methods
*/

// BitWidth gets the total bit width of this Measurement's recorded data.
func (a Measurement[T]) BitWidth() uint {
	return uint((len(a.Bytes) * 8) + len(a.Bits))
}

// GetAllBits returns a slice of the Measurement's individual bits.
func (a Measurement[T]) GetAllBits() []Bit {
	a = a.sanityCheck()
	var byteBits []Bit
	for _, b := range a.Bytes {
		bits := make([]Bit, 8)
		for i := 7; i >= 0; i-- {
			bits[7-i] = Bit((b >> i) & 1)
		}
		byteBits = append(byteBits, bits...)
	}
	return append(byteBits, a.Bits...)
}

// Append places the provided bits at the end of the Measurement.
func (a Measurement[T]) Append(bits ...Bit) Measurement[T] {
	a = a.sanityCheck(bits...)

	a.Bits = append(a.Bits, bits...)
	return a.RollUp()
}

// AppendBytes places the provided bits at the end of the Measurement.
func (a Measurement[T]) AppendBytes(bytes ...byte) Measurement[T] {
	a = a.sanityCheck()

	lastBits := a.Bits
	for _, b := range bytes {
		bits := make([]Bit, 8)

		ii := 0
		for i := byte(7); i < 8; i-- {
			bits[ii] = Bit((b >> i) & 1)
			ii++
		}

		blended := append(lastBits, bits[:8-len(lastBits)]...)
		lastBits = bits[8-len(lastBits):]

		var newByte byte
		ii = 0
		for i := byte(7); i < 8; i-- {
			newByte |= byte(blended[ii]) << i
			ii++
		}

		a.Bytes = append(a.Bytes, newByte)
	}

	a.Bits = lastBits
	return a.RollUp()
}

// AppendMeasurements appends the bits of the provided measurements to the end of the source measurement.
func (a Measurement[T]) AppendMeasurements(m ...Measurement[T]) Measurement[T] {
	for _, mmt := range m {
		a = a.Append(mmt.GetAllBits()...)
	}
	return a.RollUp()
}

// Prepend places the provided bits at the start of the Measurement.
func (a Measurement[T]) Prepend(bits ...Bit) Measurement[T] {
	a = a.sanityCheck(bits...)

	oldBits := a.Bits
	oldBytes := a.Bytes
	a.Bytes = []byte{}
	a.Bits = []Bit{}
	a = a.Append(bits...)
	a = a.AppendBytes(oldBytes...)
	a = a.Append(oldBits...)
	return a.RollUp()
}

// PrependBytes places the provided bytes at the start of the Measurement.
func (a Measurement[T]) PrependBytes(bytes ...byte) Measurement[T] {
	a = a.sanityCheck()

	oldBits := a.Bits
	oldBytes := a.Bytes
	a.Bytes = bytes
	a.Bits = []Bit{}
	a = a.AppendBytes(oldBytes...)
	a = a.Append(oldBits...)
	return a.RollUp()
}

// PrependMeasurements prepends the bits of the provided measurements at the start of the source measurement.
func (a Measurement[T]) PrependMeasurements(m ...Measurement[T]) Measurement[T] {
	if len(m) == 0 {
		return a
	}

	result := m[len(m)-1]
	for i := len(m) - 2; i >= 0; i-- {
		result = m[i].AppendBytes(result.Bytes...).Append(result.Bits...)
	}
	result = result.AppendBytes(a.Bytes...).Append(a.Bits...)
	return result.RollUp()
}

// Reverse reverses the order of all bits in the measurement.
func (a Measurement[T]) Reverse() Measurement[T] {
	reversedBytes := make([]byte, len(a.Bytes))
	reversedBits := make([]Bit, len(a.Bits))

	ii := 0
	for i := len(a.Bytes) - 1; i >= 0; i-- {
		reversedBytes[ii] = internal.ReverseByte(a.Bytes[i])
		ii++
	}

	ii = 0
	for i := len(a.Bits) - 1; i >= 0; i-- {
		reversedBits[ii] = a.Bits[i]
		ii++
	}

	a.Bytes = reversedBytes
	a.Bits = make([]Bit, 0)
	return a.Prepend(reversedBits...)
}

// BleedLastBit returns the last bit of the measurement and a measurement missing that bit.
func (a Measurement[T]) BleedLastBit() (Bit, Measurement[T]) {
	if a.BitWidth() == 0 {
		panic("cannot bleed the last bit of an empty measurement")
	}
	// TODO: Implement this
	return 0, a
}

// BleedFirstBit returns the first bit of the measurement and a measurement missing that bit.
func (a Measurement[T]) BleedFirstBit() (Bit, Measurement[T]) {
	if a.BitWidth() == 0 {
		panic("cannot bleed the first bit of an empty measurement")
	}

	// TODO: Implement this
	return 0, a
}

// RollUp combines the currently measured bits into the measured bytes if there is enough recorded.
func (a Measurement[T]) RollUp() Measurement[T] {
	for len(a.Bits) >= 8 {
		var b byte
		for i := byte(7); i < 8; i-- {
			if a.Bits[i] == 1 {
				b |= 1 << (7 - i)
			}
		}
		a.Bits = a.Bits[8:]
		a.Bytes = append(a.Bytes, b)
	}
	return a
}

/**
Utilities
*/

// sanityCheck is a measurement level sanity check that ensures the provided bits are all 1s and 0s before
// rolling the currently measured bits into bytes wherever possible.
func (a Measurement[T]) sanityCheck(bits ...Bit) Measurement[T] {
	if a.Bytes == nil {
		a.Bytes = []byte{}
	}
	if a.Bits == nil {
		a.Bits = []Bit{}
	}
	BitSanityCheck(bits...)
	return a.RollUp()
}

// String converts the measurement to a binary string entirely consisting of 1s and 0s.
func (a Measurement[T]) String() string {
	bits := a.GetAllBits()

	builder := strings.Builder{}
	builder.Grow(len(bits))
	for _, b := range bits {
		builder.WriteString(b.String())
	}
	return builder.String()
}

// StringPretty returns a measurement-formatted string of the current binary information. Measurements
// are simply formatted with a single space between digits.
func (a Measurement[T]) StringPretty() string {
	bits := a.GetAllBits()

	if len(bits) == 0 {
		return ""
	}

	builder := strings.Builder{}
	builder.Grow(len(bits)*2 - 1)

	builder.WriteString(bits[0].String())

	if len(bits) > 1 {
		for _, bit := range bits[1:] {
			builder.WriteString(" ")
			builder.WriteString(bit.String())
		}
	}

	return builder.String()
}
