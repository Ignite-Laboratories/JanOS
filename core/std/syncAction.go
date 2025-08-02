package std

import "sync"

// SyncAction represents a wait-able action.
type SyncAction struct {
	sync.WaitGroup
	Action func()
}
