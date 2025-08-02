package std

import (
	"github.com/ignite-laboratories/core/math"
	"time"
)

// TimeScale represents a pairing of duration and an abstract "height."
type TimeScale[T math.Numeric] struct {
	Duration time.Duration
	Height   T
}
