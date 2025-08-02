package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
)

// ChangeReaction creates a dimension that calls the provided reaction function if the comparator
// finds that the value has changed since the last impulse - if the provided potential returns true.
//
// Muted indicates if the stimulator of this dimension should be created muted.
func ChangeReaction[TValue any](engine *core.Engine, potential core.Potential, muted bool, target std.TargetFunc[TValue], comparator Comparator[TValue], reaction Change[TValue]) *Dimension[TValue, any] {
	d := Dimension[TValue, any]{}
	d.NamedEntity = core.NewNamedEntity()
	d.Window = core.DefaultObservanceWindow
	d.Trimmer = engine.Loop(d.ImpulseTrim, when.Frequency(&core.TrimFrequency), false)
	f := func(ctx core.Context) {
		data := std.Data[TValue]{
			Context: ctx,
			Point:   *target(),
		}

		if d.Current == nil || !comparator(d.Current.Point, data.Point) {
			d.Mutex.Lock()
			var old *std.Data[TValue]
			if d.Current != nil {
				old = d.Current
			}
			d.Timeline = append(d.Timeline, data)
			d.Current = &data
			reaction(ctx, *old, *d.Current)
			d.Mutex.Unlock()
		}
	}
	d.Stimulator = engine.Stimulate(f, potential, muted)
	return &d
}
