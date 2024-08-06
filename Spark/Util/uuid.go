package Util

import (
	"github.com/google/uuid"
	"math/rand"
)

// NewUUID returns a new UUID from the Google UUID library.
// This is merely for convenience.
func NewUUID() uuid.UUID {
	return uuid.New()
}

// NewUUIDSeeded returns a new UUID from the Google UUID library,
// but from a seeded integer value.  This allows a complex
// hash of a known identifier - likely to be used to add entropy.
func NewUUIDSeeded(seed int64) (uuid.UUID, error) {
	var id uuid.UUID
	r := rand.New(rand.NewSource(seed))
	randomizer := make([]byte, 16)
	_, err := r.Read(randomizer)
	if err != nil {
		return uuid.Nil, err
	}
	copy(id[:], randomizer[:16])
	id[6] = (id[6] & 0x0f) | 0x40 // Version 4
	id[8] = (id[8] & 0x3f) | 0x80 // Variant is 10
	return id, nil
}
