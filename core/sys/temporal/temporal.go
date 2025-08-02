package temporal

import "github.com/ignite-laboratories/core"

// Prime is the locally scoped universe.
var Prime = make(Universe)

// Universe is a collection of named worlds.
type Universe core.FilterableMap[string, World]

// World is a collection of named dimensions and any associated entities.
//
// Dimensions are stored by name.
//
// Entities are stored by ID.
type World struct {
	Dimensions core.FilterableMap[string, core.FilterableSlice[*Dimension[any, any]]]
	Entities   core.FilterableMap[uint64, *core.Entity]
}
