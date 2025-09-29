package std

import "strings"

// An Impulse represents the act of a synaptic event between a cortex and neuron.  This contains several key points:
//
// Bridge - This is the named synaptic bridge between the cortex and neuron (for tracing purposes)
//
// # TimelineOld - This holds the temporal activation information for this neuron
//
// # Count - This indicates the number of activations that this Synapse has processed
//
// Cortex - This provides a reference to the cortex that generated this impulse (the neuron can be impulsed by many cortices)
//
// Neural - This provides a reference to the neuron that this impulse terminates into.
//
// Thought - This holds the actual impulse data that -this- synaptic bridge has been accumulating across time.  Only the neuron
// understands how to interpret the information contained in the impulsive thought.
type Impulse struct {
	Bridge   Bridge
	Timeline *Timeline
	Beat     uint
	Phase    int
	Count    uint
	Cortex   *Cortex
	Neuron   Neural
	Decay    bool
	Mute     bool

	Thought *Thought

	currentEvent SynapticEvent
}

type Bridge []string

func (b Bridge) String() string {
	return strings.Join(b, " ⇝ ")
}
