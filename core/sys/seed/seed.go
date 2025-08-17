// Package seed provides access to the seed system.  This is a set of semantic ways to create transiently 'close enough'
// experiences that still evolve over time.  The idea is simple - you can create seed caches of random numbers which can be
// queried.  These caches will retain their value set for as long as you keep requesting from it, but once you have let their
// defined 'refractory period' elapse, the next call will re-seed the seed cache with new numbers.
package seed

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/num"
	"github.com/ignite-laboratories/core/sys/atlas"
	"sync"
	"time"
)

type entry[T num.Primitive] struct {
	sync.Mutex
	seed     uint64
	seeds    std.UniqueSet[T]
	promote  func() bool
	last     time.Time
	readOnly bool
}

var cache = make(map[uint64]*any)

// SetPromotion sets the seed's "promotion" function, which should indicate when the seed should promote to a
// read-only set of values which never changes.  The logic for this action is entirely contextual, thus you
// will need to implement your own rules for hippocampic memory.
func SetPromotion[T num.Primitive](seed uint64, promotion func() bool) {
	var c *entry[T]
	if v, ok := cache[seed]; ok {
		c = any(v).(*entry[T])
	} else {
		c = &entry[T]{
			seed:  seed,
			last:  time.Now(),
			seeds: std.UniqueSet[T]{},
		}
		cache[seed] = any(c).(*any)
	}
	c.Lock()
	defer c.Unlock()
	c.promote = promotion
}

// Reset clears the read-only status and resets the seed to a fresh entry.
func Reset[T num.Primitive](seed uint64) {
	var c *entry[T]
	if v, ok := cache[seed]; ok {
		c = v
	} else {
		c = &entry[T]{
			seed:  seed,
			last:  time.Now(),
			seeds: make([]any, 0, 0),
		}
		cache[seed] = c
	}
	c.Lock()
	defer c.Unlock()
	c.readOnly = false
	c.seeds = make([]any, 0, 0)
	c.last = time.Now()
}

// Random returns the requested number of random numbers from a seed which contains a cache of random values.  If the
// requested cache size changes between calls, each call will automatically truncate or fill in random numbers as it
// adjusts the cache size.  The cache will retain its current values as long as you continue to call the function
// within the refractory period, otherwise the next call will regenerate the seed cache.  Seeds also can be 'promoted'
// to 'read-only' using an anonymous method, meaning the values will then remain permanently static - see SetPromotion.
//
// NOTE: If you provide the wrong type for the provided seed, this will intentionally panic.  Please be wary of your types.
//
// NOTE: If you do not provide a refractory period, the system-wide atlas.SeedRefractoryPeriod is used.
func Random[T num.Primitive](seed uint64, count uint, cacheSize uint, refractoryPeriod ...time.Duration) []T {
	period := atlas.SeedRefractoryPeriod
	if len(refractoryPeriod) > 0 {
		period = refractoryPeriod[0]
	}

	var c *entry[T]
	if v, ok := cache[seed]; ok {
		c = v
	} else {
		c = &entry[T]{
			seed:  seed,
			last:  time.Now(),
			seeds: make([]any, 0, cacheSize),
		}
		cache[seed] = c
	}
	c.Lock()
	defer c.Unlock()

	if !c.readOnly && c.promote != nil && c.promote() {
		c.readOnly = true
	}

	if !c.readOnly {
		if time.Since(c.last) > period {
			c.seeds = make([]any, 0, cacheSize)
		}

		seen := make(map[T]struct{}, n)
		if len(c.seeds) < int(cacheSize) {
			for i := uint(0); i < cacheSize-uint(len(c.seeds)); i++ {
				c.seeds = append(c.seeds, num.Random[T]())
			}
		} else if len(c.seeds) > int(cacheSize) {
			c.seeds = c.seeds[:int(cacheSize)]
		}
	}
	c.last = time.Now()

	out := make([]T, count)
	found := make(map[T]struct{})
	var i uint
	for i < count {
		v := c.seeds[num.RandomWithinRange[uint](0, uint(len(c.seeds)-1))].(T)
		if _, ok := found[v]; !ok || i >= uint(len(c.seeds)) {
			found[v] = struct{}{}
			out[i] = v
			i++
		}
	}
	return out
}
