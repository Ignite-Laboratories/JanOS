package Systems

import (
	"fmt"
	"github.com/Ignite-Laboratories/JanOS/Logic"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type WavVizSystem struct {
	Logic.Entity

	components WavVizSystemComponents
}

type WavVizSystemComponents struct {
	Renderable *Logic.RenderableSet
	BinaryData *Logic.BinaryDataSet
}

func NewWaveformVisualizerSystem() WavVizSystem {
	return WavVizSystem{
		components: WavVizSystemComponents{
			Renderable: &Logic.RenderableSet{},
			BinaryData: &Logic.BinaryDataSet{},
		},
	}
}

func (sys WavVizSystem) GetName() string         { return "Waveform Visualizer" }
func (sys WavVizSystem) GetEntity() Logic.Entity { return sys.Entity }

func (sys WavVizSystem) Initialize(world *Logic.World) {

}

var gotSine bool

func (sys WavVizSystem) Tick(world *Logic.World, inbox Logic.Inbox) {
	var pass bool
	binaryData, ok := world.Assets.GetBinaryData("sine.1k")
	pass = ok

	if pass && !gotSine {
		log.Println("Got the sinewave")
		fmt.Println(string(binaryData.Data))
		gotSine = true
	}
}

func (sys WavVizSystem) Draw(img *ebiten.Image) {
}
