package std

import (
	"sync"
	"time"

	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/enum/lifecycle"
	"git.ignitelabs.net/janos/core/sys/id"
	"git.ignitelabs.net/janos/core/sys/rec"
)

// A Synapse represents a fixed impulsive activation between a Neuron and Cortex.  A synapse can be used to recycle
// the same action across many cortices, as it can be sparked as many times as you would like.
type Synapse func(*Impulse)

// NewSynapse creates a Neural connection between a Neuron and a Cortex.  You may optionally provide 'nil' to the potential if you'd like to imply 'always fire'.
func NewSynapse(life lifecycle.Lifecycle, neuronName string, action func(*Impulse), potential func(*Impulse) bool, cleanup ...func(*Impulse)) Synapse {
	return NewSynapseFromNeural(life, NewNeuron(neuronName, action, potential, cleanup...))
}

// NewSynapseFromNeural creates a neural Synapse which fires the provided action potential according to the provided lifecycle.Lifecycle.  You may optionally
// provide a cleanup function which is called after this Synapse finishes all neural activation (or the cortex shuts down).  For triggered
// or impulsed lifecycles, this gets called immediately - for looping or stimulative, this gets called after the cortex shuts down.
func NewSynapseFromNeural(life lifecycle.Lifecycle, neuron Neural) Synapse {
	rec.Verbosef(core.ModuleName, "%v is creating synapse '%s'\n", core.Name.Name, neuron.Named())
	count := uint(0)
	return func(imp *Impulse) {
		creation := time.Now()
		imp.Bridge = []string{(*imp.Cortex).Named(), neuron.Named()}
		imp.Neuron = neuron

		rec.Verbosef((*imp.Cortex).Named(), "wiring axon to neural endpoint '%s'\n", neuron.Named())

		panicSafeAction := func(i *Impulse) {
			defer func() {
				if r := recover(); r != nil {
					rec.Printf(i.Bridge.String(), "neural panic: %s\n", r)
				}
			}()

			neuron.Action(i)
		}

		switch life {
		case lifecycle.Looping:
			// 0 - Looping activations cyclically reactivate the same goroutine when the last finishes and the potential is high
			go func() {
				rec.Verbosef(imp.Bridge.String(), "looping\n")
				for (*imp.Cortex).Alive() {
					event := SynapticEvent{
						id:              id.Next(),
						SynapseCreation: creation,
						Inception:       time.Now(),
					}
					imp.currentEvent = event
					imp.Count = count
					imp.Beat = (*imp.Cortex).beat
					imp.Phase = (*imp.Cortex).Phase
					if neuron.Potential(imp) && !imp.Mute && (*imp.Cortex).Alive() {
						event.Activation = time.Now()
						imp.Timeline.Add(event)
						panicSafeAction(imp)
						imp.Timeline.setCompleted(event.id, time.Now())
						count++
					}

					if !imp.Decay && (*imp.Cortex).Alive() {
						(*imp.Cortex).master.Lock()
						(*imp.Cortex).clock.Wait()
						(*imp.Cortex).master.Unlock()
					} else {
						break
					}
				}
				(*imp.Cortex).hold.Add(1)
				if neuron.Cleanup != nil {
					neuron.Cleanup(imp)
				}
				rec.Verbosef(imp.Bridge.String(), "decayed\n")
				(*imp.Cortex).hold.Done()
			}()
		case lifecycle.Stimulative:
			// 1 - Stimulative activations launch new goroutines on every impulse the potential is high
			go func() {
				rec.Verbosef(imp.Bridge.String(), "stimulating\n")
				for (*imp.Cortex).Alive() {
					event := SynapticEvent{
						id:              id.Next(),
						SynapseCreation: creation,
						Inception:       time.Now(),
					}
					imp.currentEvent = event
					imp.Count = count
					imp.Beat = (*imp.Cortex).beat
					imp.Phase = (*imp.Cortex).Phase
					if neuron.Potential(imp) && !imp.Mute && (*imp.Cortex).Alive() {
						event.Activation = time.Now()
						imp.Timeline.Add(event)
						go func() {
							panicSafeAction(imp)
							imp.Timeline.setCompleted(event.id, time.Now())
						}()
						count++
					}

					if !imp.Decay && (*imp.Cortex).Alive() {
						(*imp.Cortex).master.Lock()
						(*imp.Cortex).clock.Wait()
						(*imp.Cortex).master.Unlock()
					} else {
						break
					}
				}
				(*imp.Cortex).hold.Add(1)
				if neuron.Cleanup != nil {
					neuron.Cleanup(imp)
				}
				rec.Verbosef(imp.Bridge.String(), "decayed\n")
				(*imp.Cortex).hold.Done()
			}()
		case lifecycle.Triggered:
			// 2 - Triggered activations are a one-shot GUARANTEE once the potential goes high
			go func() {
				rec.Verbosef(imp.Bridge.String(), "setting a trigger\n")
				event := SynapticEvent{
					id:              id.Next(),
					SynapseCreation: creation,
					Inception:       time.Now(),
				}
				imp.currentEvent = event
				for (*imp.Cortex).Alive() && !neuron.Potential(imp) {
					(*imp.Cortex).master.Lock()
					(*imp.Cortex).clock.Wait()
					(*imp.Cortex).master.Unlock()
				}
				if (*imp.Cortex).Alive() && !imp.Mute {
					imp.Count = count
					imp.Beat = (*imp.Cortex).beat
					imp.Phase = (*imp.Cortex).Phase
					event.Activation = time.Now()
					imp.Timeline.Add(event)
					panicSafeAction(imp)
					imp.Timeline.setCompleted(event.id, time.Now())
					count++
				}
				(*imp.Cortex).hold.Add(1)
				if neuron.Cleanup != nil {
					neuron.Cleanup(imp)
				}
				rec.Verbosef(imp.Bridge.String(), "decayed\n")
				(*imp.Cortex).hold.Done()
			}()
		case lifecycle.Impulse:
			// 3 - Impulse activations are a one-shot ATTEMPT regardless of potential
			go func() {
				rec.Verbosef(imp.Bridge.String(), "impulsing\n")
				event := SynapticEvent{
					id:              id.Next(),
					SynapseCreation: creation,
					Inception:       time.Now(),
				}
				imp.currentEvent = event
				if (*imp.Cortex).Alive() && neuron.Potential(imp) && !imp.Mute {
					imp.Count = count
					imp.Beat = (*imp.Cortex).beat
					imp.Phase = (*imp.Cortex).Phase
					event.Activation = time.Now()
					imp.Timeline.Add(event)
					panicSafeAction(imp)
					imp.Timeline.setCompleted(event.id, time.Now())
					count++
					imp.Timeline.Add(event)
				}
				(*imp.Cortex).hold.Add(1)
				if neuron.Cleanup != nil {
					neuron.Cleanup(imp)
				}
				rec.Verbosef(imp.Bridge.String(), "decayed\n")
				(*imp.Cortex).hold.Done()
			}()
		}
	}
}

