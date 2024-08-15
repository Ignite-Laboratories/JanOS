package JanOS

import (
	"math"
	"time"
)

// NewOscillatingDimension creates a dimension that oscillates over the period
// of one second and updates at the specified resolution interval per second.
func (mgr *dimensionManager) NewOscillatingDimension(name string, symbol Symbol, amplitude float64, frequency float64, resolution int64) *Dimension {
	d := Universe.Dimensions.NewDimension(name, symbol, 0)
	Universe.Printf(mgr, "Oscillating [%s] at %fhz, amplitude %f, resolution %d steps/sec", string(symbol), frequency, amplitude, resolution)
	resolutionStep := time.Duration(time.Second.Nanoseconds() / resolution)

	go func() {
		lastUpdate := time.Now()
		lastCycle := time.Now()

		for {
			if Universe.Terminate {
				break
			}

			if time.Since(lastUpdate) > resolutionStep {
				periodOffset := time.Since(lastUpdate).Seconds()

				phaseShiftInRadians := (360.0 * periodOffset * frequency) * math.Pi / 180
				angularFrequency := 2 * math.Pi * frequency
				d.Value = amplitude * math.Sin(angularFrequency*time.Second.Seconds()+phaseShiftInRadians)

				if periodOffset > 1 {
					lastCycle = lastCycle.Add(time.Second)
				}
			}
		}
		Universe.Printf(mgr, "[%s] stopped", string(symbol))
	}()

	return d
}
