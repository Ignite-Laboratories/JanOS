package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
)

// Operation represents the state of A and B that generated the resulting Value.
type Operation[TValue core.Numeric] struct {
	Value TValue
	A     std.Data[TValue]
	B     std.Data[TValue]
}

func Operator[TValue core.Numeric](engine *core.Engine, potential core.Potential, muted bool, operator Operate[TValue], a *Dimension[TValue, any], b *Dimension[TValue, any]) *Dimension[Operation[TValue], any] {
	d := Dimension[Operation[TValue], any]{}
	d.NamedEntity = core.NewNamedEntity()
	d.Window = core.DefaultObservanceWindow
	d.Trimmer = engine.Loop(d.ImpulseTrim, when.Frequency(&core.TrimFrequency), false)
	f := func(ctx core.Context) {
		operation := Operation[TValue]{
			A: *a.Current,
			B: *b.Current,
		}
		operation.Value = operator(operation.A.Point, operation.B.Point)
		data := std.Data[Operation[TValue]]{
			Context: ctx,
			Point:   operation,
		}
		d.Mutex.Lock()
		d.Timeline = append(d.Timeline, data)
		d.Current = &data
		d.Mutex.Unlock()
	}
	d.Stimulator = engine.Stimulate(f, potential, muted)
	return &d
}
