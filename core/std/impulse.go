package std

// An Impulse represents the act of a synaptic event between a cortex and neuron.  This contains several key points:
//
//   - Bridge is the named synaptic bridge between the cortex and neuron (for tracing purposes)
//   - Timeline holds the temporal activation information for this neuron
//   - Count indicates the number of activations that this Synapse has processed
//   - Cortex provides a reference to the cortex that generated this impulse (the neuron can be impulsed by many cortices)
//   - Neural provides a reference to the neuron that this impulse terminates into.
//   - Thought holds a reference to the data this synaptic bridge is maturing over time.
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
