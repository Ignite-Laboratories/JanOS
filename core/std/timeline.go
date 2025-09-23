package std

import (
	"time"

	"git.ignitelabs.net/core"
)

type SynapticEvent struct {
	// SynapseCreation represents the moment the synaptic connection was created.
	SynapseCreation time.Time

	// Inception represents the moment a neuron received a synaptic event.
	Inception time.Time

	// Activation represents the moment a neuron passed it's potential.
	Activation time.Time

	// Completion represents the moment a neuron finished execution.
	Completion time.Time
}

// Timeline represents key moments in the lifecycle of a SynapticEvent.  A synaptic event is the contextual
// activation of a Neuron from a Cortex - typically traversed along an axon in biological structure.
type Timeline struct {
	temporal TemporalBuffer[SynapticEvent]
}

func NewTimeline() *Timeline {
	inception := SynapticEvent{
		SynapseCreation: core.Inception,
		Inception:       core.Inception,
		Activation:      core.Inception,
		Completion:      core.Inception,
	}
	timeline := &Timeline{
		temporal: NewTemporalBuffer[SynapticEvent](),
	}
	timeline.Add(inception)
	return timeline
}

func (t *Timeline) Len() uint {
	return t.temporal.Len()
}

func (t *Timeline) Add(element SynapticEvent) {
	t.temporal.Add(element.Inception, element)
}

func (t *Timeline) Yield() []SynapticEvent {
	yield := t.temporal.Yield()
	out := make([]SynapticEvent, len(yield))
	for i, y := range yield {
		out[i] = y.Element
	}
	return out
}

// RefractoryPeriod represents the duration between the last activation's completion and this impulse's activation.
func (t *Timeline) RefractoryPeriod() time.Duration {
	latest := t.temporal.Latest(2)
	if len(latest) <= 1 {
		return 0
	}
	return latest[1].Element.Activation.Sub(latest[0].Element.Completion)
}

// CyclePeriod represents the duration between the last activation's start and this impulse's activation start.
func (t *Timeline) CyclePeriod() time.Duration {
	latest := t.temporal.Latest(2)
	if len(latest) <= 1 {
		return 0
	}
	return latest[1].Element.Activation.Sub(latest[0].Element.Activation)
}

// ResponseTime represents the duration between inception and activation.
func (t *Timeline) ResponseTime() time.Duration {
	latest := t.temporal.Latest()
	if len(latest) <= 1 {
		return 0
	}
	return latest[1].Element.Activation.Sub(latest[1].Element.Inception)
}

// RunTime represents the duration between activation and completion.
func (t *Timeline) RunTime() time.Duration {
	latest := t.temporal.Latest()
	if len(latest) <= 1 {
		return 0
	}
	return latest[1].Element.Completion.Sub(latest[0].Element.Activation)
}

// TotalTime represents the duration between inception and completion.
func (t *Timeline) TotalTime() time.Duration {
	latest := t.temporal.Latest()
	if len(latest) <= 1 {
		return 0
	}
	return latest[1].Element.Completion.Sub(latest[0].Element.Inception)
}
