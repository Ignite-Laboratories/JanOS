package std

import (
	"core/sys/given"
	"core/sys/given/format"
	"core/sys/id"
	"fmt"
)

// Entity provides an 'ID' field to any composite types.
//
// NOTE: Entity's String function.
type Entity struct {
	id   uint64
	Name given.Name
}

func (e Entity) GetID() uint64 {
	return e.id
}

// String returns the Entity's identifier and Name as "[ID](Name)"
func (e Entity) String() string {
	return fmt.Sprintf("[%d](%v)", e.id, e.Name.String())
}

// NewEntityNamed creates a new entity, assigns it a unique identifier, and gives it the provided name.
//
// See NewEntity and NewEntityNamed
func NewEntityNamed(name string) Entity {
	return NewEntity[format.Default](given.New(name))
}

// NewEntity creates a new entity, assigns it a unique identifier, and gives it a random name.
//
// See NewEntity and NewEntityNamed
func NewEntity[T format.Format](name ...given.Name) Entity {
	i := id.Next()
	var g given.Name
	if len(name) > 0 {
		g = name[0]
	} else {
		g, _ = given.Random[T](i)
	}

	ne := Entity{
		id:   i,
		Name: g,
	}

	return ne
}
