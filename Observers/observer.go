package Observers

import "github.com/ignite-laboratories/JanOS"

type Observer struct {
	Name          string
	OnObservation (JanOS.Sample)
}
