package std

import "time"

type Activation struct {
	id uint64

	// Creation represents the moment the synaptic connection was created.
	Creation time.Time

	// Inception represents the moment a neuron received a synaptic event.
	Inception time.Time

	// Activation represents the moment a neuron passed it's potential.
	Activation *time.Time

	// Completion represents the moment a neuron finished execution.
	Completion *time.Time

	// Step holds the activation step and infinitely increments.
	Step uint

	// Beat holds the frequency beat and is bound to the cortical frequency.
	Beat uint

	// Cycle holds the current periodic cycle number and infinitely increments.
	Cycle uint

	// Period holds the amount of time a single Frequency cycle should take.
	Period time.Duration

	// Frequency is the number of beats to pulse in a given Period.
	Frequency float64
}

func (a *Activation) ID() uint64 {
	return a.id
}
