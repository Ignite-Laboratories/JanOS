package when

import (
	"github.com/ignite-laboratories/core"
)

type beat struct{}

// Beat provides a set of helper functions for creating beat-oriented potentials.
var Beat beat = beat{}

// Downbeat provides a potential that fires when the beat is 0.
func (b beat) Downbeat(ctx core.Context) bool {
	return ctx.Beat == 0
}

// Even provides a potential that fires when the beat is even.
func (b beat) Even(ctx core.Context) bool {
	return ctx.Beat%2 == 0
}

// Odd provides a potential that returns true when the beat is odd.
func (b beat) Odd(ctx core.Context) bool {
	return ctx.Beat%2 != 0
}

// Modulo provides the following potential:
//
//	ctx.Beat % value == 0
func (b beat) Modulo(value *int) core.Potential {
	return func(ctx core.Context) bool {
		return ctx.Beat%*value == 0
	}
}

// Over provides the following potential:
//
//	ctx.Beat > value
func (b beat) Over(value *int) core.Potential {
	return func(ctx core.Context) bool {
		return ctx.Beat > *value
	}
}

// On provides the following potential:
//
//	ctx.Beat == value
func (b beat) On(beat *int) core.Potential {
	return func(ctx core.Context) bool {
		return ctx.Beat == *beat
	}
}

// Under provides the following potential:
//
//	ctx.Beat < value
func (b beat) Under(value *int) core.Potential {
	return func(ctx core.Context) bool {
		return ctx.Beat < *value
	}
}
