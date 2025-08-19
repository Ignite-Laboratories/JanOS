package recorded

// Experience is a kind of unique set which will reset its values when an arbitrary condition is met.  By default, this
// will reset after a 'refractory period' of time has passed without requesting any numbers from the underlying Unique set.
// However, you can provide your own anonymous method to this type which will contextually decide when it's appropriate to
// reset the experiential data.
//
// Abstract -
// If an entity is presented with a potential experience generated through entropically generated random numbers, then the
// experience could be recreated using the same random number set at a later time.  The entity would not have awareness of
// how the experience -could- be, only how it occurred during -that- experience.  The orchestrator (you) can determine
// if there might be a reason to re-introduce the same kind of experience at a later time.  This kind of hippocampic
// paradigm is likely similar to how memories are retrieved in the mind while providing a means of reinforcing desired
// circumstances across time.
type Experience[T any] struct {
}
