package core

import "time"

type Context struct {
	LastActivation time.Time
	Now            time.Time
	Beat           int
}
