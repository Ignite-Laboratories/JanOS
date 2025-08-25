package pattern

import "slices"

// Fixed creates a fixed pattern function that cyclically steps through the statically provided elements.
func Fixed[T any](toPad ...T) func() T {
	if len(toPad) == 0 {
		panic("cannot pad without data to pad with")
	}
	out := slices.Clone(toPad)
	i := 0
	return func() T {
		element := out[i]
		i++
		if i >= len(out) {
			i = 0
		}
		return element
	}
}
