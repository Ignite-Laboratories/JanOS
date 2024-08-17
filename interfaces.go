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

// Sampler represents anything that can take an observed slice of time from a signal.
type Sampler interface {
	named
	OnSample(*Signal, TimeSlice)
}

type Observer interface {
	named
	OnTrigger(Observation)
}

type Observation struct {
	Observer named
	Signal   *Signal
	Values   []InstantaneousValue
}
