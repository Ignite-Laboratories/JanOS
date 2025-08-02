package phrase

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/measurement"
	"github.com/ignite-laboratories/core/std/name"
)

func newPhrase[T any]() std.Phrase[T] {
	return std.Phrase[T]{
		Name: name.Tiny().Name,
		Data: make([]std.Measurement[T], 0),
	}
}

/**
Creation
*/

func Of[T any](data ...T) std.Phrase[T] {
	p := newPhrase[T]()
	for _, d := range data {
		p = p.AppendMeasurement(measurement.Of(d))
	}
	return p
}

func OfMeasurements[T any](m ...std.Measurement[T]) std.Phrase[T] {
	p := newPhrase[T]()
	p.Data = m
	return p
}

// OfBytes creates a named Phrase of the provided bytes and name.
func OfBytes(name string, bytes ...byte) std.Phrase[any] {
	p := newPhrase[any]()
	p.Name = name
	p.Data = []std.Measurement[any]{measurement.OfBytes(bytes...)}
	return p
}

// OfBits creates a named Phrase of the provided bits and name.
func OfBits(name string, bits ...std.Bit) std.Phrase[any] {
	p := newPhrase[any]()
	p.Name = name
	p.Data = []std.Measurement[any]{measurement.OfBits(bits...)}
	return p
}
