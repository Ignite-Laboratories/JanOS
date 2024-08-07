package Systems

import (
	"github.com/Ignite-Laboratories/JanOS/Spark"
	"github.com/Ignite-Laboratories/JanOS/Spark/Util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
)

type WavVizSystem struct {
	Spark.Entity

	components WavVizSystemComponents
}

type WavVizSystemComponents struct {
	Renderable *Spark.RenderableSet
	BinaryData *Spark.BinaryDataSet
	Waveforms  *WaveformSet
}

func NewWaveformVisualizerSystem() WavVizSystem {
	return WavVizSystem{
		components: WavVizSystemComponents{
			Renderable: &Spark.RenderableSet{},
			BinaryData: &Spark.BinaryDataSet{},
			Waveforms:  &WaveformSet{},
		},
	}
}

type Waveform struct {
	Entity Spark.Entity
	Cursor Spark.Entity
	Values []int
}

type WaveformSet struct {
	Spark.Components[Waveform]
}

func (sys WavVizSystem) GetName() string         { return "Waveform Visualizer" }
func (sys WavVizSystem) GetEntity() Spark.Entity { return sys.Entity }

func (sys WavVizSystem) Initialize() {

}

var gotSine bool

func (sys WavVizSystem) Visualize(cursor Spark.Entity) {
	waveform := Waveform{
		Entity: Spark.Universe.CreateEntity(),
		Cursor: cursor,
	}
	sys.components.Waveforms.Set(waveform.Entity, waveform)
}

func (sys WavVizSystem) Tick(inbox Spark.Inbox) {
}

func (sys WavVizSystem) OnDraw(entity Spark.Entity, screen *ebiten.Image) {
	waveform, ok := sys.components.Waveforms.Get(entity)
	cursorSystem, _ := Spark.Universe.GetSystem(CursoringSystem{}).(CursoringSystem)
	buffer := cursorSystem.GetCursorBuffer(waveform.Cursor)
	if ok {
		var path vector.Path
		var verticalCenter = float32(screen.Bounds().Max.Y) / 2
		var xMax = float32(screen.Bounds().Max.X)
		var xSpacing = xMax / float32(len(buffer))
		var yMax = float32(250)
		var yScaleFactor = yMax / float32(Util.GetLargest(buffer))
		path.MoveTo(0, verticalCenter)
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
