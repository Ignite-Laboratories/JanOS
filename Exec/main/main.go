package main

import (
	"JanOS"
	"JanOS/Arwen"
	"JanOS/Arwen/AI_Music"
	"JanOS/Logic"
	"JanOS/Logic/Symbol"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"time"
)

var waveformSys = Arwen.NewWaveformSystem()
var aiMusicSys = AI_Music.NewAI_MusicSystem()

var ecsWorld = Logic.NewECSWorld("Logic", waveformSys, aiMusicSys)
var window = JanOS.NewWindow("Spark", "Sample Window", 1024, 768, Update, Layout, OnDraw)

func main() {
	JanOS.Universe.Start(window, preflight, tick, ecsWorld)
}

var performance AI_Music.Performance
var binaryData AI_Music.BinaryData
var alpha *JanOS.Dimension
var omega *JanOS.Dimension
var theta *JanOS.Dimension
var gamma *JanOS.PerceptionBuffer

func preflight() {
	performance, _ = aiMusicSys.LookupPerformance(AI_Music.FamilyBrass, AI_Music.NameTrumpetInC, AI_Music.PitchA5, AI_Music.DynamicFortissimo)
	binaryData, _ = aiMusicSys.GetBinaryData(performance.Entity)
	alpha = JanOS.Universe.Dimensions.NewDimension("Alpha", Symbol.Alpha, 100)
	omega = JanOS.Universe.Dimensions.NewDimension("Omega", Symbol.Omega, 1)
	theta = JanOS.Universe.Dimensions.NewOscillatingDimension("Theta", Symbol.Theta, alpha, omega)
	gamma = JanOS.Universe.Dimensions.NewPerceptionBuffer("Gamma", Symbol.Gamma, time.Minute*10, time.Second)
}

func tick(delta time.Duration) {

}

func onDraw(screen *ebiten.Image) {

}

func Update(window *JanOS.Window) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		omega.Value = omega.Value * 0.9
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		omega.Value = omega.Value * 1.1
	}
	return nil
}

func OnDraw(window *JanOS.Window, screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", theta.Name, theta.Value), 0, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", alpha.Name, alpha.Value), 0, 15)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", omega.Name, omega.Value), 0, 30)
}

func Layout(window *JanOS.Window, outsideWidth, outsideHeight int) (int, int) {
	return window.ScreenWidth, window.ScreenHeight
}
