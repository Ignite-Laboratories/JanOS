package JanOS

import (
	"math"
	"time"
)

// SineWave starts a live loop that oscillates off of the source signal using the provided input signals.
func (signal *Signal) SineWave(amplitude *Signal, frequency *Signal) *Signal {
	now := time.Now()
	// Grab these immediately so the values don't change
	f := frequency.GetValue(now).Value
	a := amplitude.GetValue(now).Value

	Universe.Printf(signal, "Sine Wave ƒ(%s) => y = %s * sin(%s * t)", string(signal.Symbol), string(amplitude.Symbol), string(frequency.Symbol))
	go func() {
		lastUpdate := time.Now()

		for {
			now = time.Now()
			if Universe.Terminate {
				break
			}

			if time.Since(lastUpdate) >= signal.Timeline.resolution.Duration {
				f = frequency.GetValue(now).Value
				a = amplitude.GetValue(now).Value
				// Seconds() gives us a float, which acts as a scale factor
				// for the position of the period relative to 1 second.
				periodOffset := time.Since(lastUpdate).Seconds()

				phaseShiftInRadians := (360.0 * periodOffset * f) * math.Pi / 180
				angularFrequency := 2 * math.Pi * f
				calculatedValue := a * math.Sin(angularFrequency+phaseShiftInRadians)
				signal.Timeline.setValue(now, calculatedValue)
			}
		}
		Universe.Printf(signal, "%s stopped oscillating", string(signal.Symbol))
	}()
	return signal
}
