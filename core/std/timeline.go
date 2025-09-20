package std

import "time"

// Timeline represents key moments in the lifecycle of a neuron.
type Timeline struct {
	// Last represents the last activation's timeline.
	//
	// NOTE: This does not recurse beyond the last activation's timeline.
	Last *Timeline

	// Creation represents the moment the synaptic connection was created.
	Creation time.Time

	// Inception represents the moment a neuron received a synaptic event.
	Inception time.Time

	// Activation represents the moment a neuron passed it's potential.
	Activation time.Time

	// Completion represents the moment a neuron finished execution.
	Completion time.Time
}

// RefractoryPeriod represents the duration between the last activation's completion and this impulse's activation.
func (r Timeline) RefractoryPeriod() time.Duration {
	if r.Last == nil {
		return 0
	}
	return r.Activation.Sub((*r.Last).Completion)
}

// ResponseTime represents the duration between inception and activation.
func (r Timeline) ResponseTime() time.Duration {
	return r.Activation.Sub(r.Inception)
}

// RunTime represents the duration between activation and completion.
func (r Timeline) RunTime() time.Duration {
	return r.Completion.Sub(r.Activation)
}

// TotalTime represents the duration between inception and completion.
func (r Timeline) TotalTime() time.Duration {
	return r.Completion.Sub(r.Inception)
}
