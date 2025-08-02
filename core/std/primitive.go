package std

import "math/big"

// Primitive represents the primitive types provided by Go, all of which are convertable between Measurement and primitive form.
//
// In addition, the big.Int and big.Float types are considered "primitive" as they are fully interoperable with the matrix engine.
type Primitive interface {
	*big.Int | *big.Float |
		int8 | int16 | int32 | int64 |
		uint8 | uint16 | uint32 | uint64 |
		float32 | float64 |
		complex64 | complex128 |
		int | uint | uintptr |
		bool
}
