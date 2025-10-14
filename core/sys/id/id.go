package id

import (
	"sync"

	"git.ignitelabs.net/janos/core/sys/num"
)

var emitted = make(map[uint64]struct{})
var gate = &sync.Mutex{}

// Next provides a thread-safe periodically unique identifier to every caller.
//
// A "periodically unique" value is one that's guaranteed never to repeat until the available set of numbers is exhausted.
func Next() uint64 {
	gate.Lock()
	defer gate.Unlock()

	// NOTE: This doesn't care about handling exhaustion - we'll literally never see that scenario in our lifetimes =)

	r := num.Random[uint64]()
	for _, exist := emitted[r]; exist; _, exist = emitted[r] {
		r = num.Random[uint64]()
	}
	emitted[r] = struct{}{}
	return r
}
