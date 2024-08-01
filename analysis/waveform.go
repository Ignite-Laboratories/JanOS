package analysis

import (
	"fmt"
	"github.com/Ignite-Laboratories/JanOS/core"
	"github.com/go-audio/wav"
	"github.com/hajimehoshi/ebiten/v2"
	"os"
)

type WaveformSystem struct {
	WorkEntity core.Entity
}

type WaveformComponents struct {
	core.Components[Waveform]
}

type Waveform struct {
	FilePath           string
	RawData            []int
	DiscreteDerivative []int
	InflectionPoints   []*TimeRecord
}

func (x *WaveformSystem) Init(e core.Entity) {
	if e != x.WorkEntity {
		return
	}

	var waveform, _ = HomeWorld.Components.Waveforms.Get(x.WorkEntity)

	f, err := os.Open(waveform.FilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	d := wav.NewDecoder(f)

	pcmBuffer, err := d.FullPCMBuffer()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	waveform.RawData = pcmBuffer.Data
	waveform.ProcessData()
}

func (x *WaveformSystem) Draw(e core.Entity, img *ebiten.Image) {}

func (x *WaveformSystem) Update(e core.Entity) {

}

func (w *Waveform) ProcessData() {
	w.processDeltas()

	var lastPoint int
	var lastInflectionPoint int
	for i, currentPoint := range w.DiscreteDerivative {
		if i == 0 {
			lastPoint = currentPoint
			lastInflectionPoint = 0
			continue
		}

		if core.DiffSigns(lastPoint, currentPoint) {
			w.InflectionPoints = append(w.InflectionPoints, NewTimeRecord(i, w.DiscreteDerivative[lastInflectionPoint:i]))
			lastInflectionPoint = i
		}
		lastPoint = currentPoint
	}

	fmt.Println(w.InflectionPoints)
	fmt.Println(w.DiscreteDerivative[:100])
}

func (w *Waveform) processDeltas() {
	var lastPoint int
	for i, currentPoint := range w.RawData {
		// Skip the first time step
		if i == 0 {
			lastPoint = currentPoint
			w.DiscreteDerivative = append(w.DiscreteDerivative, 0)
			continue
		}

		// Store the delta between points
		delta := currentPoint - lastPoint
		w.DiscreteDerivative = append(w.DiscreteDerivative, delta)
		lastPoint = currentPoint
	}
}
