package std

import (
	"math"
	"sync"
	"time"

	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/sys/atlas"
	"git.ignitelabs.net/janos/core/sys/given/format"
	"git.ignitelabs.net/janos/core/sys/rec"
)

// A Cortex represents a source of neural impulses.  It defines the frequency which synaptic activity can fire at.
type Cortex struct {
	Entity

	// Frequency defines the minimum frequency impulses will fire.
	//
	// NOTE: If you set this to zero or negative, they will fire as fast as possible.  For a zero frequency, please mute the cortex.
	// Otherwise, we'd have to divide by zero =)
	Frequency float64

	// BeatPeriod defines the number of beats the cortex will count to before looping back to zero.
	//
	// NOTE: Set this to a negative value for an infinite phase =)
	BeatPeriod int
	beat       uint

	inception time.Time

	synapses chan Synapse
	timeline []time.Time

	deferrals    chan func(*sync.WaitGroup)
	deferralWait *sync.WaitGroup

	mute         chan any
	unmute       chan any
	impulse      chan any
	shutdown     chan any
	closed       chan any // NOTE: closed is used to 'close' signal when the cortex is shutting down
	shutdownWait *sync.WaitGroup
	hold         *sync.WaitGroup

	alive   bool
	created bool
	running bool
	limit   int

	timeLock sync.Mutex
	master   sync.Mutex
	clock    sync.Cond
}

// NewCortex creates a new named Cortex limited to the provided number of neural activations.
//
// If you'd like a randomly generated name, see given.Random[ format.Format ]
//
// NOTE: If no limit is provided, the default is 2¹⁶ - this can generally be ignored for most systems.
func NewCortex(named string, synapticLimit ...int) *Cortex {
	limit := 1 << 16
	if len(synapticLimit) > 0 {
		limit = synapticLimit[0]
	}

	c := &Cortex{
		Entity:       NewEntity[format.Default](),
		inception:    time.Now(),
		synapses:     make(chan Synapse, limit),
		deferrals:    make(chan func(*sync.WaitGroup), 1<<16),
		deferralWait: &sync.WaitGroup{},
		mute:         make(chan any, 1<<16),
		unmute:       make(chan any, 1<<16),
		impulse:      make(chan any, 1<<16),
		shutdown:     make(chan any, 1<<16),
		closed:       make(chan any, 1<<16),
		limit:        limit,
		alive:        true,
		created:      true,
	}
	c.clock = sync.Cond{L: &c.master}
	c.Entity.Name.Name = named
	c.hold = &sync.WaitGroup{}

	rec.Verbosef(core.ModuleName, "%v has created cortex '%s'\n", core.Name.Name, c.Named())
	return c
}

func (ctx *Cortex) Phase(frequency float64) float64 {
	return 1.0 * math.Sin((2*math.Pi*frequency)*(time.Now().Sub(ctx.Inception()).Seconds()))
}

// Optional: return the instantaneous phase angle in [0, 2π)
func (ctx *Cortex) PhaseAngle(frequency float64) float64 {
	if frequency <= 0 {
		return 0
	}
	t := time.Since(ctx.inception).Seconds()
	theta := 2 * math.Pi * frequency * t
	// normalize to [0, 2π)
	theta = math.Mod(theta, 2*math.Pi)
	if theta < 0 {
		theta += 2 * math.Pi
	}
	return theta
}

// Optional: return the normalized phase fraction in [0, 1)
func (ctx *Cortex) PhaseFraction(frequency float64) float64 {
	if frequency <= 0 {
		return 0
	}
	t := time.Since(ctx.inception).Seconds()
	p := math.Mod(frequency*t, 1.0)
	if p < 0 {
		p += 1.0
	}
	return p
}

// Impulse causes the cortex to fire a single impulse cycle.  Please note this is an asynchronous invocation.
//
// NOTE: If your cortex is phase-locked, this will inherently break its ability to track the phase momentarily.
// This is because phase-locking (at the cortex level) relies on tracking phase relative to the last impulse moment,
// which shifts when an impulse is fired.
func (ctx *Cortex) Impulse() {
	ctx.impulse <- nil
}

func _hertzToDuration(hz float64) time.Duration {
	if hz <= 0 {
		// No division by zero
		hz = 1e-100 // math.SmallestNonzeroFloat64 🡨 NOTE: Raspberry Pi doesn't handle this constant well
	}
	s := 1 / hz
	ns := s * 1e9
	return time.Duration(ns)
}

