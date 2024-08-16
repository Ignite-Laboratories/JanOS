package JanOS

import (
	"sync"
	"time"
)

// NOTE:
// This system is meant to be performant and has been very carefully exposed.
// We intentionally fail hard - if you are trying to look too far into the future or past then
// you need to restructure your model!  Look at the Universe.BufferLength to know what the practical
// boundaries are in your simulation and modify it if you need a different stretch of time.

type timeline struct {
	value      float64
	data       []float64
	headTime   time.Time
	lock       sync.Mutex
	resolution resolution
}

type resolution struct {
	Frequency   int
	Nanoseconds int64
	Duration    time.Duration
}

func newResolution(frequency int) resolution {
	r := resolution{
		Frequency: frequency,
	}
	r.Nanoseconds = int64(float64(time.Second.Nanoseconds()) / float64(r.Frequency))
	r.Duration = time.Duration(r.Nanoseconds + 1)
	return r
}

// SetResolution updates the timeline's resolution.  Any existing data on the timeline
// will be read in the new resolution interval - making this essentially stretch or
// shrink the existing timeline.  As such, the best practice is to set this before
// starting to put data on the timeline - but that doesn't mean you shouldn't try!
func (tl *timeline) SetResolution(frequency int) {
	tl.lock.Lock()
	defer tl.lock.Unlock()
	calculatedResolution := newResolution(frequency)
	tl.resolution.Frequency = calculatedResolution.Frequency
	tl.resolution.Nanoseconds = calculatedResolution.Nanoseconds
	tl.resolution.Duration = calculatedResolution.Duration
}

// GetResolution returns the current resolution information for the provided timeline.
func (tl *timeline) GetResolution() resolution {
	return tl.resolution
}

// timeSlice represents a slice of data relative to a moment in time.
type timeSlice struct {
	StartTime  time.Time
	Data       []float64
	Resolution resolution
}

//func (ts *timeSlice) Upsample(frequency int) timeSlice {
//
//	toReturn := timeSlice{
//		StartTime:  ts.StartTime,
//		Data:       make([]float64, 0),
//		resolution: newResolution(frequency),
//	}
//
//	lastDataPoint := ts.Data[0]
//
//	for i := 1; i < len(ts.Data); i++ {
//		currentDataPoint := ts.Data[i]
//		indicesToInterpolate := frequency - ts.resolution.Frequency
//		stride := (math.Abs(currentDataPoint) - math.Abs(lastDataPoint)) / float64(indicesToInterpolate)
//
//		toReturn.Data = append(toReturn.Data, currentDataPoint)
//		for x := 0; x < indicesToInterpolate; x++ {
//
//		}
//		lastDataPoint = currentDataPoint
//	}
//	return toReturn
//}

// Sample returns back a less resolute version of an existing timeSlice.
// NOTE: You cannot sample more data than is available in the system, by
// design.  This method provides a way to convert higher resolution data
// into lower resolution data for different kinds of mathematical calculations.
func (ts *timeSlice) Sample(frequency int) timeSlice {
	stride := int(float64(time.Second.Nanoseconds()) / float64(frequency))

	toReturn := timeSlice{
		StartTime:  ts.StartTime,
		Data:       make([]float64, 0),
		Resolution: newResolution(frequency),
	}

	for i, _ := range ts.Data {
		// if we are at a sample index...
		if i == 0 || i%stride == 0 {
			tsIndex := i * stride
			if tsIndex < len(ts.Data) {
				toReturn.Data = append(toReturn.Data, ts.Data[i*stride])
			}
		}
	}
	return toReturn
}

// ToIndexCount returns an index representation of the duration of time provided.
func (tl *timeline) ToIndexCount(duration time.Duration) int {
	// Calculate the amount of nanoseconds per index
	nanoStep := float64(time.Second.Nanoseconds()) / float64(tl.resolution.Frequency)
	// Divide the provided duration's nanoseconds by the step size to get the index count
	return int(float64(duration.Nanoseconds()) / nanoStep)
}

// ToDuration returns a time.Duration representation of the amount of indices provided
// Check Universe.Resolution for the operational frequency.
func (tl *timeline) ToDuration(steps int) time.Duration {
	// Calculate the amount of nanoseconds per index
	nanoStep := float64(time.Second.Nanoseconds()) / float64(tl.resolution.Frequency)
	// Multiply the amount of nanoseconds per step by the number of steps to get the duration in nanoseconds
	return time.Duration(nanoStep * float64(steps))
}

// GetRelativeIndex takes a moment in time and gets its index, relative to the current head time.
func (tl *timeline) GetRelativeIndex(t time.Time) int {
	return tl.ToIndexCount(t.Sub(tl.headTime))
}

// GetInstant returns the value on the timeline at a moment in time.
func (tl *timeline) GetInstant(instant time.Time) float64 {
	return tl.data[tl.GetRelativeIndex(instant)]
}

// SliceEntireFuture returns the remainder of the buffer from the provided instant in time.
func (tl *timeline) SliceEntireFuture(instant time.Time) timeSlice {
	instantIndex := tl.GetRelativeIndex(instant)
	return timeSlice{
		StartTime:  instant,
		Data:       tl.data[instantIndex:],
		Resolution: tl.resolution,
	}
}

// SliceEntirePast returns the entire buffer up to the provided instant in time.
func (tl *timeline) SliceEntirePast(instant time.Time) timeSlice {
	// We capture the head time here to ensure all calculations
	// are relative to the execution of this line of code in time.
	headTime := tl.headTime
	headIndex := tl.GetRelativeIndex(headTime)
	instantIndex := tl.GetRelativeIndex(instant)
	return timeSlice{
		StartTime:  headTime,
		Data:       tl.data[headIndex:instantIndex],
		Resolution: tl.resolution,
	}
}

