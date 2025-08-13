package pattern

import (
	"github.com/ignite-laboratories/core/std"
)

// Buffer represents any of the dimensional pattern types, each representing infinitely repeating patterns of
// ordinal dimensional data which can be traversed in any axis.
//
// NOTE: For pattern generation and predefined patterns, see the 'std/bounded/pattern' package.
//
// See Buffer, UpTo2D, UpTo3D, UpTo4D, UpTo5D, UpTo6D, and UpTo7D
type Buffer[T any] interface {
	std.Pattern[T] | std.Pattern2D[T] | std.Pattern3D[T] | std.Pattern4D[T] | std.Pattern5D[T] | std.Pattern6D[T] | std.Pattern7D[T]
}

// UpTo2D represents patterns up to 2 dimensions wide.
//
// See Buffer, UpTo2D, UpTo3D, UpTo4D, UpTo5D, UpTo6D, and UpTo7D
type UpTo2D[T any] interface {
	std.Pattern[T] | std.Pattern2D[T]
}

// UpTo3D represents patterns up to 3 dimensions wide.
//
// See Buffer, UpTo2D, UpTo3D, UpTo4D, UpTo5D, UpTo6D, and UpTo7D
type UpTo3D[T any] interface {
	std.Pattern[T] | std.Pattern2D[T] | std.Pattern3D[T]
}

// UpTo4D represents patterns up to 4 dimensions wide.
//
// See Buffer, UpTo2D, UpTo3D, UpTo4D, UpTo5D, UpTo6D, and UpTo7D
type UpTo4D[T any] interface {
	std.Pattern[T] | std.Pattern2D[T] | std.Pattern3D[T] | std.Pattern4D[T]
}

// UpTo5D represents patterns up to 5 dimensions wide.
//
// See Buffer, UpTo2D, UpTo3D, UpTo4D, UpTo5D, UpTo6D, and UpTo7D
type UpTo5D[T any] interface {
	std.Pattern[T] | std.Pattern2D[T] | std.Pattern3D[T] | std.Pattern4D[T] | std.Pattern5D[T]
}

// UpTo6D represents patterns up to 6 dimensions wide.
//
// See Buffer, UpTo2D, UpTo3D, UpTo4D, UpTo5D, UpTo6D, and UpTo7D
type UpTo6D[T any] interface {
	std.Pattern[T] | std.Pattern2D[T] | std.Pattern3D[T] | std.Pattern4D[T] | std.Pattern5D[T] | std.Pattern6D[T]
}

// UpTo7D represents patterns up to 7 dimensions wide.
//
// See Buffer, UpTo2D, UpTo3D, UpTo4D, UpTo5D, UpTo6D, and UpTo7D
type UpTo7D[T any] interface {
	std.Pattern[T] | std.Pattern2D[T] | std.Pattern3D[T] | std.Pattern4D[T] | std.Pattern5D[T] | std.Pattern6D[T] | std.Pattern7D[T]
}
