package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

func main() {
	// Initialize SDL2
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf("Failed to initialize SDL: %s\n", err)
	}
	defer sdl.Quit()

	// Create a full-screen window
	window, err := sdl.CreateWindow(
		"SDL2 Full-Screen Example",    // Window title
		sdl.WINDOWPOS_UNDEFINED,       // Window position X
		sdl.WINDOWPOS_UNDEFINED,       // Window position Y
		1920,                          // Window width (ignored in full-screen mode)
		1080,                          // Window height (ignored in full-screen mode)
		sdl.WINDOW_FULLSCREEN_DESKTOP, // Full-screen mode flag
	)
	if err != nil {
		log.Fatalf("Failed to create window: %s\n", err)
	}
	defer window.Destroy()
	running := true

	// Run the event loop to display the window and check for inputs
	for running {
		// Handle SDL events (e.g., window close or keyboard/mouse input)
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent: // Handle quit event
				running = false
			case *sdl.KeyboardEvent: // Handle keyboard events
				if e.Type == sdl.KEYDOWN { // Check for key press down
					switch e.Keysym.Sym {
					case sdl.K_ESCAPE: // If ESC key is pressed
						running = false
					default:
						fmt.Printf("Key pressed: %s\n", sdl.GetKeyName(e.Keysym.Sym))
					}
				}
			}
		}

		// Query and print the mouse position
		x, y, _ := sdl.GetGlobalMouseState()
		fmt.Printf("Mouse position: X = %d, Y = %d\n", x, y)

		// Limit the loop frequency a bit
		sdl.Delay(16) // ~60 fps
	}

	fmt.Println("Exit")
}
