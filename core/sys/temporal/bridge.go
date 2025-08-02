package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
	"time"
)

// Bridge creates a dimension that can be impulsed directly.  Because bridges are not driven by
// an impulse engine, they don't have a potential mechanic.  An engine is provided because the
// observational trimmer still requires one to be driven.
//
// The returned function acts as a means to impulse the dimension with, at your pace.
//
// While the recorded context is not tied to an engine, a bridge still tracks its own 'beat' value.
// This value will never loop over in a bridged dimension and will always count up from zero at the
// moment of inception.  Since activation or impulse statistics are not calculable from -inside-
// a dimension, they are not provided; however, the periodicity between activations is.
func Bridge[TValue any](engine *core.Engine) (func(TValue), *Dimension[TValue, any]) {
	d := Dimension[TValue, any]{}
	d.NamedEntity = core.NewNamedEntity()
	d.Window = core.DefaultObservanceWindow
	d.Trimmer = engine.Loop(d.ImpulseTrim, when.Frequency(&core.TrimFrequency), false)

	var beat int
	var lastMoment time.Time

	callback := func(value TValue) {
		now := time.Now()

		ctx := core.Context{
			Beat:   beat,
			Moment: now,
			Period: now.Sub(lastMoment),
		}

		data := std.Data[TValue]{
			Context: ctx,
			Point:   value,
		}
		d.Mutex.Lock()
		d.Timeline = append(d.Timeline, data)
		d.Current = &data
		d.Mutex.Unlock()

		beat++
		lastMoment = now
	}

	return callback, &d
}
