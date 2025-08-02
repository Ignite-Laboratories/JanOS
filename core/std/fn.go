package std

import (
	"time"
)

// DurationToHertz converts a time.Duration into Hertz.
func DurationToHertz(d time.Duration) float64 {
	if d < 0 {
		d = 0
	}
	s := float64(d) / 1e9
	hz := 1 / s
	return hz
}

// HertzToDuration converts a Hertz value to a time.Duration.
func HertzToDuration(hz float64) time.Duration {
	if hz <= 0 {
		// No division by zero
		hz = 1e-100 // math.SmallestNonzeroFloat64 🡨 NOTE: Raspberry Pi doesn't handle this constant well
	}
	s := 1 / hz
	ns := s * 1e9
	return time.Duration(ns)
}

// AbsDuration returns the absolute value of the provided duration.
func AbsDuration(d time.Duration) time.Duration {
	if d < 0 {
		d = -d
	}
	return d
}

// Action functions are provided temporal context when invoked.
type Action func(ctx Context)

// Potential functions are provided temporal context when invoked in order to make decisions.
type Potential func(ctx Context) bool
