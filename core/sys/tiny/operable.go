package tiny

import "github.com/ignite-laboratories/core/std"

// Operable represents the basic logically operable types.
//
// See Bit, Measurement, Phrase, Natural, Real, Complex, and Index
type Operable interface {
	std.Bit | byte | std.Measurement[any] | std.Phrase[any] | Natural | Real | Complex | Index
}
