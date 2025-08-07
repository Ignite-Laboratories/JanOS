package normalize

import (
	"github.com/ignite-laboratories/core/std/num"
)

// To returns a normalized float of the provided type in the range [0.0, 1.0].
//
// NOTE: If no maximum is provided, the value from std.MaxValue[TIn] is used.
func To[TIn num.Primitive, TOut num.Float](value TIn, maximum ...TIn) TOut {
	m := TIn(num.MaxValue[TIn]())
	if len(maximum) > 0 {
		m = maximum[0]
	}

	var zero TOut
	switch any(zero).(type) {
	case float32:
		return TOut(float32(value) / float32(m))
	case float64:
		return TOut(float64(value) / float64(m))
	default:
		panic("unsupported numeric type")
	}
}

// From returns a re-scaled value of the provided float in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided float is less than 0 or greater than 1.0
//
// NOTE: If no maximum is provided, the value from std.MaxValue[TOut] is used.
func From[TIn num.Float, TOut num.Primitive](value TIn, maximum ...TOut) TOut {
	if value < 0.0 || value > 1.0 {
		panic("value must be in range [0.0, 1.0]")
	}

	m := TOut(num.MaxValue[TOut]())
	if len(maximum) > 0 {
		m = maximum[0]
	}
	return TOut(value * TIn(m))
}
