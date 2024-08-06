package OLD

import (
	"fmt"
	"github.com/Ignite-Laboratories/JanOS/Logic/Util"
)

type Waveform struct {
	//ecs.BasicEntity
	RawData            []int
	DiscreteDerivative []int
	InflectionPoints   []*Util.TimeRecord
}

var waveform *Waveform = &Waveform{}

type WaveformSystem struct {
	//world       *ecs.World
	//assetSystem *AssetSystem
}

func (ws *WaveformSystem) New(w int) { // *ecs.World) {
	//ws.world = w
	//
	//for _, system := range w.Systems() {
	//	switch sys := system.(type) {
	//	case *AssetSystem:
	//		ws.assetSystem = sys
	//	}
	//}
}

func (*WaveformSystem) Remove() {} //ecs.BasicEntity) {}

func (ws *WaveformSystem) Update(dt float32) {
	//if engo.Input.Button("LoadFile").JustPressed() {
	//	var msg = NewLoadAssetMsg()
	//	msg.Load("asdf1", "audio\\sine.1k.wav")
	//	msg.Load("asdf2", "audio\\sine.1k.wav")
	//	msg.Load("asdf3", "audio\\sine.1k.wav")
	//	msg.Load("asdf4", "audio\\sine.1k.wav")
	//	msg.Load("asdf5", "audio\\sine.1k.wav")
	//	engo.Mailbox.Dispatch(msg)
	//}
	//if engo.Input.Button("Analyze").JustPressed() {
	//	asset := ws.assetSystem.assets["sine.1k"]
	//	readSeeker := bytes.NewReader(asset.FileMetaData.Contents)
	//	d := wav.NewDecoder(readSeeker)
	//
	//	pcmBuffer, err := d.FullPCMBuffer()
	//	if err != nil {
	//		fmt.Fprintln(os.Stderr, err)
	//		return
	//	}
	//
	//	waveform.RawData = pcmBuffer.Data
	//	waveform.ProcessData()
	//}
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

		if Util.DiffSigns(lastPoint, currentPoint) {
			w.InflectionPoints = append(w.InflectionPoints, Util.NewTimeRecord(i, w.DiscreteDerivative[lastInflectionPoint:i]))
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
