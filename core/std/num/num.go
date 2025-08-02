package num

import (
	"golang.org/x/exp/constraints"
	"math/big"
)

// BigPrimitive represents any Primitive, big.Int, or big.Float.
type BigPrimitive interface {
	Primitive | *big.Int | *big.Float
}

// Primitive represents any primitive Go integer or floating-point type.
type Primitive interface {
	Integer | Float
}

// ExtendedInteger represents any Integer or ExtendedPrimitive types.
type ExtendedInteger interface {
	Integer | ExtendedPrimitive
}

// Integer represents any primitive Go integer type.
type Integer interface {
	constraints.Integer
}

// Float represents any primitive Go floating-point type.
type Float interface {
	constraints.Float
}

// ExtendedPrimitive represents extensions of the primitive types with implicit overflow values.
type ExtendedPrimitive interface {
	Primitive | Crumb | Note | Nibble | Flake | Morsel | Shred | Run | Scale | Riff | Hook
}

// Crumb is an uint2, which implicitly overflows at 2²
type Crumb byte

// Note is an uint3, which implicitly overflows at 2³
type Note byte

// Nibble is an uint4, which implicitly overflows at 2⁴
type Nibble byte

// Flake is an uint5, which implicitly overflows at 2⁵
type Flake byte

// Morsel is an uint6, which implicitly overflows at 2⁶
type Morsel byte

// Shred is an uint7, which implicitly overflows at 2⁷
type Shred byte

// Run is an uint10, which implicitly overflows at 2¹⁰
type Run uint

// Scale is an uint12, which implicitly overflows at 2¹²
type Scale uint

// Riff is an uint24, which implicitly overflows at 2²⁴
type Riff uint

// Hook is an uint48, which implicitly overflows at 2⁴⁸
type Hook uint
