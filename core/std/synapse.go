package std

import (
	"time"

	"git.ignitelabs.net/core/enum/lifecycle"
	"git.ignitelabs.net/core/sys/log"
)

type synapse func(Context)

// NewSynapse creates a neural synapse which fires the provided action potential according to the provided Lifecycle.  You may optionally
// provide a cleanup function which is called after this synapse finishes all neural activation (or the cortex shuts down).  For triggered
// or impulsed lifecycles, this happens immediately - for looping or stimulative, this happens after the cortex shuts down or when the
// provided action returns false.
//
// NOTE: For stimulative activations, the action may still fire a few times after returning false - but cleanup will happen after all
//
//	activations are complete.
func NewSynapse(life lifecycle.Lifecycle, neuron Neuron) synapse {
	beat := 0
	return func(ctx Context) {
		ctx.Beat = beat
		ctx.ModuleName = (*ctx.Cortex).Named() + "." + neuron.Named()
		beat++

		log.Printf((*ctx.Cortex).Named(), "created synapse to '%s'\n", neuron.Named())

		panicSafeAction := func(ctx Context) {
			defer func() {
				if r := recover(); r != nil {
					log.Printf(ctx.ModuleName, "neural panic: %s\n", r)
				}
			}()

			neuron.Action(ctx)
		}

		switch life {
		case lifecycle.Looping:
			// 0 - Looping activations cyclically reactivate the same goroutine when the last finishes and the potential is high
			go func() {
				log.Verbosef(ctx.ModuleName, "sparked looping activation\n")
				for (*ctx.Cortex).Alive() {
					if neuron.Potential(ctx) {
						ctx.ResponseTime = time.Now().Sub(ctx.Moment)
						panicSafeAction(ctx)
					}
					(*ctx.Cortex).master.Lock()
					(*ctx.Cortex).master.Unlock()
				}

				neuron.Cleanup()
			}()
		case lifecycle.Stimulative:
			// 1 - Stimulative activations launch new goroutines on every impulse the potential is high
			go func() {
				log.Verbosef(ctx.ModuleName, "sparked stimulative activation\n")
				for (*ctx.Cortex).Alive() {
					if neuron.Potential(ctx) {
						ctx.ResponseTime = time.Now().Sub(ctx.Moment)
						go panicSafeAction(ctx)
					}
					(*ctx.Cortex).master.Lock()
					(*ctx.Cortex).master.Unlock()
				}

				neuron.Cleanup()
			}()
		case lifecycle.Triggered:
			// 2 - Triggered activations are a one-shot GUARANTEE once the potential goes high
			go func() {
				log.Verbosef(ctx.ModuleName, "sparked triggered activation\n")
				for (*ctx.Cortex).Alive() && !neuron.Potential(ctx) {
					(*ctx.Cortex).master.Lock()
					(*ctx.Cortex).master.Unlock()
				}
				if (*ctx.Cortex).Alive() {
					ctx.ResponseTime = time.Now().Sub(ctx.Moment)
					panicSafeAction(ctx)
				}

				neuron.Cleanup()
			}()
		case lifecycle.Impulse:
			// 3 - Impulse activations are a one-shot ATTEMPT regardless of potential
			go func() {
				log.Verbosef(ctx.ModuleName, "sparked impulse activation\n")
				if (*ctx.Cortex).Alive() && neuron.Potential(ctx) {
					ctx.ResponseTime = time.Now().Sub(ctx.Moment)
					panicSafeAction(ctx)
				}

				neuron.Cleanup()
			}()
		}
	}
}