// NewSynapticCluster creates a synapse which gates neural activity in a round-robin loop.  This means that synapses
// will be activated sequentially in a loop respecting the order they're received.
func NewSynapticCluster(named string, ctx *Cortex, potential func(*Impulse) bool, cleanup ...func(*Impulse)) (Synapse, chan<- Neural) {
	type endpoint struct {
		Neural
		impulse *Impulse
	}

	input := make(chan Neural, ctx.limit)
	neurons := make(chan endpoint, ctx.limit)

	ctx.Deferrals() <- func(wg *sync.WaitGroup) {
		for len(neurons) > 0 {
			n := <-neurons
			if n.Cleanup != nil {
				n.Cleanup(n.impulse)
			}
			rec.Verbosef(n.impulse.Bridge.String(), "decayed\n")
		}
		wg.Done()
	}

	return NewSynapse(lifecycle.Looping, named, func(imp *Impulse) {
		for len(input) > 0 {
			select {
			case n := <-input:
				rec.Verbosef(imp.Bridge.String(), "wiring axon to neural endpoint '%v'\n", n.Named())
				bridge := append(imp.Bridge, n.Named())

				neurons <- endpoint{
					Neural: n,
					impulse: &Impulse{
						Bridge:   bridge,
						Timeline: NewTimeline(),
						Cortex:   imp.Cortex,
					},
				}
			}
		}

		next := <-neurons
		if next.Potential(next.impulse) {
			event := imp.currentEvent

			next.impulse.Beat = ctx.beat
			next.impulse.Phase = ctx.Phase
			event.Activation = time.Now()
			next.impulse.Timeline.Add(event)
			next.Action(next.impulse)
			next.impulse.Timeline.setCompleted(event.id, time.Now())
			next.impulse.Count++
		}
		if next.impulse.Decay == false {
			neurons <- next
		} else {
			if next.Cleanup != nil {
				next.Cleanup(next.impulse)
			}
			rec.Verbosef(next.impulse.Bridge.String(), "decayed\n")
		}
	}, potential, cleanup...), input
}
