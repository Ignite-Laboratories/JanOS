package core

import "sync"

type synapse func(Context)

var master sync.Mutex
var clock = sync.NewCond(&master)

// NewSynapse creates a neural synapse which fires the provided action potential according to the provided Lifecycle.  You may optionally
// provide a cleanup function which is called after this synapse finishes all neural activation (or the cortex shuts down).  For triggered
// or impulsed lifecycles, this happens immediately - for looping or stimulative, this happens after the cortex shuts down or when the
// provided action returns false.
//
// NOTE: For stimulative activations, the action may still fire a few times after returning false - but cleanup will happen after all
//
//	activations are complete.
func NewSynapse(lifecycle Lifecycle, neuron *Neuron) synapse {
	var mu sync.Mutex
	impulse := make(chan Context)
	return func(ctx Context) {
		if mu.TryLock() {
			defer mu.Unlock()

			switch lifecycle {
			case Looping:
				// Looping activations cyclically activate the same goroutine when the last finishes and the potential is high
				go func() {
					n := *neuron
					for Alive() {
						master.Lock()
						clock.Wait()
						master.Unlock()
						if n.Potential(ctx) {
							if !n.Action(ctx) {
								n.Cleanup()
								return
							}
						}
					}
					n.Cleanup()
				}()
			case Stimulative:
				// Stimulative activations launch new goroutines while the potential is high
				go func() {
					n := *neuron
					a := true
					for Alive() && a {
						master.Lock()
						clock.Wait()
						master.Unlock()
						if a && n.Potential(ctx) {
							go func() {
								if a && !n.Action(ctx) {
									a = false
								}
							}()
						}
					}
					n.Cleanup()
				}()
			case Triggered:
				// Triggered activations are a one-shot GUARANTEE once the potential goes high
				go func() {
					n := *neuron
					for !n.Potential(ctx) {
					}
					n.Action(ctx)
					n.Cleanup()
				}()
			case Impulse:
				// Impulse activations are a one-shot ATTEMPT
				go func() {
					n := *neuron
					if n.Potential(ctx) {
						n.Action(ctx)
						n.Cleanup()
					}
				}()
			}
		} else {
			impulse <- ctx
		}
	}
}
