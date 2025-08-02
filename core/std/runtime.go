package std

import "time"

// Runtime provides information about how long something took to run relative to its inception.
type Runtime struct {
	// RefractoryPeriod is the amount of time between the end of the last and current cycle.
	//
	// For impulse statistics, this is the period between the end of the last impulse's
	// blocking execution and the current impulse's inception.
	//
	// For activation statistics, this is the period between the end of the last activation
	// and the start of the current activation.
	RefractoryPeriod time.Duration

	// Inception is the moment the impulse started.
	Inception time.Time

	// Start is the moment of activation.
	//
	// For impulse statistics, this is always the same as Inception.
	Start time.Time

	// End is the moment activation completed.
	//
	// For impulse statistics, this is the moment all blocking execution completed.
	End time.Time

	// Duration is the period between Start and End.
	Duration time.Duration
}
