package std

import (
	"github.com/ignite-laboratories/core/std/num"
	"math/big"
)

// Numeric represents any num.Primitive, big.Int, or big.Float, Measurement, or Phrase.
//
// NOTE: The type constraint is simply to ensure measurements and phrases can carry their type information
// with them, just use 'any' if you don't need that system.
type Numeric[T any] interface {
	num.Primitive | *big.Int | *big.Float | Measurement[T] | Phrase[T]
}
