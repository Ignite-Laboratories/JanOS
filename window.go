package JanOS

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Window struct {
	Name         string
	WindowTitle  string
	ScreenWidth  int
	ScreenHeight int
	OnUpdate     func(window *Window) error
	OnLayout     func(window *Window, outsideWidth, outsideHeight int) (int, int)
	OnDraw       func(window *Window, screen *ebiten.Image)
}

func NewWindow(name string, windowTitle string, width int, height int, onUpdate func(window *Window) error, onLayout func(window *Window, outsideWidth, outsideHeight int) (int, int), onDraw func(window *Window, screen *ebiten.Image)) *Window {
	return &Window{
		Name:         name,
		WindowTitle:  windowTitle,
		ScreenWidth:  width,
		ScreenHeight: height,
		OnUpdate:     onUpdate,
		OnLayout:     onLayout,
		OnDraw:       onDraw,
	}
}

func (w *Window) GetName() string {
	return w.Name
}

func (w *Window) Open() {
	Universe.Printf(w, "Launching ebiten window '%s' (%d, %d)", w.WindowTitle, w.ScreenWidth, w.ScreenHeight)
	ebiten.SetWindowSize(w.ScreenWidth, w.ScreenHeight)
	ebiten.SetWindowTitle(w.WindowTitle)
	if err := ebiten.RunGame(w); err != nil {
		log.Panic(err)
	}
	Universe.Printf(w, "ebiten window '%s' closed", w.WindowTitle)
}

func (w *Window) Update() error {
	if Universe.Terminate {
		return ebiten.Termination
	}
	return w.OnUpdate(w)
}

func (w *Window) Draw(screen *ebiten.Image) {
	w.OnDraw(w, screen)
}

func (w *Window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return w.OnLayout(w, outsideWidth, outsideHeight)
}
