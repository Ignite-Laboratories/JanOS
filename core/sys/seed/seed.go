// Package seed provides access to the seed system.  This is a set of semantic ways to create stable 'experiences' which
// still evolve over time.  The idea is simple - you can create seed caches of random numbers which can be queried.  These
// caches will retain their value set for as long as you keep requesting from it, but once you have let their defined
// 'refractory period' elapse, the next call will re-seed the seed cache with new numbers.
package seed

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/bounded"
	"github.com/ignite-laboratories/core/std/num"
	"github.com/ignite-laboratories/core/std/set"
	"github.com/ignite-laboratories/core/sys/atlas"
	"sync"
	"time"
)

type entry[T num.Primitive] struct {
	sync.Mutex
	seed     uint64
	seeds    set.Unique[T]
	promote  func() bool
	bounds   std.Bounded[T]
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
			seeds: set.Numeric[T](),
		}
		cache[seed] = any(c).(*any)
	}
	c.Lock()
	defer c.Unlock()
	c.promote = promotion
}

func SetBounds[T num.Primitive](seed uint64, bounds std.Bounded[T]) {
	var c *entry[T]
	if v, ok := cache[seed]; ok {
		c = any(v).(*entry[T])
	} else {
		c = &entry[T]{
			seed:  seed,
			last:  time.Now(),
			seeds: set.Unique[T]{},
		}
		cache[seed] = any(c).(*any)
	}
	c.Lock()
	defer c.Unlock()
	c.promote = promotion
}

// Reset clears the read-only status and resets the seed to a fresh entry.
func Reset[T num.Primitive](seed uint64) {
	b, _ := bounded.ByType[T](0)
	var c *entry[T]
	if v, ok := cache[seed]; ok {
		c = any(v).(*entry[T])
	} else {
		c = &entry[T]{
			seed: seed,
			last: time.Now(),
		}
		cache[seed] = any(c).(*any)
	}
	c.Lock()
	defer c.Unlock()
	c.readOnly = false
	c.seeds = *set.Numeric[T](b)
	c.last = time.Now()
}

// Random returns the requested number of random numbers from a std.Unique[T] while modulating it to the desired cacheSize.
// The provided seed value is associated with a single unique set, and the cache will retain its current values as long as you
// continue to call the function within the refractory period.  Otherwise the next call for that seed value will reset the unique
// set's data.  Seeds also can be 'promoted' to 'read-only' using an anonymous method, meaning the values will then remain permanently
// static - see SetPromotion.
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
		c = any(v).(*entry[T])
	} else {
		c = &entry[T]{
			seed:  seed,
			last:  time.Now(),
			seeds: *set.Numeric[T](bounded),
		}
		a := any(c)
		cache[seed] = &a
	}
	c.Lock()
	defer c.Unlock()

	if !c.readOnly && c.promote != nil && c.promote() {
		c.readOnly = true
	}

	if !c.readOnly {
		if time.Since(c.last) > period {
			c.seeds.Reset()
		}

		c.seeds.Modulate(cacheSize)
	}
	c.last = time.Now()

	return c.seeds.Random(count)
}
