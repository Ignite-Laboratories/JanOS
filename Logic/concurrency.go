package Logic

import (
	"sync/atomic"
)

var masterCount uint64

// NextId increments the internal master count maintained since execution and then returns the value
func NextId() uint64 { return atomic.AddUint64(&masterCount, 1) }
