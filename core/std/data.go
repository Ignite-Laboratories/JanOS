package std

// Data represents a contextual point value in time.
type Data[T any] struct {
	Context

	// Point is the recorded value of this contextual moment.
	Point T
}
