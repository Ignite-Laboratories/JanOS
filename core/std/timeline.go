package std

import (
	"time"

	"git.ignitelabs.net/janos/core"
)

// Timeline represents key moments in the lifecycle of a SynapticEvent.  A synaptic event is the contextual
// activation of a Neuron from a Cortex - typically traversed along an axon in biological structure.
type Timeline struct {
	temporal *TemporalBuffer[SynapticEvent]
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

func (t *Timeline) setCompleted(id uint64, moment time.Time) {
	t.temporal.master.Lock()
	defer t.temporal.master.Unlock()

	for i := len(t.temporal.buffer) - 1; i >= 0; i-- {
		if t.temporal.buffer[i].Element.id == id {
			t.temporal.buffer[i].Element.Completion = moment
			break
		}
	}
}

func (t *Timeline) Len() uint {
	return t.temporal.Len()
}

func (t *Timeline) Add(element SynapticEvent) {
	t.temporal.Add(element.Inception, element)
}

// Latest returns the latest elements from the timeline in temporal order up to the provided depth.
//
// NOTE: If no depth is provided, a depth of '1' is implied. If a negative depth is provided, all elements are returned.
func (t *Timeline) Latest(depth ...int) []SynapticEvent {
	latest := t.temporal.Latest(depth...)
	out := make([]SynapticEvent, len(latest))
	for i, l := range latest {
		out[i] = l.Element
	}
	return out
}

// LatestSince returns the latest elements from the timeline in temporal order up to the provided moment in the past.
//
// NOTE: If you provide a moment in the future, nothing will match and an empty set will be returned.
func (t *Timeline) LatestSince(moment time.Time) []SynapticEvent {
	latest := t.temporal.LatestSince(moment)
	out := make([]SynapticEvent, len(latest))
	for i, l := range latest {
		out[i] = l.Element
	}
	return out
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
	if len(latest) == 0 {
		return 0
	}
	if len(latest) == 1 {
		return latest[0].Element.Activation.Sub(latest[0].Element.SynapseCreation)
	}
	return latest[1].Element.Activation.Sub(latest[0].Element.Activation)
}

// ResponseTime represents the duration between inception and activation.
func (t *Timeline) ResponseTime() time.Duration {
	latest := t.temporal.Latest()
	if len(latest) == 0 {
		return 0
	}
	return latest[0].Element.Activation.Sub(latest[0].Element.Inception)
}

// RunTime represents the duration between activation and completion.
func (t *Timeline) RunTime() time.Duration {
	latest := t.temporal.Latest()
	if len(latest) <= 1 {
		return 0
	}
	return latest[0].Element.Completion.Sub(latest[0].Element.Activation)
}

// TotalTime represents the duration between inception and completion.
func (t *Timeline) TotalTime() time.Duration {
	latest := t.temporal.Latest()
	if len(latest) <= 1 {
		return 0
	}
	return latest[0].Element.Completion.Sub(latest[0].Element.Inception)
}
