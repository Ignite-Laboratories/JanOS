// Package measurement provides higher-order access to std.Measurement functions.
package measurement

import (
	"fmt"
	"github.com/ignite-laboratories/core/enum/endian"
	"github.com/ignite-laboratories/core/enum/traveling"
	"github.com/ignite-laboratories/core/internal"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/name"
	"reflect"
	"unsafe"
)

// To attempts to convert a std.Measurement[any] of binary information into the specified type T.
func To[T any](m std.Measurement[any]) T {
	bits := m.GetAllBits()
	var zero T
	typeOf := reflect.TypeOf(zero)

	// Handle slices
	if typeOf.Kind() == reflect.Slice {
		elemType := typeOf.Elem()
		elemSize := elemType.Size()

		numElements := len(bits) / (8 * int(elemSize))
		if numElements == 0 {
			return zero
		}

		sliceVal := reflect.MakeSlice(typeOf, numElements, numElements)
		slicePtr := unsafe.Pointer(sliceVal.UnsafePointer())
		resultBytes := unsafe.Slice((*byte)(slicePtr), numElements*int(elemSize))

		byteI := (len(bits) / 8) - 1
		i := len(bits) - 1
		for i > 0 {
			var currentByte byte
			for ii := 0; ii < 8; ii++ {
				if bits[i] == 1 {
					currentByte |= 1 << ii
				}
				i--
			}

			resultBytes[byteI] = currentByte
			byteI--
		}

		return sliceVal.Interface().(T)
	}

	// Handle non-slices
	size := typeOf.Size()
	if len(bits) > int(size)*8 {
		panic("bit slice too large for target type")
	}

	result := zero
	resultPtr := unsafe.Pointer(&result)
	resultBytes := unsafe.Slice((*byte)(resultPtr), size)

	byteI := (len(bits) / 8) - 1
	i := len(bits) - 1
	for i > 0 {
		var currentByte byte
		for ii := 0; ii < 8; ii++ {
			if bits[i] == 1 {
				currentByte |= 1 << ii
			}
			i--
		}

		resultBytes[byteI] = currentByte
		byteI--
	}

	return result
}

// Of creates a new std.Measurement[T] of the provided input data by reading it directly from memory.
func Of[T any](data T) std.Measurement[T] {
	m := std.Measurement[T]{
		Bytes:      internal.Measure[T](data)[0],
		Endianness: internal.GetArchitectureEndianness(),
	}
	m.GivenName = name.Random[name.Default]()
	return m
}

// OfZeros creates a new std.Measurement[any] of the provided bit-width consisting entirely of 0s.
func OfZeros(width int) std.Measurement[any] {
	m := std.Measurement[any]{
		Bytes:      make([]byte, width/8),
		Bits:       make([]std.Bit, width%8),
		Endianness: endian.Big,
	}.RollUp()
	m.GivenName = name.Random[name.Default]()
	return m
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

// OfBits creates a new std.Measurement[any] of the provided std.Bit slice.
func OfBits(bits ...std.Bit) std.Measurement[any] {
	std.BitSanityCheck(bits...)
	m := std.Measurement[any]{
		Bits:       bits,
		Endianness: endian.Big,
	}.RollUp()
	m.GivenName = name.Random[name.Default]()
	return m
}

// OfBytes creates a new std.Measurement[any] of the provided byte slice.
func OfBytes(bytes ...byte) std.Measurement[any] {
	m := std.Measurement[any]{
		Bytes:      bytes,
		Endianness: endian.Big,
	}.RollUp()
	m.GivenName = name.Random[name.Default]()
	return m
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
