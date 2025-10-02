package std

import (
	"strings"
	"time"
)

type Impulse struct {
	Activation

	Name       string
	Cortex     Cortex
	cortexName string
	Timeline   *TemporalBuffer[Activation]

	// Decay can be used to flag back to the cortex to decay this synapse.
	Decay bool

	Thought *Thought

	recycle func(bool)
}

func (imp Impulse) Bridge() []string {
	return []string{imp.cortexName, imp.Name}
}

func (imp Impulse) String() string {
	return strings.Join(imp.Bridge(), " ⇝ ")
}

// RefractoryPeriod represents the duration between the last activation's completion and this impulse's activation.
func (imp *Impulse) RefractoryPeriod() time.Duration {
	last := imp.Timeline.Latest()
	if len(last) <= 0 {
		return 0
	}
	return imp.Activation.Activation.Sub(*last[0].Element.Completion)
}

// CyclePeriod represents the duration between the last activation's start and this impulse's activation start.
func (imp *Impulse) CyclePeriod() time.Duration {
	last := imp.Timeline.Latest()
	if len(last) <= 0 {
		return 0
	}
	return imp.Activation.Activation.Sub(*last[0].Element.Activation)
}

// ResponseTime represents the duration between inception and activation of the current impulse's activation.
func (imp *Impulse) ResponseTime() time.Duration {
	return imp.Activation.Activation.Sub(imp.Activation.Inception)
}

// RunTime represents the duration between activation and completion of the last impulse's activation.
func (imp *Impulse) RunTime() time.Duration {
	last := imp.Timeline.Latest()
	if len(last) <= 0 {
		return 0
	}
	return last[0].Element.Completion.Sub(*last[0].Element.Activation)
}

// TotalTime represents the duration between inception and completion of the last impulse's activation.
func (imp *Impulse) TotalTime() time.Duration {
	last := imp.Timeline.Latest()
	if len(last) <= 0 {
		return 0
	}
	return last[0].Element.Completion.Sub(last[0].Element.Inception)
}
