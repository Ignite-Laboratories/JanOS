package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
)

// Reaction creates a dimension that calls the reaction function if the provided potential returns true.
//
// NOTE: The potential function gates the creation of timeline indexes!
// This can adjust the "resolution" of output data =)
//
// Muted indicates if the stimulator of this dimension should be created muted.
func Reaction[TValue any](engine *core.Engine, potential core.Potential, muted bool, target std.TargetFunc[TValue], reaction Change[TValue]) *Dimension[TValue, any] {
	d := Dimension[TValue, any]{}
	d.NamedEntity = core.NewNamedEntity()
	d.Window = core.DefaultObservanceWindow
	d.Trimmer = engine.Loop(d.ImpulseTrim, when.Frequency(&core.TrimFrequency), false)
	f := func(ctx core.Context) {
		data := std.Data[TValue]{
			Context: ctx,
			Point:   *target(),
		}
		d.Mutex.Lock()
		var old *std.Data[TValue]
		if d.Current != nil {
			old = d.Current
		}
		d.Timeline = append(d.Timeline, data)
		d.Current = &data
		// Don't 'react' to the first impulse
		if old != nil {
			reaction(ctx, *old, *d.Current)
		}
		d.Mutex.Unlock()
	}
	d.Stimulator = engine.Stimulate(f, potential, muted)
	return &d
}
