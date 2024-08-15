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
var omega *JanOS.Dimension
var theta *JanOS.Dimension

func preflight() {
	performance, _ = aiMusicSys.LookupPerformance(AI_Music.FamilyBrass, AI_Music.NameTrumpetInC, AI_Music.PitchA5, AI_Music.DynamicFortissimo)
	binaryData, _ = aiMusicSys.GetBinaryData(performance.Entity)
	omega = JanOS.Universe.Dimensions.NewDimension("Omega", Symbol.Omega, 42)
	theta = JanOS.Universe.Dimensions.NewOscillatingDimension("Theta", Symbol.Theta, 100, 1, 1000)
}

func tick(delta time.Duration) {

}

func onDraw(screen *ebiten.Image) {

}

func Update(window *JanOS.Window) error {
	return nil
}

func OnDraw(window *JanOS.Window, screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%s - %f", theta.Name, theta.Value))
}

func Layout(window *JanOS.Window, outsideWidth, outsideHeight int) (int, int) {
	return window.ScreenWidth, window.ScreenHeight
}
