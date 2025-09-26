package std

import (
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

	inception time.Time

	synapses chan Synapse
	timeline []time.Time

	deferrals    chan func(*sync.WaitGroup)
	deferralWait *sync.WaitGroup

	mute         chan any
	unmute       chan any
	impulse      chan any
	shutdown     chan any
	shutdownWait *sync.WaitGroup

	alive   bool
	created bool
	running bool

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
		deferrals:    make(chan func(*sync.WaitGroup), limit),
		deferralWait: &sync.WaitGroup{},
		mute:         make(chan any, 1<<16),
		unmute:       make(chan any, 1<<16),
		impulse:      make(chan any, 1<<16),
		shutdown:     make(chan any, 1<<16),
		alive:        true,
		created:      true,
	}
	c.clock = sync.Cond{L: &c.master}
	c.Entity.Name.Name = named

	rec.Verbosef(core.ModuleName, "created cortex '%s'\n", c.Named())
	return c
}

// Impulse causes the cortex to fire a single impulse cycle.  Please note this is an asynchronous invocation.
//
// NOTE: If your cortex is phase-locked, this will inherently break its ability to track the phase momentarily.
// This is because phase-locking (at the cortex level) relies on tracking phase relative to the last impulse moment,
// which shifts when an impulse is fired.
func (c *Cortex) Impulse() {
	c.impulse <- nil
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

func (c *Cortex) Spark(synapses ...Synapse) {
	c.sanityCheck()

	rec.Verbosef(c.Named(), "sparking neural activity\n")

	for _, syn := range synapses {
		c.synapses <- syn
	}

	if c.running {
		return
	}

	core.Deferrals() <- func(wg *sync.WaitGroup) {
		c.shutdownWait = wg
		c.Shutdown()
	}

	go func() {
		c.alive = true
		c.running = true

		defer func() {
			count := len(c.deferrals)
			if count > 0 {
				if count > 1 {
					rec.Verbosef(c.Named(), "running %d deferrals\n", count)
				} else {
					rec.Verbosef(c.Named(), "running %d deferral\n", count)
				}
				for len(c.deferrals) > 0 {
					deferral := <-c.deferrals
					if deferral != nil {
						c.deferralWait.Add(1)
						go func() {
							defer func() {
								if r := recover(); r != nil {
									rec.Printf(c.Named(), "deferral error: %v\n", r)
								}
							}()

							deferral(c.deferralWait)
						}()
					}
				}
				c.deferralWait.Wait()
			}
			rec.Verbosef(c.Named(), "cortex shut down complete\n")
			c.shutdownWait.Done()
		}()

		last := time.Now()
		var expected time.Duration
		var adjustment time.Duration
		var frequency float64

	main:
		for c.Alive() {
			if c.Frequency <= 0 {
				// This is a 'free-spin' condition
				select {
				case <-c.mute:
					select {
					case <-c.shutdown:
						break main
					case <-c.impulse:
						// NOTE: Impulse requests should not break the muted condition
						c.mute <- nil
					case <-c.unmute:
						for len(c.mute) > 0 {
							<-c.mute
						}
						for len(c.unmute) > 0 {
							<-c.unmute
						}
					}
				default:
				}
			} else {
				// This is a 'timer-step' condition
				expected = last.Add(_hertzToDuration(c.Frequency)).Sub(time.Now().Add(adjustment))
				frequency = c.Frequency
				select {
				case <-c.shutdown:
					break main
				case <-c.impulse:
				case <-time.After(expected):
					observed := time.Since(last)
					adjustment = observed - expected

					// If the frequency changed between cycles, don't try to 'adjust' it =)
					if c.Frequency != frequency {
						adjustment = 0
					}
				case <-c.mute:
					select {
					case <-c.shutdown:
						break main
					case <-c.impulse:
						// NOTE: Impulse requests should not break the muted condition
						c.mute <- nil
					case <-c.unmute:
						for len(c.mute) > 0 {
							<-c.mute
						}
						for len(c.unmute) > 0 {
							<-c.unmute
						}
					}
				}
			}

			for len(c.synapses) > 0 {
				imp := &Impulse{
					Cortex:   c,
					Timeline: NewTimeline(),
				}
				syn := <-c.synapses
				syn(imp)
			}

			c.clock.Broadcast()
			c.addToTimeline(time.Now())
			last = time.Now()
		}

		rec.Verbosef(c.Named(), "decayed\n")

		// This beat frees the synapses to complete their activation and exit
		c.clock.Broadcast()
	}()
}

func (c *Cortex) Shutdown(delay ...time.Duration) {
	c.sanityCheck()

	if !c.alive {
		return
	}

	if len(delay) > 0 {
		rec.Verbosef(c.Named(), "cortex shutting down in %v\n", delay[0])
		time.Sleep(delay[0])
	}

	c.master.Lock()
	rec.Verbosef(c.Named(), "cortex shutting down\n")
	c.running = false
	c.alive = false
	c.shutdown <- nil
	c.master.Unlock()
}

func (c *Cortex) addToTimeline(moment time.Time) {
	c.timeLock.Lock()
	defer c.timeLock.Unlock()

	c.timeline = append(c.timeline, moment)

	var trim int
	for i := range c.timeline {
		if c.timeline[i].Before(moment.Add(-atlas.ObservanceWindow)) {
			trim++
		} else {
			break
		}
	}
	c.timeline = c.timeline[trim:]
}

// TimelineOld returns the activation moments available to the cortex - limited to atlas.ObservanceWindow.
func (c *Cortex) Timeline() []time.Time {
	c.timeLock.Lock()
	defer c.timeLock.Unlock()
	return c.timeline
}

func (c *Cortex) Alive() bool {
	c.sanityCheck()

	return c.alive
}

func (c *Cortex) Mute() {
	c.sanityCheck()

	c.mute <- c.Entity
}

func (c *Cortex) Unmute() {
	c.sanityCheck()

	c.unmute <- c.Entity
}

func (c *Cortex) Inception() time.Time {
	c.sanityCheck()

	return c.inception
}

func (c *Cortex) Synapses() chan<- Synapse {
	c.sanityCheck()

	return c.synapses
}

func (c *Cortex) Deferrals() chan<- func(*sync.WaitGroup) {
	c.sanityCheck()

	return c.deferrals
}

func (c *Cortex) sanityCheck() {
	if !c.created {
		panic("cortices must be created through NewCortex")
	}
	if c.deferrals == nil {
		panic("deferrals must not be nil")
	}
	if c.synapses == nil {
		panic("synapses must not be nil")
	}
}
