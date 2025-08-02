// Package measurement provides higher-order access to std.Measurement functions.
package measurement

import (
	"fmt"
	"github.com/ignite-laboratories/core/enum/endian"
	"github.com/ignite-laboratories/core/enum/traveling"
	"github.com/ignite-laboratories/core/std"
)

// OfZeros creates a new std.Measurement[any] of the provided bit-width consisting entirely of 0s.
func OfZeros(width int) std.Measurement[any] {
	return std.Measurement[any]{
		Bytes:      make([]byte, width/8),
		Bits:       make([]std.Bit, width%8),
		Endianness: endian.Big,
	}.RollUp()
}

// OfOnes creates a new std.Measurement[any] of the provided bit-width consisting entirely of 1s.
func OfOnes(width int) std.Measurement[any] {
	zeros := OfZeros(width)
	for i := range zeros.Bytes {
		zeros.Bytes[i] = 255
	}
	for i := range zeros.Bits {
		zeros.Bits[i] = 1
	}
	return zeros.RollUp()
}

// Of creates a new std.Measurement[T] of the provided input data by reading it directly from memory.
func Of[T any](data T) std.Measurement[T] {

}

// OfBits creates a new std.Measurement[any] of the provided std.Bit slice.
func OfBits(bits ...std.Bit) std.Measurement[any] {
	std.BitSanityCheck(bits...)
	return std.Measurement[any]{
		Bits:       bits,
		Endianness: endian.Big,
	}.RollUp()
}

// OfBytes creates a new std.Measurement[any] of the provided byte slice.
func OfBytes(bytes ...byte) std.Measurement[any] {
	return std.Measurement[any]{
		Bytes:      bytes,
		Endianness: endian.Big,
	}.RollUp()
}

// OfPattern creates a new std.Measurement[T] of the provided bit-width consisting of the pattern emitted across it in the direction.Direction of travel.Traveling.
//
// Inward and outward travel directions are supported and work from the midpoint of the width, biased towards the west.
func OfPattern(w uint, t traveling.Traveling, pattern ...std.Bit) std.Measurement[any] {
	if w <= 0 || len(pattern) == 0 {
		return std.Measurement[any]{
			Endianness: endian.Big,
		}
	}

	if t == traveling.Northbound || t == traveling.Southbound {
		panic(fmt.Sprintf("cannot take a latitudinal binary measurement [%v]", t.StringFull(true)))
	}

	printer := func(width uint, tt traveling.Traveling) []std.Bit {
		bits := make([]std.Bit, width)
		patternI := 0
		for i := 0; i < int(width); i++ {
			ii := i
			if tt == traveling.Westbound {
				ii = int(width) - 1 - i
			}

			bits[ii] = pattern[patternI]
			patternI = (patternI + 1) % len(pattern)
		}
		return bits
	}

	if t == traveling.Inbound || t == traveling.Outbound {
		leftWidth := w / 2
		rightWidth := w - leftWidth

		if t == traveling.Inbound {
			left := OfBits(printer(leftWidth, traveling.Eastbound)...)
			right := OfBits(printer(rightWidth, traveling.Westbound)...)
			return left.AppendMeasurements(right)
		}
		return OfBits(printer(leftWidth, traveling.Westbound)...).Append(printer(rightWidth, traveling.Eastbound)...)
	}
	return OfBits(printer(w, t)...)
}

// OfString creates a new std.Measurement[T] from the provided binary input string.
//
// NOTE: This will panic if anything but a 1 or 0 is found in the input string.
func OfString(s string) std.Measurement[any] {
	bits := make([]std.Bit, len(s))
	for i := 0; i < len(s); i++ {
		bits[i] = std.Bit(s[i])
	}
	return OfBits(bits...)
}
