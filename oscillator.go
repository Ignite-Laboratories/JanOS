package JanOS

import (
	"math"
	"time"
)

// NewOscillatingDimension creates a dimension that oscillates over time.
func (mgr *dimensionManager) NewOscillatingDimension(name string, symbol Symbol, amplitude *Dimension, frequency *Dimension) *Dimension {
	now := time.Now()
	// Grab these immediately so the values don't change
	f := frequency.GetValue(now)
	a := amplitude.GetValue(now)

	d := Universe.Dimensions.NewDimension(name, symbol, 0)
	Universe.Printf(mgr, "Oscillating [%s] at %fhz, amplitude %f", string(symbol), f, a)
	resolutionStep := time.Duration(1 / d.Timeline.resolution.Frequency)

	go func() {
		lastUpdate := time.Now()
		lastCycle := time.Now()

		for {
			now = time.Now()
			if Universe.Terminate {
				break
			}

			if time.Since(lastUpdate) > resolutionStep {
				f = frequency.GetValue(now)
				a = amplitude.GetValue(now)
				frequencyStep := time.Duration(float64(time.Second.Nanoseconds()) / f).Nanoseconds()
				frequencyOffset := time.Since(lastCycle).Nanoseconds()
				periodOffset := time.Since(lastUpdate).Seconds()

				phaseShiftInRadians := (360.0 * periodOffset * f) * math.Pi / 180
				angularFrequency := 2 * math.Pi * f
				d.Timeline.setValue(now, a*math.Sin(angularFrequency*time.Second.Seconds()+phaseShiftInRadians))

				// This ensures that we wait until as close to the current frequency time away from the last cycle
				// before updating the inputs.  It isn't perfect, but fuzzy logic is okay.a
				if frequencyOffset > frequencyStep {
					lastCycle = now
					// Update these at the start of every cycle for smooth transition on change
					f = frequency.GetValue(now)
					a = amplitude.GetValue(now)
				}
			}
		}
		Universe.Printf(mgr, "%s [%s] stopped oscillating", name, string(symbol))
	}()

	return d
}
