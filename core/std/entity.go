package std

import (
	"github.com/ignite-laboratories/core/std/name"
	"github.com/ignite-laboratories/core/sys/id"
)

// Entity provides an 'ID' field to any composite types.
type Entity struct {
	// ID is the unique identifier for this entity, relative to its home world.
	ID uint64
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
		given = name.Random()
	}

	ne := NamedEntity{
		Given: given,
	}
	ne.ID = id.Next()

	return ne
}

// NewNamedEntityFromDB creates a new entity, assigns it a unique identifier, and gives it a random
// name from the provided name database.  If no database is provided, the default database is used.
//
// If you'd prefer to name your entity directly, please see NewNamed.
func NewNamedEntityFromDB(db ...[]name.Given) NamedEntity {
	given := name.Random(db...)

	ne := NamedEntity{
		Given: given,
	}
	ne.ID = id.Next()

	return ne
}