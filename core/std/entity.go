package std

import (
	"github.com/ignite-laboratories/core/std/name"
	"github.com/ignite-laboratories/core/std/name/format"
	"github.com/ignite-laboratories/core/sys/id"
)

// Entity provides an 'ID' field to any composite types.
type Entity struct {
	id        uint64
	GivenName name.Given
}

func (e Entity) GetID() uint64 {
	return e.id
}

func (e Entity) String() string {
	return e.GivenName.String()
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
