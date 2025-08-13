package pattern2D

import (
	"github.com/ignite-laboratories/core/std"
)

// UpTo represents patterns up to 3 dimensions wide.
type UpTo[T any] interface {
	std.Pattern[T] | std.Pattern2D[T] | std.Pattern3D[T]
}
