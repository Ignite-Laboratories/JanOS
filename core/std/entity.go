package std

import (
	"github.com/ignite-laboratories/core/std/name"
	"github.com/ignite-laboratories/core/sys/id"
)

// Entity provides an 'ID' field to any composite types.
type Entity struct {
	// ID is the unique identifier for this entity, relative to its home world.
	ID        uint64
	GivenName name.Given
}

// NewEntity creates a new entity, assigns it a unique identifier, and gives it a random name.
//
// If you'd prefer to directly name your entity, provide it as a parameter here.  Otherwise,
// a random entry from the provided name.Format database type is chosen.
func NewEntity[T name.Format](str ...name.Given) Entity {
	var given name.Given
	if len(str) > 0 {
		given = str[0]
	} else {
		given = name.Random[T]()
	}

	ne := Entity{
		ID:        id.Next(),
		GivenName: given,
	}

	return ne
}
