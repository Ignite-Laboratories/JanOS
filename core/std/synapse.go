package std

import (
	"sync"
	"time"

	"git.ignitelabs.net/core"
	"git.ignitelabs.net/core/enum/lifecycle"
	"git.ignitelabs.net/core/sys/log"
)

// A Synapse represents a fixed impulsive activation between a Neuron and Cortex.  A synapse can be used to recycle
// the same action across many cortices, as it can be sparked as many times as you would like.
type Synapse func(*Impulse)

// NewSynapse creates a Neural connection between a Neuron and a Cortex.  You may optionally provide 'nil' to the potential if you'd like to imply 'always fire'.
func NewSynapse(lifecycle lifecycle.Lifecycle, neuronName string, action func(*Impulse) bool, potential func(*Impulse) bool, cleanup ...func(*Impulse, *sync.WaitGroup)) Synapse {
	n := NewNeuron(neuronName, action, potential, cleanup...)
	return NewSynapseFromNeuron(lifecycle, n)
}

// NewSynapseFromNeuron creates a neural Synapse which fires the provided action potential according to the provided lifecycle.Lifecycle.  You may optionally
// provide a cleanup function which is called after this Synapse finishes all neural activation (or the cortex shuts down).  For triggered
// or impulsed lifecycles, this gets called immediately - for looping or stimulative, this gets called after the cortex shuts down.
func NewSynapseFromNeuron(life lifecycle.Lifecycle, neuron Neural) Synapse {
	log.Verbosef(core.ModuleName, "creating synapse '%s'\n", neuron.Named())
	beat := 0
	return func(imp *Impulse) {
		creation := time.Now()
		imp.Bridge = (*imp.Cortex).Named() + " ⇝ " + neuron.Named()
		imp.Neuron = neuron

		log.Verbosef((*imp.Cortex).Named(), "wired axon to synapse '%s'\n", neuron.Named())

		panicSafeAction := func(i *Impulse) bool {
			defer func() {
				if r := recover(); r != nil {
					log.Printf(i.Bridge, "neural panic: %s\n", r)
				}
			}()

			return neuron.Action(i)
		}

		switch life {
		case lifecycle.Looping:
			// 0 - Looping activations cyclically reactivate the same goroutine when the last finishes and the potential is high
			go func() {
				log.Verbosef(imp.Bridge, "looping\n")
				for (*imp.Cortex).Alive() {
					event := SynapticEvent{
						SynapseCreation: creation,
						Inception:       time.Now(),
					}
					imp.Beat = beat
					var keepAlive bool
					if neuron.Potential(imp) {
						event.Activation = time.Now()
						imp.Event = event
						keepAlive = panicSafeAction(imp)
						event.Completion = time.Now()
						beat++
					}
					imp.Timeline.Add(event)

					if keepAlive {
						(*imp.Cortex).master.Lock()
						(*imp.Cortex).clock.Wait()
						(*imp.Cortex).master.Unlock()
					} else {
						break
					}
				}

				wg := &sync.WaitGroup{}
				wg.Add(1)
				neuron.Cleanup(imp, wg)
				log.Verbosef(imp.Bridge, "ended\n")
			}()
		case lifecycle.Stimulative:
			// 1 - Stimulative activations launch new goroutines on every impulse the potential is high
			go func() {
				log.Verbosef(imp.Bridge, "stimulating\n")
				for (*imp.Cortex).Alive() {
					event := SynapticEvent{
						SynapseCreation: creation,
						Inception:       time.Now(),
					}
					imp.Beat = beat
					var keepAlive bool
					if neuron.Potential(imp) {
						event.Activation = time.Now()
						imp.Event = event
						panicSafeAction(imp)
						event.Completion = time.Now()
						beat++
					}
					imp.Timeline.Add(event)

					if keepAlive {
						(*imp.Cortex).master.Lock()
						(*imp.Cortex).clock.Wait()
						(*imp.Cortex).master.Unlock()
					} else {
						break
					}
				}

				wg := &sync.WaitGroup{}
				wg.Add(1)
				neuron.Cleanup(imp, wg)
				log.Verbosef(imp.Bridge, "ended\n")
			}()
		case lifecycle.Triggered:
			// 2 - Triggered activations are a one-shot GUARANTEE once the potential goes high
			go func() {
				log.Verbosef(imp.Bridge, "triggering\n")
				event := SynapticEvent{
					SynapseCreation: creation,
					Inception:       time.Now(),
				}
				for (*imp.Cortex).Alive() && !neuron.Potential(imp) {
					(*imp.Cortex).master.Lock()
					(*imp.Cortex).clock.Wait()
					(*imp.Cortex).master.Unlock()
				}
				if (*imp.Cortex).Alive() {
					imp.Beat = beat
					event.Activation = time.Now()
					imp.Event = event
					panicSafeAction(imp)
					event.Completion = time.Now()
					beat++
					imp.Timeline.Add(event)
				}

				wg := &sync.WaitGroup{}
				wg.Add(1)
				neuron.Cleanup(imp, wg)
				log.Verbosef(imp.Bridge, "ended\n")
			}()
		case lifecycle.Impulse:
			// 3 - Impulse activations are a one-shot ATTEMPT regardless of potential
			go func() {
				log.Verbosef(imp.Bridge, "impulsing\n")
				event := SynapticEvent{
					SynapseCreation: creation,
					Inception:       time.Now(),
				}
				if (*imp.Cortex).Alive() && neuron.Potential(imp) {
					imp.Beat = beat
					event.Activation = time.Now()
					imp.Event = event
					panicSafeAction(imp)
					event.Completion = time.Now()
					beat++
					imp.Timeline.Add(event)
				}

				wg := &sync.WaitGroup{}
				wg.Add(1)
				neuron.Cleanup(imp, wg)
				log.Verbosef(imp.Bridge, "ended\n")
			}()
		}
	}
}
