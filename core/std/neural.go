package std

// Neural represents any type that can fire an action-potential (see.ActionPotentials).
type Neural interface {
	// Named returns the underlying Entity.Given.
	Named() string

	// Action takes in an Impulse and should return whether to keep the neural activity running (true).
	//
	// NOTE: The returned value is ignored in life.Impulse and life.Triggered activations.
	Action(*Impulse)

	// Potential takes in an Impulse and should return whether to fire the Action or not.
	Potential(*Impulse) bool

	// Cleanup will be called once the synaptic life completes.  This will fire regardless of if
	// the underlying action fires.
	Cleanup(*Impulse)
}
