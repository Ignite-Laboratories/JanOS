package std

import (
	"core/sys/id"
	"core/sys/name"
	"core/sys/name/format"
	"fmt"
)

// Entity provides an 'ID' field to any composite types.
//
// NOTE: Entity's String function.
type Entity struct {
	id        uint64
	GivenName name.Given
}

func (e Entity) GetID() uint64 {
	return e.id
}

// String returns the Entity's identifier and GivenName as "[ID](Name)"
func (e Entity) String() string {
	return fmt.Sprintf("[%d](%v)", e.id, e.GivenName.String())
}

// NewEntity creates a new entity, assigns it a unique identifier, and gives it a random name.
//
// If you'd prefer to directly name your entity, provide it as a parameter here.  Otherwise,
// a random entry from the provided name.Format database type is chosen.
func NewEntity[T format.Format](str ...name.Given) Entity {
	i := id.Next()
	var given name.Given
	if len(str) > 0 {
		given = str[0]
	} else {
		given, _ = name.Random[T](i)
	}

	ne := Entity{
		id:        i,
		GivenName: given,
	}

	return ne
}
