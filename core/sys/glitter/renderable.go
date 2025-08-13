package glitter

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/bounded"
)

// Renderable represents the lifecycle of an impulsably renderable structure.
//
// It's functionally equivalent to Impulsable, but the Impulse method also accepts
// a generic 'surface' and size.
type Renderable[T any] interface {
	Initialize()
	Impulse(ctx std.Context, surface T, size bounded.XY[int])
	Cleanup()
	Lock()
	Unlock()
}
