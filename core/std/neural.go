package std

import "sync"

// A Neural is any type that can fire an action-potential (see.ActionPotentials).  The Cleanup function
// will be called whenever the neuron's Lifecycle has reached completion or the cortex shuts down.
type Neural interface {
	Named(...string) string

	Action(*Impulse)
	Potential(*Impulse) bool

	Mute()
	Unmute()

	// Cleanup will be called once the synaptic lifecycle completes.  This will fire regardless of if
	// the underlying action fires.
	Cleanup(*Impulse, *sync.WaitGroup)
}
