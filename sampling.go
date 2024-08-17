package JanOS

import (
	"time"
)

// SamplePoint samples the provided signal at the provided resolution frequency
// and calls onObservation with the internally stored buffer data.  This sampling
// variant only grabs a single point of data in time on each observation.  This
// would best be used for low resolution triggering.
func (signal *Signal) SamplePoint(frequency int, observer Sampler) {
	Universe.Printf(signal, "Point sampling [%s] at %dhz", string(signal.Symbol), frequency)
	r := newResolution(frequency)
	go signal.sample(r, r.Duration, observer)
}

// Sample samples the provided signal at the provided resolution frequency
// and calls onObservation with the internally stored buffer data.  This sampling
// variant samples at the provided frequency for a specified duration in time
// before calling onObservation.  This would best be used for duty cycle processing
// of complicated sets of data.
func (signal *Signal) Sample(frequency int, duration time.Duration, sampler Sampler) {
	Universe.Printf(sampler, "Sampling %s at %dhz on a %v duty cycle", string(signal.Symbol), frequency, duration)
	r := newResolution(frequency)
	go signal.sample(r, duration, sampler)
}

func (signal *Signal) sample(r resolution, duration time.Duration, sampler Sampler) {
	lastUpdate := time.Now()
	headTime := lastUpdate
	buffer := make([]PointValue, r.ToIndex(duration))
	i := 0

	for {
		now := time.Now()
		if Universe.Terminate {
			break
		}

		// Sample after the approximate amount of time has passed
		if now.Sub(lastUpdate) >= r.Duration {
			buffer[i] = signal.GetInstantValue(now)
			lastUpdate = now
			i++
		}

		// If we reach the end of our current buffer...
		if i == len(buffer) {
			// ...output the buffer data
			output := TimeSlice{
				StartTime:  headTime,
				Data:       buffer,
				Resolution: r,
			}
			go sampler.OnSample(signal, output)

			// ...and reset the buffer
			headTime = now
			buffer = make([]PointValue, r.ToIndex(duration))
			i = 0
		}
	}
	Universe.Printf(signal, "%s sampling stopped", string(signal.Symbol))
}
