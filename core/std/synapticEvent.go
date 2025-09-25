package std

import "time"

type SynapticEvent struct {
	id uint64

	// SynapseCreation represents the moment the synaptic connection was created.
	SynapseCreation time.Time

	// Inception represents the moment a neuron received a synaptic event.
	Inception time.Time

	// Activation represents the moment a neuron passed it's potential.
	Activation time.Time

	// Completion represents the moment a neuron finished execution.
	Completion time.Time
}
