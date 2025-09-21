package std

import (
	"sync"
)

type Neuron struct {
	Entity

	muted     bool
	action    func(*Impulse)
	potential func(*Impulse) bool
	cleanup   func(*Impulse, *sync.WaitGroup)
}

func NewLongRunning(named string, action func(*Impulse), potential func(*Impulse) bool, cleanup ...func(*Impulse, *sync.WaitGroup)) Neural {
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

func (lr Neuron) Mute() {
	lr.muted = true
}

func (lr Neuron) Unmute() {
	lr.muted = false
}

func (lr Neuron) Action(imp *Impulse) {
	lr.action(imp)
}

func (lr Neuron) Potential(imp *Impulse) bool {
	if lr.muted {
		return false
	}
	return lr.potential(imp)
}

func (lr Neuron) Cleanup(imp *Impulse, wg *sync.WaitGroup) {
	if lr.cleanup != nil {
		lr.cleanup(imp, wg)
	}
}
