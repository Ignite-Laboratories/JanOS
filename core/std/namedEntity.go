package std

import (
	"github.com/ignite-laboratories/core/std/name"
	"github.com/ignite-laboratories/core/sys/id"
)

// NamedEntity provides a given name to the entity.
type NamedEntity struct {
	name.Given
	Entity
}

// NewNamedEntity creates a new entity, assigns it a unique identifier, and gives it a random name.
//
// If you'd prefer to directly name your entity, provide it as a parameter here.  Otherwise,
// a random entry from core.Names is chosen.  If you'd prefer to use a different random
// name database, please see NewNamedFromDB.
func NewNamedEntity(str ...name.Given) NamedEntity {
	var given name.Given
	if len(str) > 0 {
		given = str[0]
	} else {
		given = name.Random[name.NameDB]()
	}

	ne := NamedEntity{
		Given: given,
	}
	ne.ID = id.Next()

	return ne
}
