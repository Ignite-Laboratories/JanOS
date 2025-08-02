package std

// BooleanReader represents a type that gets a ReadOnlyBool value.
type BooleanReader interface {
	Get() bool
}

// ReadOnlyBool represents a boolean that can only be set by its creator.
type ReadOnlyBool struct {
	value *bool
}

// NewReadOnlyBool creates a read-only boolean value tied to the provided boolean reference.
func NewReadOnlyBool(value *bool) *ReadOnlyBool {
	return &ReadOnlyBool{value: value}
}

// Get de-references the current value of the read-only boolean.
func (b *ReadOnlyBool) Get() bool {
	return *b.value
}
