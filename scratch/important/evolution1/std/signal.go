package std

import (
	"sync"

	"git.ignitelabs.net/janos/core/enum/life"
)

type Neural interface {
	Action(*Impulse)
	Potential(*Impulse) bool
	Cleanup(*Impulse)
}

type signalMaker byte
type signal interface{}

type decay struct {
	wait *sync.WaitGroup
}

type mute byte

type unmute byte

type spark []synapse

var Signal signalMaker

func (signalMaker) Decay(wg *sync.WaitGroup) decay {
	return decay{wg}
}

func (signalMaker) Mute() mute {
	return mute(0)
}

func (signalMaker) Unmute() unmute {
	return unmute(0)
}

func (signalMaker) Spark(synapses ...synapse) spark {
	return synapses
}

func (signalMaker) Synapse(lifecycle life.Cycle, named string, action func(*Impulse), potential func(*Impulse) bool, cleanup ...func(*Impulse)) synapse {
	return newSynapse(lifecycle, named, action, potential, cleanup...)
}
