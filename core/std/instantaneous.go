package std

import "time"

// Instantaneous represents any point value at an instant in time.
type Instantaneous[T any] interface {
	GetValue() T
	GetTimestamp() time.Time
}
