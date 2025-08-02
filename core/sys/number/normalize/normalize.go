package normalize

import "github.com/ignite-laboratories/core/sys/number"

// To returns a normalized float of the provided type in the range [0.0, 1.0].
func To[TIn number.Numeric, TOut number.Float](value TIn) TOut {
	var zero TOut
	switch any(zero).(type) {
	case float32:
		return TOut(float32(value) / float32(number.MaxValue[TIn]()))
	case float64:
		return TOut(float64(value) / float64(number.MaxValue[TIn]()))
	default:
		panic("unsupported numeric type")
	}
}

// From returns a re-scaled value of the provided float in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided float is less than 0 or greater than 1.0
func From[TIn number.Float, TOut number.Numeric](value TIn) TOut {
	if value < 0.0 || value > 1.0 {
		panic("value must be in range [0.0, 1.0]")
	}
	return TOut(value * TIn(number.MaxValue[TIn]()))
}
