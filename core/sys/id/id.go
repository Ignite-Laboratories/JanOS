package id

import "sync/atomic"

var current atomic.Uint64

// Next provides a thread-safe unique identifier to every caller.
func Next() uint64 {
	return current.Add(1)
}
