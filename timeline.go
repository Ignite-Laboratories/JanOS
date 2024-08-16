package JanOS

import (
	"sync"
	"time"
)

// NOTE:
// This system is meant to be performant and carefully exposed.
// We intentionally fail hard - if you are trying to look too far into the future or past then
// you need to restructure your model!  Look at the Universe.BufferLength to know what the practical
// boundaries are in your simulation and modify it if you need a different stretch of time.

type timeline struct {
	value    float64
	data     []float64
	headTime time.Time
	lock     sync.Mutex
}

// GetRelativeIndex takes a moment in time and gets its index, relative to the current head time.
func (tl *timeline) GetRelativeIndex(t time.Time) int {
	return ToIndexCount(t.Sub(tl.headTime))
}

// GetInstant returns the value on the timeline at a moment in time.
func (tl *timeline) GetInstant(instant time.Time) float64 {
	return tl.data[tl.GetRelativeIndex(instant)]
}

// SliceEntireFuture returns the remainder of the buffer from the provided instant in time.
func (tl *timeline) SliceEntireFuture(instant time.Time) TimeSlice {
	instantIndex := tl.GetRelativeIndex(instant)
	return TimeSlice{
		StartTime: instant,
		Data:      tl.data[instantIndex:],
	}
}

// SliceEntirePast returns the entire buffer up to the provided instant in time.
func (tl *timeline) SliceEntirePast(instant time.Time) TimeSlice {
	// We capture the head time here to ensure all calculations
	// are relative to the execution of this line of code in time.
	headTime := tl.headTime
	headIndex := tl.GetRelativeIndex(headTime)
	instantIndex := tl.GetRelativeIndex(instant)
	return TimeSlice{
		StartTime: headTime,
		Data:      tl.data[headIndex:instantIndex],
	}
}

// SliceFuture returns a slice of the future from an instant in time.
func (tl *timeline) SliceFuture(instant time.Time, duration time.Duration) TimeSlice {
	instantIndex := tl.GetRelativeIndex(instant)
	futureIndex := instantIndex + ToIndexCount(duration)
	return TimeSlice{
		StartTime: instant,
		Data:      tl.data[instantIndex:futureIndex],
	}
}

// SlicePast returns a slice of the past from an instant in time.
func (tl *timeline) SlicePast(instant time.Time, duration time.Duration) TimeSlice {
	instantIndex := tl.GetRelativeIndex(instant)
	pastIndex := instantIndex - ToIndexCount(duration)
	return TimeSlice{
		StartTime: instant.Add(-duration),
		Data:      tl.data[pastIndex:instantIndex],
	}
}

// AddValues seeks to the appropriate position in time and additively introduces the provided data to the buffer
func (tl *timeline) AddValues(instant time.Time, data ...float64) {
	tl.lock.Lock()
	defer tl.lock.Unlock()
	startIndex := tl.GetRelativeIndex(instant)

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

// newTimeline creates a buffer.
// The duration represents the total amount of time to buffer, with now being considered relative to
// the midpoint of the buffer.  The frequency tells it how often to trim/append the buffer in time.
func (mgr *dimensionManager) newTimeline(name string, symbol Symbol, defaultValue float64, duration time.Duration) *timeline {
	durationInIndex := ToIndexCount(duration)
	tl := &timeline{
		// Set the head time relative to the midpoint of the buffer
		// (we initialize with the past being empty, essentially)
		headTime: time.Now().Add(-(duration / 2)),
		data:     NewInitializedArray(defaultValue, durationInIndex),
		value:    defaultValue,
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

			// If we are above the minimum cycle time
			if now.Sub(lastUpdate) > Universe.BufferFrequency {
				tl.lock.Lock()
				// Calculate new head time
				newHead := now.Add(-(duration / 2))
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
		Universe.Printf(mgr, "%s [%s] timeline stopped", name, string(symbol))
	}()

	return tl
}
