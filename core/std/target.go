package std

// TargetFunc functions return a pointer to a value.
type TargetFunc[TValue any] func() *TValue

// Target returns a function that retrieves a reference to the target on demand.
func Target[TValue any](val *TValue) TargetFunc[TValue] {
	return func() *TValue {
		return val
	}
}
