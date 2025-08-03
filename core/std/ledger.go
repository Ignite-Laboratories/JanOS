package std

import "reflect"

// Ledger represents a deep measurement of an object.
type Ledger[T any] struct {
	Entries map[uintptr]LedgerEntry
	Start   uintptr
}

type LedgerEntry struct {
	Measurement[any]
	reflect.Type
	Address uintptr
}
