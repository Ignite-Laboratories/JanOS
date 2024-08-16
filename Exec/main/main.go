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
var window = JanOS.NewWindow("Spark", "JanOS", 1024, 768, Update, Layout, OnDraw)

func main() {
	JanOS.Universe.Start(nil, preflight, tick, ecsWorld)
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
	alpha = JanOS.Universe.Dimensions.NewDimension("Alpha", Symbol.Alpha, 100)
	omega = JanOS.Universe.Dimensions.NewDimension("Omega", Symbol.Omega, 1)
	mu = JanOS.Universe.Dimensions.NewDimension("Mu", Symbol.Mu, 1.1)
	nu = JanOS.Universe.Dimensions.NewDimension("Nu", Symbol.Nu, 1.2)
	xi = JanOS.Universe.Dimensions.NewDimension("Xi", Symbol.Xi, 1.3)
	omicron = JanOS.Universe.Dimensions.NewDimension("Omicron", Symbol.Omicron, 1.4)
	pi = JanOS.Universe.Dimensions.NewDimension("Pi", Symbol.Pi, 1.5)
	rho = JanOS.Universe.Dimensions.NewDimension("Rho", Symbol.Rho, 1.6)

	theta = JanOS.Universe.Dimensions.NewOscillatingDimension("Theta", Symbol.Theta, alpha, omega)
	sigma = JanOS.Universe.Dimensions.NewOscillatingDimension("Sigma", Symbol.Sigma, alpha, mu)
	tau = JanOS.Universe.Dimensions.NewOscillatingDimension("Tau", Symbol.Tau, alpha, nu)
	upsilon = JanOS.Universe.Dimensions.NewOscillatingDimension("Upsilon", Symbol.Upsilon, alpha, xi)
	phi = JanOS.Universe.Dimensions.NewOscillatingDimension("Phi", Symbol.Phi, alpha, omicron)
	chi = JanOS.Universe.Dimensions.NewOscillatingDimension("Chi", Symbol.Chi, alpha, pi)
	psi = JanOS.Universe.Dimensions.NewOscillatingDimension("Psi", Symbol.Psi, alpha, rho)
}

func tick(delta time.Duration) {
}

func onDraw(screen *ebiten.Image) {

}

func Update(window *JanOS.Window) error {
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

func OnDraw(window *JanOS.Window, screen *ebiten.Image) {
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

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", theta.Name, theta.Timeline.SliceFuture(now, JanOS.Universe.Resolution.Duration*10).Data), 0, 9*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", sigma.Name, sigma.Timeline.SliceFuture(now, JanOS.Universe.Resolution.Duration*10).Data), 0, 10*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", tau.Name, tau.Timeline.SliceFuture(now, JanOS.Universe.Resolution.Duration*10).Data), 0, 11*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", upsilon.Name, upsilon.Timeline.SliceFuture(now, JanOS.Universe.Resolution.Duration*10).Data), 0, 12*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", phi.Name, phi.Timeline.SliceFuture(now, JanOS.Universe.Resolution.Duration*10).Data), 0, 13*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", chi.Name, chi.Timeline.SliceFuture(now, JanOS.Universe.Resolution.Duration*10).Data), 0, 14*offset)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s - %f", psi.Name, psi.Timeline.SliceFuture(now, JanOS.Universe.Resolution.Duration*10).Data), 0, 15*offset)
}

func Layout(window *JanOS.Window, outsideWidth, outsideHeight int) (int, int) {
	return window.ScreenWidth, window.ScreenHeight
}
