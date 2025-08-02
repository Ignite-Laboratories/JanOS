package temporal

// Recorder is a type of core.Emitter that records values onto a dimension.
type Recorder[TValue any] struct {
	d *Dimension[TValue, any]
}

// NewRecorder creates a recorder that emits typed values onto the provided dimension.
func NewRecorder[TValue any](dimension *Dimension[TValue, any]) *Recorder[TValue] {
	r := Recorder[TValue]{}
	r.d = dimension
	return &r
}

// Emit records the provided value onto the dimensional timeline.
//
// NOTE: This timestamps the emission with the current instant.
// The ability to doctor -when- something is emitted directly contradicts the spirit of this library,
// and as such you should reconsider your implementation if you wish to circumvent this limitation.
//
// Remember, the amount of calculation time something takes to emit a value is a useful data point in of itself! =)
func (r *Recorder[TValue]) Emit(value TValue) {
	r.d.Write(value)
}
