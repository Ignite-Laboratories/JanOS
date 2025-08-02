package std

import "github.com/ignite-laboratories/core/std/name"

// NamedEntity provides a given name to the entity.
type NamedEntity struct {
	name.Given
	Entity
}