func (ctx *Cortex) Spark(synapses ...Synapse) {
	ctx.sanityCheck()

	select {
	case _, ok := <-ctx.closed:
		if !ok {
			panic("cannot re-spark a cortex after shutdown - please create a new cortex")
		}
	default:
	}

	rec.Verbosef(ctx.Named(), "sparking neural activity\n")

	for _, syn := range synapses {
		ctx.synapses <- syn
	}

	if ctx.running {
		return
	}

	core.Deferrals() <- func(wg *sync.WaitGroup) {
		ctx.shutdownWait = wg
		ctx.Shutdown()
	}

	go func() {
		ctx.alive = true
		ctx.running = true

		defer func() {
			count := len(ctx.deferrals)
			if count > 0 {
				if count > 1 {
					rec.Verbosef(ctx.Named(), "running %d deferrals\n", count)
				} else {
					rec.Verbosef(ctx.Named(), "running %d deferral\n", count)
				}
				for len(ctx.deferrals) > 0 {
					deferral := <-ctx.deferrals
					if deferral != nil {
						ctx.deferralWait.Add(1)
						go func() {
							defer func() {
								if r := recover(); r != nil {
									rec.Printf(ctx.Named(), "deferral error: %v\n", r)
								}
							}()

							deferral(ctx.deferralWait)
						}()
					}
				}
				ctx.deferralWait.Wait()
			}
			time.Sleep(time.Second)
			ctx.hold.Wait()
			rec.Verbosef(ctx.Named(), "cortex shut down complete\n")
			ctx.shutdownWait.Done()
		}()

		initial := true
		last := time.Now()
		var expected time.Duration
		var adjustment time.Duration
		var frequency float64

	main:
		for ctx.Alive() {
			if ctx.Frequency <= 0 {
				// This is a 'free-spin' condition
				select {
				case <-ctx.mute:
					rec.Verbosef(ctx.Named(), "muting\n")
					select {
					case <-ctx.shutdown:
						break main
					case <-ctx.impulse:
						// NOTE: Impulse requests should not break the muted condition
						ctx.mute <- nil
					case <-ctx.unmute:
						rec.Verbosef(ctx.Named(), "unmuting\n")
						for len(ctx.mute) > 0 {
							<-ctx.mute
						}
						for len(ctx.unmute) > 0 {
							<-ctx.unmute
						}
					}
				case <-ctx.unmute:
					continue // drain stray unmute signals
				default:
				}
			} else {
				// This is a 'timer-step' condition
				expected = last.Add(_hertzToDuration(ctx.Frequency)).Sub(time.Now().Add(adjustment))
				frequency = ctx.Frequency
				select {
				case <-ctx.shutdown:
					break main
				case <-ctx.impulse:
					rec.Verbosef(ctx.Named(), "impulsing\n")
				case <-time.After(expected):
					observed := time.Since(last)
					adjustment = observed - expected

					// If the frequency changed between cycles, don't try to 'adjust' it =)
					if ctx.Frequency != frequency {
						adjustment = 0
					}
				case <-ctx.mute:
					rec.Verbosef(ctx.Named(), "muting\n")
					select {
					case <-ctx.shutdown:
						break main
					case <-ctx.impulse:
						rec.Verbosef(ctx.Named(), "impulsing\n")
						// NOTE: Impulse requests should not break the muted condition
						ctx.mute <- nil
					case <-ctx.unmute:
						rec.Verbosef(ctx.Named(), "unmuting\n")
						for len(ctx.mute) > 0 {
							<-ctx.mute
						}
						for len(ctx.unmute) > 0 {
							<-ctx.unmute
						}
					}
				case <-ctx.unmute:
					continue // drain stray unmute signals
				}
			}

			for len(ctx.synapses) > 0 {
				imp := &Impulse{
					Cortex:   ctx,
					Timeline: NewTimeline(),
				}
				syn := <-ctx.synapses
				syn(imp)
			}

			if initial {
				initial = false
			} else {
				ctx.beat++
				if ctx.BeatPeriod > 0 && ctx.beat > uint(ctx.BeatPeriod) {
					ctx.beat = 0
				}
			}

			ctx.clock.Broadcast()
			ctx.addToTimeline(time.Now())
			last = time.Now()
		}

		rec.Verbosef(ctx.Named(), "decayed\n")

		// This beat frees the synapses to complete their activation and exit
		ctx.clock.Broadcast()
	}()
}

func (ctx *Cortex) Shutdown(delay ...time.Duration) {
	ctx.sanityCheck()

	if !ctx.alive {
		return
	}

	if len(delay) > 0 {
		rec.Verbosef(ctx.Named(), "cortex shutting down in %v\n", delay[0])
		time.Sleep(delay[0])
	}

	ctx.master.Lock()
	rec.Verbosef(ctx.Named(), "cortex shutting down\n")
	ctx.running = false
	ctx.alive = false
	ctx.shutdown <- nil
	close(ctx.closed)
	ctx.master.Unlock()
}

func (ctx *Cortex) addToTimeline(moment time.Time) {
	ctx.timeLock.Lock()
	defer ctx.timeLock.Unlock()

	ctx.timeline = append(ctx.timeline, moment)

	var trim int
	for i := range ctx.timeline {
		if ctx.timeline[i].Before(moment.Add(-atlas.ObservanceWindow)) {
			trim++
		} else {
			break
		}
	}
	ctx.timeline = ctx.timeline[trim:]
}

func (ctx *Cortex) Timeline() []time.Time {
	ctx.timeLock.Lock()
	defer ctx.timeLock.Unlock()
	return ctx.timeline
}

func (ctx *Cortex) Alive() bool {
	ctx.sanityCheck()

	return ctx.alive
}

func (ctx *Cortex) Mute() {
	ctx.sanityCheck()

	ctx.mute <- ctx.Entity
}

func (ctx *Cortex) Unmute() {
	ctx.sanityCheck()

	ctx.unmute <- ctx.Entity
}

func (ctx *Cortex) Inception() time.Time {
	ctx.sanityCheck()

	return ctx.inception
}

func (ctx *Cortex) Synapses() chan<- Synapse {
	ctx.sanityCheck()

	return ctx.synapses
}

func (ctx *Cortex) Deferrals() chan<- func(*sync.WaitGroup) {
	ctx.sanityCheck()

	return ctx.deferrals
}

func (ctx *Cortex) sanityCheck() {
	if !ctx.created {
		panic("cortices must be created through NewCortex")
	}
	if ctx.deferrals == nil {
		panic("deferrals must not be nil")
	}
	if ctx.synapses == nil {
		panic("synapses must not be nil")
	}
}
