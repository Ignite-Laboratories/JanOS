package neuron

import (
	"sync"

	"git.ignitelabs.net/core/std"
)

type LongRunning struct {
	std.Entity

	running bool
	mutex   *sync.Mutex

	action    func(*std.Impulse)
	potential func(*std.Impulse) bool
	cleanup   func(*std.Impulse, *sync.WaitGroup)
}

func NewLongRunning(named string, action func(*std.Impulse), potential func(*std.Impulse) bool, cleanup ...func(*std.Impulse, *sync.WaitGroup)) std.Neuron {
	if action == nil {
		panic("the action of a neuron can never be nil")
	}
	if potential == nil {
		potential = func(*std.Impulse) bool {
			return true
		}
	}

	var clean func(*std.Impulse, *sync.WaitGroup)
	if len(cleanup) > 0 {
		clean = cleanup[0]
	}
	return &LongRunning{
		Entity:    std.NewEntityNamed(named),
		mutex:     &sync.Mutex{},
		action:    action,
		potential: potential,
		cleanup:   clean,
	}
}

func (lr LongRunning) Mute() {
	lr.mutex.Lock()
}

func (lr LongRunning) Unmute() {
	lr.mutex.Unlock()
}

func (lr LongRunning) Action(imp *std.Impulse) {
	if lr.mutex == nil {
		lr.mutex = &sync.Mutex{}
	}

	lr.mutex.Lock()
	defer lr.mutex.Unlock()

	lr.running = true
	lr.action(imp)
	lr.running = false
}

func (lr LongRunning) Potential(imp *std.Impulse) bool {
	if lr.running {
		return false
	}
	return lr.potential(imp)
}

func (lr LongRunning) Cleanup(imp *std.Impulse, wg *sync.WaitGroup) {
	if lr.cleanup != nil {
		lr.cleanup(imp, wg)
	}
}
