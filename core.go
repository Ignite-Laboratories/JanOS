package JanOS

import (
	"sync/atomic"
	"time"
)

var Assets *assetManager
var Logging *logManager
var Signals *signalManager
var Terminate bool
var Frequency int
var AssetPath string
var masterCount uint64

func init() {
	AssetPath = "../Assets/"
	Frequency = 44000

	Signals = newSignalManager()
	Assets = newAssetManager()
	Logging = newLogManager()
}

// StepSize returns the current universal frequency back as a time.Duration
func StepSize() time.Duration { return FrequencyToDuration(Frequency) }

// NextId increments the internal master count maintained since execution and then returns the value.
// This happens as an atomic operation to ensure uniqueness across threads.
func NextId() uint64 { return atomic.AddUint64(&masterCount, 1) }

// ToFrequency returns the amount of JanOS.Frequency steps that elapse in the provided duration of time.
func ToFrequency(duration time.Duration) int {
	// Calculate the amount of nanoseconds per index
	nanoStep := float64(time.Second.Nanoseconds()) / float64(Frequency)
	// Divide the provided duration's nanoseconds by the step size to get the index count
	return int(float64(duration.Nanoseconds()) / nanoStep)
}

// FrequencyToDuration returns a time.Duration representation of the provided frequency in Hz
func FrequencyToDuration(frequency int) time.Duration {
	// Calculate the amount of nanoseconds per index
	nanoStep := float64(time.Second.Nanoseconds()) / float64(frequency)
	// Multiply the amount of nanoseconds per step by the number of steps to get the duration in nanoseconds
	return time.Duration((nanoStep * float64(1)) + 1) // We add 1 nanosecond to avoid '0' when converting back to frequency
}
