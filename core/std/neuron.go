package std

// A Neuron is any type that can fire an action-potential (see.ActionPotentials).  The Cleanup function
// will be called whenever the neuron's Lifecycle has reached completion or the cortex shuts down.
type Neuron interface {
	Named(...string) string

	Action(Impulse)
	Potential(Impulse) bool

	Mute()
	Unmute()

	// Cleanup will be called once the synaptic lifecycle completes.  This will -always- fire, regardless if
	// the underlying action fires.
	Cleanup()
}
