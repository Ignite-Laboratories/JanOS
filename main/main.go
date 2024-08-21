package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ignite-laboratories/JanOS"
	"github.com/ignite-laboratories/JanOS/Formula"
	"log"
)

// The reference that ebiten drives
type window int

var signalA = JanOS.Signals.Create().SetValue(1).Toggle()
var signalB = JanOS.Signals.Create().SetValue(100)
var signalC = signalA.CreateMux(Formula.Additive, signalB)
var signalD = signalA.CreateMux(Formula.Multiplicative, signalB, signalC)
var signalE = signalA.CreateMux(Formula.Divisive, signalB, signalC)
var signalG = signalA.AbsoluteValue()
var signalF = JanOS.Signals.SineWave(signalB, signalG)

func main() {
	var w *window
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Spark")
	if err := ebiten.RunGame(w); err != nil {
		log.Panic(err)
	}
}

func (w *window) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		signalA.SetValue(signalA.Value * 1.1)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		signalA.SetValue(signalA.Value * 0.9)
	}

	return nil
}

func (w *window) Draw(screen *ebiten.Image) {
	// This is used in lieu of a newline in graphics-land
	offset := 15

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", signalA.Name, signalA.Value), 0, offset*0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", signalB.Name, signalB.Value), 0, offset*1)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", signalC.Name, signalC.Value), 0, offset*2)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", signalD.Name, signalD.Value), 0, offset*3)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", signalE.Name, signalE.Value), 0, offset*4)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", signalF.Name, signalF.Value), 0, offset*5)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", signalG.Name, signalG.Value), 0, offset*6)
}

// Layout is required by ebiten
func (w *window) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
