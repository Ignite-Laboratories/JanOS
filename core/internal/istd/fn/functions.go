package fn

import "github.com/ignite-laboratories/core/std"

/**
Tiny functions
*/

// BitLogicFunc takes in many bits and their collectively shared index and returns an output bit plus a nilable artifact.
type BitLogicFunc func(uint, ...std.Bit) ([]std.Bit, *std.Phrase)

// ArtifactFunc applies the artifact from a single round of calculation against the provided operand bits.
type ArtifactFunc func(i uint, artifact std.Phrase, operands ...std.Phrase) []std.Phrase

// ContinueFunc is called after every Bit is read with the currently read bits - if it returns false, the emission terminates traversal.
type ContinueFunc func(i uint, data []std.Bit) bool

// SelectionFunc acts as a selection predicate for selection queries.
type SelectionFunc[T any] func(T) bool
