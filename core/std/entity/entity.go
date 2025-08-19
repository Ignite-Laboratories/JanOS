package entity

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/name"
	"github.com/ignite-laboratories/core/std/name/format"
)

// New creates a new entity, assigns it a unique identifier, and gives it a random name.
//
// If you'd prefer to directly name your entity, provide it as a parameter here.  Otherwise,
// a random entry from the provided name.Format database type is chosen.
func New[T format.Format](str ...name.Given) std.Entity {
	return std.NewEntity[T](str...)
}
