package JanOS

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

// Window represents an abstract graphical window.
// OnTick is called by ebiten on a fixed interval.
// OnDraw is called by ebiten whenever it is ready to draw a frame.
// The reality loop is not managed by the window.
type Window struct {
	Name         string
	WindowTitle  string
	ScreenWidth  int
	ScreenHeight int
	OnTick       func(window *Window) error
	OnDraw       func(window *Window, screen *ebiten.Image)
	OnLayout     func(outsideWidth, outsideHeight int) (int, int)
}

// NewWindow creates a window and connects up the provided tick and draw functions.
// onTick is called by ebiten's Update method, and onDraw is called by its OnDraw method.
// If you would like to override ebiten's Layout method, re-assign the OnLayout property.
func NewWindow(name string, windowTitle string, width int, height int, onTick func(window *Window) error, onDraw func(window *Window, screen *ebiten.Image)) *Window {
	return &Window{
		Name:         name,
		WindowTitle:  windowTitle,
		ScreenWidth:  width,
		ScreenHeight: height,
		OnTick:       onTick,
		OnDraw:       onDraw,
		OnLayout:     func(outsideWidth, outsideHeight int) (int, int) { return outsideWidth, outsideHeight },
	}
}

// GetNamedValue returns the assigned name to this instance.
func (w *Window) GetNamedValue() string {
	return w.Name
}

// Open opens the window in the operating system.
func (w *Window) Open() {
	Universe.Printf(w, "Launching ebiten window '%s' (%d, %d)", w.WindowTitle, w.ScreenWidth, w.ScreenHeight)
	ebiten.SetWindowSize(w.ScreenWidth, w.ScreenHeight)
	ebiten.SetWindowTitle(w.WindowTitle)
	if err := ebiten.RunGame(w); err != nil {
		log.Panic(err)
	}
	Universe.Printf(w, "ebiten window '%s' closed", w.WindowTitle)
}

// Update passes through ebiten's Update function to the provided OnTick function
func (w *Window) Update() error {
	if Universe.Terminate {
		return ebiten.Termination
	}
	return w.OnTick(w)
}

// Draw passes through ebiten's OnDraw function to the provided OnDraw function
func (w *Window) Draw(screen *ebiten.Image) {
	w.OnDraw(w, screen)
}

// Layout passes through ebiten's Layout function to the provided OnLayout function
func (w *Window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return w.OnLayout(outsideWidth, outsideHeight)
}
