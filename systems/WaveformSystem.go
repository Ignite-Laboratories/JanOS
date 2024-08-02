package systems

import (
	"fmt"
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/Ignite-Laboratories/JanOS/util"
	"github.com/go-audio/wav"
	"os"
)

type Waveform struct {
	ecs.BasicEntity
	RawData            []int
	DiscreteDerivative []int
	InflectionPoints   []*util.TimeRecord
}

var waveform *Waveform = &Waveform{}

type WaveformSystem struct {
	world       *ecs.World
	assetSystem *AssetSystem
}

func (ws *WaveformSystem) New(w *ecs.World) {
	ws.world = w

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *AssetSystem:
			ws.assetSystem = sys
		}
	}
}

func (*WaveformSystem) Remove(ecs.BasicEntity) {}

func (ws *WaveformSystem) Update(dt float32) {
	if engo.Input.Button("Analyze").JustPressed() {
		d := wav.NewDecoder(ws.assetSystem.files["sine.1k"])

		pcmBuffer, err := d.FullPCMBuffer()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		waveform.RawData = pcmBuffer.Data
		waveform.ProcessData()
	}
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

		if util.DiffSigns(lastPoint, currentPoint) {
			w.InflectionPoints = append(w.InflectionPoints, util.NewTimeRecord(i, w.DiscreteDerivative[lastInflectionPoint:i]))
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
