package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
)

// Blending represents the state of A and B that generated the resulting blended Value.
type Blending[TValue core.Numeric] struct {
	Value TValue
	A     std.Data[TValue]
	B     std.Data[TValue]
}

// Blender creates a dimension that blends the point value of two input dimensions for every impulse that the potential returns true.
//
// NOTE: The potential function gates the creation of timeline indexes!
// This can adjust the "resolution" of output data =)
//
// Muted indicates if the stimulator of this dimension should be created muted.
func Blender[TValue core.Numeric](engine *core.Engine, potential core.Potential, muted bool, blend Operate[TValue], a *Dimension[TValue, any], b *Dimension[TValue, any]) *Dimension[Blending[TValue], any] {
	d := Dimension[Blending[TValue], any]{}
	d.NamedEntity = core.NewNamedEntity()
	d.Window = core.DefaultObservanceWindow
	d.Trimmer = engine.Loop(d.ImpulseTrim, when.Frequency(&core.TrimFrequency), false)
	f := func(ctx core.Context) {
		mux := Blending[TValue]{
			A: *a.Current,
			B: *b.Current,
		}
		mux.Value = blend(mux.A.Point, mux.B.Point)
		data := std.Data[Blending[TValue]]{
			Context: ctx,
			Point:   mux,
		}
		d.Mutex.Lock()
		d.Timeline = append(d.Timeline, data)
		d.Current = &data
		d.Mutex.Unlock()
	}
	d.Stimulator = engine.Stimulate(f, potential, muted)
	return &d
}
