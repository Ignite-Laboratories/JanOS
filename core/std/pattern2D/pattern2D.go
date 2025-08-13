package pattern2D

import (
	"github.com/ignite-laboratories/core/std"
)

// UpTo represents patterns up to 2 dimensions wide.
type UpTo[T any] interface {
	std.Pattern[T] | std.Pattern2D[T]
}
