package std

// Entity provides an 'ID' field to any composite types.
type Entity struct {
	// ID is the unique identifier for this entity, relative to its home world.
	ID uint64
}
