package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

func main() {
	fmt.Println("SDL2 Mouse Position Example")

	// Initialize SDL2
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(fmt.Errorf("failed to initialize SDL: %v", err))
	}
	defer sdl.Quit()

	// Infinite loop to show global mouse position
	running := true
	for running {
		// Event polling to make sure SDL updates its internal state
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch ev := event.(type) {
			case *sdl.QuitEvent:
				// Exit loop on quit event
				fmt.Println("Quit event received!")
				running = false
			default:
				// Polling other events, if necessary (none in this example)
				_ = ev
			}
		}

		// Get global mouse state
		x, y, _ := sdl.GetGlobalMouseState()
		fmt.Printf("Global Mouse Position: X=%d, Y=%d\n", x, y)

		// Sleep briefly to not flood output
		time.Sleep(500 * time.Millisecond)
	}
}
