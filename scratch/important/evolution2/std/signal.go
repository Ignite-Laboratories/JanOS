package std

import (
	"sync"
)

type signalMaker byte
type signal interface{}

var Signal signalMaker

type decay struct {
	wait *sync.WaitGroup
}

func (signalMaker) Decay(wg *sync.WaitGroup) decay {
	return decay{wg}
}

type mute byte

func (signalMaker) Mute() mute {
	return mute(0)
}

type unmute byte

func (signalMaker) Unmute() unmute {
	return unmute(0)
}

type spark []synapse

// A Spark (when sent to a cortex) starts neural activations.  When sent to a cluster, please see its documentation for what will happen.
func (signalMaker) Spark(synapses ...synapse) spark {
	return synapses
}
