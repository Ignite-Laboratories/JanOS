package std

// Neuron is a logical unit of execution.
type Neuron struct {
	Entity

	// Muted can be used to explicitly suppress neural activation.
	Muted bool

	// Action is what the engine calls whenever the Potential returns true.
	Action Action

	// Destroyed indicates if this neuron has been destroyed and can be used to make cleanup decisions.
	//
	// To destroy a neuron, please use its Destroy() method.
	Destroyed ReadOnlyBool

	// Potential must return true when called for activation to occur.
	Potential Potential

	// LastActivation provides temporal runtime information for the last activation.
	LastActivation Runtime

	// ActivationCount provides the number of times this neuron has been activated.
	ActivationCount uint64

	destroyed bool
	executing bool
	engine    *Engine
}

// Trigger fires the neural action once, if the potential returns true.
//
// If 'async' is true, the action is called asynchronously - otherwise, it blocks the firing impulse.
func (n *Neuron) Trigger(async bool) {
	n.engine.Trigger(n.Action, n.Potential, async)
}

// Destroy removes this neuron from the engine entirely.
func (n *Neuron) Destroy() {
	n.destroyed = true
	n.engine.remove(n.ID)
}
