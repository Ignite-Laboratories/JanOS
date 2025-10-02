package std

import (
	"sync"
	"time"

	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/enum/life"
	"git.ignitelabs.net/janos/core/sys/rec"
)

// A Cluster sequentially activates all neural activity distributed across a period of time.
//
// 0 - The distributed activity should always attempt to activate in the same order
// 1 - Activations should adjust their activation time automatically when any neuron decays out or is added to the stack
// 2 - Looping activations should not reactivate if they are still running
type Cluster struct {
	loop      chan Neural
	stimulate chan Neural
	trigger   chan Neural
	impulse   chan Neural

	neuralCount uint

	Mute   bool
	Decay  bool
	Period *time.Duration // A nil or negative period should be treated as '1 s'

	spark func()

	last time.Time
}

func (cls *Cluster) sanityCheck() {
	if cls.spark == nil {
		panic("a cluster's spark cannot be nil")
	}
	if cls.loop == nil {
		panic("a cluster's loop channel can't be nil")
	}
	if cls.stimulate == nil {
		panic("a cluster's stimulate channel can't be nil")
	}
}

func (cls *Cluster) Loop() chan<- Neural {
	cls.sanityCheck()
	return cls.loop
}

func (cls *Cluster) Stimulate() chan<- Neural {
	cls.sanityCheck()
	return cls.stimulate
}

func (cls *Cluster) Trigger() chan<- Neural {
	cls.sanityCheck()
	return cls.trigger
}

func (cls *Cluster) Impulse() chan<- Neural {
	cls.sanityCheck()
	return cls.impulse
}

func (cls *Cluster) Spark() {
	cls.sanityCheck()
	cls.spark()
}

func (ctx *Cortex) CreateCluster(named string, period *time.Duration, potential func(*Impulse) bool, cleanup ...func(*Impulse)) *Cluster {
	ctx.sanityCheck()

	type endpoint struct {
		Neural

		bridgeStr string

		running     bool
		stimulative bool
		triggered   bool
		impulsed    bool

		impulse *Impulse

		fire chan *Impulse
	}

	cls := &Cluster{
		loop:      make(chan Neural, ctx.limit),
		stimulate: make(chan Neural, ctx.limit),
		trigger:   make(chan Neural, ctx.limit),
		impulse:   make(chan Neural, ctx.limit),

		Period: period,
		last:   time.Now(),
	}
	cls.spark = func() {
		endpoints := make(chan *endpoint, ctx.limit)

		ctx.Deferrals() <- func(wg *sync.WaitGroup) {
			for len(endpoints) > 0 {
				end := <-endpoints
				if cls.neuralCount > 0 {
					cls.neuralCount--
				}
				if end.Cleanup != nil {
					end.Cleanup(nil)
				}
				end.fire <- nil
				rec.Verbosef(end.bridgeStr, "decayed\n")
			}
			wg.Done()
		}

		loadFn := func(bridge Bridge) {
			receiveFn := func(n Neural) *endpoint {
				rec.Verbosef(bridge.String(), "wiring axon to neural endpoint '%v'\n", n.Named())

				end := &endpoint{
					Neural: n,
					fire:   make(chan *Impulse, 1<<16),
					impulse: &Impulse{
						Bridge:   append(bridge, n.Named()),
						Timeline: NewTimeline(),
						Cortex:   ctx,
					},
				}

				go func() {
					decay := false
					for ctx.Alive() && core.Alive() && !decay {
						impulse := <-end.fire
						if !ctx.Alive() || !core.Alive() {
							return
						}

						end.bridgeStr = impulse.Bridge.String()

						end.running = true

						activation := func() {
							end.Action(impulse)
							impulse.Timeline.setCompleted(impulse.currentEvent.id, time.Now())
							impulse.Count++
							end.running = false

							if end.triggered || end.impulsed {
								impulse.Decay = true
							}
							if impulse.Decay == true {
								if end.Cleanup != nil {
									end.Cleanup(impulse)
								}
								cls.neuralCount--
								rec.Verbosef(impulse.Bridge.String(), "decayed\n")
								decay = true
							} else if !end.stimulative {
								endpoints <- end
							}
						}

						if end.stimulative {
							endpoints <- end
							go activation()
						} else {
							activation()
						}
					}
				}()

				cls.neuralCount++
				return end
			}

			for len(cls.loop) > 0 || len(cls.stimulate) > 0 || len(cls.trigger) > 0 || len(cls.impulse) > 0 {
				select {
				case n := <-cls.loop:
					endpoints <- receiveFn(n)
				case n := <-cls.stimulate:
					end := receiveFn(n)
					end.stimulative = true
					endpoints <- end
				case n := <-cls.trigger:
					end := receiveFn(n)
					end.triggered = true
					endpoints <- end
				case n := <-cls.impulse:
					end := receiveFn(n)
					end.impulsed = true
					endpoints <- end
				}
			}
		}

		ctx.Synapses() <- NewSynapse(life.Looping, named, func(imp *Impulse) {
			if cls.Decay {
				imp.Decay = true
				return
			}

			loadFn(imp.Bridge)

			if len(endpoints) == 0 {
				return
			}

			end := <-endpoints

			end.impulse.currentEvent = imp.currentEvent

			if (end.stimulative || !end.running) && end.Potential(end.impulse) {
				end.impulse.Beat = ctx.beat
				end.impulse.BeatPeriod = ctx.BeatPeriod
				end.impulse.currentEvent.Activation = time.Now()
				end.impulse.Timeline.Add(*end.impulse.currentEvent)
				end.fire <- end.impulse
			}
		}, func(imp *Impulse) bool {
			loadFn(imp.Bridge)

			if cls.neuralCount == 0 || cls.Mute {
				return false
			}

			p := time.Second
			if cls.Period != nil && *cls.Period > 0 {
				p = *cls.Period
			}

			now := time.Now()
			offset := time.Duration(p.Nanoseconds() / int64(cls.neuralCount))
			if now.After(cls.last.Add(offset)) {
				cls.last = now

				if potential == nil {
					return true
				}
				return potential(imp)
			}
			return false
		}, cleanup...)
	}

	return cls
}
