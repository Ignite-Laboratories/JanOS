package std

// Renderable represents the lifecycle of an impulsably renderable structure.
//
// It's functionally equivalent to Impulsable, but the Impulse method also accepts
// a generic 'surface' and size.
type Renderable[T any] interface {
	Initialize()
	Impulse(ctx Context, surface T, size XY[int])
	Cleanup()
	Lock()
	Unlock()
}
