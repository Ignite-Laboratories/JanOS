package id

import "sync/atomic"

// current holds the last provided identifier.
var current uint64

// Next provides a thread-safe unique identifier to every caller.
func Next() uint64 {
	return atomic.AddUint64(&current, 1)
}
