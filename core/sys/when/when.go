package when

import (
	"time"

	"git.ignitelabs.net/janos/core/std"
	"git.ignitelabs.net/janos/core/sys/num"
)

func Always() func(*std.Impulse) bool {
	return func(*std.Impulse) bool { return true }
}

// StepMaker creates two functions for synchronizing neural activations - one that makes a potential function
// tied to a specific numeric value, and the other that increments a globally shared counter.  When the potential
// is called while the global counter matches the assigned value, it will go high - otherwise it will return false.
//
// See E0S0 for a working example.
func StepMaker(limit int) (makePotential func(step int) func(*std.Impulse) bool, step func()) {
	i := 0
	return func(step int) func(*std.Impulse) bool {
			return func(*std.Impulse) bool {
				if i == step {
					return true
				}
				return false
			}
		}, func() {
			i++
			if i >= limit {
				i = 0
			}
		}
}

func Periodically(duration *time.Duration) func(*std.Impulse) bool {
	last := time.Now()
	return func(*std.Impulse) bool {
		now := time.Now()
		if now.Sub(last) > *duration {
			last = now
			return true
		}
		return false
	}
}

func Frequency[T num.Primitive](hertz *T) func(*std.Impulse) bool {
	last := time.Now()
	return func(*std.Impulse) bool {
		now := time.Now()
		if now.Sub(last) > HertzToDuration(*hertz) {
			last = now
			return true
		}
		return false
	}
}

// Resonant provides a potential that activates at a sympathetic frequency (in Hertz) to the source frequency.
//
//	Resonance = Source / Subdivision
func Resonant[T num.Primitive](source *T, subdivision *T) func(*std.Impulse) bool {
	last := time.Now()
	return func(*std.Impulse) bool {
		now := time.Now()
		resonance := *source / *subdivision
		if now.Sub(last) > HertzToDuration(resonance) {
			last = now
			return true
		}
		return false
	}
}

// HalfSpeed provides a potential that activates at half the rate of the source frequency (Hertz).
func HalfSpeed[T num.Primitive](hertz *T) func(*std.Impulse) bool {
	subdivision := T(2.0)
	return Resonant(hertz, &subdivision)
}

// QuarterSpeed provides a potential that activates at a quarter the rate of the source frequency (Hertz).
func QuarterSpeed[T num.Primitive](hertz *T) func(*std.Impulse) bool {
	subdivision := T(4.0)
	return Resonant(hertz, &subdivision)
}

// EighthSpeed provides a potential that activates at an eighth the rate of the source frequency (Hertz).
func EighthSpeed[T num.Primitive](hertz *T) func(*std.Impulse) bool {
	subdivision := T(8.0)
	return Resonant(hertz, &subdivision)
}
