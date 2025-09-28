package std

import (
	"sync"
	"time"

	"git.ignitelabs.net/janos/core/sys/atlas"
	"git.ignitelabs.net/janos/core/sys/num/tiny"
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
func NewTemporalBuffer[T any](window ...*time.Duration) *TemporalBuffer[T] {
	w := &atlas.ObservanceWindow
	if len(window) > 0 {
		w = window[0]
	}
	return &TemporalBuffer[T]{
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
	if b.Window == nil {
		b.Window = &atlas.ObservanceWindow
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

// LatestSince will grab the latest elements after the provided moment in time, exclusively.  If you'd like to include the
// requested moment, you can set the optional 'includeMoment' parameter to true.  When performing integration and
// differentiation, you'll want to include the moment - otherwise, it makes more sense to exclude it if simply recording
// temporal signal data.
func (b *TemporalBuffer[T]) LatestSince(moment time.Time, includeMoment ...bool) []instant[T] {
	b.sanityCheck()
	include := len(includeMoment) > 0 && includeMoment[0]

	out := b.Yield()
	point := -1
	included := false
	for i := len(out) - 1; i >= 0; i-- {
		if out[i].Moment.After(moment) {
			point = i
		} else if include && !included {
			included = true
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

func (b *TemporalBuffer[T]) Record(moment time.Time, element T) {
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

// Calculate will yield the Latest(depth) - then, for each of the elements from oldest to newest, call calcFn(dt, element)
// before returning the accumulated results.
func (b *TemporalBuffer[T]) Calculate(depth int, calcFn func(time.Duration, T) any) []instant[any] {
	b.sanityCheck()
	//yield := b.Latest(depth)
	return nil
}

// CalculateSince will yield the LatestSince(moment, includeMoment) - then, for each of the elements from oldest to newest, call
// calcFn(dt, element) before returning the accumulated results.
//
// NOTE: For integration and differentiation, you'll want to include the provided moment in the calculation to create a continuous calculation =)
func (b *TemporalBuffer[T]) CalculateSince(moment time.Time, calcFn func(time.Duration, T) any, includeMoment ...bool) []instant[any] {
	b.sanityCheck()
	//yield := b.LatestSince(moment)
	return nil
}

// Integrate will perform standard temporal integration against the provided depth of elements.  This will yield the area
// between each moment and the total calculated area.  If you'd like to implement your own integration logic,
// please leverage Calculate.
//
// NOTE: If the elements are NOT implicitly parseable as defined by tiny.FilterOperands, this will panic.  In that case,
// please provide a 'parseFn' which translates the buffered information into a parseable type.
func (b *TemporalBuffer[T]) Integrate(base uint16, depth int, parseFn ...func(T) any) ([]instant[any], float64) {
	b.sanityCheck()
	area := 0.0
	last := 0.0
	return b.Calculate(depth, func(dt time.Duration, element T) any {
		var number any
		if len(parseFn) > 0 {
			number = parseFn[0](element)
		} else {
			number = element
		}
		if dt < 0 {
			last = 0.0
			return 0.0
		}
		last = tiny.ParseAs[float64](tiny.Multiply(dt, tiny.Add(number, last), 0.5))
		area = tiny.ParseAs[float64](tiny.Add(area, last))
		return last
	}), area
}

func (b *TemporalBuffer[T]) IntegrateSince(moment time.Time, parseFn ...func(T) any) []instant[T] {
	b.sanityCheck()
	// NOTE: This is INCLUSIVE of the provided moment =)
	return nil
}

func (b *TemporalBuffer[T]) Differentiate(depth int, parseFn ...func(T) any) []instant[T] {
	b.sanityCheck()
	return nil
}

func (b *TemporalBuffer[T]) DifferentiateSince(moment time.Time, parseFn ...func(T) any) []instant[T] {
	b.sanityCheck()
	// NOTE: This is INCLUSIVE of the provided moment =)
	return nil
}
