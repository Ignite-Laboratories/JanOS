package std

import (
	"time"

	"git.ignitelabs.net/core"
	"git.ignitelabs.net/core/sys/atlas"
)

// Timeline represents key moments in the lifecycle of a neuron.
type Timeline struct {
	// Last represents the last activation's timeline.
	//
	// NOTE: This does not recurse beyond the last activation's timeline.
	Last *Timeline

	// SynapticCreation represents the moment the synaptic connection was created.
	SynapticCreation time.Time

	// Inception represents the moment a neuron received a synaptic event.
	Inception time.Time

	// Activation represents the moment a neuron passed it's potential.
	Activation time.Time

	// Completion represents the moment a neuron finished execution.
	Completion time.Time
}

func newTimeline() Timeline {
	inception := Timeline{
		Last:             nil,
		SynapticCreation: core.Inception,
		Inception:        core.Inception,
		Activation:       core.Inception,
		Completion:       core.Inception,
	}
	inception.Last = &inception
	return inception
}

func timelineRollover(src *Timeline) {
	if src == nil {
		return
	}

	if atlas.TimelineDepth == 0 {
		src.Last = nil
		return
	}

	copied := new(Timeline)
	*copied = *src
	src.Last = copied

	t := src
	for i := uint(0); i < atlas.TimelineDepth; i++ {
		if t.Last == nil {
			return
		}
		if i == atlas.TimelineDepth-1 {
			t.Last = nil
			return
		}
		t = t.Last
	}
}

// RefractoryPeriod represents the duration between the last activation's completion and this impulse's activation.
func (t Timeline) RefractoryPeriod() time.Duration {
	if t.Last == nil {
		return 0
	}
	return t.Activation.Sub((*t.Last).Completion)
}

// CyclePeriod represents the duration between the last activation's start and this impulse's activation.
func (t Timeline) CyclePeriod() time.Duration {
	if t.Last == nil {
		return 0
	}
	return t.Activation.Sub((*t.Last).Activation)
}

// ResponseTime represents the duration between inception and activation.
func (t Timeline) ResponseTime() time.Duration {
	return t.Activation.Sub(t.Inception)
}

// RunTime represents the duration between activation and completion.
func (t Timeline) RunTime() time.Duration {
	return t.Completion.Sub(t.Activation)
}

// TotalTime represents the duration between inception and completion.
func (t Timeline) TotalTime() time.Duration {
	return t.Completion.Sub(t.Inception)
}
