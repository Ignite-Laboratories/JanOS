package core

// A Neuron is any type that can fire an action-potential (see.ActionPotentials).  The Cleanup function
// will be called whenever the neuron's Lifecycle has reached completion or the cortex shuts down.
type Neuron interface {
	Action(Context) bool
	Potential(Context) bool
	Cleanup()
}
