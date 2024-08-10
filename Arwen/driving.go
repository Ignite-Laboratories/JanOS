package Arwen

import (
	"sync"
	"time"
)

// Driver is a time interval firing system for samplers.
type Driver struct {
	Interval time.Duration
	ToDrive  []Drivable
}

func NewDriver(interval time.Duration, toDrive ...Drivable) *Driver {
	d := &Driver{
		Interval: interval,
		ToDrive:  toDrive,
	}
	return d
}

// Start is a blocking method to fire off goroutines that perform sampling at a fixed interval.
func (d *Driver) Start() {
	// On a fixed interval...
	ticker := time.NewTicker(d.Interval)
	for range ticker.C {

		// Queue up the samplers...
		var wg sync.WaitGroup
		for _, sampler := range d.ToDrive {
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
