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
