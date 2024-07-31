package main

import (
	"fmt"
	"github.com/Ignite-Laboratories/JanOS/entities"
	"github.com/Ignite-Laboratories/JanOS/systems"
	"github.com/go-audio/wav"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
	"os"
)

const (
	tileSize     = 64
	screenWidth  = tileSize * 30
	screenHeight = tileSize * 20
)

type World struct {
	Entities   []entities.Entity
	Components struct {
	}
	Systems []any
}

func main() {
	f, err := os.Open("c:\\Source\\Arwen-Source\\JanOS\\cmd\\waveform\\assets\\1k sine.wav")
	if err != nil {
		fmt.Println(err)
		return
	}

	d := wav.NewDecoder(f)

	pcmBuffer, err := d.FullPCMBuffer()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println(pcmBuffer.Data)

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	world := &World{}

	// Setup the Systems

	inputSystem := &systems.InputSystem{}

	world.Systems = []any{
		inputSystem,
	}

	world.Init()

	ebiten.SetWindowSize(screenWidth*4, screenHeight*4)
	ebiten.SetWindowTitle("JanOS")
	if err := ebiten.RunGame(world); err != nil {
		log.Fatal(err)
	}
}

func (g *World) Init() {
	for _, e := range g.Entities {
		for _, s := range g.Systems {
			if d, ok := s.(systems.SystemIniter); ok {
				d.Init(e)
			}
		}
	}
}

func (g *World) Update() error {
	for _, e := range g.Entities {
		for _, s := range g.Systems {
			if u, ok := s.(systems.SystemUpdater); ok {
				u.Update(e)
			}
		}
	}
	return nil
}

func (g *World) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	for _, e := range g.Entities {
		for _, s := range g.Systems {
			if d, ok := s.(systems.SystemDrawer); ok {
				d.Draw(e, screen)
			}
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Ticks/sec: %0.2f", ebiten.ActualTPS()))
}

func (g *World) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
