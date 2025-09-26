package std

import (
	"sync"
)

// Neuron represents a primitive implementation of the Neural interface.  This can be used for creating anonymous neural activity.
type Neuron struct {
	Entity

	action    func(*Impulse)
	potential func(*Impulse) bool
	cleanup   func(*Impulse, *sync.WaitGroup)
}

func NewNeuron(named string, action func(*Impulse), potential func(*Impulse) bool, cleanup ...func(*Impulse, *sync.WaitGroup)) Neural {
	if action == nil {
		panic("the action of a neuron can never be nil")
	}
	if potential == nil {
		potential = func(*Impulse) bool {
			return true
		}
	}

	var clean func(*Impulse, *sync.WaitGroup)
	if len(cleanup) > 0 {
		clean = cleanup[0]
	}
	return &Neuron{
		Entity:    NewEntityNamed(named),
		action:    action,
		potential: potential,
		cleanup:   clean,
	}
}

func (n Neuron) Action(imp *Impulse) {
	n.action(imp)
}

func (n Neuron) Potential(imp *Impulse) bool {
	return n.potential(imp)
}

func (n Neuron) Cleanup(imp *Impulse, wg *sync.WaitGroup) {
	if n.cleanup != nil {
		n.cleanup(imp, wg)
	}
}
