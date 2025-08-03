package atlas

import (
	"time"
)

type config struct {
	PrintPreamble    *bool   `json:"printPreamble"`
	ObservanceWindow string  `json:"observanceWindow"`
	TrimFrequency    float64 `json:"trimFrequency"`
}

func (c config) apply() {
	if c.PrintPreamble != nil {
		PrintPreamble = *c.PrintPreamble
	}
	if len(c.ObservanceWindow) > 0 {
		ObservanceWindow, _ = time.ParseDuration(c.ObservanceWindow)
	}
	if c.TrimFrequency > 0 {
		TrimFrequency = c.TrimFrequency
	}
}
