package main

import (
	"github.com/Ignite-Laboratories/JanOS/analysis"
	"github.com/Ignite-Laboratories/JanOS/core"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	w := &analysis.World{
		Name:     "JanOS",
		Width:    640,
		Height:   480,
		Entities: []core.Entity{},
		Components: struct {
			Waveforms analysis.WaveformComponents
		}{},
		Systems: []any{
			&analysis.WaveformSystem{},
		},
	}

	workEntity := core.NewEntity()
	w.Components.Waveforms.Set(workEntity, analysis.Waveform{FilePath: "c:\\source\\ignite\\janos\\analysis\\assets\\1k sine.wav"})
	w.AddEntity(workEntity)

	analysis.HomeWorld = w
	analysis.HomeWorld.FireBigBang()
}
