package when

import (
	"core/sys/num"
	"time"
)

func Always() bool {
	return true
}

func Never() bool {
	return false
}

func Periodically(duration *time.Duration) func() bool {
	last := time.Now()
	return func() bool {
		now := time.Now()
		if now.Sub(last) > *duration {
			last = now
			return true
		}
		return false
	}
}

func Frequency[T num.Primitive](hertz *T) func() bool {
	last := time.Now()
	return func() bool {
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
func Resonant[T num.Primitive](source *T, subdivision *T) func() bool {
	last := time.Now()
	return func() bool {
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
func HalfSpeed[T num.Primitive](hertz *T) func() bool {
	subdivision := T(2.0)
	return Resonant(hertz, &subdivision)
}

// QuarterSpeed provides a potential that activates at a quarter the rate of the source frequency (Hertz).
func QuarterSpeed[T num.Primitive](hertz *T) func() bool {
	subdivision := T(4.0)
	return Resonant(hertz, &subdivision)
}

// EighthSpeed provides a potential that activates at an eighth the rate of the source frequency (Hertz).
func EighthSpeed[T num.Primitive](hertz *T) func() bool {
	subdivision := T(8.0)
	return Resonant(hertz, &subdivision)
}
