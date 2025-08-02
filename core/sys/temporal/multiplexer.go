package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
)

// Multiplexer creates a dimension that's a blend of the point value of many input dimensions for every impulse that the potential returns true.
//
// NOTE: The potential function gates the creation of timeline indexes!
// This can adjust the "resolution" of output data =)
//
// Muted indicates if the stimulator of this dimension should be created muted.
func Multiplexer[TValue core.Numeric](engine *core.Engine, potential core.Potential, muted bool, blend Blend[TValue], dimensions ...*Dimension[any, any]) *Dimension[TValue, any] {
	d := Dimension[TValue, any]{}
	d.NamedEntity = core.NewNamedEntity()
	d.Window = core.DefaultObservanceWindow
	d.Trimmer = engine.Loop(d.ImpulseTrim, when.Frequency(&core.TrimFrequency), false)
	f := func(ctx core.Context) {
		values := make([]any, len(dimensions))
		for i, otherD := range dimensions {
			values[i] = otherD.Current
		}
		data := std.Data[TValue]{
			Context: ctx,
			Point:   blend(values),
		}
		d.Mutex.Lock()
		d.Timeline = append(d.Timeline, data)
		d.Current = &data
		d.Mutex.Unlock()
	}
	d.Stimulator = engine.Stimulate(f, potential, muted)
	return &d
}
