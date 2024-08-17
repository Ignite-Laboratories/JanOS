package JanOS

type named interface {
	// GetNamedValue returns the assigned name to this instance.
	GetNamedValue() string
}

type world interface {
	named
	Start()
}

type initializable interface {
	Initialize()
}

// Observer represents anything that can take an observed slice of time from a signal.
type Observer interface {
	OnObservation(*Signal, TimeSlice)
}
