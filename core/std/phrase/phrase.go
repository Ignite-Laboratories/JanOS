package phrase

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/measurement"
	"github.com/ignite-laboratories/core/std/name"
	"github.com/ignite-laboratories/core/std/num"
)

func newPhrase[TMmt any, TName name.Format]() std.Phrase[TMmt] {
	p := std.Phrase[TMmt]{
		Data: make([]std.Measurement[TMmt], 0),
	}
	p.GivenName = name.Random[TName]()
	return p
}

/**
Creation
*/

func Of[T any](data ...T) std.Phrase[T] {
	p := newPhrase[T, name.Default]()
	for _, d := range data {
		p = p.AppendMeasurement(measurement.Of(d))
	}
	return p
}

func OfMeasurements[T any](m ...std.Measurement[T]) std.Phrase[T] {
	p := newPhrase[T, name.Default]()
	p.Data = m
	return p
}

// OfBytes creates a named Phrase of the provided bytes and name.
func OfBytes(bytes ...byte) std.Phrase[any] {
	p := newPhrase[any, name.Default]()
	p.Data = []std.Measurement[any]{measurement.OfBytes(bytes...)}
	return p
}

// OfBits creates a named Phrase of the provided bits and name.
func OfBits(bits ...num.Bit) std.Phrase[any] {
	p := newPhrase[any, name.Default]()
	p.Data = []std.Measurement[any]{measurement.OfBits(bits...)}
	return p
}
