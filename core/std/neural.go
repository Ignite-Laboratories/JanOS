package std

import "sync"

// Neural represents any type that can fire an action-potential (see.ActionPotentials).
type Neural interface {
	// Named returns the underlying Entity.Name.
	Named() string

	// Action takes in an Impulse and should return whether to keep the neural activity running (true).
	//
	// NOTE: The returned value is ignored in lifecycle.Impulse and lifecycle.Triggered activations.
	Action(*Impulse)

	// Potential takes in an Impulse and should return whether to fire the Action or not.
	Potential(*Impulse) bool

	// Cleanup will be called once the synaptic lifecycle completes.  This will fire regardless of if
	// the underlying action fires.
	Cleanup(*Impulse, *sync.WaitGroup)
}
