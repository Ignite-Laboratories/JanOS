package internal

import (
	"encoding/binary"
	"github.com/ignite-laboratories/core/enum/endian"
	"reflect"
	"unsafe"
)

// GetArchitectureEndianness returns the Endianness of the currently executing hardware.
func GetArchitectureEndianness() endian.Endianness {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, 0xABCD)
	if buf[0] == 0xAB {
		return endian.Big
	}
	return endian.Little
}

// ReverseByte reverses all the bits of a byte.
func ReverseByte(b byte) byte {
	b = (b&0xF0)>>4 | (b&0x0F)<<4
	b = (b&0xCC)>>2 | (b&0x33)<<2
	return (b&0xAA)>>1 | (b&0x55)<<1
}

// Measure takes "raw" measurements of many objects at runtime and returns slices with each of their underlying bytes.
//
// NOTE: Be sure to explicitly provide the type parameter to ensure Go doesn't implicitly
// give you, say, all 8 bytes worth of an 'int' to represent a single 'byte' =)
func Measure[T any](values ...T) [][]byte {
	out := make([][]byte, len(values))
	for i, v := range values {
		out[i] = measure(v)
	}
	return out
}

func measure[T any](value T) []byte {
	var size uintptr
	switch any(value).(type) {
	case byte, int8, bool:
		size = 1
	case uint16, int16:
		size = 2
	case uint32, int32, float32:
		size = 4
	case uint64, int64, float64, uint, int:
		size = 8
	case complex64:
		size = 8
	case complex128:
		size = 16
	case string:
		return []byte(any(value).(string))
	default:
		// Handle other types including slices using reflection
		val := reflect.ValueOf(value)
		if val.Kind() == reflect.Slice {
			if val.Len() == 0 {
				return []byte{}
			}
			elemSize := val.Type().Elem().Size()
			totalSize := uintptr(val.Len()) * elemSize
			size = totalSize
		} else {
			size = reflect.TypeOf(value).Size()
		}
	}

	if size == 0 {
		return []byte{}
	}

	var dataPtr unsafe.Pointer
	if val := reflect.ValueOf(value); val.Kind() == reflect.Slice {
		dataPtr = unsafe.Pointer(val.UnsafePointer())
	} else {
		dataPtr = unsafe.Pointer(reflect.ValueOf(&value).Elem().UnsafeAddr())
	}

	bytes := make([]byte, size)
	copy(bytes, (*[1 << 30]byte)(dataPtr)[:size:size])
	return bytes
}
