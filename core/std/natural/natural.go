package std

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/measurement"
	"github.com/ignite-laboratories/core/sys/support"
)

func From[T std.Numeric[any]](value T) std.Natural {

}

// NaturalFrom takes a Measurement of the provided unsigned integer value as a Natural number.
func NaturalFrom(value uint) Natural {
	return Natural{
		Measurement: measurement.OfBytes(support.Measure(value)[0]...),
	}
}
