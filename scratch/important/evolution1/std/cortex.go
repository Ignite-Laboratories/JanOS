package std

import (
	"fmt"
	"sync"
	"time"

	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/sys/id"
	"git.ignitelabs.net/janos/core/sys/when"
)

type Cortex chan<- signal

// KeepAlive is a convenience passthrough function to core.KeepAlive - please see it's documentation for current information.
func (Cortex) KeepAlive(postDelay ...time.Duration) {
	core.KeepAlive(postDelay...)
}

func NewCortex(named string, period time.Duration, frequency float64, synapticLimit ...uint) Cortex {
	return NewCortexRef(named, core.Ref(period), core.Ref(frequency), synapticLimit...)
}

func NewCortexRef(named string, period *time.Duration, frequency *float64, channelLimit ...uint) Cortex {
	limit := 1 << 16
	if len(channelLimit) > 0 {
		limit = int(channelLimit[0])
	}
	control := make(chan signal, limit)

	core.Deferrals() <- func(wg *sync.WaitGroup) {
		control <- Signal.Decay(wg)
	}

	creation := time.Now()
	sparked := false
	muted := false
	adjustable := false
	step := uint(0)
	beat := uint(0)
	cycle := uint(0)
	last := time.Now()
	delay := *period
	held := make(chan signal, limit)
	synapseCount := uint(0)
	synapses := make(chan synapse, limit)

	f := when.HertzToDuration(*frequency)
	potential := func() bool {
		now := time.Now()
		current := when.HertzToDuration(*frequency)
		if f != current {
			f = current

			delay = f
			adjustable = true
			if now.After(last.Add(f)) {
				last = now
				return true
			}
		} else if now.After(last.Add(f)) {
			if adjustable {
				observed := now.Sub(last)
				adjustment := observed - f
				delay = f - now.Sub(last) - adjustment
			} else {
				delay = f - now.Sub(last)
			}
			last = now
			return true
		}
		return false
	}

	recycle := func() {
		for len(held) > 0 {
			control <- <-held
		}
	}

	go func() {
		for core.Alive() {
			if muted {
				raw := <-control
				if _, ok := raw.(unmute); !ok {
					held <- raw
				} else {
					muted = false
					recycle()
				}
			} else {
				select {
				case raw, open := <-control:
					if !open {
						panic("cortex channel closed")
					}
					// NOTE: Don't adjust the feedback loop when we get a control signal
					adjustable = false

					if _, ok := any(raw).(spark); !ok && !sparked {
						// Buffer any non-spark messages in the order they're received until a spark occurs
						held <- raw
					} else {
						switch msg := raw.(type) {
						case spark:
							fmt.Println("got spark")
							for _, spk := range msg {
								control <- spk
							}

							sparked = true
							recycle()
						case decay:
							// Do work here
							if msg.wait != nil {
								msg.wait.Done()
							}
						case mute:
							muted = true
						case unmute:
							muted = false
							recycle()
						case synapse:
							fmt.Println("got synapse")
							if muted {
								held <- msg
								continue
							}
							msg.timeline = NewTemporalBuffer[Activation]()
							synapses <- msg
							synapseCount++
						default:
							panic(fmt.Errorf("unknown signal message: %T", msg))
						}
					}
				case <-time.After(delay):
					// NOTE: Only adjust the feedback loop when we observe an uninterrupted period
					adjustable = true
				}

				if potential() {
					imp := &Impulse{
						cortexName: named,
						Activation: Activation{
							id:             id.Next(),
							CortexCreation: creation,
							Inception:      time.Now(),
							Step:           step,
							Beat:           beat,
							Cycle:          cycle,
							Period:         *period,
							Frequency:      *frequency,
						},
						Cortex: control,
					}

					if len(synapses) >= 0 {
						syn := <-synapses
						imp.Timeline = syn.timeline
						imp.recycle = func(decay bool) {
							if decay {
								if synapseCount > 0 {
									synapseCount--
								}
							} else {
								synapses <- syn
							}
						}
						syn.control <- imp
					}

					step++

					// TODO: IMPORTANT!!! Beat and cycle MUST be calculated for the system to work
					// Move back from the frequency system to just use the synaptic count - frequency is implied by period

					beat++
					if float64(beat) > *frequency {
						beat = 0
						cycle++
					}
				}
			}
		}
	}()

	return control
}
