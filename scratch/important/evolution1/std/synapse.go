package std

import (
	"time"

	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/enum/life"
	"git.ignitelabs.net/janos/core/sys/rec"
)

type synapse struct {
	lifecycle life.Cycle
	name      string
	control   chan *Impulse

	timeline *TemporalBuffer[Activation]

	action    func(*Impulse)
	potential func(*Impulse) bool
	cleanup   func(*Impulse)
}

func newSynapse(lifecycle life.Cycle, named string, action func(*Impulse), potential func(*Impulse) bool, cleanup ...func(*Impulse)) synapse {
	var c func(*Impulse)
	if len(cleanup) > 0 {
		c = cleanup[0]
	}
	if action == nil {
		panic("the synaptic action must not be nil")
	}
	control := make(chan *Impulse, 1<<16)
	go func() {
		panicSafeCleanup := func(imp *Impulse) {
			defer func() {
				if r := recover(); r != nil {
					rec.Printf(named, "neural potentialpanic: %v\n", r)
				}
			}()
			if len(cleanup) > 0 && cleanup[0] != nil {
				cleanup[0](imp)
			}
		}
		lastCycle := -1
		panicSafePotential := func(imp *Impulse) bool {
			defer func() {
				if r := recover(); r != nil {
					rec.Printf(named, "neural potentialpanic: %v\n", r)
				}
			}()
			if int(imp.Activation.Cycle) > lastCycle && (potential == nil || potential(imp)) {
				lastCycle = int(imp.Activation.Cycle)
				return true
			}
			return false
		}
		panicSafeAction := func(imp *Impulse, decay bool) bool {
			defer func() {
				if r := recover(); r != nil {
					rec.Printf(named, "neural action panic: %v\n", r)
				}
			}()
			imp.Activation.Activation = core.Ref(time.Now())
			action(imp)

			imp.Activation.Completion = core.Ref(time.Now())
			imp.Timeline.Record(imp.Inception, imp.Activation)

			imp.recycle(imp.Decay)

			if decay || imp.Decay {
				panicSafeCleanup(imp)
				return true
			}
			return false
		}

		for core.Alive() {
			imp := <-control

			imp.Name = named

			if time.Since(imp.Inception) > imp.Period {
				// This is for when an impulse comes in before the prior looping activation finished.
				// Otherwise, long-running activations beyond the cortical period would queue up and repeat from constant impulses.
				continue
			}

			if panicSafePotential(imp) {
				switch lifecycle {
				case life.Looping:
					panicSafeAction(imp, false)
				case life.Stimulative:
					go panicSafeAction(imp, false)
				case life.Triggered:
					panicSafeAction(imp, true)
				case life.Impulse:
					panicSafeAction(imp, true)
				}
			}
		}
	}()
	return synapse{
		lifecycle: lifecycle,
		name:      named,
		control:   control,
		action:    action,
		potential: potential,
		cleanup:   c,
	}
}

func (s synapse) Action(imp *Impulse) {
	s.action(imp)
}

func (s synapse) Potential(imp *Impulse) bool {
	return s.potential(imp)
}

func (s synapse) Cleanup(imp *Impulse) {
	if s.cleanup != nil {
		s.cleanup(imp)
	}
}
