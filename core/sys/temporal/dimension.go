package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"sync"
	"time"
)

// Dimension is a way of observing a target value across time, limited to a window of observance.
type Dimension[TValue any, TCache any] struct {
	core.NamedEntity

	// Current is the currently held value of this dimension.
	Current *std.Data[TValue]

	// Cache is a place where a looping stimulator can save information for the next activation of the loop.
	Cache *TCache

	// Timeline is the historical values of this dimension.
	Timeline []std.Data[TValue]

	// Window is the duration to hold onto recorded values for.
	Window time.Duration

	// Mutex should be locked for any operations that need a momentary snapshot of the timeline.
	Mutex sync.Mutex

	// Stimulator is the neuron that drives the function that populates this timeline.
	Stimulator *core.Neuron

	// Trimmer is the neuron that trims the timeline of entries beyond the window of observance.
	Trimmer *core.Neuron

	// Destroyed indicates if this dimension has been destroyed.
	Destroyed bool

	// HostAliveFunc can be set to override the default IsHostAlive check.
	HostAliveFunc func() bool

	// lastCycle is used by integration to locate timeline indexes.
	lastCycle time.Time
}

// IsHostAlive returns whether the system that hosts this dimension is alive.
//
// By default, this returns core.Alive.  If you would like different logic, you can provide
// a function to the dimension's HostAliveFunc field.
func (d *Dimension[TValue, TCache]) IsHostAlive() bool {
	if d.HostAliveFunc != nil {
		return core.Alive
	}
	return d.HostAliveFunc()
}

// Write circumvents the impulse engine and writes a value directly to the dimension.
func (d *Dimension[TValue, TCache]) Write(value TValue) {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()

	now := time.Now()
	period := now.Sub(d.Current.Moment)

	var ctx core.Context
	ctx.ID = core.NextID()
	ctx.Moment = now
	ctx.Period = period

	data := std.Data[TValue]{
		Context: ctx,
		Point:   value,
	}

	d.Current = &data
	d.Timeline = append(d.Timeline, data)
}

// Read returns a copy of the current timeline information.
func (d *Dimension[TValue, TCache]) Read() []std.Data[TValue] {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()
	result := make([]std.Data[TValue], len(d.Timeline))
	copy(result, d.Timeline)
	return result
}

// GetPastValue retrieves the value of a specific moment in time from the timeline.
func (d *Dimension[TValue, TCache]) GetPastValue(moment time.Time) *std.Data[TValue] {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()
	for _, v := range d.Timeline {
		if v.Moment == moment {
			return &v
		}
	}
	return nil
}

// GetClosestMoment retrieves the value of the closest moment in time to the provided target from the timeline,
// as well as the relative duration between the two.  The returned duration is calculated as the timeline
// value minus the target value.
func (d *Dimension[TValue, TCache]) GetClosestMoment(target time.Time) (*std.Data[TValue], time.Duration) {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()
	closest := &std.Data[TValue]{}
	closest.Moment = core.Inception
	closestOffset := closest.Moment.Sub(target)

	for _, v := range d.Timeline {
		offset := v.Moment.Sub(target)
		if core.AbsDuration(offset) < core.AbsDuration(closestOffset) {
			closest = &v
			closestOffset = offset
		}
	}
	return closest, closestOffset
}

// GetBeatValue retrieves the value of a specific beat from the timeline.
func (d *Dimension[TValue, TCache]) GetBeatValue(beat int) *std.Data[TValue] {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()
	for _, v := range d.Timeline {
		if v.Beat == beat {
			return &v
		}
	}
	return nil
}

// ImpulseTrim removes anything on the timeline that is older than the dimension's window of observance.
//
// This version allows an impulse to trigger the trimming operation - Please use Trim if calling directly.
func (d *Dimension[TValue, TCache]) ImpulseTrim(ctx core.Context) {
	d.Trim()
}

// Trim removes anything on the timeline that is older than the dimension's window of observance.
func (d *Dimension[TValue, TCache]) Trim() {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()

	var trimCount int
	for _, v := range d.Timeline {
		if time.Now().Sub(v.Moment) < d.Window {
			break
		}
		trimCount++
	}
	d.Timeline = d.Timeline[trimCount:]
}

// Destroy removes this dimension's neurons from the engine entirely.
func (d *Dimension[TValue, TCache]) Destroy() {
	d.Stimulator.Destroy()
	d.Trimmer.Destroy()
	d.Destroyed = true
}

// Mute suppresses the stimulator of this dimension.
func (d *Dimension[TValue, TCache]) Mute() {
	d.Stimulator.Muted = true
}

// Unmute un-suppresses the stimulator of this dimension.
func (d *Dimension[TValue, TCache]) Unmute() {
	d.Stimulator.Muted = false
}
