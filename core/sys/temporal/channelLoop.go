package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
)

// ChannelLoop doesn't trigger anything, rather it sends the context through a channel assigned to the dimension's Cache.
//
// You must read these messages and handle them for the activation to complete.
func ChannelLoop(engine *core.Engine, potential core.Potential, muted bool) *Dimension[core.Runtime, chan std.ChannelAction] {
	d := Dimension[core.Runtime, chan std.ChannelAction]{}
	d.NamedEntity = core.NewNamedEntity()
	d.Window = core.DefaultObservanceWindow
	d.Trimmer = engine.Loop(d.ImpulseTrim, when.Frequency(&core.TrimFrequency), false)
	c := make(chan std.ChannelAction)
	d.Cache = &c
	f := func(ctx core.Context) {
		data := std.Data[core.Runtime]{
			Context: ctx,
			Point:   d.Stimulator.LastActivation,
		}
		*d.Cache <- std.ChannelAction{Context: ctx}
		d.Mutex.Lock()
		d.Timeline = append(d.Timeline, data)
		d.Current = &data
		d.Mutex.Unlock()
	}

	d.Stimulator = engine.Loop(f, potential, muted)
	return &d
}
