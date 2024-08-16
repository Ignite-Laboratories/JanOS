package JanOS

import (
	"math"
	"time"
)

// SineWave starts a live loop that oscillates off of the source dimension using the provided input dimensions.
func (d *Dimension) SineWave(amplitude *Dimension, frequency *Dimension) *Dimension {
	now := time.Now()
	// Grab these immediately so the values don't change
	f := frequency.GetValue(now)
	a := amplitude.GetValue(now)

	Universe.Printf(d, "Oscillating [%s] at %fhz, amplitude %f", string(d.Symbol), f, a)
	go func() {
		lastUpdate := time.Now()

		for {
			now = time.Now()
			if Universe.Terminate {
				break
			}

			if time.Since(lastUpdate) >= d.Timeline.resolution.Duration {
				f = frequency.GetValue(now)
				a = amplitude.GetValue(now)
				periodOffset := time.Since(lastUpdate).Seconds()

				phaseShiftInRadians := (360.0 * periodOffset * f) * math.Pi / 180
				angularFrequency := 2 * math.Pi * f
				calculatedValue := a * math.Sin(angularFrequency*time.Second.Seconds()+phaseShiftInRadians)
				d.Timeline.setValue(now, calculatedValue)
			}
		}
		Universe.Printf(d, "%s [%s] stopped oscillating", d.Name, string(d.Symbol))
	}()
	return d
}
