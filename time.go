package JanOS

import (
	"time"
)

// TimeSlice represents a slice of data relative to a moment in time.
type TimeSlice struct {
	StartTime time.Time
	Data      []float64
}

// ToIndexCount returns an index representation of the duration of time provided.
func ToIndexCount(duration time.Duration) int {
	// Calculate the amount of nanoseconds per index
	nanoStep := float64(time.Second.Nanoseconds()) / Universe.Resolution.Frequency
	// Divide the provided duration's nanoseconds by the step size to get the index count
	return int(float64(duration.Nanoseconds()) / nanoStep)
}

// ToDuration returns a time.Duration representation of the amount of indices provided
// Check Universe.Resolution for the operational frequency.
func ToDuration(steps int) time.Duration {
	// Calculate the amount of nanoseconds per index
	nanoStep := float64(time.Second.Nanoseconds()) / Universe.Resolution.Frequency
	// Multiply the amount of nanoseconds per step by the number of steps to get the duration in nanoseconds
	return time.Duration(nanoStep * float64(steps))
}
