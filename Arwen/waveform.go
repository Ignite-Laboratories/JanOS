package Arwen

import (
	"JanOS/Logic"
	"time"
)

type waveformSystem struct {
}

func NewWaveformSystem() *waveformSystem {
	return &waveformSystem{}
}

func (sys *waveformSystem) Initialize() {

}

func (sys *waveformSystem) Tick(entity Logic.Entity, delta time.Duration) {

}

func (sys *waveformSystem) GetName() string { return "Waveform System" }
