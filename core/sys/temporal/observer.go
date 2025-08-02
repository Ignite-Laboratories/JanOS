package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
)

// Observer creates a dimension that records the target value across time, if the provided potential returns true.
//
// NOTE: The potential function gates the creation of timeline indexes!
// This can adjust the "resolution" of output data =)
//
// Muted indicates if the stimulator of this dimension should be created muted.
func Observer[TValue any](engine *core.Engine, potential core.Potential, muted bool, target std.TargetFunc[TValue]) *Dimension[TValue, any] {
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
		d.Timeline = append(d.Timeline, data)
		d.Current = &data
		d.Mutex.Unlock()
	}
	d.Stimulator = engine.Stimulate(f, potential, muted)
	return &d
}
