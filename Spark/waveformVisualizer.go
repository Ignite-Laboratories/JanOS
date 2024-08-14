package Spark

import (
	"github.com/Ignite-Laboratories/JanOS/Logic"
	"github.com/Ignite-Laboratories/JanOS/Logic/Math"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
)

type WavVizSystem struct {
	Logic.Entity

	components WavVizSystemComponents
}

type WavVizSystemComponents struct {
	Renderable *RenderableSet
	BinaryData *BinaryDataSet
	Waveforms  *WaveformSet
}

func NewWaveformVisualizerSystem() WavVizSystem {
	return WavVizSystem{
		components: WavVizSystemComponents{
			Renderable: &RenderableSet{},
			BinaryData: &BinaryDataSet{},
			Waveforms:  &WaveformSet{},
		},
	}
}

type Waveform struct {
	Entity Logic.Entity
	Cursor Logic.Entity
	Values []int
}

type WaveformSet struct {
	Logic.Components[Waveform]
}

func (sys WavVizSystem) GetName() string         { return "Waveform Visualizer" }
func (sys WavVizSystem) GetEntity() Logic.Entity { return sys.Entity }

func (sys WavVizSystem) Initialize() {

}

var gotSine bool

func (sys WavVizSystem) Visualize(cursor Logic.Entity) {
	waveform := Waveform{
		Entity: Universe.CreateEntity(),
		Cursor: cursor,
	}
	sys.components.Waveforms.Set(waveform.Entity, waveform)
}

func (sys WavVizSystem) Tick(inbox Inbox) {
}

func (sys WavVizSystem) OnDraw(entity Logic.Entity, screen *ebiten.Image) {
	waveform, ok := sys.components.Waveforms.Get(entity)
	buffer := Universe.Cursoring.GetCursorBuffer(waveform.Cursor)
	if ok {
		var path vector.Path
		var verticalCenter = float32(screen.Bounds().Max.Y) / 2
		var xMax = float32(screen.Bounds().Max.X)
		var xSpacing = xMax / float32(len(buffer))
		var yMax = float32(250)
		var yScaleFactor = yMax / float32(Math.GetLargest(buffer))
		path.MoveTo(0, verticalCenter+(float32(buffer[0])*yScaleFactor))
		for i, value := range buffer {
			x := float32(i) * xSpacing
			y := verticalCenter + (float32(value) * yScaleFactor)
			path.LineTo(x, y)
		}

		var vs []ebiten.Vertex
		var is []uint16
		op := &vector.StrokeOptions{}
		op.Width = 5
		op.LineJoin = vector.LineJoinRound
		vs, is = path.AppendVerticesAndIndicesForStroke(nil, nil, op)

		for i := range vs {
			vs[i].SrcX = 1
			vs[i].SrcY = 1
			vs[i].ColorR = 0x33 / float32(0xff)
			vs[i].ColorG = 0x66 / float32(0xff)
			vs[i].ColorB = 0xff / float32(0xff)
			vs[i].ColorA = 1
		}

		top := &ebiten.DrawTrianglesOptions{}
		top.AntiAlias = true
		top.FillRule = ebiten.NonZero
		whiteImage := ebiten.NewImage(3, 3)
		whiteSubImage := whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
		whiteImage.Fill(color.White)
		screen.DrawTriangles(vs, is, whiteSubImage, top)
	}
}
