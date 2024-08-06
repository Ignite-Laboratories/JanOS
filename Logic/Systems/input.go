package Systems

import (
	"github.com/Ignite-Laboratories/JanOS/Logic"
)

type InputSystem struct {
	Logic.Entity

	components InputSystemComponents
}

type InputSystemComponents struct {
}

func NewInputSystem() InputSystem {
	return InputSystem{
		components: InputSystemComponents{},
	}
}

func (sys InputSystem) GetName() string         { return "Input" }
func (sys InputSystem) GetEntity() Logic.Entity { return sys.Entity }

func (sys InputSystem) Initialize(world *Logic.World) {
}

func (sys InputSystem) Tick(world *Logic.World, inbox Logic.Inbox) {
}
