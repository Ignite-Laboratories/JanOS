package std

import (
	"sync"
	"time"

	"git.ignitelabs.net/core"
	"git.ignitelabs.net/core/sys/atlas"
	"git.ignitelabs.net/core/sys/given/format"
	"git.ignitelabs.net/core/sys/log"
	"git.ignitelabs.net/core/sys/when"
)

type Cortex struct {
	Entity

	// Frequency defines the minimum frequency impulses will fire.
	//
	// NOTE: If you set this negative, they will fire as fast as possible.
	Frequency float64

	inception time.Time

	synapses chan Synapse
	timeline []time.Time

	deferrals    chan func(*sync.WaitGroup)
	deferralWait *sync.WaitGroup

	mute     chan any
	unmute   chan any
	shutdown chan any

	alive   bool
	created bool
	running bool

	timeLock sync.Mutex
	master   sync.Mutex
	clock    sync.Cond
}

// NewCortex creates a new named Cortex limited to the provided number of neural activations.
//
// The named parameter takes in either a string or a name format - if you'd like a random name, please use format.Default
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
		shutdown:     make(chan any, 1<<16),
		alive:        true,
		created:      true,
	}
	c.deferralWait.Add(1)
	c.clock = sync.Cond{L: &c.master}
	c.Entity.Name.Name = named
	core.Deferrals() <- func(wg *sync.WaitGroup) {
		c.Shutdown()
		c.deferralWait.Wait()
		log.Verbosef(c.Named(), "cortex shut down complete\n")
		wg.Done()
	}

	log.Printf(core.ModuleName, "created cortex '%s'\n", c.Named())
	return c
}

func (c *Cortex) Spark(synapses ...Synapse) {
	c.sanityCheck()

	for _, syn := range synapses {
		c.synapses <- syn
	}

	if c.running {
		return
	}
	go func() {
		c.alive = true
		c.running = true

		log.Printf(c.Named(), "sparking\n")

		defer func() {
			if len(c.deferrals) > 0 {
				log.Verbosef(c.Named(), "running %d deferrals\n", len(c.deferrals))
				for len(c.deferrals) > 0 {
					deferral := <-c.deferrals
					if deferral != nil {
						go func() {
							defer func() {
								if r := recover(); r != nil {
									log.Printf(c.Named(), "deferral error: %v\n", r)
									c.deferralWait.Done()
								}
							}()

							deferral(c.deferralWait)
						}()
					}
				}
			} else {
				c.deferralWait.Done()
			}
		}()

		last := time.Now()
		var expected time.Duration
		var adjustment time.Duration
		var frequency float64
		started := false

	main:
		for c.Alive() {
			if started || len(c.mute) <= 0 {
				if c.Frequency <= 0 {
					// This is a 'free-spin' condition
					select {
					case <-c.mute:
						_ = <-c.unmute

						for _ = range c.mute {
						}
						for _ = range c.unmute {
						}
					default:
					}
				} else {
					// This is a 'timer-step' condition
					expected = last.Add(when.HertzToDuration(c.Frequency)).Sub(time.Now().Add(adjustment))
					frequency = c.Frequency
					select {
					case <-c.shutdown:
						break main
					case <-time.After(expected):
						observed := time.Since(last)
						adjustment = observed - expected

						// If the frequency changed between cycles, don't try to 'adjust' it =)
						if c.Frequency != frequency {
							adjustment = 0
						}
					case <-c.mute:
						_ = <-c.unmute
					}
				}
			}
			started = true

			for len(c.synapses) > 0 {
				imp := &Impulse{
					Cortex:   c,
					Timeline: newTimeline(),
				}
				syn := <-c.synapses
				syn(imp)
			}

			c.clock.Broadcast()
			c.addToTimeline(time.Now())
			last = time.Now()
		}
	}()
}

// Keola Frenzel;
func (c *Cortex) Shutdown(delay ...time.Duration) {
	c.sanityCheck()

	if !c.alive {
		return
	}

	if len(delay) > 0 {
		log.Verbosef(c.Named(), "cortex shutting down in %v\n", delay[0])
		time.Sleep(delay[0])
	}

	c.master.Lock()
	c.running = false
	c.alive = false
	c.shutdown <- nil

	log.Verbosef(c.Named(), "cortex shutting down\n")

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

// Timeline returns the activation moments available to the cortex - limited to atlas.ObservanceWindow.
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
