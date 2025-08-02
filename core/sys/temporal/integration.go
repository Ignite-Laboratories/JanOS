package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
)

// Integration creates a new dimension that performs calculation on sets of temporal data cyclically.
// The integral function is always called with the exact timeline data since the last analysis started.
//
// NOTE: The potential function gates analysis.
// This can adjust "reactivity" to input data =)
//
// Muted indicates if the stimulator of this dimension should be created muted.
//
// Impulsive indicates if the stimulator of this dimension should activate impulsively, or as a loop.
//
// If it activates impulsively this will ensure the results are placed accurately on the timeline.
func Integration[TSource any, TValue any, TCache any](engine *core.Engine, potential core.Potential, muted bool, impulsive bool, integrate Integral[std.Data[TSource], TValue, TCache], target *Dimension[TSource, any]) *Dimension[TValue, TCache] {
	d := Dimension[TValue, TCache]{}
	d.NamedEntity = core.NewNamedEntity()
	d.Window = core.DefaultObservanceWindow
	d.Trimmer = engine.Loop(d.ImpulseTrim, when.Frequency(&core.TrimFrequency), false)
	d.lastCycle = core.Inception
	d.Cache = new(TCache)

	f := func(ctx core.Context) {
		// Get target timeline data
		target.Mutex.Lock()
		data := make([]std.Data[TSource], len(target.Timeline))
		copy(data, target.Timeline)
		target.Mutex.Unlock()

		// Trim any indices that were handled by the last analysis
		var trimCount int
		for _, v := range data {
			if v.Moment.After(d.lastCycle) {
				break
			}
			trimCount++
		}
		data = data[trimCount:]

		// Save off the last moment for the next cycle
		lastCycle := d.lastCycle
		if len(data) > 0 {
			lastCycle = data[len(data)-1].Moment
		}

		// Perform integration
		point := integrate(ctx, d.Cache, data)
		out := std.Data[TValue]{
			Context: ctx,
			Point:   point,
		}

		// Record the result
		d.Mutex.Lock()
		d.lastCycle = lastCycle
		if !impulsive {
			// Integration execution is logically ordered - just append
			d.Timeline = append(d.Timeline, out)
			d.Current = &out
		} else {
			// Integration execution is chaotically ordered - inject appropriately
			var left []std.Data[TValue]
			right := d.Timeline
			for i := len(d.Timeline) - 1; i >= 0; i-- {
				if out.Moment.After(d.Timeline[i].Moment) {
					left = d.Timeline[:i]
					right = d.Timeline[i:]
					break
				}
			}
			d.Timeline = left
			d.Timeline = append(d.Timeline, out)
			d.Timeline = append(d.Timeline, right...)
			if d.Current.Moment.Before(out.Moment) {
				d.Current = &out
			}
		}
		d.Mutex.Unlock()
	}

	if impulsive {
		d.Stimulator = engine.Loop(f, potential, muted)
	} else {
		d.Stimulator = engine.Stimulate(f, potential, muted)
	}
	return &d
}
