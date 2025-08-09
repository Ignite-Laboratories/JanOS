package tiny

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/num"
)

// Operable represents the basic logically operable types.
//
// See Bit, Measurement, Phrase, Natural, Real, Complex, and Index
type Operable interface {
	num.Bit | byte | std.Measurement[any] | std.Phrase[any] | std.Natural | Real | Complex | Index
}
