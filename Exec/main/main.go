package main

import (
	"JanOS"
	"JanOS/Arwen"
	"JanOS/Arwen/AI_Music"
	"JanOS/Symbol"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"time"
)

var waveformSys = Arwen.NewWaveformSystem()
var aiMusicSys = AI_Music.NewAI_MusicSystem()

var ecsWorld = JanOS.NewECSWorld("Logic", waveformSys, aiMusicSys)
var window = JanOS.NewWindow("Spark", "JanOS", 1024, 768, onTick, onDraw)

func main() {
	JanOS.Universe.Start(window, preflight, onRealityUpdate, ecsWorld)
}

var performance AI_Music.Performance
var binaryData AI_Music.BinaryData
var alpha *JanOS.Dimension
var omega *JanOS.Dimension
var theta *JanOS.Dimension
var sigma *JanOS.Dimension
var tau *JanOS.Dimension
var upsilon *JanOS.Dimension
var phi *JanOS.Dimension
var chi *JanOS.Dimension
var psi *JanOS.Dimension
var mu *JanOS.Dimension
var nu *JanOS.Dimension
var xi *JanOS.Dimension
var omicron *JanOS.Dimension
var pi *JanOS.Dimension
var rho *JanOS.Dimension

func preflight() {
	performance, _ = aiMusicSys.LookupPerformance(AI_Music.FamilyBrass, AI_Music.NameTrumpetInC, AI_Music.PitchA5, AI_Music.DynamicFortissimo)
	binaryData, _ = aiMusicSys.GetBinaryData(performance.Entity)

	bufferLength := time.Duration(time.Second * 5)
	frequency := 44000

	alpha = JanOS.Universe.Dimensions.NewDimension("Alpha", Symbol.Alpha, 100, bufferLength, frequency)
	omega = JanOS.Universe.Dimensions.NewDimension("Omega", Symbol.Omega, 1, bufferLength, frequency)
	mu = JanOS.Universe.Dimensions.NewDimension("Mu", Symbol.Mu, 1.1, bufferLength, frequency)
	nu = JanOS.Universe.Dimensions.NewDimension("Nu", Symbol.Nu, 1.2, bufferLength, frequency)
	xi = JanOS.Universe.Dimensions.NewDimension("Xi", Symbol.Xi, 1.3, bufferLength, frequency)
	omicron = JanOS.Universe.Dimensions.NewDimension("Omicron", Symbol.Omicron, 1.4, bufferLength, frequency)
	pi = JanOS.Universe.Dimensions.NewDimension("Pi", Symbol.Pi, 1.5, bufferLength, frequency)
	rho = JanOS.Universe.Dimensions.NewDimension("Rho", Symbol.Rho, 1.6, bufferLength, frequency)

	theta = JanOS.Universe.Dimensions.NewOscillatingDimension("Theta", Symbol.Theta, alpha, omega, bufferLength, frequency)
	sigma = JanOS.Universe.Dimensions.NewOscillatingDimension("Sigma", Symbol.Sigma, alpha, mu, bufferLength, frequency)
	tau = JanOS.Universe.Dimensions.NewOscillatingDimension("Tau", Symbol.Tau, alpha, nu, bufferLength, frequency)
	upsilon = JanOS.Universe.Dimensions.NewOscillatingDimension("Upsilon", Symbol.Upsilon, alpha, xi, bufferLength, frequency)
	phi = JanOS.Universe.Dimensions.NewOscillatingDimension("Phi", Symbol.Phi, alpha, omicron, bufferLength, frequency)
	chi = JanOS.Universe.Dimensions.NewOscillatingDimension("Chi", Symbol.Chi, alpha, pi, bufferLength, frequency)
	psi = JanOS.Universe.Dimensions.NewOscillatingDimension("Psi", Symbol.Psi, alpha, rho, bufferLength, frequency)
}

func onRealityUpdate(delta time.Duration) {
}

func onTick(window *JanOS.Window) error {
	now := time.Now()
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		omega.SetValue(now, omega.GetValue(now)*0.9)
		mu.SetValue(now, mu.GetValue(now)*0.9)
		nu.SetValue(now, nu.GetValue(now)*0.9)
		xi.SetValue(now, xi.GetValue(now)*0.9)
		omicron.SetValue(now, omicron.GetValue(now)*0.9)
		pi.SetValue(now, pi.GetValue(now)*0.9)
		rho.SetValue(now, rho.GetValue(now)*0.9)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		omega.SetValue(now, omega.GetValue(now)*1.1)
		mu.SetValue(now, mu.GetValue(now)*1.1)
		nu.SetValue(now, nu.GetValue(now)*1.1)
		xi.SetValue(now, xi.GetValue(now)*1.1)
		omicron.SetValue(now, omicron.GetValue(now)*1.1)
		pi.SetValue(now, pi.GetValue(now)*1.1)
		rho.SetValue(now, rho.GetValue(now)*1.1)
	}

	return nil
}

func onDraw(window *JanOS.Window, screen *ebiten.Image) {
	now := time.Now()
	offset := 15
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", alpha.Name, alpha.GetValue(now)), 0, 0*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", omega.Name, omega.GetValue(now)), 0, 1*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", theta.Name, theta.GetValue(now)), 0, 2*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", mu.Name, mu.GetValue(now)), 0, 3*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", nu.Name, nu.GetValue(now)), 0, 4*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", xi.Name, xi.GetValue(now)), 0, 5*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", omicron.Name, omicron.GetValue(now)), 0, 6*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", pi.Name, pi.GetValue(now)), 0, 7*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", rho.Name, rho.GetValue(now)), 0, 8*offset)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", theta.Name, theta.Timeline.SliceFuture(now, theta.Timeline.Resolution.Duration*10).Data), 0, 9*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", sigma.Name, sigma.Timeline.SliceFuture(now, sigma.Timeline.Resolution.Duration*10).Data), 0, 10*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", tau.Name, tau.Timeline.SliceFuture(now, tau.Timeline.Resolution.Duration*10).Data), 0, 11*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", upsilon.Name, upsilon.Timeline.SliceFuture(now, upsilon.Timeline.Resolution.Duration*10).Data), 0, 12*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", phi.Name, phi.Timeline.SliceFuture(now, phi.Timeline.Resolution.Duration*10).Data), 0, 13*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", chi.Name, chi.Timeline.SliceFuture(now, chi.Timeline.Resolution.Duration*10).Data), 0, 14*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", psi.Name, psi.Timeline.SliceFuture(now, psi.Timeline.Resolution.Duration*10).Data), 0, 15*offset)
}
