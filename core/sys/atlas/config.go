package atlas

import (
	"time"
)

type config struct {
	PrintPreamble        *bool   `json:"printPreamble"`
	ObservanceWindow     string  `json:"observanceWindow"`
	TrimFrequency        float64 `json:"trimFrequency"`
	Precision            uint    `json:"precision"`
	PrecisionMinimum     uint    `json:"precisionMinimum"`
	Base                 uint16  `json:"base"`
	SeedRefractoryPeriod string  `json:"seedRefractoryPeriod"`
	IncludeNilBits       *bool   `json:"includeNilBits"`
	CompactVectors       *bool   `json:"compactVectors"`
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
	if c.Precision > 0 {
		Precision = c.Precision
	}
	if c.PrecisionMinimum > 0 {
		PrecisionMinimum = c.PrecisionMinimum
	}
	if len(c.SeedRefractoryPeriod) > 0 {
		SeedRefractoryPeriod, _ = time.ParseDuration(c.SeedRefractoryPeriod)
	}
	if c.IncludeNilBits != nil {
		IncludeNilBits = *c.IncludeNilBits
	}
	if c.CompactVectors != nil {
		CompactVectors = *c.CompactVectors
	}
}
