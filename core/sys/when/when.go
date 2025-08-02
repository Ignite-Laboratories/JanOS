package when

import (
	"github.com/ignite-laboratories/core"
	"time"
)

// Frequency provides a potential that activates at the specified frequency (in Hertz).
func Frequency(hertz *float64) core.Potential {
	return func(ctx core.Context) bool {
		d := core.HertzToDuration(*hertz)
		return ctx.Moment.Sub(ctx.LastActivation.Inception) > d
	}
}

// Resonant provides a potential that activates at a sympathetic frequency (in Hertz) to the source frequency.
//
//	Resonance = Source / Subdivision
func Resonant(source *float64, subdivision *float64) core.Potential {
	return func(ctx core.Context) bool {
		resonance := *source / *subdivision
		d := core.HertzToDuration(resonance)
		return ctx.Moment.Sub(ctx.LastActivation.Inception) > d
	}
}

// HalfSpeed provides a potential that activates at half the rate of the source frequency (Hertz).
func HalfSpeed(hertz *float64) core.Potential {
	subdivision := 2.0
	return Resonant(hertz, &subdivision)
}

// QuarterSpeed provides a potential that activates at a quarter the rate of the source frequency (Hertz).
func QuarterSpeed(hertz *float64) core.Potential {
	subdivision := 4.0
	return Resonant(hertz, &subdivision)
}

// EighthSpeed provides a potential that activates at an eighth the rate of the source frequency (Hertz).
func EighthSpeed(hertz *float64) core.Potential {
	subdivision := 8.0
	return Resonant(hertz, &subdivision)
}

// Duration provides the following potential:
//
//	time.Now().Sub(ctx.LastActivation.Inception) > duration
func Duration(duration *time.Duration) core.Potential {
	return func(ctx core.Context) bool {
		return time.Now().Sub(ctx.LastActivation.Inception) > *duration
	}
}

// Pace provides a potential that counts to the provided value before returning true.
//
// NOTE: This kind of potential is an impulse slowing operation, regardless of neural synchronicity!
func Pace(value *uint64) core.Potential {
	return func(ctx core.Context) bool {
		for i := uint64(0); i < *value; i++ {
		}
		return true
	}
}

// Always provides a potential that always returns true.
func Always(ctx core.Context) bool {
	return true
}

// Never provides a potential that never returns true.
func Never() bool {
	return false
}

// High provides a potential that dereferences the provided boolean on demand.
func High(value *bool) core.Potential {
	return func(ctx core.Context) bool {
		return *value
	}
}
