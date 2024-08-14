package Logic

import (
	"sync"
	"time"
)

// Drivable represents a system that can be driven at an interval defined by the hosting architecture.
type Drivable interface {
	Tick()
}

// Drive is a blocking method to fire off goroutines that perform sampling at a fixed interval.
func Drive(interval time.Duration, toDrive ...Drivable) {
	// On a fixed interval...
	//ticker := time.NewTicker(interval)
	for { //range ticker.C {

		// Queue up the samplers...
		var wg sync.WaitGroup
		for _, sampler := range toDrive {
			wg.Add(1)
			go func() {
				defer wg.Done()
				sampler.Tick()
			}()
		}

		// Asynchronously fire them all
		wg.Wait()
	}
}
