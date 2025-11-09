package atlas

import (
	"time"

	"git.ignitelabs.net/janos/core/sys/rec"
)

type config struct {
	PrintPreamble        *bool         `json:"printPreamble"`
	Verbose              *bool         `json:"verbose"`
	Silent               *bool         `json:"silent"`
	ShutdownTimeout      time.Duration `json:"shutdownTimeout"`
	Record               []byte        `json:"record"`
	ObservanceWindow     string        `json:"observanceWindow"`
	ObservedMinimum      uint          `json:"observedMinimum"`
	TrimFrequency        float64       `json:"trimFrequency"`
	Precision            uint          `json:"precision"`
	PrecisionMinimum     uint          `json:"precisionMinimum"`
	Radix                uint          `json:"radix"`
	SeedRefractoryPeriod string        `json:"seedRefractoryPeriod"`
	IncludeNilBits       *bool         `json:"includeNilBits"`
	CompactVectors       *bool         `json:"compactVectors"`
	SynapticChannelLimit uint          `json:"synapticChannelLimit"`
}

func (c config) apply() {
	if c.PrintPreamble != nil {
		PrintPreamble = *c.PrintPreamble
	}
	if c.Verbose != nil {
		rec.Verbose = *c.Verbose
	}
	if c.Silent != nil {
		rec.Silent = *c.Silent
	}
	if c.ShutdownTimeout != 0 {
		ShutdownTimeout = c.ShutdownTimeout
	}
	if len(c.Record) > 0 {
		Record = c.Record
	}
	if len(c.ObservanceWindow) > 0 {
		ObservanceWindow, _ = time.ParseDuration(c.ObservanceWindow)
	}
	if c.ObservedMinimum != 0 {
		ObservedMinimum = c.ObservedMinimum
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
	if c.Radix > 0 {
		Radix = c.Radix
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
	if c.SynapticChannelLimit > 0 {
		SynapticChannelLimit = c.SynapticChannelLimit
	}
}
