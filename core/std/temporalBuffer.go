package std

import (
	"sync"
	"time"

	"git.ignitelabs.net/janos/core/sys/atlas"
)

// A TemporalBuffer is a type of buffer that holds data for a period of time, rather than up to a fixed size.
// The default period for temporal buffers is defined through atlas.ObservanceWindow.
//
// There are a few notable features of a temporal buffer:
//
// 0 - When you Yield() the buffer contents, you're provided a thread-independent copy of the underlying buffer
//
// 1 - I/O of the buffer is automatically handled in a thread-safe fashion
//
// 2 - You are not required to add elements temporally sequentially - instead, you provide the moment to add
//
// 3 - Temporal buffers automatically trim on access
type TemporalBuffer[T any] struct {
	buffer []instant[T]

	Window *time.Duration

	master sync.Mutex
}

// NewTemporalBuffer creates a new instance of a temporal buffer which observes the provided window of time.  If no
// window is provided, this will default to atlas.ObservanceWindow
func NewTemporalBuffer[T any](window ...*time.Duration) TemporalBuffer[T] {
	w := &atlas.ObservanceWindow
	if len(window) > 0 {
		w = window[0]
	}
	return TemporalBuffer[T]{
		buffer: make([]instant[T], 0),
		Window: w,
	}
}

type instant[T any] struct {
	Moment  time.Time
	Element T
}

func (b *TemporalBuffer[T]) sanityCheck() {
	if b.buffer == nil {
		panic("temporal buffer set to nil - please create these through std.NewTemporalBuffer")
	}
}

func (b *TemporalBuffer[T]) trim() {
	b.sanityCheck()
	now := time.Now()
	cutoff := now.Add(-*b.Window)

	var i int
	for _, inst := range b.buffer {
		if inst.Moment.After(cutoff) {
			break
		}
		i++
	}
	maximum := len(b.buffer) - int(atlas.ObservedMinimum)
	if maximum < 0 {
		maximum = 0
	}
	if i > maximum {
		i = maximum
	}

	b.buffer = b.buffer[i:]
}

func (b *TemporalBuffer[T]) Len() uint {
	return uint(len(b.buffer))
}

func (b *TemporalBuffer[T]) LatestSince(moment time.Time) []instant[T] {
	b.sanityCheck()
	out := b.Yield()
	point := -1
	for i := len(out) - 1; i >= 0; i-- {
		if out[i].Moment.After(moment) {
			point = i
		} else {
			break
		}
	}
	if point == -1 {
		return []instant[T]{}
	}
	return out[point:]
}

// Latest returns the most recent elements in logical temporal order up to the provided depth.  If a negative depth is provided, all elements are returned.
// If no depth is provided, a depth of '1' is implied.
func (b *TemporalBuffer[T]) Latest(depth ...int) []instant[T] {
	d := 1
	if len(depth) > 0 {
		d = depth[0]
	}

	b.sanityCheck()
	out := b.Yield()
	if d <= 0 {
		return out
	}
	end := len(out) - d
	if end < 0 {
		end = 0
	}
	out = out[end:]
	if len(out) > d {
		panic("too many elements")
	}
	return out
}

func (b *TemporalBuffer[T]) Yield() []instant[T] {
	b.sanityCheck()
	b.master.Lock()
	defer b.master.Unlock()

	b.trim()
	out := make([]instant[T], len(b.buffer))
	copy(out, b.buffer)
	return out
}

func (b *TemporalBuffer[T]) Add(moment time.Time, element T) {
	b.sanityCheck()
	b.master.Lock()
	defer b.master.Unlock()

	t := len(b.buffer) - 1
	for i := len(b.buffer) - 1; i >= 0; i-- {
		if b.buffer[i].Moment.Before(moment) {
			break
		}
		t--
	}
	t++
	if t < 0 {
		t = 0
	}
	left := b.buffer[:t]
	right := b.buffer[t:]
	left = append(left, instant[T]{moment, element})
	left = append(left, right...)
	b.buffer = left

	b.trim()
}
