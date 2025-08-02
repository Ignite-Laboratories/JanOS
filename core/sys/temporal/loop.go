package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
)

// Loop creates an dimension that calls a function in a looping fashion while
// observing its own runtime information, if the provided potential returns true.
//
// NOTE: The potential function gates the creation of timeline indexes!
//
// Muted indicates if the stimulator of this dimension should be created muted.
func Loop(engine *core.Engine, potential core.Potential, muted bool, target core.Action) *Dimension[core.Runtime, any] {
	d := Dimension[core.Runtime, any]{}
	d.NamedEntity = core.NewNamedEntity()
	d.Window = core.DefaultObservanceWindow
	d.Trimmer = engine.Loop(d.ImpulseTrim, when.Frequency(&core.TrimFrequency), false)
	f := func(ctx core.Context) {
		data := std.Data[core.Runtime]{
			Context: ctx,
			Point:   d.Stimulator.LastActivation,
		}
		target(ctx)
		d.Mutex.Lock()
		d.Timeline = append(d.Timeline, data)
		d.Current = &data
		d.Mutex.Unlock()
	}
	d.Stimulator = engine.Loop(f, potential, muted)
	return &d
}
