package std

import (
	"sync"
	"time"

	"git.ignitelabs.net/core/enum/lifecycle"
	"git.ignitelabs.net/core/sys/log"
)

type Synapse func(*Impulse)

// NewSynapse creates a neural Synapse which fires the provided action potential according to the provided Lifecycle.  You may optionally
// provide a cleanup function which is called after this Synapse finishes all neural activation (or the cortex shuts down).  For triggered
// or impulsed lifecycles, this happens immediately - for looping or stimulative, this happens after the cortex shuts down or when the
// provided action returns false.
//
// NOTE: For stimulative activations, the action may still fire a few times after returning false - but cleanup will happen after all
//
//	activations are complete.
func NewSynapse(life lifecycle.Lifecycle, neuron Neuron) Synapse {
	beat := 0
	return func(imp *Impulse) {
		imp.Timeline.Creation = time.Now()
		imp.Bridge = (*imp.Cortex).Named() + " ↦ " + neuron.Named()
		imp.Neuron = neuron

		log.Printf((*imp.Cortex).Named(), "created synapse to '%s'\n", neuron.Named())

		panicSafeAction := func(i *Impulse) {
			defer func() {
				if r := recover(); r != nil {
					log.Printf(i.Bridge, "neural panic: %s\n", r)
				}
			}()

			neuron.Action(i)
		}

		switch life {
		case lifecycle.Looping:
			// 0 - Looping activations cyclically reactivate the same goroutine when the last finishes and the potential is high
			go func() {
				log.Verbosef(imp.Bridge, "initializing as a 'looping' synapse\n")
				for (*imp.Cortex).Alive() {
					inception := time.Now()
					timelineRollover(&imp.Timeline)
					imp.Timeline.Inception = inception
					imp.Beat = beat
					if neuron.Potential(imp) {
						imp.Timeline.Activation = time.Now()
						panicSafeAction(imp)
						imp.Timeline.Completion = time.Now()
						beat++
					}

					(*imp.Cortex).master.Lock()
					(*imp.Cortex).clock.Wait()
					(*imp.Cortex).master.Unlock()
				}

				wg := &sync.WaitGroup{}
				wg.Add(1)
				neuron.Cleanup(imp, wg)
				log.Verbosef(imp.Bridge, "ended\n")
			}()
		case lifecycle.Stimulative:
			// 1 - Stimulative activations launch new goroutines on every impulse the potential is high
			go func() {
				log.Verbosef(imp.Bridge, "initializing as a 'stimulative' synapse\n")
				for (*imp.Cortex).Alive() {
					inception := time.Now()
					timelineRollover(&imp.Timeline)
					imp.Timeline.Inception = inception
					imp.Beat = beat
					if neuron.Potential(imp) {
						imp.Timeline.Activation = time.Now()
						panicSafeAction(imp)
						imp.Timeline.Completion = time.Now()
						beat++
					}
					(*imp.Cortex).master.Lock()
					(*imp.Cortex).clock.Wait()
					(*imp.Cortex).master.Unlock()
				}

				wg := &sync.WaitGroup{}
				wg.Add(1)
				neuron.Cleanup(imp, wg)
				log.Verbosef(imp.Bridge, "ended\n")
			}()
		case lifecycle.Triggered:
			// 2 - Triggered activations are a one-shot GUARANTEE once the potential goes high
			go func() {
				log.Verbosef(imp.Bridge, "initializing as a 'triggered' synapse\n")
				imp.Timeline.Inception = time.Now()
				for (*imp.Cortex).Alive() && !neuron.Potential(imp) {
					(*imp.Cortex).master.Lock()
					(*imp.Cortex).clock.Wait()
					(*imp.Cortex).master.Unlock()
				}
				if (*imp.Cortex).Alive() {
					imp.Beat = beat
					imp.Timeline.Activation = time.Now()
					panicSafeAction(imp)
					imp.Timeline.Completion = time.Now()
					beat++
				}

				wg := &sync.WaitGroup{}
				wg.Add(1)
				neuron.Cleanup(imp, wg)
				log.Verbosef(imp.Bridge, "ended\n")
			}()
		case lifecycle.Impulse:
			// 3 - Impulse activations are a one-shot ATTEMPT regardless of potential
			go func() {
				log.Verbosef(imp.Bridge, "initializing as an 'impulsed' synapse\n")
				imp.Timeline.Inception = time.Now()
				if (*imp.Cortex).Alive() && neuron.Potential(imp) {
					imp.Beat = beat
					imp.Timeline.Activation = time.Now()
					panicSafeAction(imp)
					imp.Timeline.Completion = time.Now()
					beat++
				}

				wg := &sync.WaitGroup{}
				wg.Add(1)
				neuron.Cleanup(imp, wg)
				log.Verbosef(imp.Bridge, "ended\n")
			}()
		}
	}
}
