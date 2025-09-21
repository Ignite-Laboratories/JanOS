package lifecycle

// A Lifecycle defines how the neuron should be re-activated.  When a neuron is 'activated' it's provided a
// -new- goroutine.  Looping neuron recycle the same goroutine, while all others launch a new goroutine on
// every activation - so please be wary of that.  There are four ways a neuron can be activated:
//
// 0 - Looping - this will cyclically re-activate the neuron only after it finishes its current execution
//
// 1 - Stimulative - this will activate the neuron whenever its potential returns high, regardless of current execution
//
// 2 - Triggered - this will perform a single impulse once the potential goes high.
//
// 3 - Impulse - this will ATTEMPT to activate the neuron -exactly- once.
type Lifecycle byte

const (
	Looping Lifecycle = iota
	Stimulative
	Triggered
	Impulse
)
