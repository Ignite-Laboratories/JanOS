package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"runtime"
)

// DedicatedLoop is a special kind of ChannelLoop - it guarantees the target action is always
// called from the same host thread using runtime.LockOSThread()
func DedicatedLoop(engine *core.Engine, potential core.Potential, muted bool, target core.Action) *Dimension[core.Runtime, chan std.ChannelAction] {
	d := ChannelLoop(engine, potential, muted)
	go func() {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		for msg := range *d.Cache {
			target(msg.Context)
		}
	}()
	return d
}
