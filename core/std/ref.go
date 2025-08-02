package std

// Ref provides a way to create hard inline pointer references.
func Ref[T any](val T) *T {
	return &val
}
