package systems

import (
	"github.com/Ignite-Laboratories/JanOS/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

type SystemIniter interface {
	Init(e entities.Entity)
}

type SystemUpdater interface {
	Update(e entities.Entity)
}

type SystemDrawer interface {
	Draw(e entities.Entity, img *ebiten.Image)
}
