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

type BufferDimension struct {
	data     []float64
	headTime time.Time
}

func (d *BufferDimension) getRelativeIndex(t time.Time) int {
	return int(float64(t.Sub(d.headTime).Nanoseconds()) * Universe.Resolution)
}

func ToIndex(duration time.Duration) int {
	return int(float64(duration.Nanoseconds()) * Universe.Resolution)
}

func (d *BufferDimension) GetNow() float64 {
	return d.data[d.getRelativeIndex(time.Now())]
}

func (d *BufferDimension) SliceFromNow(steps int) Buffer {
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

func (d *BufferDimension) SliceIntoPast(duration time.Duration) Buffer {
	now := time.Now()
	nowIndex := d.getRelativeIndex(now)
	pastIndex := nowIndex - ToIndex(duration)
	return Buffer{
		StartTime: now.Add(-duration),
		Data:      d.data[pastIndex:nowIndex],
	}

}

func (d *BufferDimension) SliceIntoFuture(duration time.Duration) Buffer {
	now := time.Now()
	nowIndex := d.getRelativeIndex(now)
	futureIndex := nowIndex + ToIndex(duration)
	return Buffer{
		StartTime: now,
		Data:      d.data[nowIndex:futureIndex],
	}

}

// Buffer seeks to the appropriate position in time and additively introduces the provided data to the buffer
func (d *BufferDimension) Buffer(instant time.Time, data []float64) {
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

// NewBufferDimension creates a dimension that can have data buffered for it to playback in real time.
func (mgr *dimensionManager) NewBufferDimension(name string, symbol Symbol, duration time.Duration) *BufferDimension {
	durationInIndex := ToIndex(duration)
	bd := &BufferDimension{
		headTime: time.Now(),
		data:     make([]float64, durationInIndex),
	}
	mgr.bufferDimensions[name] = bd
	Universe.Printf(mgr, "Let '%s' [%s] = []float64(%fs)", name, string(symbol), duration.Seconds())

	go func() {
		cutoffDuration := time.Duration(float64(duration.Nanoseconds()) * 0.1)
		cutoffDurationInIndex := ToIndex(cutoffDuration)
		for {
			if Universe.Terminate {
				break
			}

			delta := len(bd.data) - durationInIndex
			isPositive := delta > 0
			absDelta := int(math.Abs(float64(delta)))

			// If the delta is above the cutoff duration and positive...
			if absDelta > cutoffDurationInIndex {
				// ...cutoff the beginning data
				log.Println("removing")
				bd.data = bd.data[0:ToIndex(cutoffDuration)]
			}
			// If the delta is above the cutoff duration and negative...
			if !isPositive && absDelta > cutoffDurationInIndex {
				// ...append data to the end
				log.Println("appending")
				bd.data = append(bd.data, make([]float64, cutoffDuration)...)
			}
		}
		Universe.Printf(mgr, "[%s] stopped", string(symbol))
	}()

	return bd
}
