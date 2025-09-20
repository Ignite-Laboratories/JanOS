package std

import (
	"sync"
	"time"

	"git.ignitelabs.net/core"
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

	synapses chan synapse

	deferrals chan func()

	locked  bool
	alive   bool
	created bool
	running bool
	relay   chan any
	reset   chan any
	limit   int

	master sync.Mutex
	clock  sync.Cond
}

// NewCortex creates a new named Cortex limited to the provided number of neural activations.
//
// The named parameter takes in either a string or a name format - if you'd like a random name, please use format.Default
//
// NOTE: If no limit is provided, the default is 2¹⁶ - this can generally be ignored for most systems.
func NewCortex(named string, synapticLimit ...int) *Cortex {
	limit := 2 ^ 16
	if len(synapticLimit) > 0 {
		limit = synapticLimit[0]
	}

	c := &Cortex{
		Entity:    NewEntity[format.Default](),
		inception: time.Now(),
		synapses:  make(chan synapse, limit),
		deferrals: make(chan func(), limit),
		alive:     true,
		created:   true,
		limit:     limit,
	}
	c.clock = sync.Cond{L: &c.master}
	c.Entity.Name.Name = named
	core.Deferrals() <- func() {
		c.Shutdown()
	}

	log.Printf(core.ModuleName, "created cortex '%s'\n", c.Named())
	return c
}

func (c *Cortex) Spark(synapses ...synapse) {
	c.sanityCheck()

	for _, syn := range synapses {
		c.synapses <- syn
	}

	if c.running {
		return
	}
	c.alive = true
	c.running = true

	log.Printf(c.Named(), "sparking\n")

	defer func() {
		for deferral := range c.deferrals {
			deferral()
		}
	}()

	last := time.Now()

	started := false

	for c.Alive() {
		c.master.Lock()
		c.master.Unlock()
		moment := time.Now()
		ctx := Context{
			Moment: moment,
			Cortex: c,
		}

		if !started || (c.Frequency <= 0 || time.Since(last) > when.HertzToDuration(c.Frequency)) {
			started = true
			last = moment
			for len(c.synapses) > 0 {
				syn := <-c.synapses
				syn(ctx)
			}

			c.clock.Broadcast()
		}
	}
}

func (c *Cortex) Shutdown(delay ...time.Duration) {
	c.sanityCheck()

	if len(delay) > 0 {
		log.Verbosef(c.Named(), "shutting down in %v\n", delay[0])
		time.Sleep(delay[0])
	}

	c.master.Lock()
	log.Printf(c.Named(), "shutting down\n")
	c.running = false
	c.alive = false

	log.Verbosef(c.Named(), "running %d deferrals\n", len(c.deferrals))

	wg := sync.WaitGroup{}
	for deferFn := range c.deferrals {
		wg.Add(1)
		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Printf(c.Named(), "deferral error: %v\n", r)
					wg.Done()
				}
			}()

			deferFn()
			wg.Done()
		}()
	}
	wg.Wait()

	log.Verbosef(c.Named(), "shut down\n")

	c.master.Unlock()
}

func (c *Cortex) Alive() bool {
	c.sanityCheck()

	return c.alive
}

func (c *Cortex) Mute() {
	c.sanityCheck()

	c.master.Lock()
}

func (c *Cortex) Unmute() {
	c.sanityCheck()

	c.master.Unlock()
}

func (c *Cortex) Inception() time.Time {
	c.sanityCheck()

	return c.inception
}

func (c *Cortex) Synapses() chan<- synapse {
	c.sanityCheck()

	return c.synapses
}

func (c *Cortex) Deferrals() chan<- func() {
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
