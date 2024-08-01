package core

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SystemIniter interface {
	Init(e Entity)
}

type SystemUpdater interface {
	Update(e Entity)
}

type SystemDrawer interface {
	Draw(e Entity, img *ebiten.Image)
}
