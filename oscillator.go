package JanOS

import (
	"math"
	"time"
)

// NewOscillatingDimension creates a dimension that oscillates over the period
// of one second and updates at the specified resolution interval per second.
func (mgr *dimensionManager) NewOscillatingDimension(name string, symbol Symbol, amplitude *Dimension, frequency *Dimension) *Dimension {
	d := Universe.Dimensions.NewDimension(name, symbol, 0)
	Universe.Printf(mgr, "Oscillating [%s] at %fhz, amplitude %f", string(symbol), frequency.Value, amplitude.Value)
	resolutionStep := time.Duration(1 / Universe.Resolution)

	go func() {
		lastUpdate := time.Now()
		lastCycle := time.Now()
		// Grab these immediately so the values don't change
		f := frequency.Value
		a := amplitude.Value

		for {
			if Universe.Terminate {
				break
			}

			if time.Since(lastUpdate) > resolutionStep {
				frequencyStep := time.Duration(float64(time.Second.Nanoseconds()) / f).Nanoseconds()
				frequencyOffset := time.Since(lastCycle).Nanoseconds()
				periodOffset := time.Since(lastUpdate).Seconds()

				phaseShiftInRadians := (360.0 * periodOffset * f) * math.Pi / 180
				angularFrequency := 2 * math.Pi * f
				d.Value = a * math.Sin(angularFrequency*time.Second.Seconds()+phaseShiftInRadians)

				// This ensures that we wait until as close to the current frequency time away from the last cycle
				// before updating the inputs.  It isn't perfect, but fuzzy logic is okay.a
				if frequencyOffset > frequencyStep {
					lastCycle = time.Now()
					// Update these at the start of every cycle for smooth transition on change
					f = frequency.Value
					a = amplitude.Value
				}
			}
		}
		Universe.Printf(mgr, "[%s] stopped", string(symbol))
	}()

	return d
}
