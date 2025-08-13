package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std/bounded"
	"github.com/ignite-laboratories/core/std/pattern"
)

type asdf struct {
}

func main() {
	bounded.
		pad.Plane()

	p := pattern.NilAny()

	for core.Alive() {
		if err := keyboard.Open(); err != nil {
			panic(err)
		}

		fmt.Println("Press Left/Right arrows (ESC to exit)")

		for {
			_, key, err := keyboard.GetKey()
			if err != nil {
				fmt.Println("error:", err)
				return
			}
			switch key {
			case keyboard.KeyArrowLeft:
				v := p.GoWest()
				fmt.Println(v)
			case keyboard.KeyArrowRight:
				v := p.GoEast()
				fmt.Println(v)
			case keyboard.KeyEsc:
				keyboard.Close()
				core.ShutdownNow()
				return
			default:
			}
		}
	}
}
