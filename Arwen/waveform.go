package Arwen

import (
	"JanOS"
	"time"
)

type waveformSystem struct {
}

func NewWaveformSystem() *waveformSystem {
	return &waveformSystem{}
}

//func (sys *waveformSystem) Initialize() {
//
//}

func (sys *waveformSystem) Tick(entity JanOS.Entity, delta time.Duration) {

}

func (sys *waveformSystem) GetNamedValue() string { return "Waveform System" }
