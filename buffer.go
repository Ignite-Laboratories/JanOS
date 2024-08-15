package JanOS

import "time"

type Buffer struct {
	StartTime time.Time
	Data      []float64
}

type BufferDimension struct {
	data     []float64
	headTime time.Time
}

func (d *BufferDimension) Buffer(data []float64) {

}

// NewBufferDimension creates a dimension that can have data buffered for it to playback in real time.
func (mgr *dimensionManager) NewBufferDimension(name string, symbol Symbol) *Dimension {
	d := Universe.Dimensions.NewDimension(name, symbol, 0)
	resolutionStep := time.Duration(1 / Universe.Resolution)

	go func() {
		lastUpdate := time.Now()
		for {
			if Universe.Terminate {
				break
			}

			if time.Since(lastUpdate) > resolutionStep {

			}
		}
		Universe.Printf(mgr, "[%s] stopped", string(symbol))
	}()

	return d
}
