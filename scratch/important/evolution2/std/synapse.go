package std

import "git.ignitelabs.net/janos/core/enum/life"

type neuron chan<- signal

func (signalMaker) Cluster(named string, synapses ...synapse) neuron {
	cls := make(chan signal)
	for _, syn := range synapses {
		cls <- syn
	}
	return neuron(cls)
}

type synapse struct {
	action    func(*Impulse)
	potential func(*Impulse) bool
	cleanup   func(*Impulse)
}

func Synapse(lifecycle life.Cycle, named string, action func(*Impulse), potential func(*Impulse) bool, cleanup ...func(*Impulse)) synapse {

}

func (syn synapse) Action(imp *Impulse) {
	syn.action(imp)
}

func (syn synapse) Potential(imp *Impulse) bool {
	return syn.potential(imp)
}

func (syn synapse) Cleanup(imp *Impulse) {
	if syn.cleanup != nil {
		syn.cleanup(imp)
	}
}


Types of clustered activation:

	- Burst, where all synapses are impulsed on every impulse
	- Round, where synapses are impulsed in a round on every impulse
	- Distributed, where synapses are impulsed distributed across a period

The cortex drives the cluster, but the cluster decides how to activate the synapses