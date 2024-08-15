package JanOS

import (
	"log"
	"math"
	"time"
)

type Buffer struct {
	StartTime time.Time
	Data      []float64
}

type PerceptionBuffer struct {
	data     []float64
	headTime time.Time
}

func (d *PerceptionBuffer) getRelativeIndex(t time.Time) int {
	return ToIndexCount(t.Sub(d.headTime))
}

func ToIndexCount(duration time.Duration) int {
	// Calculate the amount of nanoseconds per index
	nanoStep := float64(time.Second.Nanoseconds()) / Universe.Resolution
	// Divide the provided duration's nanoseconds by the step size to get the index count
	return int(float64(duration.Nanoseconds()) / nanoStep)
}

func ToDuration(steps int) time.Duration {
	// Calculate the amount of nanoseconds per index
	nanoStep := float64(time.Second.Nanoseconds()) / Universe.Resolution
	// Multiply the amount of nanoseconds per step by the number of steps to get the duration in nanoseconds
	return time.Duration(nanoStep * float64(steps))
}

func (d *PerceptionBuffer) GetNow() float64 {
	return d.data[d.getRelativeIndex(time.Now())]
}

func (d *PerceptionBuffer) SliceEntireFuture() Buffer {
	now := time.Now()
	relativeIndex := d.getRelativeIndex(now)
	return Buffer{
		StartTime: now,
		Data:      d.data[relativeIndex:],
	}
}

func (d *PerceptionBuffer) SliceFromNow(steps int) Buffer {
	now := time.Now()
	nowIndex := d.getRelativeIndex(now)
	relativeIndex := nowIndex + steps
	lower := int(math.Min(float64(nowIndex), float64(relativeIndex)))
	upper := int(math.Max(float64(nowIndex), float64(relativeIndex)))
	return Buffer{
		StartTime: now,
		Data:      d.data[lower:upper],
	}
}

func (d *PerceptionBuffer) SliceIntoPast(duration time.Duration) Buffer {
	now := time.Now()
	nowIndex := d.getRelativeIndex(now)
	pastIndex := nowIndex - ToIndexCount(duration)
	return Buffer{
		StartTime: now.Add(-duration),
		Data:      d.data[pastIndex:nowIndex],
	}

}

func (d *PerceptionBuffer) SliceIntoFuture(duration time.Duration) Buffer {
	now := time.Now()
	nowIndex := d.getRelativeIndex(now)
	futureIndex := nowIndex + ToIndexCount(duration)
	return Buffer{
		StartTime: now,
		Data:      d.data[nowIndex:futureIndex],
	}

}

// Buffer seeks to the appropriate position in time and additively introduces the provided data to the buffer
func (d *PerceptionBuffer) Buffer(instant time.Time, data []float64) {
	startIndex := d.getRelativeIndex(instant)

	for x := 0; x < len(data); x++ {
		// If we are out of bounds of the underlying buffer...
		if startIndex+x > len(d.data) {
			// ...append the remaining data to the buffer and exit the loop
			toAppend := data[x:]
			d.data = append(d.data, toAppend...)
			break
		}

		// Otherwise, just add the data to the existing data
		d.data[startIndex+x] = d.data[startIndex+x] + data[x]
	}
}

// NewPerceptionBuffer creates a dimension that can have data buffered for it to playback in real time.
// The duration represents the total amount of time to buffer, with now being considered relative to
// the midpoint of the buffer.  The frequency tells it how often to trim/append the buffer in time.
func (mgr *dimensionManager) NewPerceptionBuffer(name string, symbol Symbol, duration time.Duration, frequency time.Duration) *PerceptionBuffer {
	durationInIndex := ToIndexCount(duration)
	bd := &PerceptionBuffer{
		// Set the head time relative to the midpoint of the buffer
		// (we initialize with the past being empty, essentially)
		headTime: time.Now().Add(-(duration / 2)),
		data:     make([]float64, durationInIndex),
	}
	mgr.bufferDimensions[name] = bd
	Universe.Printf(mgr, "Let '%s' [%s] = []float64(%fs)", name, string(symbol), duration.Seconds())

	go func() {
		for {
			if Universe.Terminate {
				break
			}

			// Save off now so the logic operates on the same instant during this routine
			now := time.Now()
			// Calculate new head time
			newHead := now.Add(-(duration / 2))
			delta := bd.getRelativeIndex(newHead)
			log.Println(delta)
			// Update the head time
			bd.headTime = newHead
			// Slice off the past that has passed
			bd.data = bd.data[delta:]
			// Append new data to replace the old data
			bd.data = append(bd.data, make([]float64, delta)...)

			time.Sleep(frequency)
		}
		Universe.Printf(mgr, "[%s] stopped", string(symbol))
	}()

	return bd
}
