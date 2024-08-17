package JanOS

import "time"

type Formula struct {
	Operator  string
	Operation func(instant time.Time, sourceSignal *Signal, otherSignals ...*Signal) float64
}