// SliceFutureDuration returns a slice of the future from an instant in time utilizing a time.Duration.
func (tl *timeline) SliceFutureDuration(instant time.Time, duration time.Duration) timeSlice {
	instantIndex := tl.GetRelativeIndex(instant)
	futureIndex := instantIndex + tl.ToIndexCount(duration)
	return timeSlice{
		StartTime:  instant,
		Data:       tl.data[instantIndex:futureIndex],
		Resolution: tl.resolution,
	}
}

// SliceFutureIndices returns an indexCount length slice of the future from an instant in time.
func (tl *timeline) SliceFutureIndices(instant time.Time, indexCount int) timeSlice {
	instantIndex := tl.GetRelativeIndex(instant)
	return timeSlice{
		StartTime:  instant,
		Data:       tl.data[instantIndex : instantIndex+indexCount],
		Resolution: tl.resolution,
	}
}

// SlicePastDuration returns an indexCount length slice of the past up to an instant in time.
func (tl *timeline) SlicePastDuration(instant time.Time, duration time.Duration) timeSlice {
	instantIndex := tl.GetRelativeIndex(instant)
	pastIndex := instantIndex - tl.ToIndexCount(duration)
	return timeSlice{
		StartTime:  instant.Add(-duration),
		Data:       tl.data[pastIndex:instantIndex],
		Resolution: tl.resolution,
	}
}

// SlicePastIndices returns a indexCount length slice of the past up to an instant in time.
func (tl *timeline) SlicePastIndices(instant time.Time, indexCount int) timeSlice {
	instantIndex := tl.GetRelativeIndex(instant)
	return timeSlice{
		StartTime:  instant,
		Data:       tl.data[instantIndex-indexCount : instantIndex],
		Resolution: tl.resolution,
	}
}

// AddValues seeks to the appropriate position in time and additively introduces the provided data to the buffer
func (tl *timeline) AddValues(instant time.Time, data ...float64) {
	tl.lock.Lock()
	defer tl.lock.Unlock()
	startIndex := tl.GetRelativeIndex(instant)

	// If the buffer has no data between the end of its current prediction and the start of this data...
	tlDataLen := len(tl.data)
	if startIndex >= tlDataLen {
		// ...fill the gap with initialized data to the current value of the timeline
		toFill := startIndex - tlDataLen
		initialized := NewInitializedArray(tl.value, toFill)
		tl.data = append(tl.data, initialized...)
	}

	for x := 0; x < len(data); x++ {
		// If we are out of bounds of the underlying buffer...
		if startIndex+x > len(tl.data) {
			// ...append the remaining data to the buffer and exit the loop
			toAppend := data[x:]
			tl.data = append(tl.data, toAppend...)
			break
		}

		// Otherwise, just add the data to the existing data
		tl.data[startIndex+x] = tl.data[startIndex+x] + data[x]
	}
}

// setValue seeks to the appropriate position in time and replaces the values on the remainder of the buffer.
// It is not exposed because you should always set the value of the dimension from the dimension - this is a method
// that keeps the  timeline logic together.  It keeps from spreading its implementation details into the dimension
// code through which you are intended to access it.
func (tl *timeline) setValue(instant time.Time, value float64) {
	tl.lock.Lock()
	defer tl.lock.Unlock()
	tl.value = value
	startIndex := tl.GetRelativeIndex(instant)
	remainingLen := len(tl.data) - startIndex

	// Split off the array and then inject the new value (roughly) for the remaining length
	newData := NewInitializedArray(value, remainingLen)
	tl.data = append(tl.data[:startIndex], newData...)
}

// newTimeline creates a timeline buffer.
// The duration represents the total amount of time to buffer, with now being considered relative to
// the midpoint of the buffer.  The frequency tells it how often to trim/append the buffer in time.
func (d *Dimension) newTimeline(defaultValue float64) *timeline {
	nanoStep := float64(time.Second.Nanoseconds()) / float64(Universe.StdResolution)
	durationInIndex := int(float64(Universe.StdBufferLength.Nanoseconds()) / nanoStep)
	tl := &timeline{
		// Set the head time relative to the midpoint of the buffer
		// (we initialize with the past being empty, essentially)
		headTime:   time.Now().Add(-(Universe.StdBufferLength / 2)),
		data:       NewInitializedArray(defaultValue, durationInIndex),
		value:      defaultValue,
		resolution: newResolution(Universe.StdResolution),
	}

	// Spin off the loop for this dimension's timeline
	go func() {
		lastUpdate := time.Now()
		for {
			if Universe.Terminate {
				break
			}

			// Save off now so the logic operates on the same instant during this routine
			now := time.Now()

			// If we are above the minimum cycle time (arbitrarily 10ms)
			if now.Sub(lastUpdate) > time.Millisecond*10 {
				tl.lock.Lock()
				// Calculate new head time
				newHead := now.Add(-(Universe.StdBufferLength / 2))
				delta := tl.GetRelativeIndex(newHead)
				// Update the head time
				tl.headTime = newHead
				// Slice off the past that has passed
				tl.data = tl.data[delta:]
				// Append new data to replace the old data
				newData := NewInitializedArray(tl.value, delta)
				tl.data = append(tl.data, newData...)
				tl.lock.Unlock()
			}
			time.Sleep(5)
		}
		Universe.Printf(d, "%s [%s] timeline stopped", d.Name, string(d.Symbol))
	}()

	return tl
}
