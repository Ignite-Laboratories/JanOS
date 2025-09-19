package std

import (
	"time"
)

type Context struct {
	ModuleName string

	LastActivation   time.Time
	Moment           time.Time
	RefractoryPeriod time.Duration
	ResponseTime     time.Duration
	ActivationTime   time.Duration
	Beat             int

	Cortex *Cortex
}
