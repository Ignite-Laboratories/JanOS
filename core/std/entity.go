package std

import (
	"fmt"
	"time"

	"git.ignitelabs.net/janos/core/sys/given"
	"git.ignitelabs.net/janos/core/sys/given/format"
	"git.ignitelabs.net/janos/core/sys/id"
)

// Entity provides an 'ID' field to any composite types.
//
// NOTE: Entity's String function.
type Entity struct {
	id uint64
	given.Given

	Genesis time.Time
}

func (e Entity) GetID() uint64 {
	return e.id
}

// Named gets the Given name.
func (e Entity) Named() string {
	return e.Given.Name
}

// String returns the Entity's identifier and Given as "[ID](Name)"
func (e Entity) String() string {
	return fmt.Sprintf("[%d](%v)", e.id, e.Given.String())
}

// NewEntityNamed creates a new entity, assigns it a unique identifier, and gives it the provided name.
//
// See NewEntity and NewEntityNamed
func NewEntityNamed(name string) Entity {
	return NewEntity[format.Default](given.New(name))
}

// NewEntity creates a new entity, assigns it a unique identifier, gives it a random name, and sets its genesis moment.
//
// See NewEntity and NewEntityNamed
func NewEntity[T format.Format](name ...given.Given) Entity {
	i := id.Next()
	var g given.Given
	if len(name) > 0 {
		g = name[0]
	} else {
		g = given.Random[T]()
	}

	ne := Entity{
		id:      i,
		Given:   g,
		Genesis: time.Now(),
	}

	return ne
}
