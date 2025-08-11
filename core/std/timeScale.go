package std

import (
	"github.com/ignite-laboratories/core/std/num"
	"time"
)

// TimeScale represents a pairing of duration and an abstract "height."
type TimeScale[T num.Primitive] struct {
	Duration time.Duration
	Height   T
}
