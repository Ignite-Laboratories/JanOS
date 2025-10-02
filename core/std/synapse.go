package std

import (
	"time"

	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/enum/life"
	"git.ignitelabs.net/janos/core/sys/id"
	"git.ignitelabs.net/janos/core/sys/rec"
)

// A Synapse represents a fixed impulsive activation between a Neuron and Cortex.  A synapse can be used to recycle
// the same action across many cortices, as it can be sparked as many times as you would like.
type Synapse func(*Impulse)

// NewSynapse creates a Neural connection between a Neuron and a Cortex.  You may optionally provide 'nil' to the potential if you'd like to imply 'always fire'.
func NewSynapse(lifeycle life.Cycle, neuronName string, action func(*Impulse), potential func(*Impulse) bool, cleanup ...func(*Impulse)) Synapse {
	return NewSynapseFromNeural(lifeycle, NewNeuron(neuronName, action, potential, cleanup...))
}

// NewSynapseFromNeural creates a neural Synapse which fires the provided action potential according to the provided life.Cycle.  You may optionally
// provide a cleanup function which is called after this Synapse finishes all neural activation (or the cortex shuts down).  For triggered
// or impulsed lifecycles, this gets called immediately - for looping or stimulative, this gets called after the cortex shuts down.
func NewSynapseFromNeural(lifeycle life.Cycle, neuron Neural) Synapse {
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

		switch lifeycle {
		case life.Looping:
			// 0 - Looping activations cyclically reactivate the same goroutine when the last finishes and the potential is high
			go func() {
				rec.Verbosef(imp.Bridge.String(), "looping\n")
				for (*imp.Cortex).Alive() {
					event := &SynapticEvent{
						id:              id.Next(),
						SynapseCreation: creation,
						Inception:       time.Now(),
					}
					imp.currentEvent = event
					imp.Count = count
					imp.Beat = (*imp.Cortex).beat
					imp.BeatPeriod = (*imp.Cortex).BeatPeriod
					if neuron.Potential(imp) && !imp.Mute && (*imp.Cortex).Alive() {
						event.Activation = time.Now()
						imp.Timeline.Add(*event)
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
		case life.Stimulative:
			// 1 - Stimulative activations launch new goroutines on every impulse the potential is high
			go func() {
				rec.Verbosef(imp.Bridge.String(), "stimulating\n")
				for (*imp.Cortex).Alive() {
					event := &SynapticEvent{
						id:              id.Next(),
						SynapseCreation: creation,
						Inception:       time.Now(),
					}
					imp.currentEvent = event
					imp.Count = count
					imp.Beat = (*imp.Cortex).beat
					imp.BeatPeriod = (*imp.Cortex).BeatPeriod
					if neuron.Potential(imp) && !imp.Mute && (*imp.Cortex).Alive() {
						event.Activation = time.Now()
						imp.Timeline.Add(*event)
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
		case life.Triggered:
			// 2 - Triggered activations are a one-shot GUARANTEE once the potential goes high
			go func() {
				rec.Verbosef(imp.Bridge.String(), "setting a trigger\n")
				event := &SynapticEvent{
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
					imp.BeatPeriod = (*imp.Cortex).BeatPeriod
					event.Activation = time.Now()
					imp.Timeline.Add(*event)
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
		case life.Impulse:
			// 3 - Impulse activations are a one-shot ATTEMPT regardless of potential
			go func() {
				rec.Verbosef(imp.Bridge.String(), "impulsing\n")
				event := &SynapticEvent{
					id:              id.Next(),
					SynapseCreation: creation,
					Inception:       time.Now(),
				}
				imp.currentEvent = event
				if (*imp.Cortex).Alive() && neuron.Potential(imp) && !imp.Mute {
					imp.Count = count
					imp.Beat = (*imp.Cortex).beat
					imp.BeatPeriod = (*imp.Cortex).BeatPeriod
					event.Activation = time.Now()
					imp.Timeline.Add(*event)
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
		}
	}
}
